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
            </div>

            <div class="d-flex gap-2 mt-3">
              <button class="btn btn-outline-primary btn-sm" @click="viewNFTDetails(nft)">
                <i class="bi bi-info-circle me-1"></i>详情
              </button>
              <button class="btn btn-outline-secondary btn-sm" @click="transferNFT(nft)">
                <i class="bi bi-arrow-right me-1"></i>转移
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
            <button type="button" class="btn btn-primary" @click="transferNFT(selectedNFT)" v-if="selectedNFT">
              <i class="bi bi-arrow-right me-1"></i>转移NFT
            </button>
          </div>
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
    let bsDetailsModal = null

    // 从store获取数据
    const nfts = computed(() => store.state.nft.myNFTs)
    const isLoading = computed(() => store.state.nft.isLoading)
    const isConnected = computed(() => store.state.wallet.isConnected)

    // 选中的NFT
    const selectedNFT = ref(null)

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
      return nft.metadata?.image || nft.uri
    }

    // 获取NFT名称
    const getNFTName = (nft) => {
      return nft.metadata?.name || 'NFT #' + nft.tokenId
    }

    // 获取NFT描述
    const getNFTDescription = (nft) => {
      return nft.metadata?.description || 'No description'
    }

    // 获取NFT属性
    const getAttributes = (nft) => {
      return nft.metadata?.attributes || []
    }

    // 获取NFT是否有属性
    const hasAttributes = (nft) => {
      return nft.metadata?.attributes && nft.metadata.attributes.length > 0
    }

    // 查看NFT详情
    const viewNFTDetails = (nft) => {
      selectedNFT.value = nft
      if (bsDetailsModal) {
        bsDetailsModal.show()
      }
    }

    // 转移NFT
    const transferNFT = (nft) => {
      if (!nft) return

      // 保存选中的NFT
      store.commit('nft/setSelectedNFT', nft)
      router.push({ name: 'TransferNFT', params: { tokenId: nft.tokenId } })
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

    onMounted(() => {
      // 加载NFT列表
      loadNFTs()

      // 初始化模态框
      if (nftDetailsModal.value) {
        bsDetailsModal = new Modal(nftDetailsModal.value)
      }
    })

    return {
      nfts,
      isLoading,
      isConnected,
      selectedNFT,
      nftDetailsModal,
      handleImageError,
      viewNFTDetails,
      transferNFT,
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