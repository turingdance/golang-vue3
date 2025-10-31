<template>
<img :src="pic"/>
</template>

<script setup>
import {useStorageStore} from "@/store"
import { watch } from "vue";
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
</script>