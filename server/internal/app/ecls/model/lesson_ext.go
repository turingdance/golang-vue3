// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git
// @author winlion@turingdance.com

package model

import (
	"github.com/turingdance/infra/pkkit"
	"gorm.io/gorm"
)

// 创建去的钩子
func (m *Lesson) BeforeCreate(tx *gorm.DB) (err error) {

	if m.Id == "" {
		m.Id = pkkit.UseSnowflakeID()
	}
	return
}

// 创建去的钩子
func (m *Lesson) BeforeSave(tx *gorm.DB) (err error) {

	return
}

// 创建去的钩子
func (m *Lesson) BeforeDelete(tx *gorm.DB) (err error) {

	return
}
