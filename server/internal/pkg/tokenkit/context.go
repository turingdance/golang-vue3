package tokenkit

import (
	"context"
	"errors"
	"net/http"
)

// 私有key类型，防止冲突
type ctxKeyType string

// 常量key，包内私有
const realmKey ctxKeyType = "_turing_realm_data_"

// SetRealm 将 realm 存入context，返回新context
func SetRealm(ctx context.Context, realm *Realm) context.Context {
	return context.WithValue(ctx, realmKey, realm)
}

// GetRealm 从上下文取出 Realm指针
func GetRealm(ctx context.Context) (*Realm, error) {
	val := ctx.Value(realmKey)
	if val == nil {
		return nil, errors.New("当前用户尚未登陆,请先鉴权")
	}
	realm, ok := val.(*Realm)
	if !ok {
		return nil, errors.New("上下文用户数据类型异常")
	}
	return realm, nil
}

func GetRequestRealm(req *http.Request) (*Realm, error) {
	return GetRealm(req.Context())
}

func SetRequestRealm(req *http.Request, realm *Realm) context.Context {
	ctx := req.Context()
	return SetRealm(ctx, realm)
}
