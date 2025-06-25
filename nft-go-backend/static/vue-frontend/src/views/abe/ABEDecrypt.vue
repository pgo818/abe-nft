<template>
  <div class="abe-decrypt">
    <h2><i class="bi bi-unlock me-2"></i>ABE数据解密</h2>

    <div class="row mt-4">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h5 class="mb-0"><i class="bi bi-key me-2"></i>解密表单</h5>
          </div>
          <div class="card-body">
            <form @submit.prevent="decryptData">
              <div class="mb-3">
                <label for="ciphertext" class="form-label">密文数据</label>
                <textarea class="form-control" id="ciphertext" v-model="decryptForm.ciphertext" rows="4"
                  placeholder="粘贴需要解密的密文数据" required></textarea>
                <div class="d-flex gap-2 mt-2">
                  <button type="button" class="btn btn-sm btn-outline-secondary" @click="useLatestCiphertext"
                    :disabled="!latestCiphertext">
                    <i class="bi bi-arrow-down me-1"></i>使用最新密文
                  </button>
                  <button type="button" class="btn btn-sm btn-outline-secondary" @click="pasteFromClipboard">
                    <i class="bi bi-clipboard me-1"></i>从剪贴板粘贴
                  </button>
                </div>
              </div>

              <div class="mb-3">
                <label for="userKey" class="form-label">用户密钥</label>
                <textarea class="form-control" id="userKey" v-model="decryptForm.userKey" rows="4"
                  placeholder="粘贴您的ABE用户密钥" required></textarea>
                <div class="d-flex gap-2 mt-2 flex-wrap">
                  <button type="button" class="btn btn-sm btn-outline-secondary" @click="useLatestUserKey"
                    :disabled="!latestUserKey">
                    <i class="bi bi-arrow-down me-1"></i>使用最新密钥
                  </button>
                  <button type="button" class="btn btn-sm btn-outline-secondary" @click="pasteKeyFromClipboard">
                    <i class="bi bi-clipboard me-1"></i>从剪贴板粘贴
                  </button>
                  <button type="button" class="btn btn-sm btn-outline-success" @click="saveCurrentKey"
                    :disabled="!decryptForm.userKey.trim()">
                    <i class="bi bi-save me-1"></i>保存当前密钥
                  </button>
                  <button type="button" class="btn btn-sm btn-outline-info" @click="showKeyManager = !showKeyManager">
                    <i class="bi bi-key me-1"></i>密钥管理 ({{ savedKeys.length }})
                  </button>
                </div>
                <div class="form-text">输入您在ABE密钥生成页面获得的用户密钥</div>
              </div>

              <!-- 密钥管理器 -->
              <div v-if="showKeyManager" class="card mb-3">
                <div class="card-header">
                  <h6 class="mb-0"><i class="bi bi-collection me-2"></i>保存的密钥</h6>
                </div>
                <div class="card-body">
                  <div v-if="savedKeys.length === 0" class="text-center py-3 text-muted">
                    <i class="bi bi-inbox" style="font-size: 2rem;"></i>
                    <p class="mt-2 mb-0">没有保存的密钥</p>
                  </div>
                  <div v-else>
                    <div v-for="key in savedKeys" :key="key.id" class="border rounded p-3 mb-2">
                      <div class="d-flex justify-content-between align-items-start">
                        <div class="flex-grow-1">
                          <h6 class="mb-1">{{ key.name }}</h6>
                          <small class="text-muted">
                            <i class="bi bi-calendar me-1"></i>{{ formatDate(key.created_at) }}
                            <br>
                            <i class="bi bi-wallet me-1"></i>{{ key.wallet_address }}
                            <br>
                            <i class="bi bi-tags me-1"></i>{{ key.attributes.join(', ') }}
                          </small>
                        </div>
                        <div class="d-flex gap-1">
                          <button class="btn btn-sm btn-outline-primary" @click="useKey(key)">
                            <i class="bi bi-arrow-down"></i>使用
                          </button>
                          <button class="btn btn-sm btn-outline-secondary" @click="editKeyName(key)">
                            <i class="bi bi-pencil"></i>改名
                          </button>
                          <button class="btn btn-sm btn-outline-danger" @click="deleteKey(key.id)">
                            <i class="bi bi-trash"></i>删除
                          </button>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
                <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-1" role="status"
                  aria-hidden="true"></span>
                <i v-else class="bi bi-unlock me-1"></i>
                {{ isSubmitting ? '解密中...' : '解密数据' }}
              </button>
            </form>
          </div>
        </div>
      </div>

      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h5 class="mb-0"><i class="bi bi-file-text me-2"></i>解密结果</h5>
          </div>
          <div class="card-body">
            <div v-if="!decryptResult && !decryptError" class="text-center py-4 text-muted">
              <i class="bi bi-shield-unlock" style="font-size: 3rem;"></i>
              <p class="mt-2">填写表单并点击解密按钮以开始解密</p>
            </div>

            <!-- 解密成功 -->
            <div v-if="decryptResult">
              <div class="alert alert-success">
                <h6><i class="bi bi-check-circle me-2"></i>解密成功！</h6>
                <p class="mb-0"><strong>解密时间:</strong> {{ formatDate(decryptResult.decrypted_at) }}</p>
                <p class="mb-0"><strong>数据长度:</strong> {{ decryptResult.message.length }} 字符</p>
              </div>

              <div class="mb-3">
                <label class="form-label"><strong>解密后的数据:</strong></label>
                <textarea class="form-control" rows="5" readonly v-model="decryptResult.message"
                  @click="$event.target.select()"></textarea>
                <div class="d-flex gap-2 mt-2">
                  <button class="btn btn-sm btn-outline-primary" @click="copyToClipboard(decryptResult.message)">
                    <i class="bi bi-clipboard me-1"></i>复制内容
                  </button>
                  <button class="btn btn-sm btn-outline-success" @click="saveToFile">
                    <i class="bi bi-download me-1"></i>保存文件
                  </button>
                </div>
              </div>

              <div class="alert alert-info">
                <i class="bi bi-info-circle me-2"></i>
                <small>您的用户密钥满足访问策略要求，解密成功。</small>
              </div>
            </div>

            <!-- 解密失败 -->
            <div v-if="decryptError">
              <div class="alert alert-danger">
                <h6><i class="bi bi-x-circle me-2"></i>解密失败</h6>
                <p class="mb-0">{{ decryptError.message }}</p>
              </div>

              <div class="alert alert-warning">
                <i class="bi bi-exclamation-triangle me-2"></i>
                <small>可能的原因：</small>
                <ul class="mb-0 mt-1 small">
                  <li>您的用户密钥不满足访问策略要求</li>
                  <li>密文格式错误或已损坏</li>
                  <li>用户密钥格式错误或已损坏</li>
                  <li>密钥与密文不匹配</li>
                </ul>
              </div>

              <button class="btn btn-sm btn-outline-secondary" @click="clearError">
                <i class="bi bi-arrow-clockwise me-1"></i>重新尝试
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>


  </div>
</template>

<script>
import { ref, reactive, computed, onMounted } from 'vue'
import { useStore } from 'vuex'

export default {
  name: 'ABEDecrypt',
  setup() {
    const store = useStore()

    const isSubmitting = ref(false)
    const decryptResult = ref(null)
    const decryptError = ref(null)
    const showKeyManager = ref(false)

    const decryptForm = reactive({
      ciphertext: '',
      userKey: ''
    })

    // 从store获取最新的密文和用户密钥
    const latestCiphertext = computed(() => store.state.abe?.latestCiphertext)
    const latestUserKey = computed(() => store.state.abe?.latestUserKey)
    const savedKeys = computed(() => store.getters['abe/savedUserKeys'])

    onMounted(() => {
      // 加载保存的密钥
      store.dispatch('abe/loadSavedKeys')
      
      // 如果有最新密文和用户密钥，自动填入
      if (latestCiphertext.value) {
        decryptForm.ciphertext = latestCiphertext.value
      }
      if (latestUserKey.value) {
        decryptForm.userKey = latestUserKey.value
      }
    })

    // 解密数据
    const decryptData = async () => {
      if (!decryptForm.ciphertext.trim() || !decryptForm.userKey.trim()) {
        store.dispatch('app/showError', '请填写密文和用户密钥')
        return
      }

      // 清除之前的结果和错误
      decryptResult.value = null
      decryptError.value = null

      isSubmitting.value = true
      try {
        console.log('开始解密数据:', {
          ciphertextLength: decryptForm.ciphertext.length,
          userKeyLength: decryptForm.userKey.length
        })

        // 调用store中的直接解密action
        const result = await store.dispatch('abe/decryptDataDirect', {
          cipher: decryptForm.ciphertext,
          attrib_keys: decryptForm.userKey
        })

        if (result && result.message) {
          decryptResult.value = {
            message: result.message,
            decrypted_at: new Date().toISOString()
          }

          store.dispatch('app/showSuccess', '数据解密成功！')
        } else {
          throw new Error('解密返回结果为空')
        }
      } catch (error) {
        console.error('解密失败:', error)
        decryptError.value = {
          message: error.message || '解密过程中发生未知错误'
        }
        store.dispatch('app/showError', '解密失败: ' + error.message)
      } finally {
        isSubmitting.value = false
      }
    }

    // 使用最新密文
    const useLatestCiphertext = () => {
      if (latestCiphertext.value) {
        decryptForm.ciphertext = latestCiphertext.value
        store.dispatch('app/showSuccess', '已填入最新密文')
      } else {
        store.dispatch('app/showError', '没有可用的最新密文')
      }
    }

    // 从剪贴板粘贴密文
    const pasteFromClipboard = async () => {
      try {
        const text = await navigator.clipboard.readText()
        if (text) {
          decryptForm.ciphertext = text
          store.dispatch('app/showSuccess', '已从剪贴板粘贴密文')
        } else {
          store.dispatch('app/showError', '剪贴板中没有内容')
        }
      } catch (error) {
        console.error('从剪贴板读取失败:', error)
        store.dispatch('app/showError', '无法从剪贴板读取内容，请手动粘贴')
      }
    }

    // 使用最新用户密钥
    const useLatestUserKey = () => {
      if (latestUserKey.value) {
        decryptForm.userKey = latestUserKey.value
        store.dispatch('app/showSuccess', '已填入最新用户密钥')
      } else {
        store.dispatch('app/showError', '没有可用的最新用户密钥')
      }
    }

    // 从剪贴板粘贴用户密钥
    const pasteKeyFromClipboard = async () => {
      try {
        const text = await navigator.clipboard.readText()
        if (text) {
          decryptForm.userKey = text
          store.dispatch('app/showSuccess', '已从剪贴板粘贴用户密钥')
        } else {
          store.dispatch('app/showError', '剪贴板中没有内容')
        }
      } catch (error) {
        console.error('从剪贴板读取失败:', error)
        store.dispatch('app/showError', '无法从剪贴板读取内容，请手动粘贴')
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

    // 保存到文件
    const saveToFile = () => {
      try {
        const blob = new Blob([decryptResult.value.message], { type: 'text/plain' })
        const url = URL.createObjectURL(blob)
        const a = document.createElement('a')
        a.href = url
        a.download = `decrypted_data_${Date.now()}.txt`
        document.body.appendChild(a)
        a.click()
        document.body.removeChild(a)
        URL.revokeObjectURL(url)
        store.dispatch('app/showSuccess', '文件已保存')
      } catch (error) {
        console.error('保存文件失败:', error)
        store.dispatch('app/showError', '保存文件失败')
      }
    }



    // 保存当前密钥
    const saveCurrentKey = () => {
      if (!decryptForm.userKey.trim()) {
        store.dispatch('app/showError', '请先输入用户密钥')
        return
      }

      const keyName = prompt('请输入密钥名称:', `密钥-${new Date().toLocaleString()}`)
      if (!keyName) return

      const keyData = {
        id: Date.now().toString(), // 简单的ID生成
        name: keyName,
        wallet_address: store.getters['wallet/currentAccount'] || '未知地址',
        attributes: ['手动保存'],
        attrib_keys: decryptForm.userKey,
        created_at: new Date().toISOString()
      }

      store.dispatch('abe/saveUserKey', keyData)
      store.dispatch('app/showSuccess', '密钥已保存')
    }

    // 使用保存的密钥
    const useKey = (key) => {
      decryptForm.userKey = key.attrib_keys
      store.dispatch('app/showSuccess', `已加载密钥: ${key.name}`)
    }

    // 编辑密钥名称
    const editKeyName = (key) => {
      const newName = prompt('请输入新的密钥名称:', key.name)
      if (!newName || newName === key.name) return

      const updatedKey = { ...key, name: newName }
      store.dispatch('abe/saveUserKey', updatedKey)
      store.dispatch('app/showSuccess', '密钥名称已更新')
    }

    // 删除密钥
    const deleteKey = (keyId) => {
      if (!confirm('确定要删除这个密钥吗？')) return

      store.dispatch('abe/deleteSavedKey', keyId)
      store.dispatch('app/showSuccess', '密钥已删除')
    }

    // 清除错误
    const clearError = () => {
      decryptError.value = null
    }

    // 格式化日期
    const formatDate = (dateString) => {
      if (!dateString) return ''
      const date = new Date(dateString)
      return date.toLocaleString()
    }

    return {
      isSubmitting,
      decryptResult,
      decryptError,
      decryptForm,
      showKeyManager,
      latestCiphertext,
      latestUserKey,
      savedKeys,
      decryptData,
      useLatestCiphertext,
      useLatestUserKey,
      pasteFromClipboard,
      pasteKeyFromClipboard,
      saveCurrentKey,
      useKey,
      editKeyName,
      deleteKey,
      copyToClipboard,
      saveToFile,
      clearError,
      formatDate
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

.list-unstyled li {
  padding: 0.2rem 0;
}
</style> 