// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git
// @author winlion@turingdance.com

package model

import "gorm.io/datatypes"

const TableNameLesson = "ecls_lesson"
const TableTitleLesson = "课件资源"
const TableLesson = "lesson"

//课件资源数据库模型
type Lesson struct {
	Id string `json:"id" form:"id" gorm:"type:varchar;size:32;primaryKey;comment:视频ID;"`

	Media string `json:"media" form:"media" gorm:"type:varchar;size:10;comment:资源类型;"`

	Cover string `json:"cover" form:"cover" gorm:"type:varchar;size:255;comment:封面;"`

	Title string `json:"title" form:"title" gorm:"type:varchar;size:255;comment:标题;"`

	Memo string `json:"memo" form:"memo" gorm:"type:varchar;size:255;comment:描述;"`

	Uri  string         `json:"uri" form:"uri" gorm:"type:varchar;size:255;comment:资源key;"`
	Cate string         `json:"cate" form:"cate" gorm:"type:varchar;size:20;comment:分类;"`
	Tag  datatypes.JSON `json:"tag" form:"tag" gorm:"type:json;comment:标签;"`
	Read int            `json:"read" form:"read" gorm:"type:int(11);comment:阅读数;default:250"`
}

// TableName Lesson's table name
func (*Lesson) TableName() string {
	return TableNameLesson
}

func init() {
	RegisterModel(&Lesson{})
}
