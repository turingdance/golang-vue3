package rest

import (
	"net/http"

	"turingdance.com/turing/internal/app/sys/args"
	"turingdance.com/turing/internal/app/sys/logic"
	"turingdance.com/turing/internal/app/sys/model"

	"github.com/turingdance/infra/wraper"

	"strconv"
)

// 机构信息控制器
type Org struct {
}

func init() {
	//RegisterRestCtrl(&Org{})
}

// 搜索机构信息
func (ctrl *Org) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Org{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, total, err := logic.SearchOrg(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {

		wraper.OkRows(result, total).Encode(w)
	}
}

// 搜索机构信息
func (ctrl *Org) Mine(w http.ResponseWriter, req *http.Request) {

	_, err := logic.ParseUserId(req)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	wraper.Error("not surport").Encode(w)
}

// 创建机构信息
func (ctrl *Org) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Org{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.CreateOrg(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 更新机构信息
func (ctrl *Org) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Org{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.UpdateOrg(property)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 删除机构信息,系统默认都是逻辑删除
func (ctrl *Org) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.Org{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.LogicDeleteOrg(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 获取机构信息
func (ctrl *Org) GetOne(w http.ResponseWriter, req *http.Request) {
	instanceId := req.URL.Query().Get("orgId")

	id, _ := strconv.Atoi(instanceId)
	if d, e := logic.FindOrg(uint(id)); e != nil {

		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkData(d).Encode(w)
	}
}
