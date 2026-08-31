// 机构信息
package args

import (
	"gorm.io/gorm"
	"turingdance.com/turing/internal/types"
)

type Tenant struct {
	types.PageArg

	TenantId string `json:"tenantId" form:"tenantId"`

	Name string `json:"name" form:"name"`

	Status int `json:"status" form:"status"`

	Province string `json:"province" form:"province"`

	City string `json:"city" form:"city"`

	District string `json:"district" form:"district"`

	Address string `json:"address" form:"address"`

	Cate string `json:"cate" form:"cate"`

	CertNo string `json:"certNo" form:"certNo"`

	UserId string `json:"userId" form:"userId"`
}

// 分页
func (p *Tenant) Condtions() func(db *gorm.DB) *gorm.DB {
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
