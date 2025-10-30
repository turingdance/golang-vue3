package rest

import (
	"net/http"

	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/app/sys/logic"
	"turingdance.com/reliable/internal/app/sys/model"

	"github.com/turingdance/infra/wraper"

	"strconv"

	"github.com/gorilla/mux"
)

// 控制器
type Role struct {
}

func init() {
	//RegisterRestCtrl(&Role{})
}

// 搜索
func (ctrl *Role) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Role{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, total, err := logic.SearchRole(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {

		wraper.OkRows(result, total).Encode(w)
	}
}

// 创建
func (ctrl *Role) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Role{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.CreateRole(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 授权
func (ctrl *Role) Grant(w http.ResponseWriter, req *http.Request) {
	property := model.Role{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}

	if err := logic.Grant(property.Id, property.RightIds); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkMsg("授权成功").Encode(w)
	}
}

// 更新
func (ctrl *Role) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Role{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.UpdateRole(property)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 启用角色
func (ctrl *Role) Enable(w http.ResponseWriter, req *http.Request) {
	property := model.Role{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	property.State = 1
	if result, err := logic.UpdateRoleState(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).WithMsg("角色已经启用").Encode(w)
	}
}

// 停用角色
func (ctrl *Role) Disable(w http.ResponseWriter, req *http.Request) {
	property := model.Role{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	property.State = 1
	if result, err := logic.UpdateRoleState(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).WithMsg("角色已经启用").Encode(w)
	}
}

// 获取
// @Router /role/{roleId}  [POST,GET]
func (ctrl *Role) GetOne(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req) // 获取参数
	instanceId := vars["roleId"]
	id, _ := strconv.Atoi(instanceId)
	if d, e := logic.FindRole(uint(id)); e != nil {
		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkData(d).Encode(w)
	}
}
