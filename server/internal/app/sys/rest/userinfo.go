package rest

import (
	"net/http"

	"github.com/turingdance/infra/wraper"
	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/app/sys/logic"
	"turingdance.com/reliable/internal/app/sys/model"
	"turingdance.com/reliable/internal/pkg/utils"
	"turingdance.com/reliable/internal/types"
)

// 控制器
type Userinfo struct {
}

func init() {
	//RegisterRestCtrl(&Userinfo{})
}

// 搜索
func (ctrl *Userinfo) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Userinfo{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, total, err := logic.SearchUserinfo(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {

		wraper.OkRows(result, total).Encode(w)
	}
}

// 搜索
func (ctrl *Userinfo) Count(w http.ResponseWriter, req *http.Request) {
	arg := args.Userinfo{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if total, err := logic.CountUserinfo(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {

		wraper.OkData(total).Encode(w)
	}
}

func (ctrl *Userinfo) Updatedeptandname(w http.ResponseWriter, req *http.Request) {
	arg := args.Userinfo{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if arg.DeptId == "" {
		wraper.Error("请选择部门").Encode(w)
		return
	}
	if err := logic.Updatedeptandname(arg.UserId, arg.DeptId, arg.Nickname); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(arg.OrgId).Encode(w)
	}
}

// 搜索
func (ctrl *Userinfo) UpdateUserinfo(w http.ResponseWriter, req *http.Request) {
	arg := args.Userinfo{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if arg.DeptId == "" {
		wraper.Error("请选择部门").Encode(w)
		return
	}
	userId, err := logic.ParseUserId(req)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	} else {
		logic.UpdateUserinfo(model.Userinfo{
			UserId:   userId,
			Nickname: arg.Nickname,
			Pic:      arg.Pic,
			DeptId:   arg.DeptId,
			RoleId:   arg.RoleId,
		})
		wraper.OkData(arg.OrgId).Encode(w)
	}
}

// 创建默认用户
func (ctrl *Userinfo) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Userinfo{}
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
		wraper.Error("只有管理员账户才可进行此操作").Encode(w)
		return
	}
	property.Password, _ = utils.GeneratePassword(property.Password)
	//property. = model.RoleMember
	property.Status = types.EnableStatusTrue
	if result, err := logic.CreateUserinfo(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 更新
func (ctrl *Userinfo) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Userinfo{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.UpdateUserinfo(property)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 删除,系统默认都是逻辑删除
func (ctrl *Userinfo) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.Userinfo{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.LogicDeleteUserinfo(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 获取
func (ctrl *Userinfo) GetOne(w http.ResponseWriter, req *http.Request) {
	instanceId := req.URL.Query().Get("userId")

	if d, e := logic.FindUserinfo(instanceId); e != nil {

		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkData(d).Encode(w)
	}
}

// 获取
func (ctrl *Userinfo) Snsinfo(w http.ResponseWriter, req *http.Request) {
	instanceId := req.URL.Query().Get("userId")
	if d, e := logic.FindSimpleInfo(instanceId); e != nil {
		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkData(d).Encode(w)
	}
}
