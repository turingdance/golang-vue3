package model

import (
	"turingdance.com/turing/internal/pkg/utils"
	"turingdance.com/turing/internal/types"

	"gorm.io/gorm"
)

// 租户信息数据库模型
type Tenant struct {
	Id          string `json:"id" form:"id" gorm:"primaryKey;type:varchar(32);comment:租户ID,32位字符串"`
	UserId      string `json:"userId" form:"userId" gorm:"type:varchar(32);index;comment:创建者,谁创建的这个"`
	Name        string `json:"name" form:"name" gorm:"type:varchar(128);comment:客户名称,3‑64位"`
	ShortName   string `json:"shortName" form:"shortName" gorm:"type:varchar(64);default:NULL;comment:简称,比如中行"`
	AppKey      string `json:"appKey" form:"appKey" gorm:"type:varchar(48);comment:AppKey,授权码访问AppKey"`
	AppSecret   string `json:"appSecret" form:"appSecret" gorm:"type:varchar(255);default:NULL;comment:AppSecret,授权码访问密钥"`
	SvcpkgId    string `json:"svcpkgId" form:"svcpkgId" gorm:"type:varchar(32);default:NULL;comment:套餐ID,当前使用的服务"`
	MaxTask     int    `json:"maxTask" form:"maxTask" gorm:"type:int(11);default:NULL;comment:任务上限,并发处理任务数上限,默认2"`
	MaxUsers    int    `json:"maxUsers" form:"maxUsers" gorm:"type:int(11);default:NULL;comment:账号上线,员工账号上限"`
	MaxModels   int    `json:"maxModels" form:"maxModels" gorm:"type:int(11);default:NULL;comment:私有模型上限,正整数"`
	ExpiresAt   string `json:"expiresAt" form:"expiresAt" gorm:"type:datetime(3);default:NULL;comment:到期时间,超过这个时间账号不可用"`
	Ctname      string `json:"ctname" form:"ctname" gorm:"type:varchar(64);default:NULL;comment:联系人,姓名"`
	Ctphone     string `json:"ctphone" form:"ctphone" gorm:"type:varchar(32);default:NULL;comment:联系电话,手机号"`
	Ctemail     string `json:"ctemail" form:"ctemail" gorm:"type:varchar(128);default:NULL;comment:联系邮箱,联系邮箱"`
	Province    string `json:"province" form:"province" gorm:"type:varchar(32);default:NULL;comment:省"`
	City        string `json:"city" form:"city" gorm:"type:varchar(32);default:NULL;comment:市"`
	Address     string `json:"address" form:"address" gorm:"type:varchar(512);default:NULL;comment:详细地址"`
	Remark      string `json:"remark" form:"remark" gorm:"type:text;default:NULL;comment:备注"`
	Status      int    `json:"status" form:"status" gorm:"type:tinyint(4);default:NULL;comment:租户状态,1=激活 2=过期 3=停用"`
	CertNo      string `json:"certNo" form:"certNo" gorm:"type:varchar(32);default:NULL;comment:机构编码,组织机构代码证"`
	types.Model `gorm:"embedded"`
}

func (r Tenant) TableName() string {
	return "sys_tenant"
}

// 模型注册后系统能自动处理表变动
func init() {
	RegisterModel(&Tenant{})
}

// 创建前钩子，完全复刻 Tenant 的逻辑
func (m *Tenant) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateAt = types.DateTimeNow()
	m.Deleted = 0
	m.DeleteAt = types.DateTimeNow()
	if m.Id == "" {
		m.Id = utils.NextIntString()
	}
	return
}
