package logic

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"

	"turingdance.com/turing/internal/app/sys/model"
	"turingdance.com/turing/internal/pkg/cache"
	"turingdance.com/turing/internal/pkg/dysms"
	"turingdance.com/turing/internal/pkg/utils"
	"turingdance.com/turing/internal/types"
)

// 发送短信
var defaultDysmsService *dysms.DysmsService

func InitDysmsService(smsConf dysms.DysmsConf) *dysms.DysmsService {
	defer utils.TimeCost("InitDysmsService")()
	defaultDysmsService = dysms.NewDysmsService(smsConf, dysms.WithResultCallBack(OnsmsResultCallback), dysms.SetDefaultService())
	defaultDysmsService.Start()
	return defaultDysmsService
}
func SmsKey(phone string) string {
	return "sendsms" + phone
}

func OnsmsResultCallback(resp dysms.SmsTaskResponse) {
	// 这里用来处理结果
	if nil != resp.Error {
		log.Error(resp.Error.Error())
	} else {
		smstask := model.Smstask{
			Id:      resp.TaskId,
			Result:  resp.Resp.BaseResponse.GetHttpContentString(),
			Success: 0,
			SendAt:  types.DateTimeNow(),
		}
		if resp.Resp.IsSuccess() {
			smstask.Success = 1
		}
		UpdateSmstask(smstask)
	}
}

// 俩分钟有效
func SendSms(phone string) error {
	code := utils.RandNumber(4)
	r, err := cache.CacheEngin.Get(phone)

	if r != nil {
		return errors.New("请稍候再试")
	}
	smstask := model.Smstask{}
	smstask.PhoneNumbers = phone
	smstask.Code = code

	result, err := CreateSmstask(smstask)
	if err != nil {
		log.Error(err.Error())
	}
	cache.CacheEngin.Set(SmsKey(phone), code, time.Second*120)
	defaultDysmsService.Publish(dysms.SmsTask{TaskId: result.Id, PhoneNumbers: []string{phone}, Params: map[string]string{"code": code}})
	return nil
}

// 俩分钟有效
func VerifySms(phone, code string) error {
	if code == "" {
		return errors.New("缺少验证码")
	}
	code0, err := cache.CacheEngin.Get(SmsKey(phone))
	cache.CacheEngin.Set(SmsKey(phone), "", 0)
	if err != nil {
		return errors.New("验证码已过期")
	} else {
		if code0 != code {
			return errors.New("验证码不正确")
		} else {
			return nil
		}
	}
}
