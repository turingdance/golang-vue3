<template>
   <el-form :model="form" ref="formRef" label-width="auto" label-position="top">
      
       <el-form-item   label="业务名称(英文)" 
       prop="biz"
       :rules="[{ required: true, message: '请填写业务名称' }]"
       >
           <el-input   v-model="form.biz"  placeholder="sys:biz"/>
       </el-form-item>  
       <el-form-item   label="权限描述,中文" prop="title" :rules="[{ required: true, message: '请描述业务' }]">
           <el-input  v-model="form.title" placeholder="这个权限是用来做什么的" />
       </el-form-item>

       <el-form-item  label="图标" >
         <IconSelect  v-model="form.icon" style="width:100%" />
       </el-form-item>

       <el-form-item   label="uri地址,按钮级别可携带协议如get:/sys/biz/add" prop="uri"  :rules="[{ required: true, message: '请填写连接地址' }]">
         <el-input  v-model="form.uri" placeholder="/sys/biz" />
       </el-form-item>

       <el-form-item    >
         <template #label>
         节点等级
         </template>
         <el-switch v-model="islevel1"   :active-value="true" :inactive-value="false" active-text="这是1级节点" inactive-text="这不是1级节点" />
       </el-form-item>

       <el-form-item label="节点挂载位置"  v-if="!islevel1" prop="pid"  :rules="[{ required: true, message: '请选择业务节点' }]" >
         
         <el-tree-select

            v-model="form.pid"
            :data="scope.tabledata"
            :render-after-expand="false"
            style="width: 100%"
            node-key="id"
            :check-strictly="true"
            :props="{'label':'title'}"
         >
         <template #default="{ data: { title,type } }">
                  <div v-if="type!='api'">{{ title }}</div>>
         </template>
         </el-tree-select>
       </el-form-item>
       
       <el-form-item   label="权限类型" prop="type"  :rules="[{ required: true, message: '请选择业务类型' }]">
         <DictDom dictname="sys_rights_type" style="width:100%"  v-model="form.type"> </DictDom>
       </el-form-item>  
      

       
       <el-form-item v-if="form.type=='view'" prop="pid"  :rules="[{ required: true, message: '请填写业务组件' }]" label="组件地址,相对于src/views/的目录地址,不需要带.vue后缀 分组类型填Layout" >
         <el-input  v-model="form.component" placeholder="/sys/biz/index" />
       </el-form-item>
       <el-form-item v-if="form.type=='view'||form.type=='group'"  label="跳转url" >
         <el-input  v-model="form.redirect" placeholder="/sys/biz/index" />
       </el-form-item>


       <el-form-item v-if="form.type=='view'||form.type=='group'"  label="是否启用常驻菜单栏功能" >
         <el-switch  v-model="form.alwaysShow" :active-value="true" :inactive-value="false" active-text="启用"  inactive-text="停用" />
       </el-form-item>
       <el-form-item v-if="form.type=='view'||form.type=='group'"  label="是否启用菜单栏隐身功能" >
         <el-switch  v-model="form.hidden" :active-value="true" :inactive-value="false" active-text="启用"  inactive-text="停用" />
       </el-form-item>
        
       <el-form-item   label="排序位,越小越靠前" >
         <el-input  v-model.number="form.sortIndex" placeholder="1/2/3/1.5/1.6" />
       </el-form-item>

   </el-form>
   <el-row class="flex-x-end m-2">
       <el-button type="primary" @click="confirm" >确认</el-button>
     </el-row>
   </template>
   <script setup >
import { onMounted } from "vue";
   const form = ref({})
   const formRef = ref()
   const confirm = ()=>{
      formRef.value.validate((valid) => {
         if(valid){
            appprops.scope.action(form.value)
            console.log(appprops.scope)
         }
      })
   }
   const appprops = defineProps({
      scope:{
         type:Object
      }
   })
   const islevel1 = ref(true)
   watch(()=>appprops.scope,(scope)=>{
    if(!!scope.formdata){
      form.value = Object.assign({},scope.formdata)
      islevel1.value = !form.value.pid
    }
   })
   onMounted(()=>{
    console.log('appprops.scope.formdata',appprops.scope.formdata)
    if(!!appprops.scope.formdata){
      form.value = Object.assign({},appprops.scope.formdata)
      islevel1.value = !form.value.pid
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
   