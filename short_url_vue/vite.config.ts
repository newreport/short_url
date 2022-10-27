import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { Plugin as importToCDN } from "vite-plugin-cdn-import";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";

export default defineConfig({
  plugins: [
    vue(),
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
    importToCDN({//CDN只会在build时生效，yarn dev不生效
      modules: [
        {
          name:"axios",
          var: "Axios",
          path:"https://unpkg.com/axios/dist/axios.min.js"
        },
        {
          name: "vue",
          var: "Vue",
          path: "//unpkg.com/vue@next",
        },
        {
          name: "vuex",
          var: "Vuex",
          path: "//unpkg.com/vuex@next",
        },
        {
          name: "vue-class-component",
          var: "VueClassComponent",
          path: "//unpkg.com/vue-class-component@next",
        },
        {
          name: "element-plus",
          var: "ElementPlus",
          path: "//unpkg.com/element-plus",
          css: "//unpkg.com/element-plus/dist/index.css",
        },
      ],
    }),
  ],
});
