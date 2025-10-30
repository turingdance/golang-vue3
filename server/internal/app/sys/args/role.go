package args

import (
	"gorm.io/gorm"
	"turingdance.com/reliable/internal/types"
)

type Role struct {
	types.PageArg
	Total int `json:"total" form:"total"`

	Id uint `json:"id" form:"id"`

	Name  string `json:"name" form:"name"`
	Code  string `json:"code" form:"code"`
	State *int   `json:"state" form:"state"`
}

// 分页
func (p *Role) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if p.Name != "" {
			db = db.Where("name like ?", "%"+p.Name+"%")
		}
		if p.Code != "" {
			db = db.Where("code like ?", "%"+p.Code+"%")
		}
		if p.State != nil {
			db = db.Where("state = ?", *p.State)
		}
		return db
	}
}

//结束
