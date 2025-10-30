package rest

import (
	"net/http"

	"github.com/turingdance/infra/wraper"
)

// 字典控制器
type Flow struct {
}

// 搜索流程
func (ctrl *Flow) Search(w http.ResponseWriter, req *http.Request) {
	result := make([]string, 0)
	wraper.OkRows(result, 0).Encode(w)

}

// 创建流程
func (ctrl *Flow) Create(w http.ResponseWriter, req *http.Request) {
	wraper.Error("NOT SURPORT").Encode(w)
}

// 更新流程
func (ctrl *Flow) Update(w http.ResponseWriter, req *http.Request) {
	wraper.Error("NOT SURPORT").Encode(w)
}

// 删除流程,系统默认都是逻辑删除
func (ctrl *Flow) Delete(w http.ResponseWriter, req *http.Request) {
	wraper.Error("NOT SURPORT").Encode(w)
}

// 获取流程
func (ctrl *Flow) GetOne(w http.ResponseWriter, req *http.Request) {
	wraper.Error("NOT SURPORT").Encode(w)
}
