<template>
  <div class="abe-dashboard">
    <!-- ABE平台导航 -->
    <nav class="navbar navbar-expand-lg navbar-light bg-light mb-4">
      <div class="container">
        <span class="navbar-brand">
          <i class="bi bi-shield-lock me-2"></i>ABE管理
        </span>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#abeNavbarNav"
          aria-controls="abeNavbarNav" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="abeNavbarNav">
          <div class="navbar-nav">
            <router-link class="nav-link" :to="{ name: 'ABESetup' }">
              <i class="bi bi-gear me-1"></i>系统初始化
            </router-link>
            <router-link class="nav-link" :to="{ name: 'ABEKeyGen' }">
              <i class="bi bi-key me-1"></i>密钥生成
            </router-link>
            <!-- 暂未实现的功能 -->
            <a class="nav-link disabled" href="#" tabindex="-1" aria-disabled="true">
              <i class="bi bi-lock me-1"></i>加密
            </a>
            <a class="nav-link disabled" href="#" tabindex="-1" aria-disabled="true">
              <i class="bi bi-unlock me-1"></i>解密
            </a>
            <a class="nav-link disabled" href="#" tabindex="-1" aria-disabled="true">
              <i class="bi bi-journal-text me-1"></i>操作日志
            </a>
          </div>
        </div>
      </div>
    </nav>

    <!-- 主内容区 -->
    <div class="container">
      <router-view />
    </div>
  </div>
</template>

<script>
export default {
  name: 'ABEDashboard'
}
</script>

<style scoped>
.abe-dashboard {
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