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

export enum Strategy {
  Filename = 1,
  UUID = 2,
  Timestamp = 3,
  Func = 4
}
export type UploadResult = {
  url?: string
  name?: string
  key?: string
  host?: string
  size?: number
}
