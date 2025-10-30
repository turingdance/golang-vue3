import  request  from "./index"
import {ICond} from "../types"
import {useStorageStore} from "@/store"
const prefix = "/storage"

export type StorageConf = {
        driver: "local"|"oss"
        bucket:string
        datapath:string
        ssl:boolean
        host: string
        authRequired:boolean,
        primary: boolean
}
export type IStorageConf = Partial<StorageConf>

export function config() {
    return request({
      url: prefix+'/config',
      method: 'get',
    })
}

export function shareUrl(filekey) {
    return request({
      url: prefix+'/shareUrl',
      method: 'get',
      params:{filekey}
    }).then(res=>res.data)
}

export default {
    config,shareUrl
}  