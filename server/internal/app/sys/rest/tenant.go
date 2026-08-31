package rest

import (
	"net/http"

	"turingdance.com/turing/internal/app/sys/args"
	"turingdance.com/turing/internal/app/sys/logic"
	"turingdance.com/turing/internal/app/sys/model"
	"turingdance.com/turing/internal/pkg/tokenkit"

	"github.com/turingdance/infra/wraper"
)

// 机构信息控制器
type Tenant struct {
}

func init() {
	//RegisterRestCtrl(&Tenant{})
}

// 搜索机构信息
func (ctrl *Tenant) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Tenant{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, total, err := logic.SearchTenant(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {

		wraper.OkRows(result, total).Encode(w)
	}
}

// 搜索机构信息
func (ctrl *Tenant) Mine(w http.ResponseWriter, req *http.Request) {
	_, err := tokenkit.ParseUserId(req)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	wraper.Error("not surport").Encode(w)
}

// 创建机构信息
func (ctrl *Tenant) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Tenant{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.CreateTenant(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 更新机构信息
func (ctrl *Tenant) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Tenant{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.UpdateTenant(property)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 删除机构信息,系统默认都是逻辑删除
func (ctrl *Tenant) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.Tenant{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.LogicDeleteTenant(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 获取机构信息
func (ctrl *Tenant) GetOne(w http.ResponseWriter, req *http.Request) {
	instanceId := req.URL.Query().Get("tenantId")
	if d, e := logic.FindTenant(instanceId); e != nil {

		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkData(d).Encode(w)
	}
}
