<template>
  <el-button @click="_open=true" type="primary" >批量添加</el-button>
  <el-dialog v-model="_open" >
       <el-form :model="form" ref="formRef" label-width="auto" label-position="top">
       <el-form-item   label="封面" prop="cover" :rules="[{ required: true, message: '请上传封面' }]">
           <PickImage :multiple="true" style="width:180px;height:150px;"  v-model="form.cover" placeholder="这个权限是用来做什么的" />
       </el-form-item>
       <el-form-item   label="资源文件" prop="uris"  :rules="[{ required: true, message: '请上传资源' }]">
        <PickFile  
         :multiple="true" 
         :limit="10"
         v-model:filekeys="form.uris"
         v-model:filelist="form.lists"
        ></PickFile>
       </el-form-item>
       <el-form-item   label="标签" :rules="[{ required: true, message: '请上传资源' }]">
         <!-- <el-input-tag  v-model="form.tag" placeholder="请输入标签" /> -->
         <DictDom dictname="sys_lesson_cate" v-model="form.cate"></DictDom>
       </el-form-item>
   </el-form>
   <el-row class="flex-x-end m-2">
       <el-button type="primary" @click="confirm" >确认</el-button>
     </el-row>
    </el-dialog>
   
   </template>
   <script setup >
import lessonApi from "@/api/ecls/lesson"
import {getFileType} from "@/utils/index"
   const form = ref({
    lists:[],
    urls:[],
    cover:"",
    tag:[],
    cate:"",
   })
   const formRef = ref()
   const confirm = ()=>{
      formRef.value.validate((valid) => {
         if(valid){
          /*
          `id` varchar(32) NOT NULL COMMENT '视频ID',
          `media` varchar(10) DEFAULT NULL COMMENT '资源类型',
          `cover` varchar(255) DEFAULT NULL COMMENT '封面',
          `title` varchar(255) DEFAULT NULL COMMENT '标题',
          `memo` varchar(255) DEFAULT NULL COMMENT '描述',
          `uri` varchar(255) DEFAULT NULL COMMENT '资源key',
          `tag` json DEFAULT NULL COMMENT '标签',
          */
          let postdata = form.value.lists.map((item,index)=>{
            return {
              media: getFileType(item.name),
              cover:form.value.cover,
              title:item.name,
              uri:form.value.lists[index].key,
              cate:form.value.cate,
            }
          })
           lessonApi.batchAdd(postdata)
         }
      })
   }
   const _open = ref(false)
   </script>
   <style scope>
   .row{
     display:flex;
     flex-direction:row;
     justify-content:space-between;
   }
   </style>
   