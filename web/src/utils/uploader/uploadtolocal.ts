import  {sysrequest as request}  from "../request"
import { base64ToFile } from "../base64"
import { UploadResult } from "./types"

export function uploadbase64(b64data: string,bucket="mnt"):Promise<UploadResult[]> {
  return upload(base64ToFile(b64data),bucket)
}

/*
 * 往Oss上传文件
 * @param file File对象
 * @returns {Promise<{name:string,host:string}>}
 */
export function upload(obj,bucket="mnt"):Promise<UploadResult[]> {
  return new Promise((resolve, reject) => {
    const config = {
      timeout: 0,
      headers: { "Content-Type": "multipart/form-data" }
    }
    
    const formData = new FormData()
    formData.append("bucket",bucket)
    for(var i=0; i<obj.length; i++){
        formData.append("file", obj[i])
    }
    
    request({
      url: "storage/upload",
      method: "post",
      ...config,
      data: formData
    })
      .then((resp) => {
        resolve(resp.data as UploadResult[])
      })
      .catch((err) => reject(err))
  })
}
export default {
  upload,
  uploadbase64
}
