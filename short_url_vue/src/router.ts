import { createRouter, createWebHistory } from 'vue-router'

import Login from './components/Login.vue'
import Index from './components/Index.vue'
//注册路由
const routes = [
    {
        path: '/',
        name: 'login',
        component: Login
    },
    {
        path: '/index',
        name: 'index',
        component: Index
    }
];
const router = createRouter({
// createWebHashHistory()是使用hash模式路由
// createWebHistory()是使用history模式路由
    history: createWebHistory(),
    routes
});
  

export default router;
