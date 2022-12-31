import { createRouter, createWebHistory } from "vue-router";
import myStore from "@store/index"; // 引入useStore 方法
import base64 from "@utils/base64";

const store = myStore;
//注册路由
const routes = [
  {
    path: "/",
    name: 'Login',
    component: () => import("@views/Login.vue"),
  },
  {
    path: "/index",
    name: 'Index',
    component: () => import("@views/Index.vue"),
  },
];
const router = createRouter({
  // createWebHashHistory()是使用hash模式路由
  // createWebHistory()是使用history模式路由
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from) => {
  if (store.state.token.id == -1) {
    //F5进入页面
    //重新读取token
    let base = new base64();
    let refreshToken = localStorage.getItem("refreshToken");
    let accessToken = localStorage.getItem("accessToken");
    if (refreshToken) {
      refreshToken = base.decode(refreshToken);
      store.commit("refreshToken", refreshToken);
      if (accessToken) {
        accessToken = base.decode(accessToken);
        store.commit("accessToken", accessToken);
      }
    } else {
      store.commit("cleanToken");
    }
    if (to.name == "Login"&&store.state.token.id >0) {
      //登录了但是又返回登录页面
      return { name: "Index" };
    }else if (to.name != "Login"&&store.state.token.id == -1){
      console.log("去Login")
      return {name:"Login"}
    }
  }
});

export default router;
