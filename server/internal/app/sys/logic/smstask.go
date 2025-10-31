package logic

//短信发送记录
import (
	"turingdance.com/turing/internal/app/sys/args"
	"turingdance.com/turing/internal/app/sys/model"

	"turingdance.com/turing/internal/types"
)

// 创建短信发送记录
func CreateSmstask(property model.Smstask) (result model.Smstask, err error) {
	err = DbEngin.Create(&property).Error
	return property, err
}

// 更新短信发送记录
func UpdateSmstask(property model.Smstask) (result model.Smstask, err error) {
	err = DbEngin.Where("id = ?", property.Id).Updates(&property).Error
	return property, err
}

// 删除短信发送记录
func DeleteSmstask(property model.Smstask) (result model.Smstask, err error) {
	err = DbEngin.Where("id = ?", property.Id).Delete(&property).Error
	return property, err
}

// 逻辑删除短信发送记录
func LogicDeleteSmstask(property model.Smstask) (result model.Smstask, err error) {
	err = DbEngin.Where("id = ?", property.Id).Model(&property).Updates(map[string]interface{}{
		"deleted":   types.HasBeenDeleted,
		"delete_at": types.DateTimeNow().String(),
	}).Error
	return property, err
}

// 搜索短信发送记录
func SearchSmstask(arg args.Smstask) (result []model.Smstask, total int64, err error) {
	objs := make([]model.Smstask, 0)

	err = DbEngin.Model(new(model.Smstask)).Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Find(&objs).Error
	total = int64(arg.Total)
	if arg.Total == -1 {
		DbEngin.Model(new(model.Smstask)).Scopes(arg.Condtions()).Count(&total)
	}
	return objs, total, err
}

// 查询一条短信发送记录
func FindSmstask(id uint) (result model.Smstask, err error) {
	result = model.Smstask{}
	err = DbEngin.Model(new(model.Smstask)).Where("id = ?", id).First(&result).Error
	return result, err
}
