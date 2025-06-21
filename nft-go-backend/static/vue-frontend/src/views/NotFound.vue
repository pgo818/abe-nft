<template>
  <div class="not-found">
    <div class="container text-center py-5">
      <div class="error-code">404</div>
      <h1 class="mt-4">页面未找到</h1>
      <p class="lead">您访问的页面不存在或已被移除</p>
      <div class="mt-4">
        <router-link to="/" class="btn btn-primary">
          <i class="bi bi-house-door me-2"></i>返回首页
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'NotFound'
}
</script>

<style scoped>
.not-found {
  min-height: 80vh;
  display: flex;
  align-items: center;
  justify-content: center;
}

.error-code {
  font-size: 8rem;
  font-weight: 700;
  color: #e9ecef;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
}
</style> 