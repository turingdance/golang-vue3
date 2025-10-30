
import  request  from "./index"
import {ICond} from "../types"
const prefix = "/sms"

export type Sms = {
      id : number   // ID
    
      code : string   // 验证码
    
      phoneNumbers : string   // 手机号
    
      templeteId : string   // 模板ID
    
      result : string   // 发送结果
    
      createAt :  any    // 发布时间
    
      success : number   // 是否成功
    
      sendAt :  any    // 发送时间
    }
// 
export type ISms = Partial<Sms>
//查询,搜索短信对象
export function search(data:ICond) {
  return request({
    url: prefix +'/search',
    method: 'post',
    data:data
  })
}

//创建短信对象
export function create(data:ISms) {
  return request({
    url: prefix+'/create',
    method: 'post',
    data
  })
}
//更新短信
export function update(data:ISms) {
  return request({
    url: prefix+'/update',
    method: 'post',
    data
  })
}
//获取一条短信记录
export function getOne(id : string|number) {
  return request({
    url: prefix+`/getOne`,
    method: 'get',
    params:{ id }
  })
}

//删除某一条短信
export function deleteIt(id : string|number) {

  return request({
    url: prefix+'/delete',
    method: 'post',
    data:{
      id
    }
  })
}

//删除多条短信
export function deleteIts(ids : string[]|number[]) {
  return request({
    url: prefix+'/deleteIts',
    method: 'post',
    data:{
      ids
    }
  })
}

//根据条件导出短信
export function exportxls(cond:any) {
  return request({
    url: prefix+'/export',
    method: 'post',
    responseType: 'blob',
    data:cond
  })
}
//导出短信的meta信息
export function meta(id:any) {
  return request({
    url: prefix+'/meta',
    method: 'post',
    data:{
      id
    }
  })
}
//导出短信的meta信息
export function send(data:{mobile:string,captchaKey:string,captchaCode:string}) {
  return request({
    url: prefix+'/send',
    method: 'post',
    data:data
  })
}
//默认导出全部API
export default {search,create,update,deleteIt,deleteIts,getOne,exportxls,meta,send}
