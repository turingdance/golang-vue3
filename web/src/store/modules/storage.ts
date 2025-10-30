import { store } from "@/store";
import storageApi,{IStorageConf} from "@/api/sys/storage"
// 保存到storage
export const useStorageStore = defineStore("storage", () => {
  const config = ref<IStorageConf[]>([]);

  /**
   * 登录
   *
   * @param {LoginData}
   * @returns
   */
  function init() {
  return  storageApi.config().then(res=>{
    config.value = res.data ||[]
      return config.value
    })
  }

 function shareUrl(key){
    if(key.indexOf("https://")>-1 || key.indexOf("http://")>-1){
      return Promise.resolve(key)
    }
    let _config = config.value.find(cfg=>key.indexOf(cfg.bucket+"/")==0)
    if(!_config){
      return Promise.reject(new Error("配置文件不存在"))
    }
    //拿到config
    if(!_config!.authRequired){
        var url = (_config!.ssl?"https://":"http://")+_config!.host+"/"+key
        return Promise.resolve(url)
    }else{
      return storageApi.shareUrl(key)
    }
}

function getConfig(bucket){
  return config.value.find(cfg=>cfg.bucket==bucket)
}

  return {
    config,getConfig,
    init,shareUrl,
  };
});

// 非setup
export function useStorageStoreHook() {
  return useStorageStore(store);
}
