<template>
  <template v-if="!!dictItem">
    <slot :item="dictItem">
      <font v-if="appprops.tag=='font'">{{ dictItem.label }}</font>
      <el-tag v-else-if="appprops.tag=='tag'">{{ dictItem.label }}</el-tag>
      <template v-else>{{ dictItem.label }}</template>
    </slot>
  </template>
</template>
<script lang="ts" setup module="smstask">
import { ref } from "vue"
import { useDictStore } from "@/store/modules/dict"
import { IDictItem } from "@/api/sys/dict"
const appprops = defineProps({
  dictname: {
    type: String,
    require: true
  },
  value: {
    type: [String,Number],
    requiew: true
  },
  tag:{
    type:String,
    default:'tag'
  },
  label:{
    type:String,
    default:"",
  }
})
const dictStore = useDictStore()
const dictItem = ref<IDictItem>({
  label:appprops.label,value:"",
})
onMounted(()=>{
  refreshdict()
})
const refreshdict = ()=>{
  var dictValue = dictStore.getDictValue(appprops.dictname as string) as IDictItem[]
  (dictValue || []).forEach((oo:IDictItem) => {
  if (oo.value == appprops.value) {
    dictItem.value = oo
  }
  })
}
watch(()=>([appprops.dictname,appprops.value]),()=>{
  refreshdict()
},{
  deep:true,
})

//console.log("dictItem.value2", dictValue)
</script>
