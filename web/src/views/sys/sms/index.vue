
<template>
	<div class="app-container">
		  <TableApp :handlers="handlers" :context="context" :meta="meta"></TableApp>
	</div>
  </template>
  
  <script setup lang="ts">
  import {ref} from "vue"
  import smsApi from "@/api/sys/sms"
  defineOptions({
    name:"sys-sms"
  })
  const context = ref({
	perm:'sys:sms',
	name:"sms",
	title:"短信",
	primaryKey:"id",
  }) 
  const meta = ref([
  {type:'selection','column-key':"id",prop:"id","width":80},
	{ prop:"id",domType:"number",hidden: false ,dataType:"int64",label:"ID",sortable: true ,suportSearch: false ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"code",domType:"text",hidden: false ,dataType:"string",label:"验证码",sortable: true ,suportSearch: false ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"phoneNumbers",domType:"text",hidden: false ,dataType:"string",label:"手机号",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"templeteId",domType:"text",hidden: false ,dataType:"string",label:"模板ID",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"result",domType:"text",hidden: false ,dataType:"string",label:"发送结果",sortable: true ,suportSearch: true,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"createAt",domType:"text",hidden: false ,dataType:"types.Date",label:"发布时间",sortable: true ,suportSearch:false ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"success",domType:"dict",option:"sms_status",hidden: false ,dataType:"int32",label:"是否成功",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"sendAt",domType:"text",hidden: false ,dataType:"types.Date",label:"发送时间",sortable: true ,suportSearch: false ,suportCreate: true ,suportUpdate: true  },
	
  ])
  const handlers = ref(smsApi)
  </script>
