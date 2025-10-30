package rest

import (
	"net/http"

	"github.com/turingdance/infra/wraper"
)

// 岗位控制器
type Instance struct{}

func init() {
	//RegisterRestCtrl(&Instance{})
}

// 搜索流程
func (ctrl *Instance) Search(w http.ResponseWriter, req *http.Request) {
	result := make([]string, 0)
	wraper.OkRows(result, 0).Encode(w)

}

// 创建流程
func (ctrl *Instance) Create(w http.ResponseWriter, req *http.Request) {
	wraper.Error("NOT SURPORT").Encode(w)
}

// 更新流程
func (ctrl *Instance) Update(w http.ResponseWriter, req *http.Request) {
	wraper.Error("NOT SURPORT").Encode(w)
}

// 删除流程,系统默认都是逻辑删除
func (ctrl *Instance) Delete(w http.ResponseWriter, req *http.Request) {
	wraper.Error("NOT SURPORT").Encode(w)
}

// 获取流程
func (ctrl *Instance) GetOne(w http.ResponseWriter, req *http.Request) {
	wraper.Error("NOT SURPORT").Encode(w)
}
