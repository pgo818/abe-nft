<template>
  <div class="home">
    <!-- 平台选择器 -->
    <div class="platform-switcher">
      <div class="container">
        <div class="text-center text-white mb-4">
          <h1><i class="bi bi-layers me-3"></i>选择管理平台</h1>
          <p class="lead">请选择您要使用的管理平台</p>
        </div>
        <div class="row justify-content-center">
          <div class="col-md-4 mb-4">
            <div class="card platform-card nft-platform h-100" @click="navigateTo('/nft')">
              <div class="card-body text-center">
                <i class="bi bi-collection platform-icon"></i>
                <h3>NFT管理平台</h3>
                <p class="text-muted">管理您的NFT资产，包括铸造、转移、元数据管理等功能</p>
                <div class="mt-3">
                  <span class="badge bg-primary me-2">铸造NFT</span>
                  <span class="badge bg-info me-2">元数据管理</span>
                  <span class="badge bg-warning">子NFT</span>
                </div>
              </div>
            </div>
          </div>
          <div class="col-md-4 mb-4">
            <div class="card platform-card abe-platform h-100" @click="navigateTo('/abe')">
              <div class="card-body text-center">
                <i class="bi bi-shield-lock platform-icon"></i>
                <h3>ABE加密管理</h3>
                <p class="text-muted">基于属性的加密系统，提供细粒度的访问控制和数据保护</p>
                <div class="mt-3">
                  <span class="badge bg-success me-2">属性加密</span>
                  <span class="badge bg-secondary me-2">密钥管理</span>
                  <span class="badge bg-danger">访问控制</span>
                </div>
              </div>
            </div>
          </div>
          <div class="col-md-4 mb-4">
            <div class="card platform-card did-platform h-100" @click="navigateTo('/did')">
              <div class="card-body text-center">
                <i class="bi bi-person-circle platform-icon"></i>
                <h3>DID和VC管理</h3>
                <p class="text-muted">管理您的DID和VC，包括创建、验证、管理等功能</p>
                <div class="mt-3">
                  <span class="badge bg-info me-2">DID创建</span>
                  <span class="badge bg-warning">VC管理</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 连接钱包提示 -->
    <div v-if="!isConnected" class="container mt-5">
      <div class="alert alert-info text-center">
        <h4>请连接MetaMask钱包</h4>
        <p>要使用平台的全部功能，您需要先连接MetaMask钱包</p>
        <button @click="connectWallet" class="btn btn-primary">
          <i class="bi bi-wallet2 me-1"></i>连接钱包
        </button>
      </div>
    </div>
    
    <!-- 系统介绍 -->
    <div class="container mt-5">
      <div class="row">
        <div class="col-md-4">
          <div class="card mb-4">
            <div class="card-body">
              <h5 class="card-title"><i class="bi bi-collection me-2"></i>NFT管理</h5>
              <p class="card-text">NFT管理平台提供了铸造、转移、元数据管理等功能，支持主NFT和子NFT的创建与管理。</p>
              <router-link to="/nft" class="btn btn-outline-primary">进入NFT平台</router-link>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div class="card mb-4">
            <div class="card-body">
              <h5 class="card-title"><i class="bi bi-shield-lock me-2"></i>ABE加密</h5>
              <p class="card-text">基于属性的加密系统，提供细粒度的访问控制，可以对敏感数据进行加密保护。</p>
              <router-link to="/abe" class="btn btn-outline-success">进入ABE平台</router-link>
            </div>
          </div>
        </div>
        <div class="col-md-4">
          <div class="card mb-4">
            <div class="card-body">
              <h5 class="card-title"><i class="bi bi-person-circle me-2"></i>DID和VC</h5>
              <p class="card-text">去中心化身份标识和可验证凭证系统，提供身份管理和凭证颁发、验证功能。</p>
              <router-link to="/did" class="btn btn-outline-warning">进入DID平台</router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 系统状态 -->
    <div class="container mt-4 mb-5">
      <div class="card">
        <div class="card-header">
          <h5><i class="bi bi-info-circle me-2"></i>系统状态</h5>
        </div>
        <div class="card-body">
          <div class="row">
            <div class="col-md-4">
              <div class="d-flex align-items-center mb-3">
                <div class="me-3">
                  <i class="bi bi-wallet2 fs-3"></i>
                </div>
                <div>
                  <h6 class="mb-0">钱包状态</h6>
                  <span :class="isConnected ? 'text-success' : 'text-danger'">
                    {{ isConnected ? '已连接' : '未连接' }}
                  </span>
                </div>
              </div>
            </div>
            <div class="col-md-4">
              <div class="d-flex align-items-center mb-3">
                <div class="me-3">
                  <i class="bi bi-person-badge fs-3"></i>
                </div>
                <div>
                  <h6 class="mb-0">DID状态</h6>
                  <span :class="did ? 'text-success' : 'text-danger'">
                    {{ did ? '已创建' : '未创建' }}
                  </span>
                </div>
              </div>
            </div>
            <div class="col-md-4">
              <div class="d-flex align-items-center mb-3">
                <div class="me-3">
                  <i class="bi bi-hdd-network fs-3"></i>
                </div>
                <div>
                  <h6 class="mb-0">网络状态</h6>
                  <span class="text-success">已连接</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'

export default {
  name: 'Home',
  
  setup() {
    const store = useStore()
    const router = useRouter()
    
    // 从store获取钱包状态
    const isConnected = computed(() => store.state.wallet.isConnected)
    const account = computed(() => store.state.wallet.account)
    const did = computed(() => store.state.wallet.did)
    
    // 导航到指定路由
    const navigateTo = (path) => {
      router.push(path)
    }
    
    // 连接钱包
    const connectWallet = async () => {
      await store.dispatch('wallet/connectWallet')
    }
    
    return {
      isConnected,
      account,
      did,
      navigateTo,
      connectWallet
    }
  }
}
</script>

<style scoped>
.platform-switcher {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 2rem 0;
  margin-bottom: 2rem;
}

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
</style> 