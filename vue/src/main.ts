// main.ts
import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

<<<<<<< HEAD
const app= createApp(App)
app.mount('#app')


=======
const app = createApp(App)

app.use(ElementPlus)



app.mount('#app')
>>>>>>> main
