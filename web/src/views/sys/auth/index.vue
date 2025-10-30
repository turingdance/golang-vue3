
<template>
	<div class="app-container">
		  <TableApp :handlers="handlers" :context="context" :meta="meta"></TableApp>
	</div>
  </template>
  
  <script setup lang="ts">
  import {ref} from "vue"
  import authApi from "@/api/sys/auth"
  defineOptions({name:"sys-auth"})
  const context = ref({
	perm:'sys:auth',
	name:"auth",
	title:"权限",
	primaryKey:"id",
  }) 
  const meta = ref([
  {type:'selection','column-key':"id",prop:"id","width":80},
	{ prop:"id",domType:"number",hidden: false ,dataType:"int64",label:"ID",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"accId",domType:"number",hidden: false ,dataType:"int64",label:"用户ID",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"roleId",domType:"number",hidden: false ,dataType:"int64",label:"角色ID",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"teamId",domType:"number",hidden: false ,dataType:"int64",label:"所属团队",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"status",domType:"number",hidden: false ,dataType:"int32",label:"权限状态",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"expireAt",domType:"text",hidden: false ,dataType:"types.Date",label:"权限过期时间",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
  ])
  const handlers = ref(authApi)
  </script>