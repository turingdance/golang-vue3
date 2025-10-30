package model

import (
	"gorm.io/gorm"
	"turingdance.com/reliable/internal/types"
)

// 区域数据库模型
type Area struct {
	AreaId uint `json:"areaId" form:"areaId" gorm:"primaryKey;comment:areaId"` //areaId

	Pid uint `json:"pid" form:"pid" gorm:"index;comment:父级ID"` //父级ID

	Name string `json:"name" form:"name" gorm:"type:string;size:120;comment:名称"` //名称

	Type string `json:"type" form:"type" gorm:"index;type:string;size:40;;comment:类型"` //类型

	Rank        int `json:"rank" form:"rank" gorm:"index;type:int;comment:级别"` //级别
	types.Model `gorm:"embedded"`
}

// 模型注册后系统能自动处理表变动
func init() {
	RegisterModel(&Area{})
}

// 表名
func (m Area) TableName() string {
	return "sys_area"
}

// 创建去的钩子
func (m *Area) BeforeCreate(tx *gorm.DB) (err error) {

	m.CreateAt = types.DateTimeNow()
	m.DeleteAt = types.DateTimeNow()
	m.Deleted = 0
	return
}

// 创建去的钩子
func (m *Area) BeforeSave(tx *gorm.DB) (err error) {

	m.CreateAt = types.DateTimeNow()
	m.DeleteAt = types.DateTimeNow()
	m.Deleted = 0
	return
}

// 创建去的钩子
func (m *Area) BeforeDelete(tx *gorm.DB) (err error) {
	m.DeleteAt = types.DateTimeNow()
	m.Deleted = 1
	return
}
