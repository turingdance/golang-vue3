<template>
    <slot>
      <template v-if="appprops.datatype.indexOf('int')>-1">
        <el-select :disabled="disabled"  v-model.number="status" :style="{'min-width':appprops.minWidth+'px',width:width}">
          <el-option v-for="item in dictValue" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </template>
      <template v-else>
        <el-select :disabled="disabled" v-model="status" :style="{'min-width':appprops.minWidth+'px',width:width}">
          <el-option v-for="item in dictValue" :key="item.value" :label="item.label" :value="item.value" />
        </el-select>
      </template>
    </slot>
</template>
<script lang="ts" setup module="smstask">
import { computed, onMounted, ref } from "vue"
import dictApi, { IDictItem } from "@/api/sys/dict"
import { useVModel } from "@vueuse/core"
const appprops = defineProps({
  modelValue: {
    type: String
  },
  datatype:{
    type:String,
    default: "string" 
  },
  disabled: {
    type: Boolean,
    default: false
  },
  dictname: {
    type: String,
    require: true
  },
  width:{
    type:[Number,String],
    default:'100%'
  },
  minWidth:{
    type:[Number],
    default:120,
  }
})
const width = computed(()=>{
  if(typeof appprops.width == "number"){
    return appprops.width+'px'
  }
  else if(typeof appprops.width == "string"){
    return appprops.width
  }else{
    return 'calc( 100% - 5px )'
  }
})
const emits = defineEmits(["update:modelValue"])
const status = useVModel(appprops, "modelValue", emits)
const dictValue = ref<IDictItem[]>([])
const initdata = (name:string) => {
  dictApi.option(name).then((res) => {
    res.data.forEach((oo:any)=>{
      if(appprops.datatype.indexOf('int')>-1){
        oo.value = + oo.value
      }
    })
    dictValue.value = res.data
  })
}
onMounted(() => {
  initdata(appprops.dictname as string)
})
</script>
<style scope>
</style>
