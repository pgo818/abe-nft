import abeService from '@/services/abeService';

// ABE状态模块
export default {
    namespaced: true,

    state: {
        systemStatus: {
            initialized: false,
            attributeCount: 0,
            keyCount: 0,
            setupTime: null
        },
        logs: [],
        isLoading: false,
        error: null
    },

    mutations: {
        SET_SYSTEM_STATUS(state, status) {
            state.systemStatus = status;
        },
        SET_LOGS(state, logs) {
            state.logs = logs;
        },
        SET_LOADING(state, isLoading) {
            state.isLoading = isLoading;
        },
        SET_ERROR(state, error) {
            state.error = error;
        }
    },

    actions: {
        // 获取系统状态
        async getSystemStatus({ commit }) {
            commit('SET_LOADING', true);
            try {
                const status = await abeService.getSystemStatus();
                commit('SET_SYSTEM_STATUS', status);
                return status;
            } catch (error) {
                commit('SET_ERROR', error.message || '获取系统状态失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },
        
        // 初始化系统
        async setupSystem({ commit }, setupData) {
            commit('SET_LOADING', true);
            try {
                const result = await abeService.setupSystem(setupData);
                // 更新系统状态
                commit('SET_SYSTEM_STATUS', {
                    initialized: true,
                    attributeCount: setupData.universe ? setupData.universe.split('\n').filter(Boolean).length : 0,
                    keyCount: 0,
                    setupTime: new Date().toISOString()
                });
                return result;
            } catch (error) {
                commit('SET_ERROR', error.message || '初始化系统失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },
        
        // 生成密钥
        async generateKey({ commit }, keyData) {
            commit('SET_LOADING', true);
            try {
                const result = await abeService.generateKey(keyData);
                return result;
            } catch (error) {
                commit('SET_ERROR', error.message || '生成密钥失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },
        
        // 加载日志
        async loadLogs({ commit }, { page = 1, limit = 10 } = {}) {
            commit('SET_LOADING', true);
            try {
                const logs = await abeService.getLogs(page, limit);
                commit('SET_LOGS', logs);
                return logs;
            } catch (error) {
                commit('SET_ERROR', error.message || '加载日志失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        }
    },

    getters: {
        isSystemInitialized: state => state.systemStatus.initialized,
        systemAttributes: state => state.systemStatus.attributeCount,
        systemKeyCount: state => state.systemStatus.keyCount
    }
} 