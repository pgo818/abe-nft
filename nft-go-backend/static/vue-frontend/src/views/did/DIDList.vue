<template>
  <div class="container mt-4">
    <div class="d-flex justify-content-between align-items-center mb-4">
      <h2><i class="bi bi-person-badge me-2"></i>我的DID信息</h2>
      <button class="btn btn-outline-primary" @click="refreshData">
        <i class="bi bi-arrow-clockwise me-1"></i>刷新
      </button>
    </div>

    <!-- 钱包连接状态 -->
    <div v-if="!isWalletConnected" class="alert alert-warning">
      <i class="bi bi-exclamation-triangle me-2"></i>
      请先连接钱包以查看您的DID信息
    </div>

    <!-- 当前钱包信息 -->
    <div v-if="isWalletConnected" class="card mb-4">
      <div class="card-body">
        <h5 class="card-title">
          <i class="bi bi-wallet2 me-2"></i>当前钱包
        </h5>
        <p class="card-text">
          <strong>地址：</strong>
          <span class="text-monospace">{{ currentAccount }}</span>
          <button class="btn btn-sm btn-outline-secondary ms-2" @click="copyToClipboard(currentAccount)">
            <i class="bi bi-clipboard"></i>
          </button>
        </p>
      </div>
    </div>

    <div v-if="loading" class="text-center my-5">
      <div class="spinner-border text-primary" role="status">
        <span class="visually-hidden">加载中...</span>
      </div>
      <p class="mt-2">加载DID信息...</p>
    </div>

    <!-- DID信息展示 -->
    <div v-else-if="isWalletConnected">
      <!-- 医生DID信息 -->
      <div v-if="doctorDID" class="card mb-4">
        <div class="card-header bg-success text-white">
          <h5 class="mb-0">
            <i class="bi bi-person-badge me-2"></i>医生DID信息
          </h5>
        </div>
        <div class="card-body">
          <div class="row">
            <div class="col-md-6">
              <div class="mb-3">
                <label class="form-label fw-bold">DID标识符</label>
                <div class="input-group">
                  <input type="text" class="form-control" :value="doctorDID.didString" readonly>
                  <button class="btn btn-outline-secondary" @click="copyToClipboard(doctorDID.didString)">
                    <i class="bi bi-clipboard"></i>
                  </button>
                </div>
              </div>
              <div class="mb-3">
                <label class="form-label fw-bold">医生姓名</label>
                <p class="mb-0">{{ doctorDID.name }}</p>
              </div>
              <div class="mb-3">
                <label class="form-label fw-bold">执业编号</label>
                <p class="mb-0">{{ doctorDID.licenseNumber }}</p>
              </div>
            </div>
            <div class="col-md-6">
              <div class="mb-3">
                <label class="form-label fw-bold">创建时间</label>
                <p class="mb-0">{{ formatDate(doctorDID.createdAt) }}</p>
              </div>
              <div class="mb-3">
                <label class="form-label fw-bold">状态</label>
                <span class="badge bg-success">{{ doctorDID.status }}</span>
              </div>
              <div class="mb-3">
                <label class="form-label fw-bold">关联凭证</label>
                <p class="mb-0">{{ doctorVCs.length }} 个</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 医生可验证凭证 -->
      <div v-if="doctorVCs.length > 0" class="card mb-4">
        <div class="card-header">
          <h5 class="mb-0">
            <i class="bi bi-card-checklist me-2"></i>医生可验证凭证 ({{ doctorVCs.length }})
          </h5>
        </div>
        <div class="card-body">
          <div class="row">
            <div class="col-md-6 mb-3" v-for="vc in doctorVCs" :key="vc.id">
              <div class="card border-primary">
                <div class="card-header bg-primary text-white">
                  <h6 class="mb-0">{{ vc.type }}</h6>
                </div>
                <div class="card-body">
                  <div class="mb-2">
                    <strong>凭证ID：</strong>
                    <span class="text-monospace small">{{ vc.vcid }}</span>
                    <button class="btn btn-sm btn-outline-secondary ms-1" @click="copyToClipboard(vc.vcid)">
                      <i class="bi bi-clipboard"></i>
                    </button>
                  </div>
                  <div class="mb-2">
                    <strong>颁发者：</strong> {{ vc.issuerDID }}
                  </div>
                  <div class="mb-2">
                    <strong>颁发时间：</strong> {{ formatDate(vc.issuedAt) }}
                  </div>
                  <div class="mb-2">
                    <strong>过期时间：</strong> {{ formatDate(vc.expiresAt) }}
                  </div>
                  <div class="mb-2">
                    <strong>状态：</strong>
                    <span :class="['badge', vc.status === 'active' ? 'bg-success' : 'bg-warning']">
                      {{ vc.status === 'active' ? '有效' : '无效' }}
                    </span>
                  </div>
                  <div v-if="vc.content" class="mb-2">
                    <strong>内容：</strong>
                    <details>
                      <summary class="text-primary" style="cursor: pointer;">点击查看详情</summary>
                      <pre class="mt-2 bg-light p-2 rounded small">{{ formatVCContent(vc.content) }}</pre>
                    </details>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- 普通DID列表 -->
      <div v-if="normalDIDs.length > 0" class="card mb-4">
        <div class="card-header">
          <h5 class="mb-0">
            <i class="bi bi-list-ul me-2"></i>其他DID信息 ({{ normalDIDs.length }})
          </h5>
        </div>
        <div class="card-body">
          <div class="table-responsive">
            <table class="table">
              <thead>
                <tr>
                  <th>DID标识符</th>
                  <th>创建时间</th>
                  <th>状态</th>
                  <th>操作</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="did in normalDIDs" :key="did.id">
                  <td>
                    <span class="text-monospace">{{ truncateDID(did.didString) }}</span>
                    <button class="btn btn-sm btn-outline-secondary ms-1" @click="copyToClipboard(did.didString)">
                      <i class="bi bi-clipboard"></i>
                    </button>
                  </td>
                  <td>{{ formatDate(did.createdAt) }}</td>
                  <td>
                    <span :class="['badge', did.status === 'active' ? 'bg-success' : 'bg-secondary']">
                      {{ did.status === 'active' ? '活跃' : '非活跃' }}
                    </span>
                  </td>
                  <td>
                    <button class="btn btn-sm btn-outline-primary" @click="viewDIDDetails(did)">
                      <i class="bi bi-eye"></i> 详情
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- 无DID信息时的提示 -->
      <div v-if="!doctorDID && normalDIDs.length === 0 && !loading" class="text-center my-5">
        <i class="bi bi-person-badge fs-1 text-muted"></i>
        <p class="lead mt-3">您还没有创建任何DID</p>
        <p class="text-muted">DID（去中心化身份标识符）可以帮助您管理数字身份</p>
        <router-link to="/did/doctor" class="btn btn-primary mt-2">
          <i class="bi bi-hospital me-1"></i>创建医生DID
        </router-link>
      </div>
    </div>

    <!-- DID详情模态框 -->
    <div class="modal fade" ref="detailsModal" tabindex="-1">
      <div class="modal-dialog modal-lg">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">DID详情</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body" v-if="selectedDID">
            <div class="mb-3">
              <label class="form-label fw-bold">DID标识符</label>
              <div class="input-group">
                <input type="text" class="form-control" :value="selectedDID.didString" readonly>
                <button class="btn btn-outline-secondary" @click="copyToClipboard(selectedDID.didString)">
                  <i class="bi bi-clipboard"></i>
                </button>
              </div>
            </div>

            <div class="mb-3">
              <label class="form-label fw-bold">钱包地址</label>
              <p>{{ selectedDID.walletAddress }}</p>
            </div>

            <div class="mb-3">
              <label class="form-label fw-bold">创建时间</label>
              <p>{{ formatDate(selectedDID.createdAt) }}</p>
            </div>

            <div class="mb-3">
              <label class="form-label fw-bold">状态</label>
              <span :class="['badge', selectedDID.status === 'active' ? 'bg-success' : 'bg-secondary']">
                {{ selectedDID.status === 'active' ? '活跃' : '非活跃' }}
              </span>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useStore } from 'vuex'
import { Modal } from 'bootstrap'

export default {
  name: 'DIDList',
  
  setup() {
    const store = useStore()
    const detailsModal = ref(null)
    let bsDetailsModal = null

    // 状态
    const loading = ref(false)
    const allDIDs = ref([])
    const doctorVCs = ref([])
    const selectedDID = ref(null)

    // 计算属性
    const isWalletConnected = computed(() => store.state.wallet.isConnected)
    const currentAccount = computed(() => store.state.wallet.account)

    // 分离医生DID和普通DID
    const doctorDID = computed(() => {
      if (!Array.isArray(allDIDs.value)) return null
      return allDIDs.value.find(did => 
        did.didString && did.didString.includes('ethr') && 
        did.walletAddress === currentAccount.value
      )
    })

    const normalDIDs = computed(() => {
      if (!Array.isArray(allDIDs.value)) return []
      return allDIDs.value.filter(did => 
        !(did.didString && did.didString.includes('ethr') && 
          did.walletAddress === currentAccount.value)
      )
    })

    // 方法
    const loadData = async () => {
      console.log('loadData 开始执行')
      console.log('isWalletConnected:', isWalletConnected.value)
      console.log('currentAccount:', currentAccount.value)
      console.log('store.state.wallet:', store.state.wallet)
      
      if (!isWalletConnected.value) {
        console.log('钱包未连接，跳过加载数据')
        return
      }

      loading.value = true
      try {
        // 获取钱包相关的DID列表
        console.log('开始获取钱包相关的DID列表...')
        const dids = await store.dispatch('did/getDIDsByWallet')
        console.log('获取到的DID列表:', dids, '类型:', typeof dids, '是否为数组:', Array.isArray(dids))
        
        // 确保dids是数组类型
        if (Array.isArray(dids)) {
          allDIDs.value = dids
        } else {
          console.warn('返回的DID数据不是数组格式:', dids)
          allDIDs.value = []
        }

        // 如果有医生DID，获取其凭证
        if (doctorDID.value) {
          console.log('找到医生DID，开始获取凭证:', doctorDID.value.didString)
          try {
            const vcs = await store.dispatch('did/getDoctorVCs', doctorDID.value.didString)
            console.log('获取到的医生凭证:', vcs, '类型:', typeof vcs, '是否为数组:', Array.isArray(vcs))
            
            // 确保vcs是数组类型
            if (Array.isArray(vcs)) {
              doctorVCs.value = vcs
            } else {
              console.warn('返回的凭证数据不是数组格式:', vcs)
              doctorVCs.value = []
            }
          } catch (error) {
            console.warn('获取医生凭证失败:', error)
            doctorVCs.value = []
          }
        } else {
          console.log('没有找到医生DID')
          doctorVCs.value = []
        }
      } catch (error) {
        console.error('加载DID数据失败:', error)
        // 发生错误时确保数组状态
        allDIDs.value = []
        doctorVCs.value = []
        store.dispatch('app/showError', '加载DID数据失败: ' + error.message)
      } finally {
        loading.value = false
      }
    }

    const refreshData = () => {
      loadData()
    }

    const viewDIDDetails = (did) => {
      selectedDID.value = did
      if (bsDetailsModal) {
        bsDetailsModal.show()
      }
    }

    const copyToClipboard = (text) => {
      navigator.clipboard.writeText(text).then(() => {
        store.dispatch('app/showSuccess', '已复制到剪贴板')
      }).catch(() => {
        store.dispatch('app/showError', '复制失败')
      })
    }

    const formatDate = (dateString) => {
      if (!dateString) return '未知'
      const date = new Date(dateString)
      return date.toLocaleString()
    }

    const truncateDID = (did) => {
      if (!did) return ''
      if (did.length <= 30) return did
      return did.substring(0, 15) + '...' + did.substring(did.length - 10)
    }

    const formatVCContent = (content) => {
      try {
        if (typeof content === 'string') {
          return JSON.stringify(JSON.parse(content), null, 2)
        }
        return JSON.stringify(content, null, 2)
      } catch (error) {
        return content
      }
    }

    onMounted(() => {
      console.log('DIDList 组件挂载')
      console.log('挂载时钱包状态:', {
        isConnected: isWalletConnected.value,
        account: currentAccount.value
      })
      
      // 初始化模态框
      if (detailsModal.value) {
        bsDetailsModal = new Modal(detailsModal.value)
      }

      // 如果钱包已连接，立即加载数据
      if (isWalletConnected.value && currentAccount.value) {
        console.log('钱包已连接，开始加载数据')
        loadData()
      }

      // 监听钱包连接状态变化
      store.watch(
        (state) => state.wallet.isConnected,
        (isConnected, oldValue) => {
          console.log('钱包连接状态变化:', { oldValue, isConnected })
          if (isConnected && currentAccount.value) {
            console.log('钱包已连接，重新加载数据')
            loadData()
          } else if (!isConnected) {
            console.log('钱包已断开，清空数据')
            allDIDs.value = []
            doctorVCs.value = []
          }
        }
      )

      // 监听钱包账户变化
      store.watch(
        (state) => state.wallet.account,
        (account, oldAccount) => {
          console.log('钱包账户变化:', { oldAccount, account })
          if (account && isWalletConnected.value) {
            console.log('账户变化且钱包已连接，重新加载数据')
            loadData()
          }
        }
      )
    })

    return {
      loading,
      isWalletConnected,
      currentAccount,
      doctorDID,
      normalDIDs,
      doctorVCs,
      selectedDID,
      detailsModal,
      refreshData,
      viewDIDDetails,
      copyToClipboard,
      formatDate,
      truncateDID,
      formatVCContent
    }
  }
}
</script>

<style scoped>
.text-monospace {
  font-family: 'Courier New', monospace;
}

pre {
  max-height: 200px;
  overflow-y: auto;
  font-size: 0.875rem;
}

.card {
  transition: all 0.3s ease;
}

.card:hover {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

details summary {
  outline: none;
}

details[open] summary {
  margin-bottom: 0.5rem;
}
</style>
