package logic

//机构信息
import (
	"errors"

	"gorm.io/gorm"
	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/app/sys/model"
	"turingdance.com/reliable/internal/types"
)

// 创建机构信息
func CreateOrg(property model.Org) (result model.Org, err error) {
	if property.UserId == "" {
		err = errors.New("请登陆后再试")
		return
	}
	if property.Title == "" {
		err = errors.New("请输入企业名称")
		return
	}
	org, err1 := FindOrgByUserId(property.UserId)
	if err1 != nil {
		err = err1
		return
	}
	if org.OrgId > 0 {
		err = errors.New("每个账户只能创建一个商家")
		return
	}
	err = DbEngin.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&property).Error
	})
	return property, err
}

// 更新机构信息
func UpdateOrg(property model.Org) (result model.Org, err error) {
	err = DbEngin.Where("org_id = ?", property.OrgId).Updates(&property).Error
	return property, err
}

// 删除机构信息
func DeleteOrg(property model.Org) (result model.Org, err error) {
	err = DbEngin.Where("org_id = ?", property.OrgId).Delete(&property).Error
	return property, err
}

// 逻辑删除机构信息
func LogicDeleteOrg(property model.Org) (result model.Org, err error) {
	err = DbEngin.Where("org_id = ?", property.OrgId).Model(&property).Updates(map[string]interface{}{
		"deleted":   types.HasBeenDeleted,
		"delete_at": types.DateTimeNow().String(),
	}).Error
	return property, err
}

// 搜索机构信息
func SearchOrg(arg args.Org) (result []model.Org, total int64, err error) {
	objs := make([]model.Org, 0)

	err = DbEngin.Model(new(model.Org)).Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Find(&objs).Error
	DbEngin.Model(new(model.Org)).Scopes(arg.Condtions()).Count(&total)
	return objs, total, err
}

// 查询一条机构信息
func FindOrg(orgId uint) (result model.Org, err error) {
	result = model.Org{}
	err = DbEngin.Model(new(model.Org)).Where("org_id = ?", orgId).First(&result).Error
	return result, err
}

// 查询一条机构信息
func FindOrgByUserId(userId string) (result model.Org, err error) {
	result = model.Org{}
	err = DbEngin.Model(new(model.Org)).Where("user_id = ?", userId).First(&result).Error
	return result, err
}
