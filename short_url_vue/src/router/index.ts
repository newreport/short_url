import { createRouter, createWebHistory } from 'vue-router'

//注册路由
const routes = [
    {
        path: '/',
        name: 'login',
        component: () => import('../components/login.vue')
    },
    {
        path: '/index',
        name: 'index',
        component: () => import('../components/index.vue')
    }
];
const router = createRouter({
// createWebHashHistory()是使用hash模式路由
// createWebHistory()是使用history模式路由
    history: createWebHistory(),
    routes
});
router.beforeEach(async (to, from) => {
      // 将用户重定向到登录页面
      return { name: 'index' }
  })
  

export default router;
