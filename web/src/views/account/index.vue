<script lang="ts" setup>
import { ref, watch } from "vue"
import wwwheader from "./components/wwwhead/index.vue"
import wwwfooter from "./components/wwwfoot/index.vue"
import Register from "./components/register/index.vue"
import Login from "./components/login/index.vue"
import ResetPwd from "./components/resetpwd/index.vue"
import { useRoute } from "vue-router"
const componentName = ref<string>("login")
const route = useRoute()
watch(
  () => route.name,
  (v) => {
    componentName.value = v as string
  },
  {
    immediate: true
  }
)
</script>
<template>
  <div class="page-container">
    <wwwheader style="height: 60px" class="wwwheader" />
    <div class="main">
      <div class="main-container">
        <div class="left">
          <!-- <img class="img" src="../../assets/common/banner1.png" />
          <H1>保护环境</H1>
      <h3>造福人类,构建人类命运共同体</h3> -->
          <div style="width: 40vw">
            <el-carousel indicator-position="outside">
              <el-carousel-item>
                <img src="../../assets/banner/3.png" />
              </el-carousel-item>
              <el-carousel-item>
                <img src="../../assets/banner/4.png" />
              </el-carousel-item>
              <el-carousel-item>
                <img src="../../assets/banner/5.png" />
              </el-carousel-item>
            </el-carousel>
          </div>
        </div>
        <div class="right">
          <Register v-if="componentName == 'register'" />
          <ResetPwd v-else-if="componentName == 'resetpwd'" />
          <Login v-else />
        </div>
      </div>
    </div>
    <wwwfooter style="height: 220px" class="wwwfooter" />
  </div>
</template>

<style lang="scss" scoped>
.page-container {
  background-image: url("@/assets/common/backgroud1.png");
  background-size: cover;
}
.main {
  height: calc(100vh - 280px);
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  width: 100%;
  .main-container {
    width: 80vw;
    display: flex;
    flex-direction: row;
    .left {
      flex: 1;
      .img {
        height: 60vh;
      }
    }
    .right {
      width: 380px;
    }
  }
}
</style>
