<template>
  <div class="header">
    <div class="header-container">
      <div class="left">
        <AuthImage class="logo" v-if="configStore.sitemap['site.logo0']" :src="configStore.sitemap['site.logo0']"></AuthImage>
        <img class="logo" v-else src="@/assets/common/logo.png">
        <H1 class="logo">{{configStore.sitemap['site.title']}}</H1>
        <h3 class="slogo">{{configStore.sitemap['site.memo']}}</h3>
      </div>
      <div class="right">
        <div class="item">首页</div>
        <div class="item">控制台</div>
      </div>
    </div>
  </div>
</template>
<script lang="ts" setup>
import { useConfigStore } from '@/store';
const configStore = useConfigStore()
</script>
<style lang="scss" scoped>
@import "@/styles/mixins.scss";
.header {
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #ffffff;
  .header-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    height: 100%;
    width: 80vw;
    .left {
      display: flex;
      flex-direction: row;
      justify-content: flex-start;
      align-items: center;
      .logo {
        height: 85%;
      }
    }
    .right {
      display: flex;
      justify-content: flex-start;
      flex-direction: row;
      align-items: center;
      .item {
        width: 80px;
        text-align: center;
        cursor: pointer;
      }
    }
  }
}
</style>
