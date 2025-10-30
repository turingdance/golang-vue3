import type { App } from "vue";
import { createPinia } from "pinia";
import { useDictStore } from "./modules/dict";
import { useConfigStore } from "./modules/config";
import { useStorageStore } from "./modules/storage";

const store = createPinia();

// 全局注册 store
export function setupStore(app: App<Element>) {
  app.use(store);
}

export * from "./modules/app";
export * from "./modules/permission";
export * from "./modules/settings";
export * from "./modules/tagsView";
export * from "./modules/user";
export * from "./modules/storage"
export * from "./modules/config"
export * from "./modules/dict"
export function initStore(){
  const dictStore = useDictStore()
  const configStore = useConfigStore()
  const storageStore = useStorageStore()
  dictStore.init()
  configStore.init()
  storageStore.init()
}
export { store };
