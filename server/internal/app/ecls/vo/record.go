
// gen by codectl ,donot modify ,https://github.com/turingdance/codectl.git
// @author winlion@turingdance.com

package vo
import 
(
	
	"github.com/turingdance/infra/types"
	
	
)
//上课记录参数请求接口
type Record struct{
	
        Id int64 `json:"id"  form:"id"  validate:"required"`
    
        CreateAt types.Date `json:"createAt" time_format:"2006-01-02" time_utc:"1"  form:"createAt"  validate:"required"`
    
        UserId string `json:"userId"  form:"userId"  validate:"required"`
    
        LessonId string `json:"lessonId"  form:"lessonId"  validate:"required"`
    
        CreateDate types.Date `json:"createDate" time_format:"2006-01-02" time_utc:"1"  form:"createDate"  validate:"required"`
    
}
// 上课记录批量处理
type RecordKeyBatch struct {
	Ids []int64 `json:"ids"  form:"ids"  validate:"min=1" errmsg:"缺少ID"`
}
