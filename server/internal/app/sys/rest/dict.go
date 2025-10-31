package rest

import (
	"net/http"

	"turingdance.com/turing/internal/app/sys/args"
	"turingdance.com/turing/internal/app/sys/logic"
	"turingdance.com/turing/internal/app/sys/model"

	"github.com/turingdance/infra/wraper"
)

// 字典控制器
type Dict struct {
}

// 搜索字典
func (ctrl *Dict) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Dict{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, total, err := logic.SearchDict(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {

		wraper.OkRows(result, total).Encode(w)
	}
}

// 搜索字典
func (ctrl *Dict) ListAll(w http.ResponseWriter, req *http.Request) {
	arg := args.Dict{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, total, err := logic.SearchAllDict(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {

		wraper.OkRows(result, total).Encode(w)
	}
}

// 创建字典
func (ctrl *Dict) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Dict{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.CreateDict(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 更新字典
func (ctrl *Dict) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Dict{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.UpdateDict(property)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 删除字典,系统默认都是逻辑删除
func (ctrl *Dict) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.Dict{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.LogicDeleteDict(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 获取字典
func (ctrl *Dict) GetOne(w http.ResponseWriter, req *http.Request) {
	property := model.Dict{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result := model.Dict{}
	if property.Id > 0 {
		result, _ = logic.FindDict(property.Id)
	} else {
		result, _ = logic.FindDict(property.Name)
	}
	wraper.OkData(result).Encode(w)

}
