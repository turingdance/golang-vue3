package logic

//机构信息
import (
	"errors"

	"gorm.io/gorm"
	"turingdance.com/turing/internal/app/sys/args"
	"turingdance.com/turing/internal/app/sys/model"
	"turingdance.com/turing/internal/types"
)

// 创建机构信息
func CreateTenant(property model.Tenant) (result model.Tenant, err error) {

	if property.Name == "" {
		err = errors.New("请输入企业名称")
		return
	}
	err = DbEngin.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&property).Error
	})
	return property, err
}

// 更新机构信息
func UpdateTenant(property model.Tenant) (result model.Tenant, err error) {
	err = DbEngin.Where("id = ?", property.Id).Updates(&property).Error
	return property, err
}

// 删除机构信息
func DeleteTenant(property model.Tenant) (result model.Tenant, err error) {
	err = DbEngin.Where("id = ?", property.Id).Delete(&property).Error
	return property, err
}

// 逻辑删除机构信息
func LogicDeleteTenant(property model.Tenant) (result model.Tenant, err error) {
	err = DbEngin.Where("id = ?", property.Id).Model(&property).Updates(map[string]interface{}{
		"deleted":   types.HasBeenDeleted,
		"delete_at": types.DateTimeNow().String(),
	}).Error
	return property, err
}

// 搜索机构信息
func SearchTenant(arg args.Tenant) (result []model.Tenant, total int64, err error) {
	objs := make([]model.Tenant, 0)

	err = DbEngin.Model(new(model.Tenant)).Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Find(&objs).Error
	DbEngin.Model(new(model.Tenant)).Scopes(arg.Condtions()).Count(&total)
	return objs, total, err
}

// 查询一条机构信息
func FindTenant(tenantId string) (result model.Tenant, err error) {
	result = model.Tenant{}
	err = DbEngin.Model(new(model.Tenant)).Where("id = ?", tenantId).First(&result).Error
	return result, err
}

// 查询一条机构信息
func FindTenantByUserId(userId string) (result model.Tenant, err error) {
	result = model.Tenant{}
	err = DbEngin.Model(new(model.Tenant)).Where("user_id = ?", userId).First(&result).Error
	return result, err
}
