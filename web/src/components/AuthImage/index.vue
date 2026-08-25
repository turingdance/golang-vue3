<template>
<img 
:src="pic"
@error="handleImageError"
/>
</template>

<script setup>
import {useStorageStore} from "@/store"
import { watch } from "vue";
import fallbackImage from '@/assets/images/logo.png'
const storageStore = useStorageStore()
const appprops = defineProps({
    src:{
        type:String
    }
})
const pic = ref('data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==')
watch(()=>appprops.src,async (v)=>{
    if(v){
        pic.value = await storageStore.shareUrl(v)
    }
})
onMounted(async ()=>{
    pic.value = await storageStore.shareUrl(appprops.src)
})
const handleImageError = (e) => {
  // 防止死循环：如果已经是兜底图就不再赋值
  if (e.target.src !== fallbackImage) {
    e.target.src = fallbackImage
  }
}
</script>