import  request  from "./index"

/** 获取登录验证码 */
export function getCaptcha() {
  return request({
    url: "/captcha/get",
    method: "get"
  })
}
export default {
  getCaptcha,
}
