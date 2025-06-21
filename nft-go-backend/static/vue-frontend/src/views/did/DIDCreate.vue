<template>
  <div class="container mt-4">
    <div class="d-flex justify-content-between align-items-center mb-4">
      <h2>创建 DID</h2>
      <router-link to="/did/list" class="btn btn-outline-primary">
        <i class="bi bi-arrow-left me-1"></i>返回列表
      </router-link>
    </div>

    <div class="row">
      <div class="col-lg-8">
        <div class="card">
          <div class="card-header">
            创建新的分布式身份标识符
          </div>
          <div class="card-body">
            <form @submit.prevent="createDID">
              <div class="mb-3">
                <label for="didName" class="form-label">DID 名称</label>
                <input type="text" class="form-control" id="didName" v-model="didForm.name"
                  placeholder="为您的 DID 提供一个易于识别的名称" required>
                <div class="form-text">此名称仅用于显示，不会包含在 DID 文档中</div>
              </div>

              <div class="mb-3">
                <label for="didType" class="form-label">DID 类型</label>
                <select class="form-select" id="didType" v-model="didForm.type" required>
                  <option value="personal">个人身份</option>
                  <option value="professional">专业身份</option>
                  <option value="medical">医疗身份</option>
                  <option value="organization">组织身份</option>
                </select>
                <div class="form-text">选择 DID 的用途类型</div>
              </div>

              <div class="mb-3">
                <label for="didMethod" class="form-label">DID 方法</label>
                <select class="form-select" id="didMethod" v-model="didForm.method" required>
                  <option value="key">did:key (本地密钥)</option>
                  <option value="web">did:web (网站关联)</option>
                  <option value="ethr">did:ethr (以太坊)</option>
                </select>
                <div class="form-text">选择 DID 的底层实现方法</div>
              </div>

              <div class="mb-3" v-if="didForm.method === 'ethr'">
                <label for="ethereumAddress" class="form-label">以太坊地址</label>
                <div class="input-group">
                  <span class="input-group-text">
                    <i class="bi bi-wallet2"></i>
                  </span>
                  <input type="text" class="form-control" id="ethereumAddress" v-model="didForm.ethereumAddress"
                    placeholder="0x..." required>
                  <button type="button" class="btn btn-outline-secondary" @click="useCurrentWallet"
                    v-if="isWalletConnected">
                    使用当前钱包
                  </button>
                </div>
                <div class="form-text">用于创建和控制 DID 的以太坊地址</div>
              </div>

              <div class="mb-3" v-if="didForm.method === 'web'">
                <label for="domain" class="form-label">域名</label>
                <div class="input-group">
                  <span class="input-group-text">https://</span>
                  <input type="text" class="form-control" id="domain" v-model="didForm.domain" placeholder="example.com"
                    required>
                </div>
                <div class="form-text">与 DID 关联的域名</div>
              </div>

              <div class="mb-3">
                <label class="form-label">高级选项</label>
                <div class="form-check mb-2">
                  <input class="form-check-input" type="checkbox" id="useKeyAgreement"
                    v-model="didForm.useKeyAgreement">
                  <label class="form-check-label" for="useKeyAgreement">
                    添加密钥协商密钥
                  </label>
                </div>
                <div class="form-check mb-2">
                  <input class="form-check-input" type="checkbox" id="useAssertionMethod"
                    v-model="didForm.useAssertionMethod">
                  <label class="form-check-label" for="useAssertionMethod">
                    添加断言方法
                  </label>
                </div>
                <div class="form-check">
                  <input class="form-check-input" type="checkbox" id="useCapabilityInvocation"
                    v-model="didForm.useCapabilityInvocation">
                  <label class="form-check-label" for="useCapabilityInvocation">
                    添加能力调用
                  </label>
                </div>
                <div class="form-text">添加额外的密钥用途</div>
              </div>

              <div class="d-grid gap-2">
                <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
                  <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-1" role="status"
                    aria-hidden="true"></span>
                  <i v-else class="bi bi-plus-circle me-1"></i>创建 DID
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>

      <div class="col-lg-4">
        <div class="card mb-4">
          <div class="card-header">
            DID 方法说明
          </div>
          <div class="card-body">
            <div v-if="didForm.method === 'key'">
              <h5>did:key</h5>
              <p>基于密码学密钥的 DID 方法，不依赖任何特定的分布式账本或网络。适用于：</p>
              <ul>
                <li>需要快速创建 DID</li>
                <li>不需要更新 DID 文档</li>
                <li>需要离线工作的场景</li>
              </ul>
              <div class="alert alert-info">
                <i class="bi bi-info-circle-fill me-2"></i>
                将自动生成一个新的 Ed25519 密钥对
              </div>
            </div>

            <div v-if="didForm.method === 'web'">
              <h5>did:web</h5>
              <p>基于网站域名的 DID 方法，利用现有的 DNS 和 HTTPS 基础设施。适用于：</p>
              <ul>
                <li>组织和企业</li>
                <li>已拥有域名的实体</li>
                <li>需要简单解析机制的场景</li>
              </ul>
              <div class="alert alert-warning">
                <i class="bi bi-exclamation-triangle-fill me-2"></i>
                您需要在域名的 .well-known 目录下托管 DID 文档
              </div>
            </div>

            <div v-if="didForm.method === 'ethr'">
              <h5>did:ethr</h5>
              <p>基于以太坊区块链的 DID 方法，利用智能合约实现 DID 文档管理。适用于：</p>
              <ul>
                <li>需要在区块链上管理身份</li>
                <li>需要可更新的 DID 文档</li>
                <li>与以太坊生态系统集成</li>
              </ul>
              <div class="alert alert-info">
                <i class="bi bi-info-circle-fill me-2"></i>
                需要连接以太坊钱包并支付少量 gas 费用
              </div>
            </div>
          </div>
        </div>

        <div class="card">
          <div class="card-header">
            DID 预览
          </div>
          <div class="card-body">
            <div v-if="didPreview">
              <p><strong>DID:</strong></p>
              <div class="bg-light p-2 rounded mb-3">
                <code>{{ didPreview }}</code>
              </div>

              <p><strong>DID 文档 (预览):</strong></p>
              <pre class="bg-light p-2 rounded"><code>{{ didDocumentPreview }}</code></pre>
            </div>
            <div v-else class="text-center py-3 text-muted">
              <i class="bi bi-person-badge fs-1"></i>
              <p class="mt-2">填写表单以查看 DID 预览</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 创建成功模态框 -->
    <div class="modal fade" ref="successModal" tabindex="-1">
      <div class="modal-dialog modal-lg">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">DID 创建成功</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body" v-if="createdDID">
            <div class="alert alert-success">
              <h5 class="alert-heading"><i class="bi bi-check-circle-fill me-2"></i>您的 DID 已成功创建！</h5>
              <p>您现在可以使用此 DID 进行身份验证和凭证管理。</p>
            </div>

            <div class="mb-3">
              <label class="form-label fw-bold">DID</label>
              <div class="input-group">
                <input type="text" class="form-control" :value="createdDID.did" readonly>
                <button class="btn btn-outline-secondary" @click="copyToClipboard(createdDID.did)">
                  <i class="bi bi-clipboard"></i>
                </button>
              </div>
            </div>

            <div class="mb-3">
              <label class="form-label fw-bold">私钥</label>
              <div class="alert alert-danger">
                <i class="bi bi-exclamation-triangle-fill me-2"></i>
                <strong>重要提示：</strong> 请安全保存此私钥，它不会再次显示！
              </div>
              <div class="input-group">
                <input :type="showPrivateKey ? 'text' : 'password'" class="form-control" :value="createdDID.privateKey"
                  readonly>
                <button class="btn btn-outline-secondary" @click="togglePrivateKeyVisibility">
                  <i :class="showPrivateKey ? 'bi bi-eye-slash' : 'bi bi-eye'"></i>
                </button>
                <button class="btn btn-outline-secondary" @click="copyToClipboard(createdDID.privateKey)">
                  <i class="bi bi-clipboard"></i>
                </button>
              </div>
            </div>

            <div class="mb-3">
              <label class="form-label fw-bold">DID 文档</label>
              <pre class="bg-light p-3 rounded"><code>{{ JSON.stringify(createdDID.document, null, 2) }}</code></pre>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">关闭</button>
            <button type="button" class="btn btn-primary" @click="exportDIDDocument">
              <i class="bi bi-download me-1"></i>导出 DID 文档
            </button>
            <button type="button" class="btn btn-success" @click="goToDIDList">
              <i class="bi bi-list-check me-1"></i>查看 DID 列表
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex';

export default {
  name: 'DIDCreate',
  data() {
    return {
      isSubmitting: false,
      didForm: {
        name: '',
        type: 'personal',
        method: 'key',
        ethereumAddress: '',
        domain: '',
        useKeyAgreement: false,
        useAssertionMethod: false,
        useCapabilityInvocation: false
      },
      createdDID: null,
      showPrivateKey: false,
      successModal: null
    };
  },
  computed: {
    ...mapState('wallet', ['isConnected', 'address']),
    isWalletConnected() {
      return this.isConnected;
    },
    didPreview() {
      if (!this.didForm.name) return '';

      switch (this.didForm.method) {
        case 'key':
          return 'did:key:z6MkhaXgBZDvotDkL5257faiztiGiC2QtKLGpbnnEGta2doK';
        case 'web':
          if (!this.didForm.domain) return '';
          return `did:web:${this.didForm.domain}`;
        case 'ethr': {
          if (!this.didForm.ethereumAddress) return '';
          const address = this.didForm.ethereumAddress.startsWith('0x')
            ? this.didForm.ethereumAddress.substring(2)
            : this.didForm.ethereumAddress;
          return `did:ethr:${address}`;
        }
        default:
          return '';
      }
    },
    didDocumentPreview() {
      if (!this.didPreview) return '';

      const document = {
        "@context": "https://www.w3.org/ns/did/v1",
        "id": this.didPreview,
        "verificationMethod": [
          {
            "id": `${this.didPreview}#keys-1`,
            "type": "Ed25519VerificationKey2020",
            "controller": this.didPreview,
            "publicKeyMultibase": "z6MkhaXgBZDvotDkL5257faiztiGiC2QtKLGpbnnEGta2doK"
          }
        ],
        "authentication": [
          `${this.didPreview}#keys-1`
        ]
      };

      if (this.didForm.useKeyAgreement) {
        document.keyAgreement = [
          `${this.didPreview}#keys-2`
        ];
        document.verificationMethod.push({
          "id": `${this.didPreview}#keys-2`,
          "type": "X25519KeyAgreementKey2020",
          "controller": this.didPreview,
          "publicKeyMultibase": "z6LSbysY2xFMRpGMhb7tFTLMpeuPRaqaWM1yECx2AtzE3KCc"
        });
      }

      if (this.didForm.useAssertionMethod) {
        document.assertionMethod = [
          `${this.didPreview}#keys-1`
        ];
      }

      if (this.didForm.useCapabilityInvocation) {
        document.capabilityInvocation = [
          `${this.didPreview}#keys-1`
        ];
      }

      return JSON.stringify(document, null, 2);
    }
  },
  mounted() {
    // 初始化Bootstrap模态框
    this.successModal = new this.$bootstrap.Modal(this.$refs.successModal);
  },
  methods: {
    useCurrentWallet() {
      if (this.isWalletConnected) {
        this.didForm.ethereumAddress = this.address;
      }
    },
    async createDID() {
      this.isSubmitting = true;

      try {
        // 验证表单
        if (this.didForm.method === 'ethr' && !this.didForm.ethereumAddress) {
          throw new Error('请输入以太坊地址');
        }

        if (this.didForm.method === 'web' && !this.didForm.domain) {
          throw new Error('请输入域名');
        }

        // 模拟API调用
        // 实际应用中应该替换为真实的API调用
        await new Promise(resolve => setTimeout(resolve, 1500));

        // 模拟创建结果
        this.createdDID = {
          did: this.didPreview,
          privateKey: 'z6MkhaXgBZDvotDkL5257faiztiGiC2QtKLGpbnnEGta2doK_private_key_example',
          document: JSON.parse(this.didDocumentPreview),
          name: this.didForm.name,
          type: this.didForm.type,
          method: this.didForm.method,
          createdAt: new Date().toISOString()
        };

        // 显示成功模态框
        this.successModal.show();

        this.$store.commit('notifications/add', {
          type: 'success',
          message: 'DID 创建成功！'
        });
      } catch (error) {
        console.error('创建DID失败:', error);
        this.$store.commit('notifications/add', {
          type: 'danger',
          message: '创建DID失败: ' + error.message
        });
      } finally {
        this.isSubmitting = false;
      }
    },
    togglePrivateKeyVisibility() {
      this.showPrivateKey = !this.showPrivateKey;
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
    exportDIDDocument() {
      if (!this.createdDID) return;

      const documentStr = JSON.stringify(this.createdDID.document, null, 2);
      const blob = new Blob([documentStr], { type: 'application/json' });
      const url = URL.createObjectURL(blob);

      const a = document.createElement('a');
      a.href = url;
      a.download = `${this.createdDID.did.replace(/:/g, '-')}.json`;
      document.body.appendChild(a);
      a.click();
      document.body.removeChild(a);
      URL.revokeObjectURL(url);

      this.$store.commit('notifications/add', {
        type: 'success',
        message: 'DID 文档已导出'
      });
    },
    goToDIDList() {
      this.successModal.hide();
      this.$router.push('/did/list');
    }
  }
};
</script>

<style scoped>
pre {
  max-height: 200px;
  overflow-y: auto;
  font-size: 0.85rem;
}

.card {
  margin-bottom: 1.5rem;
}
</style>