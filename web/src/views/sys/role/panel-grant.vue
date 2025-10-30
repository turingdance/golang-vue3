<template>
  <el-tree
    ref="treeRef"
    style="max-width: 600px"
    :data="treedata"
    show-checkbox
    node-key="value"
    :check-strictly="notlinkchanged"
    highlight-current
    :props="defaultProps"
  >
  <template #default="{ node, data }">
        <div class="custom-tree-node">
          <div>
            <svg-icon :size="16" :icon-class="data.rights.type"></svg-icon>
            <span class="pl-1">{{ node.label }}</span>
        </div>
          <div>
            <el-button type="primary" v-if="!!data.children && data.children.length>0"  link @click="checkall(data)">
              全选
            </el-button>
            <el-button
              style="margin-left: 4px"
              v-if="!!data.children && data.children.length>0"
              link
              @click="uncheckall(data)"
            >
              取消
            </el-button>
          </div>
        </div>
      </template>
</el-tree>
  <slot name="footer" :checked="checked">
    <el-row class="m-5 mt-10 flex-x-end">
      <el-button type="primary" @click="confirm">确认</el-button>
      </el-row>
  </slot>
</template>
<style scoped lang="scss">
.custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 14px;
  padding-right: 8px;
}
</style>
<script setup type="ts">
import resApi from "@/api/sys/rights"
import { onMounted,watch } from "vue";
const defaultProps = {
  children: 'children',
  label: 'label',
}
const model = defineModel({required: true})
const treeRef = ref()

const setCheckedKeys = (ids)=>{
    treeRef.value && treeRef.value.setCheckedKeys(ids,false)
}
const notlinkchanged  = ref(true)
const checkall = (data)=>{
  notlinkchanged.value = false
  nextTick(()=>{
    treeRef.value && treeRef.value.setChecked(data.value,true,true)
    notlinkchanged.value = true
  })
  
}

const uncheckall = (data)=>{
  notlinkchanged.value = false
  nextTick(()=>{
    treeRef.value && treeRef.value.setChecked(data.value,false,true)
    notlinkchanged.value = true
  })
}

// 父子不相互关联
// 这个参数非常重要
// 初始化阶段需要设置为true,
// 设置阶段需要设置为false
// 在显示复选框的情况下，
// 是否严格的遵循父子不互相关联的做法，默认为 false
const emits = defineEmits(['confirm'])
const treedata = ref([])
const initdata =async  ()=>{
  return resApi.tree().then(({data})=>{
    return data
  })
}

const confirm = ()=>{
  var  fullchecked = treeRef.value.getCheckedKeys()
  var  halfchecked = treeRef.value.getHalfCheckedKeys()
  let tmp = fullchecked.concat(halfchecked)
  emits('confirm',tmp)
}
onMounted(async ()=>{
  treedata.value = await initdata()
  setCheckedKeys(model.value)
})
</script>
