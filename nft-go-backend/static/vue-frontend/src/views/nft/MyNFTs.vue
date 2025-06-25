<template>
  <div class="my-nfts">
    <h2><i class="bi bi-wallet me-2"></i>我的NFT</h2>

    <div class="alert alert-info" v-if="!isConnected">
      <i class="bi bi-info-circle me-2"></i>请先连接钱包以查看您的NFT
    </div>

    <!-- 加载状态 -->
    <div v-else-if="isLoading" class="text-center my-5">
      <div class="spinner-border text-primary" role="status">
        <span class="visually-hidden">加载中...</span>
      </div>
      <p class="mt-2">加载NFT中...</p>
    </div>

    <!-- 没有NFT时显示 -->
    <div v-else-if="!nfts.length" class="text-center my-5">
      <i class="bi bi-collection fs-1 text-muted"></i>
      <p class="lead mt-3">您还没有NFT</p>
      <router-link :to="{ name: 'MintNFT' }" class="btn btn-primary mt-2">
        <i class="bi bi-plus-circle me-1"></i>铸造NFT
      </router-link>
    </div>

    <!-- NFT列表 -->
    <div v-else class="row">
      <div v-for="nft in nfts" :key="nft.tokenId" class="col-md-4 mb-4">
        <div class="card nft-card h-100">
          <img :src="getNFTImage(nft)" class="card-img-top" :alt="getNFTName(nft)" @error="handleImageError">
          <div class="card-body">
            <!-- 元数据加载状态 -->
            <div v-if="!nft.metadata && nft.uri" class="text-center mb-2">
              <div class="spinner-border spinner-border-sm text-primary" role="status">
                <span class="visually-hidden">加载中...</span>
              </div>
              <small class="text-muted ms-1">加载元数据中...</small>
            </div>

            <h5 class="card-title">
              {{ getNFTName(nft) }}
            </h5>
            <p class="card-text">{{ getNFTDescription(nft) }}</p>

            <!-- 元数据属性展示 -->
            <div v-if="hasAttributes(nft)" class="mt-3">
              <h6 class="border-bottom pb-2">属性</h6>
              <div class="row g-2 mt-2">
                <div v-for="(attr, index) in getAttributes(nft)" :key="index" class="col-6">
                  <div class="attribute-box p-2 border rounded">
                    <div class="text-muted small">{{ attr.trait_type }}</div>
                    <!-- 特殊处理密文显示 -->
                    <div v-if="attr.trait_type === 'Encrypted_ciphertext'" class="fw-bold">
                      <div class="ciphertext-container">
                        <code class="small text-wrap">{{ truncateText(attr.value, 30) }}</code>
                        <button v-if="attr.value && attr.value.length > 30" 
                                class="btn btn-sm btn-outline-secondary mt-1" 
                                @click="viewNFTDetails(nft)">
                          <i class="bi bi-eye me-1"></i>查看完整密文
                        </button>
                      </div>
                    </div>
                    <!-- 其他属性正常显示 -->
                    <div v-else class="fw-bold">{{ attr.value }}</div>
                  </div>
                </div>
              </div>
            </div>

            <div class="mt-3">
              <p class="card-text"><small class="text-muted">Token ID: {{ nft.tokenId }}</small></p>
              <p class="card-text"><small class="text-muted">Owner: {{ formatAddress(nft.owner) }}</small></p>
            </div>

            <div class="d-flex flex-wrap gap-2 mt-3">
              <button class="btn btn-outline-primary btn-sm" @click="viewNFTDetails(nft)">
                <i class="bi bi-info-circle me-1"></i>详情
              </button>
              <!-- 只有主NFT才能创建子NFT -->
              <button v-if="!isChildNFT(nft)" class="btn btn-outline-success btn-sm" @click="createChildNFT(nft.tokenId)">
                <i class="bi bi-plus-circle me-1"></i>创建子NFT
              </button>
              <!-- 只有主NFT才能更新元数据 -->
              <button v-if="!isChildNFT(nft)" class="btn btn-outline-warning btn-sm" @click="updateMetadata(nft)">
                <i class="bi bi-pencil-square me-1"></i>更新元数据
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- NFT详情模态框 -->
    <div class="modal fade" ref="nftDetailsModal" tabindex="-1">
      <div class="modal-dialog modal-lg">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">NFT详情</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body" v-if="selectedNFT">
            <div class="row">
              <div class="col-md-5">
                <img :src="getNFTImage(selectedNFT)" class="img-fluid rounded" :alt="getNFTName(selectedNFT)"
                  @error="handleImageError">
              </div>
              <div class="col-md-7">
                <h4>{{ getNFTName(selectedNFT) }}</h4>
                <p>{{ getNFTDescription(selectedNFT) }}</p>

                <div class="mt-3">
                  <h6>基本信息</h6>
                  <table class="table table-sm">
                    <tbody>
                      <tr>
                        <th>Token ID</th>
                        <td>{{ selectedNFT.tokenId }}</td>
                      </tr>
                      <tr>
                        <th>所有者</th>
                        <td>{{ selectedNFT.owner }}</td>
                      </tr>
                      <tr>
                        <th>URI</th>
                        <td class="text-break">{{ selectedNFT.uri }}</td>
                      </tr>
                      <tr v-if="selectedNFT.isChildNft">
                        <th>父NFT ID</th>
                        <td>{{ selectedNFT.parentTokenId }}</td>
                      </tr>
                    </tbody>
                  </table>
                </div>

                <!-- 元数据属性展示 -->
                <div v-if="getAttributes(selectedNFT) && getAttributes(selectedNFT).length > 0" class="mt-3">
                  <h6>属性</h6>
                  <div class="row g-2">
                    <div v-for="(attr, index) in getAttributes(selectedNFT)" :key="index" class="col-12">
                      <div class="attribute-box p-3 border rounded">
                        <div class="text-muted small mb-1">{{ attr.trait_type }}</div>
                        <!-- 特殊处理密文显示 -->
                        <div v-if="attr.trait_type === 'Encrypted_ciphertext'" class="fw-bold">
                          <div class="ciphertext-detail">
                            <div class="d-flex justify-content-between align-items-start mb-2">
                              <span class="badge bg-secondary">密文 ({{ attr.value?.length || 0 }} 字符)</span>
                              <button class="btn btn-sm btn-outline-primary" 
                                      @click="copyToClipboard(attr.value)"
                                      title="复制密文">
                                <i class="bi bi-clipboard"></i>
                              </button>
                            </div>
                            <div class="ciphertext-content p-2 border rounded bg-light" style="max-height: 200px; overflow-y: auto;">
                              <code class="small text-wrap d-block">{{ attr.value }}</code>
                            </div>
                          </div>
                        </div>
                        <!-- 策略属性特殊显示 -->
                        <div v-else-if="attr.trait_type === 'Policy'" class="fw-bold">
                          <div class="policy-display">
                            <span class="badge bg-info mb-2">访问策略</span>
                            <div class="policy-content p-2 border rounded bg-light">
                              <code class="text-primary">{{ attr.value }}</code>
                            </div>
                          </div>
                        </div>
                        <!-- 其他属性正常显示 -->
                        <div v-else class="fw-bold">{{ attr.value }}</div>
                      </div>
                    </div>
                  </div>
                </div>


              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">关闭</button>
            <!-- 只有主NFT才能更新元数据 -->
            <button v-if="selectedNFT && !isChildNFT(selectedNFT)" 
                    type="button" 
                    class="btn btn-warning" 
                    @click="updateMetadata(selectedNFT)">
              <i class="bi bi-pencil-square me-1"></i>更新元数据
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- 更新URI模态框 -->
    <div class="modal fade" ref="updateUriModal" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">更新NFT元数据URI</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <form @submit.prevent="submitUriUpdate">
            <div class="modal-body">
              <div class="mb-3">
                <label for="current-uri" class="form-label">当前URI</label>
                <input type="text" class="form-control" id="current-uri" :value="updateForm.currentUri" readonly>
              </div>
              <div class="mb-3">
                <label for="new-uri" class="form-label">新URI</label>
                <input type="text" class="form-control" id="new-uri" v-model="updateForm.newUri" 
                       placeholder="输入新的NFT元数据URI..." required>
                <div class="form-text">请输入有效的IPFS URI或HTTP URL</div>
              </div>
              <div class="mb-3">
                <label for="token-id" class="form-label">Token ID</label>
                <input type="text" class="form-control" id="token-id" :value="updateForm.tokenId" readonly>
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">取消</button>
              <button type="submit" class="btn btn-warning" :disabled="isUpdatingUri">
                <span v-if="isUpdatingUri" class="spinner-border spinner-border-sm me-1" role="status" aria-hidden="true"></span>
                <i v-else class="bi bi-pencil-square me-1"></i>
                {{ isUpdatingUri ? '更新中...' : '更新URI' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import { Modal } from 'bootstrap'
import { formatAddress } from '@/utils/api'

export default {
  name: 'MyNFTs',

  setup() {
    const store = useStore()
    const router = useRouter()
    const nftDetailsModal = ref(null)
    const updateUriModal = ref(null)
    let bsDetailsModal = null
    let bsUpdateUriModal = null

    // 从store获取数据
    const nfts = computed(() => store.state.nft.myNFTs)
    const isLoading = computed(() => store.state.nft.isLoading)
    const isConnected = computed(() => store.state.wallet.isConnected)

    // 选中的NFT
    const selectedNFT = ref(null)

    // 更新URI表单
    const updateForm = ref({
      tokenId: '',
      currentUri: '',
      newUri: ''
    })

    // 更新状态
    const isUpdatingUri = ref(false)

    // 加载NFT列表
    const loadNFTs = async () => {
      if (isConnected.value) {
        await store.dispatch('nft/loadMyNFTs')
      }
    }

    // 处理图片加载错误
    const handleImageError = (event) => {
      event.target.src = 'data:image/svg+xml;charset=UTF-8,%3Csvg%20width%3D%22200%22%20height%3D%22200%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22%23f0f0f0%22%2F%3E%3Ctext%20x%3D%2250%25%22%20y%3D%2250%25%22%20font-family%3D%22Arial%22%20font-size%3D%2224%22%20text-anchor%3D%22middle%22%20dominant-baseline%3D%22middle%22%20fill%3D%22%23999%22%3ENFT%20%E5%8D%A0%E4%BD%8D%E5%9B%BE%3C%2Ftext%3E%3C%2Fsvg%3E'
      event.target.onerror = null
    }

    // 获取NFT图像
    const getNFTImage = (nft) => {
      console.log('获取NFT图像:', nft.tokenId, nft.uri, nft.metadata)

      // 如果有元数据且有图像URL
      if (nft.metadata && nft.metadata.image) {
        console.log('使用元数据中的图像:', nft.metadata.image)

        // 处理IPFS链接
        if (nft.metadata.image.startsWith('ipfs://')) {
          const ipfsUrl = nft.metadata.image.replace('ipfs://', 'https://ipfs.io/ipfs/')
          console.log('转换IPFS链接为HTTP:', ipfsUrl)
          return ipfsUrl
        }

        return nft.metadata.image
      }

      // 如果URI是IPFS链接，转换为HTTP链接
      if (nft.uri && nft.uri.startsWith('ipfs://')) {
        const ipfsUrl = nft.uri.replace('ipfs://', 'https://ipfs.io/ipfs/')
        console.log('使用转换后的URI作为图像:', ipfsUrl)
        return ipfsUrl
      }

      // 如果URI是JSON字符串，尝试提取图像URL
      if (nft.uri && nft.uri.trim().startsWith('{') && nft.uri.trim().endsWith('}')) {
        try {
          const uriData = JSON.parse(nft.uri)
          if (uriData.image) {
            console.log('从URI JSON中提取图像:', uriData.image)

            // 处理IPFS链接
            if (uriData.image.startsWith('ipfs://')) {
              const ipfsUrl = uriData.image.replace('ipfs://', 'https://ipfs.io/ipfs/')
              console.log('转换IPFS链接为HTTP:', ipfsUrl)
              return ipfsUrl
            }

            return uriData.image
          }
        } catch (error) {
          console.error('解析URI JSON失败:', error)
        }
      }

      // 回退到URI或占位图
      console.log('使用原始URI作为图像:', nft.uri)
      return nft.uri || 'data:image/svg+xml;charset=UTF-8,%3Csvg%20width%3D%22200%22%20height%3D%22200%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22%23f0f0f0%22%2F%3E%3Ctext%20x%3D%2250%25%22%20y%3D%2250%25%22%20font-family%3D%22Arial%22%20font-size%3D%2224%22%20text-anchor%3D%22middle%22%20dominant-baseline%3D%22middle%22%20fill%3D%22%23999%22%3ENFT%20%E5%8D%A0%E4%BD%8D%E5%9B%BE%3C%2Ftext%3E%3C%2Fsvg%3E'
    }

    // 获取NFT名称
    const getNFTName = (nft) => {
      console.log('获取NFT名称:', nft.tokenId, nft.uri)

      // 如果有元数据且有名称
      if (nft.metadata && nft.metadata.name) {
        console.log('使用元数据中的名称:', nft.metadata.name)
        return nft.metadata.name
      }

      // 如果URI是JSON字符串，尝试提取名称
      if (nft.uri && nft.uri.trim().startsWith('{') && nft.uri.trim().endsWith('}')) {
        try {
          const uriData = JSON.parse(nft.uri)
          if (uriData.name) {
            console.log('从URI JSON中提取名称:', uriData.name)
            return uriData.name
          }
        } catch (error) {
          console.error('解析URI JSON失败:', error)
        }
      }

      // 如果URI是IPFS链接，提取哈希作为名称的一部分
      if (nft.uri && nft.uri.startsWith('ipfs://')) {
        const ipfsHash = nft.uri.replace('ipfs://', '')
        console.log('使用IPFS哈希作为名称的一部分:', ipfsHash)
        return `NFT #${nft.tokenId} (${ipfsHash.substring(0, 8)}...)`
      }

      // 回退到默认名称
      return 'NFT #' + nft.tokenId
    }

    // 获取NFT描述
    const getNFTDescription = (nft) => {
      console.log('获取NFT描述:', nft.tokenId, nft.uri)

      // 如果有元数据且有描述
      if (nft.metadata && nft.metadata.description) {
        console.log('使用元数据中的描述:', nft.metadata.description)
        return nft.metadata.description
      }

      // 如果URI是JSON字符串，尝试提取描述
      if (nft.uri && nft.uri.trim().startsWith('{') && nft.uri.trim().endsWith('}')) {
        try {
          const uriData = JSON.parse(nft.uri)
          if (uriData.description) {
            console.log('从URI JSON中提取描述:', uriData.description)
            return uriData.description
          }
        } catch (error) {
          console.error('解析URI JSON失败:', error)
        }
      }

      // 回退到默认描述
      return 'This is an NFT created on our platform'
    }

    // 获取NFT属性
    const getAttributes = (nft) => {
      console.log('获取NFT属性:', nft.tokenId, nft.uri)

      // 如果有元数据且有属性
      if (nft.metadata && nft.metadata.attributes && Array.isArray(nft.metadata.attributes)) {
        console.log('使用元数据中的属性:', nft.metadata.attributes)
        return nft.metadata.attributes
      }

      // 如果URI是JSON字符串，尝试提取属性
      if (nft.uri && nft.uri.trim().startsWith('{') && nft.uri.trim().endsWith('}')) {
        try {
          const uriData = JSON.parse(nft.uri)
          if (uriData.attributes && Array.isArray(uriData.attributes)) {
            console.log('从URI JSON中提取属性:', uriData.attributes)
            return uriData.attributes
          }
        } catch (error) {
          console.error('解析URI JSON失败:', error)
        }
      }

      // 回退到默认属性
      return [
        { trait_type: 'Type', value: nft.contractType === 'child' ? 'Child NFT' : 'Main NFT' },
        { trait_type: 'Rarity', value: 'Common' }
      ]
    }

    // 获取NFT是否有属性
    const hasAttributes = (nft) => {
      const attributes = getAttributes(nft)
      return attributes && attributes.length > 0
    }

    // 查看NFT详情
    const viewNFTDetails = (nft) => {
      selectedNFT.value = nft
      if (bsDetailsModal) {
        bsDetailsModal.show()
      }
    }

    // 检查NFT是否有访问控制信息
    const hasAccessControl = (nft) => {
      if (!nft || !nft.metadata) return false

      // 检查属性中是否有Policy或Encrypted_ciphertext
      if (nft.metadata.attributes && Array.isArray(nft.metadata.attributes)) {
        return nft.metadata.attributes.some(attr =>
          attr.trait_type === 'Policy' || attr.trait_type === 'Encrypted_ciphertext'
        )
      }

      return false
    }

    // 获取访问策略
    const getPolicy = (nft) => {
      if (!nft || !nft.metadata || !nft.metadata.attributes) return null

      const policyAttr = nft.metadata.attributes.find(attr => attr.trait_type === 'Policy')
      return policyAttr ? policyAttr.value : null
    }

    // 获取密文
    const getCiphertext = (nft) => {
      if (!nft || !nft.metadata || !nft.metadata.attributes) return null

      const ciphertextAttr = nft.metadata.attributes.find(attr => attr.trait_type === 'Encrypted_ciphertext')
      return ciphertextAttr ? ciphertextAttr.value : null
    }

    // 截断文本
    const truncateText = (text, maxLength) => {
      if (!text) return ''
      return text.length > maxLength ? text.substring(0, maxLength) + '...' : text
    }

    // 复制到剪贴板
    const copyToClipboard = async (text) => {
      try {
        await navigator.clipboard.writeText(text)
        store.dispatch('app/showSuccess', '已复制到剪贴板')
      } catch (error) {
        console.error('复制失败:', error)
        store.dispatch('app/showError', '复制失败')
      }
    }

    // 创建子NFT
    const createChildNFT = (parentTokenId) => {
      // 跳转到创建子NFT页面或打开相应的模态框
      router.push({ 
        name: 'AllNFTs', 
        query: { action: 'createChild', tokenId: parentTokenId } 
      })
    }

    // 检查NFT是否为子NFT
    const isChildNFT = (nft) => {
      return nft.isChildNft || nft.contractType === 'child' || nft.parentTokenId
    }

    // 更新元数据
    const updateMetadata = (nft) => {
      // 关闭详情模态框（如果打开的话）
      if (bsDetailsModal) {
        bsDetailsModal.hide()
      }
      
      // 填充更新表单
      updateForm.value.tokenId = nft.tokenId
      updateForm.value.currentUri = nft.uri
      updateForm.value.newUri = ''
      
      // 打开更新URI模态框
      if (bsUpdateUriModal) {
        bsUpdateUriModal.show()
      }
    }

    // 提交更新URI
    const submitUriUpdate = async () => {
      if (!updateForm.value.newUri.trim()) {
        store.dispatch('app/showError', '请输入新的URI')
        return
      }

      isUpdatingUri.value = true
      try {
        // 调用API更新NFT的URI
        await store.dispatch('nft/updateNFTMetadataURI', {
          tokenId: updateForm.value.tokenId,
          newUri: updateForm.value.newUri
        })
        
        store.dispatch('app/showSuccess', 'NFT元数据URI更新成功')
        
        // 关闭模态框
        if (bsUpdateUriModal) {
          bsUpdateUriModal.hide()
        }
        
        // 重置表单
        updateForm.value = {
          tokenId: '',
          currentUri: '',
          newUri: ''
        }
        
        // 刷新NFT列表
        await loadNFTs()
        
      } catch (error) {
        console.error('更新NFT元数据URI失败:', error)
        const errorMessage = error.response?.data?.error || error.message || '更新NFT元数据URI失败'
        store.dispatch('app/showError', errorMessage)
      } finally {
        isUpdatingUri.value = false
      }
    }

    onMounted(() => {
      // 加载NFT列表
      loadNFTs()

      // 初始化模态框
      if (nftDetailsModal.value) {
        bsDetailsModal = new Modal(nftDetailsModal.value)
      }
      if (updateUriModal.value) {
        bsUpdateUriModal = new Modal(updateUriModal.value)
      }
    })

    return {
      nfts,
      isLoading,
      isConnected,
      selectedNFT,
      nftDetailsModal,
      updateUriModal,
      handleImageError,
      viewNFTDetails,
      hasAccessControl,
      getPolicy,
      getCiphertext,
      truncateText,
      formatAddress,
      getNFTImage,
      getNFTName,
      getNFTDescription,
      getAttributes,
      hasAttributes,
      createChildNFT,
      copyToClipboard,
      isChildNFT,
      updateMetadata,
      updateForm,
      isUpdatingUri,
      submitUriUpdate
    }
  }
}
</script>

<style scoped>
.nft-card {
  transition: transform 0.3s ease;
}

.nft-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1);
}

.card-img-top {
  height: 200px;
  object-fit: cover;
}

.attribute-box {
  background-color: rgba(0, 0, 0, 0.03);
  transition: all 0.2s;
}

.attribute-box:hover {
  background-color: rgba(0, 0, 0, 0.05);
}

/* 密文相关样式 */
.ciphertext-container {
  max-width: 100%;
}

.ciphertext-container code {
  word-break: break-all;
  white-space: pre-wrap;
  color: #666;
}

.ciphertext-detail .ciphertext-content {
  font-family: 'Courier New', monospace;
  font-size: 0.85rem;
  line-height: 1.4;
}

.ciphertext-detail .ciphertext-content code {
  color: #2c3e50;
  background: transparent;
  border: none;
  padding: 0;
}

.policy-display .policy-content code {
  color: #0d6efd;
  font-weight: 600;
}

/* 响应式调整 */
@media (max-width: 768px) {
  .col-6 {
    flex: 0 0 100%;
    max-width: 100%;
  }
  
  /* 移动设备上的按钮样式 */
  .d-flex.flex-wrap {
    justify-content: center;
  }
  
  .btn-sm {
    font-size: 0.8rem;
    padding: 0.375rem 0.75rem;
  }
}

@media (max-width: 576px) {
  .d-flex.flex-wrap .btn {
    flex: 1 1 auto;
    min-width: 100px;
  }
}
</style>