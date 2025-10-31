package model

import (
	"gorm.io/gorm"
	"turingdance.com/turing/internal/types"
)

// 资源管理数据库模型
type Resource struct {
	ResId uint `json:"resId" form:"resId" gorm:"res_id bigint(20) unsigned;primaryKey;autoIncrement;comment:资源ID"` //资源ID

	ResKey string `json:"resKey" form:"resKey" gorm:"res_key varchar(255);comment:资源ID"` //资源ID

	InstanceId string `json:"instanceId" form:"instanceId" gorm:"instanceId;type:string;size:40;comment:绑定流程ID"` //绑定流程ID

	ResType string `json:"resType" form:"resType" gorm:"res_type varchar(20);comment:类型"` //类型

	CreateAt types.DateTime `json:"createAt" form:"createAt" gorm:"create_at datetime;comment:状态"` //状态

	types.Model `gorm:"embedded"`
}

// 模型注册后系统能自动处理表变动
func init() {
	RegisterModel(&Resource{})
}

// 表名
func (m Resource) TableName() string {
	return "doc_resource"
}

// 创建去的钩子
func (m *Resource) BeforeCreate(tx *gorm.DB) (err error) {

	m.CreateAt = types.DateTimeNow()
	m.DeleteAt = types.DateTimeNow()
	m.Deleted = 0
	return
}

// 创建去的钩子
func (m *Resource) BeforeSave(tx *gorm.DB) (err error) {

	m.CreateAt = types.DateTimeNow()
	m.DeleteAt = types.DateTimeNow()
	m.Deleted = 0
	return
}

// 创建去的钩子
func (m *Resource) BeforeDelete(tx *gorm.DB) (err error) {
	m.DeleteAt = types.DateTimeNow()
	m.Deleted = 1
	return
}
