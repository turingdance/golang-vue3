// 机构信息
package args

import (
	"gorm.io/gorm"
	"turingdance.com/reliable/internal/types"
)

type Org struct {
	types.PageArg
	Total int `json:"total" form:"total"`

	OrgId uint `json:"orgId" form:"orgId"`

	Title string `json:"title" form:"title"`

	Status int `json:"status" form:"status"`

	CreateAt types.DateTime `json:"createAt" form:"createAt" time_format:"2006-01-02 15:04:05" time_utc:"1" `

	Linker string `json:"linker" form:"linker"`

	Pic string `json:"pic" form:"pic"`

	Memo string `json:"memo" form:"memo"`

	Content string `json:"content" form:"content"`

	Tel string `json:"tel" form:"tel"`

	Province string `json:"province" form:"province"`

	City string `json:"city" form:"city"`

	District string `json:"district" form:"district"`

	Address string `json:"address" form:"address"`

	Lat  string `json:"lat" form:"lat"`
	Cate string `json:"cate" form:"cate"`
	Lng  string `json:"lng" form:"lng"`

	OrgNo string `json:"orgNo" form:"orgNo"`

	Deleted int `json:"deleted" form:"deleted"`

	UserId string `json:"userId" form:"userId"`

	DeleteAt types.DateTime `json:"deleteAt" form:"deleteAt" time_format:"2006-01-02 15:04:05" time_utc:"1" `
}

// 分页
func (p *Org) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !p.Datefrom.IsZero() {
			db = db.Where("create_at  >= ?", p.Datefrom.String())
		}
		if !p.Dateto.IsZero() {
			db = db.Where("create_at  < ?", p.Dateto.String())
		}
		if p.Cate != "" {
			db = db.Where("cate  = ?", p.Cate)
		}
		db = db.Where("deleted  = ?", p.Deleted)

		return db
	}
}

//结束
