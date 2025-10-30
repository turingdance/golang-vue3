package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

var svcConfig = &service.Config{
	Name:        PARAM_APP_NAME,              //服务显示名称
	DisplayName: PARAM_APP_NAME + " service", //服务名称
	Description: PARAM_APP_TTTLE + "应用服务程序",  //服务描述
}

var prg = &program{}

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	ch := make(chan os.Signal)
	<-ch
}

func (p *program) Stop(s service.Service) error {
	return nil
}

var CmdServiceInstall = &cobra.Command{
	Use:   "install",
	Short: "install:注册服务",
	Long:  PARAM_APP_TTTLE + `客户端服务注册工具集`,
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value.String()
		memo := cmd.Flag("memo").Value.String()
		if len(name) == 0 {
			cmd.Help()
			return
		}
		svcConfig = &service.Config{
			Name:        name,              //服务显示名称
			DisplayName: name + " Service", //服务名称
			Description: memo,              //服务描述
		}
		s, _ := service.New(prg, svcConfig)
		if err := s.Install(); err != nil {
			fmt.Printf("install service %s error %s\n", name, err.Error())
		} else {
			fmt.Printf(`install %s ok\n`, name)
		}
		// 这里完成注册逻辑
	},
}

var CmdServiceUninstall = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall:卸载服务",
	Long:  PARAM_APP_TTTLE + `客户端服务注册工具集`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.Join(args, "--"))
		if len(args) == 0 {
			cmd.Help()
			return
		}
		// name := cmd.Flag("name").Value.String()
		// svcConfig = &service.Config{
		// 	Name:        name,              //服务显示名称
		// 	DisplayName: name + " Service", //服务名称
		// }
		// s, _ := service.New(prg, svcConfig)
		// if err := s.Uninstall(); err != nil {
		// 	fmt.Errorf("%s", err.Error())
		// } else {
		// 	fmt.Println(`register ok`)
		// }
		// 这里完成注册逻辑
	},
}

var CmdService = &cobra.Command{
	Use:   "service",
	Short: "service:服务相关处理",
	Long:  PARAM_APP_TTTLE + `服务注册/卸载工具集`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.Join(args, "-"))
		if len(args) == 0 {
			cmd.Help()
		} else {
			fmt.Printf("no args %+v", args)
		}
	},
}

func init() {
	CmdServiceInstall.PersistentFlags().StringP(PARAM_KEY_NAME, "n", PARAM_APP_NAME, "service name")
	CmdServiceInstall.PersistentFlags().StringP(PARAM_KEY_MEMO, "m", PARAM_APP_NAME, "service name")
	CmdServiceUninstall.PersistentFlags().StringP(PARAM_KEY_MEMO, "m", PARAM_APP_TTTLE+"客户端助手", "descript service")
	CmdService.AddCommand(CmdServiceInstall)
	CmdService.AddCommand(CmdServiceUninstall)
}
