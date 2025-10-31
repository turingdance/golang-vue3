package auth

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/turingdance/infra/slicekit"
	"turingdance.com/turing/internal/pkg/utils"
)

type RouteMatch struct {
	exp     *regexp.Regexp
	Methods []string
}
type Authorize struct {
	routes []RouteMatch
}

// /post/\w+ GET,POST
// /post/10086
func NewAuthorize(list []string) *Authorize {
	re, _ := regexp.Compile(`[\s\,]+`)
	tmp := make([]RouteMatch, 0)
	for _, v := range list {
		arr := re.Split(v, -1)
		patern := arr[0]
		if (len(arr)) < 2 {
			arr = append(arr, "POST", "GET")
		}
		methods := arr[1:]
		retmp := regexp.MustCompile(patern)

		tmp = append(tmp, RouteMatch{
			Methods: methods,
			exp:     retmp,
		})

	}
	return &Authorize{
		routes: tmp,
	}
}

// 正常匹配
func (auth *Authorize) CheckInWhiteList(uri string, method string) bool {
	match := false
	for _, r := range auth.routes {
		// 正则是否匹配
		if r.exp.MatchString(uri) && slicekit.HasSubStrIgnoreCase(r.Methods, method) {
			match = true
			break
		}
	}
	return match
}

// 鉴权
func (auth *Authorize) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// 如果通过白名单
		if auth.CheckInWhiteList(req.RequestURI, req.Method) {
			next.ServeHTTP(w, req)
			return
		}
		// post:/a/{}

		token := utils.GetAuthorizationFromRequest(req)

		//判断段是否满足
		// 如果token 是空的
		if token == "" {
			resp := utils.ResultError("请登录后再试")
			resp.Code = http.StatusForbidden
			err := resp.ResponseJson(w)
			if err != nil {
				fmt.Println(err.Error())
			}
			return
		}
		if _, err := ParseToken(token); err != nil {
			resp := utils.ResultError(err.Error())
			resp.Code = http.StatusForbidden
			resp.ResponseJson(w)
			return
		} else {
			next.ServeHTTP(w, req)
		}
	})
}
