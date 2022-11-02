//https://juejin.cn/post/7029609093539037197
import { createI18n } from "vue-i18n";

const messages = {
    en: {
      message: {
        hello: 'hello world'
      }
    },
    ja: {
      message: {
        hello: 'こんにちは、世界'
      }
    },
    zh:{
        message: {
            hello: '你好，世界'
          } 
    }
  }
  // 通过选项创建 VueI18n 实例
const i18n =   createI18n({
    locale:"ja", // 设置地区
    messages, // 设置地区信息
  })
  
  
  export default i18n