<template>
    <el-button type="primary" v-if="!contextStore.haslogin" @click="loginshow=true" >登录</el-button>
    <el-button type="primary" v-else @click="router.push('/dashboard')" >个人中心</el-button>
     <!-- 登录模态框 -->
 <el-dialog v-model="loginshow" append-to-body :width="dlgwidth" title="用户登录" :z-index="999998">
                 
                 
                 <el-form id="loginForm" :model="loginForm" ref="refForm" label-width="80px" label-position="top">
                     <el-form-item label="用户名" prop="username" :rules="[{ required: true, message: '请输入用户名' }]">
                         <el-input v-model="loginForm.username" placeholder="请输入用户名"></el-input>
                     </el-form-item>
                     
                     <el-form-item label="密码" prop="password" :rules="[{ required: true, message: '请输入密码' }]">
                         <el-input 
                             v-model="loginForm.password" 
                             placeholder="请输入密码" 
                             type="password"
                         >
                         </el-input>
                     </el-form-item>
                     
                     
 
                     <el-form-item>
                             <div class="fm-field w-full" style="margin-top: 30px; margin-bottom: 20px" >
                                 <el-button type="primary" :loading="loading" @click="trylogin" class="w-full">登录</el-button>
                             </div>
                     </el-form-item>
                     
                     <el-form-item class="text-center"  >
                         <span>还没有账号？</span>
                         <el-link href="#" type="primary">立即注册</el-link>
                     </el-form-item>
                 </el-form>
 </el-dialog>
 
 <el-dialog append-to-body v-model="outerVisible"  :width="dlgwidth" :z-index="999999">
                             <Captcha
                             @refresh="handlerefresh"
                             @close="handleclose"
                             @confirm="handleconfirm"
                             :thumb-base64="appdata.thumb"
                             :image-base64="appdata.bgimage"
                             />
 </el-dialog>
 
 </template>
 <script setup>
import { useUserStore } from "@/store/modules/user"
import { getCaptcha as getLoginCodeApi } from "@/api/sys/login"
import Captcha from "@/components/Captcha/index.vue"
import {ref,reactive,watch} from "vue"
import {ElMessage} from "element-plus"
import emiter  from "@/utils/eventbus"
const  loginshow  = ref(false)
const router = useRouter()
const route = useRoute()
emiter.on("showlogin",()=>{
     loginshow.value = true
})
const contextStore = useUserStore()
 const appdata = reactive({
   wxloginIcon:"data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNTAiIGhlaWdodD0iNTAiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiPjxkZWZzPjxmaWx0ZXIgeD0iLTUuNyUiIHk9Ii04LjclIiB3aWR0aD0iMTEyLjQlIiBoZWlnaHQ9IjExNy40JSIgZmlsdGVyVW5pdHM9Im9iamVjdEJvdW5kaW5nQm94IiBpZD0iYSI+PGZlT2Zmc2V0IGR4PSIyIiBpbj0iU291cmNlQWxwaGEiIHJlc3VsdD0ic2hhZG93T2Zmc2V0T3V0ZXIxIi8+PGZlR2F1c3NpYW5CbHVyIHN0ZERldmlhdGlvbj0iNy41IiBpbj0ic2hhZG93T2Zmc2V0T3V0ZXIxIiByZXN1bHQ9InNoYWRvd0JsdXJPdXRlcjEiLz48ZmVDb2xvck1hdHJpeCB2YWx1ZXM9IjAgMCAwIDAgMCAwIDAgMCAwIDAgMCAwIDAgMCAwIDAgMCAwIDAuMTUgMCIgaW49InNoYWRvd0JsdXJPdXRlcjEiLz48L2ZpbHRlcj48cGF0aCBkPSJNMiAwaDMyNy41ODZhMSAxIDAgMDEuNzA3LjI5M2w0OS40MTQgNDkuNDE0YTEgMSAwIDAxLjI5My43MDdWMjY4YTIgMiAwIDAxLTIgMkgyYTIgMiAwIDAxLTItMlYyYTIgMiAwIDAxMi0yeiIgaWQ9ImIiLz48L2RlZnM+PGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTMzMCkiIGZpbGw9Im5vbmUiIGZpbGwtcnVsZT0iZXZlbm9kZCI+PHJlY3QgZmlsbD0iI0ZGRiIgd2lkdGg9IjM4MCIgaGVpZ2h0PSIyNzAiIHJ4PSIyIi8+PHBhdGggZD0iTTM2Ni44MzMgMTAuODMzdjIuMzM0aDIuMzM0di0yLjMzNGgtMi4zMzR6bS03IDE2LjMzNGgyLjMzNFYyOS41aC0yLjMzNHYtMi4zMzN6bTkuMzM0LTkuMzM0aDIuMzMzdjIuMzM0aC0yLjMzM3YtMi4zMzR6bS05LjMzNCA0LjY2N2gyLjMzNHYyLjMzM2gtMi4zMzRWMjIuNXptNC42NjctNC42NjdoMi4zMzN2Mi4zMzRIMzY0LjV2LTIuMzM0ek0zNTEuNjY3IDguNUgzNjFjLjY0NCAwIDEuMTY3LjUyMiAxLjE2NyAxLjE2N1YxOWMwIC42NDQtLjUyMyAxLjE2Ny0xLjE2NyAxLjE2N2gtOS4zMzNBMS4xNjcgMS4xNjcgMCAwMTM1MC41IDE5VjkuNjY3YzAtLjY0NS41MjItMS4xNjcgMS4xNjctMS4xNjd6bTEuMTY2IDIuMzMzdjdoN3YtN2gtN3ptMi4zMzQgMi4zMzRoMi4zMzNWMTUuNWgtMi4zMzN2LTIuMzMzem0xMC41LTQuNjY3aDQuNjY2Yy42NDUgMCAxLjE2Ny41MjIgMS4xNjcgMS4xNjd2NC42NjZjMCAuNjQ1LS41MjIgMS4xNjctMS4xNjcgMS4xNjdoLTQuNjY2YTEuMTY3IDEuMTY3IDAgMDEtMS4xNjctMS4xNjdWOS42NjdjMC0uNjQ1LjUyMi0xLjE2NyAxLjE2Ny0xLjE2N3ptMCAxNGg0LjY2NmMuNjQ1IDAgMS4xNjcuNTIyIDEuMTY3IDEuMTY3djQuNjY2YzAgLjY0NS0uNTIyIDEuMTY3LTEuMTY3IDEuMTY3aC00LjY2NmExLjE2NyAxLjE2NyAwIDAxLTEuMTY3LTEuMTY3di00LjY2NmMwLS42NDUuNTIyLTEuMTY3IDEuMTY3LTEuMTY3em0xLjE2NiAyLjMzM3YyLjMzNGgyLjMzNHYtMi4zMzRoLTIuMzM0ek0zNTEuNjY3IDIyLjVoNC42NjZjLjY0NSAwIDEuMTY3LjUyMiAxLjE2NyAxLjE2N3Y0LjY2NmMwIC42NDUtLjUyMiAxLjE2Ny0xLjE2NyAxLjE2N2gtNC42NjZhMS4xNjcgMS4xNjcgMCAwMS0xLjE2Ny0xLjE2N3YtNC42NjZjMC0uNjQ1LjUyMi0xLjE2NyAxLjE2Ny0xLjE2N3ptMS4xNjYgMi4zMzN2Mi4zMzRoMi4zMzR2LTIuMzM0aC0yLjMzNHoiIG9wYWNpdHk9Ii41IiBmaWxsPSIjMDAwIiBmaWxsLW9wYWNpdHk9Ii45Ii8+PHVzZSBmaWxsPSIjMDAwIiBmaWx0ZXI9InVybCgjYSkiIHhsaW5rOmhyZWY9IiNiIi8+PHVzZSBmaWxsPSIjRkZGIiB4bGluazpocmVmPSIjYiIvPjwvZz48L3N2Zz4=",
   mode: "qrcode",
   qrcode: "",
   thumb: "",
   bgimage: "",
   height: 0
 })
 

 function quit(){
    contextStore.quit()
    router.push("/")
 }
 const dlgwidth = ref(480)
 
 const  refForm = ref()
 /** 登录按钮 Loading */
 const loading = ref(false)
 /** 登录表单数据 */
 const outerVisible = ref(false)
 const loginForm= reactive({
   username: "",
   password: "",
   code: "",
   uuid: "",
   mobile:"",
   type:"username",
 })
 watch(
   () => outerVisible.value,
   (visible) => {
     if (visible) {
       handlerefresh()
     }
   }
 )
 const trylogin = () => {
     refForm.value.validate(r=>{
         console.log('r',r)
         if(r){
             outerVisible.value = true      
         }
     })
 }


 
 /** 登录逻辑 */
 const handleLogin = () => {
   loading.value = true
   if (/^1[\d]{1}[0-9]{9}$/.test(loginForm.username)) {
     loginForm.mobile = loginForm.username
     loginForm.type = "mobile"
   }
   console.log('dlgwidth',loginForm)
   contextStore.login(loginForm)
     .then(() => {
      loginshow.value = false  
      outerVisible.value = false   
     })
     .catch((err) => {
       ElMessage.error(err.message)
       createCode()
       //loginForm.password = ""
     })
     .finally(() => {
       loading.value = false
     })
 }
 
 const handleclose = () => {
   outerVisible.value = false
 }
 const handleconfirm = (params) => {
   //console.log("params", params)
   const dotsA = [] 
   params.forEach(function (value) {
     const tmp = [value.x, value.y]
     dotsA.push(tmp.join(","))
   })
   if (dotsA.length < 1) {
     ElMessage.error("请输入验证码")
     return
   }
   loginForm.code = dotsA.join(",")
   handleLogin()
 }
 const handlerefresh = () => {
   createCode()
 }
 /** 创建验证码 */
 const createCode = () => {
   // 先清空验证码的输入
   loginForm.code = ""
   // 获取验证码
   loginForm.uuid = ""
   getLoginCodeApi().then((res) => {
     appdata.bgimage = res.data.image_base64
     appdata.thumb = res.data.thumb_base64
     loginForm.uuid = res.data.captcha_key
   })
 }
 /** 解析 redirect 字符串 为 path 和  queryParams */
 function parseRedirect(){
   const query = route.query;
   const redirect = (query.redirect ) ?? "/learn";
   const url = new URL(redirect, window.location.origin);
   const path = url.pathname;
   const queryParams = {};
 
   url.searchParams.forEach((value, key) => {
     queryParams[key] = value;
   });
 
   return { path, queryParams };
 }
 
 
     </script>
 
