package args

import (
	"gorm.io/gorm"
	"turingdance.com/reliable/internal/types"
)

type AccountMethod string

const UseMobile AccountMethod = "mobile"
const UseUserName AccountMethod = "username"
const UseSmsCode AccountMethod = "smscode"
const UseWxcode AccountMethod = "wxcode"

// 账户相关的结构体
type Account struct {
	types.PageArg
	Username    string        `json:"username" form:"username"`
	Nickname    string        `json:"nickname" form:"nickname"`
	Realname    string        `json:"realname" form:"realname"`
	Mobile      string        `json:"mobile" form:"mobile"`
	Email       string        `json:"email" form:"email"`
	PasswordOld string        `json:"passwordOld" form:"passwordOld"`
	UserId      string        `json:"userId" form:"userId"`
	Pic         string        `json:"pic" form:"pic"`
	Password    string        `json:"password" form:"password"`
	UUid        string        `json:"uuid" form:"uuid"`
	OpenId      string        `json:"openId" form:"openId"`
	Code        string        `json:"code" form:"code"`
	Role        string        `json:"role" form:"role"`
	Token       string        `json:"token" form:"token"`
	Orgname     string        `json:"orgname" form:"orgname"`
	Tel         string        `json:"tel" form:"tel"`
	Linker      string        `json:"linker" form:"linker"`
	Type        AccountMethod `json:"type" form:"type"`
	Idcard      string        `json:"idcard" form:"idcard"`
	Autologin   bool          `json:"autologin" form:"autologin"`
}

// 分页
func (p *Account) Condtions() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !p.Datefrom.IsZero() {
			db = db.Where("create_at  >= ?", p.Datefrom.String())
		}
		if !p.Dateto.IsZero() {
			db = db.Where("create_at  < ?", p.Dateto.String())
		}
		db = db.Where("deleted  = ?", p.Deleted)

		return db
	}
}
func (a Account) UseMobile() bool {
	return a.Type == UseMobile
}

func (a Account) UseUsername() bool {
	return a.Type == UseUserName
}

func (a Account) UseSmsCode() bool {
	return a.Type == UseSmsCode
}
func (a Account) UseWxCode() bool {
	return a.Type == UseWxcode
}

func (a Account) UseType(t AccountMethod) bool {
	return a.Type == t
}

// 课件资源批量处理
type AccountKeyBatch struct {
	UserIds []string `json:"userIds"  form:"userIds"  errmsg:"缺少用戶ID"`
}
