package model

import (
	"errors"

	"github.com/turingdance/infra/types"
	"gorm.io/gorm"
)

// 创建去的钩子
func (m *Record) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreateAt = types.DateTimeNow()
	m.CreateDate = types.DateNow()
	if m.LessonId == "" {
		err = errors.New("缺少课程ID")
	}
	if m.UserId == "" {
		err = errors.New("缺少用户ID")
	}
	return
}
