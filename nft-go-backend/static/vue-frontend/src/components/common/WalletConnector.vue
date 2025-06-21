<template>
  <div class="wallet-connector">
    <button v-if="!isConnected" @click="connectWallet" class="btn btn-outline-dark">
      <i class="bi bi-wallet2 me-1"></i>连接钱包
    </button>
    <div v-else class="d-flex align-items-center">
      <div class="wallet-info me-3">
        <div class="wallet-address" :title="account">{{ shortAccount }}</div>
        <div v-if="did" class="did-info small text-muted">DID: {{ shortDid }}</div>
      </div>
      <button @click="disconnectWallet" class="btn btn-outline-dark btn-sm">
        <i class="bi bi-box-arrow-right me-1"></i>断开
      </button>
    </div>
    
    <!-- 钱包连接提示模态框 -->
    <div class="modal fade" ref="connectModal" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">连接钱包</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <p class="text-center">请连接MetaMask钱包以使用平台功能</p>
            <div class="text-center">
              <button @click="connectWallet" class="btn btn-primary">
                <i class="bi bi-wallet2 me-1"></i>连接MetaMask
              </button>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">取消</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { computed, ref, onMounted, watch } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { Modal } from 'bootstrap'

export default {
  name: 'WalletConnector',
  
  setup() {
    const store = useStore()
    const router = useRouter()
    const connectModal = ref(null)
    let bsModal = null
    
    // 从store获取钱包状态
    const isConnected = computed(() => store.state.wallet.isConnected)
    const account = computed(() => store.state.wallet.account)
    const did = computed(() => store.state.wallet.did)
    const showConnectPrompt = computed(() => store.state.wallet.showConnectPrompt)
    const redirectRoute = computed(() => store.state.wallet.redirectRoute)
    
    // 格式化地址和DID
    const shortAccount = computed(() => {
      if (!account.value) return ''
      return `${account.value.substring(0, 6)}...${account.value.substring(account.value.length - 4)}`
    })
    
    const shortDid = computed(() => {
      if (!did.value) return ''
      if (did.value.length <= 20) return did.value
      return `${did.value.substring(0, 10)}...${did.value.substring(did.value.length - 10)}`
    })
    
    // 连接钱包
    const connectWallet = async () => {
      const result = await store.dispatch('wallet/connectWallet')
      
      if (result) {
        // 如果有重定向路由，连接成功后跳转
        if (redirectRoute.value) {
          router.push(redirectRoute.value)
          store.commit('wallet/clearRedirectRoute')
        }
        
        // 隐藏模态框
        if (bsModal) {
          bsModal.hide()
        }
      }
    }
    
    // 断开钱包连接
    const disconnectWallet = () => {
      store.dispatch('wallet/disconnectWallet')
      // 重定向到首页
      if (router.currentRoute.value.meta.requiresWallet) {
        router.push('/')
      }
    }
    
    // 监听钱包连接提示状态
    watch(showConnectPrompt, (show) => {
      if (show && bsModal) {
        bsModal.show()
        // 显示后重置状态
        store.commit('wallet/setShowConnectPrompt', false)
      }
    })
    
    // 组件挂载后检查钱包连接状态和初始化模态框
    onMounted(async () => {
      // 检查钱包连接状态
      await store.dispatch('wallet/checkWalletConnection')
      
      // 初始化模态框
      if (connectModal.value) {
        bsModal = new Modal(connectModal.value)
        
        // 如果需要显示连接提示，则显示模态框
        if (showConnectPrompt.value) {
          bsModal.show()
          store.commit('wallet/setShowConnectPrompt', false)
        }
      }
      
      // 设置MetaMask事件监听器
      if (window.ethereum) {
        window.ethereum.on('accountsChanged', handleAccountsChanged)
        window.ethereum.on('chainChanged', () => {
          // 网络变化时重新加载页面
          window.location.reload()
        })
      }
    })
    
    // 处理账户变更
    const handleAccountsChanged = async (accounts) => {
      if (accounts.length === 0) {
        // 用户断开了所有账户
        store.dispatch('wallet/disconnectWallet')
        if (router.currentRoute.value.meta.requiresWallet) {
          router.push('/')
        }
      } else if (accounts[0] !== account.value) {
        // 用户切换了账户
        await store.dispatch('wallet/connectWallet')
      }
    }
    
    return {
      isConnected,
      account,
      did,
      shortAccount,
      shortDid,
      connectWallet,
      disconnectWallet,
      connectModal
    }
  }
}
</script>

<style scoped>
.wallet-connector {
  display: flex;
  align-items: center;
}

.wallet-info {
  text-align: right;
}

.wallet-address {
  font-family: 'Courier New', monospace;
  font-size: 0.9rem;
  font-weight: 500;
}

.did-info {
  font-family: 'Courier New', monospace;
  word-break: break-all;
}
</style> 