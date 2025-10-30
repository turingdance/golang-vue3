// 短信发送记录
package args

import (
	"gorm.io/gorm"
	"turingdance.com/reliable/internal/types"
)

type Smstask struct {
	types.PageArg
	Total int `json:"total" form:"total"`

	Code string `json:"code" form:"code"`

	PhoneNumbers string `json:"phoneNumbers" form:"phoneNumbers"`

	TempleteId string `json:"templeteId" form:"templeteId"`

	Result string `json:"result" form:"result"`

	SendAt types.DateTime `json:"sendAt" form:"sendAt" time_format:"2006-01-02 15:04:05" time_utc:"1" `
}

// 分页
func (p *Smstask) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !p.Datefrom.IsZero() {
			db = db.Where("create_at  >= ?", p.Datefrom.String())
		}
		if !p.Dateto.IsZero() {
			db = db.Where("create_at  < ?", p.Dateto.String())
		}
		db = db.Where("deleted  = ?", p.Deleted)

		return db
	}
}

//结束
