package logic

//
import (
	"errors"

	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/app/sys/model"
)

// 创建
func CreateRights(property model.Rights) (result model.Rights, err error) {
	err = DbEngin.Create(&property).Error
	return property, nil
}

// 更新
func UpdateRights(property model.Rights) (result model.Rights, err error) {
	err = DbEngin.Where("id = ?", property.Id).Updates(&property).Error
	return property, err
}

// 删除
func DeleteRights(property model.Rights) (result model.Rights, err error) {
	err = DbEngin.Where("id = ?", property.Id).Delete(&property).Error
	return property, err
}

// 逻辑删除
func LogicDeleteRights(property model.Rights) (result model.Rights, err error) {
	var total int64 = 0
	err = DbEngin.Where("pid = ?", property.Id).Model(&property).Count(&total).Error
	if err != nil {
		return
	}
	if total > 0 {
		err = errors.New("当前节点存在子节点无法被删除")
		return
	}
	err = DbEngin.Model(&property).Delete(&property, "id = ?", property.Id).Error
	return property, err
}

// 搜索
func SearchRights(arg args.Rights) (result []model.Rights, total int64, err error) {
	objs := make([]model.Rights, 0)

	err = DbEngin.Model(new(model.Rights)).Scopes(arg.Paginate(), arg.Condtions()).Order("sort_index asc,id asc").Find(&objs).Error
	total = int64(arg.Total)
	if arg.Total == -1 {
		DbEngin.Model(new(model.Rights)).Scopes(arg.Condtions()).Count(&total)
	}
	return objs, total, err
}

// 查询一条
func FindRights(RightsId uint) (result model.Rights, err error) {
	result = model.Rights{}
	err = DbEngin.Model(new(model.Rights)).Where("id = ?", RightsId).First(&result).Error
	return result, err
}

// 查询一条
func FindRightsByIds(rightsIds []uint) (result []model.Rights, err error) {
	results := []model.Rights{}
	err = DbEngin.Model(new(model.Rights)).Find(&results, "id IN ?", rightsIds).Error
	return results, err
}

type RightsNode struct {
	Value    uint         `json:"value"`
	Label    string       `json:"label"`
	Rights   model.Rights `json:"rights"`
	Children []RightsNode `json:"children"`
}

func AllRights() []model.Rights {
	objs := make([]model.Rights, 0)
	DbEngin.Model(new(model.Rights)).Find(&objs).Limit(1024)
	return objs
}

func BuildRightsNodeFrom(input []model.Rights) []RightsNode {
	res := make([]RightsNode, 0)
	for _, v := range input {
		res = append(res, RightsNode{
			Value:  v.Id,
			Label:  v.Title,
			Rights: v,
		})
	}
	return res
}
func RightsMap() map[uint]model.Rights {
	rights := AllRights()
	result := map[uint]model.Rights{}
	for _, v := range rights {
		result[v.Id] = v
	}
	return result
}

// 搜索岗位
func RightsTree(nodes []RightsNode, rootId uint) (result []RightsNode) {
	res := make([]RightsNode, 0)
	for _, node := range nodes {
		if node.Rights.Pid == rootId {
			node.Children = RightsTree(nodes, node.Value)
			res = append(res, node)
		}
	}
	return res
}
