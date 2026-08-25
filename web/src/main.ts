import { createApp } from 'vue'; 
import App from "./App.vue";
import setupPlugins from "@/plugins";
// 本地SVG图标
import "virtual:svg-icons-register";
// 样式
import "element-plus/dist/index.css"
import "element-plus/theme-chalk/dark/css-vars.css";
import "@/styles/index.scss";
import './assets/css/tailwind.css' // 引入包含 Tailwind 指令的 CSS
import "animate.css";
const app = createApp(App);

app.use(setupPlugins);
app.mount("#app");
