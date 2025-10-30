
import  request  from "./index"
import {ICond} from "../types"
const prefix = "/config"

export type Config = {
      name : string   // 角色名称
      label : string   // 状态
      value : string   // 资源ID
      id : number   // ID    
      orgId: number
    }
// 
export type IConfig = Partial<Config>
//查询,搜索角色对象
export function search(data:ICond) {
  return request({
    url: prefix +'/search',
    method: 'post',
    data:data
  })
}

//创建角色对象
export function create(data:IConfig) {
  return request({
    url: prefix+'/create',
    method: 'post',
    data
  })
}
//更新角色
export function update(data:IConfig) {
  return request({
    url: prefix+'/update',
    method: 'post',
    data
  })
}
//更新角色
export function grant(data:IConfig) {
  return request({
    url: prefix+'/grant',
    method: 'post',
    data
  })
}
//获取一条角色记录
export function getOne(id : string|number) {
  return request({
    url: prefix+`/getOne`,
    method: 'get',
    params:{ id }
  })
}

//删除某一条角色
export function deleteIt(id : string|number) {

  return request({
    url: prefix+'/delete',
    method: 'post',
    data:{
      id
    }
  })
}

//删除多条角色
export function deleteIts(ids : string[]|number[]) {
  return request({
    url: prefix+'/deleteIts',
    method: 'post',
    data:{
      ids
    }
  })
}

//根据条件导出角色
export function exportxls(cond:any) {
  return request({
    url: prefix+'/export',
    method: 'post',
    responseType: 'blob',
    data:cond
  })
}
//导出角色的meta信息
export function meta(id:any) {
  return request({
    url: prefix+'/meta',
    method: 'post',
    data:{
      id
    }
  })
}
//导出角色的meta信息
export function name(name:any) {
  return request({
    url: prefix+'/name',
    method: 'post',
    data:{
      name
    }
  })
}
//导出角色的meta信息
export function site() {
  return request({
    url: prefix+'/site',
    method: 'post',
    data:{
      
    }
  })
}
//默认导出全部API
export default {search,create,update,deleteIt,deleteIts,grant,getOne,exportxls,meta,site,name}
