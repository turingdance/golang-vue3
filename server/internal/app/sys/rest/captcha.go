package rest

import (
	"net/http"

	"turingdance.com/turing/internal/app/sys/logic"

	"github.com/turingdance/infra/wraper"
)

type Captcha struct {
}

func init() {
	//RegisterRestCtrl(&Captcha{})
}

// 这是router
// @Router /captcha/get [GET]
func (ctrl *Captcha) Get(w http.ResponseWriter, r *http.Request) {
	if _, b64, tb64, key, err := logic.GenerateCaptcha(); err != nil {
		wraper.Error(err).Encode(w)
	} else {
		wraper.OkData(map[string]interface{}{
			"image_base64": b64,
			"thumb_base64": tb64,
			"captcha_key":  key,
		}).Encode(w)
	}
}
