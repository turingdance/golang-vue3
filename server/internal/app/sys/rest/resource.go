package rest

import (
	"net/http"

	"github.com/turingdance/infra/wraper"
	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/app/sys/logic"
	"turingdance.com/reliable/internal/app/sys/model"
)

// 资源管理控制器
type Resource struct {
}

func init() {
	//RegisterRestCtrl(&Resource{})
}

// 搜索资源管理
func (ctrl *Resource) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Resource{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, total, err := logic.SearchResource(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {

		wraper.OkRows(result, total).Encode(w)
	}
}

// 创建资源管理
func (ctrl *Resource) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Resource{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.CreateResource(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 更新资源管理
func (ctrl *Resource) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Resource{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.UpdateResource(property)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 删除资源管理,系统默认都是逻辑删除
func (ctrl *Resource) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.Resource{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.LogicDeleteResource(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 获取资源管理
func (ctrl *Resource) GetOne(w http.ResponseWriter, req *http.Request) {
	property := model.Resource{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	instanceId := property.ResId
	if d, e := logic.FindResource(instanceId); e != nil {

		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkData(d).Encode(w)
	}
}
