package tokenkit

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"turingdance.com/turing/internal/pkg/utils"
)

type TokenManager struct {
	jwtSecret string
}

func NewTokenManager(jwtSecret string) *TokenManager {
	return &TokenManager{
		jwtSecret: jwtSecret,
	}
}

// 注册
func (mgr *TokenManager) Parse(token string) (result Realm, err error) {

	tokenObj, err := jwt.Parse(token, func(tok *jwt.Token) (interface{}, error) {
		// 防御：校验签名算法，禁止 none 攻击
		if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", tok.Header["alg"])
		}
		return []byte(mgr.jwtSecret), nil
	})
	if err != nil {
		return result, fmt.Errorf("parse token failed: %w", err)
	}

	claims, ok := tokenObj.Claims.(jwt.MapClaims)
	if !ok {
		return result, errors.New("claims type convert failed")
	}

	// 校验过期时间
	if err = claims.Valid(); err != nil {
		return result, errors.New("token invalid or expired")
	}

	// 安全读取字符串字段
	user, _ := claims["user"].(string)
	accId, _ := claims["accId"].(string)
	role, _ := claims["role"].(string)
	nickname, _ := claims["nickname"].(string)
	avatar, _ := claims["avatar"].(string)
	// json数字默认float64，安全转换 uint64
	var tenantId uint64
	if v, ok := claims["tenantId"].(float64); ok {
		tenantId = uint64(v)
	}

	// 使用 RealmBuilder 链式构造，BuildPtr 返回指针
	realm := NewRealmBuilder().
		User(user).
		AccId(accId).
		TenantId(tenantId).
		RoleKey(role).
		Avatar(avatar).
		NickName(nickname).Build()
	return realm, nil

}

func (mgr *TokenManager) Generate(realm Realm) (string, error) {
	claim := jwt.MapClaims{
		"iat": time.Now().Unix(),
		"iss": "turing-microapp-server",
		"aud": "turing-microapp-server",
		"nbf": time.Now().Unix(),
		"exp": time.Now().Add(72 * time.Hour).Unix(),
		"sub": "user",
		// 按需放入，白名单
		"user":     realm.User,
		"accId":    realm.AccId,
		"tenantId": realm.TenantId,
		"avatar":   realm.Avatar,
		"nickname": realm.NickName,
		"role":     realm.RoleKey,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString([]byte(mgr.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

var defaultTokenManager *TokenManager = &TokenManager{
	jwtSecret: "gateway@turingdance",
}

func GenerateToken(values Realm) (string, error) {
	return defaultTokenManager.Generate(values)
}
func ParseToken(in interface{}) (result Realm, err error) {
	if token, ok := in.(string); ok {
		return defaultTokenManager.Parse(token)
	} else if req, ok := in.(*http.Request); ok {
		token := utils.GetAuthorizationFromRequest(req)
		return defaultTokenManager.Parse(token)
	} else {
		return result, errors.New("不支持的数据类型")
	}

}
func ParseUserId(req *http.Request) (userId string, e error) {
	token := utils.GetAuthorizationFromRequest(req)
	userInfo, err := ParseToken(token)
	if err != nil {
		return "", err
	} else {
		return userInfo.AccId, err
	}
}
func ParseRole(req *http.Request) (roletype string, e error) {
	token := utils.GetAuthorizationFromRequest(req)
	realm, err := ParseToken(token)
	return realm.RoleKey, err
}
