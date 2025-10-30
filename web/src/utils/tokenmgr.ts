/** 统一处理 Cookie */

import {TOKEN_KEY} from "@/enums/CacheEnum"
 const getToken = () => {
  return sessionStorage.getItem(TOKEN_KEY)
}
 const setToken = (token: string) => {
  sessionStorage.setItem(TOKEN_KEY, token)
}
 const removeToken = () => {
  sessionStorage.removeItem(TOKEN_KEY)
}

export default {
  getToken,setToken,removeToken
}
