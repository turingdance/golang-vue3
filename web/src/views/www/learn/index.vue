<template>

    <!-- 主内容区 -->
    <main class="container mx-auto px-4 py-8">

      
      <!-- 分类按钮 -->
      <div class="mb-10">
        <div class="flex justify-center gap-4">
          <a 
           href="zhengtianLaser://fromline"
                  size="large"
                  type="primary"
                  class="cursor-pointer bg-primary text-white btn-hover px-8 py-3 rounded-lg font-semibold text-lg flex-1 max-w-xs">
              <i class="fa fa-cut mr-2"></i>光纤激光金属切割机仿真
          </a>
          <a href="zhengtianLaser://fromline" size="large" type="primary" class="cursor-pointer bg-secondary text-white btn-hover px-8 py-3 rounded-lg font-semibold text-lg flex-1 max-w-xs">
            <i class="fa fa-paint-brush mr-2"></i>金属激光雕切机仿真
          </a>
        </div>
      </div>
       <!-- 搜索和筛选区域 -->
       <div class="mb-8 bg-white p-4 rounded-xl shadow-sm">
        <!-- 标签筛选 -->
        <div class="flex flex-wrap gap-2 pointer-cussor use-select-none">
          <el-button  v-for="item,key in sys_lesson_cate" @click="currentCategory = item.value" :type="currentCategory === item.value ? 'primary' : ''" effect="dark" >{{item.label}}</el-button>
        </div>
      </div>

      <!-- 视频列表 -->
      <div class="grid grid-cols-4 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
        <div v-for="lesson,index in filteredLesson" 
        @click="dispatchLesson(lesson)" 
        :key="index" 
        class="bg-white rounded-xl overflow-hidden card-shadow transform transition-all duration-300 hover:-translate-y-1">
          <!-- 视频封面 -->
          <div class="relative">
            <img :src="buildUrl(lesson.cover)" alt="封面" class="w-full h-48 object-cover">
            <div class="absolute bottom-2 right-2 bg-black/70 text-white text-xs px-2 py-1 rounded">
              <span v-if="lesson.media=='video'">视频</span>
              <span v-else-if="lesson.media=='exe'">应用软件</span>
              <span v-else-if="lesson.media=='ppt'">ppt</span>
              <span v-else-if="lesson.media=='url'">外链资源</span>
              <span v-else>{{lesson.media}}</span>
            </div>
          </div>
          
          <!-- 视频信息 -->
          <div class="p-4">
            <h3 class="font-semibold text-lg mb-2 line-clamp-2 h-12">{{ lesson.title }}</h3>
            <div class="flex items-center text-sm text-gray-500">
              <i class="fa fa-eye mr-1"></i> {{ lesson.read }} 次观看
              <span class="mx-2">•</span>
              <i class="fa fa-calendar mr-1"></i> {{ lesson.date }}
            </div>
          </div>
        </div>
      </div>
      
      <!-- 加载更多 -->
      <!-- <div class="mt-10 text-center" v-if="false">
        <button class="btn-hover bg-gray-100 hover:bg-gray-200 text-gray-700 px-6 py-3 rounded-lg font-medium">
          加载更多视频
        </button>
      </div> -->

    </main>
    <Preview v-model="openit" :lesson="lesson" v-if="lesson"></Preview>
  </template>
  <script setup>
import recordApi from "@/api/ecls/record"
import lessonApi from "@/api/ecls/lesson"
import {useDictStore} from "@/store/modules/dict"
import {buildUrl} from "@/utils/index"

const dictStore = useDictStore()

const sys_lesson_cate = dictStore.getDictValue('sys_lesson_cate')

const lesson = ref({})
const openit = ref(false)  
  const dispatchLesson = (v)=>{
    lesson.value = Object.assign({},v)
    openit.value = true
    track(v.id)
  }
const track = (lessonId)=>{
  recordApi.create({
    lessonId
  })
}

const  listall = ()=>{
  lessonApi.search().then(res=>{
    videos.value = res.data ||[]
  })
}
listall()
const currentCategory = ref( 'metal')
const videos = ref([])

const  filteredLesson = computed(()=>{
    return videos.value.filter(video => video.cate === currentCategory.value);
})
</script>
    