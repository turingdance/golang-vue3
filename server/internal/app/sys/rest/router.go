// don't modify !!!!
// create at 2024-11-22 15:45:08
// creeate by winlion
//
//go:generate  codectl router -t ./router.tpl -a winlion -s . -d . -n router.go
package rest

import (
	"github.com/turingdance/infra/restkit"
)

var (

	// 账号管理模块
	accountCtrl = &Account{}

	// 区域控制器
	areaCtrl = &Area{}

	// 这一层里面实现router
	storageCtrl = &Storage{}

	//
	captchaCtrl = &Captcha{}

	// 控制器
	configCtrl = &Config{}

	// 岗位控制器
	deptCtrl = &Dept{}

	// 字典控制器
	dictCtrl = &Dict{}

	// 机构信息控制器
	orgCtrl = &Org{}

	//
	ossCtrl = &Oss{}

	// 资源管理控制器
	resourceCtrl = &Resource{}

	// 控制器
	rightsCtrl = &Rights{}

	// 控制器
	roleCtrl = &Role{}

	// 短信发送记录控制器
	smstaskCtrl = &Smstask{}

	// 控制器
	userinfoCtrl = &Userinfo{}
)
var MapCtrl map[string]any = map[string]any{

	"Account": accountCtrl, // 账号管理模块

	"Area": areaCtrl, // 区域控制器

	"Storage": storageCtrl, // 这一层里面实现router

	"Captcha": captchaCtrl, //

	"Config": configCtrl, // 控制器

	"Dept": deptCtrl, // 岗位控制器

	"Dict": dictCtrl, // 字典控制器

	"Org": orgCtrl, // 机构信息控制器

	"Oss": ossCtrl, //

	"Resource": resourceCtrl, // 资源管理控制器

	"Rights": rightsCtrl, // 控制器

	"Role": roleCtrl, // 控制器

	"Smstask": smstaskCtrl, // 短信发送记录控制器

	"Userinfo": userinfoCtrl, // 控制器

}
var Routes []restkit.Route = []restkit.Route{

	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.Register, Func: "Register", Path: "/account/register", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "用户注册"},

	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.Bindwxuser, Func: "Bindwxuser", Path: "/account/bindwxuser", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "如果存在店铺名称,则创建一个门店"},

	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.Update, Func: "Update", Path: "/account/update", Method: []string{"PUT", "POST", "OPTIONS"}, Comment: ""},

	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.Login, Func: "Login", Path: "/account/login", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建"},

	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.ResetPwd, Func: "ResetPwd", Path: "/account/resetPwd", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建"},

	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.UpdateMyPwd, Func: "UpdateMyPwd", Path: "/account/updateMyPwd", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "用户专用"},

	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.UpdatePwd, Func: "UpdatePwd", Path: "/account/updatePwd", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建"},

	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.Resetmobile, Func: "Resetmobile", Path: "/account/resetmobile", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "修改当前账号手机号"},

	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.UpdateUserName, Func: "UpdateUserName", Path: "/account/updateUserName", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "修改当前账号手机号"},
	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.Search, Func: "Search", Path: "/account/search", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "查询"},

	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.GetInfo, Func: "GetInfo", Path: "/account/getInfo", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "更新"},
	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.GetMenu, Func: "GetMenu", Path: "/account/getMenu", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "可用菜单"},

	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.Renewal, Func: "Renewal", Path: "/account/renewal", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "token 续期"},

	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.Enable, Func: "Enable", Path: "/account/enable", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "管理员修改用户密码"},

	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.Disable, Func: "Disable", Path: "/account/disable", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "更新"},
	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.Template, Func: "Template", Path: "/account/template", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "导出模板"},
	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.Import, Func: "Import", Path: "/account/import", Method: []string{"POST", "OPTIONS"}, Comment: "批量导入"},
	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.DeleteIt, Func: "DeleteIt", Path: "/account/deleteIt", Method: []string{"DELETE", "POST", "OPTIONS"}, Comment: "删除单条记录"},
	{Package: "rest", Module: "Account", HandlerFunc: accountCtrl.DeleteIts, Func: "DeleteIts", Path: "/account/deleteIts", Method: []string{"DELETE", "POST", "OPTIONS"}, Comment: "删除多条记录"},

	{Package: "rest", Module: "Area", HandlerFunc: areaCtrl.Search, Func: "Search", Path: "/area/search", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "RegisterRestCtrl(&Area{})"},

	{Package: "rest", Module: "Area", HandlerFunc: areaCtrl.Create, Func: "Create", Path: "/area/create", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建区域"},

	{Package: "rest", Module: "Area", HandlerFunc: areaCtrl.Update, Func: "Update", Path: "/area/update", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "更新区域"},

	{Package: "rest", Module: "Area", HandlerFunc: areaCtrl.Delete, Func: "Delete", Path: "/area/delete", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "删除区域,系统默认都是逻辑删除"},

	{Package: "rest", Module: "Area", HandlerFunc: areaCtrl.GetOne, Func: "GetOne", Path: "/area/getOne", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获取区域"},

	{Package: "rest", Module: "Storage", HandlerFunc: storageCtrl.Render, Func: "Render", Path: "/storage/render", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "处理/upload 逻辑"},
	{Package: "rest", Module: "Storage", HandlerFunc: storageCtrl.ShareUrl, Func: "ShareUrl", Path: "/storage/shareUrl", Method: []string{"GET", "OPTIONS"}, Comment: "暴漏"},

	{Package: "rest", Module: "Storage", HandlerFunc: storageCtrl.Upload, Func: "Upload", Path: "/storage/upload", Method: []string{"POST", "OPTIONS"}, Comment: "处理/upload 逻辑"},
	{Package: "rest", Module: "Storage", HandlerFunc: storageCtrl.Config, Func: "Config", Path: "/storage/config", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "配置文件"},

	{Package: "rest", Module: "Captcha", HandlerFunc: captchaCtrl.Get, Func: "Get", Path: "/captcha/get", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "验证码"},

	{Package: "rest", Module: "Config", HandlerFunc: configCtrl.Search, Func: "Search", Path: "/config/search", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "资源加载"},

	{Package: "rest", Module: "Config", HandlerFunc: configCtrl.Create, Func: "Create", Path: "/config/create", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建"},

	{Package: "rest", Module: "Config", HandlerFunc: configCtrl.Save, Func: "Save", Path: "/config/save", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建"},

	{Package: "rest", Module: "Config", HandlerFunc: configCtrl.Update, Func: "Update", Path: "/config/update", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "更新"},

	{Package: "rest", Module: "Config", HandlerFunc: configCtrl.Delete, Func: "Delete", Path: "/config/delete", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "删除,系统默认都是逻辑删除"},

	{Package: "rest", Module: "Config", HandlerFunc: configCtrl.GetOne, Func: "GetOne", Path: "/config/getOne", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获取"},

	{Package: "rest", Module: "Config", HandlerFunc: configCtrl.Value, Func: "Value", Path: "/config/value", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获取"},
	{Package: "rest", Module: "Config", HandlerFunc: configCtrl.Site, Func: "Site", Path: "/config/site", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获取"},
	{Package: "rest", Module: "Config", HandlerFunc: configCtrl.Prefix, Func: "Prefix", Path: "/config/name", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获取"},

	{Package: "rest", Module: "Dept", HandlerFunc: deptCtrl.Search, Func: "Search", Path: "/dept/search", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "RegisterRestCtrl(&Dept{})"},

	{Package: "rest", Module: "Dept", HandlerFunc: deptCtrl.Create, Func: "Create", Path: "/dept/create", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建岗位"},

	{Package: "rest", Module: "Dept", HandlerFunc: deptCtrl.Update, Func: "Update", Path: "/dept/update", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "更新岗位"},

	{Package: "rest", Module: "Dept", HandlerFunc: deptCtrl.Delete, Func: "Delete", Path: "/dept/delete", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "删除岗位,系统默认都是逻辑删除"},

	{Package: "rest", Module: "Dept", HandlerFunc: deptCtrl.Tree, Func: "Tree", Path: "/dept/tree", Method: []string{"GET", "POST", "OPTIONS"}, Comment: ""},

	{Package: "rest", Module: "Dept", HandlerFunc: deptCtrl.GetOne, Func: "GetOne", Path: "/dept/getOne", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获取岗位"},

	{Package: "rest", Module: "Dict", HandlerFunc: dictCtrl.Search, Func: "Search", Path: "/dict/search", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "搜索字典"},
	{Package: "rest", Module: "Dict", HandlerFunc: dictCtrl.ListAll, Func: "ListAll", Path: "/dict/listAll", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "搜索字典"},

	{Package: "rest", Module: "Dict", HandlerFunc: dictCtrl.Create, Func: "Create", Path: "/dict/create", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建字典"},

	{Package: "rest", Module: "Dict", HandlerFunc: dictCtrl.Update, Func: "Update", Path: "/dict/update", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "更新字典"},

	{Package: "rest", Module: "Dict", HandlerFunc: dictCtrl.Delete, Func: "Delete", Path: "/dict/delete", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "删除字典,系统默认都是逻辑删除"},

	{Package: "rest", Module: "Dict", HandlerFunc: dictCtrl.GetOne, Func: "GetOne", Path: "/dict/getOne", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获取字典"},

	{Package: "rest", Module: "Org", HandlerFunc: orgCtrl.Search, Func: "Search", Path: "/org/search", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "RegisterRestCtrl(&Org{})"},

	{Package: "rest", Module: "Org", HandlerFunc: orgCtrl.Mine, Func: "Mine", Path: "/org/mine", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "搜索机构信息"},

	{Package: "rest", Module: "Org", HandlerFunc: orgCtrl.Create, Func: "Create", Path: "/org/create", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建机构信息"},

	{Package: "rest", Module: "Org", HandlerFunc: orgCtrl.Update, Func: "Update", Path: "/org/update", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "更新机构信息"},

	{Package: "rest", Module: "Org", HandlerFunc: orgCtrl.Delete, Func: "Delete", Path: "/org/delete", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "删除机构信息,系统默认都是逻辑删除"},

	{Package: "rest", Module: "Org", HandlerFunc: orgCtrl.GetOne, Func: "GetOne", Path: "/org/getOne", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获取机构信息"},

	{Package: "rest", Module: "Oss", HandlerFunc: ossCtrl.Policy, Func: "Policy", Path: "/oss/policy", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "RegisterRestCtrl(&Oss{})"},

	{Package: "rest", Module: "Oss", HandlerFunc: ossCtrl.Callback, Func: "Callback", Path: "/oss/callback", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获得token"},

	{Package: "rest", Module: "Resource", HandlerFunc: resourceCtrl.Search, Func: "Search", Path: "/resource/search", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "RegisterRestCtrl(&Resource{})"},

	{Package: "rest", Module: "Resource", HandlerFunc: resourceCtrl.Create, Func: "Create", Path: "/resource/create", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建资源管理"},

	{Package: "rest", Module: "Resource", HandlerFunc: resourceCtrl.Update, Func: "Update", Path: "/resource/update", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "更新资源管理"},

	{Package: "rest", Module: "Resource", HandlerFunc: resourceCtrl.Delete, Func: "Delete", Path: "/resource/delete", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "删除资源管理,系统默认都是逻辑删除"},

	{Package: "rest", Module: "Resource", HandlerFunc: resourceCtrl.GetOne, Func: "GetOne", Path: "/resource/getOne", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获取资源管理"},

	{Package: "rest", Module: "Rights", HandlerFunc: rightsCtrl.Search, Func: "Search", Path: "/rights/search", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "RegisterRestCtrl(&Rights{})"},

	{Package: "rest", Module: "Rights", HandlerFunc: rightsCtrl.Tree, Func: "Tree", Path: "/rights/tree", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "搜索"},

	{Package: "rest", Module: "Rights", HandlerFunc: rightsCtrl.Create, Func: "Create", Path: "/rights/create", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建"},

	{Package: "rest", Module: "Rights", HandlerFunc: rightsCtrl.Update, Func: "Update", Path: "/rights/update", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "更新"},

	{Package: "rest", Module: "Rights", HandlerFunc: rightsCtrl.Delete, Func: "Delete", Path: "/rights/delete", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "删除,系统默认都是逻辑删除"},

	{Package: "rest", Module: "Rights", HandlerFunc: rightsCtrl.GetOne, Func: "GetOne", Path: "/rights/getOne", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获取"},

	{Package: "rest", Module: "Role", HandlerFunc: roleCtrl.GetOne, Func: "GetOne", Path: "/role/{roleId}", Method: []string{"POST", "GET"}, Comment: "获取"},

	{Package: "rest", Module: "Role", HandlerFunc: roleCtrl.Search, Func: "Search", Path: "/role/search", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "RegisterRestCtrl(&Role{})"},

	{Package: "rest", Module: "Role", HandlerFunc: roleCtrl.Create, Func: "Create", Path: "/role/create", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建"},

	{Package: "rest", Module: "Role", HandlerFunc: roleCtrl.Grant, Func: "Grant", Path: "/role/grant", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "授权"},

	{Package: "rest", Module: "Role", HandlerFunc: roleCtrl.Update, Func: "Update", Path: "/role/update", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "更新"},

	{Package: "rest", Module: "Role", HandlerFunc: roleCtrl.Enable, Func: "Enable", Path: "/role/enable", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "启用角色"},

	{Package: "rest", Module: "Role", HandlerFunc: roleCtrl.Disable, Func: "Disable", Path: "/role/disable", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "停用角色"},

	{Package: "rest", Module: "Smstask", HandlerFunc: smstaskCtrl.Send, Func: "Send", Path: "/smstask/send", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "RegisterRestCtrl(&Smstask{})"},

	{Package: "rest", Module: "Smstask", HandlerFunc: smstaskCtrl.Search, Func: "Search", Path: "/smstask/search", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "搜索短信发送记录"},

	{Package: "rest", Module: "Smstask", HandlerFunc: smstaskCtrl.Create, Func: "Create", Path: "/smstask/create", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建短信发送记录"},

	{Package: "rest", Module: "Smstask", HandlerFunc: smstaskCtrl.Update, Func: "Update", Path: "/smstask/update", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "更新短信发送记录"},

	{Package: "rest", Module: "Smstask", HandlerFunc: smstaskCtrl.Delete, Func: "Delete", Path: "/smstask/delete", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "删除短信发送记录,系统默认都是逻辑删除"},

	{Package: "rest", Module: "Smstask", HandlerFunc: smstaskCtrl.GetOne, Func: "GetOne", Path: "/smstask/getOne", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获取短信发送记录"},

	{Package: "rest", Module: "Userinfo", HandlerFunc: userinfoCtrl.Search, Func: "Search", Path: "/userinfo/search", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "RegisterRestCtrl(&Userinfo{})"},

	{Package: "rest", Module: "Userinfo", HandlerFunc: userinfoCtrl.Count, Func: "Count", Path: "/userinfo/count", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "搜索"},

	{Package: "rest", Module: "Userinfo", HandlerFunc: userinfoCtrl.Updatedeptandname, Func: "Updatedeptandname", Path: "/userinfo/updatedeptandname", Method: []string{"GET", "POST", "OPTIONS"}, Comment: ""},

	{Package: "rest", Module: "Userinfo", HandlerFunc: userinfoCtrl.UpdateUserinfo, Func: "UpdateUserinfo", Path: "/userinfo/updateUserinfo", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "搜索"},

	{Package: "rest", Module: "Userinfo", HandlerFunc: userinfoCtrl.Create, Func: "Create", Path: "/userinfo/create", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "创建默认用户"},

	{Package: "rest", Module: "Userinfo", HandlerFunc: userinfoCtrl.Update, Func: "Update", Path: "/userinfo/update", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "property. = model.RoleMember"},

	{Package: "rest", Module: "Userinfo", HandlerFunc: userinfoCtrl.Delete, Func: "Delete", Path: "/userinfo/delete", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "删除,系统默认都是逻辑删除"},

	{Package: "rest", Module: "Userinfo", HandlerFunc: userinfoCtrl.GetOne, Func: "GetOne", Path: "/userinfo/getOne", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获取"},

	{Package: "rest", Module: "Userinfo", HandlerFunc: userinfoCtrl.Snsinfo, Func: "Snsinfo", Path: "/userinfo/snsinfo", Method: []string{"GET", "POST", "OPTIONS"}, Comment: "获取"},
}
