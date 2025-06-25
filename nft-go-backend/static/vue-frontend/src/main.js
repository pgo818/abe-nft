import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

// 引入Bootstrap样式
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap-icons/font/bootstrap-icons.css'
// 引入Bootstrap JS
import * as bootstrap from 'bootstrap/dist/js/bootstrap.bundle.min.js'

// 应用启动时初始化
async function initializeApp() {
    try {
        // 检查钱包连接状态
        console.log('正在检查钱包连接状态...')
        await store.dispatch('wallet/checkWalletConnection')
        console.log('钱包连接状态检查完成')
    } catch (error) {
        console.error('初始化应用失败:', error)
    }
}

// 创建Vue应用
const app = createApp(App)

// 注册全局属性
app.config.globalProperties.$apiBaseUrl = '/api'
app.config.globalProperties.$bootstrap = bootstrap

// 使用路由和状态管理
app.use(router)
app.use(store)

// 等待DOM加载完成后初始化并挂载应用
document.addEventListener('DOMContentLoaded', async () => {
    // 先初始化钱包状态
    await initializeApp()

    // 然后挂载应用
    app.mount('#app')
}) 