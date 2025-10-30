package rest

import (
	"net/http"

	"github.com/turingdance/infra/wraper"
	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/app/sys/logic"
	"turingdance.com/reliable/internal/app/sys/model"
)

// 岗位控制器
type Dept struct {
}

func init() {
	//RegisterRestCtrl(&Dept{})
}

// 搜索岗位
func (ctrl *Dept) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Dept{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, total, err := logic.SearchDept(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {

		wraper.OkRows(result, total).Encode(w)
	}
}

// 创建岗位
func (ctrl *Dept) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Dept{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.CreateDept(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).WithMsg("部门创建成功").Encode(w)
	}
}

// 更新岗位
func (ctrl *Dept) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Dept{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.UpdateDept(property)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).WithMsg("部门修改成功").Encode(w)
	}
}

// 删除岗位,系统默认都是逻辑删除
func (ctrl *Dept) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.Dept{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if err := logic.LogicDeleteDept(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkMsg("删除成功").Encode(w)
	}
}

func (ctrl *Dept) Tree(w http.ResponseWriter, req *http.Request) {
	depts := logic.AllDept()
	nodes := logic.Transfer(depts)
	d := logic.DeptTree(nodes, "")
	wraper.OkData(d).Encode(w)
}

// 获取岗位
func (ctrl *Dept) GetOne(w http.ResponseWriter, req *http.Request) {
	property := model.Dept{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	instanceId := property.Id
	if d, e := logic.FindDept(instanceId); e != nil {
		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkData(d).Encode(w)
	}
}
