<template>
   <H5>请配置筛选条件</H5>
    <el-form :model="conds">
    <el-form-item  v-for="(item,index) in conds" :key="index" >
      <div class="row" style="width:80%;display:flex;flex-direction:row;justify-content:space-between">
        <div style="width:80%;display:flex;flex-direction:row;">
              <el-select v-model="item.field" style="width:30%">
                <template v-for="f,i in meta.filter(o=>!!o.suportSearch)" :key="i" >
                <el-option 
                  v-if="!!f.searchable"
                  :label="f.label" 
                  :value="f.prop" 
                  >
                </el-option>
              </template>
              </el-select>    
              <el-select v-model="item.op" style="width:30%">
                <el-option v-for="op in ops" :key="op.value" :label="op.label" :value="op.value"></el-option>
              </el-select>
              <div  style="max-width:40%">  
              <el-input  v-model="item.value" v-if="item.domtype=='text'" />
              <el-input v-model.number="item.value" v-else-if="item.domtype=='number'" />
              <el-date-picker type="datetime" format="YYYY-MM-DD hh:mm:ss" v-model="item.value" v-else-if="item.domtype=='datetime'" />
              <el-date-picker type="date" format="YYYY-MM-DD" v-model="item.value" v-if="item.domtype=='date'" />
              <PickImage v-model="item.value" v-else-if="item.domtype=='PickImage'"/>
              <PickFile v-model="item.value" v-else-if="item.domtype=='PickFile'"/>
              <PickAddress v-model="item.value" v-else-if="item.domtype=='PickAddress'"/>
              <el-input v-model="item.value" v-else/>
            </div>
      </div>
      <div class="row" style="width:20%">
        <el-button icon="delete" @click="deleteit(index)" cirlce></el-button>
        <el-button @click="add" icon="plus" cirlce></el-button>
      </div>
    </div>
    </el-form-item>

    

  </el-form>
  <H5>请选择要导出的数据列</H5>
  <el-table ref="reftable" :data="meta.filter(o=>(!(['selection','index'].includes(o.type))))">
      <el-table-column type="selection"></el-table-column>
      <el-table-column label="字段" prop="prop"></el-table-column>
      <el-table-column label="标题" prop="label"></el-table-column>
  </el-table>
   
  <el-row class="flex-x-end m-2">
    <el-button type="primary"  @click="confirm">导出</el-button>
  </el-row>
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
  meta:{
    type:Array,
    default :()=>([])
  }
})
const emits = defineEmits(['confirm'])
const reftable = ref()

const confirm = ()=>{
  let condsrows=conds.value.filter(oo=>(!!oo.field&&!!oo.op && !!oo.value))
   let rows = reftable.value.getSelectionRows()
   if(rows.length<1){
    ElMessage.error("请选择要导出的数据列")
    return 
   }
   let result = {
    meta:rows,
    cond:{
      conds:condsrows,
      pager:{pagegrom:0,pagesize:1024*1024},
      order:{field:"",method:""},
    }
   }
   emits('confirm',result)
}

const conds = ref([])

const ops = ref([
  {label:"等于(=)",value:"eq"},
  {label:"小于等于(<=)",value:"leq"},
  {label:"小于(<)",value:"le"},
  {label:"大于(>)",value:"ge"},
  {label:"大于或等于(>=)",value:"geq"},
  {label:"包含(like)",value:"like"},
  {label:"范围(between)",value:"between"},
  {label:"范围(in)",value:"in"},
])
const add = ()=>{
  //console.log('props',conds.value)
  conds.value.push({
    field:"",op:"",value:""
  })
  conds.value = conds.value
}
if(conds.value.length<1){
  add()
}
const deleteit = (index)=>{
  if(conds.value.length<2){
    conds.value[index] ={
    field:"",op:"",value:""
  }
    return 
  }
  conds.value.splice(index,1)
}
</script>
<style scope>
.row{
  display:flex;
  flex-direction:row;
  justify-content:space-between;
}
</style>
