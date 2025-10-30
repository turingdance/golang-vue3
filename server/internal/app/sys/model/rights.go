package model

import (
	"errors"

	"gorm.io/gorm"
)

// 资源管理数据库模型
type Rights struct {
	Id         uint       `json:"id" form:"id" gorm:"id bigint(20) unsigned;primaryKey;autoIncrement;comment:ID"` //资源ID
	Biz        string     `json:"biz" form:"biz" gorm:"column:biz;uniqueIndex;type:string;size:40;comment:业务标识符"` //资源ID
	Title      string     `json:"title" form:"title" gorm:"column:title;type:string;size:40;comment:权限名称"`        //资源ID
	Pid        uint       `json:"pid" form:"pid" gorm:"column:pid;type:bigint;size:20;comment:父级"`                //资源ID
	Uri        string     `json:"uri" form:"uri" gorm:"column:uri;type:string;size:40;comment:资源路径"`
	Type       RightsType `json:"type" form:"type" gorm:"column:type;type:string;size:40;comment:资源类型"`
	SortIndex  float64    `json:"sortIndex" form:"sortIndex" gorm:"column:sort_index;type:decimal;size:10;default:0;comment:排序位"`
	Component  string     `json:"component" form:"component" gorm:"column:component;type:string;size:240;comment:組件路径"`
	Icon       string     `json:"icon" form:"icon" gorm:"column:icon;type:string;size:20;comment:图标"`
	Hidden     bool       `json:"hidden" form:"hidden" gorm:"column:hidden;type:int;size:11;comment:是否隐藏"`
	AlwaysShow bool       `json:"alwaysShow" form:"alwaysShow" gorm:"column:always_show;type:int;size:11;comment:是否始终展示"`
	Redirect   string     `json:"redirect" form:"redirect" gorm:"column:redirect;type:string;size:140;comment:如果跳转"`
}
type RightsType string

const (
	RightsTypeGroup  RightsType = "group"  //分组
	RightsTypeView   RightsType = "view"   //页面
	RightsTypeWidget RightsType = "widget" //按钮
	RightsTypeApi    RightsType = "api"    //接口级
)

// 模型注册后系统能自动处理表变动
func init() {
	RegisterModel(&Rights{})
}

// 表名
func (m Rights) TableName() string {
	return "sys_rights"
}

// 创建去的钩子
func (m *Rights) BeforeCreate(tx *gorm.DB) (err error) {
	if m.Type == RightsTypeView {
		if m.Component == "" {
			return errors.New("缺少组件路径")
		}
	}
	return
}

// 创建去的钩子
func (m *Rights) BeforeSave(tx *gorm.DB) (err error) {

	return
}

// 创建去的钩子
func (m *Rights) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
