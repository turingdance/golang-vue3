import uploadtolocal from "./uploadtolocal"
import uploadtooss from "./uploadtooss"
import {useStorageStore} from "@/store"
import {UploadResult} from "./types"
var obj = uploadtooss
const upload = (file,bucket="mnt"):Promise<UploadResult[]>=>{
  const storageStore = useStorageStore()
  const config = storageStore.getConfig(bucket)
  if(config!.driver=="oss"){
    return uploadtooss.upload(file,bucket)
  }else{
    return uploadtolocal.upload(file,bucket)
  }
}
const uploadbase64 = (file,bucket="mnt"):Promise<UploadResult[]>=>{
  const storageStore = useStorageStore()
  const config = storageStore.getConfig(bucket)
  if(config!.driver=="oss"){
    return uploadtooss.uploadbase64(file)
  }else{
    return uploadtolocal.uploadbase64(file,bucket)
  }
}

export default {
  upload,uploadbase64
}
