package logic

//
import (
	"errors"

	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/app/sys/model"
	"turingdance.com/reliable/internal/types"
)

// 创建
func CreateConfig(property model.Config) (result model.Config, err error) {
	err = DbEngin.Create(&property).Error
	return property, err
}

// 创建
func SaveConfig(property model.Config) (result model.Config, err error) {
	var tmp model.Config
	DbEngin.Model(&tmp).Where("org_id = ?", property.OrgId).Where("name = ?", property.Name).First(&tmp)
	if tmp.Id > 0 {
		tmp.Value = property.Value
	}
	tmp.Name = property.Name
	tmp.OrgId = property.OrgId
	tmp.Label = property.Label

	err = DbEngin.Model(&tmp).Save(&tmp).Error
	return tmp, err
}

// 更新
func UpdateConfig(property model.Config) (result model.Config, err error) {
	err = DbEngin.Where("id = ?", property.Id).Updates(&property).Error
	return property, err
}

// 删除
func DeleteConfig(property model.Config) (result model.Config, err error) {
	err = DbEngin.Where("id = ?", property.Id).Delete(&property).Error
	return property, err
}

// 逻辑删除
func LogicDeleteConfig(property model.Config) (result model.Config, err error) {
	err = DbEngin.Where("id = ?", property.Id).Model(&property).Updates(map[string]interface{}{
		"deleted":   types.HasBeenDeleted,
		"delete_at": types.DateTimeNow().String(),
	}).Error
	return property, err
}

// 搜索
func SearchConfig(arg args.Config) (result []model.Config, total int64, err error) {
	objs := make([]model.Config, 0)

	err = DbEngin.Model(new(model.Config)).Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Find(&objs).Error
	total = int64(arg.Total)
	if arg.Total == -1 {
		DbEngin.Model(new(model.Config)).Scopes(arg.Condtions()).Count(&total)
	}
	return objs, total, err
}

// 查询一条
func FindConfig(id uint) (result model.Config, err error) {
	result = model.Config{}
	err = DbEngin.Model(new(model.Config)).Where("id = ?", id).First(&result).Error
	return result, err
}

// 查询一条
func FindConfigValue(orgId uint, name string) (result model.Config, err error) {
	if orgId == 0 {
		err = errors.New("缺少门店信息")
		return
	}
	if name == "" {
		err = errors.New("缺少参数信息")
		return
	}
	result = model.Config{}
	err = DbEngin.Model(new(model.Config)).Where("org_id = ?", orgId).Where("name = ?", name).First(&result).Error
	return result, err
}

// 查询一条
func FindConfigValueByPrefix(prefix string) (result []model.Config, err error) {
	result = []model.Config{}
	err = DbEngin.Model(new(model.Config)).Where("name like ?", prefix+".%").Limit(1000).Find(&result).Error
	return result, err
}
