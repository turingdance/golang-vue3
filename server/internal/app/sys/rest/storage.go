package rest

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/turingdance/infra/slicekit"
	"github.com/turingdance/infra/wraper"
	"turingdance.com/reliable/internal/app/sys/args"
	"turingdance.com/reliable/internal/pkg/storage"
	"turingdance.com/reliable/internal/server/middleware"
)

var storageManager *storage.LocalStorage
var ossClient *oss.Client

// 这一层里面实现router
type Storage struct{}

// 处理/upload 逻辑
// /attach/render?key=filekey
func (ctrl *Storage) Render(w http.ResponseWriter, req *http.Request) {
	arg := args.Attach{}
	if err := wraper.Bind(req, &arg); err != nil {
		wraper.Error(err).Encode(w)
		return
	}
	vars, _ := url.ParseQuery(req.RequestURI)
	reskey := vars.Get("key")
	fmt.Printf("%+v", vars)
	dstfile := strings.ReplaceAll(reskey, appCfg.Local.Alias, appCfg.Local.Datadir)
	http.ServeFile(w, req, dstfile)
}

var signer *middleware.Signer

// /attach/shareUrl?filekey=filekey
func (ctrl *Storage) ShareUrl(w http.ResponseWriter, req *http.Request) {

	reskey := req.URL.Query().Get("filekey")
	conf, ok := slicekit.Find[storage.StorageConf](appCfg.Storage, func(item storage.StorageConf, index int, slice []storage.StorageConf) bool {
		return strings.HasPrefix(reskey, item.Bucket+"/")
	})
	if !ok {
		wraper.Error("暂时不支持").Encode(w)
		return
	}
	// 判断是否是oss
	if signer == nil {
		signer = middleware.NewSigner(appCfg.Enpryt)
	}
	// 如果是本地路径
	url := ""
	if conf.Driver == storage.DriverLocal {
		if conf.Ssl {
			url += "https://"
		} else {
			url += "http://"
		}
		url += conf.Host + "/"
		url += reskey
		if conf.AuthRequired {
			// 签名
			sign, expireat := signer.Sign(reskey)
			url += fmt.Sprintf("?sign=%s&expireAt=%d", sign, expireat)
		}
	}
	// oss 暂时不支持
	if conf.Driver == storage.DriverOSS {
		if ossClient == nil {
			_ossClient, err := oss.New(appCfg.Oss.Endpoint, appCfg.Oss.AccessKeyId, appCfg.Oss.AccessKeySecret)
			if err != nil {
				wraper.Error(err.Error()).Encode(w)
				return
			}
			ossClient = _ossClient
		}
		bucket, err := ossClient.Bucket(conf.Bucket)
		if err != nil {
			wraper.Error(err.Error()).Encode(w)
			return
		}
		signedURL, err := bucket.SignURL(reskey, oss.HTTPGet, 3600)
		if err != nil {
			wraper.Error(err.Error()).Encode(w)
			return
		}
		url = signedURL
	}
	wraper.OkData(url).Encode(w)
}

// 处理/upload 逻辑
// Router /storage/upload [POST,PUT]
func (ctrl *Storage) Upload(w http.ResponseWriter, r *http.Request) {
	if storageManager == nil {
		storageManager = storage.NewLocalStorage(appCfg.Storage)
	}
	resp, err := storageManager.Upload(r)
	if err != nil {
		wraper.Error(err.Error()).Encode(w)
	} else {
		wraper.OkData(resp).WithMsg("上传成功").Encode(w)
	}
}

// 处理/upload 逻辑
func (ctrl *Storage) Config(w http.ResponseWriter, r *http.Request) {
	wraper.OkData(appCfg.Storage).Encode(w)
}
