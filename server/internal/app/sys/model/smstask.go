package model

import (
	"turingdance.com/reliable/internal/types"

	"gorm.io/gorm"
)

// 短信发送记录数据库模型
type Smstask struct {
	Id uint `json:"id" form:"id" gorm:"primaryKey;type:bigint unsigned not null  auto_increment ;comment:ID" `

	Code string `json:"code" form:"code" gorm:"type:varchar(30);comment:验证码" `

	PhoneNumbers string `json:"phoneNumbers" form:"phoneNumbers" gorm:"type:varchar(1024);comment:手机号" `

	TempleteId string `json:"templeteId" form:"templeteId" gorm:"type:varchar(20);comment:模板ID" `

	Result string `json:"result" form:"result" gorm:"type:varchar(255);comment:发送结果" `

	CreateAt types.DateTime `json:"createAt" form:"createAt" time_format:"2006-01-02 15:04:05" time_utc:"1"   gorm:"type:datetime;comment:发布时间" `

	Success int `json:"success" form:"success" gorm:"type:integer;default:0;comment:是否成功" `

	SendAt types.DateTime `json:"sendAt" form:"sendAt" time_format:"2006-01-02 15:04:05" time_utc:"1"   gorm:"type:datetime;comment:发送时间" `
}

func (r Smstask) TableName() string {
	return "sys_smstask"
}
func init() {
	RegisterModel(&Smstask{})
}
func (m *Smstask) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateAt = types.DateTimeNow()
	m.SendAt = types.DateTimeNow()
	return
}
