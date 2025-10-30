// 岗位
package args

import (
	"gorm.io/gorm"
	"turingdance.com/reliable/internal/types"
)

type Dept struct {
	types.PageArg
	Total int `json:"total" form:"total"`

	Title string `json:"title" form:"title"`
	Pid   string `json:"pid" form:"pid"`
	Level int    `json:"level" form:"level"`
}

// 分页
func (p *Dept) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if len(p.Title) > 0 {
			db = db.Where("title like ?", "%"+p.Title+"%")
		}

		if p.Level > 0 {
			db = db.Where("level = ?", p.Level)
		}
		if p.Pid != "" {
			db = db.Where("pid = ?", p.Pid)
		}
		return db
	}
}

//结束
