import  request  from "./index"
const prefix = `/websocket`

const context = (data: any) => {
  return request({
    url: prefix + "/context",
    method: "get",
    params: data
  })
}

const verify = (data: any) => {
  return request({
    url: prefix + "/verify",
    method: "post",
    data
  })
}

const publish = (data: any) => {
  return request({
    url: prefix + "/publish",
    method: "post",
    data
  })
}

//默认输出api对象
export default { context, verify, publish }
