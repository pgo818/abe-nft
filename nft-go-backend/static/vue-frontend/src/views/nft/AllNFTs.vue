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
                    <!-- 特殊处理密文显示 -->
                    <div v-if="attr.trait_type === 'Encrypted_ciphertext'" class="fw-bold">
                      <div class="ciphertext-detail">
                        <div class="d-flex justify-content-between align-items-start mb-2">
                          <span class="badge bg-secondary">密文 ({{ attr.value?.length || 0 }} 字符)</span>
                          <span class="badge bg-warning">
                            <i class="bi bi-shield-lock me-1"></i>受保护
                          </span>
                        </div>
                        <div class="ciphertext-content p-2 border rounded bg-light">
                          <div class="text-center text-muted py-3">
                            <i class="bi bi-shield-lock fs-1 mb-2"></i>
                            <p class="mb-1">密文内容已加密保护</p>
                            <small>只有NFT所有者才能查看完整密文</small>
                          </div>
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
                    <div v-for="(attr, index) in getAttributes(selectedNFT)" :key="index" class="col-12">
                      <div class="attribute-box p-3 border rounded">
                        <div class="text-muted small mb-1">{{ attr.trait_type }}</div>
                        <!-- 特殊处理密文显示 -->
                        <div v-if="attr.trait_type === 'Encrypted_ciphertext'" class="fw-bold">
                          <div class="ciphertext-detail">
                            <div class="d-flex justify-content-between align-items-start mb-2">
                              <span class="badge bg-secondary">密文 ({{ attr.value?.length || 0 }} 字符)</span>
                              <span class="badge bg-warning">
                                <i class="bi bi-shield-lock me-1"></i>受保护
                              </span>
                            </div>
                            <div class="ciphertext-content p-2 border rounded bg-light">
                              <div class="text-center text-muted py-3">
                                <i class="bi bi-shield-lock fs-1 mb-2"></i>
                                <p class="mb-1">密文内容已加密保护</p>
                                <small>只有NFT所有者才能查看完整密文</small>
                              </div>
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
                  readonly>
                <div class="form-text">子NFT URI将自动使用父NFT的URI，确保保持一致</div>
              </div>
              
              <!-- 新增：自动审核选项 -->
              <div class="mb-3">
                <div class="form-check">
                  <input class="form-check-input" type="checkbox" id="auto-approve-check" v-model="requestForm.autoApprove">
                  <label class="form-check-label" for="auto-approve-check">
                    使用VC凭证自动审核
                  </label>
                  <div class="form-text">如果您的VC凭证满足访问策略，系统将自动审核通过</div>
                </div>
                
                <!-- 调试按钮 -->
                <div class="mt-2">
                  <button type="button" class="btn btn-sm btn-outline-info" @click="debugVCCredentials">
                    <i class="bi bi-bug me-1"></i>测试VC凭证加载
                  </button>
                  <small class="text-muted ms-2">点击此按钮检查VC凭证状态</small>
                </div>
              </div>
              
              <!-- 新增：VC凭证选择 -->
              <div v-if="requestForm.autoApprove" class="mb-3">
                <label for="vc-select" class="form-label">选择VC凭证</label>
                <select class="form-select" id="vc-select" v-model="requestForm.selectedVCId" required>
                  <option value="">请选择一个VC凭证</option>
                  <option v-for="vc in userVCs" :key="vc.vcId" :value="vc.vcId">
                    {{ vc.type }} - {{ vc.content }} ({{ formatDate(vc.issuedAt) }})
                  </option>
                </select>
                <div class="form-text">
                  <small class="text-muted">选择的VC凭证将用于验证是否满足NFT的访问策略</small>
                </div>
              </div>
              
              <!-- 显示选中VC的详细信息 -->
              <div v-if="selectedVCInfo" class="mb-3">
                <div class="card">
                  <div class="card-header bg-light">
                    <h6 class="mb-0">VC凭证信息</h6>
                  </div>
                  <div class="card-body">
                    <div class="row">
                      <div class="col-sm-4"><strong>姓名:</strong></div>
                      <div class="col-sm-8">{{ selectedVCInfo.name || '未设置' }}</div>
                    </div>
                    <div class="row">
                      <div class="col-sm-4"><strong>科室:</strong></div>
                      <div class="col-sm-8">{{ selectedVCInfo.department || '未设置' }}</div>
                    </div>
                    <div class="row">
                      <div class="col-sm-4"><strong>医院:</strong></div>
                      <div class="col-sm-8">{{ selectedVCInfo.hospital || '未设置' }}</div>
                    </div>
                    <div class="row">
                      <div class="col-sm-4"><strong>职称:</strong></div>
                      <div class="col-sm-8">{{ selectedVCInfo.title || '未设置' }}</div>
                    </div>
                    <div class="row">
                      <div class="col-sm-4"><strong>专长:</strong></div>
                      <div class="col-sm-8">{{ selectedVCInfo.specialty || '未设置' }}</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">取消</button>
              <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
                <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-1" role="status"
                  aria-hidden="true"></span>
                {{ requestForm.autoApprove ? '提交自动审核申请' : '提交申请' }}
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
      uri: '',
      autoApprove: false,
      selectedVCId: ''
    })

    // 提交状态
    const isSubmitting = ref(false)
    
    // 用户VC凭证
    const userVCs = ref([])
    
    // 计算选中VC的详细信息
    const selectedVCInfo = computed(() => {
      if (!requestForm.value.selectedVCId) return null
      
      const selectedVC = userVCs.value.find(vc => vc.vcId === requestForm.value.selectedVCId)
      if (!selectedVC) return null
      
      try {
        const content = JSON.parse(selectedVC.content)
        return content.credentialSubject || content
      } catch (error) {
        console.error('解析VC内容失败:', error)
        return null
      }
    })

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
    const requestChildNFT = async (parentTokenId) => {
      if (!isConnected.value) {
        store.dispatch('app/showWarning', '请先连接钱包')
        store.commit('wallet/setShowConnectPrompt', true)
        return
      }

      // 从NFT列表中找到对应的主NFT
      const parentNFT = nfts.value.find(nft => nft.tokenId === parentTokenId)
      
      requestForm.value.parentTokenId = parentTokenId
      // 自动使用父NFT的URI
      requestForm.value.uri = parentNFT ? parentNFT.uri : ''
      requestForm.value.autoApprove = false
      requestForm.value.selectedVCId = ''

      // 加载用户的VC凭证
      await loadUserVCs()

      if (bsRequestModal) {
        bsRequestModal.show()
      }
    }

    // 提交子NFT申请
    const submitChildNFTRequest = async () => {
      if (isSubmitting.value) return
      
      try {
        isSubmitting.value = true
        
        if (!isConnected.value) {
          store.dispatch('app/showError', '请先连接钱包')
          return
        }
        
        // 验证URI
        if (!requestForm.value.uri.trim()) {
          store.dispatch('app/showError', '请输入有效的子NFT URI')
          return
        }
        
        // 如果启用自动审核，验证VC凭证
        let vcCredentials = null
        if (requestForm.value.autoApprove) {
          if (!requestForm.value.selectedVCId) {
            store.dispatch('app/showError', '请选择一个VC凭证进行自动审核')
            return
          }
          
          // 获取选中的VC凭证
          const selectedVC = userVCs.value.find(vc => vc.vcId === requestForm.value.selectedVCId)
          if (!selectedVC) {
            store.dispatch('app/showError', '找不到选中的VC凭证')
            return
          }
          
          vcCredentials = selectedVC.content
        }
        
        console.log('提交子NFT申请:', {
          parentTokenId: requestForm.value.parentTokenId,
          uri: requestForm.value.uri,
          autoApprove: requestForm.value.autoApprove,
          hasVCCredentials: !!vcCredentials
        })
        
        // 调用store中的requestChildNFT方法
        const result = await store.dispatch('nft/requestChildNFT', {
          parentTokenId: requestForm.value.parentTokenId,
          uri: requestForm.value.uri,
          autoApprove: requestForm.value.autoApprove,
          vcCredentials: vcCredentials
        })
        
        console.log('子NFT申请结果:', result)
        
        // 处理响应
        if (result.autoApproved) {
          store.dispatch('app/showSuccess', 
            `恭喜！您的VC凭证验证通过，子NFT已自动创建。交易哈希: ${result.transactionHash}`)
        } else {
          // 检查是否有策略验证失败的详细信息
          if (result.policyResult && result.policyResult.detailedReason) {
            store.dispatch('app/showWarning', 
              `自动审核未通过：${result.policyResult.detailedReason}。申请已提交，等待手动审核。`)
          } else {
            store.dispatch('app/showInfo', result.message || '子NFT申请已提交，等待父NFT持有者审批')
          }
        }
        
        // 关闭模态框
        if (bsRequestModal) {
          bsRequestModal.hide()
        }
        
        // 重置表单
        requestForm.value = {
          parentTokenId: '',
          uri: '',
          autoApprove: false,
          selectedVCId: ''
        }
        
        // 刷新数据
        await store.dispatch('nft/loadAllNFTs')
        
      } catch (error) {
        console.error('提交子NFT申请失败:', error)
        
        // 检查是否有详细的策略验证错误信息
        if (error.response?.data?.policyResult?.detailedReason) {
          store.dispatch('app/showError', 
            `策略验证失败：${error.response.data.policyResult.detailedReason}`)
        } else if (error.response?.data?.error?.includes('还没有创建元数据')) {
          // 特殊处理元数据缺失的情况
          store.dispatch('app/showWarning', 
            `${error.response.data.error}。点击此处可跳转到创建元数据页面。`)
          
          // 可以考虑添加一个跳转按钮或自动跳转逻辑
          setTimeout(() => {
            // 询问用户是否要跳转到元数据创建页面
            if (confirm('是否现在就去创建元数据？')) {
              // 跳转到元数据创建页面，这里需要根据您的路由配置调整
              window.location.href = '/metadata-manager'
            }
          }, 2000)
        } else {
          store.dispatch('app/showError', error.response?.data?.error || error.message || '提交申请失败')
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
    const getAccessPolicy = (nft) => {
      if (!nft || !nft.metadata || !nft.metadata.attributes) return '无策略'

      const policyAttr = nft.metadata.attributes.find(attr => attr.trait_type === 'Policy')
      return policyAttr ? policyAttr.value : '无策略'
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

    // 加载用户的VC凭证
    const loadUserVCs = async () => {
      try {
        if (!isConnected.value) {
          console.log('钱包未连接，无法加载VC凭证')
          userVCs.value = []
          return
        }
        
        const walletAddress = store.state.wallet.account
        console.log('开始加载VC凭证，当前钱包地址:', walletAddress)
        
        if (!walletAddress) {
          console.log('钱包地址为空，无法加载VC凭证')
          userVCs.value = []
          return
        }
        
        // 调用store中的loadDoctorVCs方法
        const vcs = await store.dispatch('did/loadDoctorVCs', {
          walletAddress: walletAddress
        })
        
        console.log('从store获取的VC凭证:', vcs)
        
        // 确保userVCs是响应式更新的
        userVCs.value = Array.isArray(vcs) ? vcs : []
        
        if (userVCs.value.length === 0) {
          console.log('没有找到VC凭证，可能需要先创建医生DID和凭证')
          store.dispatch('app/showInfo', '您还没有VC凭证，请先在医生DID页面创建您的身份凭证')
        } else {
          console.log(`成功加载${userVCs.value.length}个VC凭证`)
          
          // 显示第一个VC的详细信息用于调试
          if (vcs.length > 0) {
            console.log('第一个VC的字段信息:')
            console.log('- vcId:', vcs[0].vcId)
            console.log('- type:', vcs[0].type)
            console.log('- content:', vcs[0].content)
            console.log('- issuedAt:', vcs[0].issuedAt)
          }
        }
        
      } catch (error) {
        console.error('加载VC凭证失败:', error)
        userVCs.value = []
        store.dispatch('app/showError', '加载VC凭证失败: ' + error.message)
      }
    }

    // 格式化日期
    const formatDate = (dateString) => {
      try {
        return new Date(dateString).toLocaleString('zh-CN')
      } catch (error) {
        return '未知日期'
      }
    }

    // 调试VC凭证
    const debugVCCredentials = async () => {
      console.log('=== 开始VC凭证调试 ===')
      
      try {
        // 1. 检查钱包连接状态
        console.log('1. 钱包连接状态:', isConnected.value)
        console.log('2. 钱包地址:', store.state.wallet.account)
        
        if (!isConnected.value) {
          store.dispatch('app/showError', '钱包未连接，请先连接钱包')
          return
        }
        
        const walletAddress = store.state.wallet.account
        if (!walletAddress) {
          store.dispatch('app/showError', '无法获取钱包地址')
          return
        }
        
        // 2. 测试医生DID API
        console.log('3. 测试获取医生DID列表...')
        try {
          const didResponse = await fetch('/api/did/doctor/list')
          const didData = await didResponse.json()
          console.log('医生DID API响应:', didData)
          
          if (!didResponse.ok) {
            console.error('医生DID API错误:', didData)
            store.dispatch('app/showError', `医生DID API错误: ${didData.error}`)
            return
          }
          
          const doctors = didData.doctors || []
          console.log('找到的医生数量:', doctors.length)
          
          // 查找当前钱包对应的医生
          const currentDoctor = doctors.find(doctor => 
            doctor.walletAddress && doctor.walletAddress.toLowerCase() === walletAddress.toLowerCase()
          )
          
          if (!currentDoctor) {
            console.log('4. 没有找到对应的医生DID')
            store.dispatch('app/showWarning', '您还没有创建医生DID，请先在医生DID页面创建您的身份')
            return
          }
          
          console.log('4. 找到医生DID:', currentDoctor)
          
          // 3. 测试VC凭证API
          console.log('5. 测试获取VC凭证...')
          const vcResponse = await fetch(`/api/vc/doctor/${encodeURIComponent(currentDoctor.didString)}`)
          const vcData = await vcResponse.json()
          console.log('VC API响应:', vcData)
          
          if (!vcResponse.ok) {
            console.error('VC API错误:', vcData)
            store.dispatch('app/showError', `VC API错误: ${vcData.error}`)
            return
          }
          
          const vcs = vcData.verifiableCredentials || []
          console.log('6. 找到的VC凭证数量:', vcs.length)
          console.log('VC凭证详情:', vcs)
          
          if (vcs.length === 0) {
            store.dispatch('app/showWarning', '您的医生DID还没有VC凭证，请在医生DID页面申请凭证')
          } else {
            // 更新userVCs
            userVCs.value = vcs
            store.dispatch('app/showSuccess', `成功加载${vcs.length}个VC凭证`)
            
            // 显示第一个VC的详细信息用于调试
            if (vcs.length > 0) {
              console.log('第一个VC的字段信息:')
              console.log('- vcId:', vcs[0].vcId)
              console.log('- type:', vcs[0].type)
              console.log('- content:', vcs[0].content)
              console.log('- issuedAt:', vcs[0].issuedAt)
            }
          }
          
        } catch (apiError) {
          console.error('API调用失败:', apiError)
          store.dispatch('app/showError', 'API调用失败: ' + apiError.message)
        }
        
      } catch (error) {
        console.error('调试过程出错:', error)
        store.dispatch('app/showError', '调试过程出错: ' + error.message)
      }
      
      console.log('=== VC凭证调试完成 ===')
    }

    // 复制到剪贴板
    const copyToClipboard = async (text) => {
      try {
        await navigator.clipboard.writeText(text)
        store.dispatch('app/showSuccess', '密文已复制到剪贴板')
      } catch (error) {
        console.error('复制失败:', error)
        // 降级方案：创建临时文本框
        const textArea = document.createElement('textarea')
        textArea.value = text
        document.body.appendChild(textArea)
        textArea.focus()
        textArea.select()
        try {
          document.execCommand('copy')
          store.dispatch('app/showSuccess', '密文已复制到剪贴板')
        } catch (fallbackError) {
          console.error('降级复制方案也失败:', fallbackError)
          store.dispatch('app/showError', '复制失败，请手动选择并复制')
        }
        document.body.removeChild(textArea)
      }
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
      getAccessPolicy,
      getCiphertext,
      truncateText,
      formatAddress,
      getNFTImage,
      getNFTName,
      getNFTDescription,
      getAttributes,
      hasAttributes,
      userVCs,
      selectedVCInfo,
      loadUserVCs,
      formatDate,
      debugVCCredentials,
      copyToClipboard
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
  .ciphertext-detail .d-flex {
    flex-direction: column;
    align-items: flex-start;
  }
  
  .ciphertext-detail .btn {
    margin-top: 0.5rem;
  }
}
</style>