
import  request  from "./index"
import {ICond} from "../types"
const prefix = "/lesson"

export type Lesson = {
      id : string   // 视频ID
    
      media : string   // 资源类型
    
      cover : string   // 封面
    
      title : string   // 标题
    
      memo : string   // 描述
    
      uri : string   // 资源key
    
      tag : string   // 标签
    }
// 
export type ILesson = Partial<Lesson>
//查询,搜索课件资源对象
export function search(data:ICond) {
  return request({
    url: prefix +'/search',
    method: 'post',
    data:data
  })
}

//创建课件资源对象
export function create(data:ILesson) {
  return request({
    url: prefix+'',
    method: 'post',
    data
  })
}
//创建课件资源对象
export function batchAdd(data:ILesson[]) {
  return request({
    url: prefix+'/batchAdd',
    method: 'post',
    data:{data}
  })
}
//更新课件资源
export function update(data:ILesson) {
  return request({
    url: prefix,
    method: 'put',
    data
  })
}
//获取一条课件资源记录
export function getOne(id : string|number) {
  return request({
    url: prefix+`/`+ id ,
    method: 'get'
  })
}

//删除某一条课件资源
export function deleteIt(id : string|number) {

  return request({
    url: prefix+'/id',
    method: 'delete'
  })
}

//删除多条课件资源
export function deleteIts(ids : string[]|number[]) {
  return request({
    url: prefix,
    method: 'delete',
    data:ids
  })
}

//根据条件导出课件资源
export function exportxls(cond:any) {
  return request({
    url: prefix+'/export',
    method: 'post',
    responseType: 'blob',
    data:cond
  })
}

//默认导出全部API
export default {search,create,batchAdd,update,deleteIt,deleteIts,getOne,exportxls}
