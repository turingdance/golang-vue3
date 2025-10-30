
<template>
	<div class="app-container">
		<TableApp :handlers="handlers" :context="context" :meta="meta">
			<template #toolbar>
				<ComBatchAdd></ComBatchAdd>
			</template>
		</TableApp>
	</div>
  </template>
  
  <script setup>
  import {ref} from "vue"
  import ComBatchAdd from "./com-batchadd.vue"
  import lessonApi from "@/api/ecls/lesson"

  const context = ref({
	perm:'ecls:lesson',
	name:"lesson",
	title:"课件资源",
	primaryKey:"id",
  }) 
  const meta = ref([
  {type:'selection','column-key':"id",prop:"id","width":80},
	{ prop:"id",domType:"text",hidden: false ,dataType:"string",label:"视频ID",sortable: true ,suportSearch: true ,suportCreate: false ,suportUpdate: true  },
	{ prop:"media",domType:"dict",option:"sys_media_type",hidden: false ,dataType:"string",label:"资源类型",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"cover",domType:"PickImage",hidden: false ,dataType:"string",label:"封面",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"title",domType:"text",hidden: false ,dataType:"string",label:"标题",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"memo",domType:"text",hidden: false ,dataType:"string",label:"描述",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"uri",domType:"PickFile",hidden: false ,dataType:"string",label:"资源key",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
	{ prop:"cate",domType:"dict",option:"sys_lesson_cate",hidden: false ,dataType:"string",label:"类型",sortable: true ,suportSearch: true ,suportCreate: true ,suportUpdate: true  },
	
  ])
  const handlers = ref(lessonApi)
  </script>
