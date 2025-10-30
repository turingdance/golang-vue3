// don't modify !!!!
// create at 2025-06-01 09:29:22
// creeate by winlion
package rest

import (
	"github.com/turingdance/infra/restkit"
)

var (

	// 控制器
	lessonCtrl = &Lesson{}

	// 控制器
	recordCtrl = &Record{}
)
var MapCtrl map[string]any = map[string]any{

	"Lesson": lessonCtrl, // 控制器

	"Record": recordCtrl, // 控制器

}
var Routes []restkit.Route = []restkit.Route{

	{Package: "rest", Module: "Lesson", HandlerFunc: lessonCtrl.Delete, Func: "Delete", Path: "/lesson/{pkId}", Method: []string{"DELETE"}, Comment: "根据主键删除课件资源"},

	{Package: "rest", Module: "Lesson", HandlerFunc: lessonCtrl.GetOne, Func: "GetOne", Path: "/lesson/{pkId}", Method: []string{"GET"}, Comment: "创建一个工作表"},

	{Package: "rest", Module: "Lesson", HandlerFunc: lessonCtrl.Create, Func: "Create", Path: "/lesson", Method: []string{"POST"}, Comment: "创建课件资源"},
	{Package: "rest", Module: "Lesson", HandlerFunc: lessonCtrl.BatchAdd, Func: "BatchAdd", Path: "/lesson/batchAdd", Method: []string{"POST"}, Comment: "创建课件资源"},

	{Package: "rest", Module: "Lesson", HandlerFunc: lessonCtrl.Update, Func: "Update", Path: "/lesson", Method: []string{"PUT"}, Comment: "修改课件资源"},

	{Package: "rest", Module: "Lesson", HandlerFunc: lessonCtrl.Search, Func: "Search", Path: "/lesson/search", Method: []string{"POST", "GET"}, Comment: "根据条件查询课件资源"},

	{Package: "rest", Module: "Lesson", HandlerFunc: lessonCtrl.DeleteBatch, Func: "DeleteBatch", Path: "/lesson", Method: []string{"DELETE"}, Comment: "根据主键批量删除课件资源"},

	{Package: "rest", Module: "Lesson", HandlerFunc: lessonCtrl.Export, Func: "Export", Path: "/lesson/export", Method: []string{"POST", "GET"}, Comment: "导出课件资源"},

	{Package: "rest", Module: "Record", HandlerFunc: recordCtrl.Delete, Func: "Delete", Path: "/record/{pkId}", Method: []string{"DELETE"}, Comment: "根据主键删除上课记录"},

	{Package: "rest", Module: "Record", HandlerFunc: recordCtrl.GetOne, Func: "GetOne", Path: "/record/{pkId}", Method: []string{"GET"}, Comment: "创建一个工作表"},

	{Package: "rest", Module: "Record", HandlerFunc: recordCtrl.Create, Func: "Create", Path: "/record", Method: []string{"POST"}, Comment: "创建上课记录"},

	{Package: "rest", Module: "Record", HandlerFunc: recordCtrl.Update, Func: "Update", Path: "/record", Method: []string{"PUT"}, Comment: "修改上课记录"},

	{Package: "rest", Module: "Record", HandlerFunc: recordCtrl.Search, Func: "Search", Path: "/record/search", Method: []string{"POST", "GET"}, Comment: "根据条件查询上课记录"},
	{Package: "rest", Module: "Record", HandlerFunc: recordCtrl.Mine, Func: "Mine", Path: "/record/mine", Method: []string{"POST", "GET"}, Comment: "根据条件查询用户上课数据"},

	{Package: "rest", Module: "Record", HandlerFunc: recordCtrl.DeleteBatch, Func: "DeleteBatch", Path: "/record", Method: []string{"DELETE"}, Comment: "根据主键批量删除上课记录"},

	{Package: "rest", Module: "Record", HandlerFunc: recordCtrl.Export, Func: "Export", Path: "/record/export", Method: []string{"POST", "GET"}, Comment: "导出上课记录"},
}
