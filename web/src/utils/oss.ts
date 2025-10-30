import request from "./request"
import storage from "./storage"
import { fileHash, fileSuffix } from "./file"
import { base64ToFile } from "./base64"
import { datetimeformat } from "./index"
import { customAlphabet } from "nanoid"
const nanoid = customAlphabet("0123456789abcdefghijklmnopqrstuvwxyz", 24)
export interface BaseResponse<T = undefined> {
  status: number
  data: T
  msg: string
}
// oss配置
export interface OssConfig {
  accessid: string
  host: string
  expire: number
  signature: string
  policy: string
  dir: string
  xvar: {}
}

/**
 * 从服务端获取OSS配置
 * @returns {Promise<AxiosResponse<BaseResponse<OssConfig>>>>}
 */
export function policy() {
  return request({
    url: "/oss/policy",
    method: "get"
  })
}

/**
 * 初始化Oss配置
 * @returns {Promise<OssConfig>}
 */
function initOssPolicy() {
  return new Promise((resolve, reject) => {
    // 从本地localstore从获取配置
    const conf = storage.get<OssConfig>("ossConfig")
    if (conf) {
      //  配置存在并且距离过期时间还大于3秒则返回此配置
      const now = new Date().getTime() / 1000
      if (conf.expire - 3 > now) {
        resolve(conf)
        return
      }
    }
    // 请求接口返回配置数据
    policy()
      .then((resp) => {
        if (resp.code === 200) {
          resolve(resp.data)
          // 配置数据写入本地store
          storage.set("ossConfig", resp.data)
        } else {
          reject(resp)
        }
      })
      .catch((err) => {
        //console.log(err)
        reject(err)
      })
  })
}

export function uploadbase64(b64data: string) {
  return upload(base64ToFile(b64data))
}
export enum Strategy {
  Filename = 1,
  UUID = 2,
  Timestamp = 3,
  Func = 4
}

const buildFilename = (
  file: File,
  strategy: Strategy = Strategy.UUID,
  func: ((File) => string | undefined) | undefined = undefined
) => {
  const sufix = "." + file.name.split(".").pop()
  if (strategy == Strategy.Filename) {
    return file.name
  } else if (strategy == Strategy.UUID) {
    return nanoid() + sufix
  } else if (strategy == Strategy.Timestamp) {
    return new Date().getTime() + sufix
  } else if (strategy == Strategy.Func) {
    return func!(file)
  } else {
    return nanoid() + sufix
  }
}
/**
 * 往Oss上传文件
 * @param file File对象
 * @returns {Promise<{name:string,host:string}>}
 */
export function upload(file: File, keystrategy: Strategy = Strategy.UUID) {
  return new Promise((resolve, reject) => {
    initOssPolicy()
      .then(async (conf) => {
        let ossConf = conf as OssConfig
        const config = {
          timeout: 1000 * 60 * 10,
          headers: { "Content-Type": "multipart/form-data" }
        }
        // 计算文件Hash 避免多余的文件上传，这样做的目的是尽量少占用OSS的空间
        Object.keys(ossConf.xvar).forEach((k): void => {
          config.headers[`${k}`] = ossConf.xvar[`${k}`]
        })
        const name = buildFilename(file, keystrategy)
        const formData = new FormData()
        const ymd = datetimeformat(new Date(), "YYYYMMDD")
        const key = `${ossConf.dir}${ymd}/${name}`
        formData.append("name", name as string)
        formData.append("key", key)
        formData.append("policy", ossConf?.policy)
        formData.append("OSSAccessKeyId", ossConf?.accessid)
        formData.append("success_action_status", "200")
        formData.append("signature", ossConf.signature)
        formData.append("file", file)
        request({
          url: ossConf.host,
          method: "post",
          ...config,
          data: formData
        })
          .then((resp) => {
            resolve({ name: key, host: ossConf.host, url: ossConf.host + "/" + key, key: key })
          })
          .catch((err) => reject(err))
      })
      .catch((err) => {
        //console.log(err)
        reject({ msg: "上传失败" })
      })
  })
}

export default {
  upload,
  uploadbase64
}
