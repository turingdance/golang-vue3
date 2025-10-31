package handle

import (
	"strings"

	"turingdance.com/turing/internal/pkg/utils"
)

type ServiceMeta struct {
	AppName string `json:"appname"` //业务名称
	Model   string `json:"model"`   //
	Action  string `json:"action"`
}

// {service:}
type RequestParam struct {
	Service string      `json:"service"`
	Params  interface{} `json:"params"`
}

// base.model.action
func (r *RequestParam) ServiceMeta() ServiceMeta {
	arr := strings.Split(r.Service, ".")
	arr = append(arr, "empty", "empty", "empty")
	return ServiceMeta{
		arr[0], utils.UcFirst(arr[1]), utils.UcFirst(arr[2]),
	}
}
