<template>
    <el-row v-for="(item,index) in model" :key="index">
      <el-col :span="9">
        <el-form-item  >
          <el-input type="text"  v-model="item.value" placeholder="编码" />
        </el-form-item>
      </el-col>
      <el-col :span="9">
        <el-form-item  >
          <el-input  type="text" v-model="item.label" placeholder="名称" />
        </el-form-item>
      </el-col>
      <el-col :span="6">
        <el-button icon="delete" @click="deleteit(index)" cirlce></el-button>
        <el-button @click="add" icon="plus" cirlce></el-button>
      </el-col>
  </el-row>
</template>
<script setup>
import { onMounted } from 'vue';

const model = defineModel({required:true,default:()=>([])})
const add = ()=>{
  model.value.push({
    label:"",value:""
  })
  model.value = model.value
}

const deleteit = (index)=>{
  if(model.value.length<2){
    model.value[index] ={label:"",value:""}
    return 
  }
  model.value.splice(index,1)
}

onMounted(()=>{
  if(model.value.length<1){
    add()
  }
})

</script>
<style scope>
.row{
  display:flex;
  flex-direction:row;
  justify-content:space-between;
}
</style>
