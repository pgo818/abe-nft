import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// 引入Bootstrap样式
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap-icons/font/bootstrap-icons.css'
// 引入Bootstrap JS
import * as bootstrap from 'bootstrap/dist/js/bootstrap.bundle.min.js'

// 创建Vue应用
const app = createApp(App)

// 注册全局属性
app.config.globalProperties.$apiBaseUrl = '/api'
app.config.globalProperties.$bootstrap = bootstrap

// 使用路由和状态管理
app.use(router)
app.use(store)

// 挂载应用
app.mount('#app') 