
import  request  from "./index"
import {ICond} from "../types"
const prefix = "/auth"

export type Auth = {
      expireAt :  any    // 权限过期时间
    
      id : number   // ID
    
      accId : number   // 用户ID
    
      roleId : number   // 角色ID
    
      teamId : number   // 所属团队
    
      status : number   // 权限状态
    }
// 
export type IAuth = Partial<Auth>
//查询,搜索权限对象
export function search(data:ICond) {
  return request({
    url: prefix +'/search',
    method: 'post',
    data:data
  })
}

//创建权限对象
export function create(data:IAuth) {
  return request({
    url: prefix+'/create',
    method: 'post',
    data
  })
}
//更新权限
export function update(data:IAuth) {
  return request({
    url: prefix+'/update',
    method: 'post',
    data
  })
}
//获取一条权限记录
export function getOne(id : string|number) {
  return request({
    url: prefix+`/getOne`,
    method: 'get',
    params:{ id }
  })
}

//删除某一条权限
export function deleteIt(id : string|number) {

  return request({
    url: prefix+'/delete',
    method: 'post',
    data:{
      id
    }
  })
}

//删除多条权限
export function deleteIts(ids : string[]|number[]) {
  return request({
    url: prefix+'/deleteIts',
    method: 'post',
    data:{
      ids
    }
  })
}

//根据条件导出权限
export function exportxls(cond:any) {
  return request({
    url: prefix+'/export',
    method: 'post',
    responseType: 'blob',
    data:cond
  })
}
//导出权限的meta信息
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
export default {search,create,update,deleteIt,deleteIts,getOne,exportxls,meta}
