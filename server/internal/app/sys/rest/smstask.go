package rest

import (
	"net/http"

	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/app/sys/logic"
	"turingdance.com/reliable/internal/app/sys/model"

	"github.com/turingdance/infra/wraper"

	"strconv"
)

// 短信发送记录控制器
type Smstask struct {
}

func init() {
	//RegisterRestCtrl(&Smstask{})
}

// 搜索
func (ctrl *Smstask) Send(w http.ResponseWriter, req *http.Request) {
	arg := args.Account{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	logic.SendSms(arg.Mobile)
	wraper.OkMsg("短信已发送,请稍候").Encode(w)

}

// 搜索短信发送记录
func (ctrl *Smstask) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Smstask{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, total, err := logic.SearchSmstask(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {

		wraper.OkRows(result, total).Encode(w)
	}
}

// 创建短信发送记录
func (ctrl *Smstask) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Smstask{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.CreateSmstask(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 更新短信发送记录
func (ctrl *Smstask) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Smstask{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.UpdateSmstask(property)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 删除短信发送记录,系统默认都是逻辑删除
func (ctrl *Smstask) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.Smstask{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.LogicDeleteSmstask(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 获取短信发送记录
func (ctrl *Smstask) GetOne(w http.ResponseWriter, req *http.Request) {
	instanceId := req.URL.Query().Get("id")

	id, _ := strconv.Atoi(instanceId)
	if d, e := logic.FindSmstask(uint(id)); e != nil {

		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkData(d).Encode(w)
	}
}
