<template>
  <div class="app-container">
    <el-row :gutter="12">
      <el-col :span="12">
        <el-card shadow="never" class="search-wrapper">
          <template #header>
            <div class="card-header">
              <span>基本信息</span>
            </div>
          </template>
          <el-form label-width="100px" label-position="left">
            <el-form-item label="头像">
              <PickImage
                v-model="formData.pic"
                @success="updateprofile('pic')"
                style="width: 60px; height: 60px"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="昵称">
              <div class="form-item" v-if="!state.nickname">
                <label>{{ userStore.user.nickname }}</label>
                <el-button link @click="() => (state.nickname = true)">编辑</el-button>
              </div>
              <div class="form-item" v-if="state.nickname">
                <el-input v-model="formData.nickname" style="width: 120px" type="text" placeholder="设置昵称" />
                <div>
                  <el-button type="primary" @click="updateprofile('nickname')">保存</el-button>
                  <el-button type="default" @click="() => (state.nickname = false)">取消</el-button>
                </div>
              </div>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="never" class="search-wrapper">
          <template #header>
            <div class="card-header">
              <span>账户安全</span>
            </div>
          </template>
          <el-form ref="formRef" :model="formData" label-width="100px" label-position="left">
            <el-form-item label="手机号">
              <div class="form-item">
                <label class="data-label" v-desensitize-phone="userStore.user.mobile"></label>
                <el-button link @click="() => (state.mobile = true)">更换</el-button>
              </div>
            </el-form-item>

            <el-form-item label="用户名">
              <div class="form-item">
                <label class="data-label"  v-desensitize="{data:userStore.user.username,start:1,end:1,mask:'*'}"></label>
                <el-button link @click="() => (state.name = true)">更换</el-button>
              </div>
            </el-form-item>

            <el-form-item label="邮箱">
              <div class="form-item">
                <label class="data-label" v-desensitize-email="userStore.user.email">{{ userStore.user.email }}</label>
                <el-button link @click="() => (state.email = true)">更换</el-button>
              </div>
            </el-form-item>

            <el-form-item label="密码">
              <div class="form-item">
                <label class="data-label" />
                <el-button link @click="() => (state.password = true)">更换</el-button>
              </div>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>
    </el-row>

    <el-dialog :close-on-click-modal="false" title="修改登录手机" v-model="state.mobile">
      <updatemobile />
    </el-dialog>

    

    <el-dialog :close-on-click-modal="false" title="修改登录账号" v-model="state.name">
      <updateusername />
    </el-dialog>

    <el-dialog :close-on-click-modal="false" title="修改登录邮箱" v-model="state.email">
      <updateemail />
    </el-dialog>

    <el-dialog :close-on-click-modal="false" title="修改登录密码" v-model="state.password">
      <updatepwd />
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { type FormInstance } from "element-plus"
import accApi,{ IAcc } from "@/api/sys/acc"
import updatemobile from "./components/updatemobile.vue"
import updateemail from "./components/updateemail.vue"
import updateusername from "./components/updateusername.vue"
import updatepwd from "./components/updatepwd.vue"
import { useUserStore } from "@/store/modules/user"
const formRef = ref<FormInstance | null>(null)
const formData = ref<IAcc>({
  id:undefined,
  pic: "",
  nickname: "胡大力",
  createAt: undefined,
  username: "",
  mobile:"",
  email:"",
})
const userStore = useUserStore()
formData.value.pic = userStore.user.pic
const state = ref({
  nickname: false,
  mobile: false,
  password: false,
  name: false,
  email:false,
})
const updateprofile = (field:string)=>{
    var postdata = {}
    postdata[`${field}`] = formData.value[`${field}`]
    accApi.update(postdata)
}
</script>

<style lang="scss" scoped>
.form-item {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  .data-label {
    width: 220px;
    text-align: right;
    padding-right: 5px;
  }
}
.right {
  justify-content: flex-end;
}
</style>
@/api/sys/sys/acc