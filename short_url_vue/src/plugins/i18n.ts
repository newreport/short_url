//https://juejin.cn/post/7029609093539037197
import { createI18n } from "vue-i18n";

const messages = {
  en: {
    menu: {
      menu1: "Func",
      menu1_1: "ShortURL",
      menu2: "System",
      menu2_1: "UserSetting"
    },
    message: {
      hello: "hello world",
    },
  },
  ja: {
    menu: {
      menu1: "機能",
      menu1_1: "短链接设置",
      menu2: "系统",
      menu2_1: "用户设置"
    },
    message: {
      hello: "こんにちは、世界",
    },
  },
  zh: {
    menu: {
      menu1: "功能",
      menu1_1: "短链接设置",
      menu2: "系统",
      menu2_1: "用户设置"
    },
    message: {
      hello: "你好，世界",
    },
  },
};
// 通过选项创建 VueI18n 实例
const i18n = createI18n({
  locale: "ja", // 设置地区
  messages, // 设置地区信息
});

export default i18n;
