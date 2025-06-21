<template>
  <div class="notifications-container">
    <transition-group name="notification">
      <div 
        v-for="notification in notifications" 
        :key="notification.id"
        :class="['notification', `notification-${notification.type}`]"
      >
        <div class="notification-content">
          <i :class="getIconClass(notification.type)" class="notification-icon"></i>
          <span class="notification-message">{{ notification.message }}</span>
        </div>
        <button 
          class="notification-close" 
          @click="removeNotification(notification.id)"
        >
          &times;
        </button>
      </div>
    </transition-group>
  </div>
</template>

<script>
import { computed } from 'vue'
import { useStore } from 'vuex'

export default {
  name: 'Notifications',
  
  setup() {
    const store = useStore()
    
    // 从store获取通知列表
    const notifications = computed(() => store.state.app.notifications)
    
    // 根据通知类型获取图标类名
    const getIconClass = (type) => {
      switch (type) {
        case 'success':
          return 'bi bi-check-circle-fill'
        case 'error':
          return 'bi bi-exclamation-circle-fill'
        case 'warning':
          return 'bi bi-exclamation-triangle-fill'
        case 'info':
        default:
          return 'bi bi-info-circle-fill'
      }
    }
    
    // 移除通知
    const removeNotification = (id) => {
      store.commit('app/removeNotification', id)
    }
    
    return {
      notifications,
      getIconClass,
      removeNotification
    }
  }
}
</script>

<style scoped>
.notifications-container {
  position: fixed;
  top: 20px;
  right: 20px;
  z-index: 9999;
  display: flex;
  flex-direction: column;
  gap: 10px;
  max-width: 350px;
}

.notification {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 15px;
  border-radius: 4px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  background-color: white;
  color: #333;
  animation: slideIn 0.3s ease-out;
}

.notification-content {
  display: flex;
  align-items: center;
  gap: 10px;
}

.notification-icon {
  font-size: 1.2rem;
}

.notification-message {
  font-size: 0.95rem;
}

.notification-close {
  background: none;
  border: none;
  font-size: 1.2rem;
  cursor: pointer;
  padding: 0 5px;
  color: #666;
}

.notification-close:hover {
  color: #333;
}

.notification-success {
  border-left: 4px solid #28a745;
}

.notification-success .notification-icon {
  color: #28a745;
}

.notification-error {
  border-left: 4px solid #dc3545;
}

.notification-error .notification-icon {
  color: #dc3545;
}

.notification-warning {
  border-left: 4px solid #ffc107;
}

.notification-warning .notification-icon {
  color: #ffc107;
}

.notification-info {
  border-left: 4px solid #17a2b8;
}

.notification-info .notification-icon {
  color: #17a2b8;
}

/* 过渡动画 */
.notification-enter-active,
.notification-leave-active {
  transition: all 0.3s ease;
}

.notification-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.notification-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
</style> 