<template>
  <div class="login-card">
    <div class="box-card content">
      <div class="fm-field"><h3>注册</h3></div>
      <div class="fm-field row">
        <input
          v-model.trim="loginForm.username"
          placeholder="请输入手机号"
          readonly
          type="text"
          tabindex="1"
          v-pure
          size="large"
          class="line-input"
        />
      </div>
      <div class="fm-field row">
        <input
          v-model.trim="loginForm.smscode"
          placeholder="请输入验证码"
          style="width: 50%"
          type="text"
          tabindex="1"
          v-pure
          size="large"
          class="line-input"
        />
        <Smstask :mobile="loginForm.username" @error="onerror($event)" />
      </div>
      <div class="fm-field">
        <input
          v-model.trim="loginForm.password"
          placeholder="请设置密码"
          readonly
          auto-complete="none"
          type="password"
          tabindex="2"
          size="large"
          v-pure
          class="line-input"
        />
      </div>

      <div class="fm-field">
        <input
          v-model.trim="loginForm.orgname"
          placeholder="团队/组织机构名称"
          type="text"
          tabindex="3"
          size="large"
          v-pure
          class="line-input"
        />
      </div>
      <div class="fm-field">
        <input
          v-model.trim="loginForm.linker"
          placeholder="团队联系人信息"
          type="text"
          tabindex="3"
          size="large"
          v-pure
          class="line-input"
        />
      </div>

      <div class="fm-field" style="display: none">
        <input
          v-model.trim="loginForm.tel"
          placeholder="联系方式,手机号/座机"
          type="text"
          tabindex="3"
          size="large"
          v-pure
          class="line-input"
        />
      </div>

      <el-popover
        :width="400"
        placement="left"
        v-model:visible="outerVisible"
        trigger="click"
        popper-style="box-shadow: rgb(14 18 22 / 35%) 0px 10px 38px -10px, rgb(14 18 22 / 20%) 0px 10px 20px -15px; padding: 20px;"
      >
        <Captcha
          @refresh="handlerefresh"
          @close="handleclose"
          @confirm="handleconfirm"
          :thumb-base64="appdata.thumb"
          :image-base64="appdata.bgimage"
        />
        <template #reference>
          <div class="fm-field" style="margin-top: 30px; margin-bottom: 20px">
            <el-button :loading="loading" style="width: 100%" type="primary" size="large"> 注 册 </el-button>
          </div>
        </template>
      </el-popover>
      <div class="btn-wraper fm-field">
        <el-link type="primary">点击这里找回密码</el-link>
        <RouterLink to="/login">
          <el-link type="primary">这里用登录账号</el-link>
        </RouterLink>
      </div>
      <div class="fm-field tip">
        <p>注册视为您已同意第三方账号绑定协议、服务条款、隐私政策</p>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { onMounted, onUnmounted, reactive, ref } from "vue"
import { useRouter } from "vue-router"
import Smstask from "@/components/Smstask/index.vue"
import { useUserStore } from "@/store/modules/user"
import { type FormInstance, FormRules, ElMessageBox, ElMessage } from "element-plus"
import { type ILoginData, getCaptcha as getLoginCodeApi } from "@/api/sys/login"
import Captcha from "@/components/Captcha/index.vue"
import websocketUtil from "@/utils/websocket"
import eventbus from "@/utils/eventbus"
const appdata = reactive({
  wxloginIcon:
    "data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNTAiIGhlaWdodD0iNTAiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgeG1sbnM6eGxpbms9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkveGxpbmsiPjxkZWZzPjxmaWx0ZXIgeD0iLTUuNyUiIHk9Ii04LjclIiB3aWR0aD0iMTEyLjQlIiBoZWlnaHQ9IjExNy40JSIgZmlsdGVyVW5pdHM9Im9iamVjdEJvdW5kaW5nQm94IiBpZD0iYSI+PGZlT2Zmc2V0IGR4PSIyIiBpbj0iU291cmNlQWxwaGEiIHJlc3VsdD0ic2hhZG93T2Zmc2V0T3V0ZXIxIi8+PGZlR2F1c3NpYW5CbHVyIHN0ZERldmlhdGlvbj0iNy41IiBpbj0ic2hhZG93T2Zmc2V0T3V0ZXIxIiByZXN1bHQ9InNoYWRvd0JsdXJPdXRlcjEiLz48ZmVDb2xvck1hdHJpeCB2YWx1ZXM9IjAgMCAwIDAgMCAwIDAgMCAwIDAgMCAwIDAgMCAwIDAgMCAwIDAuMTUgMCIgaW49InNoYWRvd0JsdXJPdXRlcjEiLz48L2ZpbHRlcj48cGF0aCBkPSJNMiAwaDMyNy41ODZhMSAxIDAgMDEuNzA3LjI5M2w0OS40MTQgNDkuNDE0YTEgMSAwIDAxLjI5My43MDdWMjY4YTIgMiAwIDAxLTIgMkgyYTIgMiAwIDAxLTItMlYyYTIgMiAwIDAxMi0yeiIgaWQ9ImIiLz48L2RlZnM+PGcgdHJhbnNmb3JtPSJ0cmFuc2xhdGUoLTMzMCkiIGZpbGw9Im5vbmUiIGZpbGwtcnVsZT0iZXZlbm9kZCI+PHJlY3QgZmlsbD0iI0ZGRiIgd2lkdGg9IjM4MCIgaGVpZ2h0PSIyNzAiIHJ4PSIyIi8+PHBhdGggZD0iTTM2Ni44MzMgMTAuODMzdjIuMzM0aDIuMzM0di0yLjMzNGgtMi4zMzR6bS03IDE2LjMzNGgyLjMzNFYyOS41aC0yLjMzNHYtMi4zMzN6bTkuMzM0LTkuMzM0aDIuMzMzdjIuMzM0aC0yLjMzM3YtMi4zMzR6bS05LjMzNCA0LjY2N2gyLjMzNHYyLjMzM2gtMi4zMzRWMjIuNXptNC42NjctNC42NjdoMi4zMzN2Mi4zMzRIMzY0LjV2LTIuMzM0ek0zNTEuNjY3IDguNUgzNjFjLjY0NCAwIDEuMTY3LjUyMiAxLjE2NyAxLjE2N1YxOWMwIC42NDQtLjUyMyAxLjE2Ny0xLjE2NyAxLjE2N2gtOS4zMzNBMS4xNjcgMS4xNjcgMCAwMTM1MC41IDE5VjkuNjY3YzAtLjY0NS41MjItMS4xNjcgMS4xNjctMS4xNjd6bTEuMTY2IDIuMzMzdjdoN3YtN2gtN3ptMi4zMzQgMi4zMzRoMi4zMzNWMTUuNWgtMi4zMzN2LTIuMzMzem0xMC41LTQuNjY3aDQuNjY2Yy42NDUgMCAxLjE2Ny41MjIgMS4xNjcgMS4xNjd2NC42NjZjMCAuNjQ1LS41MjIgMS4xNjctMS4xNjcgMS4xNjdoLTQuNjY2YTEuMTY3IDEuMTY3IDAgMDEtMS4xNjctMS4xNjdWOS42NjdjMC0uNjQ1LjUyMi0xLjE2NyAxLjE2Ny0xLjE2N3ptMCAxNGg0LjY2NmMuNjQ1IDAgMS4xNjcuNTIyIDEuMTY3IDEuMTY3djQuNjY2YzAgLjY0NS0uNTIyIDEuMTY3LTEuMTY3IDEuMTY3aC00LjY2NmExLjE2NyAxLjE2NyAwIDAxLTEuMTY3LTEuMTY3di00LjY2NmMwLS42NDUuNTIyLTEuMTY3IDEuMTY3LTEuMTY3em0xLjE2NiAyLjMzM3YyLjMzNGgyLjMzNHYtMi4zMzRoLTIuMzM0ek0zNTEuNjY3IDIyLjVoNC42NjZjLjY0NSAwIDEuMTY3LjUyMiAxLjE2NyAxLjE2N3Y0LjY2NmMwIC42NDUtLjUyMiAxLjE2Ny0xLjE2NyAxLjE2N2gtNC42NjZhMS4xNjcgMS4xNjcgMCAwMS0xLjE2Ny0xLjE2N3YtNC42NjZjMC0uNjQ1LjUyMi0xLjE2NyAxLjE2Ny0xLjE2N3ptMS4xNjYgMi4zMzN2Mi4zMzRoMi4zMzR2LTIuMzM0aC0yLjMzNHoiIG9wYWNpdHk9Ii41IiBmaWxsPSIjMDAwIiBmaWxsLW9wYWNpdHk9Ii45Ii8+PHVzZSBmaWxsPSIjMDAwIiBmaWx0ZXI9InVybCgjYSkiIHhsaW5rOmhyZWY9IiNiIi8+PHVzZSBmaWxsPSIjRkZGIiB4bGluazpocmVmPSIjYiIvPjwvZz48L3N2Zz4=",
  mode: "qrcode",
  qrcode: "",
  thumb: "",
  bgimage: "",
  height: 0
})
const router = useRouter()
const loginFormRef = ref<FormInstance | null>(null)

/** 登录按钮 Loading */
const loading = ref(false)
/** 登录表单数据 */
const outerVisible = ref(false)
const loginForm: ILoginData = reactive({
  username: "",
  password: "",
  orgname: "",
  code: "",
  uuid: "",
  linker: "",
  tel: "",
  smscode: "",
  type: "username"
})
const trylogin = () => {
  validate().then((valid: boolean) => {
    if (valid) {
      outerVisible.value = true
    }
  })
}
const validate = () => {
  if (!loginForm.username) {
    return Promise.reject(new Error("请输入用户名"))
  }
  if (/^1[\d]{10}$/.test(loginForm.username))
    if (!loginForm.username) {
      return Promise.reject(new Error("请输入正确的手机号"))
    }
  loginForm.tel = loginForm.username
  if (!loginForm.password) {
    return Promise.reject(new Error("请输入密码"))
  }
  if (!loginForm.orgname) {
    return Promise.reject(new Error("请输入团队名称"))
  }
  if (!loginForm.linker) {
    return Promise.reject(new Error("请输入联系人"))
  }
  return Promise.resolve(true)
}
/** 登录逻辑 */
const handleLogin = () => {
  loading.value = true
  useUserStore()
    .login(loginForm)
    .then(() => {
      router.push({ path: "/" })
    })
    .catch(() => {
      createCode()
      //loginForm.password = ""
    })
    .finally(() => {
      loading.value = false
    })
}

const onerror = (m) => {
  ElMessage.error(m)
}
const handleclose = () => {
  outerVisible.value = false
}
const handleconfirm = (params: any) => {
  ElMessage.error("当前暂不开放注册")
  return
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
  getLoginCodeApi().then((res: any) => {
    appdata.bgimage = res.data.image_base64
    appdata.thumb = res.data.thumb_base64
    loginForm.uuid = res.data.captcha_key
  })
}

interface IData {
  scene: string
  token: string
}
interface IMessage {
  code: number
  data: IData
}
</script>

<style lang="scss" scoped>
.login-card {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  background-color: #ffffff;
  padding: 5px 5px 5px 5px;

  .content {
    display: flex;
    flex-direction: column;
    margin: 5px 5px 5px 5px;

    .fm-field {
      margin: 5px 10px 5px 10px;
      flex: 1;
      height: 35px;
      &.row {
        display: flex;
        flex-direction: row;
        justify-content: flex-start;
        align-items: center;
      }
      input {
        flex: 1;
        width: 100%;
      }
    }
    .qrcode {
      width: 120px;
      height: 120px;
    }
  }
}
.line-input {
  border: none;
  font-size: 14px;
  line-height: 22px;
  height: 35px;
  margin-top: 0;
  padding-left: 0;
  background: #fff;
  box-shadow: none;
  border: none;
  border-bottom: 1px solid #d7d8d9;
}
.btn-wraper {
  display: flex;
  justify-content: space-around;
  align-items: center;
  flex-direction: row;
}
.tip {
  font-size: 12px;
  color: #b3b3b3;
}
</style>
