
// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git
// @author winlion
package rest
import (
	"net/http"
	"turingdance.com/reliable/internal/app/testapp/logic"
	"turingdance.com/reliable/internal/app/testapp/model"
	"turingdance.com/reliable/internal/app/testapp/vo"
	"github.com/turingdance/infra/cond"
	"github.com/turingdance/infra/slicekit"
	"github.com/turingdance/infra/wraper"
	"github.com/turingdance/infra/xlskit"
)
// 控制器
type SysArea struct{}
// 创建区域信息
// @Router /sysArea [POST]
func (ctrl *SysArea) Create(w http.ResponseWriter, req *http.Request) {
	instance := &model.SysArea{}
	if err := wraper.Bind(req, instance);err!=nil{
		wraper.Error(err).Encode(w)
		return 
	}
	if instance, err := logic.Create(instance);err!=nil{
		wraper.Error(err).Encode(w)
	}else{
		wraper.OkData(instance).WithMsg("区域信息创建成功").Encode(w)
	}
}
// 修改区域信息
// @Router /sysArea [PUT]
func (ctrl *SysArea) Update(w http.ResponseWriter, req *http.Request) {
	instance := &model.SysArea{}
	if err := wraper.Bind(req, instance);err!=nil{
		wraper.Error(err).Encode(w)
		return
	}
	if instance, err := logic.Update(instance, "area_id = ?", instance.Area_id);err!=nil{
		wraper.Error(err).Encode(w)
	}else{
		wraper.OkData(instance).WithMsg("区域信息修改成功").Encode(w)
	}
}

// 根据条件查询区域信息
// @Router /sysArea/search [POST,GET]
func (ctrl *SysArea) Search(w http.ResponseWriter, req *http.Request){
	condwraper := cond.NewCondWrapper()
	if err := wraper.Bind(req, condwraper); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	instance := &model.SysArea{}
	if rows, total, err := logic.Search(instance, condwraper);err!=nil{
		wraper.Error(err).Encode(w)
	}else{
		wraper.OkData(rows).WithTotal(total).Encode(w)
	}
}

// 根据主键删除区域信息
// @Router /sysArea/{pkId} [DELETE]
func (ctrl *SysArea) Delete(w http.ResponseWriter, req *http.Request) {
	
	pkId, _ := wraper.MuxIntVar(req, "pkId", int32(0))
	
	instance := &model.SysArea{}
	if total, err := logic.Delete(instance, "area_id = ?", pkId);err!=nil{
		wraper.Error(err).Encode(w)
	}else{
		wraper.OkData(total).WithMsg("区域信息删除成功").Encode(w)
	}
}

// 根据主键批量删除区域信息
// @Router /sysArea [DELETE]
func (ctrl *SysArea) DeleteBatch(w http.ResponseWriter, req *http.Request){
	cond := &vo.SysAreaKeyBatch{}
	if err := wraper.Bind(req, cond) ;err!=nil {
		wraper.Error(err).Encode(w)
		return 
	}
	if len(cond.Area_ids)==0{
		wraper.OkData(0).Encode(w)
		return 
	}
	instance := &model.SysArea{}
	if total, err := logic.Delete(instance, "area_id in ?", cond.Area_ids);err!=nil{
		wraper.Error(err).Encode(w)
	}else{
		wraper.OkData(total).WithMsg("区域信息删除成功").Encode(w)
	}
	
	
}
// 导出区域信息
// @Router /sysArea/export [POST,GET]
func (ctrl *SysArea) Export(w http.ResponseWriter, req *http.Request){
	condwraper := cond.NewExport()
	
	if err:=wraper.Bind(req,condwraper) ;err != nil {
		wraper.Error(err).Encode(w)
		return 
	}
	instance := &model.SysArea{}
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
	dataMap := slicekit.ObjListToMapList[model.SysArea](rows)
	// 创建一个工作表
	xlsxctrl := xlskit.NewXlsCtrl(model.TableTitleSysArea)
	if buf, err := xlsxctrl.Render(metaArr, dataMap); err != nil {
		wraper.Error(err).Encode(w)
	}else{
		wraper.Blob(wraper.BlobDef{
			File:        buf.Bytes(),
			Name:        model.TableTitleSysArea+".xls",
			ContentType: "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
		}).Encode(w)
	}
}


// 根据主键获取1条区域信息
// @Router /sysArea/{pkId} [GET]
func (ctrl *SysArea) GetOne(w http.ResponseWriter, req *http.Request) {
	
	pkId, _ := wraper.MuxIntVar(req, "pkId", int64(0))
	
	instance := &model.SysArea{
		AreaId:pkId,
	}
	if instance, err := logic.TakeByPrimaryKey(instance);err!=nil{
		wraper.Error(err).Encode(w)
	}else{
		wraper.OkData(instance).Encode(w)
	}
}
