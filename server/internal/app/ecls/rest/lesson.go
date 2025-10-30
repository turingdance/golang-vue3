// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git
// @author winlion@turingdance.com
package rest

import (
	"net/http"

	"github.com/turingdance/infra/cond"
	"github.com/turingdance/infra/slicekit"
	"github.com/turingdance/infra/wraper"
	"github.com/turingdance/infra/xlskit"
	"turingdance.com/reliable/internal/app/ecls/logic"
	"turingdance.com/reliable/internal/app/ecls/model"
	"turingdance.com/reliable/internal/app/ecls/vo"
)

// 控制器
type Lesson struct{}

// 创建课件资源
// @Router /lesson [POST]
func (ctrl *Lesson) Create(w http.ResponseWriter, req *http.Request) {
	instance := &model.Lesson{}
	if err := wraper.Bind(req, instance); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if instance, err := logic.Create(instance); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(instance).WithMsg("课件资源创建成功").Encode(w)
	}
}

// 创建课件资源
// @Router /lesson/batchAdd [POST]
func (ctrl *Lesson) BatchAdd(w http.ResponseWriter, req *http.Request) {
	instance := &vo.AddLessonBatch{}
	if err := wraper.Bind(req, instance); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if instance, err := logic.CreateBatch(instance.Data); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(instance).WithMsg("课件资源创建成功").Encode(w)
	}
}

// 修改课件资源
// @Router /lesson [PUT]
func (ctrl *Lesson) Update(w http.ResponseWriter, req *http.Request) {
	instance := &model.Lesson{}
	if err := wraper.Bind(req, instance); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if instance, err := logic.Update(instance, "id = ?", instance.Id); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(instance).WithMsg("课件资源修改成功").Encode(w)
	}
}

// 根据条件查询课件资源
// @Router /lesson/search [POST,GET]
func (ctrl *Lesson) Search(w http.ResponseWriter, req *http.Request) {
	condwraper := cond.NewCondWrapper()
	if err := wraper.Bind(req, condwraper); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	instance := &model.Lesson{}
	if rows, total, err := logic.Search(instance, condwraper); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(rows).WithTotal(total).Encode(w)
	}
}

// 根据主键删除课件资源
// @Router /lesson/{pkId} [DELETE]
func (ctrl *Lesson) Delete(w http.ResponseWriter, req *http.Request) {

	pkId, _ := wraper.MuxStringVar(req, "pkId", "")

	instance := &model.Lesson{}
	if total, err := logic.Delete(instance, "id = ?", pkId); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(total).WithMsg("课件资源删除成功").Encode(w)
	}
}

// 根据主键批量删除课件资源
// @Router /lesson [DELETE]
func (ctrl *Lesson) DeleteBatch(w http.ResponseWriter, req *http.Request) {
	cond := &vo.LessonKeyBatch{}
	if err := wraper.Bind(req, cond); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	instance := &model.Lesson{}
	if total, err := logic.Delete(instance, "id in ?", cond.Ids); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(total).WithMsg("课件资源删除成功").Encode(w)
	}

}

// 导出课件资源
// @Router /lesson/export [POST,GET]
func (ctrl *Lesson) Export(w http.ResponseWriter, req *http.Request) {
	condwraper := cond.NewExport()

	if err := wraper.Bind(req, condwraper); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	instance := &model.Lesson{}
	rows, _, err := logic.Search(instance, condwraper.Cond)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	//

	metaArr := make([]xlskit.Meta, 0)
	for _, v := range condwraper.Meta {
		metaArr = append(metaArr, xlskit.Meta{
			Field: v.Prop,
			Title: v.Label,
		})
	}
	dataMap := slicekit.ObjListToMapList[model.Lesson](rows)
	// 创建一个工作表
	xlsxctrl := xlskit.NewXlsCtrl(model.TableTitleLesson)
	if buf, err := xlsxctrl.Render(metaArr, dataMap); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.Blob(wraper.BlobDef{
			File:        buf.Bytes(),
			Name:        model.TableTitleLesson + ".xls",
			ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		}).Encode(w)
	}
}

// 根据主键获取1条课件资源
// @Router /lesson/{pkId} [GET]
func (ctrl *Lesson) GetOne(w http.ResponseWriter, req *http.Request) {

	pkId, _ := wraper.MuxStringVar(req, "pkId", "")

	instance := &model.Lesson{
		Id: pkId,
	}
	if instance, err := logic.TakeByPrimaryKey(instance); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(instance).Encode(w)
	}
}
