// main.ts
import { createApp } from "vue";
import App from "@/App.vue";
import router from "@/router";
import i18n from "@/plugins/i18n";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";

const app = createApp(App);
app.use(router);
app.use(i18n);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}
console.log(123);
app.mount("#app");
