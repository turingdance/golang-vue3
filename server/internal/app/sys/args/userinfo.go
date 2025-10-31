package args

import (
	"gorm.io/gorm"
	"turingdance.com/turing/internal/types"
)

type Userinfo struct {
	types.PageArg
	Total int `json:"total" form:"total"`

	UserId string `json:"userId" form:"userId"`

	Username string `json:"username" form:"username"`

	Password string `json:"password" form:"password"`

	Pic string `json:"pic" form:"pic"`

	Nickname string `json:"nickname" form:"nickname"`

	Mobile string `json:"mobile" form:"mobile"`

	Email string `json:"email" form:"email"`

	Status int `json:"status" form:"status"`

	CreateAt types.DateTime `json:"createAt" form:"createAt" time_format:"2006-01-02 15:04:05" time_utc:"1" `

	RoleId  uint     `json:"roleId" form:"roleId"`
	OrgId   uint     `json:"orgId" form:"orgId"`
	DeptId  string   `json:"deptId" form:"deptId"`
	DeptIds []string `json:"deptIds" form:"deptIds"`
}

// 分页
func (p *Userinfo) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !p.Datefrom.IsZero() {
			db = db.Where("create_at  >= ?", p.Datefrom.String())
		}
		if !p.Dateto.IsZero() {
			db = db.Where("create_at  < ?", p.Dateto.String())
		}
		if p.RoleId > 0 {
			db = db.Where("role_id  = ?", p.RoleId)
		}

		if p.DeptId != "" {
			db = db.Where("dept_id  = ?", p.DeptId)
		}
		if len(p.DeptIds) > 0 {
			db = db.Where("dept_id  in ?", p.DeptIds)
		}
		db = db.Where("deleted  = ?", p.Deleted)

		return db
	}
}

//结束
