package model

import (
	"errors"

	"gorm.io/gorm"
)

// 数据库模型
type Role struct {
	Id       uint     `json:"id" form:"id" gorm:"primaryKey;type:bigint not null  auto_increment ;comment:角色ID"`
	Name     string   `json:"name" form:"name" gorm:"comment:角色名称;type:varchar(40)"`
	Code     RoleType `json:"code" form:"code" gorm:"comment:索引admin/patient/worker;type:varchar(20)"`
	RightIds []uint   `json:"rightIds" form:"rightIds" gorm:"serializer:json;comment:权限列表;type:text;"`
	State    int      `json:"state" form:"state" gorm:"type:int(11);default:1;comment:状态1可用,0不可用"`
}
type RoleType string

func (r Role) TableName() string {
	return "sys_role"
}
func init() {
	RegisterModel(&Role{})
}
func (m *Role) BeforeCreate(tx *gorm.DB) (err error) {
	if m.RightIds == nil {
		m.RightIds = make([]uint, 0)
	}
	if m.Code == "" {
		return errors.New("缺少编码")
	}
	if m.Name == "" {
		return errors.New("缺少名称")
	}
	m.State = 1
	return
}

const (
	// 会员
	RoleGuest  RoleType = "guest"  //普通游客
	RoleMember RoleType = "member" //普通用户
	RoleTenant RoleType = "tenant" // 注册租户
	RoleAdmin  RoleType = "admin"  // 管理员用户
)
