package logic

//
import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/turingdance/infra/pkkit"
	"turingdance.com/turing/internal/app/sys/args"
	"turingdance.com/turing/internal/app/sys/model"
	"turingdance.com/turing/internal/pkg/utils"
	"turingdance.com/turing/internal/server/auth"
	"turingdance.com/turing/internal/types"
)

type IAccount interface {
	Register(userName, password string) (bool, error)
	Login(userName, password string) (string, error)
}

type account struct {
}

func NewAccountSve() *account {
	return &account{}
}

// 注册
func (acc *account) RegisterUseMobile(mobile, plainpwd string) (userId string, err error) {
	var user model.Userinfo
	DbEngin.Where("mobile = ?", mobile).First(&user)
	if !user.Empty() {
		return "", errors.New("该用户已经存在")
	}
	user.Mobile = mobile
	user.Status = types.EnableStatusTrue
	user.CreateAt = types.DateTime(time.Now())
	user.RoleId = 0
	user.Password, err = utils.GeneratePassword(plainpwd)
	if err != nil {
		return
	}
	if err = DbEngin.Create(&user).Error; err != nil {
		return user.UserId, err
	} else {
		return user.UserId, err
	}
}

// 注册
func (acc *account) RegisterUseUsername(username, plainpwd string) (result string, err error) {
	var user model.Userinfo
	DbEngin.Where("username = ?", username).First(&user)
	if !user.Empty() {
		return "", errors.New("该用户已经存在")
	}
	user.Status = types.EnableStatusTrue
	user.CreateAt = types.DateTime(time.Now())
	user.RoleId = 0
	user.Password, err = utils.GeneratePassword(plainpwd)
	if err != nil {
		return
	}
	if err = DbEngin.Create(&user).Error; err != nil {
		return user.UserId, err
	} else {
		return user.UserId, err
	}
}

// 注册
func (acc *account) Register(arg args.Account) (userId string, err error) {
	if arg.UseMobile() {
		// 需要校验短信
		if err := VerifySms(arg.Mobile, arg.Code); err != nil {
			return "", err
		}
		return acc.RegisterUseMobile(arg.Mobile, arg.Password)
	} else {
		// 需要校验图片验证码
		// 需要校验短信
		r, err := CheckCaptcha(arg.Code, arg.UUid)
		if err != nil {
			return "", err
		}
		if !r {
			return "", errors.New("验证码不正确")
		}
		return acc.RegisterUseUsername(arg.Username, arg.Password)
	}
}

// 注册
func (acc *account) BatchCreate(users []model.Userinfo) (total int, err error) {
	createAt := types.DateTime(time.Now())
	for i, v := range users {
		users[i].Password, _ = utils.GeneratePassword(v.Password)
		users[i].CreateAt = createAt
		users[i].Status = types.EnableStatusTrue
		users[i].RoleId = 0
		users[i].UserId = pkkit.UseSnowflakeID()
		users[i].Status = types.EnableStatusTrue
	}
	err = DbEngin.CreateInBatches(&users, 10).Error
	total = len(users)
	return
}

// 注册
func (a *account) LoginUseSmsCode(acc args.Account) (token string, err error) {
	code := acc.Code
	phone := acc.Mobile
	// 首先检查sms有消息
	if code == "" {
		err = errors.New("请输入短信验证码")
		return
	}
	if phone == "" {
		err = errors.New("请输入手机号")
		return
	}
	// 验证码不正确
	if err = VerifySms(phone, code); err != nil {
		return
	}
	var user model.Userinfo
	DbEngin.Where("mobile = ?", phone).First(&user)
	if user.Empty() {
		// 用户不存在,那么创建一个用户
		user.Nickname = "user" + utils.RandSeq(3)
		user.Mobile = phone
		user.Status = types.EnableStatusTrue
		user.CreateAt = types.DateTime(time.Now())
		user.RoleId = 0
		user.Password, err = utils.GeneratePassword(utils.RandSeq(12))
		if err != nil {
			return
		}
		err = DbEngin.Create(&user).Error
		if err != nil {
			return "", errors.New("该用户不存在")
		}
	}
	//
	if user.Status != types.EnableStatusTrue {
		return "", errors.New("该账号暂未激活请联系管理员")
	}
	orgInfo, _ := FindOrgByUserId(user.UserId)
	// 登录成功了,生成redis
	// jwt 只处理 userId,username,pic,nickname
	//initial.RedisEngin.
	return a.buildToken(orgInfo, user)
}
func (acc *account) buildToken(orgInfo model.Org, user model.Userinfo) (token string, err error) {
	clain := map[string]interface{}{
		"userId":   user.UserId,
		"nickname": user.Nickname,
		// "username": user.Username,
		"pic":    user.Pic,
		"roleId": user.RoleId,
		"orgId":  orgInfo.OrgId,
	}
	return auth.GenerateToken(clain)
}

func ParseAccountFrom(input any) (user model.Userinfo, e error) {
	user = model.Userinfo{}
	if resultMap, err := auth.ParseToken(input); err == nil {
		user.UserId = resultMap["userId"].(string)
		user.Nickname = resultMap["nickname"].(string)
		user.OrgId = utils.Atoi[uint](fmt.Sprintf("%f", resultMap["orgId"].(float64)))
		user.RoleId = uint(resultMap["roleId"].(float64))
	} else {
		e = err
	}
	return user, e
}

func ParseUserId(req *http.Request) (userId string, e error) {
	token := utils.GetAuthorizationFromRequest(req)
	userInfo, err := ParseAccountFrom(token)
	if err != nil {
		return "", err
	} else {
		return userInfo.UserId, err
	}
}
func ParseRole(req *http.Request) (userId model.RoleType, e error) {
	token := utils.GetAuthorizationFromRequest(req)
	userInfo, err := ParseAccountFrom(token)
	if err != nil {
		return "", err
	} else {
		tmp := model.Role{}
		DbEngin.Where("id = ?", userInfo.RoleId).Find(&tmp)
		return tmp.Code, err
	}
}

// 注册
func (acc *account) LoginUseUserId(userId string) (token string, err error) {
	var user model.Userinfo
	DbEngin.Where("user_id = ?", userId).First(&user)
	if user.Empty() {
		return "", errors.New("该用户不存在")
	}
	orgInfo, _ := FindOrgByUserId(user.UserId)
	// 登录成功了,生成redis
	// jwt 只处理 userId,username,pic,nickname
	//initial.RedisEngin.
	return acc.buildToken(orgInfo, user)
}

// 用户名 密码登录
func (a *account) RestPwdUseUserId(acc args.Account) (err error) {
	var user model.Userinfo
	DbEngin.Where("user_id = ?", acc.UserId).First(&user)
	if r, _ := utils.ComparePassword(user.Password, acc.PasswordOld); !r {
		return errors.New("原始密码不正确")
	}
	if user.Password, err = utils.GeneratePassword(acc.Password); err != nil {
		return errors.New(err.Error())
	}
	err = DbEngin.Where("user_id = ?", user.UserId).Model(&user).UpdateColumn("password", user.Password).Error
	return err
}

// 注册
func (a *account) RestPwdUseSmsCode(acc args.Account) (err error) {
	code := acc.Code
	phone := acc.Mobile
	// 首先检查sms有消息
	if code == "" {
		err = errors.New("请输入短信验证码")
		return
	}
	if phone == "" {
		err = errors.New("请输入手机号")
		return
	}
	// 验证码不正确
	if err = VerifySms(phone, code); err != nil {
		return
	}
	var user model.Userinfo
	DbEngin.Where("mobile = ?", phone).First(&user)
	if user.Empty() {
		return errors.New("该用户不存在")
	}
	if user.Password, err = utils.GeneratePassword(acc.Password); err != nil {
		return errors.New(err.Error())
	}
	err = DbEngin.Where("user_id = ?", user.UserId).Model(&user).UpdateColumn("password", user.Password).Error
	return err
}

// 用户名 密码登录
func (a *account) LoginUseUsername(acc args.Account) (token string, err error) {
	r, err := CheckCaptcha(acc.Code, acc.UUid)
	if err != nil {
		return "", errors.New("验证码不正确")
	}
	if !r {
		return "", errors.New("验证码不正确")
	}
	var user model.Userinfo
	DbEngin.Where("username = ?", acc.Username).First(&user)
	if user.Empty() {
		return "", errors.New("该用户不存在")
	}
	if user.Status != types.EnableStatusTrue {
		return "", errors.New("该账号暂未激活请联系管理员")
	}
	if r, err := utils.ComparePassword(user.Password, acc.Password); err != nil {
		return "", errors.New(err.Error())
	} else if !r {
		return "", errors.New("该用户密码不正确")

	}
	orgInfo, _ := FindOrgByUserId(user.UserId)
	// 登录成功了,生成redis
	// jwt 只处理 userId,username,pic,nickname
	//initial.RedisEngin.
	return a.buildToken(orgInfo, user)
}

// 用户名 手机号密码登录
func (a *account) LoginUseMobile(acc args.Account) (token string, err error) {
	r, err := CheckCaptcha(acc.Code, acc.UUid)
	if err != nil {
		return "", errors.New("验证码不正确")
	}
	if !r {
		return "", errors.New("验证码不正确")
	}
	var user model.Userinfo
	DbEngin.Where("mobile = ?", acc.Mobile).First(&user)
	if user.Empty() {
		return "", errors.New("该用户不存在")
	}
	if user.Status != types.EnableStatusTrue {
		return "", errors.New("该账号暂未激活请联系管理员")
	}
	if r, err := utils.ComparePassword(user.Password, acc.Password); err != nil {
		return "", errors.New(err.Error())
	} else if !r {
		return "", errors.New("该用户密码不正确")

	}
	orgInfo, _ := FindOrgByUserId(user.UserId)
	// 登录成功了,生成redis
	// jwt 只处理 userId,username,pic,nickname
	//initial.RedisEngin.
	return a.buildToken(orgInfo, user)
}

// 查询一条
func Login(acc args.Account) (result string, err error) {
	accsvc := NewAccountSve()
	if acc.UseUsername() {
		return accsvc.LoginUseUsername(acc)
	} else if acc.UseMobile() {
		return accsvc.LoginUseMobile(acc)
	} else if acc.UseSmsCode() {
		return accsvc.LoginUseSmsCode(acc)
	} else {
		return "", errors.New("不支持的登录方法")
	}
}

// 查询一条
func RestPwd(acc args.Account) (err error) {
	accsvc := NewAccountSve()
	if acc.UseUsername() {
		return accsvc.RestPwdUseUserId(acc)
	} else if acc.UseSmsCode() {
		return accsvc.RestPwdUseSmsCode(acc)
	} else {
		return errors.New("不支持的密码重置方法")
	}
}

func UpdatePassword(acc args.Account, compareold bool) (err error) {
	user, err := FindUserinfo(acc.UserId)
	if user.UserId == "" {
		return errors.New("缺少用戶ID")
	}
	if compareold {
		ok, err := utils.ComparePassword(user.Password, acc.PasswordOld)
		if err != nil {
			return err
		}
		if !ok {
			return errors.New("原始密码不正确")
		}
	}
	hashedpassword, err := utils.GeneratePassword(acc.Password)
	if err != nil {
		return err
	}
	return DbEngin.Model(new(model.Userinfo)).Where("user_id", acc.UserId).UpdateColumn("password", hashedpassword).Error

}

// 查询一条
func RestMobile(acc args.Account) (err error) {
	if err := VerifySms(acc.Mobile, acc.Code); err != nil {
		return err
	}
	user, err := FindUserinfoByMobile(acc.Mobile)

	if user.UserId != "" {
		return errors.New("当前手机号已存在,无法设置")
	}
	user.Mobile = acc.Mobile
	return DbEngin.Model(new(model.Userinfo)).UpdateColumn("mobile", acc.Mobile).Where("user_id", acc.UserId).Error
}

// 查询一条
func LoginUseUserId(userId string) (result string, err error) {
	accsvc := NewAccountSve()
	return accsvc.LoginUseUserId(userId)
}

// 查询一条
func Register(acc args.Account) (userId string, err error) {
	accsvc := NewAccountSve()
	return accsvc.Register(acc)
}

// 解析
func Userinfo(token string) (result model.Userinfo, err error) {
	mapresult, err := auth.ParseToken(token)
	if err != nil {
		return
	}
	var userInfo model.Userinfo
	bts, _ := json.Marshal(mapresult)
	err = json.Unmarshal(bts, &userInfo)
	ret, err := FindUserinfo(userInfo.UserId)
	ret.Password = ""
	ret.Mobile = ""
	return ret, err
}

// 解析
func UserId(token string) (userId string, err error) {
	mapresult, err := auth.ParseToken(token)
	if err != nil {
		return
	}
	var userInfo model.Userinfo
	bts, _ := json.Marshal(mapresult)
	err = json.Unmarshal(bts, &userInfo)
	return userInfo.UserId, err
}
