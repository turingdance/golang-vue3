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
type Rights struct {
}

func init() {
	//RegisterRestCtrl(&Rights{})
}

// 搜索
func (ctrl *Rights) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Rights{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, total, err := logic.SearchRights(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {

		wraper.OkRows(result, total).Encode(w)
	}
}

// 搜索
func (ctrl *Rights) Tree(w http.ResponseWriter, req *http.Request) {
	nodes := logic.AllRights()
	// treenodes, err := treekit.ToTree(nodes, "Id", "Pid")
	// if err != nil {
	// 	wraper.Error(err).Encode(w)
	// } else {
	// 	wraper.OkRows(treenodes, len(treenodes)).Encode(w)
	// }
	treenode := logic.BuildRightsNodeFrom(nodes)
	result := logic.RightsTree(treenode, 0)
	wraper.OkData(result).Encode(w)

}

// 创建
func (ctrl *Rights) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Rights{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.CreateRights(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).WithMsg("节点添加成功").Encode(w)
	}
}

// 更新
func (ctrl *Rights) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Rights{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.UpdateRights(property)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).WithMsg("节点修改成功").Encode(w)
	}
}

// 删除,系统默认都是逻辑删除
func (ctrl *Rights) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.Rights{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.LogicDeleteRights(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 获取
func (ctrl *Rights) GetOne(w http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req) // 获取参数
	instanceId := vars["instanceId"]
	id, _ := strconv.Atoi(instanceId)
	if d, e := logic.FindRights(uint(id)); e != nil {

		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkData(d).Encode(w)
	}
}
