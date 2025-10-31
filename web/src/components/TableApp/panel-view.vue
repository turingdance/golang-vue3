<template>
  <el-form :model="form" label-width="auto" :disabled="true" :labelPosition="labelPosition">
    <el-form-item disabled  v-for="(item,index) in appdata" :label="item.label" :key="index" >
      <el-input  v-model="item.value" v-if="item.domType=='text'" />
        <el-input v-model.number="item.value" v-else-if="item.domType=='number'" />
        <el-input  v-model="item.value" type="textarea" v-else-if="item.domType=='textarea'" />
        <el-date-picker type="datetime" v-bind="item.attr" format="YYYY-MM-DD hh:mm:ss" v-model="item.value" v-else-if="item.domType=='datetime'" />
        <el-date-picker type="date" v-bind="item.attr" format="YYYY-MM-DD" v-model="item.value" v-else-if="item.domType=='date'" />
        <el-switch v-model="item.value" v-bind="item.attr" :style="item.style" v-else-if="item.domType=='switch'"  :active-text="item.option.activeText" :inactive-text="item.option.inactiveText" />
        <WangEditor  v-model="item.value" v-bind="item.attr" type="richtext" v-else-if="item.domType=='richtext'" />
        <PickImage v-model="item.value" v-bind="item.attr" :style="item.style" v-else-if="item.domType=='PickImage'"/>
        <PickImage v-model="item.value" v-bind="item.attr" v-else-if="item.domType=='PickImage'"/>
        <PickFile v-model="item.value" v-else-if="item.domType=='PickFile'"/>
        <PickAddress v-model="item.value" v-else-if="item.domType=='PickAddress'"/>
        <DictOption v-model="item.value" v-else-if="item.domType=='DictOption'" />
        <template v-else-if="item.domType=='dict'">
          <DictDom v-model="item.value" :datatype="item.dataType" :dictname="item.option" />
        </template>
        <el-input v-model="item.value" v-else/>
    </el-form-item>
  </el-form>
</template>
<script setup>
import { ElMessage } from 'element-plus'

const appprops = defineProps({
  context:{
    type:Object,
    default:()=>({
      title:"",
      name:"",
      primaryKey:"",
    }),
  },
  action:{
    type:Function,
    required:true,
  },
  labelPosition:{
    type:String,
    default:()=>"top",
  },
  modelValue:{
    type:Array,
    default :()=>([{field:"",op:"eq",value:""}])
  },
  formdata:{
    type:Object,
    default:()=>({})
  },
  meta:{
    type:Array,
    default :()=>([])
  }
})
const emits = defineEmits(['update:modelValue','confirm'])

const appdata = ref([])
const initappdata = ()=>{
  let tmp = appprops.meta.filter(o=>!!o.suportUpdate).map(oo=>{
    let  value = appprops.formdata[`${oo.prop}`]
    return {domType:oo.domType,label:oo.label,style:oo.style,attr:oo.attr,field:oo.prop,value:value,option:oo.option,dataType:oo.dataType}
  })
  //console.log('tmp',tmp,appprops.meta)
  appdata.value = tmp
}
initappdata()
watch(()=>appprops.formdata,()=>{
  initappdata()
})
const confirm = ()=>{
    let postdata = {}
    appdata.value.reduce((sum,curr)=>{
      //console.log('reduce',sum,curr)
      if(typeof curr.value!="undefined"){
        sum[curr.field] = curr.value
      }
        return sum
    },postdata)
    postdata[`${appprops.context.primaryKey}`] = appprops.formdata[`${appprops.context.primaryKey}`]
    if(!postdata[`${appprops.context.primaryKey}`]){
      ElMessage.error('缺少参数ID')
      return 
    }
    appprops.action(postdata).then(res=>{
      emits('confirm',res.data)
      if(res.msg){
        ElMessage.success(res.msg)
      }
    }).catch(err=>{
      ElMessage.error(err.message)
    })
}
</script>
<style scope>
.row{
  display:flex;
  flex-direction:row;
  justify-content:space-between;
}
</style>
