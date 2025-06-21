// Wallet状态模块
export default {
    namespaced: true,

    state: {
        isConnected: false,
        account: null,
        did: null,
        showConnectPrompt: false,
        redirectRoute: null
    },

    mutations: {
        setConnected(state, isConnected) {
            state.isConnected = isConnected
        },
        setAccount(state, account) {
            state.account = account
        },
        setDID(state, did) {
            state.did = did
        },
        setShowConnectPrompt(state, show) {
            state.showConnectPrompt = show
        },
        setRedirectRoute(state, route) {
            state.redirectRoute = route
        },
        clearRedirectRoute(state) {
            state.redirectRoute = null
        }
    },

    actions: {
        // 连接钱包
        async connectWallet({ commit, dispatch }) {
            try {
                // 检查是否安装了MetaMask
                if (!window.ethereum) {
                    throw new Error('请安装MetaMask钱包')
                }

                // 请求连接钱包
                const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' })
                const account = accounts[0]

                if (account) {
                    // 保存账户信息
                    commit('setAccount', account)
                    commit('setConnected', true)

                    // 保存到localStorage
                    localStorage.setItem('walletConnected', 'true')
                    localStorage.setItem('walletAccount', account)

                    // 尝试获取或创建DID
                    dispatch('fetchOrCreateDID', account)

                    return account
                }

                return null
            } catch (error) {
                console.error('连接钱包失败:', error)
                dispatch('app/showError', '连接钱包失败: ' + error.message, { root: true })
                return null
            }
        },

        // 断开钱包连接
        disconnectWallet({ commit }) {
            commit('setConnected', false)
            commit('setAccount', null)
            commit('setDID', null)

            // 清除localStorage
            localStorage.removeItem('walletConnected')
            localStorage.removeItem('walletAccount')
            localStorage.removeItem('userDID')
        },

        // 检查钱包连接状态
        async checkWalletConnection({ commit, dispatch }) {
            try {
                // 首先尝试从localStorage恢复状态
                const wasConnected = localStorage.getItem('walletConnected') === 'true'
                const savedAccount = localStorage.getItem('walletAccount')
                const savedDID = localStorage.getItem('userDID')

                if (wasConnected && savedAccount) {
                    // 检查MetaMask当前连接的账户
                    const accounts = await window.ethereum.request({ method: 'eth_accounts' })

                    if (accounts.length > 0) {
                        const currentAccount = accounts[0]

                        if (savedAccount === currentAccount) {
                            // localStorage中保存的账户与MetaMask当前账户一致
                            commit('setAccount', currentAccount)
                            commit('setConnected', true)
                            if (savedDID) {
                                commit('setDID', savedDID)
                            } else {
                                // 尝试获取DID
                                dispatch('fetchOrCreateDID', currentAccount)
                            }
                            return true
                        } else {
                            // 账户不一致，使用新账户
                            commit('setAccount', currentAccount)
                            commit('setConnected', true)
                            localStorage.setItem('walletAccount', currentAccount)

                            // 尝试获取DID
                            dispatch('fetchOrCreateDID', currentAccount)
                            return true
                        }
                    } else {
                        // MetaMask没有连接账户，清除过期状态
                        localStorage.removeItem('walletConnected')
                        localStorage.removeItem('walletAccount')
                    }
                }

                return false
            } catch (error) {
                console.error('检查钱包连接失败:', error)
                return false
            }
        },

        // 获取或创建DID
        async fetchOrCreateDID({ commit, dispatch }, walletAddress) {
            try {
                // 先尝试获取现有DID
                const response = await fetch(`/api/did/wallet/${walletAddress}`)

                if (response.ok) {
                    const data = await response.json()
                    if (data.did) {
                        // 保存DID
                        commit('setDID', data.did)
                        localStorage.setItem('userDID', data.did)
                        return data.did
                    }
                }

                // 如果没有找到DID，则创建新DID
                const createResponse = await fetch(`/api/did/wallet/${walletAddress}`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    }
                })

                if (createResponse.ok) {
                    const data = await createResponse.json()
                    if (data.did) {
                        // 保存DID
                        commit('setDID', data.did)
                        localStorage.setItem('userDID', data.did)
                        return data.did
                    }
                }

                throw new Error('创建DID失败')
            } catch (error) {
                console.error('获取或创建DID失败:', error)
                dispatch('app/showError', '获取或创建DID失败: ' + error.message, { root: true })
                return null
            }
        },

        // 签名消息
        async signMessage({ state, dispatch }, message) {
            try {
                if (!state.isConnected || !state.account) {
                    throw new Error('钱包未连接')
                }

                // 使用MetaMask签名消息
                return await window.ethereum.request({
                    method: 'personal_sign',
                    params: [message, state.account]
                })
            } catch (error) {
                console.error('签名消息失败:', error)
                dispatch('app/showError', '签名消息失败: ' + error.message, { root: true })
                throw error
            }
        }
    },

    getters: {
        // 获取格式化的钱包地址（简短显示）
        shortAccount: state => {
            if (!state.account) return ''
            return `${state.account.substring(0, 6)}...${state.account.substring(state.account.length - 4)}`
        }
    }
} 