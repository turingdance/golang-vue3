package model

import (
	"gorm.io/gorm"
	"turingdance.com/reliable/internal/types"
)

// 数据库模型
type Config struct {
	Id          uint   `json:"id" form:"id" gorm:"primaryKey;type:bigint(20) unsigned not null  auto_increment;comment:ID"`
	Name        string `json:"name" form:"name" gorm:"comment:名称;type:varchar(255)"`
	Label       string `json:"label" form:"label" gorm:"comment:展示;type:varchar(255)"`
	Value       string `json:"value" form:"value" gorm:"comment:值;type:varchar(1024)"`
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
