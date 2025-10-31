package model

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
	"turingdance.com/turing/internal/types"
)

type DictValue struct {
	Json []types.DataItem `json:"json"`
}

func DefaultDictValue() DictValue {
	return DictValue{
		Json: []types.DataItem{
			{
				Value: "", Label: "",
			},
		},
	}
}
func (ds *DictValue) Scan(val interface{}) error {
	if val == nil || *&val == "" {
		*ds = DefaultDictValue()
		return nil
	} else {
		b, _ := val.([]byte)
		return json.Unmarshal(b, ds)
	}
}

func (ds DictValue) Value() (driver.Value, error) {
	return json.Marshal(ds)
}

// 字典数据库模型
type Dict struct {
	Id int `json:"id" form:"id" gorm:"primaryKey;type:unit;primaryKey;autoIncrement;comment:ID" `

	Name string `json:"name" form:"name" gorm:"type:string;size:255;comment:名称" `

	Title string `json:"title" form:"title" gorm:"type:string;size:255;comment:名字" `

	Value       []types.DataItem `json:"value" form:"value" gorm:"serializer:json;type:text;comment:值得" `
	types.Model `gorm:"embedded"`
}

func (r Dict) TableName() string {
	return "sys_dict"
}

func (m *Dict) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateAt = types.DateTimeNow()
	m.DeleteAt = types.DateTimeNow()
	m.Deleted = 0
	return
}

// 模型注册后系统能自动处理表变动
func init() {
	RegisterModel(&Dict{})
}
