import { useUserStoreHook } from "@/store/modules/user"
import { hasAuth } from "@/plugins/permission";

/**
 * 按钮权限
 */
export const hasPerm = {
  mounted(el: HTMLElement, binding: any) {
    // DOM绑定需要的按钮权限标识
    const { value: requiredPerms } = binding;
    if (requiredPerms) {
      if (!hasAuth(requiredPerms)) {
        el.parentNode && el.parentNode.removeChild(el);
      }
    } else {
      throw new Error(
        "need perms! Like v-has-perm=\"['sys:user:add','sys:user:edit']\""
      );
    }
  },
};

/**
 * 角色权限
 */
export const hasRole = {
  mounted(el: HTMLElement, binding:any) {
    // DOM绑定需要的角色编码
    const { value: requiredRoles } = binding;
    if (requiredRoles) {
      if (!hasAuth(requiredRoles, "role")) {
        el.parentNode && el.parentNode.removeChild(el);
      }
    } else {
      throw new Error("need roles! Like v-has-role=\"['admin','test']\"");
    }
  },
};
/** 权限指令 */
export const permission = {
  mounted(el:any, binding:any) {
    const { value } = binding
    const roles = useUserStoreHook().user.roles as string[]
    if (value && value instanceof Array && value.length > 0) {
      const permissionRoles = value
      const hasPermission = roles.some((role:string) => {
        return permissionRoles.includes(role)
      })
      if (!hasPermission) {
        el.style.display = "none"
      }
    } else {
      throw new Error(`need roles! Like v-permission="['admin','editor']"`)
    }
  }
}
