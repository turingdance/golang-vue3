package logic

//岗位
import (
	"errors"

	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/app/sys/model"
)

// 创建岗位
func CreateDept(property model.Dept) (result model.Dept, err error) {
	err = DbEngin.Create(&property).Error
	return property, err
}

// 更新岗位
func UpdateDept(property model.Dept) (result model.Dept, err error) {
	err = DbEngin.Where("id = ?", property.Id).Updates(&property).Error
	return property, err
}

// 删除岗位
func DeleteDept(property model.Dept) (result model.Dept, err error) {
	err = DbEngin.Where("id = ?", property.Id).Delete(&property).Error
	return property, err
}

// 逻辑删除岗位
func LogicDeleteDept(property model.Dept) (err error) {
	var total int64 = 0
	DbEngin.Table(property.TableName()).Where("dept_id = ?", property.Id).Count(&total)
	if total > 0 {
		err = errors.New("当前部门已被使用,无法删除")
		return
	}
	DbEngin.Where("id = ?", property.Id).Delete(&property)
	return
}

type DeptNode struct {
	Value    string     `json:"value"`
	Label    string     `json:"label"`
	Dept     model.Dept `json:"dept"`
	Children []DeptNode `json:"children"`
}

func AllDept() []model.Dept {
	objs := make([]model.Dept, 0)
	DbEngin.Model(new(model.Dept)).Where("status = 1").Find(&objs).Limit(1024)
	return objs
}

func Transfer(depts []model.Dept) []DeptNode {
	nodes := make([]DeptNode, 0)
	for _, v := range depts {
		nodes = append(nodes, DeptNode{
			Dept:  v,
			Value: v.Id,
			Label: v.Title,
		})
	}
	return nodes
}

// 搜索岗位
func DeptTree(nodes []DeptNode, pid string) (result []DeptNode) {
	res := make([]DeptNode, 0)
	for _, node := range nodes {
		if node.Dept.Pid == pid {
			node.Children = DeptTree(nodes, node.Dept.Id)
			res = append(res, node)
		}
	}
	return res
}

// 搜索岗位
func SearchDept(arg args.Dept) (result []model.Dept, total int64, err error) {
	objs := make([]model.Dept, 0)

	err = DbEngin.Model(new(model.Dept)).Scopes(arg.Paginate(), arg.Condtions()).Order(arg.Sort()).Find(&objs).Error
	total = int64(arg.Total)
	if arg.Total < 1 {
		DbEngin.Model(new(model.Dept)).Scopes(arg.Condtions()).Count(&total)
	}
	return objs, total, err
}

// 查询一条岗位
func FindDept(deptId any) (result model.Dept, err error) {
	result = model.Dept{}
	err = DbEngin.Model(new(model.Dept)).Where("id = ?", deptId).First(&result).Error
	return result, err
}
