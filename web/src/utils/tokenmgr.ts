/** 统一处理 Cookie */

import {TOKEN_KEY} from "@/enums/CacheEnum"
import store from "./storage"

const getToken = () => {
  return store.get(TOKEN_KEY)
}
 const setToken = (token: string) => {
  store.set(TOKEN_KEY, token)
}
 const removeToken = () => {
  store.remove(TOKEN_KEY)
}

export default {
  getToken,setToken,removeToken
}
