<template>
  <div class="abe-encrypt">
    <h2><i class="bi bi-lock me-2"></i>ABE数据加密</h2>

    <div class="row mt-4">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h5 class="mb-0"><i class="bi bi-pencil-square me-2"></i>加密表单</h5>
          </div>
          <div class="card-body">
            <form @submit.prevent="encryptData">
              <div class="mb-3">
                <label for="message" class="form-label">要加密的数据</label>
                <textarea class="form-control" id="message" v-model="encryptForm.message" rows="4"
                  placeholder="输入需要加密的数据内容" required></textarea>
                <div class="form-text">输入任何文本数据，系统将使用ABE算法进行加密</div>
              </div>

              <div class="mb-3">
                <label class="form-label">访问策略</label>
                <div class="form-control bg-light text-monospace p-2" style="min-height: 38px;">
                  {{ currentPolicy }}
                </div>
                <div class="form-text">
                  <i class="bi bi-info-circle me-1"></i>
                  固定策略：只有当前连接的钱包地址才能解密数据
                </div>
              </div>

              <div class="mb-3" v-if="!walletAddress">
                <div class="alert alert-warning">
                  <i class="bi bi-exclamation-triangle me-2"></i>
                  请先连接钱包才能进行加密操作
                </div>
              </div>

              <div class="mb-3">
                <div class="form-check">
                  <input class="form-check-input" type="checkbox" id="saveToIPFS" v-model="encryptForm.saveToIPFS">
                  <label class="form-check-label" for="saveToIPFS">
                    完整IPFS+NFT工作流程
                  </label>
                  <div class="form-text">
                    启用后将执行：1. 上传原文到IPFS → 2. 加密文件Hash → 3. 上传密文到IPFS → 4. 创建NFT
                  </div>
                </div>
              </div>

              <!-- NFT元数据字段 (仅在启用IPFS+NFT工作流程时显示) -->
              <div v-if="encryptForm.saveToIPFS" class="border rounded p-3 mb-3">
                <h6 class="mb-3"><i class="bi bi-collection me-2"></i>NFT元数据信息</h6>
                
                <div class="mb-3">
                  <label for="nftName" class="form-label">NFT名称</label>
                  <input type="text" class="form-control" id="nftName" v-model="nftMetadata.name"
                    placeholder="输入NFT名称" required>
                </div>

                <div class="mb-3">
                  <label for="nftDescription" class="form-label">NFT描述</label>
                  <textarea class="form-control" id="nftDescription" v-model="nftMetadata.description" rows="2"
                    placeholder="输入NFT描述" required></textarea>
                </div>

                <div class="mb-3">
                  <label for="nftImage" class="form-label">NFT图像</label>
                  
                  <!-- 图片上传区域 -->
                  <div class="image-upload-section">
                    <!-- 上传按钮 -->
                    <div class="mb-2">
                      <input type="file" class="d-none" ref="imageInput" @change="handleImageSelect" 
                        accept="image/jpeg,image/jpg,image/png,image/gif,image/webp,image/svg+xml,image/bmp">
                      <button type="button" class="btn btn-outline-primary btn-sm" @click="$refs.imageInput.click()"
                        :disabled="isUploadingImage">
                        <span v-if="isUploadingImage" class="spinner-border spinner-border-sm me-1"></span>
                        <i v-else class="bi bi-cloud-upload me-1"></i>
                        {{ isUploadingImage ? '上传中...' : '选择图片文件' }}
                      </button>
                      <small class="text-muted ms-2">支持: JPG, PNG, GIF, WebP, SVG, BMP (最大10MB)</small>
                    </div>

                    <!-- 图片预览 -->
                    <div v-if="nftMetadata.image" class="mb-2">
                      <div class="image-preview-container">
                        <img :src="getImageDisplayUrl(nftMetadata.image)" 
                          alt="NFT图片预览" 
                          class="image-preview"
                          @error="handleImageError">
                        <button type="button" class="btn btn-danger btn-sm position-absolute top-0 end-0 m-1"
                          @click="removeImage" title="移除图片">
                          <i class="bi bi-x"></i>
                        </button>
                      </div>
                    </div>

                    <!-- 手动输入URL -->
                    <div class="mb-2">
                      <input type="url" class="form-control form-control-sm" 
                        v-model="nftMetadata.image"
                        placeholder="或手动输入图片URL: https://example.com/image.jpg 或 ipfs://...">
                    </div>

                    <div class="form-text">
                      <i class="bi bi-info-circle me-1"></i>
                      可以上传图片文件或手动输入图片URL。上传的图片将自动存储到IPFS网络。
                    </div>
                  </div>
                </div>
              </div>

              <button type="submit" class="btn btn-primary" :disabled="isSubmitting || !walletAddress">
                <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-1" role="status"
                  aria-hidden="true"></span>
                <i v-else class="bi bi-lock me-1"></i>
                {{ isSubmitting ? '加密中...' : '加密数据' }}
              </button>
            </form>
          </div>
        </div>
      </div>

      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h5 class="mb-0"><i class="bi bi-file-lock me-2"></i>加密结果</h5>
          </div>
          <div class="card-body">
            <div v-if="!encryptResult" class="text-center py-4 text-muted">
              <i class="bi bi-shield-lock" style="font-size: 3rem;"></i>
              <p class="mt-2">填写表单并点击加密按钮以开始加密</p>
            </div>

            <div v-else>
              <div class="alert alert-success">
                <h6><i class="bi bi-check-circle me-2"></i>
                  {{ encryptResult.isFullWorkflow ? 'IPFS+NFT工作流程完成！' : '加密成功！' }}
                </h6>
                <p class="mb-0"><strong>密文ID:</strong> {{ encryptResult.ciphertext_id }}</p>
                <p class="mb-0"><strong>访问策略:</strong> {{ encryptResult.policy }}</p>
                <p class="mb-0"><strong>创建时间:</strong> {{ formatDate(encryptResult.created_at) }}</p>
              </div>

              <!-- 完整工作流程结果 -->
              <div v-if="encryptResult.isFullWorkflow" class="mb-3">
                <div class="alert alert-info">
                  <h6><i class="bi bi-diagram-3 me-2"></i>工作流程详情</h6>
                  <div class="row">
                    <div class="col-md-6">
                      <p class="mb-1"><strong>1. 原文IPFS Hash:</strong></p>
                      <div class="bg-light p-2 rounded text-monospace small">{{ encryptResult.originalFileHash }}</div>
                    </div>
                    <div class="col-md-6">
                      <p class="mb-1"><strong>2. 密文IPFS Hash:</strong></p>
                      <div class="bg-light p-2 rounded text-monospace small">{{ encryptResult.cipherFileHash }}</div>
                    </div>
                  </div>
                  <div class="row mt-2">
                    <div class="col-md-6">
                      <p class="mb-1"><strong>3. NFT元数据Hash:</strong></p>
                      <div class="bg-light p-2 rounded text-monospace small">{{ encryptResult.nftMetadataHash }}</div>
                    </div>
                    <div class="col-md-6">
                      <p class="mb-1"><strong>4. NFT交易Hash:</strong></p>
                      <div class="bg-light p-2 rounded text-monospace small">{{ encryptResult.nftTransactionHash || '待铸造' }}</div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="mb-3">
                <label class="form-label"><strong>加密密文:</strong></label>
                <textarea class="form-control" rows="4" readonly v-model="encryptResult.cipher"
                  @click="$event.target.select()"></textarea>
                <div class="d-flex gap-2 mt-2 flex-wrap">
                  <button class="btn btn-sm btn-outline-primary" @click="copyToClipboard(encryptResult.cipher)">
                    <i class="bi bi-clipboard me-1"></i>复制密文
                  </button>
                  <button class="btn btn-sm btn-outline-success" @click="useForDecrypt">
                    <i class="bi bi-arrow-right me-1"></i>去解密
                  </button>
                  <button v-if="encryptResult.isFullWorkflow" class="btn btn-sm btn-outline-info" @click="viewOnIPFS(encryptResult.originalFileHash)">
                    <i class="bi bi-box-arrow-up-right me-1"></i>查看原文
                  </button>
                  <button v-if="encryptResult.isFullWorkflow" class="btn btn-sm btn-outline-warning" @click="viewOnIPFS(encryptResult.cipherFileHash)">
                    <i class="bi bi-box-arrow-up-right me-1"></i>查看密文文件
                  </button>
                </div>
              </div>

              <div class="alert alert-info">
                <i class="bi bi-info-circle me-2"></i>
                <small>
                  {{ encryptResult.isFullWorkflow ? 
                    '数据已完成完整的IPFS存储和NFT创建流程。NFT元数据包含密文的IPFS地址。' : 
                    '数据已成功加密。只有当前钱包地址才能解密此数据。' 
                  }}
                </small>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 钱包信息说明 -->
    <div class="card mt-4">
      <div class="card-header">
        <h5 class="mb-0"><i class="bi bi-wallet2 me-2"></i>访问控制说明</h5>
      </div>
      <div class="card-body">
        <div class="row">
          <div class="col-md-6">
            <h6>当前钱包地址：</h6>
            <p class="text-monospace bg-light p-2 rounded">{{ walletAddress || '未连接' }}</p>
          </div>
          <div class="col-md-6">
            <h6>当前访问策略：</h6>
            <p class="text-monospace bg-light p-2 rounded">{{ currentPolicy }}</p>
          </div>
        </div>
        <div class="row">
          <div class="col-md-12">
            <h6>策略说明：</h6>
            <div class="alert alert-info">
              <i class="bi bi-info-circle me-2"></i>
              <small>
                使用基于属性的加密(ABE)技术，数据加密采用固定策略模式。
                访问策略固定为 "mainNFT:当前钱包地址"，只有当前连接的钱包地址才能解密数据。
                这确保了数据的安全性和访问控制的简单性。
              </small>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, reactive, computed } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import ipfsService from '@/services/ipfsService'
import nftService from '@/services/nftService'
import abeService from '@/services/abeService'

export default {
  name: 'ABEEncrypt',
  setup() {
    const store = useStore()
    const router = useRouter()

    const isSubmitting = ref(false)
    const encryptResult = ref(null)
    const isUploadingImage = ref(false)

    const encryptForm = reactive({
      message: '',
      policy: '',
      saveToIPFS: false
    })

    const nftMetadata = reactive({
      name: '',
      description: '',
      image: ''
    })





    // 获取钱包地址
    const walletAddress = computed(() => store.getters['wallet/currentAccount'])

    // 当前策略（固定格式）
    const currentPolicy = computed(() => {
      if (!walletAddress.value) {
        return '请先连接钱包'
      }
      return `mainNFT:${walletAddress.value}`
    })

    // 加密数据
    const encryptData = async () => {
      if (!encryptForm.message.trim()) {
        store.dispatch('app/showError', '请填写要加密的数据')
        return
      }

      if (!walletAddress.value) {
        store.dispatch('app/showError', '请先连接钱包')
        return
      }

      // 自动设置固定策略
      encryptForm.policy = currentPolicy.value

      // 如果启用了IPFS+NFT工作流程，验证NFT元数据
      if (encryptForm.saveToIPFS) {
        if (!nftMetadata.name.trim() || !nftMetadata.description.trim()) {
          store.dispatch('app/showError', '请填写完整的NFT元数据信息')
          return
        }
      }

      console.log('使用的ABE策略:', encryptForm.policy)

      isSubmitting.value = true
      try {
        if (encryptForm.saveToIPFS) {
          // 执行完整的IPFS+NFT工作流程
          await executeFullWorkflow()
        } else {
          // 只执行简单加密
          await executeSimpleEncryption()
        }
      } catch (error) {
        console.error('处理失败:', error)
        store.dispatch('app/showError', '处理失败: ' + error.message)
      } finally {
        isSubmitting.value = false
      }
    }

    // 执行简单加密
    const executeSimpleEncryption = async () => {
      console.log('开始简单加密流程')

      const result = await store.dispatch('abe/encryptData', {
        message: encryptForm.message,
        policy: encryptForm.policy,
        saveToNFT: false
      })

      if (result) {
        encryptResult.value = {
          ciphertext_id: result.ciphertext_id || `cipher_${Date.now()}`,
          cipher: result.cipher,
          policy: encryptForm.policy,
          created_at: new Date().toISOString(),
          isFullWorkflow: false
        }

        store.dispatch('app/showSuccess', '数据加密成功！')
      }
    }

    // 执行完整的IPFS+NFT工作流程
    const executeFullWorkflow = async () => {
      console.log('开始完整IPFS+NFT工作流程')
      
      let originalFileHash, cipherFileHash, nftMetadataHash

      try {
        // 步骤1: 上传原文到IPFS
        store.dispatch('app/showInfo', '步骤1/4: 上传原文到IPFS...')
        const originalFileResult = await ipfsService.uploadToIPFS(encryptForm.message, 'original_data.txt')
        originalFileHash = originalFileResult.hash
        console.log('原文IPFS Hash:', originalFileHash)

        // 步骤2: 加密文件Hash
        store.dispatch('app/showInfo', '步骤2/4: 加密文件Hash...')
        const encryptResult = await store.dispatch('abe/encryptData', {
          message: originalFileHash,  // 加密的是文件Hash，而不是原文
          policy: encryptForm.policy,
          saveToNFT: false
        })

        if (!encryptResult || !encryptResult.cipher) {
          throw new Error('文件Hash加密失败')
        }

        // 步骤3: 上传密文到IPFS
        store.dispatch('app/showInfo', '步骤3/4: 上传密文到IPFS...')
        const cipherFileResult = await ipfsService.uploadToIPFS(encryptResult.cipher, 'encrypted_hash.txt')
        cipherFileHash = cipherFileResult.hash
        console.log('密文IPFS Hash:', cipherFileHash)

        // 步骤4: 创建NFT元数据并上传到IPFS
        store.dispatch('app/showInfo', '步骤4/4: 创建NFT元数据...')
        const metadata = {
          name: nftMetadata.name,
          description: nftMetadata.description,
          external_url: '',
          image: nftMetadata.image || '',
          policy: encryptForm.policy,
          ciphertext: `ipfs://${cipherFileHash}` // 存储密文的IPFS地址
        }

        const metadataResult = await nftService.createMetadata(metadata)
        nftMetadataHash = metadataResult.ipfs_hash
        console.log('NFT元数据IPFS Hash:', nftMetadataHash)

        // 成功完成所有步骤
        encryptResult.value = {
          ciphertext_id: `cipher_${Date.now()}`,
          cipher: encryptResult.cipher,
          policy: encryptForm.policy,
          created_at: new Date().toISOString(),
          isFullWorkflow: true,
          originalFileHash,
          cipherFileHash,
          nftMetadataHash,
          nftTransactionHash: null // 稍后可以添加铸造NFT的功能
        }

        store.dispatch('app/showSuccess', 'IPFS+NFT工作流程完成！')

      } catch (error) {
        console.error('工作流程失败:', error)
        // 提供详细的错误信息
        let errorMsg = '工作流程执行失败: ' + error.message
        if (originalFileHash) {
          errorMsg += `\n已完成步骤：原文上传(${originalFileHash})`
        }
        if (cipherFileHash) {
          errorMsg += `，密文上传(${cipherFileHash})`
        }
        if (nftMetadataHash) {
          errorMsg += `，元数据创建(${nftMetadataHash})`
        }
        throw new Error(errorMsg)
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
      // 将密文存储到store中，供解密页面使用
      store.commit('abe/setLatestCiphertext', encryptResult.value.cipher)
      router.push({ name: 'ABEDecrypt' })
    }

    // 在IPFS上查看文件
    const viewOnIPFS = (hash) => {
      if (hash) {
        const url = ipfsService.getIPFSUrl(hash)
        window.open(url, '_blank')
      }
    }

    // 格式化日期
    const formatDate = (dateString) => {
      if (!dateString) return ''
      const date = new Date(dateString)
      return date.toLocaleString()
    }

    // 处理图片选择
    const handleImageSelect = async (event) => {
      const file = event.target.files[0]
      if (!file) return

      isUploadingImage.value = true
      try {
        console.log('开始上传图片:', file.name)
        const result = await abeService.uploadImage(file)
        
        // 使用主要的HTTP URL作为图片地址
        nftMetadata.image = result.primary_url
        
        store.dispatch('app/showSuccess', `图片上传成功！IPFS Hash: ${result.hash}`)
        console.log('图片上传结果:', result)
      } catch (error) {
        console.error('图片上传失败:', error)
        store.dispatch('app/showError', error.message)
      } finally {
        isUploadingImage.value = false
        // 清空input的value，允许重新选择同一个文件
        event.target.value = ''
      }
    }

    // 移除图片
    const removeImage = () => {
      nftMetadata.image = ''
      store.dispatch('app/showSuccess', '图片已移除')
    }

    // 获取图片显示URL
    const getImageDisplayUrl = (imageUrl) => {
      if (!imageUrl) return ''
      
      // 如果是IPFS链接，转换为HTTP网关链接
      if (imageUrl.startsWith('ipfs://')) {
        const hash = imageUrl.replace('ipfs://', '')
        return `https://dweb.link/ipfs/${hash}`
      }
      
      return imageUrl
    }

    // 处理图片加载错误
    const handleImageError = (event) => {
      console.error('图片加载失败:', event.target.src)
      if (!event.target.src.includes('placeholder')) {
        // 使用占位图片
        event.target.src = 'data:image/svg+xml;charset=UTF-8,%3Csvg%20width%3D%22200%22%20height%3D%22120%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22%23f0f0f0%22%2F%3E%3Ctext%20x%3D%2250%25%22%20y%3D%2250%25%22%20font-family%3D%22Arial%22%20font-size%3D%2214%22%20text-anchor%3D%22middle%22%20dominant-baseline%3D%22middle%22%20fill%3D%22%23999%22%3E%E5%9B%BE%E7%89%87%E5%8A%A0%E8%BD%BD%E5%A4%B1%E8%B4%A5%3C%2Ftext%3E%3C%2Fsvg%3E'
        event.target.onerror = null
      }
    }



    return {
      encryptForm,
      nftMetadata,
      isSubmitting,
      isUploadingImage,
      encryptResult,
      walletAddress,
      currentPolicy,
      encryptData,
      executeSimpleEncryption,
      executeFullWorkflow,
      copyToClipboard,
      useForDecrypt,
      viewOnIPFS,
      formatDate,
      handleImageSelect,
      removeImage,
      getImageDisplayUrl,
      handleImageError,

    }
  }
}
</script>

<style scoped>
.card {
  height: 100%;
}

.form-text {
  font-size: 0.875rem;
}

.alert {
  border: none;
  border-radius: 0.5rem;
}

code {
  background-color: #f8f9fa;
  padding: 0.2rem 0.4rem;
  border-radius: 0.25rem;
  font-size: 0.875rem;
}

.btn-sm {
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

.image-upload-section {
  position: relative;
}

.image-preview-container {
  position: relative;
  display: inline-block;
  max-width: 200px;
}

.image-preview {
  max-width: 100%;
  max-height: 200px;
  border-radius: 0.375rem;
  border: 1px solid #dee2e6;
}

</style> 