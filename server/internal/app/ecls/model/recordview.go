package model

import (
	"github.com/turingdance/infra/types"
	"gorm.io/datatypes"
)

// 上课记录数据库模型
type RecordView struct {
	Id int64 `json:"id" form:"id" gorm:"type:bigint;size:20;primaryKey;autoIncrement;comment:ID;"`

	CreateAt types.DateTime `json:"createAt" form:"createAt" gorm:"type:datetime;time_format:2006-01-02 15:04:05;time_utc:1;comment:上课时间;"`

	UserId string `json:"userId" form:"userId" gorm:"type:string;size:40;comment:用户ID;"`

	LessonId string `json:"lessonId" form:"lessonId" gorm:"type:string;size:40;comment:课程ID;"`

	CreateDate types.Date `json:"createDate" form:"createDate" gorm:"type:date;time_format:2006-01-02;time_utc:1;comment:创建日期;"`

	Media string `json:"media" form:"media" gorm:"type:varchar;size:10;comment:资源类型;"`

	Cover string `json:"cover" form:"cover" gorm:"type:varchar;size:255;comment:封面;"`

	Title string `json:"title" form:"title" gorm:"type:varchar;size:255;comment:标题;"`

	Memo string `json:"memo" form:"memo" gorm:"type:varchar;size:255;comment:描述;"`

	Uri  string         `json:"uri" form:"uri" gorm:"type:varchar;size:255;comment:资源key;"`
	Cate string         `json:"cate" form:"cate" gorm:"type:varchar;size:20;comment:分类;"`
	Tag  datatypes.JSON `json:"tag" form:"tag" gorm:"type:json;comment:标签;"`
	Read int            `json:"read" form:"read" gorm:"type:int(11);comment:阅读数;default:250"`
}

// TableName Record's table name
func (*RecordView) TableName() string {
	return "record_view"
}
