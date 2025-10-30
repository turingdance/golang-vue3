
<template>
	<div class="app-container">
		  <TableApp :handlers="handlers" :context="context" :meta="meta">
          
      </TableApp>
	</div>
  </template>
  
  <script setup lang="ts">
  import {ref} from "vue"
  import dictApi from "@/api/sys/dict"
  defineOptions({name:"sys-dict"})
  const context = ref({
	perm:'sys:dict',
	name:"dict",
	title:"字典",
	primaryKey:"id",
  }) 
  const meta = ref([
  {type:'selection','column-key':"id",prop:"id","width":80},
	{ prop:"id",domType:"number",hidden: false ,dataType:"int64",label:"ID",sortable: true ,suportSearch: false ,suportCreate: false ,suportUpdate: false  },
	{ prop:"name",domType:"text",hidden: false ,dataType:"string",label:"编码",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	{ prop:"title",domType:"text",hidden: false ,dataType:"string",label:"名称",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	{ prop:"value",domType:"DictOption",hidden: false ,dataType:"array",label:"值",sortable: false ,suportSearch: false ,suportCreate: true ,suportUpdate: true },
  ])
  const handlers = ref(dictApi)
  </script>
