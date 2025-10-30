<template>
  <section class="app-main" :style="{ minHeight: minHeight }">
    <!-- <transition
    enter-active-class="animate__animated animate__fadeIn"
    mode="out-in"
  >
    <keep-alive :include="cachedViews">
      <router-view ></router-view>
    </keep-alive>
    </transition> -->
    <router-view ></router-view>
  </section>
</template>

<script setup lang="ts">
import { useSettingsStore, useTagsViewStore } from "@/store";
import variables from "@/styles/variables.module.scss";
const cachedViews = computed(() => useTagsViewStore().cachedViews); // 缓存页面集合
const minHeight = computed(() => {
  if (useSettingsStore().tagsView) {
    return `calc(100vh - ${variables["navbar-height"]} - ${variables["tags-view-height"]})`;
  } else {
    return `calc(100vh - ${variables["navbar-height"]})`;
  }
});
</script>

<style lang="scss" scoped>
.app-main {
  position: relative;
  background-color: var(--el-bg-color-page);
}
</style>
<!-- <template #default="{ Component, route }">
  route.fullPath={{ route.fullPath }}
  <transition
    enter-active-class="animate__animated animate__fadeIn"
    mode="out-in"
  >
    <keep-alive :include="cachedViews">
      <component :is="Component" :key="route.fullPath" />
      </keep-alive>
  </transition>
</template> -->
