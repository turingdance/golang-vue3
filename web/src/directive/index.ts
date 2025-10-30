import { type App } from "vue"
import { hasPerm, hasRole } from "./permission";
import { resize } from "./resize"
import { cdn } from "./cdn"
import { pure } from "./pure"
import { xlsx } from "./xlsx"
import {string,phone,email} from "./desensitize"

/** 挂载自定义指令 */
export function setupDirective(app: App) {
  app.directive("hasPerm", hasPerm)
  app.directive("perm", hasPerm)
  app.directive("role", hasRole)
  app.directive("resize", resize)
  app.directive("cdn", cdn)
  app.directive("pure", pure)
  app.directive("xlsx", xlsx)
  app.directive("desensitize-phone",phone)
  app.directive("desensitize-email",email)
  app.directive("desensitize",string)
}
