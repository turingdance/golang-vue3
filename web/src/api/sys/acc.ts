
import  request  from "./index"
import {ICond} from "../types"
const prefix = "/account"

export type Acc = {
      expireAt :  any    // 账户过期时间
    
      username : string   // 用户名
    
      password : string   // 密码
    
      wxOpenId : string   // 微信OPENID
    
      email : string   // 邮箱
    
      createAt :  any    // 创建时间
    
      id : number   // 用户ID
    
      status : number   // 账号状态
    
      mobile : string   // 手机号
      pic:string
      nickname:string
    }
// 
export type IAcc = Partial<Acc>

export type UserInfo = {
    roles: string[],
    perms: string[],
    authId:number,
    pic:string,
    nickname:string,
    teamId: number
    email:string
    username:string
    mobile:string
}
export type IUserInfo = Partial<UserInfo>
//查询,搜索账号对象
export function search(data:ICond) {
  return request({
    url: prefix +'/search',
    method: 'post',
    data:data
  })
}

//创建账号对象
export function create(data:IAcc) {
  return request({
    url: prefix+'/create',
    method: 'post',
    data
  })
}
//更新账号
export function update(data:IAcc) {
  return request({
    url: prefix+'/update',
    method: 'post',
    data
  })
}
//更新账号
export function importbatch(formdata) {
  return request({
    url: prefix+'/import',
    headers:{
      "Content-Type":"multipart/form-data"
    },
    method: 'post',
    data:formdata
  })
}
//更新账号
export function template() {
  return request({
    url: prefix+'/template',
    method: 'get',
    responseType: 'blob',
  })
}
//更新账号
export function updatePwd(data:{userId:string|number,password:string,passwordr?:string,passwordOld?:string}) {
  return request({
    url: prefix+'/updatePwd',
    method: 'post',
    data
  })
}
//更新账号
export function updateMobile(data:{mobile:string,code:string}) {
  return request({
    url: prefix+'/updateMobile',
    method: 'post',
    data
  })
}
//更新账号
export function updateEmail(data:{email:string,uuid:string,code:string}) {
  return request({
    url: prefix+'/updateEmail',
    method: 'post',
    data
  })
}
//获取一条账号记录
export function getOne(id : string|number) {
  return request({
    url: prefix+`/getOne`,
    method: 'get',
    params:{ id }
  })
}

//删除某一条账号
export function deleteIt(userId : string|number) {

  return request({
    url: prefix+'/deleteIt',
    method: 'delete',
    data:{
      userId
    }
  })
}

//删除多条账号
export function deleteIts(userIds : string[]|number[]) {
  return request({
    url: prefix+'/deleteIts',
    method: 'delete',
    data:userIds
  })
}

//根据条件导出账号
export function exportxls(cond:any) {
  return request({
    url: prefix+'/export',
    method: 'post',
    responseType: 'blob',
    data:cond
  })
}
//导出账号的meta信息
export function meta(id:any) {
  return request({
    url: prefix+'/meta',
    method: 'post',
    data:{
      id
    }
  })
}


export function updateUsername(username:any){
  return request({
    url: prefix+'/updateUsername',
    method: 'post',
    data:{username},
  })
}

export function resetemail (email:string){
  return request({
    url: prefix+'/resetemail',
    method: 'post',
    data:{email},
  })
}
//默认导出全部API
export default {template,importbatch,search,create,update,deleteIt,deleteIts,getOne,exportxls,meta,updatePwd,updateUsername,resetemail,resetmobile:updateMobile}
