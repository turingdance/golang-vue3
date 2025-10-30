
import  request  from "./index"
import {ICond} from "../types"
const prefix = "/org"

export type Team = {
      creator : number   // 创建者
    
      trade : string   // 行业
    
      id : number   // 组织id
    
      name : string   // 团队名称
      memo:string
      content:string
      pic:string
    }
// 
export type ITeam = Partial<Team>
//查询,搜索组织对象
export function search(data:ICond) {
  return request({
    url: prefix +'/search',
    method: 'post',
    data:data
  })
}

//创建组织对象
export function create(data:ITeam) {
  return request({
    url: prefix+'/create',
    method: 'post',
    data
  })
}
//更新组织
export function update(data:ITeam) {
  return request({
    url: prefix+'/update',
    method: 'post',
    data
  })
}
//获取一条组织记录
export function getOne(id : string|number) {
  return request({
    url: prefix+`/getOne`,
    method: 'get',
    params:{ id }
  })
}

//删除某一条组织
export function deleteIt(id : string|number) {

  return request({
    url: prefix+'/delete',
    method: 'post',
    data:{
      id
    }
  })
}

//删除多条组织
export function deleteIts(ids : string[]|number[]) {
  return request({
    url: prefix+'/deleteIts',
    method: 'post',
    data:{
      ids
    }
  })
}

//根据条件导出组织
export function exportxls(cond:any) {
  return request({
    url: prefix+'/export',
    method: 'post',
    responseType: 'blob',
    data:cond
  })
}
//导出组织的meta信息
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
