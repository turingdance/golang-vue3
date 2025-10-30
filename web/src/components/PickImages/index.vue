<template>
  <div class="file-upload-wraper">
  <div class="image-list">
    <draggable :list="picArr" class="image-list" :disabled="!appprops.editable">
      <template #item="{ element,index }">
        <div class="image-item" :style="styles">
        <img :src="element"   />
         <el-button class="btn-icon" v-if="appprops.editable" @click="handledelete(index)" icon="delete" circle type="danger"></el-button>
        </div>
      </template>
    </draggable>
    <div class="file-btn" v-if="picArr.length<appprops.limit && appprops.editable">
   <slot name="btnupload">
      <div class="cover" :style="styles">
        <el-icon size="32"><Plus/></el-icon>
      </div>
    </slot>
    <input
      :multiple="appprops.multiple"
      class="file-upload-btn"
      type="file"
      :accept="appprops.accept"
      @change="upload($event)"
    />
  </div>
  </div>
   
  </div>
</template>
<script lang="ts" setup>
import { computed, CSSProperties, ref } from "vue"
import uploader from "@/utils/uploader/index"
import { useVModel } from "@vueuse/core"
import draggable from 'vuedraggable'
const appprops = defineProps({
  modelValue: {
    type: [Array],
    default:()=>([])
  },
  multiple: {
    type: Boolean,
    default: true
  },
  accept: {
    type: String,
    default: "image/*"
  },
  width: { type: Number, default: 240 },
  height: { type: Number, default: 200 },
  limit:{type:Number,default:9},
  editable:{type:Boolean,default:true},
})
const emits = defineEmits(["update:modelValue", "data", "response"])
const picArr = useVModel(appprops,"modelValue",emits)

const upload = ($e) => {
  const promisArr = [] as Promise<unknown>[]
  for (let i = 0; i < Math.min(appprops.limit-picArr.value.length,$e.target.files.length); i++) {
    const promis = uploader.upload($e.target.files[i])
    promisArr.push(promis)
  }
  Promise.all(promisArr).then((resarr) => {
    picArr.value.push(...(resarr.map((oo:any)=>oo.url)))
  })
}

const handledelete = (index)=>{
  picArr.value.splice(index,1)
}
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
</script>

<style lang="scss" scope>
.file-btn {
  display: flex;
  position: relative;
  border: 1px dashed rgba($color: #000000, $alpha: 0.2);
  .cover {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 100%;
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
.image-list{
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  flex-wrap: wrap;
  .image-item{
    margin: 5px 5px 5px 5px;
    position: relative;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background-color: rgba($color: #000000, $alpha: 0.1);
    .btn-icon{
      position: absolute;
      right: 5px;;
      top: 5px;
    }
    img{
      max-width:100% ;
      max-height: 100%;
    }
  }
}
</style>
