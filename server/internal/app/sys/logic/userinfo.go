package logic

//
import (
	"fmt"

	"turingdance.com/turing/internal/app/sys/args"
	"turingdance.com/turing/internal/app/sys/model"
)

// 创建
func CreateUserinfo(property model.Userinfo) (result model.Userinfo, err error) {
	var tmp model.Userinfo
	if property.Username == "" {
		return property, fmt.Errorf("缺少用户名")
	}
	err = DbEngin.Where("username = ?", property.Username).Find(&tmp).Error
	if tmp.UserId == "" {
		if property.Username == "" {
			property.Username = property.Mobile
		}
		err = DbEngin.Create(&property).Error
	} else {
		return property, fmt.Errorf("该手机号已经注册")
	}
	return property, nil
}

// 更新
func UpdateUserinfo(property model.Userinfo) (result model.Userinfo, err error) {
	err = DbEngin.Where("user_id = ?", property.UserId).Updates(&property).Error
	return property, err
}

// 删除
func DeleteUserinfo(property model.Userinfo) (result model.Userinfo, err error) {
	err = DbEngin.Where("user_id = ?", property.UserId).Delete(&property).Error
	return property, nil
}

// 逻辑删除
func LogicDeleteUserinfo(property model.Userinfo) (result model.Userinfo, err error) {
	err = DbEngin.Where("user_id = ?", property.UserId).Model(&property).Updates(map[string]interface{}{
		"deleted": true,
	}).Error
	return property, nil
}

// 搜索
func SearchUserinfo(arg args.Userinfo) (result []model.Userinfo, total int64, err error) {
	objs := make([]model.Userinfo, 0)
	err = DbEngin.Model(new(model.Userinfo)).Preload("Role").Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Count(&total).Find(&objs).Error
	return objs, total, nil
}

// 搜索
func Updatedeptandname(userId, deptId, name string) (err error) {
	userInfo := &model.Userinfo{
		UserId:   userId,
		DeptId:   deptId,
		Nickname: name,
	}
	err = DbEngin.Model(new(model.Userinfo)).Where("user_id = ?", userId).Updates(userInfo).Error
	return err
}

// 搜索
func CountUserinfo(arg args.Userinfo) (total int64, err error) {

	err = DbEngin.Model(new(model.Userinfo)).Scopes(arg.Paginate(), arg.Condtions()).Count(&total).Error

	return total, err
}

// 搜索
func SetDefaultOrgId(userId string, refId uint) (err error) {
	err = DbEngin.Model(new(model.Userinfo)).Where("user_id = ?", userId).UpdateColumn("default_ref_id", refId).Error
	return err
}

// 查询一条
func FindUserinfo(userId string) (result model.Userinfo, err error) {
	result = model.Userinfo{}
	err = DbEngin.Model(new(model.Userinfo)).Where("user_id = ?", userId).Preload("Role").First(&result).Error
	return result, err
}
func UpdateUserName(userId string, username string) (err error) {

	err = DbEngin.Model(new(model.Userinfo)).Where("user_id = ?", userId).UpdateColumn("username", username).Error
	return err
}

func UpdateMobile(userId string, mobile string) (err error) {
	err = DbEngin.Model(new(model.Userinfo)).Where("user_id = ?", userId).UpdateColumn("mobile", mobile).Error
	return err
}

// 查询一条机构信息
func FindUserinfoByMobile(mobile string) (result model.Userinfo, err error) {
	result = model.Userinfo{}
	err = DbEngin.Where("mobile = ?", mobile).First(&result).Error
	return result, err
}

// 查询一条
func FindSimpleInfo(userId string) (result model.Userinfo, err error) {
	result = model.Userinfo{}
	err = DbEngin.Model(new(model.Userinfo)).Select("nickname,pic,user_id").Where("user_id = ?", userId).First(&result).Error
	return result, err
}
