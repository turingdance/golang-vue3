<template>
  <div>
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
          <div class="fm-field" style="margin-top: 5px; margin-bottom: 5px">
            <div
              :class="{ 'btn-send': true, disable: appdata.smscode.time > 0, enable: appdata.smscode.time }"
            >
              {{ appdata.smscode.time == 0 ? "" : appdata.smscode.time }}{{ appdata.smscode.text }}
            </div>
          </div>
        </template>
      </el-popover>
  </div>
</template>

<script lang="ts" setup module="smstask">
import smsApi from "@/api/sys/sms"
import { ElMessage } from "element-plus"
import captchaApi from "@/api/sys/captcha"
import Captcha from "@/components/Captcha/index.vue"
import { reactive } from "vue"
const appprops = defineProps({
  mobile: {
    type: String,
    require: true
  },
  duration: {
    type: Number,
    default: 60
  }
})
const appdata = reactive({
  thumb:"",
  bgimage:"",
  smscode: {
    text: "获取验证码",
    time: 0,
    enable: true
  }
})
const emits = defineEmits(["error", "data"])
var postdata = {mobile:"",captchaKey:"",captchaCode:""}
const outerVisible = ref(false)

const handleclose = () => {
  outerVisible.value = false
}
const handleconfirm = (params: any) => {
  //console.log("params", params)
  const dotsA = [] as any[]
  params.forEach(function (value: any) {
    const tmp = [value.x, value.y]
    dotsA.push(tmp.join(","))
  })
  if (dotsA.length < 1) {
    ElMessage.error("请输入验证码")
    return
  }
  postdata.captchaCode = dotsA.join(",")
  sendSms()
}
const handlerefresh = () => {
  createCode()
}
/** 创建验证码 */
const createCode = () => {
  // 先清空验证码的输入
  postdata.captchaCode = ""
  // 获取验证码
  postdata.captchaKey = ""
  captchaApi.getCaptcha().then((res: any) => {
    appdata.bgimage = res.data.image_base64
    appdata.thumb = res.data.thumb_base64
    postdata.captchaKey = res.data.captcha_key
  })
}
const sendSms = () => {
  if (!appdata.smscode.enable) {
    emits("error", "请稍后再试")
    return
  }
  if (!appprops.mobile) {
    emits("error", "请输入手机号")
    return
  }
  if (!/^1[3456789]{1}\d{9}$/.test(appprops.mobile as string)) {
    emits("error", "手机号不正确")
    return
  }
  postdata.mobile = appprops.mobile
  smsApi.send(postdata)
    .then((res) => {
      emits("data", res)
      appdata.smscode.enable = false
      setInterValFunc()
    })
    .catch((err) => {
      emits("error", err.message || err.msg || "出错了")
    })
}
const setInterValFunc = () => {
  appdata.smscode.time = appprops.duration
  appdata.smscode.text = "秒"
  const setTime = setInterval(() => {
    if (appdata.smscode.time - 1 == 0) {
      appdata.smscode.time = 0
      appdata.smscode.text = "重新获取"
      appdata.smscode.enable = true
      clearInterval(setTime)
    } else {
      appdata.smscode.time--
    }
  }, 1000)
}
</script>

<style lang="css" scoped>
.btn-send {
  height: 100rpx;
  line-height: 100rpx;
  float: right;
  font-size: 26rpx;
  text-align: center;
  font-family: PingFangSC-Regular, PingFang SC;
  font-weight: 400;
  cursor: pointer;
}
.disable {
  color: calc(--v3-color-disable);
}
.enable {
  color: calc(--v3-color-primary);
}
</style>
