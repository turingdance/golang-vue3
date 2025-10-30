<template>
    <el-input v-model="filekey" placeholder="请输入">
      <template #append>
        <div class="relative">
      <el-button type="primary">
        上传
      </el-button>
      <input
        class="file-upload-btn"
        type="file"
        :accept="appprops.accept"
        @change="upload($event)"
        />
      </div>  
      </template>
    </el-input>
</template>
<script lang="ts" setup>
import { ref ,onMounted} from "vue"
import uploader from "@/utils/uploader/index"
import { UploadResult } from "@/utils/uploader/types";
const appprops = defineProps({
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

  .file-upload-btn {
    opacity: 0;
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
  }
</style>
