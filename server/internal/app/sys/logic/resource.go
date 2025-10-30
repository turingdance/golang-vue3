package logic

//资源管理
import (
	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/app/sys/model"
	"turingdance.com/reliable/internal/types"
)

// 创建资源管理
func CreateResource(property model.Resource) (result model.Resource, err error) {
	err = DbEngin.Create(&property).Error
	return property, err
}

// 更新资源管理
func UpdateResource(property model.Resource) (result model.Resource, err error) {
	err = DbEngin.Where("res_id = ?", property.ResId).Updates(&property).Error
	return property, err
}

// 删除资源管理
func DeleteResource(property model.Resource) (result model.Resource, err error) {
	err = DbEngin.Where("res_id = ?", property.ResId).Delete(&property).Error
	return property, err
}

// 逻辑删除资源管理
func LogicDeleteResource(property model.Resource) (result model.Resource, err error) {
	err = DbEngin.Where("res_id = ?", property.ResId).Model(&property).Updates(map[string]interface{}{
		"deleted":   types.HasBeenDeleted,
		"delete_at": types.DateTimeNow().String(),
	}).Error
	return property, err
}

// 搜索资源管理
func SearchResource(arg args.Resource) (result []model.Resource, total int64, err error) {
	objs := make([]model.Resource, 0)

	err = DbEngin.Model(new(model.Resource)).Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Find(&objs).Error
	total = int64(arg.Total)
	if arg.Total < 1 {
		DbEngin.Model(new(model.Resource)).Scopes(arg.Condtions()).Count(&total)
	}
	return objs, total, err
}

// 查询一条资源管理
func FindResource(resId uint) (result model.Resource, err error) {
	result = model.Resource{}
	err = DbEngin.Model(new(model.Resource)).Where("res_id = ?", resId).First(&result).Error
	return result, err
}
