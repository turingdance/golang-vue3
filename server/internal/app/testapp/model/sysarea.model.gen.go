// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git
// @author winlion

package model

import 
(
	
	"github.com/turingdance/infra/types"
	
	
	"gorm.io/gorm"
)

const TableNameSysArea = "sys_area"
const TableTitleSysArea = "区域信息"
const TableSysArea = "sysArea"
var sysAreaKeys []string= []string{"areaId","pid","name","type","rank","createAt","deleteAt","deleted", }
//区域信息数据库模型
type SysArea struct{
        AreaId int64 `json:"areaId" form:"areaId" gorm:"type:bigint;size:20;primaryKey;autoIncrement;"`
    
        Pid int64 `json:"pid" form:"pid" gorm:"type:bigint;size:19;index;comment:父级ID;"`
    
        Name string `json:"name" form:"name" gorm:"type:string;size:120;comment:名称;"`
    
        Type string `json:"type" form:"type" gorm:"type:string;size:40;index;comment:类型;"`
    
        Rank int64 `json:"rank" form:"rank" gorm:"type:bigint;size:19;index;comment:级别;"`
    
        CreateAt types.DateTime `json:"createAt" form:"createAt" gorm:"type:time;time_format:2006-01-02 15:04:05;time_utc:1;comment:创建时间;"`
    
        DeleteAt types.DateTime `json:"deleteAt" form:"deleteAt" gorm:"type:time;time_format:2006-01-02 15:04:05;time_utc:1;comment:删除时间;"`
    
        Deleted int32 `json:"deleted" form:"deleted" gorm:"type:int;size:10;comment:是否删除;"`
    
}

// TableName SysArea's table name
func (*SysArea) TableName() string {
	return TableNameSysArea
}
func (obj *SysArea) BeforeCreate(tx *gorm.DB) (err error) {
		
		
	
		
		
	
		
		
	
		
		
	
		
		
	
		
		
	
		
		
	
		
		
	
	return
}
func (obj *SysArea) BeforeUpdate(tx *gorm.DB) (err error) {
		
	
		
	
		
	
		
	
		
	
		
	
		
	
		
	
	return
}
func init() {
	RegisterModel(&SysArea{})
}

