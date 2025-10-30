<!-- 混合布局菜单(top) -->
<template>
  <div class="left">
    <el-tooltip v-for="(route, index) in mixTopMenus" placement="right" :key="index" :content="route.meta!.title">

    <div
      :class="{'item':true,'active':index==currentIndex}"
      @click="handleMenuSelect(route.path,index)"
    >
      <svg-icon size="24px" v-if="route.meta && route.meta.icon" class="icon"  :icon-class="route.meta!.icon" />
      <!-- <svg-icon size="24px" v-else icon-class="home" class="icon" /> -->
      <span v-else>{{ route.meta!.title }}</span>
    </div>
  </el-tooltip>
  </div>
  <div class="right" v-if="appStore.sidebar.opened">
    <el-menu style="flex:1;background-color: transparent;" :router="true">
      <template v-for="item,index in permissionStore.mixLeftMenus">
        <el-menu-item :index="item.path" class="menu-item"  v-if="item.meta && item.meta.title && !item.meta.hidden">
          <SidebarMenuItemTitle
              style="color:white"
              :icon="item.meta.icon"
              :title="item.meta.title"
            />
        </el-menu-item>
      </template>
      </el-menu>
  </div>
</template>

<script lang="ts" setup>
import { usePermissionStore, useAppStore } from "@/store";
import { computed, ref, onMounted } from "vue";
import { RouteRecordRaw, useRoute, useRouter } from "vue-router";
const appStore = useAppStore();
const permissionStore = usePermissionStore();
const router = useRouter();

// 避免 activeTopMenuPath 缓存被清理，从当前路由路径获取顶部菜单路径，eg. /system/user → /system
const activeTopMenuPath = useRoute().path.replace(/\/[^\/]+$/, "") || "/";
appStore.activeTopMenu(activeTopMenuPath);

// 激活的顶部菜单路径
const activePath = computed(() => appStore.activeTopMenuPath);

// 混合模式顶部菜单集合
const mixTopMenus = ref<RouteRecordRaw[]>([]);

const currentIndex = ref(0)


/**
 * 菜单选择事件
 */
const handleMenuSelect = (routePath: string,index:number) => {
  currentIndex.value = index
  appStore.activeTopMenu(routePath);
  permissionStore.setMixLeftMenus(routePath);
  // 获取左侧菜单集合，默认跳转到第一个菜单
  const mixLeftMenus = permissionStore.mixLeftMenus;
  if(mixLeftMenus.length<1){
      appStore.closeSideBar()
      router.push(routePath)
  }else{
      appStore.openSideBar()
      goToFirstMenu(mixLeftMenus);
  }
};

/**
 * 默认跳转到左侧第一个菜单
 */
const goToFirstMenu = (menus: RouteRecordRaw[]) => {
  if (menus.length === 0) return;
  const [first] = menus;
  if (first.children && first.children.length > 0) {
    goToFirstMenu(first.children as RouteRecordRaw[]);
  } else if (first.name) {
    router.push({
      name: first.name,
    });
  }
};
// 初始化顶部菜单
onMounted(() => {
  mixTopMenus.value = permissionStore.routes.filter(
    (item) => (!!item.meta && !item.meta.hidden)
  )
})
</script>
<style lang="scss" scoped>
$width:32px;
.left {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  height: 100%;
  background-color:transparent;
  .item {
    width: $width;
    height: $width;
    margin: 5px 5px 5px 5px;
    //background-color: rgba(255,255,255,0.2);
    border-radius: 5px;
    text-align: center;
    justify-content: center;
    align-items: center;
    display: flex;
    flex-direction: column;
    &.active{
      background-color: rgba(255,255,255,0.8);  
    }
    &:hover{
      background-color: rgba(255,255,255,0.8);
    }
  }
}
.right{
  height: calc( 100vh - 50px );
  flex:1;
  background-color: rgba(255,255,255,0.6);
  border-radius: 5px;
}
.menu-item{
  &:hover{
    background-color: rgba(255,255,255,0.6);
  }
}
</style>
