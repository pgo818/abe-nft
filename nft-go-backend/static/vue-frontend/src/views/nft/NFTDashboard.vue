<template>
  <div class="nft-dashboard">
    <!-- NFT平台导航 -->
    <nav class="navbar navbar-expand-lg navbar-light bg-light mb-4">
      <div class="container">
        <span class="navbar-brand">
          <i class="bi bi-collection me-2"></i>NFT管理
        </span>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
          data-bs-target="#nftNavbarNav" aria-controls="nftNavbarNav" aria-expanded="false"
          aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="nftNavbarNav">
          <div class="navbar-nav">
            <router-link class="nav-link" :to="{ name: 'AllNFTs' }">
              <i class="bi bi-grid me-1"></i>所有NFT
            </router-link>
            <router-link class="nav-link" :to="{ name: 'MintNFT' }">
              <i class="bi bi-plus-circle me-1"></i>铸造NFT
            </router-link>
            <router-link class="nav-link" :to="{ name: 'MetadataManager' }">
              <i class="bi bi-file-text me-1"></i>创建元数据
            </router-link>
            <router-link class="nav-link" :to="{ name: 'MyNFTs' }">
              <i class="bi bi-wallet me-1"></i>我的NFT
            </router-link>
            <router-link class="nav-link" :to="{ name: 'RequestManager' }">
              <i class="bi bi-list-check me-1"></i>申请管理
            </router-link>
          </div>
        </div>
      </div>
    </nav>

    <!-- 主内容区 -->
    <div class="container">
      <router-view />
    </div>
    
    <!-- 结果提示模态框 -->
    <div class="modal fade" id="result-modal" tabindex="-1">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="result-title">操作结果</h5>
            <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
          </div>
          <div class="modal-body">
            <p id="result-message"></p>
            <div id="tx-hash-container" class="d-none">
              <p>交易哈希:</p>
              <code id="tx-hash" class="d-block p-2 bg-light"></code>
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
export default {
  name: 'NFTDashboard'
}
</script>

<style scoped>
.nft-dashboard {
  min-height: 80vh;
}

.navbar-brand {
  font-weight: 600;
}

.nav-link.router-link-active {
  color: #0d6efd;
  font-weight: 500;
}
</style> 