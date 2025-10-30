<template>
    <el-dialog  v-model="openIt" :title="lesson.title"  :close-on-mask="false">
        <div class="video-wraper" v-if="lesson.media=='video'">
          <video     controls  autoplay :poster="buildUrl(lesson.cover)" :src="buildUrl(lesson.uri)" ></video>
        </div>
        <div class="video-wraper" v-else-if="lesson.media=='image'">
          <el-image :src="buildUrl(lesson.uri)" fit="fit"></el-image>
        </div>
        <div class="video-wraper" v-else-if="lesson.uri.indexOf('http')>-1">
          <el-link :href="lesson.uri" target="_blank" type="primary">系统即将跳转到第三方站点,你确定要跳转吗?</el-link>
        </div>
        <div class="video-wraper" v-else>
          <el-link :href="buildUrl(lesson.uri)" target="_blank" type="primary">点击这里下载资源</el-link>
        </div>
      </el-dialog>
</template>
<script setup>
import { toRefs } from 'vue';
import {buildUrl} from "@/utils/index"
const openIt = defineModel()
const apppops = defineProps(['lesson'])
const {lesson} = toRefs(apppops)

</script>

<style lang="scss" scoped>
.video-wraper{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
  video{
    max-width: 100%;
    height: 650px;
  }
}
</style>