import  request  from "./index"
export interface ILoginData {
  /** admin 或 editor */
  username?: string
  /** 密码 */
  password: string
  passwordOld?: string
  passwordr?: string
  mobile?: string
  smscode?: string
  /** 验证码 */
  captchaKey?: string
  captchaCode?: string
  orgname?: string
  linker?: string
  tel?: string
  uuid?:string
  code?:string
  type: "username" | "mobile" | "smscode" | "userId"
}

/** 获取登录验证码 */
export function getCaptcha() {
  return request({
    url: "/captcha/get",
    method: "get"
  })
}
/** 登录并返回 Token */
export function loginApi(data: ILoginData) {
  return request({
    url: "/account/login",
    method: "post",
    data
  })
}
/** 登录并返回 Token */
export function logoutApi() {
    return  Promise.resolve()
}
/** 登录并返回 Token */
export function registerApi(data: ILoginData) {
  return request({
    url: "/account/register",
    method: "post",
    data
  })
}

/** 登录并返回 Token */
export function resetMinePwd(data: ILoginData) {
  return request({
    url: "/account/resetMinePwd",
    method: "post",
    data
  })
}

/** 获取用户详情 */
export function getInfoApi() {
  return request({
    url: "/account/getInfo",
    method: "get"
  })
}

  /**
   * 获取路由列表
   *
   * @returns 路由列表
   */
  export function getRoutes(authId:string|number){
    return request({
      url: `/account/getMenu`,
      method: "post",
      data:{authId},
    });
  }

export function updateMine(data:any) {
  return request({
    url: "userinfo/updateMine",
    method: "post",
    data: data
  })
}

export function resetOtherPassword(data:any) {
  return request({
    url: "userinfo/resetPwd",
    method: "post",
    data: data
  })
}
export default {
  getCaptcha,
  loginApi,logoutApi,
  registerApi,
  getInfoApi,
  resetMinePwd,
  updateMine,
  resetOtherPassword,
  getRoutes,
}
