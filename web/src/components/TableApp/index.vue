<template>
  
  
    <el-card shadow="hover" style="min-height: calc(100vh - 120px)">
    <template #header>
      <div class="card-header" v-if="model!='chose'">
        <el-button v-perm="context.perm+':search'" @click="_filter=true" icon="filter">筛选</el-button>
        <el-button v-perm="context.perm+':search'" @click="handlesearch" icon="Refresh">刷新</el-button>
        <el-button type="primary" v-perm="context.perm+':create'" @click="showcreate" icon="plus">添加</el-button>
              <el-popconfirm title="确定要删除选中的记录吗?"
              v-perm="context.perm+':delete'" 
              @confirm="deletebatch"
              >
                <template #reference>
                   <el-button type="danger" icon="delete">删除</el-button>
                </template>
              </el-popconfirm>
        <el-button type="warning" v-perm="context.perm+':update'" @click="updateOne" icon="edit">修改</el-button>
        <el-button type="primary" v-perm="context.perm+':export'" @click="_export=true" icon="download">导出</el-button>
        <slot name="toolbar" :selected="selected" :tabledata="tabledata">

        </slot>
      </div>
    </template>
      <el-table
      :data="tabledata"
      ref="reftable"
      stripe 
      style="width: 100%"
      @sort-change="onsortchange"
      @selection-change="onselectionchange"
      v-loading="loading"
      :row-key="context.primaryKey"
      >
          <slot>
            <template v-for="item,index in meta.filter(oo=>!oo.hidden)" :key="index">
              item.domType={{ item.domType }}
            <el-table-column  v-if="item.domType=='dict'" :label="item.label">
                <template #default="scope">
                  <DictItem :dictname="item.option" :value="scope.row[`${item.prop}`]"></DictItem>
                </template>
            </el-table-column>
            <el-table-column  v-else-if="item.domType=='DictOption'" :label="item.label">
                <template #default="scope">
                   <TagList :value="scope.row[`${item.prop}`]"></TagList>
                </template>
            </el-table-column>
            <el-table-column  v-else-if="item.domType=='PickImage'" :label="item.label">
                <template #default="scope">
                   <el-image v-if="item.option && item.option.ref" :src="scope.row[item.option.ref]"></el-image>
                   <AuthImage v-else :src="item.value"/>
                </template>
            </el-table-column>
            <el-table-column  v-else v-bind="item"></el-table-column>
          </template>
          </slot>
          <el-table-column label="操作" fixed="right" :width="200">
            <template #default="scope">
              <el-button-group>
              <el-popconfirm title="确定要删除该记录吗?"
              v-perm="context.perm+':delete'" v-if="model!='chose'" 
              @confirm="handledelete(scope.row,scope.index)"
              >
                <template #reference>
                          <el-button type="danger" circle icon="delete"></el-button>
                </template>
              </el-popconfirm>
              <el-button type="warning" v-perm="context.perm+':update'" @click="showupdate(scope.row)" v-if="model!='chose'" circle icon="edit"></el-button>
              <el-button type="primary" v-perm="context.perm+':search'" circle icon="view" @click="showview(scope.row)" ></el-button>
                <el-button type="primary" v-perm="context.perm+':search'" v-if="model=='chose'" @click="handleemit('data',scope.row)" icon="edit">选择</el-button>
                <slot name="action" :row="scope.row" :$index="scope.$index" :column="scope.column">

                </slot>  
            </el-button-group>
              
            </template>
           </el-table-column>
      </el-table>
    <template #footer>
      <el-pagination
      v-if="total>0"
      v-model:current-page="pager.pagefrom"
      v-model:page-size="pager.pagesize"
      :page-sizes="pagesizes"
      :size="size"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total"
      @size-change="handlesizechange"
      @current-change="handlecurrentchange"
    ></el-pagination>

    </template>
  </el-card>
<el-dialog v-model="_filter" :show-close="false"  :close-on-click-modal="false">
  <template #header="{ close, titleId, titleClass }">
      <div class="header">
        <div :id="titleId" :class="titleClass">请配置筛选条件</div>
        <el-button type="danger" circle icon="CircleCloseFilled" @click="close"></el-button>
      </div>
    </template>
    <slot name="filter" :meta="meta" :context="context" :condarr="condarr" :confirm="handlefilterconfirm">
       <PanelFilter :meta="meta"  :context="context" v-model="condarr" @confirm="handlefilterconfirm"></PanelFilter>
    </slot> 
</el-dialog>
<el-dialog v-model="_export" :show-close="false"  :close-on-click-modal="false">
  <template #header="{ close, titleId, titleClass }">
      <div class="header">
        <div :id="titleId" :class="titleClass">请配置导出参数</div>
        <el-button type="danger" circle icon="CircleCloseFilled" @click="close"></el-button>
      </div>
    </template>
    <slot name="export" :meta="meta" :context="context" :formdata="formdata" :action="actionfunc">
    <PanelExport :meta="meta" :context="context" @confirm="handleexportconfirm"></PanelExport>
</slot>

</el-dialog>
<!-- 修改 -->
<el-dialog v-model="_update" :show-close="false"  :close-on-click-modal="false">
  <template #header="{ close, titleId, titleClass }">
      <div class="header">
        <div :id="titleId" :class="titleClass">请配置{{context.title}}</div>
        <el-button type="danger" circle icon="CircleCloseFilled" @click="close"></el-button>
      </div>
    </template>
    <slot name="update" :meta="meta" :tabledata="tabledata" :context="context" :formdata="formdata" :action="actionfunc">
         <PanelUpdate 
         :labelPosition="labelPosition"
         :meta="meta" @confirm="()=>{
        _update=false
        handlesearch()
        }" :context="context" :formdata="formdata" :action="actionfunc" ></PanelUpdate>
      </slot>
</el-dialog>
<!-- 创建 -->
<el-dialog v-model="_create" :show-close="false"  :close-on-click-modal="false">
  <template #header="{ close, titleId, titleClass }">
      <div class="header">
        <div :id="titleId" :class="titleClass">请配置{{context.title}}</div>
        <el-button type="danger" circle icon="CircleCloseFilled" @click="close"></el-button>
      </div>
    </template>
    <slot name="create" 
    :meta="meta" 
    :context="context" 
    :formdata="formdata" 
    :tabledata="tabledata"
    :action="actionfunc">
      <PanelCreate 
      :labelPosition="labelPosition"
       :meta="meta"  
        @confirm="()=>{
        _create=false
        handlesearch()
        }" 
        :context="context" 
        :formdata="formdata" 
        :action="actionfunc" >
      </PanelCreate>
    </slot>
</el-dialog>

<el-dialog v-model="_view" :show-close="false"  :close-on-click-modal="false">
  <template #header="{ close, titleId, titleClass }">
      <div class="header">
        <div :id="titleId" :class="titleClass">{{context.title}}</div>
        <el-button type="danger" circle icon="CircleCloseFilled" @click="close"></el-button>
      </div>
    </template>
    <slot name="view" :meta="meta" :context="context" :formdata="formdata" :action="actionfunc">
      <PanelView :labelPosition="labelPosition" :meta="meta"   :context="context" :formdata="formdata" :action="actionfunc" ></PanelView>
    </slot>
</el-dialog>
</template>
<script setup>
import {ref,onMounted} from "vue"
import {camelToUnderscore} from "@/utils/index"
import PanelFilter from "./panel-fileter.vue"
import PanelCreate from "./panel-create.vue"
import PanelUpdate from "./panel-update.vue"
import PanelExport from "./panel-export.vue"
import PanelView from "./panel-view.vue"
import FileSaver from "file-saver"
const model = ref('')
// 筛选
const condarr = ref([])
const _filter = ref(false)
const handlefilterconfirm = ($e)=>{
    _filter.value = false
    handlesearch()
}

const selected = ref([])
const onselectionchange = (value)=>{
  //console.log(value)
  selected.value=value
}
const appprops=defineProps({
  context:{
    type:Object,
    default:()=>({
      title:"",
      name:"",
      primaryKey:"",
      perm:"",
    }),
  },
  meta:{
    type:Array,
    default:()=>{
      return [
        {field:"",title:"",domtype:"",option:{},datatype:""}
      ]
    }
  },
  size:{
    type:String,
    default:"small",
  },
  tree:{
    type:Object,
  },
  pagesizes:{
    type:Array,
    default:()=>([10,20,50,100,200,500,1000,2000,5000,10000,100000])
  },
  labelPosition:{
    type:String,
    default:()=>"top",
  },
  handlers:{
    type:Object,
    default:()=>{
      return {
        search:()=>{
          return  Promise.reject(new Error("not implents"))
        },create:()=>{
          return  Promise.reject(new Error("not implents"))
        },deleteIt:()=>{
          return  Promise.reject(new Error("not implents"))
        },update:()=>{
          return  Promise.reject(new Error("not implents"))
        },exportxls:()=>{
          return  Promise.reject(new Error("not implents"))
        }
      }
    }
  }
})

const deletebatch = ()=>{
   let rows = reftable.value.getSelectionRows()
    if(rows.length==0){
      ElMessage.warning("请选择要操作的数据")
      return 
    }
    let key = appprops.context.primaryKey
    let value = rows.map(o=>o[`${key}`])
    let postdata = {}
    postdata[`${key}s`] = value
    appprops.handlers.deleteIts(postdata)
}

const loading = ref(false)
const _export = ref(false)
const handleexportconfirm = (data)=>{
    _filter.value = false
    handleexport(data)
}


const updateOne = ()=>{
  let rows = reftable.value.getSelectionRows()
    //console.log(rows)
    if(rows.length==0){
      ElMessage.warning("请选择要操作的数据")
      return 
    }
    if(rows.length>1){
      ElMessage.warning("一次只能操作一行数据")
      return 
    }
    showupdate(rows[0])
}

const reftable = ref()
const handledelete =(row,index)=>{
    var key = appprops.context.primaryKey
    let primaryValue = row[`${key}`]
    if(!primaryValue){
      ElMessage.error('请输入ID标识')
      return 
    }
    appprops.handlers.deleteIt(primaryValue).then(res=>{
      tabledata.value.splice(index,1)
    })
}

const _create = ref(false)
const _update = ref(false)
const _view = ref(false)
const showcreate = ()=>{
  actionfunc.value = appprops.handlers.create
  _create.value = true
  formdata.value = {}
}
const formdata = ref({})
const showupdate = (row)=>{
  actionfunc.value = appprops.handlers.update
  _update.value = true
  formdata.value = row
}

const showview  = (row)=>{
  actionfunc.value = appprops.handlers.update
  _view.value = true
  formdata.value = row
}

const actionfunc = ref()

const emits = defineEmits(['data'])
const handleemit = (name,data)=>{
  emits(name,data)
}
// 分页

const onsortchange = (res)=>{
  let _order = res.order == "descending" ? 'desc': (res.order=='ascending'?'asc':null)
  let _prop = camelToUnderscore(res.prop)
      if(!!_order){
          order.value = {
            field:_prop,order:_order
          }
          handlesearch()
      }else{
        order.value = {
            field:"",order:""
          } 
    }
}
const handlesizechange = (size)=>{
  pager.value.pagesize = size
  pager.value.pagefrom = 0
  handlesearch()
}
const handlecurrentchange = (value)=>{
  pager.value.pagefrom = value
  handlesearch()
}
const pager = ref({pagesize:20,pagefrom:0,})
const order = ref({field:"",method:""})
const tabledata = ref()
const total = ref(0)

const condition = ()=>{
  return {
    pager:pager.value,
    order:order.value,
    conds:condarr.value
  }
}
function arrayToTree(array, options = {}) {
  const {
    idKey = 'id',
    parentKey = 'pid',
    childrenKey = 'children'
  } = options;

  // 创建一个映射表，用于快速查找节点
  const nodeMap = new Map();
  
  // 首先将每个节点添加到映射表中
  array.forEach(item => {
    const newItem = { ...item, [childrenKey]: [] };
    nodeMap.set(item[idKey], newItem);
  });

  // 然后构建树结构
  const roots = [];
  array.forEach(item => {
    const current = nodeMap.get(item[idKey]);
    const parentId = item[parentKey];
    // 如果没有父节点ID或者找不到父节点，则作为根节点
    if (parentId === undefined || parentId === null || !nodeMap.has(parentId)) {
      roots.push(current);
    } else {
      // 否则，将当前节点添加到父节点的children数组中
      const parent = nodeMap.get(parentId);
      parent[childrenKey].push(current);
    }
  });

  return roots;
}
const handlesearch =()=>{
  loading.value = true
  appprops.handlers.search(condition()).then((res)=>{
    let datalist = res.rows||res.data||[]
    if (appprops.tree){
      let {childrenKey,idKey,parentKey} = appprops.tree 
      tabledata.value = arrayToTree(datalist,{childrenKey,idKey,parentKey})
      console.log('tabledata.value',tabledata.value)
      total.value = tabledata.value.length
    }else{
      tabledata.value=datalist
      total.value = res.total
    }
    
  }).finally(()=>{
    loading.value = false
  })
}
const handleexport =(cond)=>{
  appprops.handlers.exportxls(cond).then((res)=>{
      ////console.log('download',res)
      FileSaver.saveAs(res)
  })
}
onMounted(()=>{
  handlesearch()
})


</script>
<style scope>
.header{
  display:flex;
  flex-direction:row;
  justify-content:space-between;
  align-items:center;
}
.header div{
  font-size:14px;
}
</style>
