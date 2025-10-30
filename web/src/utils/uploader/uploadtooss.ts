import  request  from "../request"
import store from "../storage"
import { base64ToFile } from "../base64"
import { datetimeformat } from "../index"
import { customAlphabet } from "nanoid"
import { Strategy, UploadResult } from "./types"
const nanoid = customAlphabet("0123456789abcdefghijklmnopqrstuvwxyz", 24)
export interface BaseResponse<T = undefined> {
  status: number
  data: T
  msg: string
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
    const conf = store.get<OssConfig>("ossConfig")
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
          store.set("ossConfig", resp.data)
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

// oss配置
export interface OssConfig {
  accessid: string
  host: string
  expire: number
  signature: string
  policy: string
  dir: string
  xvar: Record<string,any>
}

export function uploadbase64(b64data: string):Promise<UploadResult[]> {
  return upload(base64ToFile(b64data))
}

const buildFilename = (
  file: File,
  strategy: Strategy = Strategy.UUID,
  func: ((file:File) => string | undefined) | undefined = undefined) => {
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
export function upload(file: File,bucket="etrain", keystrategy: Strategy = Strategy.UUID):Promise<UploadResult[]> {
  return new Promise((resolve, reject) => {
    initOssPolicy()
      .then(async (ossConf: any) => {
        const config = {
          timeout: 1000 * 60 * 10,
          headers: { "Content-Type": "multipart/form-data" }
        } as any
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
          .then(() => {
            resolve([{ name: key, key: key, host: ossConf.host, url: ossConf.host + "/" + key }] as UploadResult[])
          })
          .catch((err) => reject(new Error(err)))
      })
      .catch((err) => {
        reject(new Error(err))
      })
  })
}

export default {
  upload,
  uploadbase64
}
