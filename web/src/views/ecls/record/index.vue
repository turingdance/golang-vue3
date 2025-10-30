
<template>
	<div class="app-container">
		  <TableApp :handlers="handlers" :context="context" :meta="meta"></TableApp>
	</div>
  </template>
  
  <script setup>
  import {ref} from "vue"
  import recordApi from "@/api/ecls/record"
  const context = ref({
	perm:'ecls:record',
	name:"record",
	title:"上课记录",
	primaryKey:"id",
  }) 
  const meta = ref([
  {type:'selection','column-key':"id",prop:"id","width":80},
	{ prop:"id",domType:"number",hidden: false ,dataType:"int64",label:"ID",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"createAt",domType:"text",hidden: false ,dataType:"types.Date",label:"上课时间",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"userId",domType:"text",hidden: false ,dataType:"string",label:"用户ID",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"lessonId",domType:"text",hidden: false ,dataType:"string",label:"课程ID",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"createDate",domType:"text",hidden: false ,dataType:"types.Date",label:"创建日期",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
  ])
  const handlers = ref(recordApi)
  </script>
