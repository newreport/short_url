import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import { Plugin as importToCDN } from "vite-plugin-cdn-import";


export default defineConfig({
  plugins: [
    vue(),
    importToCDN({
      modules: [
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
