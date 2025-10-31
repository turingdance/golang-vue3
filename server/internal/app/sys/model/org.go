package model

import (
	"turingdance.com/turing/internal/types"

	"gorm.io/gorm"
)

// 机构信息数据库模型
type Org struct {
	OrgId uint `json:"orgId" form:"orgId" gorm:"primaryKey;type:bigint not null  auto_increment ;comment:" `

	Title  string `json:"title" form:"title" gorm:"type:string;size:60;comment:名称" `
	Pic    string `json:"pic" form:"pic" gorm:"type:string;size:255;comment:logo" `
	Memo   string `json:"memo" form:"memo" gorm:"type:string;size:255;comment:sologo" `
	Status string `json:"status" form:"status" gorm:"type:string;size:40;default:NULL;comment:使能状态" `

	Linker string `json:"linker" form:"linker" gorm:"type:string;size:40;comment:联系人" `

	Content string `json:"content" form:"content" gorm:"type:text;comment:详细描述" `

	Tel string `json:"tel" form:"tel" gorm:"type:string;size:60;comment:联系方式" `

	Province string `json:"province" form:"province" gorm:"type:string;size:40;comment:省份" `

	City string `json:"city" form:"city" gorm:"type:string;size:255;comment:市" `

	District string `json:"district" form:"district" gorm:"type:string;size:255;comment:区" `

	Address string `json:"address" form:"address" gorm:"type:string;size:255;comment:详细地址" `

	Lat string `json:"lat" form:"lat" gorm:"type:string;size:20;comment:纬度" `

	Lng         string `json:"lng" form:"lng" gorm:"type:string;size:20;comment:经度" `
	Cate        string `json:"cate" form:"cate" gorm:"type:string;size:10;comment:类型" `
	OrgNo       string `json:"orgNo" form:"orgNo" gorm:"type:string;size:40;comment:统一社会信用代码" `
	UserId      string `json:"userId" form:"userId" gorm:"index;type:string;size:40;comment:机构创建者" `
	types.Model `gorm:"embedded"`
}

func (r Org) TableName() string {
	return "sys_org"
}

// 模型注册后系统能自动处理表变动
func init() {
	RegisterModel(&Org{})
}

func (m *Org) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateAt = types.DateTimeNow()
	m.Deleted = 0
	m.DeleteAt = types.DateTimeNow()
	return
}
