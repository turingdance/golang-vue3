package model

import (
	"gorm.io/gorm"
	"turingdance.com/turing/internal/types"
)

// 数据库模型
type Config struct {
	Id          uint   `json:"id" form:"id" gorm:"primaryKey;type:unit;primaryKey;autoIncrement;comment:ID"`
	Name        string `json:"name" form:"name" gorm:"comment:名称;type:string;size:255"`
	Label       string `json:"label" form:"label" gorm:"comment:展示;type:string;size:255"`
	Value       string `json:"value" form:"value" gorm:"comment:值;type:string;size:1024"`
	OrgId       uint   `json:"orgId" form:"orgId" gorm:"index;comment:组织ID"`
	types.Model `gorm:"embedded"`
}

// 模型注册后系统能自动处理表变动
func init() {
	RegisterModel(&Config{})
}
func (r Config) TableName() string {
	return "sys_config"
}
func (m *Config) BeforeCreate(tx *gorm.DB) (err error) {

	m.Deleted = 0
	return
}
