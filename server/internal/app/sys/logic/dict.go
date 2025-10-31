package logic

//字典
import (
	"turingdance.com/turing/internal/app/sys/args"
	"turingdance.com/turing/internal/app/sys/model"
	"turingdance.com/turing/internal/types"
)

// 创建字典
func CreateDict(property model.Dict) (result model.Dict, err error) {
	err = DbEngin.Create(&property).Error
	return property, err
}

// 更新字典
func UpdateDict(property model.Dict) (result model.Dict, err error) {
	err = DbEngin.Where("id = ?", property.Id).Updates(&property).Error
	return property, err
}

// 删除字典
func DeleteDict(property model.Dict) (result model.Dict, err error) {
	err = DbEngin.Where("id = ?", property.Id).Delete(&property).Error
	return property, err
}

// 逻辑删除字典
func LogicDeleteDict(property model.Dict) (result model.Dict, err error) {
	err = DbEngin.Model(&property).Updates(map[string]interface{}{
		"deleted":   types.HasBeenDeleted,
		"delete_at": types.DateTimeNow().String(),
	}).Error
	return property, err
}

// 搜索字典
func SearchDict(arg args.Dict) (result []model.Dict, total int64, err error) {
	objs := make([]model.Dict, 0)

	err = DbEngin.Model(new(model.Dict)).Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Count(&total).Find(&objs).Error
	total = int64(arg.Total)
	DbEngin.Model(new(model.Dict)).Scopes(arg.Condtions()).Count(&total)
	return objs, total, err
}

// 搜索字典
func SearchAllDict(arg args.Dict) (result []model.Dict, total int64, err error) {
	objs := make([]model.Dict, 0)
	err = DbEngin.Model(new(model.Dict)).Scopes(arg.Condtions()).Order(arg.Sort()).Count(&total).Find(&objs).Error
	return objs, total, err
}

// 查询一条字典
func FindDict(param any) (result model.Dict, err error) {
	result = model.Dict{}
	cond := make(map[string]interface{})
	if param, yes := param.(int); yes {
		cond["id"] = param
	}
	if param, yes := param.(string); yes {
		cond["name"] = param
	}

	err = DbEngin.Model(new(model.Dict)).First(&result, cond).Error
	return result, err
}
