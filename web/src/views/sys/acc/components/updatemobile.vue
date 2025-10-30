<template>
  <el-form ref="searchFormRef" :model="formData" label-position="top">
    <el-form-item label="手机号">
      <el-input v-model="formData.mobile" type="number" placeholder="输入新的手机号" />
    </el-form-item>

    <el-form-item label="验证码">
      <el-input v-model="formData.code" type="number" placeholder="输入验证码">
        <template #append>
          <Smstask :mobile="formData.mobile" @error="onerror($event)" />
        </template>
      </el-input>
    </el-form-item>
    <div style="display: flex; justify-content: flex-end">
      <el-button type="primary" @click="handleSubmit">提交</el-button>
    </div>
  </el-form>
</template>
<script setup lang="ts">
import { reactive } from "vue"

import accountApi from "@/api/sys/acc"
import Smstask from "@/components/Smstask/index.vue"
import { ElMessage } from "element-plus"

const handleSubmit = () => {
  accountApi.resetmobile(formData).then((res:any) => {
      ElMessage.success(res.msg)
    })
    .catch((err:Error) => {
      ElMessage.error(err.message)
    })
}
const formData = reactive({
  mobile: "",
  code: "",
  uuid: ""
})

const onerror = (m:string) => {
  ElMessage.error(m)
}
</script>
@/api/sys/sys/acc