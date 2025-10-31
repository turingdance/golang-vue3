package model

import (
	"encoding/json"

	"github.com/turingdance/infra/pkkit"
	"gorm.io/gorm"
	"turingdance.com/turing/internal/types"
)

// 统一流程了
// 首先注册系统账号
// 然后完善资料
// 完成平台注册
type Userinfo struct {
	UserId      string             `json:"userId" form:"userId" gorm:"primaryKey;type:varchar(40) not null ;comment:用户ID" `
	Username    string             `json:"username" form:"username" gorm:"type:varchar(40);uniqueIndex;comment:用户名" `
	Mobile      string             `json:"mobile" form:"mobile" gorm:"type:varchar(20);index;comment:手机号" `
	Password    string             `json:"-" form:"password" gorm:"type:varchar(128);comment:密码" `
	Nickname    string             `json:"nickname" form:"nickname" gorm:"type:varchar(40);comment:昵称" `
	Pic         string             `json:"pic" form:"pic" gorm:"type:varchar(256);comment:头像" `
	OrgId       uint               `json:"orgId" form:"orgId" gorm:"-"`
	RoleId      uint               `json:"roleId" form:"roleId" gorm:"type:bigint;size:20;comment:角色ID"`
	Role        Role               `json:"role" gorm:"foreignkey:RoleId;references:Id"`
	Scope       []string           `json:"scope" gorm:"-"` // 权限
	DeptId      string             `json:"deptId" form:"deptId" gorm:"index;type:string;size:40;comment:所属部门"`
	Status      types.EnableStatus `json:"status" form:"status" gorm:"type:varchar(11);default:1;comment:是否可用" `
	types.Model `gorm:"embedded"`
}

func init() {
	RegisterModel(&Userinfo{})
}
func (r Userinfo) TableName() string {
	return "sys_userinfo"
}
func (u Userinfo) Empty() bool {
	return u.UserId == ""
}
func (m *Userinfo) Desensitize() (result Userinfo) {
	result = Userinfo{}
	bts, _ := json.Marshal(m)
	json.Unmarshal(bts, &result)
	result.Password = ""

	return result
}
func (m *Userinfo) PKID() (userId string) {
	m.UserId = pkkit.UseSnowflakeID()
	return m.UserId
}

func (m *Userinfo) BeforeCreate(tx *gorm.DB) (err error) {
	if m.UserId == "" {
		m.UserId = pkkit.UseSnowflakeID()
	}
	m.Deleted = 0

	m.CreateAt = types.DateTimeNow()
	m.DeleteAt = types.DateTimeNow()
	// 初始化注册是不可用状态
	m.Status = types.EnableStatusFalse
	return
}

func NewUserInfo(role RoleType) *Userinfo {
	return &Userinfo{}
}
