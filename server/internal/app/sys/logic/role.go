package logic

//
import (
	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/app/sys/model"
)

// 创建
func CreateRole(property model.Role) (result model.Role, err error) {
	err = DbEngin.Create(&property).Error
	return property, err
}

// 更新
func UpdateRole(property model.Role) (result model.Role, err error) {
	err = DbEngin.Where("id = ?", property.Id).Updates(&property).Error
	return property, err
}

// 删除
func DeleteRole(property model.Role) (result model.Role, err error) {
	err = DbEngin.Where("id = ?", property.Id).Delete(&property).Error
	return property, err
}

// 逻辑删除
func UpdateRoleState(property model.Role) (result model.Role, err error) {
	err = DbEngin.Where("id = ?", property.Id).Model(&property).Updates(map[string]interface{}{
		"state": property.State,
	}).Error
	return property, err
}

// 更新
func Grant(roleId uint, rightsIds []uint) (err error) {
	err = DbEngin.Where("id = ?", roleId).Updates(&model.Role{
		RightIds: rightsIds,
	}).Error
	return err
}

// 搜索
func SearchRole(arg args.Role) (result []model.Role, total int64, err error) {
	objs := make([]model.Role, 0)

	err = DbEngin.Model(new(model.Role)).Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Find(&objs).Error
	if err != nil {
		return
	}
	total = int64(arg.Total)
	if arg.Total == -1 {
		DbEngin.Model(new(model.Role)).Scopes(arg.Condtions()).Count(&total)
	}
	return objs, total, nil
}

// 查询一条
func FindRole(roleId uint) (result model.Role, err error) {
	result = model.Role{}
	err = DbEngin.Model(new(model.Role)).Where("id = ?", roleId).First(&result).Error
	return result, err
}
