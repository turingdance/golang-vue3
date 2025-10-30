import  request  from "./index"
import { ICond } from "../types";
const prefix = "/rights";

import { MenuTypeEnum } from "@/enums/MenuTypeEnum";

/** 菜单查询参数 */
export interface MenuQuery {
  /** 搜索关键字 */
  keywords?: string;
}

export type Res = {
      type: MenuTypeEnum; // 资源类型CATALOG,MENU,EXTLINK,BUTTON
    
      component: string; // 组件地址
    
      hidden: number; // 是否可见

  sortIndex: number; // 排序位

  pid: number; // 父级资源

  name: string; // 路由名称

  title: string; // 资源名称

  icon: string; // 图标名称

  redirect: string; // 跳转地址

  uri: string; // 路由地址
  biz: string; // 权限标识
  id: number; // 资源ID
  children: Res[];
};
//

export type Rights =Partial<Res>

interface KeyValue {
  key: string;
  value: string;
}

/** RouteVO，路由对象 */
export interface RouteVO {
  /** 子路由列表 */
  children: RouteVO[];
  /** 组件路径 */
  component?: string;
  /** 路由属性 */
  meta?: Meta;
  /** 路由名称 */
  name?: string;
  /** 路由路径 */
  path?: string;
  /** 跳转链接 */
  redirect?: string;
}
// 自由适配
export function BuildResTORoute(resTree: any) {
  if (resTree instanceof Array) {
    const result = [] as any[];
    for (let i = 0; i < resTree.length; i++) {
      const tmp = BuildOneResTORoute(resTree[i]);
      result.push(tmp);
    }
    return result;
  } else {
    const result = BuildOneResTORoute(resTree);
    return result;
  }
}
// 单个
export function BuildOneResTORoute(resTree: any) {
  const result = {};
  _buildOneResTORoute(resTree, result);
  return result;
}
// 递归调用
function _buildOneResTORoute(tree: any, routeVo: any) {
  let resTree = Object.assign(tree,tree.rights)

  const meta = {
    alwaysShow: !!resTree.alwaysShow,
    hidden: !!resTree.hidden,
    icon: resTree.icon,
    keepAlive: !!resTree.keepAlive,
    title: resTree.title,
  } as Meta;
  routeVo.meta = meta;
  routeVo.name = resTree.biz;
  routeVo.path = resTree.uri;
  routeVo.redirect = resTree.redirect;
  routeVo.component = resTree.component;
  if (!!resTree.children && resTree.children.length > 0) {
    for (let i = 0; i < resTree.children.length; i++) {
      const tmp = resTree.children[i];
      const tmp2 = {} as RouteVO;
      const meta1 = {
        alwaysShow: !!tmp.alwaysShow,
        hidden: !!tmp.hidden,
        icon: tmp.icon,
        keepAlive: !!tmp.keepAlive,
        title: tmp.title,
      } as Meta;
      tmp2.meta = meta1;
      tmp2.name = resTree.biz;
      tmp2.path = tmp.uri;
      tmp2.redirect = tmp.redirect;
      tmp2.component = tmp.component;
      if (!routeVo.children) {
        routeVo.children = [] as RouteVO[];
      }
      routeVo.children.push(tmp2);
      _buildOneResTORoute(resTree.children[i], tmp2);
    }
  }
}

/** Meta，路由属性 */
export interface Meta {
  /** 【目录】只有一个子路由是否始终显示 */
  alwaysShow?: boolean;
  /** 是否隐藏(true-是 false-否) */
  hidden?: boolean;
  /** ICON */
  icon?: string;
  /** 【菜单】是否开启页面缓存 */
  keepAlive?: boolean;
  /** 拥有路由权限的角色编码 */
  roles?: string[];
  /** 路由title */
  title?: string;
}
export type IRes = Partial<Res>;
//查询,搜索资源对象
export function search(data:ICond) {
  return request({
    url: prefix +'/search',
    method: 'post',
    data:data
  })
}

//创建资源对象
export function create(data:IRes) {
  return request({
    url: prefix+'/create',
    method: 'post',
    data
  })
}
//更新资源
export function update(data:IRes) {
  return request({
    url: prefix+'/update',
    method: 'post',
    data
  })
}
//获取一条资源记录
export function getOne(id : string|number) {
  return request({
    url: prefix+`/getOne`,
    method: 'get',
    params:{ id }
  })
}

//删除某一条资源
export function deleteIt(id : string|number) {

  return request({
    url: prefix+'/delete',
    method: 'post',
    data:{
      id
    }
  })
}

//删除多条资源
export function deleteIts(ids : string[]|number[]) {
  return request({
    url: prefix+'/deleteIts',
    method: 'post',
    data:{
      ids
    }
  })
}

//根据条件导出资源
export function exportxls(cond:any) {
  return request({
    url: prefix+'/export',
    method: 'post',
    responseType: 'blob',
    data:cond
  })
}
//导出资源的meta信息
export function meta(id:any) {
  return request({
    url: prefix+'/meta',
    method: 'post',
    data:{
      id
    }
  })
}
//导出资源的meta信息
export function tree() {
  return request({
    url: prefix+'/tree',
    method: 'post',
    data:{},
  })
}
//默认导出全部API
export default {
  search,
  create,
  update,
  deleteIt,
  getOne,
  exportxls,
  meta,
  tree,
  BuildResTORoute,
};
