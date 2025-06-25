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
        latestCiphertext: null,
        latestUserKey: null,
        savedUserKeys: [], // 保存的用户密钥列表
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
        setLatestCiphertext(state, ciphertext) {
            state.latestCiphertext = ciphertext;
        },
        setLatestUserKey(state, userKey) {
            state.latestUserKey = userKey;
        },
        SET_LOADING(state, isLoading) {
            state.isLoading = isLoading;
        },
        SET_ERROR(state, error) {
            state.error = error;
        },
        ADD_USER_KEY(state, key) {
            state.latestUserKey = key.attrib_keys;
        },
        SAVE_USER_KEY(state, keyData) {
            // 添加到保存的密钥列表
            const existingIndex = state.savedUserKeys.findIndex(k => k.id === keyData.id);
            if (existingIndex >= 0) {
                // 更新现有密钥
                state.savedUserKeys[existingIndex] = keyData;
            } else {
                // 添加新密钥
                state.savedUserKeys.push(keyData);
            }
            // 保存到localStorage
            localStorage.setItem('abeUserKeys', JSON.stringify(state.savedUserKeys));
        },
        LOAD_SAVED_KEYS(state) {
            // 从localStorage加载保存的密钥
            const saved = localStorage.getItem('abeUserKeys');
            if (saved) {
                try {
                    state.savedUserKeys = JSON.parse(saved);
                } catch (error) {
                    console.error('加载保存的密钥失败:', error);
                    state.savedUserKeys = [];
                }
            }
        },
        DELETE_SAVED_KEY(state, keyId) {
            state.savedUserKeys = state.savedUserKeys.filter(k => k.id !== keyId);
            // 更新localStorage
            localStorage.setItem('abeUserKeys', JSON.stringify(state.savedUserKeys));
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

        // 生成密钥 - 使用钱包地址
        async generateKey({ commit }, { wallet_address, attributes }) {
            try {
                console.log('ABE Store: 生成密钥请求', { wallet_address, attributes })

                const response = await abeService.generateKey({
                    wallet_address,
                    attributes
                })

                console.log('ABE Store: 密钥生成响应', response)

                if (response.user_key_id) {
                    commit('setLatestUserKey', response.attrib_keys)

                    const keyData = {
                        id: response.user_key_id,
                        wallet_address: response.wallet_address,
                        attributes: response.attributes,
                        nft_count: response.nft_count,
                        attrib_keys: response.attrib_keys,
                        created_at: new Date().toISOString(),
                        name: `密钥-${new Date().toLocaleString()}` // 默认名称
                    };

                    commit('ADD_USER_KEY', keyData)
                    // 自动保存密钥
                    commit('SAVE_USER_KEY', keyData)
                }

                return response
            } catch (error) {
                console.error('ABE Store: 生成密钥失败', error)
                throw error
            }
        },

        // 加密数据 - 使用固定的mainNFT:钱包地址格式
        async encryptData({ commit }, { message, policy, saveToNFT = false }) {
            commit('SET_LOADING', true);
            try {
                if (!policy || !policy.startsWith('mainNFT:')) {
                    throw new Error('访问策略格式错误，必须是 mainNFT:钱包地址 格式');
                }

                const result = await abeService.encryptData({
                    message,
                    policy
                });

                // 保存最新的密文
                commit('setLatestCiphertext', result.cipher);

                // 如果需要保存到NFT，则调用相关API
                if (saveToNFT && result.cipher) {
                    // 这里可以添加保存到NFT的逻辑
                    console.log('保存加密结果到NFT:', result);
                }

                return result;
            } catch (error) {
                commit('SET_ERROR', error.message || '加密数据失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 解密数据 - 使用钱包地址自动生成密钥
        async decryptData({ commit, rootGetters }, { ciphertext, walletAddress, autoGenerateKey = true }) {
            commit('SET_LOADING', true);
            try {
                let attribKeys = null;

                // 如果启用自动生成密钥，先为钱包地址生成用户密钥
                if (autoGenerateKey && walletAddress) {
                    const keyResult = await abeService.generateKey({
                        wallet_address: walletAddress
                    });
                    attribKeys = keyResult.attrib_keys;

                    // 保存最新的用户密钥
                    commit('setLatestUserKey', attribKeys);
                } else {
                    // 使用已有的用户密钥
                    attribKeys = this.state.abe.latestUserKey;
                    if (!attribKeys) {
                        throw new Error('没有可用的用户密钥，请先生成密钥或启用自动生成');
                    }
                }

                // 执行解密
                const result = await abeService.decryptData({
                    cipher: ciphertext,
                    attrib_keys: attribKeys
                });

                return result;
            } catch (error) {
                commit('SET_ERROR', error.message || '解密数据失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 直接解密（使用已有的密钥）
        async decryptDataDirect({ commit }, { cipher, attrib_keys }) {
            commit('SET_LOADING', true);
            try {
                if (!cipher || !attrib_keys) {
                    throw new Error('密文和用户密钥不能为空');
                }

                const result = await abeService.decryptData({
                    cipher,
                    attrib_keys
                });

                return result;
            } catch (error) {
                commit('SET_ERROR', error.message || '解密数据失败');
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
        },

        // 保存用户密钥
        saveUserKey({ commit }, keyData) {
            commit('SAVE_USER_KEY', keyData);
        },

        // 加载保存的密钥
        loadSavedKeys({ commit }) {
            commit('LOAD_SAVED_KEYS');
        },

        // 删除保存的密钥
        deleteSavedKey({ commit }, keyId) {
            commit('DELETE_SAVED_KEY', keyId);
        }
    },

    getters: {
        isSystemInitialized: state => state.systemStatus.initialized,
        systemAttributes: state => state.systemStatus.attributeCount,
        systemKeyCount: state => state.systemStatus.keyCount,
        latestCiphertext: state => state.latestCiphertext,
        latestUserKey: state => state.latestUserKey,
        savedUserKeys: state => state.savedUserKeys
    }
} 