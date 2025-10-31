<template>
  <div class="file-upload-wraper">
    <slot :filekeys="filekeys" :filelist="filelist" :filekey="filekey">
      <div class="cover" :style="styles" v-if="!filekey">
            <el-icon size="32"><Plus/></el-icon>
      </div>
      <div class="cover" :style="styles" v-else>
        <AuthImage :src="filekey"></AuthImage>
        <div class="action-wraper">
          <el-button class="btn-icon" @click="handledelete(0)" icon="delete" circle type="danger"></el-button>
        </div>
      </div>
  </slot>
    <input
      v-if="!filekey"
      class="file-upload-btn"
      type="file"
      :accept="appprops.accept"
      @change="upload($event)"
    />
  </div>
</template>
<script lang="ts" setup>
import { computed, CSSProperties, ref } from "vue"
import uploader from "@/utils/uploader/index"
import { UploadResult } from "@/utils/uploader/types";
const appprops = defineProps({
  multiple: {
    type: Boolean,
    default: false
  },
  accept: {
    type: String,
    default: "image/*"
  },
  bucket: {
    type: String,
    default: "mnt"
  },
  limit:{
    type:Number,
    default:1,
  },
  width: { type: Number, default: 0 },
  height: { type: Number, default: 0 }
})
const emits = defineEmits(["update:modelValue",'success', "data", "response"])
const filekey = defineModel<string>()
const filekeys = defineModel<string[]>("filekeys",{
  default:()=>([])
})

const filelist = ref<UploadResult[]>([])
const upload = ($e) => {
  uploader.upload($e.target.files,appprops.bucket).then(rows=>{
      filekey.value = rows[0].key
      filekeys.value = rows.map(pp=>pp.key as string)as string[]
      filelist.value = rows || []
  })
  
}
onMounted(()=>{
  filelist.value = (filekeys.value||[]).map(key=>{
    return {
      key
    }
  })
  if(filekey.value){
    filekeys.value.push(filekey.value)
  }
  
})
const styles = computed(() => {
  const obj = {} as CSSProperties
  if (appprops.width > 0) {
    obj.width = appprops.width + "px"
  }
  if (appprops.height > 0) {
    obj.height = appprops.height + "px"
  }
  return obj
})
const handledelete = (index)=>{
  filekeys.value.splice(index,1)
  filelist.value.splice(index,1)
  if(filekeys.value.length==0){
    filekey.value = ""
  }
}
</script>

<style lang="scss">
.file-upload-wraper {
  display: flex;
  position: relative;
  border: 1px dashed var(--el-border-color-darker);
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
