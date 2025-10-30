package logic

import (
	"gorm.io/gorm"
	"turingdance.com/reliable/internal/app/ecls/model"
)

func ReadIt(who string, lessonId string, num int) (err error) {
	return DbEngin.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&model.Lesson{Id: lessonId}).UpdateColumn("read", gorm.Expr("`read` + ?", num)).Error
		if err != nil {
			return err
		}
		err = tx.Create(&model.Record{
			UserId: who, LessonId: lessonId,
		}).Error
		return err
	})

}
