<template>
  <div class="abe-keygen">
    <h2><i class="bi bi-key me-2"></i>ABE密钥生成</h2>

    <div class="card mt-4">
      <div class="card-header">
        <h5 class="mb-0"><i class="bi bi-plus-circle me-2"></i>生成用户密钥</h5>
      </div>
      <div class="card-body">
        <div v-if="!walletAddress" class="alert alert-warning">
          <i class="bi bi-exclamation-triangle me-2"></i>
          请先连接钱包才能生成用户密钥
        </div>
        
        <div v-else>
          <div class="mb-3">
            <label class="form-label">当前钱包地址</label>
            <input type="text" class="form-control" :value="walletAddress" readonly disabled>
            <div class="form-text">
              <i class="bi bi-info-circle me-1"></i>
              系统将为此钱包地址生成专用的ABE密钥
            </div>
          </div>

          <!-- NFT选择区域 -->
          <div class="mb-3">
            <label class="form-label">选择您的子NFT</label>
            <div v-if="loadingNFTs" class="text-center py-3">
              <div class="spinner-border spinner-border-sm text-primary" role="status">
                <span class="visually-hidden">加载中...</span>
              </div>
              <span class="ms-2">正在加载您的子NFT...</span>
            </div>
            
            <div v-else-if="userNFTs.length === 0" class="alert alert-info">
              <i class="bi bi-info-circle me-2"></i>
              您还没有拥有任何子NFT。请先申请子NFT。
            </div>
            
            <div v-else>
              <div class="row">
                <div v-for="nft in userNFTs" :key="nft.id" class="col-md-6 col-lg-4 mb-3">
                  <div class="card nft-card" :class="{ 'border-primary': selectedNFTs.includes(nft.id) }">
                    <div class="card-body">
                      <div class="form-check">
                        <input 
                          class="form-check-input" 
                          type="checkbox" 
                          :id="`nft-${nft.id}`"
                          :value="nft.id"
                          v-model="selectedNFTs"
                        >
                        <label class="form-check-label w-100" :for="`nft-${nft.id}`">
                          <div class="d-flex align-items-center">
                            <div class="nft-icon me-2">
                              <i class="bi bi-collection text-primary" style="font-size: 1.5rem;"></i>
                            </div>
                            <div class="flex-grow-1">
                              <h6 class="mb-1">{{ nft.name || `NFT #${nft.id}` }}</h6>
                              <p class="mb-1 text-muted small">
                                <strong>主地址:</strong> 
                                <code class="text-monospace">{{ formatAddress(nft.mainAddress) }}</code>
                              </p>
                              <p class="mb-0 text-muted small">
                                <strong>类型:</strong> {{ nft.isChild ? '子NFT' : '主NFT' }}
                                <span v-if="nft.isChild" class="badge bg-secondary ms-1">子NFT</span>
                              </p>
                            </div>
                          </div>
                        </label>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              
              <div class="form-text">
                <i class="bi bi-info-circle me-1"></i>
                选择您要生成密钥的子NFT。密钥将基于该子NFT对应的主NFT拥有者地址生成，格式为 mainNFT:主地址
              </div>
            </div>
          </div>

          <!-- 生成的属性预览 -->
          <div v-if="selectedNFTs.length > 0" class="mb-3">
            <label class="form-label">将生成的用户属性</label>
            <div class="bg-light p-3 rounded">
              <div v-for="attribute in generatedAttributes" :key="attribute" class="mb-1">
                <span class="badge bg-primary me-1">{{ attribute }}</span>
              </div>
            </div>
            <div class="form-text">
              <i class="bi bi-info-circle me-1"></i>
              这些属性基于子NFT对应的主NFT拥有者地址生成，将用于生成您的ABE密钥
            </div>
          </div>

          <button 
            type="button" 
            class="btn btn-primary" 
            @click="generateKey" 
            :disabled="isSubmitting || selectedNFTs.length === 0"
          >
            <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-1" role="status"
              aria-hidden="true"></span>
            <i v-else class="bi bi-key-fill me-1"></i>
            {{ isSubmitting ? '生成中...' : `生成密钥 (${selectedNFTs.length}个子NFT)` }}
          </button>
        </div>
      </div>
    </div>

    <!-- 生成结果 -->
    <div v-if="generatedKey" class="card mt-4">
      <div class="card-header">
        <h5 class="mb-0"><i class="bi bi-check-circle me-2"></i>密钥生成成功</h5>
      </div>
      <div class="card-body">
        <div class="alert alert-success">
          <h6><i class="bi bi-check-circle me-2"></i>密钥已生成！</h6>
          <p class="mb-0"><strong>密钥ID:</strong> {{ generatedKey.user_key_id }}</p>
          <p class="mb-0"><strong>钱包地址:</strong> {{ generatedKey.wallet_address }}</p>
          <p class="mb-0"><strong>包含子NFT数量:</strong> {{ generatedKey.attributes.length }}</p>
        </div>

        <div class="mb-3">
          <label class="form-label"><strong>生成的属性:</strong></label>
          <div class="bg-light p-3 rounded">
            <div v-for="attribute in generatedKey.attributes" :key="attribute" class="mb-1">
              <span class="badge bg-primary me-1">{{ attribute }}</span>
            </div>
          </div>
        </div>

        <div class="mb-3">
          <label class="form-label"><strong>用户密钥:</strong></label>
          <textarea class="form-control" rows="5" readonly :value="generatedKey.attrib_keys"
            @click="$event.target.select()"></textarea>
          <div class="d-flex gap-2 mt-2">
            <button class="btn btn-sm btn-outline-primary" @click="copyToClipboard(generatedKey.attrib_keys)">
              <i class="bi bi-clipboard me-1"></i>复制密钥
            </button>
            <button class="btn btn-sm btn-outline-success" @click="useForDecrypt">
              <i class="bi bi-arrow-right me-1"></i>去解密
            </button>
          </div>
        </div>

        <div class="alert alert-info">
          <i class="bi bi-info-circle me-2"></i>
          <small>请妥善保存此密钥，它将用于解密与您选择的子NFT对应的主NFT相关的加密数据。</small>
        </div>
      </div>
    </div>

    <!-- 密钥历史记录 -->
    <div class="card mt-4">
      <div class="card-header">
        <h5 class="mb-0"><i class="bi bi-list me-2"></i>密钥历史记录</h5>
      </div>
      <div class="card-body">
        <div v-if="loading" class="text-center py-3">
          <div class="spinner-border text-primary" role="status">
            <span class="visually-hidden">加载中...</span>
          </div>
          <p class="mt-2">加载密钥列表...</p>
        </div>
        <div v-else-if="keys.length === 0" class="text-center py-3">
          <i class="bi bi-key" style="font-size: 3rem; color: #dee2e6;"></i>
          <p class="text-muted mt-2">暂无密钥记录</p>
        </div>
        <div v-else>
          <div class="table-responsive">
            <table class="table table-hover">
              <thead>
                <tr>
                  <th>密钥ID</th>
                  <th>钱包地址</th>
                  <th>子NFT数量</th>
                  <th>创建时间</th>
                  <th>操作</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="key in keys" :key="key.id">
                  <td><code>{{ key.id }}</code></td>
                  <td>
                    <span class="text-monospace">{{ formatAddress(key.wallet_address) }}</span>
                  </td>
                  <td>
                    <span class="badge bg-info">{{ key.nft_count }}个子NFT</span>
                  </td>
                  <td>{{ formatDate(key.createdAt) }}</td>
                  <td>
                    <button class="btn btn-sm btn-outline-primary me-1" @click="viewKey(key)">
                      <i class="bi bi-eye"></i>
                    </button>
                    <button class="btn btn-sm btn-outline-secondary" @click="copyToClipboard(key.attrib_keys)">
                      <i class="bi bi-clipboard"></i>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted, watch } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import nftService from '@/services/nftService'

export default {
  name: 'ABEKeyGen',
  setup() {
    const store = useStore()
    const router = useRouter()

    const loading = ref(false)
    const loadingNFTs = ref(false)
    const isSubmitting = ref(false)
    const generatedKey = ref(null)
    const keys = ref([])
    const userNFTs = ref([])
    const selectedNFTs = ref([])

    // 获取钱包地址
    const walletAddress = computed(() => store.getters['wallet/currentAccount'])

    // 获取用户的NFT
    const fetchUserNFTs = async () => {
      if (!walletAddress.value) return

      loadingNFTs.value = true
      try {
        console.log('获取用户NFT:', walletAddress.value)
        
        // 创建要签名的消息
        const message = JSON.stringify({
          action: 'get_my_nfts',
          address: walletAddress.value,
          timestamp: Date.now()
        })

        console.log('准备签名的消息:', message)

        // 临时使用测试模式，跳过签名验证
        const signature = 'dummy'
        const testMessage = 'dummy'
        
        console.log('使用测试模式获取NFT，跳过签名验证')
        
        // 获取用户的NFT
        const response = await nftService.getMyNFTs(walletAddress.value, signature, testMessage)
        console.log('用户NFT响应:', response)
        
        if (response && response.nfts) {
          // 只处理子NFT，过滤掉主NFT
          const childNFTs = response.nfts.filter(nft => 
            nft.isChildNft === true || nft.contractType === 'child'
          )
          
          console.log('过滤后的子NFT列表:', childNFTs)
          
          // 处理子NFT数据，需要获取对应主NFT的拥有者地址
          userNFTs.value = await Promise.all(childNFTs.map(async (nft) => {
            let mainNFTOwnerAddress = walletAddress.value // 默认使用当前钱包地址
            
            // 如果有父NFT的TokenID，尝试获取父NFT的拥有者地址
            if (nft.parentTokenId) {
              try {
                // 这里应该调用API获取父NFT的拥有者信息
                // 暂时使用当前钱包地址，后续可以优化
                console.log('子NFT的父TokenID:', nft.parentTokenId)
                
                // 调用API获取父NFT信息
                const parentNFTResponse = await nftService.getNFTDetails(nft.parentTokenId)
                if (parentNFTResponse && parentNFTResponse.data) {
                  mainNFTOwnerAddress = parentNFTResponse.data.owner
                  console.log('获取到父NFT拥有者地址:', mainNFTOwnerAddress)
                } else {
                  // 如果获取失败，使用当前钱包地址
                  console.warn('无法获取父NFT信息，使用当前钱包地址')
                  mainNFTOwnerAddress = walletAddress.value
                }
              } catch (error) {
                console.error('获取父NFT信息失败:', error)
                mainNFTOwnerAddress = walletAddress.value
              }
            }
            
            return {
              id: nft.tokenId,
              name: nft.metadata?.name || `子NFT #${nft.tokenId}`,
              mainAddress: mainNFTOwnerAddress, // 使用主NFT拥有者的地址
              isChild: true,
              metadata: nft.metadata,
              tokenId: nft.tokenId,
              owner: nft.owner,
              contractType: nft.contractType,
              parentTokenId: nft.parentTokenId
            }
          }))
          
          console.log('处理后的子NFT列表:', userNFTs.value)
        } else {
          userNFTs.value = []
        }
      } catch (error) {
        console.error('获取用户NFT失败:', error)
        store.dispatch('app/showError', '获取NFT列表失败: ' + error.message)
        userNFTs.value = []
      } finally {
        loadingNFTs.value = false
      }
    }

    // 生成的属性
    const generatedAttributes = computed(() => {
      return selectedNFTs.value.map(nftId => {
        const nft = userNFTs.value.find(n => n.id === nftId)
        return nft ? `mainNFT:${nft.mainAddress}` : ''
      }).filter(attr => attr)
    })

    // 监听钱包地址变化
    watch(walletAddress, (newAddress) => {
      if (newAddress) {
        fetchUserNFTs()
      } else {
        userNFTs.value = []
        selectedNFTs.value = []
      }
    }, { immediate: true })

    // 生成密钥
    const generateKey = async () => {
      if (!walletAddress.value) {
        store.dispatch('app/showError', '请先连接钱包')
        return
      }

      if (selectedNFTs.value.length === 0) {
        store.dispatch('app/showError', '请选择至少一个子NFT')
        return
      }

      isSubmitting.value = true
      try {
        console.log('发送密钥生成请求:', { 
          wallet_address: walletAddress.value,
          attributes: generatedAttributes.value
        })

        // 调用store中的密钥生成action
        const result = await store.dispatch('abe/generateKey', {
          wallet_address: walletAddress.value,
          attributes: generatedAttributes.value
        })

        console.log('密钥生成结果:', result)

        generatedKey.value = {
          ...result,
          attributes: generatedAttributes.value
        }

        // 添加到本地列表（用于显示）
        keys.value.unshift({
          id: String(result.user_key_id || Date.now()),
          wallet_address: walletAddress.value,
          attributes: generatedAttributes.value,
          nft_count: selectedNFTs.value.length,
          attrib_keys: result.attrib_keys,
          createdAt: new Date().toISOString()
        })

        // 清空选择
        selectedNFTs.value = []

        // 显示成功消息
        store.dispatch('app/showSuccess', '密钥生成成功！')
      } catch (error) {
        console.error('密钥生成失败:', error)
        store.dispatch('app/showError', '密钥生成失败: ' + error.message)
      } finally {
        isSubmitting.value = false
      }
    }

    // 复制到剪贴板
    const copyToClipboard = async (text) => {
      try {
        await navigator.clipboard.writeText(text)
        store.dispatch('app/showSuccess', '已复制到剪贴板')
      } catch (error) {
        console.error('复制失败:', error)
        store.dispatch('app/showError', '复制失败，请手动复制')
      }
    }

    // 跳转到解密页面
    const useForDecrypt = () => {
      // 将用户密钥存储到store中，供解密页面使用
      store.commit('abe/setLatestUserKey', generatedKey.value.attrib_keys)
      router.push({ name: 'ABEDecrypt' })
    }

    // 查看密钥详情
    const viewKey = (key) => {
      const attributesList = key.attributes.join('\n')
      store.dispatch('app/showInfo', `密钥ID: ${key.id}\n钱包地址: ${key.wallet_address}\n子NFT数量: ${key.nft_count}\n属性:\n${attributesList}`)
    }

    // 格式化地址
    const formatAddress = (address) => {
      if (!address) return ''
      return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`
    }

    // 格式化日期
    const formatDate = (dateString) => {
      if (!dateString) return '未知'
      const date = new Date(dateString)
      return date.toLocaleString()
    }

    // 获取密钥列表
    const fetchKeys = async () => {
      loading.value = true
      try {
        // 这里可以调用API获取历史密钥记录
        // 暂时使用模拟数据
        await new Promise(resolve => setTimeout(resolve, 800))
        
        // 模拟数据
        keys.value = []
      } catch (error) {
        console.error('获取密钥列表失败:', error)
        store.dispatch('app/showError', '获取密钥列表失败: ' + error.message)
      } finally {
        loading.value = false
      }
    }

    // 组件挂载时检查钱包连接状态
    onMounted(() => {
      console.log('ABEKeyGen 组件挂载')
      console.log('钱包连接状态:', walletAddress.value)
      console.log('Store wallet state:', store.state.wallet)
      
      if (!walletAddress.value) {
        store.dispatch('app/showWarning', '请连接钱包以使用ABE密钥生成功能')
      }
      fetchKeys()
    })

    return {
      loading,
      loadingNFTs,
      isSubmitting,
      generatedKey,
      keys,
      userNFTs,
      selectedNFTs,
      walletAddress,
      generatedAttributes,
      fetchUserNFTs,
      generateKey,
      copyToClipboard,
      useForDecrypt,
      viewKey,
      formatAddress,
      formatDate
    }
  }
}
</script>

<style scoped>
.bi {
  vertical-align: middle;
}

.badge {
  font-size: 0.8rem;
}

.text-monospace {
  font-family: 'Courier New', monospace;
  font-size: 0.9rem;
}

input:disabled {
  background-color: #f8f9fa !important;
  opacity: 0.8;
}

code {
  background-color: #f8f9fa;
  padding: 0.2rem 0.4rem;
  border-radius: 0.25rem;
  font-size: 0.875rem;
}

.nft-card {
  transition: all 0.2s ease;
  cursor: pointer;
}

.nft-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.nft-card.border-primary {
  border-width: 2px !important;
}

.form-check-label {
  cursor: pointer;
}

.nft-icon {
  flex-shrink: 0;
}
</style>
