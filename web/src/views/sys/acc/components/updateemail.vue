<template>
  <el-form ref="searchFormRef" :model="formData" label-position="top">
    <el-form-item label="新的邮箱">
      <el-input v-model="formData.email" type="number" placeholder="输入新的邮箱好吗" />
    </el-form-item>
    <div style="display: flex; justify-content: flex-end">
      <el-button type="primary" @click="handleSubmit">提交</el-button>
    </div>
  </el-form>
</template>
<script setup lang="ts">
import { reactive } from "vue"
import accountApi from "@/api/sys/acc"
import { ElMessage } from "element-plus"
const handleSubmit = () => {
  accountApi
    .resetemail(formData.email)
    .then((res) => {
      ElMessage.success(res.msg)
    })
    .catch((err) => {
      ElMessage.error(err.message)
    })
}
const formData = reactive({
  email: "",
  code: "",
  uuid: ""
})

const onerror = (m:string) => {
  ElMessage.error(m)
}
</script>
@/api/sys/sys/acc