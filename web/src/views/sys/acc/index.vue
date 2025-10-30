
<template>
	<div class="app-container">
		  <TableApp :handlers="handlers" :context="context" :meta="meta">
        <template #action="scope">
          <el-button type="danger" @click="updatepassword(scope.row.userId)">重置密码</el-button>
        </template>
        <template #toolbar="scope">
          <el-button type="danger" @click="showdialog=true">批量导入</el-button>
        </template>
      </TableApp>
      <el-dialog v-model="showdialog" title="批量上传">
        <el-upload
            class="upload-demo"
            drag
            :http-request="importbatch"
            :multiple="false"
            accept=".xls,.xlsx, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet,application/vnd.ms-excel"
            :auto-upload="true"
          >
            <el-icon class="el-icon--upload">
              <upload-filled /></el-icon>
            <div class="el-upload__text">
             拖拽文件到这里 <em>或者</em>点击这里上传文件
            </div>
            <template #tip>
              <div class="el-upload__tip">
                <label> 文件大小不超超过1MB</label> <el-button type="primary" link  @click="downtemplate">点击这里下载账户导入模板</el-button>
              </div>
            </template>
          </el-upload>
          <template #footer>
              <el-button   @click="showdialog=false">取消</el-button>
          </template>
      </el-dialog>
	</div>
  </template>
  
  <script setup lang="ts">
  import {ref} from "vue"
  import accApi from "@/api/sys/acc"
  defineOptions({name:"sys-acc"})
  const context = ref({
	perm:'sys:acc',
	name:"acc",
	title:"账号",
	primaryKey:"userId",
  }) 
  const meta = ref([
  {type:'selection','column-key':"id",prop:"id","width":80},
	{ prop:"userId",domType:"text",hidden: false ,dataType:"int64",label:"用户ID",sortable: true ,suportSearch: false ,suportCreate: false ,suportUpdate: false  },
	{ prop:"username",domType:"text",hidden: false ,dataType:"string",label:"用户名",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	{ prop:"nickname",domType:"text",hidden: false ,dataType:"string",label:"昵称",sortable: true ,suportSearch: false ,suportCreate: false ,suportUpdate: false  },
	{ prop:"role.name",domType:"text",hidden: true ,dataType:"string",label:"角色",sortable: true ,suportSearch: false ,suportCreate: false ,suportUpdate: false  },
	{ prop:"createAt",domType:"datetime",hidden: false ,dataType:"types.DateTime",label:"创建时间",sortable: true ,suportSearch: false ,suportCreate: false ,suportUpdate: false  },
  ])
  const handlers = ref(accApi)

  const updatepassword = (userId:number|string)=>{
      ElMessageBox.prompt("请输入新的密码","设置密码", {
confirmButtonText: '确定',
cancelButtonText: '取消',
inputPlaceholder:"密码6-32位,必须包含大写或小写字母+数字+特殊字符",
inputType: 'text',
inputPattern: /^[a-zA-Z0-9\,\!\?\:\'\"\(\)\$\%\^\*\@]{6,32}$/,
inputErrorMessage: '密码至少长6位'
}).then(({ value }) => {
    accApi.updatePwd({userId:userId,password:value})
}).catch(() => {

})
  }

const  showdialog = ref(false)
function downtemplate(){
  accApi.template().then(blob=>{
    var downloadLink = document.createElement('a');
      downloadLink.href = URL.createObjectURL(blob);
      // 模拟用户点击下载链接
      downloadLink.download = '导入账号模板.xlsx';
      downloadLink.click();
  })
}  

function importbatch(param){
  let {file} = param
  const formatData = new FormData()
  formatData.append("file",file)
  return accApi.importbatch(formatData).then(res=>{ 
      ElMessage.success(res.msg)
      return res
  }).catch(err=>{
      ElMessage.error(err.message)
  })

}
  </script>