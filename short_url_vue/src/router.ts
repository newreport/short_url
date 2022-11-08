import { createRouter, createWebHistory } from 'vue-router'

//注册路由
const routes = [
    {
        path: '/',
        name: 'Login',
        component: ()=>import('@views/Login.vue')
    },
    {
        path: '/index',
        name: 'Index',
        component:()=>import('@views/Index.vue') 
    }
];
const router = createRouter({
// createWebHashHistory()是使用hash模式路由
// createWebHistory()是使用history模式路由
    history: createWebHistory(),
    routes
});
  
// router.beforeEach((to, from) => {
//     if(to.name!='Login')
//     // ...
//     // 返回 false 以取消导航
//     return false
//   })

export default router;
