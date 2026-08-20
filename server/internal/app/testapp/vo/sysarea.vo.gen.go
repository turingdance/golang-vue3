
// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git
// @author winlion

package vo
import 
(
	
	"github.com/turingdance/infra/types"
	
	
)
//区域信息参数请求接口
type SysArea struct{
	
        AreaId int64 `json:"areaId"  form:"areaId"  validate:"required"`
    
        Pid int64 `json:"pid"  form:"pid"  validate:"required"`
    
        Name string `json:"name"  form:"name"  validate:"required"`
    
        Type string `json:"type"  form:"type"  validate:"required"`
    
        Rank int64 `json:"rank"  form:"rank"  validate:"required"`
    
        CreateAt types.DateTime `json:"createAt" time_format:"2006-01-02 15:04:05" time_utc:"1"  form:"createAt"  validate:"required"`
    
        DeleteAt types.DateTime `json:"deleteAt" time_format:"2006-01-02 15:04:05" time_utc:"1"  form:"deleteAt"  validate:"required"`
    
        Deleted int32 `json:"deleted"  form:"deleted"  validate:"required"`
    
}
// 区域信息批量处理
type SysAreaKeyBatch struct {
	Area_ids []int64  `json:"area_ids"  form:"area_ids"  validate:"min=1" errmsg:"缺少"`
}
