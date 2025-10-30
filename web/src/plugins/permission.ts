import router from "@/router";
import { useUserStore, usePermissionStore } from "@/store";
import { showlogin } from "@/utils";
import NProgress from "@/utils/nprogress";
import tokenmgr from "@/utils/tokenmgr";
import { RouteRecordRaw } from "vue-router";

export function setupPermission() {
  // 白名单路由
  const whiteList = ["/login", "/index","/resetpwd"];

  router.beforeEach(async (to, from, next) => {
    NProgress.start();
    const hasToken = tokenmgr.getToken();
    if (hasToken) {
      if (to.path === "/login") {
        // 如果已登录，跳转首页
        next();
        NProgress.done();
      } else {
        const userStore = useUserStore();
        const hasRoles =   userStore.user.roles && userStore.user.roles.length > 0;
        if (hasRoles) {
          // 未匹配到任何路由，跳转404
          if (to.matched.length === 0) {
            from.name ? next({ name: from.name }) : next("/404");
          } else {
            // 如果路由参数中有 title，覆盖路由元信息中的 title
            const title =
              (to.params.title as string) || (to.query.title as string);
            if (title) {
              to.meta.title = title;
            }

            next();
          }
        } else {
          const permissionStore = usePermissionStore();
          try {
            const { authId, roles } = await userStore.getUserInfo();
            const accessRoutes = await permissionStore.generateRoutes(
              authId as number,
              roles as string[]
            );
            console.log('accessRoutes',accessRoutes)
            accessRoutes.forEach((route: RouteRecordRaw) => {
              router.addRoute(route);
            });
            next({ ...to, replace: true });
          } catch (error) {
            console.error(error);
            // 移除 token 并跳转登录页
            await userStore.resetToken();
            // 重定向到登录页，并携带当前页面路由和参数，作为登录成功后跳转的页面
            const params = new URLSearchParams(
              to.query as Record<string, string>
            );
            const queryString = params.toString();
            const redirect = queryString
              ? `${to.path}?${queryString}`
              : to.path;
            showlogin()
            NProgress.done();
          }
        }
      }
    } else {
      // 未登录
      if (whiteList.indexOf(to.path) !== -1) {
        // 在白名单，直接进入
        next();
      } else {
        // 不在白名单，重定向到登录页
        const params = new URLSearchParams(to.query as Record<string, string>);
        const queryString = params.toString();
        const redirect = queryString ? `${to.path}?${queryString}` : to.path;
        //next(`/login?redirect=${encodeURIComponent(redirect)}`);
        showlogin()
        NProgress.done();
      }
    }
  });

  router.afterEach(() => {
    NProgress.done();
  });
}

// 是否有权限
export function hasAuth(
  value: string | string[],
  type: "button" | "role" = "button"
) {
  const { roles, perms } = useUserStore().user;
  //「超级管理员」拥有所有的按钮权限
  if (type === "button" && (roles as string[]).includes("admin")) {
      return true;
  }
  const auths = (type === "button" ? perms : roles) as string[];
  return typeof value === "string"
    ? auths.includes(value)
    : auths.some((perm:string) => {
        return value.includes(perm);
      });
}
