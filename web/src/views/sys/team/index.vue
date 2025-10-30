
<template>
	<div class="app-container">
		  <TableApp :handlers="handlers" :context="context" :meta="meta"></TableApp>
	</div>
  </template>
  
  <script setup lang="ts">
  import {ref} from "vue"
  import teamApi from "@/api/sys/team"
  defineOptions({name:"sys-team"})
  const context = ref({
	perm:'sys:team',
	name:"team",
	title:"组织",
	primaryKey:"id",
  }) 
  const meta = ref([
  {type:'selection','column-key':"id",prop:"id","width":80},
	{ prop:"id",domType:"number",hidden: false ,dataType:"int64",label:"组织id",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"name",domType:"text",hidden: false ,dataType:"string",label:"团队名称",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"creator",domType:"number",hidden: false ,dataType:"int64",label:"创建者",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"trade",domType:"text",hidden: false ,dataType:"string",label:"行业",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
  ])
  const handlers = ref(teamApi)
  </script>
