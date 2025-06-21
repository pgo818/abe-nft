// App状态模块
export default {
    namespaced: true,

    state: {
        isLoading: false,
        notifications: [],
        appTitle: 'NFT+ABE集成管理平台'
    },

    mutations: {
        setLoading(state, isLoading) {
            state.isLoading = isLoading
        },
        addNotification(state, notification) {
            // 添加通知，包含类型、消息和ID
            const id = Date.now()
            state.notifications.push({
                id,
                type: notification.type || 'info',
                message: notification.message,
                autoClose: notification.autoClose !== false
            })

            // 自动关闭通知
            if (notification.autoClose !== false) {
                setTimeout(() => {
                    this.commit('app/removeNotification', id)
                }, notification.duration || 5000)
            }
        },
        removeNotification(state, id) {
            const index = state.notifications.findIndex(n => n.id === id)
            if (index !== -1) {
                state.notifications.splice(index, 1)
            }
        },
        clearNotifications(state) {
            state.notifications = []
        }
    },

    actions: {
        showSuccess({ commit }, message) {
            commit('addNotification', {
                type: 'success',
                message
            })
        },
        showError({ commit }, message) {
            commit('addNotification', {
                type: 'error',
                message,
                duration: 8000
            })
        },
        showInfo({ commit }, message) {
            commit('addNotification', {
                type: 'info',
                message
            })
        },
        showWarning({ commit }, message) {
            commit('addNotification', {
                type: 'warning',
                message,
                duration: 7000
            })
        }
    }
} 