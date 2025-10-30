import { store } from "@/store";
import configApi,{IConfig} from "@/api/sys/config"
// 保存到storage
export const useConfigStore = defineStore("config", () => {
  const sitemap = ref<Record<string,string>>({});
  /**
   * 登录
   *
   * @param {LoginData}
   * @returns
   */
  function init() {
  return  configApi.site().then(res=>{
    let tmp = res.rows.reduce((ret,cur)=>{
      ret[cur.name]=cur.value
      return ret
    },{})
    sitemap.value = tmp
    return tmp
    })
  }

  return {
    sitemap,
    init
  };
});

// 非setup
export function useConfigStoreHook() {
  return useConfigStore(store);
}
