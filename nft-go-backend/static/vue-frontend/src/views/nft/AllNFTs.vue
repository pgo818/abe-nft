<template>
  <div class="all-nfts">
    <h2><i class="bi bi-grid me-2"></i>所有NFT</h2>

    <!-- 加载状态 -->
    <div v-if="isLoading" class="text-center my-5">
      <div class="spinner-border text-primary" role="status">
        <span class="visually-hidden">加载中...</span>
      </div>
      <p class="mt-2">加载NFT列表...</p>
    </div>

    <!-- 没有NFT时显示 -->
    <div v-else-if="!nfts.length" class="text-center my-5">
      <i class="bi bi-collection fs-1 text-muted"></i>
      <p class="lead mt-3">暂无NFT</p>
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
                    <div class="fw-bold">{{ attr.value }}</div>
                  </div>
                </div>
              </div>
            </div>

            <!-- 访问控制信息 -->
            <div v-if="hasAccessControl(nft)" class="mt-3">
              <h6 class="border-bottom pb-2">访问控制</h6>
              <div v-if="getPolicy(nft)" class="mt-2 mb-2">
                <div class="attribute-box p-2 border rounded bg-light">
                  <div class="text-muted small">访问策略</div>
                  <div class="fw-bold">{{ getPolicy(nft) }}</div>
                </div>
              </div>
              <div v-if="getCiphertext(nft)" class="mt-2">
                <div class="attribute-box p-2 border rounded bg-light">
                  <div class="text-muted small">密文</div>
                  <div class="fw-bold text-truncate" :title="getCiphertext(nft)">
                    {{ truncateText(getCiphertext(nft), 20) }}
                  </div>
                </div>
              </div>
            </div>

            <div class="mt-3">
              <p class="card-text"><small class="text-muted">Token ID: {{ nft.tokenId }}</small></p>
              <p class="card-text"><small class="text-muted">Owner: {{ formatAddress(nft.owner) }}</small></p>
            </div>

            <div class="d-flex gap-2 mt-3">
              <button class="btn btn-outline-primary btn-sm" @click="viewNFTDetails(nft)">
                <i class="bi bi-info-circle me-1"></i>详情
              </button>
              <button v-if="isConnected" class="btn btn-outline-primary btn-sm" @click="requestChildNFT(nft.tokenId)">
                <i class="bi bi-plus-circle me-1"></i>申请子NFT
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
                <div v-if="hasAttributes(selectedNFT)" class="mt-3">
                  <h6>属性</h6>
                  <div class="row g-2">
                    <div v-for="(attr, index) in getAttributes(selectedNFT)" :key="index" class="col-6">
                      <div class="attribute-box p-2 border rounded">
                        <div class="text-muted small">{{ attr.trait_type }}</div>
                        <div class="fw-bold">{{ attr.value }}</div>
                      </div>
                    </div>
                  </div>
                </div>

                <!-- 访问控制信息 -->
                <div v-if="hasAccessControl(selectedNFT)" class="mt-3">
                  <h6>访问控制</h6>
                  <div v-if="getPolicy(selectedNFT)" class="mt-2">
                    <div class="attribute-box p-2 border rounded bg-light">
                      <div class="text-muted small">访问策略</div>
                      <div class="fw-bold">{{ getPolicy(selectedNFT) }}</div>
                    </div>
                  </div>
                  <div v-if="getCiphertext(selectedNFT)" class="mt-2">
                    <div class="attribute-box p-2 border rounded bg-light">
                      <div class="text-muted small">密文</div>
                      <div class="fw-bold text-break">{{ getCiphertext(selectedNFT) }}</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>

    <!-- 申请子NFT模态框 -->
    <div class="modal fade" ref="requestChildModal" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">申请子NFT</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <form @submit.prevent="submitChildNFTRequest">
            <div class="modal-body">
              <div class="mb-3">
                <label for="request-parent-token-id" class="form-label">父NFT ID</label>
                <input type="text" class="form-control" id="request-parent-token-id" v-model="requestForm.parentTokenId"
                  readonly>
              </div>
              <div class="mb-3">
                <label for="request-child-uri" class="form-label">子NFT URI</label>
                <input type="text" class="form-control" id="request-child-uri" v-model="requestForm.uri"
                  placeholder="ipfs://..." required>
                <div class="form-text">指向子NFT元数据的URI</div>
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">取消</button>
              <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
                <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-1" role="status"
                  aria-hidden="true"></span>
                提交申请
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
import { Modal } from 'bootstrap'
import { formatAddress } from '@/utils/api'

export default {
  name: 'AllNFTs',

  setup() {
    const store = useStore()
    const requestChildModal = ref(null)
    const nftDetailsModal = ref(null)
    let bsRequestModal = null
    let bsDetailsModal = null

    // 从store获取数据
    const nfts = computed(() => store.state.nft.allNFTs)
    const isLoading = computed(() => store.state.app.isLoading)
    const isConnected = computed(() => store.state.wallet.isConnected)

    // 选中的NFT
    const selectedNFT = ref(null)

    // 申请子NFT表单
    const requestForm = ref({
      parentTokenId: '',
      uri: ''
    })

    // 提交状态
    const isSubmitting = ref(false)

    // 加载NFT列表
    const loadNFTs = async () => {
      await store.dispatch('nft/loadAllNFTs')
    }

    // 处理图片加载错误
    const handleImageError = (event) => {
      event.target.src = 'data:image/svg+xml;charset=UTF-8,%3Csvg%20width%3D%22200%22%20height%3D%22200%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22%23f0f0f0%22%2F%3E%3Ctext%20x%3D%2250%25%22%20y%3D%2250%25%22%20font-family%3D%22Arial%22%20font-size%3D%2224%22%20text-anchor%3D%22middle%22%20dominant-baseline%3D%22middle%22%20fill%3D%22%23999%22%3ENFT%20%E5%8D%A0%E4%BD%8D%E5%9B%BE%3C%2Ftext%3E%3C%2Fsvg%3E'
    }

    // 查看NFT详情
    const viewNFTDetails = (nft) => {
      selectedNFT.value = nft
      if (bsDetailsModal) {
        bsDetailsModal.show()
      }
    }

    // 申请子NFT
    const requestChildNFT = (parentTokenId) => {
      if (!isConnected.value) {
        store.dispatch('app/showWarning', '请先连接钱包')
        store.commit('wallet/setShowConnectPrompt', true)
        return
      }

      requestForm.value.parentTokenId = parentTokenId
      requestForm.value.uri = ''

      if (bsRequestModal) {
        bsRequestModal.show()
      }
    }

    // 提交子NFT申请
    const submitChildNFTRequest = async () => {
      if (!requestForm.value.parentTokenId || !requestForm.value.uri) {
        store.dispatch('app/showError', '请填写所有必填字段')
        return
      }

      try {
        isSubmitting.value = true

        const result = await store.dispatch('nft/requestChildNFT', {
          parentTokenId: requestForm.value.parentTokenId,
          uri: requestForm.value.uri
        })

        if (result) {
          // 隐藏模态框
          if (bsRequestModal) {
            bsRequestModal.hide()
          }

          // 重置表单
          requestForm.value.parentTokenId = ''
          requestForm.value.uri = ''
        }
      } finally {
        isSubmitting.value = false
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

    onMounted(() => {
      // 加载NFT列表
      loadNFTs()

      // 初始化模态框
      if (requestChildModal.value) {
        bsRequestModal = new Modal(requestChildModal.value)
      }

      if (nftDetailsModal.value) {
        bsDetailsModal = new Modal(nftDetailsModal.value)
      }
    })

    return {
      nfts,
      isLoading,
      isConnected,
      selectedNFT,
      requestForm,
      isSubmitting,
      requestChildModal,
      nftDetailsModal,
      handleImageError,
      viewNFTDetails,
      requestChildNFT,
      submitChildNFTRequest,
      hasAccessControl,
      getPolicy,
      getCiphertext,
      truncateText,
      formatAddress,
      getNFTImage,
      getNFTName,
      getNFTDescription,
      getAttributes,
      hasAttributes
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
</style>