// main.ts
import { createApp } from "vue";
import App from "@/App.vue";
import router from "@/router";
import i18n from "@plugins/i18n";
import * as ElementPlusIconsVue from "@element-plus/icons-vue";
import base64 from '@utils/base64';
import store from '@store/index'
import moment from 'moment'

const app = createApp(App);
app.use(router);
app.use(i18n);
app.use(store)

app.config.globalProperties.$base64=new base64()
app.config.globalProperties.$moment = moment
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component);
}

app.mount("#app");
