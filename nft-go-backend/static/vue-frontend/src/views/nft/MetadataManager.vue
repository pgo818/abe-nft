<template>
  <div class="container mt-4">
    <h2>NFT 元数据管理</h2>

    <div class="row">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            创建新元数据
          </div>
          <div class="card-body">
            <form @submit.prevent="createMetadata">
              <div class="mb-3">
                <label for="name" class="form-label">名称</label>
                <input type="text" class="form-control" id="name" v-model="newMetadata.name" required>
              </div>

              <div class="mb-3">
                <label for="description" class="form-label">描述</label>
                <textarea class="form-control" id="description" v-model="newMetadata.description" rows="3"
                  required></textarea>
              </div>

              <div class="mb-3">
                <label for="image" class="form-label">图片 URL</label>
                <input type="text" class="form-control" id="image" v-model="newMetadata.image" required>
                <div class="form-text">IPFS 或 HTTP URL</div>
              </div>

              <div class="mb-3">
                <label for="policy" class="form-label">访问策略</label>
                <input type="text" class="form-control" id="policy" v-model="newMetadata.policy" required>
                <div class="form-text">ABE访问策略，例如: "department:CS AND year:2023"</div>
              </div>

              <div class="mb-3">
                <label for="ciphertext" class="form-label">密文</label>
                <textarea class="form-control" id="ciphertext" v-model="newMetadata.ciphertext" rows="3"
                  required></textarea>
                <div class="form-text">加密后的内容（如果没有，请输入"none"）</div>
              </div>

              <div class="d-grid">
                <button type="submit" class="btn btn-primary" :disabled="isSubmitting">
                  <span v-if="isSubmitting" class="spinner-border spinner-border-sm me-1" role="status"
                    aria-hidden="true"></span>
                  创建元数据
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>

      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            预览
          </div>
          <div class="card-body">
            <div v-if="newMetadata.name || newMetadata.description || newMetadata.image" class="metadata-preview">
              <img v-if="newMetadata.image" :src="newMetadata.image" class="img-fluid mb-3" alt="NFT Preview"
                @error="handleImageError">
              <img v-else
                src="data:image/svg+xml;charset=UTF-8,%3Csvg%20width%3D%22200%22%20height%3D%22200%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22%23f0f0f0%22%2F%3E%3Ctext%20x%3D%2250%25%22%20y%3D%2250%25%22%20font-family%3D%22Arial%22%20font-size%3D%2224%22%20text-anchor%3D%22middle%22%20dominant-baseline%3D%22middle%22%20fill%3D%22%23999%22%3ENFT%20%E5%8D%A0%E4%BD%8D%E5%9B%BE%3C%2Ftext%3E%3C%2Fsvg%3E"
                class="img-fluid mb-3" alt="NFT Preview">

              <h4>{{ newMetadata.name || 'NFT 名称' }}</h4>
              <p>{{ newMetadata.description || '暂无描述' }}</p>

              <div>
                <h5 class="mt-3">访问控制</h5>
                <div class="row g-2">
                  <!-- 策略属性 -->
                  <div v-if="newMetadata.policy" class="col-12">
                    <div class="attribute-box p-2 border rounded bg-light">
                      <div class="text-muted small">访问策略</div>
                      <div class="fw-bold">{{ newMetadata.policy }}</div>
                    </div>
                  </div>

                  <!-- 密文 -->
                  <div v-if="newMetadata.ciphertext" class="col-12 mt-2">
                    <div class="attribute-box p-2 border rounded bg-light">
                      <div class="text-muted small">密文</div>
                      <div class="fw-bold">{{ newMetadata.ciphertext.length > 20 ? newMetadata.ciphertext.substring(0,
                        20) + '...' : newMetadata.ciphertext }}</div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div v-else class="text-center p-5 text-muted">
              <i class="bi bi-card-image fs-1"></i>
              <p class="mt-3">填写表单以查看预览</p>
            </div>
          </div>
        </div>

        <div class="card mt-3" v-if="metadataUrl">
          <div class="card-header">
            元数据 URL
          </div>
          <div class="card-body">
            <div class="input-group">
              <input type="text" class="form-control" :value="metadataUrl" readonly>
              <button class="btn btn-outline-secondary" type="button" @click="copyToClipboard(metadataUrl)">
                <i class="bi bi-clipboard"></i>
              </button>
            </div>
            <div class="mt-3">
              <a :href="metadataUrl" target="_blank" class="btn btn-sm btn-outline-primary">
                <i class="bi bi-box-arrow-up-right me-1"></i>打开链接
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'MetadataManager',
  data() {
    return {
      newMetadata: {
        name: '',
        description: '',
        image: '',
        policy: '',
        ciphertext: 'none',
      },
      isSubmitting: false,
      metadataUrl: ''
    };
  },
  methods: {
    handleImageError(event) {
      if (!event.target.src.includes('placeholder')) {
        event.target.src = 'data:image/svg+xml;charset=UTF-8,%3Csvg%20width%3D%22200%22%20height%3D%22200%22%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%3E%3Crect%20width%3D%22100%25%22%20height%3D%22100%25%22%20fill%3D%22%23f0f0f0%22%2F%3E%3Ctext%20x%3D%2250%25%22%20y%3D%2250%25%22%20font-family%3D%22Arial%22%20font-size%3D%2224%22%20text-anchor%3D%22middle%22%20dominant-baseline%3D%22middle%22%20fill%3D%22%23999%22%3ENFT%20%E5%8D%A0%E4%BD%8D%E5%9B%BE%3C%2Ftext%3E%3C%2Fsvg%3E';
        event.target.onerror = null;
      }
    },
    async createMetadata() {
      if (!this.newMetadata.name || !this.newMetadata.description || !this.newMetadata.image || !this.newMetadata.policy || !this.newMetadata.ciphertext) {
        this.$store.commit('notifications/add', {
          type: 'danger',
          message: '请填写所有必填字段'
        });
        return;
      }

      this.isSubmitting = true;

      try {
        const metadata = {
          name: this.newMetadata.name,
          description: this.newMetadata.description,
          image: this.newMetadata.image,
          external_url: '',
          policy: this.newMetadata.policy,
          ciphertext: this.newMetadata.ciphertext
        };

        console.log('提交的元数据:', metadata);

        // 调用API上传元数据
        const response = await fetch(`${this.$apiBaseUrl}/metadata`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(metadata)
        });

        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(errorData.error || '上传元数据失败');
        }

        const data = await response.json();
        this.metadataUrl = `ipfs://${data.ipfs_hash}`;

        this.$store.commit('notifications/add', {
          type: 'success',
          message: '元数据创建成功! IPFS哈希: ' + data.ipfs_hash
        });
      } catch (error) {
        console.error('创建元数据失败:', error);
        this.$store.commit('notifications/add', {
          type: 'danger',
          message: '创建元数据失败: ' + error.message
        });
      } finally {
        this.isSubmitting = false;
      }
    },
    copyToClipboard(text) {
      navigator.clipboard.writeText(text).then(
        () => {
          this.$store.commit('notifications/add', {
            type: 'success',
            message: 'URL已复制到剪贴板'
          });
        },
        () => {
          this.$store.commit('notifications/add', {
            type: 'danger',
            message: '复制到剪贴板失败'
          });
        }
      );
    }
  }
};
</script>

<style scoped>
.attribute-box {
  background-color: rgba(0, 0, 0, 0.03);
  transition: all 0.2s;
}

.attribute-box:hover {
  background-color: rgba(0, 0, 0, 0.05);
}

.metadata-preview img {
  max-height: 300px;
  object-fit: contain;
  width: 100%;
}
</style>