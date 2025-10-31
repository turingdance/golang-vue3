package rest

import (
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/turingdance/infra/cond"
	"github.com/turingdance/infra/slicekit"
	"github.com/turingdance/infra/wraper"
	"github.com/xuri/excelize/v2"
	"turingdance.com/turing/internal/app/sys/args"
	"turingdance.com/turing/internal/app/sys/logic"
	"turingdance.com/turing/internal/app/sys/model"
	"turingdance.com/turing/internal/pkg/utils"
	"turingdance.com/turing/internal/server/auth"
	"turingdance.com/turing/internal/types"
)

// @Summary 账号管理模块
// @Router /account
// @Tags Account
type Account struct {
}

// 用户注册
func (ctrl *Account) Register(w http.ResponseWriter, req *http.Request) {
	arg := args.Account{}
	if err := wraper.Bind(req, &arg); err != nil {
		return
	}
	if userId, err := logic.Register(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		// 如果存在店铺名称,则创建一个门店
		org := model.Org{}
		org.UserId = userId
		// 注册时不注册团队信息
		if arg.Orgname != "" {
			org.Title = arg.Orgname
		}
		if arg.Tel != "" {
			org.Tel = arg.Tel
		}
		if arg.Linker != "" {
			org.Linker = arg.Linker
		}
		if len(org.Title) > 0 {
			_, err := logic.CreateOrg(org)
			if err != nil {
				log.Error(err.Error())
			}
		}
		userAttr := model.Userinfo{}

		if len(arg.Nickname) > 0 {
			userAttr.Nickname = arg.Nickname
		} else {
			userAttr.Nickname = "用户" + utils.RandNumber(4)
		}

		if len(arg.Pic) > 0 {
			userAttr.Pic = arg.Pic
		}

		userAttr.UserId = userId
		logic.UpdateUserinfo(userAttr)
		resp := ""
		if arg.Autologin {
			if token, err := logic.NewAccountSve().LoginUseUserId(userId); err == nil {
				resp = token
			}
		}
		wraper.OkData(resp).WithMsg("注册成功,请登录").Encode(w)
	}
}

// 搜索
func (ctrl *Account) Bindwxuser(w http.ResponseWriter, req *http.Request) {

	railm, err := auth.ParseToken(req)
	userId := ""
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	userId = railm["userId"].(string)
	if userId == "" {
		wraper.Error("鉴权信息有误,请重试").Encode(w)
		return
	}
	wraper.Error("not surport").Encode(w)
}

// 导入excel 的模板
func (ctrl *Account) Template(w http.ResponseWriter, r *http.Request) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheetname := "导入账号模板"
	f.SetSheetName("Sheet1", sheetname)
	// 创建一个工作表
	index, err := f.GetSheetIndex(sheetname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 设置工作簿的默认工作表
	f.SetActiveSheet(index)
	// 设置单元格的值
	f.SetCellValue(sheetname, "A1", "用户名")
	f.SetCellValue(sheetname, "B1", "密码")
	f.SetCellValue(sheetname, "C1", "姓名")
	f.SetCellValue(sheetname, "A2", "zhang001")
	f.SetCellValue(sheetname, "B2", "123456")
	f.SetCellValue(sheetname, "C2", "张三")

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.Header().Set("Content-Disposition", "attachment; filename="+sheetname+time.Now().Format("20060102")+".xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	w.Header().Set("Cache-Control", "must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Pragma", "public")

	if err := f.Write(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// 导入excel
func (ctrl *Account) Import(w http.ResponseWriter, r *http.Request) {
	// 解析 multipart/form-data 请求
	if err := r.ParseMultipartForm(10 << 20); // 10 MB 最大内存
	err != nil {
		wraper.Error(err).Encode(w)
		return
	}

	// 获取上传的文件
	file, _, err := r.FormFile("file")
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	defer file.Close()

	// 创建临时文件保存上传的 Excel
	tmpFile, err := os.CreateTemp("", "upload-*.xlsx")
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	defer os.Remove(tmpFile.Name()) // 处理完后删除临时文件

	// 将上传的文件内容复制到临时文件
	if _, err := tmpFile.ReadFrom(file); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	tmpFile.Close()
	// 登录账号|登录密码|用户姓名
	// 解析 Excel 文件
	if users, err := parseExcel(tmpFile.Name()); err != nil {
		wraper.Error(err).Encode(w)
		return
	} else {
		if total, err := logic.NewAccountSve().BatchCreate(users); err != nil {
			wraper.Error(err).Encode(w)
		} else {
			wraper.OkMsg(fmt.Sprintf("成功导入%d条记录", total)).Encode(w)
		}
	}

}
func parseExcel(filePath string) (users []model.Userinfo, err error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return users, fmt.Errorf("打开 Excel 文件失败: %v", err)
	}
	defer f.Close()

	// 获取第一个工作表的名称
	sheetName := f.GetSheetName(0)

	// 获取工作表中的所有行
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return users, fmt.Errorf("获取行数据失败: %v", err)
	}
	// 第一行  用户名  第二行 密码
	// 遍历所有行
	users = make([]model.Userinfo, 0)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if slicekit.Some[string](row, func(arr []string, ele string) bool {
			return ele == ""
		}) {
			return users, fmt.Errorf("第%d行数据不全", i)
		}

		user := model.Userinfo{}
		user.Username = row[0]
		user.Password = row[1]
		user.Nickname = row[2]
		users = append(users, user)

	}

	return users, nil
}

// 这个只用来修改社交信息
func (ctrl *Account) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Userinfo{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	updatObj := model.Userinfo{}
	updatObj.Pic = property.Pic
	updatObj.Nickname = property.Nickname
	userId, err := logic.ParseUserId(req)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	updatObj.UserId = userId
	if result, err := logic.UpdateUserinfo(updatObj); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).WithMsg("信息更新成功").Encode(w)
	}
}

// 创建
func (ctrl *Account) Login(w http.ResponseWriter, req *http.Request) {
	property := args.Account{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.Login(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).WithMsg("登录成功").Encode(w)
	}
}

// 创建
func (ctrl *Account) ResetPwd(w http.ResponseWriter, req *http.Request) {
	property := args.Account{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if err := logic.RestPwd(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkMsg("密码重置成功").Encode(w)
	}
}

// 用户专用
func (ctrl *Account) UpdateMyPwd(w http.ResponseWriter, req *http.Request) {
	property := args.Account{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	userId, err := logic.ParseUserId(req)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	property.UserId = userId

	if err := logic.UpdatePassword(property, true); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkMsg("用户密码修改成功").Encode(w)
	}
}

// 创建
// 管理员专用
func (ctrl *Account) UpdatePwd(w http.ResponseWriter, req *http.Request) {
	property := args.Account{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	role, err := logic.ParseRole(req)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if role != model.RoleAdmin {
		wraper.Error("只有管理员才有此权限进行操作").Encode(w)
		return
	}
	if property.UserId == "" {
		wraper.Error("缺少用户ID").Encode(w)
		return
	}
	if err := logic.UpdatePassword(args.Account{UserId: property.UserId, Password: property.Password}, false); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkMsg("密码重置成功").Encode(w)
	}
}

// 修改当前账号手机号
func (ctrl *Account) Resetmobile(w http.ResponseWriter, req *http.Request) {
	property := args.Account{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	userId, err := logic.ParseUserId(req)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	property.UserId = userId
	if err := logic.RestMobile(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkMsg("手机号重置成功").Encode(w)
	}
}

// 修改当前账号手机号
func (ctrl *Account) UpdateUserName(w http.ResponseWriter, req *http.Request) {
	property := args.Account{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	role, err := logic.ParseRole(req)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if role != model.RoleAdmin {
		wraper.OkMsg("只有管理员才可以操作该接口").Encode(w)
		return
	}

	if err := logic.UpdateUserName(property.UserId, property.Username); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkMsg("用户名重置成功").Encode(w)
	}
}

// 更新
func (ctrl *Account) GetInfo(w http.ResponseWriter, req *http.Request) {

	userId, err := logic.ParseUserId(req)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.FindUserinfo(userId)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		// 加载用户的rights.biz
		result.Scope = make([]string, 0)
		rightIds := result.Role.RightIds
		rightMap := logic.RightsMap()
		for _, id := range rightIds {
			if r, ok := rightMap[id]; ok {
				result.Scope = append(result.Scope, r.Biz)
			}
		}
		wraper.OkData(result.Desensitize()).Encode(w)
	}
}

// 更新
func (ctrl *Account) GetMenu(w http.ResponseWriter, req *http.Request) {

	userId, err := logic.ParseUserId(req)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.FindUserinfo(userId)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {

		role, _ := logic.FindRole(result.RoleId)
		rights, _ := logic.FindRightsByIds(role.RightIds)
		// 分组或者分页面都可
		nodes := slicekit.Filter[model.Rights](rights, func(right model.Rights, index int, slice []model.Rights) bool {
			return right.Type == model.RightsTypeView || right.Type == model.RightsTypeGroup
		})
		nodes = slicekit.Sort[model.Rights](nodes, func(item1, item2 model.Rights) bool {
			return item1.SortIndex < item2.SortIndex
		})
		// 构建树形节点
		treenode := logic.BuildRightsNodeFrom(nodes)
		results := logic.RightsTree(treenode, 0)

		wraper.OkData(results).WithTotal(len(nodes)).Encode(w)
	}
}

// 搜索区域
// Router /account/search [POST,GET]
func (ctrl *Account) Search(w http.ResponseWriter, req *http.Request) {
	arg := cond.NewCondWrapper()
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, total, err := logic.Search(&model.Userinfo{}, arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkRows(result, total).Encode(w)
	}
}

// 更新
func (ctrl *Account) Permit(w http.ResponseWriter, req *http.Request) {
	userId, err := logic.ParseUserId(req)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.FindUserinfo(userId)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {

		role, _ := logic.FindRole(result.RoleId)
		rights, _ := logic.FindRightsByIds(role.RightIds)
		results := slicekit.Map[model.Rights, string](rights, func(right model.Rights, index int, slice []model.Rights) string {
			return right.Biz
		})
		wraper.OkData(results).WithTotal(len(results)).Encode(w)
	}
}

// token 续期
func (ctrl *Account) Renewal(w http.ResponseWriter, req *http.Request) {
	token := utils.GetAuthorizationFromRequest(req)
	result, err := logic.ParseAccountFrom(token)
	if err != nil {

		wraper.Error(err).Encode(w)
	} else {

		if token, err := logic.NewAccountSve().LoginUseUserId(result.UserId); err == nil {
			wraper.OkData(token).Encode(w)
		} else {
			wraper.Error(err).Encode(w)
		}

	}
}

// 管理员修改用户密码
// 更新
func (ctrl *Account) Enable(w http.ResponseWriter, req *http.Request) {
	property := model.Userinfo{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}

	result, err := logic.UpdateUserinfo(model.Userinfo{
		UserId: property.UserId,
		Status: types.EnableStatusTrue,
	})
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).WithMsg("账号已激活").Encode(w)
	}
}

// 更新
func (ctrl *Account) Disable(w http.ResponseWriter, req *http.Request) {
	property := model.Userinfo{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	property.Status = types.EnableStatusFalse
	result, err := logic.UpdateUserinfo(model.Userinfo{
		UserId: property.UserId,
		Status: types.EnableStatusFalse,
	})
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).WithMsg("账号已停用").Encode(w)
	}
}

// 删除,系统默认都是逻辑删除
// Router /account/deleteIt [POST,DELETE]
func (ctrl *Account) DeleteIt(w http.ResponseWriter, req *http.Request) {
	property := model.Userinfo{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.Delete(&property, "user_id = ?", property.UserId); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// Router /account/deleteIts [POST,DELETE]
func (ctrl *Account) DeleteIts(w http.ResponseWriter, req *http.Request) {
	property := args.AccountKeyBatch{
		UserIds: make([]string, 0),
	}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if _, err := logic.Delete(&model.Userinfo{}, "user_id in ?", property.UserIds); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkMsg("刪除成功").Encode(w)
	}
}
