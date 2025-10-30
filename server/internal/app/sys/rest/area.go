package rest

import (
	"net/http"

	"github.com/turingdance/infra/wraper"
	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/app/sys/logic"
	"turingdance.com/reliable/internal/app/sys/model"
)

// 区域控制器
type Area struct {
}

func init() {
	//RegisterRestCtrl(&Area{})
}

// 搜索区域
func (ctrl *Area) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Area{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, total, err := logic.SearchArea(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkRows(result, total).Encode(w)
	}
}

// 创建区域
func (ctrl *Area) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Area{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.CreateArea(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 更新区域
func (ctrl *Area) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Area{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.UpdateArea(property)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 删除区域,系统默认都是逻辑删除
func (ctrl *Area) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.Area{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.LogicDeleteArea(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 获取区域
func (ctrl *Area) GetOne(w http.ResponseWriter, req *http.Request) {
	property := model.Area{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	instanceId := property.AreaId
	if d, e := logic.FindArea(instanceId); e != nil {
		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkData(d).Encode(w)
	}
}
