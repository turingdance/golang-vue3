<template>
  <el-form ref="searchFormRef" :model="formData" label-position="top">
    <el-form-item label="新的用户名">
      <el-input v-model="formData.username" type="text" placeholder="输入新的账号" />
    </el-form-item>

    <!-- <el-form-item label="验证码">
        <el-input v-model="formData.code" type="number" placeholder="输入验证码">
                    <template #append>
                                <Smstask :mobile="formData.mobile"   @error="onerror($event)"></Smstask>
                            
            </template>
        </el-input>
    </el-form-item> -->
    <div style="display: flex; justify-content: flex-end">
      <el-button type="primary" @click="handleSubmit">提交</el-button>
    </div>
  </el-form>
</template>
<script setup lang="ts">
import { reactive } from "vue"

import accountApi from "@/api/sys/acc"


const handleSubmit = () => {
  accountApi.updateUsername(formData.username)
}
const formData = reactive({
  username: "",
  code: "",
  uuid: ""
})
</script>
@/api/sys/sys/acc