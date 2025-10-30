
<template>
	<div class="app-container">
		  <TableApp :handlers="handlers" :context="context" :meta="meta"></TableApp>
	</div>
  </template>
  
  <script setup>
  import {ref} from "vue"
  import lessonApi from "@/api/sys/config"
  const context = ref({
	perm:'system:config',
	name:"config",
	title:"系統參數",
	primaryKey:"id",
  }) 
  const meta = ref([
  	{type:'selection','column-key':"id",prop:"id","width":80},
	{ prop:"id",domType:"text",hidden: false ,dataType:"string",label:"配置ID",sortable: true ,suportSearch: true ,suportCreate: false ,suportUpdate: true  },
	{ prop:"name",domType:"text",hidden: false ,dataType:"string",label:"参数名称(英文)",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	{ prop:"label",domType:"text",hidden: false ,dataType:"string",label:"参数描述(中文)",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	{ prop:"value",domType:"text",hidden: false ,dataType:"string",label:"参值",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
  ])
  const handlers = ref(lessonApi)
  </script>
