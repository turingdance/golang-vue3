<template>
  <el-form ref="searchFormRef" :model="formData" label-position="top">
    <el-form-item label="原密码">
      <el-input v-model="formData.passwordOld" required type="password" placeholder="输入原始密码" />
    </el-form-item>

    <el-form-item label="新密码">
      <el-input v-model="formData.password" required type="password" placeholder="请输入新密码" />
    </el-form-item>

    <el-form-item label="新密码">
      <el-input v-model="formData.passwordr" required type="password" placeholder="请在此输入新密码" />
    </el-form-item>

    <div style="display: flex; justify-content: flex-end">
      <el-button type="primary" @click="handleSubmit">提交</el-button>
    </div>
  </el-form>
</template>
<script setup lang="ts">
import { reactive } from "vue"
import { type ILoginData } from "@/api/sys/login"
import accApi from "@/api/sys/acc"
import { ElMessage } from "element-plus"
const handleSubmit = () => {
  if (formData.password != formData.passwordr) {
    ElMessage.error("密码不一致")
    return
  }
  if (!formData.password) {
    ElMessage.error("请输入密码")
    return
  }
  if (formData.password.length < 6) {
    ElMessage.error("密码太简单了")
    return
  }
  accApi.updatePwd(formData)
}
const formData = reactive({
  userId:"",
  passwordOld: "",
  password: "",
  passwordr: "",
  type: "userId"
})
</script>
@/api/sys/sys/acc