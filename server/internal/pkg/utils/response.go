package utils

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Rows  []interface{} `json:"rows,omitempty"`
	Msg   string      `json:"msg,omitempty"`
	Total int64       `json:"total,omitempty"`
}
type ResultBytes struct {
	bytes []byte
	Code int
}
const jsonheader = "application/json; charset=utf-8"
func (r Result) ResponseJson(w http.ResponseWriter) error{
	w.Header().Set("content-type", jsonheader)
	w.WriteHeader(r.Code)
	return json.NewEncoder(w).Encode(r)
}
func (r ResultBytes)ResponseJson(w http.ResponseWriter) (err error){
	w.Header().Set("content-type", jsonheader)
	if r.Code==0{
		r.Code = http.StatusOK
	}
	w.WriteHeader(r.Code)
	_,err = w.Write(r.bytes)
	return err
}
func ResultRows(data interface{}, total int64) Result {
	return Result{
		Code:  http.StatusOK,
		Data:  data,
		Rows: make([]interface{}, 0),
		Total: total,
	}
}
func ResultOk(data interface{}) Result {
	return Result{
		Code: http.StatusOK,
		Data: data,
	}
}
func ResultError(msg any) Result {
	str := ""
	if str1, ok := msg.(string); ok {
		str = str1
	} else if str1, ok := msg.(error); ok {
		str = str1.Error()
	}
	return Result{
		Code: http.StatusNotFound,
		Msg:  str,
	}
}
func ResultByte(in []byte,code ...int) ResultBytes {
	r := ResultBytes{
		bytes: in,
	}
	if len(code)>0{
		r.Code = code[0]
	}
	return r
}
