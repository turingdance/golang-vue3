import {createService,createRequestFunction} from "@/utils/request"
/** 用于网络请求的实例 */
const service = createService()
export const request = createRequestFunction(service, "/sys")
export default request
