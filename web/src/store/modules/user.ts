import { resetRouter } from "@/router";
import { store } from "@/store";

import authAPI,{ ILoginData } from "@/api/sys/login";
import { IUserInfo } from "@/api/sys/acc";
import tokenmgr from "@/utils/tokenmgr"
import {useDictStore} from "./dict"
export const useUserStore = defineStore("user", () => {
  const user = ref<IUserInfo>({
    roles: [],
    perms: [],
    authId:0,
    pic:"",
    nickname:"",
    teamId:0,
    username:"",
    mobile:"",
    email:"",
  });

  const haslogin = ref(false)

  /**
   * 登录
   *
   * @param {LoginData}
   * @returns
   */
  function login(loginData: ILoginData) {
    return new Promise<void>((resolve, reject) => {
      authAPI.loginApi(loginData)
        .then((res) => {
          const { data }= res
          tokenmgr.setToken(data); // Bearer eyJhbGciOiJIUzI1NiJ9.xxx.xxx
          haslogin.value = true
          resolve();
        })
        .catch((error) => {
          reject(error);
        });
    });
  }

  // 获取信息(用户昵称、头像、角色集合、权限集合)
  function getUserInfo() {
    return new Promise<IUserInfo>((resolve, reject) => {
      authAPI.getInfoApi()
        .then(({data}) => {
           
          if (!data) {
            reject("校验失败,请重试");
            return;
          }
          haslogin.value = true
          var {pic,nickname,username,mobile,email,orgId,roleId,scope} = data
          var roles = []as string[]
          var perms =[] as string[]
          var authId = 0
          var teamId = 0
          var auth = data.role
             roles.push(auth.code)
             perms = perms.concat(scope||[])
             authId = roleId
             teamId = orgId
          if (roles.length <= 0) {
            reject("当前用户尚未分配权限");
            return;
          }
          
          Object.assign(user.value, { roles,perms,authId,teamId,pic,nickname,username,mobile,email});
          useDictStore().init()
          resolve({ roles,perms,authId,teamId,pic,nickname,username,mobile,email});
        })
        .catch((error) => {
          reject(error);
        });
    });
  }
  
  // user logout
  function logout() {
    return new Promise<void>((resolve, reject) => {
      authAPI.logoutApi()
        .then(() => {
          tokenmgr.removeToken()
          haslogin.value = false 
          location.reload(); // 清空路由
          resolve();
        })
        .catch((error) => {
          reject(error);
        });
    });
  }

  // remove token
  function resetToken() {
    //console.log("resetToken");
    return new Promise<void>((resolve) => {
      tokenmgr.setToken('')
      resetRouter();
      resolve();
    });
  }

  return {
    user,
    haslogin,
    teamId:user.value.teamId,
    login,
    getUserInfo,
    logout,
    resetToken,
  };
});

// 非setup
export function useUserStoreHook() {
  return useUserStore(store);
}
