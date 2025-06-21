<template>
  <div class="did-dashboard">
    <!-- DID平台导航 -->
    <nav class="navbar navbar-expand-lg navbar-light bg-light mb-4">
      <div class="container">
        <span class="navbar-brand">
          <i class="bi bi-person-badge me-2"></i>DID管理
        </span>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#didNavbarNav"
          aria-controls="didNavbarNav" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="didNavbarNav">
          <div class="navbar-nav">
            <router-link class="nav-link" :to="{ name: 'DIDList' }">
              <i class="bi bi-list-ul me-1"></i>DID列表
            </router-link>
            <router-link class="nav-link" :to="{ name: 'DIDCreate' }">
              <i class="bi bi-person-plus me-1"></i>创建DID
            </router-link>
            <router-link class="nav-link" :to="{ name: 'DoctorDID' }">
              <i class="bi bi-hospital me-1"></i>医生DID
            </router-link>
            <!-- 暂未实现的功能 -->
            <a class="nav-link disabled" href="#" tabindex="-1" aria-disabled="true">
              <i class="bi bi-file-earmark-check me-1"></i>凭证颁发
            </a>
            <a class="nav-link disabled" href="#" tabindex="-1" aria-disabled="true">
              <i class="bi bi-card-checklist me-1"></i>凭证管理
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
  name: 'DIDDashboard'
}
</script>

<style scoped>
.did-dashboard {
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