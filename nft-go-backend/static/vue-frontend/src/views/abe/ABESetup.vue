<template>
  <div class="abe-setup">
    <h2>ABE系统初始化</h2>

    <div class="card mt-4">
      <div class="card-header">
        <h5 class="mb-0">系统状态</h5>
      </div>
      <div class="card-body">
        <div v-if="loading" class="text-center py-3">
          <div class="spinner-border text-primary" role="status">
            <span class="visually-hidden">加载中...</span>
          </div>
          <p class="mt-2">加载系统状态...</p>
        </div>
        <div v-else>
          <div class="alert" :class="systemInitialized ? 'alert-success' : 'alert-warning'">
            <i :class="systemInitialized ? 'bi bi-check-circle-fill' : 'bi bi-exclamation-triangle-fill'"
              class="me-2"></i>
            <span v-if="systemInitialized">ABE 系统已初始化</span>
            <span v-else>ABE 系统尚未初始化，请先完成系统初始化</span>
          </div>

          <div class="row mt-3" v-if="systemInitialized">
            <div class="col-md-6">
              <h5>系统信息</h5>
              <ul class="list-group">
                <li class="list-group-item d-flex justify-content-between align-items-center">
                  初始化时间
                  <span>{{ formatDate(systemInfo.setupTime) }}</span>
                </li>
                <li class="list-group-item d-flex justify-content-between align-items-center">
                  密钥数量
                  <span class="badge bg-primary rounded-pill">{{ systemInfo.keyCount || 0 }}</span>
                </li>
                <li class="list-group-item d-flex justify-content-between align-items-center">
                  加密操作次数
                  <span class="badge bg-success rounded-pill">{{ systemInfo.encryptCount || 0 }}</span>
                </li>
                <li class="list-group-item d-flex justify-content-between align-items-center">
                  解密操作次数
                  <span class="badge bg-info rounded-pill">{{ systemInfo.decryptCount || 0 }}</span>
                </li>
              </ul>
            </div>
            <div class="col-md-6">
              <h5>最近操作</h5>
              <div class="list-group">
                <div v-if="recentOperations.length === 0" class="text-center py-3 text-muted">
                  暂无操作记录
                </div>
                <a href="#" class="list-group-item list-group-item-action" v-for="(op, index) in recentOperations"
                  :key="index">
                  <div class="d-flex w-100 justify-content-between">
                    <h6 class="mb-1">{{ op.type }}</h6>
                    <small>{{ formatDate(op.time) }}</small>
                  </div>
                  <p class="mb-1">{{ op.description }}</p>
                  <small>操作者: {{ op.user }}</small>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="card mt-4">
      <div class="card-header">
        <h5 class="mb-0">初始化系统</h5>
      </div>
      <div class="card-body">
        <form @submit.prevent="setupSystem">
          <button type="submit" class="btn btn-primary" :disabled="isSubmitting || systemInitialized">
            <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-1" role="status"
              aria-hidden="true"></span>
            <i v-else class="bi bi-gear-fill me-1"></i>
            初始化ABE系统
          </button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ABESetup',
  data() {
    return {
      loading: true,
      isSubmitting: false,
      systemInitialized: false,
      systemInfo: {
        setupTime: null,
        keyCount: 0,
        encryptCount: 0,
        decryptCount: 0
      },
      recentOperations: []
    };
  },
  mounted() {
    this.fetchSystemStatus();
  },
  methods: {
    async fetchSystemStatus() {
      this.loading = true;
      try {
        // 模拟API调用
        // 实际应用中应该替换为真实的API调用
        await new Promise(resolve => setTimeout(resolve, 800));

        // 模拟数据
        this.systemInitialized = true;
        this.systemInfo = {
          setupTime: new Date().toISOString(),
          keyCount: 5,
          encryptCount: 12,
          decryptCount: 8
        };

        this.recentOperations = [
          {
            type: '密钥生成',
            time: new Date(Date.now() - 3600000).toISOString(),
            description: '为用户生成了属性为 "department:HR AND role:manager" 的密钥',
            user: '0x1234...5678'
          },
          {
            type: '加密',
            time: new Date(Date.now() - 7200000).toISOString(),
            description: '使用策略 "(department:HR OR department:IT) AND role:manager" 加密数据',
            user: '0x8765...4321'
          }
        ];
      } catch (error) {
        console.error('获取系统状态失败:', error);
        this.$store.commit('notifications/add', {
          type: 'danger',
          message: '获取系统状态失败: ' + error.message
        });
      } finally {
        this.loading = false;
      }
    },
    async setupSystem() {
      this.isSubmitting = true;
      try {
        // 模拟API调用
        await new Promise(resolve => setTimeout(resolve, 1500));

        // 更新状态
        this.systemInitialized = true;
        this.systemInfo.setupTime = new Date().toISOString();

        // 显示成功消息
        this.$store.commit('notifications/add', {
          type: 'success',
          message: 'ABE系统初始化成功！'
        });

        // 刷新状态
        this.fetchSystemStatus();
      } catch (error) {
        console.error('系统初始化失败:', error);
        this.$store.commit('notifications/add', {
          type: 'danger',
          message: '系统初始化失败: ' + error.message
        });
      } finally {
        this.isSubmitting = false;
      }
    },
    formatDate(dateString) {
      if (!dateString) return '未知';
      const date = new Date(dateString);
      return date.toLocaleString();
    }
  }
};
</script>

<style scoped>
.bi {
  vertical-align: middle;
}
</style>