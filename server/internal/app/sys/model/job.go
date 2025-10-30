package model

import (
	"gorm.io/gorm"
	"turingdance.com/reliable/internal/types"
)

// 岗位数据库模型
type Job struct {
	JobId uint `json:"jobId" form:"jobId" gorm:"job_id bigint(20) unsigned;primaryKey;comment:ID"` //ID

	Title string `json:"title" form:"title" gorm:"title varchar(255);comment:岗位名称"` //岗位名称

	Rank int `json:"rank" form:"rank" gorm:"rank int(11);comment:级别"` //级别

	types.Model `gorm:"embedded"`
}

// 模型注册后系统能自动处理表变动
func init() {
	RegisterModel(&Job{})
}

// 表名
func (m Job) TableName() string {
	return "doc_job"
}

// 创建去的钩子
func (m *Job) BeforeCreate(tx *gorm.DB) (err error) {

	m.CreateAt = types.DateTimeNow()
	m.DeleteAt = types.DateTimeNow()
	m.Deleted = 0
	return
}

// 创建去的钩子
func (m *Job) BeforeSave(tx *gorm.DB) (err error) {

	m.CreateAt = types.DateTimeNow()
	m.DeleteAt = types.DateTimeNow()
	m.Deleted = 0
	return
}

// 创建去的钩子
func (m *Job) BeforeDelete(tx *gorm.DB) (err error) {
	m.DeleteAt = types.DateTimeNow()
	m.Deleted = 1
	return
}
