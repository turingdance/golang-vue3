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
	syslogic "turingdance.com/reliable/internal/app/sys/logic"
)

// 控制器
type Record struct{}

// 创建上课记录
// @Router /record [POST]
func (ctrl *Record) Create(w http.ResponseWriter, req *http.Request) {
	instance := &model.Record{}
	if err := wraper.Bind(req, instance); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	userId, err := syslogic.ParseUserId(req)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if err := logic.ReadIt(userId, instance.LessonId, 1); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkMsg("上课记录创建成功").Encode(w)
	}
}

// 修改上课记录
// @Router /record [PUT]
func (ctrl *Record) Update(w http.ResponseWriter, req *http.Request) {
	instance := &model.Record{}
	if err := wraper.Bind(req, instance); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	if instance, err := logic.Update(instance, "id = ?", instance.Id); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(instance).WithMsg("上课记录修改成功").Encode(w)
	}
}

// 根据条件查询上课记录
// @Router /record/search [POST,GET]
func (ctrl *Record) Search(w http.ResponseWriter, req *http.Request) {
	condwraper := cond.NewCondWrapper()
	if err := wraper.Bind(req, condwraper); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	instance := &model.RecordView{}
	if rows, total, err := logic.Search(instance, condwraper); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(rows).WithTotal(total).Encode(w)
	}
}

// 根据条件查询上课记录
// @Router /record/mine [POST,GET]
func (ctrl *Record) Mine(w http.ResponseWriter, req *http.Request) {
	condwraper := cond.NewCondWrapper()
	if err := wraper.Bind(req, condwraper); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	userId, err := syslogic.ParseUserId(req)
	if err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	condwraper.AddCond(cond.Cond{
		Field: "user_id",
		Op:    cond.OPEQ,
		Value: userId,
	})
	instance := &model.RecordView{}
	if rows, total, err := logic.Search(instance, condwraper); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(rows).WithTotal(total).Encode(w)
	}
}

// 根据主键删除上课记录
// @Router /record/{pkId} [DELETE]
func (ctrl *Record) Delete(w http.ResponseWriter, req *http.Request) {

	pkId, _ := wraper.MuxIntVar(req, "pkId", int32(0))

	instance := &model.Record{}
	if total, err := logic.Delete(instance, "id = ?", pkId); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(total).WithMsg("上课记录删除成功").Encode(w)
	}
}

// 根据主键批量删除上课记录
// @Router /record [DELETE]
func (ctrl *Record) DeleteBatch(w http.ResponseWriter, req *http.Request) {
	cond := &vo.RecordKeyBatch{}
	if err := wraper.Bind(req, cond); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	instance := &model.Record{}
	if total, err := logic.Delete(instance, "id = ?", cond.Ids); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(total).WithMsg("上课记录删除成功").Encode(w)
	}

}

// 导出上课记录
// @Router /record/export [POST,GET]
func (ctrl *Record) Export(w http.ResponseWriter, req *http.Request) {
	condwraper := cond.NewExport()

	if err := wraper.Bind(req, condwraper); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	instance := &model.Record{}
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
	dataMap := slicekit.ObjListToMapList[model.Record](rows)
	// 创建一个工作表
	xlsxctrl := xlskit.NewXlsCtrl(model.TableTitleRecord)
	if buf, err := xlsxctrl.Render(metaArr, dataMap); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.Blob(wraper.BlobDef{
			File:        buf.Bytes(),
			Name:        model.TableTitleRecord + ".xls",
			ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		}).Encode(w)
	}
}

// 根据主键获取1条上课记录
// @Router /record/{pkId} [GET]
func (ctrl *Record) GetOne(w http.ResponseWriter, req *http.Request) {

	pkId, _ := wraper.MuxIntVar(req, "pkId", int64(0))

	instance := &model.Record{
		Id: pkId,
	}
	if instance, err := logic.TakeByPrimaryKey(instance); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(instance).Encode(w)
	}
}
