
import  request  from "./index"
import {ICond} from "../types"
const prefix = "/role"

export type Role = {
      name : string   // 角色名称
    
      state : number   // 状态
      rightIds : any[]   // 资源ID
    
      id : number   // ID
    
      code : string   // 角色码
    }
// 
export type IRole = Partial<Role>
//查询,搜索角色对象
export function search(data:ICond) {
  return request({
    url: prefix +'/search',
    method: 'post',
    data:data
  })
}

//创建角色对象
export function create(data:IRole) {
  return request({
    url: prefix+'/create',
    method: 'post',
    data
  })
}
//更新角色
export function update(data:IRole) {
  return request({
    url: prefix+'/update',
    method: 'post',
    data
  })
}
//更新角色
export function grant(data:IRole) {
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
//默认导出全部API
export default {search,create,update,deleteIt,deleteIts,grant,getOne,exportxls,meta}
