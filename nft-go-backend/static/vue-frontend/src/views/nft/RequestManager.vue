<template>
  <div class="request-manager">
    <h2><i class="bi bi-file-earmark-text me-2"></i>NFT请求管理</h2>

    <!-- 请求列表 -->
    <div class="card">
      <div class="card-header d-flex justify-content-between align-items-center">
        <h5 class="mb-0"><i class="bi bi-list-ul me-2"></i>请求列表</h5>
        <div>
          <div class="btn-group me-2" role="group">
            <button class="btn btn-sm"
              :class="{ 'btn-primary': filter === 'all', 'btn-outline-primary': filter !== 'all' }"
              @click="setFilter('all')">
              全部
            </button>
            <button class="btn btn-sm"
              :class="{ 'btn-primary': filter === 'pending', 'btn-outline-primary': filter !== 'pending' }"
              @click="setFilter('pending')">
              待处理
            </button>
            <button class="btn btn-sm"
              :class="{ 'btn-primary': filter === 'approved', 'btn-outline-primary': filter !== 'approved' }"
              @click="setFilter('approved')">
              已批准
            </button>
            <button class="btn btn-sm"
              :class="{ 'btn-primary': filter === 'rejected', 'btn-outline-primary': filter !== 'rejected' }"
              @click="setFilter('rejected')">
              已拒绝
            </button>
          </div>
          <button class="btn btn-sm btn-outline-primary" @click="loadRequests">
            <i class="bi bi-arrow-clockwise me-1"></i>刷新
          </button>
        </div>
      </div>
      <div class="card-body">
        <!-- 钱包未连接提示 -->
        <div v-if="!isConnected" class="text-center my-3">
          <i class="bi bi-wallet2 text-warning" style="font-size: 3rem;"></i>
          <h4 class="mt-3">钱包未连接</h4>
          <p class="lead">请先连接您的钱包以查看NFT请求</p>
        </div>

        <!-- 加载状态 -->
        <div v-else-if="isLoading" class="text-center my-3">
          <div class="spinner-border text-primary" role="status">
            <span class="visually-hidden">加载中...</span>
          </div>
          <p class="mt-2">加载请求列表...</p>
        </div>

        <!-- 没有请求时显示 -->
        <div v-else-if="!filteredRequests.length" class="text-center my-3">
          <i class="bi bi-inbox text-muted" style="font-size: 3rem;"></i>
          <h4 class="mt-3">暂无符合条件的请求</h4>
          <p class="lead">当前没有找到相关的NFT请求</p>
        </div>

        <!-- 请求列表 -->
        <div v-else class="table-responsive">
          <table class="table">
            <thead>
              <tr>
                <th>Token ID</th>
                <th>请求者</th>
                <th>状态</th>
                <th>创建时间</th>
                <th>操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="request in filteredRequests" :key="request.id">
                <td><strong>{{ request.token_id || request.tokenId || request.parentTokenId || '新铸造' }}</strong></td>
                <td><code>{{ truncateAddress(request.requester || request.applicantAddress) }}</code></td>
                <td>
                  <span :class="[
                    'badge',
                    request.status === 'pending' ? 'bg-warning' :
                      request.status === 'approved' ? 'bg-success' :
                        'bg-danger'
                  ]">
                    {{ getStatusText(request.status) }}
                  </span>
                </td>
                <td>{{ formatDate(request.created_at || request.createdAt) }}</td>
                <td>
                  <div class="btn-group">
                    <button class="btn btn-sm btn-outline-primary me-1" @click="viewRequest(request)" title="查看详情">
                      <i class="bi bi-eye"></i>
                    </button>
                    <button v-if="request.status === 'pending'" class="btn btn-sm btn-success me-1"
                      @click="handleApproveRequest(request.id || request.requestId)" title="批准">
                      <i class="bi bi-check-lg"></i>
                    </button>
                    <button v-if="request.status === 'pending'" class="btn btn-sm btn-danger"
                      @click="handleRejectRequest(request.id || request.requestId)" title="拒绝">
                      <i class="bi bi-x-lg"></i>
                    </button>
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- 请求详情模态框 -->
    <div class="modal fade" ref="requestDetailModal" tabindex="-1">
      <div class="modal-dialog modal-lg">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">请求详情</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body" v-if="selectedRequest">
            <div class="row">
              <div class="col-md-6">
                <p><strong>请求ID:</strong> {{ selectedRequest.id || selectedRequest.requestId || '未知' }}</p>
                <p><strong>请求者:</strong> {{ selectedRequest.requester || selectedRequest.applicantAddress || '未知' }}</p>
                <p><strong>状态:</strong>
                  <span :class="[
                    'badge',
                    selectedRequest.status === 'pending' ? 'bg-warning' :
                      selectedRequest.status === 'approved' ? 'bg-success' :
                        'bg-danger'
                  ]">
                    {{ getStatusText(selectedRequest.status) }}
                  </span>
                </p>
              </div>
              <div class="col-md-6">
                <p><strong>Token ID:</strong> {{ selectedRequest.token_id || selectedRequest.tokenId ||
                  selectedRequest.parentTokenId || '新铸造' }}</p>
                <p><strong>接收者:</strong> {{ selectedRequest.recipient || selectedRequest.receiverAddress || '不适用' }}</p>
                <p><strong>URI:</strong> {{ selectedRequest.uri || '不适用' }}</p>
                <p><strong>创建时间:</strong> {{ formatDate(selectedRequest.created_at || selectedRequest.createdAt) }}</p>
                <p v-if="selectedRequest.processed_at || selectedRequest.processedAt"><strong>处理时间:</strong> {{
                  formatDate(selectedRequest.processed_at || selectedRequest.processedAt) }}</p>
              </div>
            </div>
            <div class="mt-3" v-if="selectedRequest.transaction_hash || selectedRequest.transactionHash">
              <h6>交易哈希:</h6>
              <code
                class="d-block p-2 bg-light">{{ selectedRequest.transaction_hash || selectedRequest.transactionHash }}</code>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">关闭</button>
            <div v-if="selectedRequest && selectedRequest.status === 'pending'" class="btn-group">
              <button type="button" class="btn btn-success"
                @click="handleApproveRequest(selectedRequest.id || selectedRequest.requestId)">
                <i class="bi bi-check me-1"></i>批准
              </button>
              <button type="button" class="btn btn-danger"
                @click="handleRejectRequest(selectedRequest.id || selectedRequest.requestId)">
                <i class="bi bi-x me-1"></i>拒绝
              </button>
            </div>
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

export default {
  name: 'RequestManager',

  setup() {
    const store = useStore()
    const requestDetailModal = ref(null)
    let bsRequestDetailModal = null

    // 从store获取数据
    const requests = computed(() => store.state.nft.requests)
    const isLoading = computed(() => store.state.nft.isLoading || store.state.app.isLoading)
    const isAdmin = computed(() => store.state.wallet.isAdmin)
    const isConnected = computed(() => store.state.wallet.isConnected)
    const account = computed(() => store.state.wallet.account)

    // 过滤器
    const filter = ref('all')
    const filteredRequests = computed(() => {
      if (filter.value === 'all') {
        return requests.value
      }
      return requests.value.filter(request => request.status === filter.value)
    })

    // 选中的请求
    const selectedRequest = ref(null)

    // 提交状态
    const isSubmitting = ref(false)

    // 加载请求列表
    const loadRequests = async () => {
      // 检查钱包是否已连接
      if (!isConnected.value || !account.value) {
        console.log('钱包未连接，无法加载请求列表')
        store.dispatch('app/showError', '请先连接钱包')
        return
      }

      try {
        console.log('开始加载请求列表，钱包地址:', account.value)
        await store.dispatch('nft/loadRequests')
        console.log('请求列表数据:', store.state.nft.requests)
        
        // 如果请求列表为空，显示提示信息
        if (!store.state.nft.requests || store.state.nft.requests.length === 0) {
          console.log('请求列表为空')
        }
      } catch (error) {
        console.error('加载请求列表失败:', error);
        store.dispatch('app/showError', '加载请求列表失败: ' + error.message);
      }
    }

    // 设置过滤器
    const setFilter = (newFilter) => {
      filter.value = newFilter
    }

    // 查看请求详情
    const viewRequest = (request) => {
      console.log('查看请求详情:', request);
      selectedRequest.value = request;

      if (bsRequestDetailModal) {
        bsRequestDetailModal.show();
      }
    }

    // 批准请求
    const handleApproveRequest = async (requestId) => {
      if (!requestId) {
        store.dispatch('app/showError', '请求ID不能为空');
        return;
      }

      isSubmitting.value = true;
      console.log('批准请求:', requestId);

      try {
        const success = await store.dispatch('nft/approveRequest', requestId);

        if (success) {
          // 显示成功消息
          store.dispatch('app/showSuccess', '请求已批准');

          // 关闭模态框（如果打开）
          if (bsRequestDetailModal) {
            bsRequestDetailModal.hide();
          }

          // 刷新请求列表
          await loadRequests();
        }
      } catch (error) {
        console.error('批准请求失败:', error);
        store.dispatch('app/showError', '批准请求失败: ' + error.message);
      } finally {
        isSubmitting.value = false;
      }
    };

    // 拒绝请求
    const handleRejectRequest = async (requestId) => {
      if (!requestId) {
        store.dispatch('app/showError', '请求ID不能为空');
        return;
      }

      isSubmitting.value = true;
      console.log('拒绝请求:', requestId);

      try {
        const success = await store.dispatch('nft/rejectRequest', requestId);

        if (success) {
          // 显示成功消息
          store.dispatch('app/showSuccess', '请求已拒绝');

          // 关闭模态框（如果打开）
          if (bsRequestDetailModal) {
            bsRequestDetailModal.hide();
          }

          // 刷新请求列表
          await loadRequests();
        }
      } catch (error) {
        console.error('拒绝请求失败:', error);
        store.dispatch('app/showError', '拒绝请求失败: ' + error.message);
      } finally {
        isSubmitting.value = false;
      }
    };

    // 获取请求类型文本
    const getRequestTypeText = (type) => {
      switch (type) {
        case 'mint': return '铸造NFT'
        case 'update': return '更新元数据'
        case 'transfer': return '转移NFT'
        case 'child': return '创建子NFT'
        default: return '未知类型'
      }
    }

    // 获取状态文本
    const getStatusText = (status) => {
      switch (status) {
        case 'pending': return '待处理'
        case 'approved': return '已批准'
        case 'rejected': return '已拒绝'
        default: return '未知状态'
      }
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

    // 截断地址
    const truncateAddress = (address) => {
      if (!address) return ''
      return address.substring(0, 6) + '...' + address.substring(address.length - 4)
    }

    // 监听钱包连接状态变化
    watch([isConnected, account], ([newIsConnected, newAccount], [oldIsConnected, oldAccount]) => {
      console.log('钱包状态变化:', { 
        isConnected: newIsConnected, 
        account: newAccount,
        oldIsConnected: oldIsConnected,
        oldAccount: oldAccount
      })
      
      // 如果钱包从未连接变为已连接，或者账户发生变化，重新加载数据
      if (newIsConnected && newAccount && (
        !oldIsConnected || 
        newAccount !== oldAccount
      )) {
        console.log('钱包连接状态改变，重新加载请求列表')
        loadRequests()
      }
    }, { immediate: false })

    onMounted(() => {
      // 初始化模态框
      if (requestDetailModal.value) {
        bsRequestDetailModal = new Modal(requestDetailModal.value)
      }

      // 如果钱包已经连接，立即加载数据
      if (isConnected.value && account.value) {
        console.log('组件挂载时钱包已连接，立即加载请求列表')
        loadRequests()
      } else {
        console.log('组件挂载时钱包未连接，等待钱包连接后再加载数据')
      }
    })

    return {
      requests,
      filteredRequests,
      isLoading,
      isAdmin,
      isConnected,
      account,
      filter,
      selectedRequest,
      isSubmitting,
      requestDetailModal,
      loadRequests,
      setFilter,
      viewRequest,
      handleApproveRequest,
      handleRejectRequest,
      getRequestTypeText,
      getStatusText,
      formatDate,
      truncateText,
      truncateAddress
    }
  }
}
</script>

<style scoped>
.table td {
  vertical-align: middle;
}
</style>