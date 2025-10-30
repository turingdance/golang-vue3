import { defineStore } from "pinia"
import storage from "@/utils/storage"
import dictApi, { IDictItem } from "@/api/sys/dict"
const datamapkey = "dictdata"
export const useDictStore = defineStore("dict", () => {
  let dictMap = ref({} as Record<string, IDictItem[]>)
  const init = () => {
    dictApi.listAll().then((res) => {
      let tmp = (res.rows||[]).reduce((ret,cur)=>{
        ret[cur.name]=cur.value
        return ret
      },{})
      dictMap.value = tmp
      storage.set(datamapkey, dictMap.value)
    })
  }

  const getDictValue = function (name: string) {
    //console.log('dictMap',dictMap)
    return dictMap.value[name]
  }
  return { init, getDictValue, dictMap }
})
