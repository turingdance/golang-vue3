// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git
// @author winlion@turingdance.com

package model

import (
	"github.com/turingdance/infra/types"
)

const TableNameRecord = "ecls_record"
const TableTitleRecord = "上课记录"
const TableRecord = "record"

var recordKeys []string = []string{"id", "createAt", "userId", "lessonId", "createDate"}

// 上课记录数据库模型
type Record struct {
	Id int64 `json:"id" form:"id" gorm:"type:bigint;size:20;primaryKey;autoIncrement;comment:ID;"`

	CreateAt types.DateTime `json:"createAt" form:"createAt" gorm:"type:datetime;time_format:2006-01-02 15:04:05;time_utc:1;comment:上课时间;"`

	UserId string `json:"userId" form:"userId" gorm:"type:string;size:40;comment:用户ID;"`

	LessonId string `json:"lessonId" form:"lessonId" gorm:"type:string;size:40;comment:课程ID;"`

	CreateDate types.Date `json:"createDate" form:"createDate" gorm:"type:date;time_format:2006-01-02;time_utc:1;comment:创建日期;"`
}

// TableName Record's table name
func (*Record) TableName() string {
	return TableNameRecord
}

func init() {
	RegisterModel(&Record{})
}
