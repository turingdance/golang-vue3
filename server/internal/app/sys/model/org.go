package model

import (
	"turingdance.com/reliable/internal/types"

	"gorm.io/gorm"
)

// 机构信息数据库模型
type Org struct {
	OrgId uint `json:"orgId" form:"orgId" gorm:"primaryKey;type:bigint not null  auto_increment ;comment:" `

	Title  string `json:"title" form:"title" gorm:"type:varchar(60);comment:名称" `
	Pic    string `json:"pic" form:"pic" gorm:"type:varchar(255);comment:logo" `
	Memo   string `json:"memo" form:"memo" gorm:"type:varchar(255);comment:sologo" `
	Status string `json:"status" form:"status" gorm:"type:varchar(10);default:NULL;comment:使能状态" `

	Linker string `json:"linker" form:"linker" gorm:"type:varchar(40);comment:联系人" `

	Content string `json:"content" form:"content" gorm:"type:text;comment:详细描述" `

	Tel string `json:"tel" form:"tel" gorm:"type:varchar(60);comment:联系方式" `

	Province string `json:"province" form:"province" gorm:"type:varchar(40);comment:省份" `

	City string `json:"city" form:"city" gorm:"type:varchar(255);comment:市" `

	District string `json:"district" form:"district" gorm:"type:varchar(255);comment:区" `

	Address string `json:"address" form:"address" gorm:"type:varchar(255);comment:详细地址" `

	Lat string `json:"lat" form:"lat" gorm:"type:varchar(20);comment:纬度" `

	Lng         string `json:"lng" form:"lng" gorm:"type:varchar(20);comment:经度" `
	Cate        string `json:"cate" form:"cate" gorm:"type:varchar(10);comment:类型" `
	OrgNo       string `json:"orgNo" form:"orgNo" gorm:"type:varchar(40);comment:统一社会信用代码" `
	UserId      string `json:"userId" form:"userId" gorm:"index;type:varchar(40);comment:机构创建者" `
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
