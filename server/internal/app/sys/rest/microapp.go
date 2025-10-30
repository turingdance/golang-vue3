package rest

import (
	"net/http"

	"github.com/turingdance/infra/wraper"
)

// 岗位控制器
type Microapp struct {
}

func init() {
	//RegisterRestCtrl(&Microapp{})
}

// 搜索岗位
func (ctrl *Microapp) Search(w http.ResponseWriter, req *http.Request) {
	result := make([]string, 0)
	wraper.OkRows(result, 0).Encode(w)

}

// 创建岗位
func (ctrl *Microapp) Create(w http.ResponseWriter, req *http.Request) {
	wraper.Error("NOT SURPORT").Encode(w)
}

// 更新岗位
func (ctrl *Microapp) Update(w http.ResponseWriter, req *http.Request) {
	wraper.Error("NOT SURPORT").Encode(w)
}

// 删除岗位,系统默认都是逻辑删除
func (ctrl *Microapp) Delete(w http.ResponseWriter, req *http.Request) {
	wraper.Error("NOT SURPORT").Encode(w)
}

// 获取岗位
func (ctrl *Microapp) GetOne(w http.ResponseWriter, req *http.Request) {
	wraper.Error("NOT SURPORT").Encode(w)
}
