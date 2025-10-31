package rest

import (
	"net/http"

	config "turingdance.com/turing/internal/conf"
	"turingdance.com/turing/internal/server/auth"

	"github.com/turingdance/infra/alikit/osskit"
	"github.com/turingdance/infra/wraper"
)

type Oss struct {
}

func init() {
	//RegisterRestCtrl(&Oss{})
}

// 初始化上传策略
func (ctrl *Oss) Policy(w http.ResponseWriter, req *http.Request) {
	if policy, err := osskit.GeneratePolicy(config.AppConf.Oss); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		// 获得token
		token, _ := auth.ParseToken(req)
		if userId, ok := token["userId"]; ok {
			policy.Xvar["x-oss-callback-userId"] = userId
		}
		if orgId, ok := token["orgId"]; ok {
			policy.Xvar["x-oss-callback-orgId"] = orgId
		}
		wraper.OkData(policy).Encode(w)
	}
}

// 初始化上传策略
func (ctrl *Oss) Callback(w http.ResponseWriter, req *http.Request) {
	/*resource := args.RequestResourceCallBack{}
	if err := wraper.Bind(req, &resource); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	media := strings.Split(resource.MimeType,"/")
	// 管你咋的 统一上传再说
	media = append(media, "unknow")
	media = append(media, "unknow")
	// 创建用户基础信息
	logic.CreateResource(model.Resource{
		UserId: resource.UserId,
		ObjKey: resource.Object,
		OrgId: resource.OrgId,
		MimeType: resource.MimeType,
		Size: resource.Size,
		Media: media[0],
		Suffix: media[1],
	})
	oss.ResponseSuccess(w)
	*/
}
