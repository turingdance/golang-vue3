<template>
  <div class="file-upload-wraper flex flex-col">
    <div class="flex flex-col">
      <div class="relative">
      <el-button type="primary">
        上传
      </el-button>
      <input
        v-if="(multiple && filekeys.length<limit) || (!multiple && filekeys.length==0)"
        :multiple="appprops.multiple"
        class="file-upload-btn"
        type="file"
        :accept="appprops.accept"
        @change="upload($event)"
        />
      </div>
      <div>
      <slot :filekeys="filekeys" :filelist="filelist" :filekey="filekey">
          <el-table :data="filelist">
            <el-table-column prop="name" :width="250"></el-table-column>
            <el-table-column prop="size" >
                <template #default="scope">
                  {{ formatBytes(scope.row.size) }}
                </template>
            </el-table-column>
            <el-table-column >
                <template #default="scope">
                  <el-button @click="removeitem(scope.$index)" icon="delete" type="danger">删除</el-button>
                </template>
            </el-table-column>
          </el-table>
      </slot>
      </div>
  </div>
  
  </div>
</template>
<script lang="ts" setup>
import { ref ,onMounted} from "vue"
import uploader from "@/utils/uploader/index"
import { UploadResult } from "@/utils/uploader/types";
function  formatBytes(bytes, decimals = 2) {
    if (bytes === 0) return '0 Bytes';

    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];

    const i = Math.floor(Math.log(bytes) / Math.log(k));

    return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
}
const  removeitem = (index)=>{
  
  filekeys.value.splice(index,1)
  filelist.value.splice(index,1)
  if(filekeys.value.length==0){
    filekey.value = ""
  }
  
}
const appprops = defineProps({
  multiple: {
    type: Boolean,
    default: false
  },
  accept: {
    type: String,
    default: "*/*"
  },
  bucket: {
    type: String,
    default: "mnt"
  },
  limit:{
    type:Number,
    default:9,
  },
  width: { type: Number, default: 0 },
  height: { type: Number, default: 0 }
})
const emits = defineEmits(["update:modelValue",'success', "data", "response"])
const filekey = defineModel<string>()
const filekeys = defineModel<string[]>("filekeys",{
  default:()=>([])
})
const filelist = defineModel<UploadResult[]>("filelist",{
  default:()=>([] as UploadResult[])
})
const upload = ($e) => {
  uploader.upload($e.target.files,appprops.bucket).then(rows=>{
      filekey.value = rows[0].key
      filekeys.value = rows.map(pp=>pp.key as string)as string[]
      filelist.value = rows
  })
  
}
onMounted(()=>{
  filelist.value = (filekeys.value||[]).map(key=>{
    return {
      key
    }
  })
})

</script>

<style lang="scss">
.file-upload-wraper {
  display: flex;
  position: relative;
  .cover {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 100%;
    position:relative;
    .action-wraper{
        display:none !important;
        position:absolute;
    }
    &:hover{   
      .action-wraper{
        position:absolute;
        display:flex !important;
        left:0;
        right:0;
        top:0;
        bottom:0;
        flex-direction:column;
        justify-content:space-around;
        align-items:center;
        background-color:rgba(0,0,0,0.4);
      }
    }
    img {
      max-width: 100%;
      max-height: 100%;
    }
  }
  .file-upload-btn {
    opacity: 0;
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
  }
}
</style>
