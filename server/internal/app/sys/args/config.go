package args

import (
	"gorm.io/gorm"
	"turingdance.com/turing/internal/types"
)

type Config struct {
	types.PageArg
	Total int `json:"total" form:"total"`

	Id uint `json:"id" form:"id"`

	Name string `json:"name" form:"name"`

	Label string `json:"label" form:"label"`

	Value string `json:"value" form:"value"`

	OrgId uint `json:"orgId" form:"orgId"`
}

// 分页
func (p *Config) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if p.Id > 0 {
			db = db.Where("id = ?", p.Id)
		}

		if len(p.Name) > 0 {
			db = db.Where("name = ?", p.Name)
		}

		if len(p.Value) > 0 {
			db = db.Where("value = ?", p.Value)
		}

		if p.OrgId > 0 {
			db = db.Where("org_id = ?", p.OrgId)
		}

		db = db.Where("deleted  = ?", p.Deleted)
		return db
	}
}

//结束
