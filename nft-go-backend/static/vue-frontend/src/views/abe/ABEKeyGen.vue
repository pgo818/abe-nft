<template>
  <div class="abe-keygen">
    <h2>ABE密钥生成</h2>

    <div class="card mt-4">
      <div class="card-header">
        <h5 class="mb-0">生成密钥</h5>
      </div>
      <div class="card-body">
        <form @submit.prevent="generateKey">
          <div class="mb-3">
            <label for="attributes" class="form-label">用户属性</label>
            <textarea class="form-control" id="attributes" v-model="keygenForm.attributes" rows="5"
              placeholder="输入用户属性，每行一个属性，例如：&#10;department:HR&#10;role:manager" required></textarea>
            <div class="form-text">定义用户的属性，每行一个属性</div>
          </div>

          <div class="mb-3">
            <label for="userId" class="form-label">用户ID/钱包地址</label>
            <input type="text" class="form-control" id="userId" v-model="keygenForm.userId" placeholder="0x..."
              required>
            <div class="form-text">用户的唯一标识符或钱包地址</div>
          </div>

          <div class="mb-3 form-check">
            <input type="checkbox" class="form-check-input" id="storeKey" v-model="keygenForm.storeKey">
            <label class="form-check-label" for="storeKey">在服务器上存储密钥</label>
            <div class="form-text">是否在服务器上保存用户密钥的副本</div>
          </div>

          <button type="submit" class="btn btn-primary" :disabled="isSubmitting || !systemInitialized">
            <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-1" role="status"
              aria-hidden="true"></span>
            <i v-else class="bi bi-key-fill me-1"></i>
            生成密钥
          </button>
        </form>
      </div>
    </div>

    <div class="card mt-4">
      <div class="card-header">
        <h5 class="mb-0">已生成的密钥</h5>
      </div>
      <div class="card-body">
        <div v-if="loading" class="text-center py-3">
          <div class="spinner-border text-primary" role="status">
            <span class="visually-hidden">加载中...</span>
          </div>
          <p class="mt-2">加载密钥列表...</p>
        </div>
        <div v-else-if="keys.length === 0" class="text-center py-3">
          <p class="text-muted">暂无密钥记录</p>
        </div>
        <div v-else>
          <div class="table-responsive">
            <table class="table table-hover">
              <thead>
                <tr>
                  <th>用户ID</th>
                  <th>属性</th>
                  <th>创建时间</th>
                  <th>操作</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(key, index) in keys" :key="index">
                  <td>{{ formatAddress(key.userId) }}</td>
                  <td>
                    <span v-for="(attr, i) in key.attributes.split(',')" :key="i" class="badge bg-primary me-1 mb-1">
                      {{ attr }}
                    </span>
                  </td>
                  <td>{{ formatDate(key.createdAt) }}</td>
                  <td>
                    <button class="btn btn-sm btn-outline-primary me-1" @click="viewKey(key)">
                      <i class="bi bi-eye"></i>
                    </button>
                    <button class="btn btn-sm btn-outline-danger" @click="revokeKey(key)">
                      <i class="bi bi-trash"></i>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
/* eslint-disable no-unused-vars */
export default {
  name: 'ABEKeyGen',
  data() {
    return {
      loading: true,
      isSubmitting: false,
      systemInitialized: false,
      keygenForm: {
        attributes: '',
        userId: '',
        storeKey: true
      },
      keys: []
    };
  },
  mounted() {
    this.fetchSystemStatus();
    this.fetchKeys();
  },
  methods: {
    async fetchSystemStatus() {
      try {
        // 模拟API调用
        await new Promise(resolve => setTimeout(resolve, 500));

        // 模拟数据
        this.systemInitialized = true;
      } catch (error) {
        console.error('获取系统状态失败:', error);
        this.$store.commit('notifications/add', {
          type: 'danger',
          message: '获取系统状态失败: ' + error.message
        });
      }
    },
    async fetchKeys() {
      this.loading = true;
      try {
        // 模拟API调用
        await new Promise(resolve => setTimeout(resolve, 800));

        // 模拟数据
        this.keys = [
          {
            id: '1',
            userId: '0x1234567890abcdef1234567890abcdef12345678',
            attributes: 'department:HR,role:manager',
            createdAt: new Date(Date.now() - 86400000).toISOString()
          },
          {
            id: '2',
            userId: '0xabcdef1234567890abcdef1234567890abcdef12',
            attributes: 'department:IT,role:developer,level:senior',
            createdAt: new Date(Date.now() - 172800000).toISOString()
          }
        ];
      } catch (error) {
        console.error('获取密钥列表失败:', error);
        this.$store.commit('notifications/add', {
          type: 'danger',
          message: '获取密钥列表失败: ' + error.message
        });
      } finally {
        this.loading = false;
      }
    },
    async generateKey() {
      if (!this.keygenForm.attributes || !this.keygenForm.userId) {
        this.$store.commit('notifications/add', {
          type: 'danger',
          message: '请填写所有必填字段'
        });
        return;
      }

      this.isSubmitting = true;
      try {
        // 解析属性
        const attributes = this.keygenForm.attributes
          .split('\n')
          .map(attr => attr.trim())
          .filter(attr => attr);

        if (attributes.length === 0) {
          throw new Error('属性不能为空');
        }

        // 模拟API调用
        await new Promise(resolve => setTimeout(resolve, 1500));

        // 添加到列表
        this.keys.unshift({
          id: String(this.keys.length + 1),
          userId: this.keygenForm.userId,
          attributes: attributes.join(','),
          createdAt: new Date().toISOString()
        });

        // 重置表单
        this.keygenForm.attributes = '';

        // 显示成功消息
        this.$store.commit('notifications/add', {
          type: 'success',
          message: '密钥生成成功！'
        });
      } catch (error) {
        console.error('密钥生成失败:', error);
        this.$store.commit('notifications/add', {
          type: 'danger',
          message: '密钥生成失败: ' + error.message
        });
      } finally {
        this.isSubmitting = false;
      }
    },
    viewKey(key) {
      // eslint-disable-next-line no-unused-vars
      this.$store.commit('notifications/add', {
        type: 'info',
        message: `查看密钥功能即将推出！`
      });
    },
    revokeKey(key) {
      if (confirm(`确定要撤销用户 ${this.formatAddress(key.userId)} 的密钥吗？`)) {
        // 模拟撤销
        this.keys = this.keys.filter(k => k.id !== key.id);

        this.$store.commit('notifications/add', {
          type: 'success',
          message: '密钥已成功撤销'
        });
      }
    },
    formatDate(dateString) {
      if (!dateString) return '未知';
      const date = new Date(dateString);
      return date.toLocaleString();
    },
    formatAddress(address) {
      if (!address) return '';
      return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`;
    }
  }
};
</script>

<style scoped>
.bi {
  vertical-align: middle;
}

.badge {
  font-size: 0.8rem;
}
</style>
