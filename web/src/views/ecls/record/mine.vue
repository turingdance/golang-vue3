<template>
<div class="app-container">
       <el-table :data="recordList">
        <el-table-column prop="createAt" label="学习时间"  width="220" />
        <el-table-column prop="lessonId" label="课程ID" />
        <el-table-column prop="title" label="课程名称" />
        <el-table-column prop="uri" label="课程查看" >
          <template #default="scope">
            <el-button @click="showlesson(scope.row)">观看</el-button>
          </template>
        </el-table-column>
       </el-table>
</div>
    
      <Preview v-model="openit" :lesson="lesson" v-if="openit"></Preview>
  </template>
  <script setup lang="js">
  import recordApi from '@/api/ecls/record';
  const recordList = ref([])
  const openit = ref(false)

  function initdata(){
      recordApi.mine({pageSize:1024}).then(res=>{
        recordList.value = res.data ||[]
      })
  }
const lesson = ref({})
function showlesson(row){
  lesson.value = row
  openit.value = true
}
 initdata()
  </script>
    