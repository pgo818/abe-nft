<template>
  <div class="app-container">
    <!-- 导航栏 -->
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
      <div class="container">
        <router-link class="navbar-brand" to="/">
          <i class="bi bi-shield-lock me-2"></i>
          NFT+ABE集成平台
        </router-link>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
          <ul class="navbar-nav me-auto">
            <li class="nav-item">
              <router-link class="nav-link" to="/">
                <i class="bi bi-house me-1"></i>首页
              </router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link" to="/nft">
                <i class="bi bi-collection me-1"></i>NFT管理
              </router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link" to="/abe">
                <i class="bi bi-shield-lock me-1"></i>ABE加密
              </router-link>
            </li>
            <li class="nav-item">
              <router-link class="nav-link" to="/did">
                <i class="bi bi-person-circle me-1"></i>DID/VC
              </router-link>
            </li>
          </ul>
          <wallet-connector />
        </div>
      </div>
    </nav>

    <!-- 加载指示器 -->
    <div v-if="isLoading" class="loading-overlay">
      <div class="spinner-border text-light" role="status">
        <span class="visually-hidden">加载中...</span>
      </div>
    </div>

    <!-- 主内容区 -->
    <div class="main-content">
      <router-view />
    </div>
    
    <!-- 通知组件 -->
    <notifications />
  </div>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import WalletConnector from '@/components/common/WalletConnector.vue'
import Notifications from '@/components/common/Notifications.vue'

export default {
  name: 'App',
  components: {
    WalletConnector,
    Notifications
  },
  setup() {
    const store = useStore()
    
    // 从store获取加载状态
    const isLoading = computed(() => store.state.app.isLoading)
    
    return {
      isLoading
    }
  }
}
</script>

<style>
.app-container {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

.main-content {
  flex: 1;
  padding: 20px 0;
}

.loading-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 9999;
}

/* 全局样式 */
.platform-card {
  transition: all 0.3s ease;
  cursor: pointer;
  border: none;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
}

.platform-card:hover {
  transform: translateY(-10px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15);
}

.platform-icon {
  font-size: 3rem;
  margin-bottom: 1rem;
}

.nft-platform .platform-icon {
  color: #007bff;
}

.abe-platform .platform-icon {
  color: #28a745;
}

.did-platform .platform-icon {
  color: #fd7e14;
}

.wallet-address-short {
  font-family: 'Courier New', monospace;
  font-size: 0.9em;
}

.did-string {
  font-family: 'Courier New', monospace;
  font-size: 0.85em;
  word-break: break-all;
}
</style> 