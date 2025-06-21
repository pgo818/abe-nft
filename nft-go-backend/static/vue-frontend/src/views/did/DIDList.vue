<template>
  <div class="container mt-4">
    <div class="d-flex justify-content-between align-items-center mb-4">
      <h2>DID 列表</h2>
      <router-link to="/did/create" class="btn btn-primary">
        <i class="bi bi-plus-circle me-1"></i>创建新 DID
      </router-link>
    </div>

    <div class="card mb-4">
      <div class="card-body">
        <div class="row">
          <div class="col-md-8">
            <div class="input-group">
              <span class="input-group-text">
                <i class="bi bi-search"></i>
              </span>
              <input type="text" class="form-control" placeholder="搜索 DID 或名称" v-model="searchQuery"
                @input="filterDIDs">
            </div>
          </div>
          <div class="col-md-4">
            <select class="form-select" v-model="filterType" @change="filterDIDs">
              <option value="all">所有 DID</option>
              <option value="owned">我创建的</option>
              <option value="shared">与我共享的</option>
            </select>
          </div>
        </div>
      </div>
    </div>

    <div v-if="loading" class="text-center my-5">
      <div class="spinner-border text-primary" role="status">
        <span class="visually-hidden">加载中...</span>
      </div>
      <p class="mt-2">加载 DID 列表...</p>
    </div>

    <div v-else-if="filteredDIDs.length === 0" class="text-center my-5">
      <i class="bi bi-person-badge fs-1 text-muted"></i>
      <p class="lead mt-3">没有找到 DID</p>
      <p class="text-muted">创建一个新的 DID 或调整搜索条件</p>
      <router-link to="/did/create" class="btn btn-primary mt-2">
        <i class="bi bi-plus-circle me-1"></i>创建新 DID
      </router-link>
    </div>

    <div v-else>
      <div class="row row-cols-1 row-cols-md-2 row-cols-lg-3 g-4">
        <div class="col" v-for="did in filteredDIDs" :key="did.id">
          <div class="card h-100">
            <div class="card-header d-flex justify-content-between align-items-center">
              <h5 class="mb-0">{{ did.name }}</h5>
              <span class="badge" :class="getStatusBadgeClass(did.status)">
                {{ getStatusText(did.status) }}
              </span>
            </div>
            <div class="card-body">
              <div class="mb-3">
                <label class="form-label fw-bold">DID</label>
                <div class="input-group">
                  <input type="text" class="form-control" :value="did.did" readonly>
                  <button class="btn btn-outline-secondary" @click="copyToClipboard(did.did)">
                    <i class="bi bi-clipboard"></i>
                  </button>
                </div>
              </div>

              <div class="mb-3">
                <label class="form-label fw-bold">创建时间</label>
                <p class="mb-0">{{ formatDate(did.createdAt) }}</p>
              </div>

              <div class="mb-3">
                <label class="form-label fw-bold">类型</label>
                <p class="mb-0">{{ did.type }}</p>
              </div>

              <div class="mb-3">
                <label class="form-label fw-bold">关联凭证</label>
                <p class="mb-0">{{ did.vcCount || 0 }} 个凭证</p>
              </div>
            </div>
            <div class="card-footer">
              <div class="d-flex justify-content-between">
                <button class="btn btn-sm btn-outline-primary" @click="viewDetails(did)">
                  <i class="bi bi-eye me-1"></i>详情
                </button>
                <div class="btn-group">
                  <button class="btn btn-sm btn-outline-success" @click="issueVC(did)">
                    <i class="bi bi-file-earmark-plus me-1"></i>颁发凭证
                  </button>
                  <button class="btn btn-sm btn-outline-danger" @click="confirmDelete(did)">
                    <i class="bi bi-trash me-1"></i>删除
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- DID 详情模态框 -->
    <div class="modal fade" ref="detailsModal" tabindex="-1">
      <div class="modal-dialog modal-lg">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">DID 详情</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body" v-if="selectedDID">
            <div class="mb-3">
              <label class="form-label fw-bold">名称</label>
              <p>{{ selectedDID.name }}</p>
            </div>

            <div class="mb-3">
              <label class="form-label fw-bold">DID</label>
              <div class="input-group">
                <input type="text" class="form-control" :value="selectedDID.did" readonly>
                <button class="btn btn-outline-secondary" @click="copyToClipboard(selectedDID.did)">
                  <i class="bi bi-clipboard"></i>
                </button>
              </div>
            </div>

            <div class="mb-3">
              <label class="form-label fw-bold">DID 文档</label>
              <pre class="bg-light p-3 rounded"><code>{{ JSON.stringify(selectedDID.document, null, 2) }}</code></pre>
            </div>

            <div class="mb-3">
              <label class="form-label fw-bold">关联凭证</label>
              <div v-if="selectedDID.vcs && selectedDID.vcs.length > 0">
                <div class="list-group">
                  <a href="#" class="list-group-item list-group-item-action" v-for="(vc, index) in selectedDID.vcs"
                    :key="index">
                    <div class="d-flex w-100 justify-content-between">
                      <h6 class="mb-1">{{ vc.type }}</h6>
                      <small>{{ formatDate(vc.issuanceDate) }}</small>
                    </div>
                    <p class="mb-1">{{ vc.issuer }}</p>
                    <small>ID: {{ vc.id }}</small>
                  </a>
                </div>
              </div>
              <p v-else class="text-muted">暂无关联凭证</p>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">关闭</button>
            <button type="button" class="btn btn-primary" @click="exportDID">
              <i class="bi bi-download me-1"></i>导出 DID 文档
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- 删除确认模态框 -->
    <div class="modal fade" ref="deleteModal" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">确认删除</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body" v-if="selectedDID">
            <div class="alert alert-danger">
              <i class="bi bi-exclamation-triangle-fill me-2"></i>
              <strong>警告：</strong> 删除操作不可逆，将永久删除此 DID 及其相关数据。
            </div>
            <p>您确定要删除以下 DID 吗？</p>
            <p><strong>名称：</strong> {{ selectedDID.name }}</p>
            <p><strong>DID：</strong> {{ selectedDID.did }}</p>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">取消</button>
            <button type="button" class="btn btn-danger" @click="deleteDID" :disabled="isDeleting">
              <span v-if="isDeleting" class="spinner-border spinner-border-sm me-1" role="status"
                aria-hidden="true"></span>
              <i v-else class="bi bi-trash me-1"></i>确认删除
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DIDList',
  data() {
    return {
      loading: true,
      isDeleting: false,
      dids: [],
      filteredDIDs: [],
      searchQuery: '',
      filterType: 'all',
      selectedDID: null,
      detailsModal: null,
      deleteModal: null
    };
  },
  mounted() {
    this.fetchDIDs();
    // 初始化Bootstrap模态框
    this.detailsModal = new this.$bootstrap.Modal(this.$refs.detailsModal);
    this.deleteModal = new this.$bootstrap.Modal(this.$refs.deleteModal);
  },
  methods: {
    async fetchDIDs() {
      this.loading = true;
      try {
        // 模拟API调用
        // 实际应用中应该替换为真实的API调用
        await new Promise(resolve => setTimeout(resolve, 800));

        // 模拟数据
        this.dids = [
          {
            id: '1',
            name: '个人身份',
            did: 'did:example:123456789abcdefghi',
            type: '个人',
            status: 'active',
            createdAt: new Date(Date.now() - 3600000).toISOString(),
            vcCount: 3,
            document: {
              "@context": "https://www.w3.org/ns/did/v1",
              "id": "did:example:123456789abcdefghi",
              "verificationMethod": [
                {
                  "id": "did:example:123456789abcdefghi#keys-1",
                  "type": "Ed25519VerificationKey2020",
                  "controller": "did:example:123456789abcdefghi",
                  "publicKeyMultibase": "zH3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV"
                }
              ],
              "authentication": [
                "did:example:123456789abcdefghi#keys-1"
              ]
            },
            vcs: [
              {
                id: 'vc:example:123',
                type: '身份证明',
                issuer: 'did:example:issuer123',
                issuanceDate: new Date(Date.now() - 86400000).toISOString()
              },
              {
                id: 'vc:example:456',
                type: '医疗记录访问权限',
                issuer: 'did:example:hospital123',
                issuanceDate: new Date(Date.now() - 172800000).toISOString()
              }
            ]
          },
          {
            id: '2',
            name: '医疗身份',
            did: 'did:example:abcdefghi123456789',
            type: '医疗',
            status: 'active',
            createdAt: new Date(Date.now() - 86400000).toISOString(),
            vcCount: 1,
            document: {
              "@context": "https://www.w3.org/ns/did/v1",
              "id": "did:example:abcdefghi123456789",
              "verificationMethod": [
                {
                  "id": "did:example:abcdefghi123456789#keys-1",
                  "type": "Ed25519VerificationKey2020",
                  "controller": "did:example:abcdefghi123456789",
                  "publicKeyMultibase": "zH3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV"
                }
              ],
              "authentication": [
                "did:example:abcdefghi123456789#keys-1"
              ]
            },
            vcs: [
              {
                id: 'vc:example:789',
                type: '医生资格证明',
                issuer: 'did:example:medboard123',
                issuanceDate: new Date(Date.now() - 259200000).toISOString()
              }
            ]
          },
          {
            id: '3',
            name: '工作身份',
            did: 'did:example:9876543210abcdefghi',
            type: '专业',
            status: 'inactive',
            createdAt: new Date(Date.now() - 172800000).toISOString(),
            vcCount: 0,
            document: {
              "@context": "https://www.w3.org/ns/did/v1",
              "id": "did:example:9876543210abcdefghi",
              "verificationMethod": [
                {
                  "id": "did:example:9876543210abcdefghi#keys-1",
                  "type": "Ed25519VerificationKey2020",
                  "controller": "did:example:9876543210abcdefghi",
                  "publicKeyMultibase": "zH3C2AVvLMv6gmMNam3uVAjZpfkcJCwDwnZn6z3wXmqPV"
                }
              ],
              "authentication": [
                "did:example:9876543210abcdefghi#keys-1"
              ]
            },
            vcs: []
          }
        ];

        this.filterDIDs();
      } catch (error) {
        console.error('获取DID列表失败:', error);
        this.$store.commit('notifications/add', {
          type: 'danger',
          message: '获取DID列表失败: ' + error.message
        });
      } finally {
        this.loading = false;
      }
    },
    filterDIDs() {
      if (!this.searchQuery && this.filterType === 'all') {
        this.filteredDIDs = [...this.dids];
        return;
      }

      let filtered = [...this.dids];

      // 应用类型过滤
      if (this.filterType !== 'all') {
        if (this.filterType === 'owned') {
          filtered = filtered.filter(did => did.type === '个人' || did.type === '专业');
        } else if (this.filterType === 'shared') {
          filtered = filtered.filter(did => did.type === '医疗');
        }
      }

      // 应用搜索过滤
      if (this.searchQuery) {
        const query = this.searchQuery.toLowerCase();
        filtered = filtered.filter(did =>
          did.name.toLowerCase().includes(query) ||
          did.did.toLowerCase().includes(query) ||
          did.type.toLowerCase().includes(query)
        );
      }

      this.filteredDIDs = filtered;
    },
    formatDate(dateString) {
      if (!dateString) return '未知';
      const date = new Date(dateString);
      return date.toLocaleString();
    },
    getStatusBadgeClass(status) {
      switch (status) {
        case 'active':
          return 'bg-success';
        case 'inactive':
          return 'bg-secondary';
        case 'revoked':
          return 'bg-danger';
        case 'pending':
          return 'bg-warning';
        default:
          return 'bg-primary';
      }
    },
    getStatusText(status) {
      switch (status) {
        case 'active':
          return '活跃';
        case 'inactive':
          return '未激活';
        case 'revoked':
          return '已撤销';
        case 'pending':
          return '待处理';
        default:
          return '未知';
      }
    },
    viewDetails(did) {
      this.selectedDID = did;
      this.detailsModal.show();
    },
    issueVC(did) {
      // 实现颁发凭证的逻辑
      this.$router.push({
        path: '/did/vc/issue',
        query: { did: did.did }
      });
    },
    confirmDelete(did) {
      this.selectedDID = did;
      this.deleteModal.show();
    },
    async deleteDID() {
      if (!this.selectedDID) return;

      this.isDeleting = true;

      try {
        // 模拟API调用
        // 实际应用中应该替换为真实的API调用
        await new Promise(resolve => setTimeout(resolve, 800));

        // 从列表中移除
        this.dids = this.dids.filter(did => did.id !== this.selectedDID.id);
        this.filterDIDs();

        // 关闭模态框
        this.deleteModal.hide();

        this.$store.commit('notifications/add', {
          type: 'success',
          message: 'DID 已成功删除'
        });
      } catch (error) {
        console.error('删除DID失败:', error);
        this.$store.commit('notifications/add', {
          type: 'danger',
          message: '删除DID失败: ' + error.message
        });
      } finally {
        this.isDeleting = false;
      }
    },
    copyToClipboard(text) {
      navigator.clipboard.writeText(text).then(
        () => {
          this.$store.commit('notifications/add', {
            type: 'success',
            message: '已复制到剪贴板'
          });
        },
        () => {
          this.$store.commit('notifications/add', {
            type: 'danger',
            message: '复制到剪贴板失败'
          });
        }
      );
    },
    exportDID() {
      if (!this.selectedDID) return;

      const documentStr = JSON.stringify(this.selectedDID.document, null, 2);
      const blob = new Blob([documentStr], { type: 'application/json' });
      const url = URL.createObjectURL(blob);

      const a = document.createElement('a');
      a.href = url;
      a.download = `${this.selectedDID.did.replace(/:/g, '-')}.json`;
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
      URL.revokeObjectURL(url);

      this.$store.commit('notifications/add', {
        type: 'success',
        message: 'DID 文档已导出'
      });
    }
  }
};
</script>

<style scoped>
pre {
  max-height: 300px;
  overflow-y: auto;
}

.card {
  transition: transform 0.3s;
}

.card:hover {
  transform: translateY(-5px);
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
}
</style>
