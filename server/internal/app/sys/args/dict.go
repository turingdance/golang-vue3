// 字典
package args

import (
	"gorm.io/gorm"
	"turingdance.com/reliable/internal/types"
)

type Dict struct {
	types.PageArg
	Total int `json:"total" form:"total"`

	Id int `json:"id" form:"id"`

	Name string `json:"name" form:"name"`

	Title string `json:"title" form:"title"`

	Value string `json:"value" form:"value"`
}

// 分页
func (p *Dict) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		db = db.Where("deleted  = ?", p.Deleted)

		return db
	}
}

//结束
