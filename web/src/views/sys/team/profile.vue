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
          <el-form ref="formRef" :model="formData" label-width="100px" label-position="left">
            <el-form-item label="Logo(100*100)">
              <PickImage
                v-model="formData.pic"
                @success="updateprofile('pic')"
                style="width: 80px; height: 80px"
                placeholder=""
              />
            </el-form-item>
            <el-form-item label="名称">
              <div class="form-item" v-if="!state.name">
                <label>{{ formData.name }}</label>
                <el-button link @click="() => (state.name = true)">编辑</el-button>
              </div>
              <div class="form-item" v-if="state.name">
                <el-input v-model="formData.name" type="text" placeholder="设置名称" />
                <div class="btnopt">
                  <el-button type="primary" link @click="updateprofile('name')">保存</el-button>
                  <el-button type="default" link @click="() => (state.name = false)">取消</el-button>
                </div>
              </div>
            </el-form-item>
            <el-form-item label="愿景">
              <div class="form-item" v-if="!state.memo">
                <label>{{ formData.memo }}</label>
                <el-button link @click="() => (state.memo = true)">编辑</el-button>
              </div>
              <div class="form-item" v-if="state.memo">
                <el-input v-model="formData.memo" type="textarea" placeholder="愿景" />
                <div class="btnopt">
                  <el-button type="primary" link @click="updateprofile('memo')">保存</el-button>
                  <el-button type="default" link @click="() => (state.memo = false)">取消</el-button>
                </div>
              </div>
            </el-form-item>
            <!-- <el-form-item label="联系人">
              <div class="form-item" v-if="!state.linker">
                <label>{{ formData.linker }}</label>
                <el-button link @click="() => (state.linker = true)">编辑</el-button>
              </div>
              <div class="form-item" v-if="state.linker">
                <el-input v-model="formData.linker" type="text" placeholder="联系人" />
                <div class="btnopt">
                  <el-button type="primary" link @click="updateprofile('linker')">保存</el-button>
                  <el-button type="default" link @click="() => (state.linker = false)">取消</el-button>
                </div>
              </div>
            </el-form-item>
            <el-form-item label="联系电话">
              <div class="form-item" v-if="!state.tel">
                <label>{{ formData.tel }}</label>
                <el-button link @click="() => (state.tel = true)">编辑</el-button>
              </div>
              <div class="form-item" v-if="state.tel">
                <el-input v-model="formData.tel" type="text" placeholder="联系电话" />
                <div class="btnopt">
                  <el-button type="primary" link @click="updateprofile('tel')">保存</el-button>
                  <el-button type="default" link @click="() => (state.tel = false)">取消</el-button>
                </div>
              </div>
            </el-form-item> -->
          </el-form>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="never" class="search-wrapper">
          <template #header>
            <div class="row">
              <div class="card-header">
                <span>详情</span>
              </div>
              <div>
                <div class="form-item" v-if="!state.content">
                  <el-button link @click="() => (state.content = true)">编辑</el-button>
                </div>
                <div class="form-item" v-if="state.content">
                  <div>
                    <el-button link @click="updateprofile('content')">保存</el-button>
                    <el-button link @click="() => (state.content = false)">取消</el-button>
                  </div>
                </div>
              </div>
            </div>
          </template>
          <div class="content-wraper" v-html="formData.content" />
        </el-card>
      </el-col>
    </el-row>
    <el-dialog :close-on-click-modal="false" title="详情" v-model="state.content">
      <Tinymce v-if="state.content" v-model="formData.content" placeholder="输入" />
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue"
import { type FormInstance, ElMessage } from "element-plus"
import teamApi,{ ITeam } from "@/api/sys/team"
import Tinymce from "@/components/Tinymce/index.vue"
import { useUserStore } from "@/store/modules/user"

const formRef = ref<FormInstance | null>(null)
const formData = ref<ITeam>({
  id: undefined,
  name: "",
  creator: undefined,
  trade:"",
  pic:"",
  memo:"",
  content:"",

})
const state = ref({
  name: false,
  memo: false,
  content: false
})
const updateprofile = (field:string) => {
  if (field) {
    let tmp = {}
    //tmp[`${field}`] = false
    //Object.assign(state.value,tmp)
  }
  teamApi.update(formData.value)
    .then((res:any) => {
      ElMessage.success(res.msg || "更新成功")
    })
    .catch((err:Error) => {
      ElMessage.error(err.message)
    })
}
const initdata = () => {
  const orgId = useUserStore().teamId
  teamApi.getOne(orgId as number).then((res:any) => {
    formData.value = res.data
  })
}
initdata()
</script>

<style lang="scss">
.row {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}
.btnopt {
  width: 120px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
.form-item {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  label {
    width: 210px;
  }
  .data-label {
    width: 220px;
    text-align: right;
  }
}
.right {
  justify-content: flex-end;
}
.el-card__body {
  .content-wraper {
    p {
      img {
        width: 100%;
      }
    }
  }
}
</style>
