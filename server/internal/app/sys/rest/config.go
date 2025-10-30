package rest

import (
	"net/http"

	"github.com/turingdance/infra/wraper"
	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/app/sys/logic"
	"turingdance.com/reliable/internal/app/sys/model"

	"strconv"
)

// 控制器
type Config struct {
}

func init() {
	//RegisterRestCtrl(&Config{})
}

// 搜索
func (ctrl *Config) Search(w http.ResponseWriter, req *http.Request) {
	arg := args.Config{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, total, err := logic.SearchConfig(arg); err != nil {
		wraper.Error(err).Encode(w)
	} else {

		wraper.OkRows(result, total).Encode(w)
	}
}

// 创建
func (ctrl *Config) Create(w http.ResponseWriter, req *http.Request) {
	property := model.Config{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.CreateConfig(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 创建
func (ctrl *Config) Save(w http.ResponseWriter, req *http.Request) {
	property := model.Config{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.SaveConfig(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 更新
func (ctrl *Config) Update(w http.ResponseWriter, req *http.Request) {
	property := model.Config{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	result, err := logic.UpdateConfig(property)
	if err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 删除,系统默认都是逻辑删除
func (ctrl *Config) Delete(w http.ResponseWriter, req *http.Request) {
	property := model.Config{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if result, err := logic.LogicDeleteConfig(property); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(result).Encode(w)
	}
}

// 获取
func (ctrl *Config) GetOne(w http.ResponseWriter, req *http.Request) {
	instanceId := req.URL.Query().Get("id")

	id, _ := strconv.Atoi(instanceId)
	if d, e := logic.FindConfig(uint(id)); e != nil {

		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkData(d).Encode(w)
	}
}

// 获取
func (ctrl *Config) Value(w http.ResponseWriter, req *http.Request) {
	property := model.Config{}
	if err := wraper.Bind(req, &property); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if d, e := logic.FindConfigValue(property.OrgId, property.Name); e != nil {

		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkData(d).Encode(w)
	}
}

// 获取万战配置
func (ctrl *Config) Site(w http.ResponseWriter, req *http.Request) {
	if d, e := logic.FindConfigValueByPrefix("site"); e != nil {
		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkRows(d, len(d)).Encode(w)
	}
}

// 获取万战配置
func (ctrl *Config) Prefix(w http.ResponseWriter, req *http.Request) {
	prefix := req.URL.Query().Get("name")
	if d, e := logic.FindConfigValueByPrefix(prefix); e != nil {
		wraper.Error(e.Error()).Encode(w)
	} else {
		wraper.OkRows(d, len(d)).Encode(w)
	}
}
