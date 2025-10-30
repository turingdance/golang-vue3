package model

import (
	"fmt"

	"turingdance.com/reliable/internal/pkg/utils"

	"gorm.io/gorm"
)

// 机构信息数据库模型
type Dept struct {
	Id       string   `json:"id" form:"id" gorm:"primaryKey;type:string;size:40;comment:部门ID" `
	Title    string   `json:"title" form:"title" gorm:"type:varchar(60);comment:名称名称" `
	Pid      string   `json:"pid" form:"pid" gorm:"type:string;size:40;index;comment:父级ID" `
	MasterId string   `json:"masterId" form:"masterId" gorm:"type:string;size:40;index;comment:部门负责人ID"`
	Userinfo Userinfo `json:"userinfo" form:"userinfo" gorm:"foreignkey:MasterId;references:UserId"`
	Status   int      `json:"status" form:"status" gorm:"type:int(11);default:1;comment:部门状态" `
	Level    int      `json:"level" form:"level" gorm:"type:int(11);default:1;comment:部门等级" `
}

func (r Dept) TableName() string {
	return "sys_dept"
}

// 模型注册后系统能自动处理表变动
func init() {
	RegisterModel(&Dept{})
}

func (m *Dept) BeforeCreate(tx *gorm.DB) (err error) {
	m.Status = 1
	m.Level = 1
	if m.Id == "" {
		m.Id = utils.PKID()
	}
	if m.Title == "" {
		err = fmt.Errorf("缺少部门名称")
	}
	if m.MasterId == "" {
		err = fmt.Errorf("缺少部门负责人")
	}
	return
}
