package args

import (
	"gorm.io/gorm"
	"turingdance.com/reliable/internal/types"
)

type Rights struct {
	types.PageArg

	Type string `json:"type" form:"type"`
}

// 分页
func (p *Rights) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if p.Type != "" {
			db = db.Where("type = ?", p.Type)
		}
		return db
	}
}

//结束
