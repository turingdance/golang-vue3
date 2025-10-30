<template>
<img :src="dsturl" v-if="image"/>
<video :src="dsturl" v-else-if="media=='video'"></video>
<audio :src="dsturl" v-else-if="media=='audio'"></audio>
<div v-else style="position: relative;">
    <svg-icon :icon-class="'media-'+media" size="64"></svg-icon>
    <div style="position: absolute;bottom: 0;height: 10px;background-color: rgba(0,0,0,0,0.5);left: 0;right: 0" >{{ title }}</div>
</div>
</template>

<script setup>
import {useStorageStore} from "@/store"
const storageStore = useStorageStore()
const extmap = [
  {"media":"image","ext":[".jpg",".png",".jpeg",".bmp",".gif"]},
  {"media":"video","ext":[".mp4",".avi",".mov",".mkv",".wmv"]},
  {"media":"audio","ext":[".mp3",".wav",".ogg",".flac",".aac"]},
  {"media":"docx","ext":[".docx",".doc"]},
  {"media":"pptx","ext":[".pptx",".ppt"]},
  {"media":"pdf","ext":[".pdf"]},
  {"media":"xlsx","ext":[".xlsx",".xls"]}
]
const appprops = defineProps({
    src:{
        type:String
    },
    title:{
        type:String
    }
})
const dsturl = ref('data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==')
watch(()=>appprops.src,async (v)=>{
    if(v){
        dsturl.value = await storageStore.shareUrl(v)
        media.value = findmedia(v)
    }
})
const findmedia = (src)=>{
    let item = extmap.find(oo=>oo.ext.some(ext=>src.lastIndexOf(ext)>-1))
    return item?item.media:'unkown'
}
const media = ref('unkown')
onMounted(async ()=>{
    dsturl.value = await storageStore.shareUrl(appprops.src)
    media.value = findmedia(appprops.src)
})
</script>