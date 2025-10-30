
<template>
	<div class="app-container">
    <TableApp :handlers="handlers" :context="context" :meta="meta">
      <template #action="scope">
        <el-button type="primary" @click="trygrant(scope)">授权</el-button>
      </template>
    </TableApp>
	</div>
  <el-dialog v-model="_grant" :close-on-click-modal="false">
    <PanelGrant v-if="_grant" v-model="resids" @confirm="confirmgrant"></PanelGrant>
  </el-dialog>
   
  </template>
  
  <script setup lang="ts">
  import {ref} from "vue"
  import roleApi from "@/api/sys/role"
  import PanelGrant from "./panel-grant.vue"
  defineOptions({name:"sys-role"})
  const context = ref({
	perm:'sys:role',
	name:"role",
	title:"角色",
	primaryKey:"id",
  }) 
  const meta = ref([
  {type:'selection','column-key':"id",prop:"id","width":80},
	{ prop:"id",domType:"number",hidden: true ,dataType:"int64",label:"ID",sortable: false ,suportSearch: false ,suportCreate: false ,suportUpdate: false  },
	
	{ prop:"code",domType:"text",hidden: false ,dataType:"string",label:"角色码",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"name",domType:"text",hidden: false ,dataType:"string",label:"角色名称",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"state",domType:"dict",hidden: false ,option:"sys_enable",dataType:"int32",label:"状态",sortable: false ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"resId",domType:"text",hidden: true ,dataType:"string",label:"资源ID",sortable: false ,suportSearch: false ,suportCreate: false ,suportUpdate: false  },
	
  ])
  const handlers = ref(roleApi)
  // 鉴权
  var residarr = [] as number[]
  const resids = ref(residarr)
  const  _grant = ref(false)
  var  currentId = 0
  const trygrant = ({row}:any)=>{
    resids.value = row.rightIds||[]
    currentId = row.id
    _grant.value = true
  }
  const confirmgrant = (rightIds:number[])=>{
    roleApi.grant({id:currentId,rightIds}).then(res=>{
      resids.value = rightIds
      _grant.value = false
    })
  }
  </script>
