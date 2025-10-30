<template>
  <!-- 如果菜单项没有隐藏则显示 -->
  <div v-if="(!item.meta || !item.meta.hidden) ">
    <!-- 显示只有一个子路由或没有子路由的菜单项 -->
    <!-- alwaysShow 这是一个伪命题
    !item.meta?.alwaysShow={{ !item.meta?.alwaysShow }} -->
    <!-- 没有子菜单 -->
    <!-- 有一个需要展示的子菜单 -->
    <template v-if="!item.children">
      <AppLink
        :to="{
          path: resolvePath(item.path),
          query: (item.meta||{}).params,
        }"
      >
        <el-menu-item
          :index="resolvePath(item.path)"
          :class="{ 'submenu-title-noDropdown': false }"
        >
          <SidebarMenuItemTitle
            :icon="(item.meta||{}).icon"
            :title="(item.meta||{}).title"
          />
        </el-menu-item>
      </AppLink>
    </template>

    <template v-else-if="hasOneShowingChild(item.children, item as RouteRecordRaw) && (onlyOneChild && (!onlyOneChild.children || onlyOneChild.noShowingChildren))">
      <AppLink
        v-if="onlyOneChild.meta"
        :to="{
          path: resolvePath(onlyOneChild.path),
          query: onlyOneChild.meta.params,
        }"
      >
        <el-menu-item
          :index="resolvePath(onlyOneChild.path)"
          :class="{ 'submenu-title-noDropdown': !isNest }"
        >
          <SidebarMenuItemTitle
            :icon="onlyOneChild.meta.icon || (item.meta && item.meta.icon)"
            :title="onlyOneChild.meta.title"
          />
        </el-menu-item>
      </AppLink>
    </template>
    
    <!-- 显示具有多个子路由的父菜单项 -->
    <el-sub-menu v-else :index="resolvePath(item.path)" teleported>
      <template #title>
        <SidebarMenuItemTitle
          v-if="item.meta"
          :icon="item.meta && item.meta.icon"
          :title="item.meta && item.meta.title"
        />
      </template>
      <SidebarMenuItem
        v-for="child in item.children"
        :key="child.path"
        :is-nest="child.children && child.children.length > 0"
        :item="child"
        :base-path="resolvePath(child.path)"
      />
    </el-sub-menu>
  </div>
</template>
<script setup lang="ts">
defineOptions({
  name: "SidebarMenuItem",
  inheritAttrs: false,
});

import { isExternal , resolve} from "@/utils/index";
import { RouteRecordRaw } from "vue-router";

const props = defineProps({
  /**
   * 当前路由项对象
   */
  item: {
    type: Object,
    required: true,
  },

  /**
   * 父层级完整路由路径
   */
  basePath: {
    type: String,
    required: true,
  },
  /**
   * 是否为嵌套路由
   */
  isNest: {
    type: Boolean,
    default: false,
  },
});

const onlyOneChild = ref();

/**
 * 判断当前路由是否只有一个显示的子路由
 *
 * @param children 子路由数组
 * @param parent 父级路由对象
 * @returns 布尔值，表示是否只有一个显示的子路由
 */
function hasOneShowingChild(
  children: RouteRecordRaw[] = [],
  parent: RouteRecordRaw
) {
  // 筛选出需要显示的子路由
  const showingChildren = children.filter((route: RouteRecordRaw) => {
    if (!route.meta) {
      return true;
    }
    // 如果指明hidden 则直接过滤掉
    if (typeof route.meta?.hidden=="undefined" ) {
      return true;
    }
    if (!route.meta?.hidden) {
      route.meta!.hidden = false;
      onlyOneChild.value = route;
      return true;
    }else {
      // 否则
      return false;;
    }
  });
  console.log('children',children)
  // 如果只有一个或没有显示的子路由
  console.log('showingChildren',showingChildren)
  return showingChildren.length==1
  if (showingChildren.length == 1) {
    return true;
  }

  // 如果没有子路由，显示父级路由
  if (showingChildren.length === 0) {
    onlyOneChild.value = { ...parent, path: "", noShowingChildren: true };
    return true;
  }
  return false;
}

/**
 * 解析路由路径，将相对路径转换为绝对路径
 *
 * @param routePath 路由路径
 * @returns 绝对路径
 */
function resolvePath(routePath: string) {
  if (isExternal(routePath)) {
    return routePath;
  }
  if (isExternal(props.basePath)) {
    return props.basePath;
  }

  // 完整路径(/system/user) = 父级路径(/system) + 路由路径(user)
  const fullPath = resolve(props.basePath, routePath);
  return fullPath;
}
</script>

<style lang="scss">
.hideSidebar {
  .submenu-title-noDropdown {
    position: relative;
    padding: 0 !important;

    .el-tooltip {
      padding: 0 !important;

      .sub-el-icon {
        margin-left: 19px;
      }
    }

    & > span {
      display: inline-block;
      width: 0;
      height: 0;
      overflow: hidden;
      visibility: hidden;
    }
  }

  .el-sub-menu {
    overflow: hidden;

    & > .el-sub-menu__title {
      padding: 0 !important;

      .sub-el-icon {
        margin-left: 19px;
      }

      .el-sub-menu__icon-arrow {
        display: none;
      }
    }
  }

  .el-menu--collapse {
    width: $sidebar-width-collapsed;

    .el-sub-menu {
      & > .el-sub-menu__title {
        & > span {
          display: inline-block;
          width: 0;
          height: 0;
          overflow: hidden;
          visibility: hidden;
        }
      }
    }
  }
}

.el-menu-item:hover {
  background-color: $menu-hover;
}
</style>
