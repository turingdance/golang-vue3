import {defineAsyncComponent,App} from 'vue';
const components = import.meta.glob('./**/index.vue'); // 异步方式
export default function install(app:App) {
  for (const [key, value] of Object.entries(components)) {
    const name = key.slice(2,key.lastIndexOf('/') );
    app.component(name, defineAsyncComponent(value as any));
  }
}
