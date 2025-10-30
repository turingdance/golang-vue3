
<template>
	<div class="app-container">
		  <TableApp :handlers="handlers" :context="context" :tree="tree" :meta="meta">
			<template #create="scope">
				<RightsItem :scope="scope"></RightsItem>
			</template>
			<template #update="scope">
				<RightsItem :scope="scope"></RightsItem>
			</template>
		</TableApp>
	</div>
  </template>
  
  <script setup lang="ts">
  import RightsItem from "./com-rightsitem.vue"
  import {ref} from "vue"
  import rightsApi from "@/api/sys/rights"
  defineOptions({name:"sys-rights"})
  const context = ref({
	perm:'sys:rights',
	name:"rights",
	title:"资源",
	primaryKey:"id",
  }) 
  const meta = ref([
  {type:'selection','column-key':"id",prop:"id","width":80},
	{ prop:"id",domType:"number",hidden: false ,dataType:"int64",label:"资源ID",sortable: true ,suportSearch: false ,suportCreate: false ,suportUpdate: false  },
	
	{ prop:"type",domType:"dict",option:"sys_rights_type",hidden: false ,dataType:"string",label:"资源类型",sortable: true ,suportSearch: false ,suportCreate: true ,suportUpdate: true  },
	
		
		
	{ prop:"icon",domType:"text",hidden: false ,dataType:"string",label:"图标名称",sortable: false ,suportSearch: false ,suportCreate: true ,suportUpdate: true  },
	{ prop:"biz",domType:"text",hidden: false ,dataType:"string",label:"权限标识",sortable: true ,suportSearch: false ,suportCreate: true ,suportUpdate: true  },

	{ prop:"title",domType:"text",hidden: false ,dataType:"string",label:"资源名称",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
		
	{ prop:"pid",domType:"number",hidden: false ,dataType:"int64",label:"父级资源",sortable: false ,suportSearch: false ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"uri",domType:"text",hidden: false ,dataType:"string",label:"路由地址",sortable: true ,suportSearch: false ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"component",domType:"text",hidden: false ,dataType:"string",label:"组件地址",sortable: true ,suportSearch: false ,suportCreate: true ,suportUpdate: true  },	
	{ prop:"redirect",domType:"text",hidden: false ,dataType:"string",label:"跳转地址",sortable: false ,suportSearch: false ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"sortIndex",domType:"number",hidden: false ,dataType:"int32",label:"排序位",sortable: false ,suportSearch: false ,suportCreate: true ,suportUpdate: true  },
	{ prop:"alwaysShow",domType:"number",hidden: false ,dataType:"bool",label:"是否可见",sortable: true ,suportSearch: false ,suportCreate: true ,suportUpdate: true  },
	{ prop:"hidden",domType:"number",hidden: false ,dataType:"bool",label:"是否可见",sortable: true ,suportSearch: false ,suportCreate: true ,suportUpdate: true  },
		
  ])
  const handlers = ref(rightsApi)
  const tree = ref({
	idKey:"id",parentKey:"pid",childrenKey:"children"
  })
  </script>
