package storage

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/turingdance/infra/alikit/osskit"
	"github.com/turingdance/infra/slicekit"
	"turingdance.com/turing/internal/server/middleware"
)

var ossClient *oss.Client = nil

func SignUrl(storageConfs []StorageConf, enpryt middleware.Enpryt, ossConf osskit.OssConf, reskey string) (url string, err error) {

	conf, ok := slicekit.Find(storageConfs, func(item StorageConf, index int, slice []StorageConf) bool {
		return strings.HasPrefix(reskey, item.Bucket+"/")
	})
	if !ok {
		err = errors.New("暂时不支持")
		return
	}
	// 判断是否是oss

	signer := middleware.NewSigner(enpryt)

	// 如果是本地路径
	url = ""
	if conf.Driver == DriverLocal {
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
	if conf.Driver == DriverOSS {
		if ossClient == nil {
			_ossClient, _err := oss.New(ossConf.Endpoint, ossConf.AccessKeyId, ossConf.AccessKeySecret)
			if _err != nil {
				err = _err
				return
			}
			ossClient = _ossClient
		}
		bucket, _err := ossClient.Bucket(conf.Bucket)
		if err != nil {
			err = _err
			return
		}
		url, err = bucket.SignURL(reskey, oss.HTTPGet, 3600)
	}
	return url, err
}
