package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/turingdance/infra/cond"
	"turingdance.com/reliable/internal/app/sys"
	"turingdance.com/reliable/internal/app/sys/logic"
	"turingdance.com/reliable/internal/app/sys/model"
	"turingdance.com/reliable/internal/initial/conf"
	"turingdance.com/reliable/internal/pkg/log"
)

// 子命令定义 运行方法 go run main.go version 编译后 ./hugo version
var menuCmd = &cobra.Command{
	Use:   "menu", // Use这里定义的就是命令的名称
	Short: "generate menu by annotation",
	Long:  `generate menu by annotation`,
	Run: func(cmd *cobra.Command, args []string) { //这里是命令的执行方法
		//首先扫描目标文件夹,获得route数组
		routers, err := parseroute(appsrc)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tree := buildroutetree(routers)
		//然后链接数据库进行比对 创建
		file := fmt.Sprintf("app-%s.yml", env)
		conf.AppConf = conf.ParseConf(file)
		//初始化日志
		log.Initialize(conf.AppConf.Log)
		sys.Initialize(conf.AppConf)
		//

		// 创建1级菜单
		app := &model.Rights{
			Biz: appname,
		}
		condwraper := cond.NewCondWrapper()
		condwraper.AddCond(cond.Cond{
			Field: "biz", Op: "eq", Value: appname,
		}, cond.Cond{
			Field: "type", Op: "eq", Value: model.RightsTypeGroup,
		})
		app, err = logic.Take(app, *condwraper)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if app == nil || app.Id == 0 {
			//创建一个
			app.Pid = 0
			app.SortIndex = 0
			app.Title = apptitle
			app.Type = model.RightsTypeGroup
			app.Uri = "/" + appname
			app, err = logic.Create(app)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		//创建页面级菜单
		cond2 := cond.NewCondWrapper()
		cond2.AddOneCond("pid", "eq", app.Id)
		pages, _, err := logic.Search(&model.Rights{}, cond2)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		rigmap1 := map[string]*model.Rights{}
		for _, v := range pages {
			rigmap1[v.Biz] = &v
		}
		tmppages := make([]*model.Rights, 0)
		for _, page := range tree {
			biz := strings.ToLower(app.Biz + ":" + page.Module)
			if _, ok := rigmap1[biz]; !ok {
				//创建
				rightpage := &model.Rights{
					Biz:   biz,
					Title: page.Comment,
					Pid:   app.Id,
					Uri:   strings.ToLower(page.Module),
					Type:  model.RightsTypeView,
				}
				tmppages = append(tmppages, rightpage)
				rigmap1[biz] = rightpage
			}
		}
		for _, page := range tmppages {
			_, err = logic.Create(page)
			// 创建页面
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		//再次构建标签
		tmpbtns := make([]*model.Rights, 0)
		for _, page := range tree {
			for _, btn := range page.Children {
				biz := strings.ToLower(app.Biz + ":" + page.Module + ":" + btn.Func)
				pright := rigmap1[strings.ToLower(app.Biz+":"+page.Module)]
				rightbtn := &model.Rights{
					Biz:   biz,
					Title: btn.Comment,
					Pid:   pright.Id,
					Uri:   btn.Path,
					Type:  model.RightsTypeApi,
				}
				tmpbtns = append(tmpbtns, rightbtn)
			}
		}
		for _, btn := range tmpbtns {
			_, err = logic.Create(btn)
			// 创建页面
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		fmt.Println("全部处理完成")

	},
	PreRun: func(cmd *cobra.Command, args []string) {
		//这个在命令执行前执行
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		//这个在命令执行后执行
	},
	// 还有其他钩子函数
}

var appname string = ""
var apptitle string = ""
var appsrc string = ""

func init() {
	menuCmd.Flags().StringVarP(&appname, "appname", "n", "app", "app name")
	menuCmd.Flags().StringVarP(&apptitle, "apptitle", "t", "", "app name")
	menuCmd.Flags().StringVarP(&appsrc, "src", "o", "./", "app name")
	rootCmd.AddCommand(menuCmd)
}
