<template>
  <el-form :model="form" label-width="auto">
    
    <el-form-item  v-for="(item,index) in conds" :key="index" >
      <div class="row" style="width:80%;display:flex;flex-direction:row;justify-content:space-between">
        <div style="width:80%;display:flex;flex-direction:row;">
              <el-select v-model.underline="item.field" style="width:30%" @change="chosecurrentfield(item,index)" >
                <template v-for="f,i in meta" :key="i" >
                <el-option 
                  v-if="!!f.suportSearch"
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
              <el-input  v-model="item.value" v-if="item.meta.domType=='text'" />
              <el-input v-model.number="item.value" v-else-if="item.meta.domType=='number'" />
              <el-date-picker type="datetime" format="YYYY-MM-DD hh:mm:ss" v-model="item.value" v-else-if="item.meta.domType=='datetime'" />
              <el-date-picker type="date" format="YYYY-MM-DD" v-model="item.value" v-else-if="item.meta.domType=='date'" />
              <PickImage v-model="item.value" v-else-if="item.meta.domType=='PickImage'"/>
              <PickFile v-model="item.value" v-else-if="item.meta.domType=='PickFile'"/>
              <DictDom v-model="item.value" v-else-if="item.meta.domType=='dict'" :datatype="item.meta.dataType" :dictname="item.meta.option" />
              <PickAddress v-model="item.value" v-else-if="item.meta.domType=='PickAddress'"/>
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
  <el-row class="flex-x-end m-2">
    <el-button type="primary" @click="confirm" >搜索</el-button>
  </el-row>
</template>
<script setup>
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

const condarr = defineModel()
/*
* 驼峰转换下划线
*/ 
function underline(name) {
  return name.replace(/([A-Z])/g,"_$1").toLowerCase();
}

const emits = defineEmits(['update:modelValue','confirm'])

const conds = ref(Object.assign([],condarr.value))
const  chosecurrentfield = (item,index)=>{
  conds.value[index].meta = appprops.meta.find(o=>o.prop==item.field)
}
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
    field:"",op:"",value:"",meta:{}
  })
  conds.value = conds.value
}
if(conds.value.length<1){
  add()
}
const deleteit = (index)=>{
  if(conds.value.length<2){
    conds.value[index] ={
    field:"",op:"",value:"",meta:{}
  }
    return 
  }
  conds.value.splice(index,1)
}
const confirm = ()=>{
  let result=conds.value.filter(oo=>(!!oo.field&&!!oo.op && !!oo.value))
  result.forEach(o=>{
    o.field = underline(o.field)
  })
  conds.value = result
  condarr.value=result
  emits('confirm',result)
}
</script>
<style scope>
.row{
  display:flex;
  flex-direction:row;
  justify-content:space-between;
}
</style>
