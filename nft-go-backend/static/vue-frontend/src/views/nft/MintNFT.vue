<template>
  <div class="mint-nft">
    <h2><i class="bi bi-plus-circle me-2"></i>铸造新NFT</h2>

    <div class="row">
      <div class="col-md-6">
        <div class="card">
          <div class="card-body">
            <form @submit.prevent="mintNFT">
              <div class="mb-3">
                <label for="nft-uri" class="form-label">NFT URI (IPFS或HTTP)</label>
                <div class="input-group">
                  <input type="text" class="form-control" id="nft-uri" v-model="uri" placeholder="ipfs://..." required
                    @input="debouncedFetchMetadata">
                  <button class="btn btn-outline-secondary" type="button" @click="openMetadataSelector">
                    <i class="bi bi-search me-1"></i>选择元数据
                  </button>
                </div>
                <div class="form-text">指向NFT元数据的URI</div>
              </div>
              <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
                <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-1" role="status"
                  aria-hidden="true"></span>
                <i v-else class="bi bi-plus-circle me-1"></i>铸造NFT
              </button>
            </form>
          </div>
        </div>

        <!-- 最近创建的元数据 -->
        <div class="card mt-4" v-if="metadata.length">
          <div class="card-header d-flex justify-content-between align-items-center">
            <h5 class="mb-0"><i class="bi bi-list me-2"></i>最近创建的元数据</h5>
            <button class="btn btn-sm btn-outline-primary" @click="loadMetadata">
              <i class="bi bi-arrow-clockwise me-1"></i>刷新
            </button>
          </div>
          <div class="card-body">
            <div class="table-responsive">
              <table class="table">
                <thead>
                  <tr>
                    <th>名称</th>
                    <th>描述</th>
                    <th>IPFS哈希</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in metadata.slice(0, 5)" :key="item.id">
                    <td>{{ item.name }}</td>
                    <td :title="item.description">{{ truncateText(item.description, 30) }}</td>
                    <td>
                      <code class="small">{{ item.ipfs_hash }}</code>
                    </td>
                    <td>
                      <button class="btn btn-sm btn-success" @click="useAsURI(item.ipfs_hash)">
                        <i class="bi bi-check-circle me-1"></i>用作URI
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div class="text-center mt-2">
              <router-link :to="{ name: 'MetadataManager' }" class="btn btn-outline-primary btn-sm">
                <i class="bi bi-arrow-right me-1"></i>查看全部元数据
              </router-link>
            </div>
          </div>
        </div>
      </div>

      <!-- 元数据预览卡片 -->
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h5 class="mb-0">元数据预览</h5>
          </div>
          <div class="card-body">
            <!-- 加载中状态 -->
            <div v-if="isLoadingPreview" class="text-center my-4">
              <div class="spinner-border text-primary" role="status">
                <span class="visually-hidden">加载中...</span>
              </div>
              <p class="mt-2">解析元数据中...</p>
            </div>

            <!-- 无元数据状态 -->
            <div v-else-if="!previewMetadata" class="text-center my-4">
              <i class="bi bi-card-image fs-1 text-muted"></i>
              <p class="mt-3">输入有效的URI以查看元数据预览</p>
            </div>

            <!-- 元数据预览 -->
            <div v-else class="metadata-preview">
              <img v-if="previewMetadata.image" :src="previewMetadata.image" class="img-fluid mb-3 rounded"
                alt="NFT Preview" @error="handleImageError">
              <img v-else
                src="data:image/svg+xml;charset=UTF-8,%3Csvg%20width%3D%22200%22%20height%3D%22200%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22%23f0f0f0%22%2F%3E%3Ctext%20x%3D%2250%25%22%20y%3D%2250%25%22%20font-family%3D%22Arial%22%20font-size%3D%2224%22%20text-anchor%3D%22middle%22%20dominant-baseline%3D%22middle%22%20fill%3D%22%23999%22%3ENFT%20%E5%8D%A0%E4%BD%8D%E5%9B%BE%3C%2Ftext%3E%3C%2Fsvg%3E"
                class="img-fluid mb-3 rounded" alt="NFT Preview">

              <h4>{{ previewMetadata.name || 'NFT名称' }}</h4>
              <p>{{ previewMetadata.description || '暂无描述' }}</p>

              <!-- 元数据属性展示 -->
              <div v-if="previewMetadata.attributes && previewMetadata.attributes.length > 0" class="mt-3">
                <h6 class="border-bottom pb-2">属性</h6>
                <div class="row g-2 mt-2">
                  <div v-for="(attr, index) in previewMetadata.attributes" :key="index" class="col-6">
                    <div class="attribute-box p-2 border rounded">
                      <div class="text-muted small">{{ attr.trait_type }}</div>
                      <div class="fw-bold">{{ attr.value }}</div>
                    </div>
                  </div>
                </div>
              </div>


            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 元数据选择模态框 -->
    <div class="modal fade" ref="metadataModal" tabindex="-1">
      <div class="modal-dialog modal-lg">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">选择元数据</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <div v-if="isLoadingMetadata" class="text-center my-4">
              <div class="spinner-border text-primary" role="status">
                <span class="visually-hidden">加载中...</span>
              </div>
              <p class="mt-2">加载元数据列表...</p>
            </div>
            <div v-else-if="!metadata.length" class="text-center my-4">
              <p class="lead">暂无元数据</p>
              <router-link :to="{ name: 'MetadataManager' }" class="btn btn-primary">
                <i class="bi bi-plus-circle me-1"></i>创建元数据
              </router-link>
            </div>
            <div v-else class="table-responsive">
              <table class="table table-hover">
                <thead>
                  <tr>
                    <th>名称</th>
                    <th>描述</th>
                    <th>IPFS哈希</th>
                    <th>操作</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="item in metadata" :key="item.id" @click="selectMetadata(item)">
                    <td>{{ item.name }}</td>
                    <td :title="item.description">{{ truncateText(item.description, 30) }}</td>
                    <td><code class="small">{{ item.ipfs_hash }}</code></td>
                    <td>
                      <button class="btn btn-sm btn-primary" @click.stop="selectMetadata(item)">
                        <i class="bi bi-check-circle me-1"></i>选择
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">取消</button>
          </div>
        </div>
      </div>
    </div>

    <!-- 交易结果模态框 -->
    <div class="modal fade" ref="resultModal" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">铸造结果</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <div class="alert alert-success">
              <h5 class="alert-heading">NFT铸造交易已提交!</h5>
              <p>您的NFT铸造交易已成功提交到区块链，请等待确认。</p>
              <hr>
              <p class="mb-0">交易哈希:</p>
              <code class="d-block mt-2 p-2 bg-light">{{ transactionHash }}</code>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">关闭</button>
            <router-link :to="{ name: 'MyNFTs' }" class="btn btn-primary">
              <i class="bi bi-wallet me-1"></i>查看我的NFT
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted, watch } from 'vue'
import { useStore } from 'vuex'
import { Modal } from 'bootstrap'
import nftService from '@/services/nftService'

export default {
  name: 'MintNFT',

  setup() {
    const store = useStore()
    const metadataModal = ref(null)
    const resultModal = ref(null)
    let bsMetadataModal = null
    let bsResultModal = null

    // 表单数据
    const uri = ref('')
    const isSubmitting = ref(false)
    const transactionHash = ref('')

    // 预览相关
    const previewMetadata = ref(null)
    const isLoadingPreview = ref(false)
    let fetchTimeout = null

    // 从store获取数据
    const metadata = computed(() => store.state.nft.metadata)
    const isLoading = computed(() => store.state.app.isLoading)

    // 元数据加载状态
    const isLoadingMetadata = ref(false)

    // 加载元数据列表
    const loadMetadata = async () => {
      isLoadingMetadata.value = true
      try {
        await store.dispatch('nft/loadMetadata')
      } finally {
        isLoadingMetadata.value = false
      }
    }

    // 打开元数据选择器
    const openMetadataSelector = () => {
      if (bsMetadataModal) {
        bsMetadataModal.show()
      }
    }

    // 选择元数据
    const selectMetadata = (item) => {
      uri.value = `ipfs://${item.ipfs_hash}`
      if (bsMetadataModal) {
        bsMetadataModal.hide()
      }
      fetchMetadata(uri.value)
    }

    // 使用IPFS哈希作为URI
    const useAsURI = (hash) => {
      uri.value = `ipfs://${hash}`
      fetchMetadata(uri.value)
    }

    // 获取元数据内容
    const fetchMetadata = async (uriValue) => {
      if (!uriValue) {
        previewMetadata.value = null
        return
      }

      isLoadingPreview.value = true
      try {
        const data = await nftService.fetchMetadataFromURI(uriValue)
        previewMetadata.value = data
      } catch (error) {
        console.error('获取元数据失败:', error)
        store.dispatch('app/showError', '获取元数据失败: ' + error.message)
        previewMetadata.value = null
      } finally {
        isLoadingPreview.value = false
      }
    }

    // 防抖处理，避免频繁请求
    const debouncedFetchMetadata = () => {
      if (fetchTimeout) {
        clearTimeout(fetchTimeout)
      }
      fetchTimeout = setTimeout(() => {
        fetchMetadata(uri.value)
      }, 500)
    }

    // 铸造NFT
    const mintNFT = async () => {
      if (!uri.value) {
        store.dispatch('app/showError', '请输入NFT URI')
        return
      }

      isSubmitting.value = true

      try {
        // 检查钱包连接
        if (!store.state.wallet.isConnected) {
          store.dispatch('app/showError', '请先连接钱包')
          return
        }

        // 直接传递URI字符串
        const hash = await store.dispatch('nft/mintNFT', uri.value)

        if (hash) {
          transactionHash.value = hash

          // 显示结果模态框
          if (bsResultModal) {
            bsResultModal.show()
          }

          // 重置表单
          uri.value = ''
          previewMetadata.value = null
        }
      } catch (error) {
        console.error('铸造NFT失败:', error)
        store.dispatch('app/showError', '铸造NFT失败: ' + (error.message || '未知错误'))
      } finally {
        isSubmitting.value = false
      }
    }

    // 处理图片加载错误
    const handleImageError = (event) => {
      event.target.src = 'data:image/svg+xml;charset=UTF-8,%3Csvg%20width%3D%22200%22%20height%3D%22200%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22%23f0f0f0%22%2F%3E%3Ctext%20x%3D%2250%25%22%20y%3D%2250%25%22%20font-family%3D%22Arial%22%20font-size%3D%2224%22%20text-anchor%3D%22middle%22%20dominant-baseline%3D%22middle%22%20fill%3D%22%23999%22%3ENFT%20%E5%8D%A0%E4%BD%8D%E5%9B%BE%3C%2Ftext%3E%3C%2Fsvg%3E'
      event.target.onerror = null
    }

    // 检查NFT是否有访问控制信息
    const hasAccessControl = (metadata) => {
      if (!metadata || !metadata.attributes) return false

      // 检查属性中是否有Policy或Encrypted_ciphertext
      if (metadata.attributes && Array.isArray(metadata.attributes)) {
        return metadata.attributes.some(attr =>
          attr.trait_type === 'Policy' || attr.trait_type === 'Encrypted_ciphertext'
        )
      }

      return false
    }

    // 获取访问策略
    const getPolicy = (metadata) => {
      if (!metadata || !metadata.attributes) return null

      const policyAttr = metadata.attributes.find(attr => attr.trait_type === 'Policy')
      return policyAttr ? policyAttr.value : null
    }

    // 获取密文
    const getCiphertext = (metadata) => {
      if (!metadata || !metadata.attributes) return null

      const ciphertextAttr = metadata.attributes.find(attr => attr.trait_type === 'Encrypted_ciphertext')
      return ciphertextAttr ? ciphertextAttr.value : null
    }

    // 格式化日期
    const formatDate = (dateString) => {
      if (!dateString) return ''
      const date = new Date(dateString)
      return date.toLocaleString()
    }

    // 截断文本
    const truncateText = (text, length) => {
      if (!text) return ''
      return text.length > length ? text.substring(0, length) + '...' : text
    }

    // 监听URI变化
    watch(uri, (newUri) => {
      if (!newUri) {
        previewMetadata.value = null
      }
    })

    onMounted(() => {
      // 加载元数据列表
      loadMetadata()

      // 初始化模态框
      if (metadataModal.value) {
        bsMetadataModal = new Modal(metadataModal.value)
      }

      if (resultModal.value) {
        bsResultModal = new Modal(resultModal.value)
      }
    })

    return {
      uri,
      metadata,
      isLoading,
      isSubmitting,
      isLoadingMetadata,
      isLoadingPreview,
      previewMetadata,
      transactionHash,
      metadataModal,
      resultModal,
      mintNFT,
      loadMetadata,
      openMetadataSelector,
      selectMetadata,
      useAsURI,
      debouncedFetchMetadata,
      handleImageError,
      hasAccessControl,
      getPolicy,
      getCiphertext,
      formatDate,
      truncateText
    }
  }
}
</script>

<style scoped>
.table tbody tr {
  cursor: pointer;
}

.table tbody tr:hover {
  background-color: rgba(0, 0, 0, 0.05);
}

.metadata-preview img {
  max-height: 300px;
  object-fit: contain;
  width: 100%;
}

.attribute-box {
  background-color: rgba(0, 0, 0, 0.03);
  transition: all 0.2s;
}

.attribute-box:hover {
  background-color: rgba(0, 0, 0, 0.05);
}
</style>