// 区域
package args

import (
	"gorm.io/gorm"
	"turingdance.com/turing/internal/types"
)

type Area struct {
	types.PageArg
	Total int `json:"total" form:"total"`

	AreaId uint `json:"areaId" form:"areaId"`

	Pid uint `json:"pid" form:"pid"`

	Name string `json:"name" form:"name"`

	Type string `json:"type" form:"type"`

	Rank int `json:"rank" form:"rank"`
}

// 分页
func (p *Area) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if p.Pid > 0 {
			db = db.Where("pid = ?", p.Pid)
		}

		if len(p.Name) > 0 {
			db = db.Where("name like ?", "%"+p.Name+"%")
		}

		if len(p.Type) > 0 {
			db = db.Where("type = ?", p.Type)
		}

		if p.Rank > 0 {
			db = db.Where("rank = ?", p.Rank)
		}

		db = db.Where("deleted  = ?", p.Deleted)
		return db
	}
}

//结束
