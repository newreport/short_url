import { createRouter, createWebHistory } from "vue-router";
import myStore from "@store/index"; // 引入useStore 方法
import base64 from "@utils/base64";

const store = myStore;
//注册路由
const routes = [
  {
    path: "/",
    name: "Login",
    component: () => import("@views/Login.vue"),
  },
  {
    path: "/index",
    name: "Index",
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
  console.log(store);
  if (store.state.token.id == -1) {
    //F5进入页面
    //重新读取token
    let base = new base64();
    let refreshToken = base.decode(localStorage.getItem("refreshToken"));
    let accessToken = base.decode(localStorage.getItem("accessToken"));
    if (refreshToken) {
      store.commit("refreshToken", refreshToken);
      if (accessToken) {
        store.commit("accessToken", accessToken);
      }
    } else {
      store.commit("cleanToken");
      return { name: "Login" };
    }
    if(to.name == 'Login'){//登录了但是又返回登录页面
      return { name: "Index" };
    }
  }
});

export default router;
