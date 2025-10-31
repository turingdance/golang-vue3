// 资源管理
package args

import (
	"gorm.io/gorm"
	"turingdance.com/turing/internal/types"
)

type Resource struct {
	types.PageArg
	Total int `json:"total" form:"total"`

	ResId uint `json:"resId" form:"resId"`

	ResKey string `json:"resKey" form:"resKey"`

	ProgressId uint `json:"progressId" form:"progressId"`

	ResType string `json:"resType" form:"resType"`

	CreateAt types.DateTime `json:"createAt" form:"createAt"`
}

// 分页
func (p *Resource) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		if len(p.ResKey) > 0 {
			db = db.Where("resKey = ?", p.ResKey)
		}

		if p.ProgressId > 0 {
			db = db.Where("progress_id = ?", p.ProgressId)
		}

		if len(p.ResType) > 0 {
			db = db.Where("resType = ?", p.ResType)
		}

		db = db.Where("deleted  = ?", p.Deleted)
		return db
	}
}

//结束
