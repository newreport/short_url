import { createRouter, createWebHistory } from 'vue-router'
import { useStore } from 'vuex' // 引入useStore 方法


const store = useStore()  // 该方法用于返回store 实例
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
  
router.beforeEach((to, from,next) => {
    if(to.name!='Login'&&store.state.id==-1){
        console.log("鉴权失败")
        // return { name: 'Login' }
    }
  })

export default router;
