import { createRouter, createWebHistory } from 'vue-router'

//注册路由
const routes = [
    {
        path: '/',
        name: 'login',
        component: () => import('../components/Login.vue')
    },
    {
        path: '/index',
        name: 'index',
        component: () => import('../components/index/Index.vue')
    }
];
const router = createRouter({
// createWebHashHistory()是使用hash模式路由
// createWebHistory()是使用history模式路由
    history: createWebHistory(),
    routes
});

export default router;
