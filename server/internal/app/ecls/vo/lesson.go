// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git
// @author winlion@turingdance.com

package vo

import "turingdance.com/reliable/internal/app/ecls/model"

//课件资源参数请求接口
type Lesson struct {
	Id string `json:"id"  form:"id"  validate:"required"`

	Media string `json:"media"  form:"media"  validate:"required"`

	Cover string `json:"cover"  form:"cover"  validate:"required"`

	Title string `json:"title"  form:"title"  validate:"required"`

	Memo string `json:"memo"  form:"memo"  validate:"required"`

	Uri string `json:"uri"  form:"uri"  validate:"required"`

	Tag string `json:"tag"  form:"tag"  validate:"required"`
}

// 课件资源批量处理
type LessonKeyBatch struct {
	Ids []string `json:"ids"  form:"ids"  errmsg:"缺少视频ID"`
}

// 课件资源批量处理
type AddLessonBatch struct {
	Data []model.Lesson `json:"data"`
}
