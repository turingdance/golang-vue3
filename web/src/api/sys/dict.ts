
import  request  from "./index"
import {ICond} from "../types"
import {useDictStore} from "@/store/modules/dict"
const prefix = "/dict"

export type Dict = {
      id : number   // ID
    
      code : string   // 编码
      name : string   // 名称
    
      option : any[]   // 值
    }
export type IDictItem = {
    
      label : string   // 编码
    
      value : string   // 名称
    }
// 
export type IDict = Partial<Dict>
//查询,搜索字典对象
export function search(data:ICond) {
  return request({
    url: prefix +'/search',
    method: 'post',
    data:data
  })
}

//查询,搜索字典对象
export function listAll() {
  return request({
    url: prefix +'/listAll',
    method: 'post',
    data:{}
  })
}

//创建字典对象
export function create(data:IDict) {
  return request({
    url: prefix+'/create',
    method: 'post',
    data
  })
}
//更新字典
export function update(data:IDict) {
  return request({
    url: prefix+'/update',
    method: 'post',
    data
  })
}
//获取一条字典记录
export function getOne(code : string|number) {
  return request({
    url: prefix+`/getOne`,
    method: 'get',
    params:{ code }
  })
}

//获取一条字典记录
export function option(code : string|number) {
  let dictstore = useDictStore()
  let dictItems = dictstore.getDictValue(code as string)
  if(!dictItems){
    return request({
      url: prefix+`/option`,
      method: 'get',
      params:{ code }
    })
  }else{
    return Promise.resolve({code:200,data:dictItems})
  }
  
}

//删除某一条字典
export function deleteIt(id : string|number) {

  return request({
    url: prefix+'/delete',
    method: 'post',
    data:{
      id
    }
  })
}

//删除多条字典
export function deleteIts(ids : string[]|number[]) {
  return request({
    url: prefix+'/deleteIts',
    method: 'post',
    data:{
      ids
    }
  })
}

//根据条件导出字典
export function exportxls(cond:any) {
  return request({
    url: prefix+'/export',
    method: 'post',
    responseType: 'blob',
    data:cond
  })
}
//导出字典的meta信息
export function meta(id:any) {
  return request({
    url: prefix+'/meta',
    method: 'post',
    data:{
      id
    }
  })
}
//默认导出全部API
export default {search,create,update,deleteIt,deleteIts,getOne,exportxls,meta,option,listAll}
