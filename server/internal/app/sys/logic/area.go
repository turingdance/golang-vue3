package logic

//区域
import (
	"turingdance.com/turing/internal/app/sys/args"
	"turingdance.com/turing/internal/app/sys/model"
	"turingdance.com/turing/internal/types"
)

// 创建区域
func CreateArea(property model.Area) (result model.Area, err error) {
	err = DbEngin.Create(&property).Error
	return property, err
}

// 更新区域
func UpdateArea(property model.Area) (result model.Area, err error) {
	err = DbEngin.Where("area_id = ?", property.AreaId).Updates(&property).Error
	return property, err
}

// 删除区域
func DeleteArea(property model.Area) (result model.Area, err error) {
	err = DbEngin.Where("area_id = ?", property.AreaId).Delete(&property).Error
	return property, err
}

// 逻辑删除区域
func LogicDeleteArea(property model.Area) (result model.Area, err error) {
	err = DbEngin.Where("area_id = ?", property.AreaId).Model(&property).Updates(map[string]interface{}{
		"deleted":   types.HasBeenDeleted,
		"delete_at": types.DateTimeNow().String(),
	}).Error
	return property, err
}

// 搜索区域
func SearchArea(arg args.Area) (result []model.Area, total int64, err error) {
	objs := make([]model.Area, 0)

	err = DbEngin.Model(new(model.Area)).Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Find(&objs).Error
	total = int64(arg.Total)
	if arg.Total < 1 {
		DbEngin.Model(new(model.Area)).Scopes(arg.Condtions()).Count(&total)
	}
	return objs, total, err
}

// 查询一条区域
func FindArea(areaId uint) (result model.Area, err error) {
	result = model.Area{}
	err = DbEngin.Model(new(model.Area)).Where("areaId = ?", areaId).First(&result).Error
	return result, err
}
