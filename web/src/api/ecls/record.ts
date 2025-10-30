
import  request  from "./index"
import {ICond} from "../types"
const prefix = "/record"

export type Record = {
      id : number   // ID
    
      createAt :  any    // 上课时间
    
      userId : string   // 用户ID
    
      lessonId : string   // 课程ID
    
      createDate :  any    // 创建日期
    }
// 
export type IRecord = Partial<Record>
//查询,搜索上课记录对象
export function search(data:ICond) {
  return request({
    url: prefix +'/search',
    method: 'post',
    data:data
  })
}

//创建上课记录对象
export function create(data:IRecord) {
  return request({
    url: prefix+'',
    method: 'post',
    data
  })
}
//更新上课记录
export function update(data:IRecord) {
  return request({
    url: prefix,
    method: 'put',
    data
  })
}

export function mine(data:ICond) {
  return request({
    url: `${prefix}/mine`,
    method: "post",
    data
  })
}
//获取一条上课记录记录
export function getOne(id : string|number) {
  return request({
    url: prefix+`/`+ id ,
    method: 'get'
  })
}

//删除某一条上课记录
export function deleteIt(id : string|number) {

  return request({
    url: prefix+'/id',
    method: 'delete'
  })
}

//删除多条上课记录
export function deleteIts(ids : string[]|number[]) {
  return request({
    url: prefix,
    method: 'delete',
    data:{
      ids
    }
  })
}

//根据条件导出上课记录
export function exportxls(cond:any) {
  return request({
    url: prefix+'/export',
    method: 'post',
    responseType: 'blob',
    data:cond
  })
}

//默认导出全部API
export default {search,create,update,deleteIt,deleteIts,getOne,exportxls,mine}
