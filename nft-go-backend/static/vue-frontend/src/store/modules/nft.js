// NFT状态模块
import nftService from '@/services/nftService'
import { createSignMessage } from '@/utils/api'

export default {
    namespaced: true,

    state: {
        allNFTs: [],
        myNFTs: [],
        metadata: [],
        requests: [],
        selectedNFT: null,
        isLoading: false,
        error: null
    },

    mutations: {
        setAllNFTs(state, nfts) {
            state.allNFTs = nfts
        },
        setMyNFTs(state, nfts) {
            state.myNFTs = nfts
        },
        setMetadata(state, metadata) {
            state.metadata = metadata
        },
        setRequests(state, requests) {
            state.requests = requests
        },
        setSelectedNFT(state, nft) {
            state.selectedNFT = nft
        },
        SET_LOADING(state, isLoading) {
            state.isLoading = isLoading
        },
        SET_ERROR(state, error) {
            state.error = error
        },
        updateNFTMetadata(state, { tokenId, metadata }) {
            // 更新allNFTs中的元数据
            const nftIndex = state.allNFTs.findIndex(nft => nft.tokenId === tokenId)
            if (nftIndex !== -1) {
                state.allNFTs[nftIndex].metadata = metadata
            }

            // 更新myNFTs中的元数据
            const myNftIndex = state.myNFTs.findIndex(nft => nft.tokenId === tokenId)
            if (myNftIndex !== -1) {
                state.myNFTs[myNftIndex].metadata = metadata
            }
        }
    },

    actions: {
        // 加载所有NFT
        async loadAllNFTs({ commit, dispatch }) {
            commit('SET_LOADING', true)
            try {
                const response = await nftService.getAllNFTs()
                console.log('NFT数据:', response)
                const nfts = response.nfts || []

                // 处理每个NFT的元数据
                const processedNfts = nfts.map(nft => {
                    // 如果后端已经提供了元数据，直接使用
                    if (nft.metadata) {
                        return nft
                    }

                    // 否则，尝试从URI生成一个基本的元数据
                    return {
                        ...nft,
                        metadata: {
                            name: `NFT #${nft.tokenId}`,
                            description: 'This is an NFT created on our platform',
                            image: nft.uri,
                            attributes: [
                                { trait_type: 'Type', value: nft.contractType === 'child' ? 'Child NFT' : 'Main NFT' },
                                { trait_type: 'Rarity', value: 'Common' }
                            ]
                        }
                    }
                })

                commit('setAllNFTs', processedNfts)

                // 加载每个NFT的元数据
                dispatch('loadNFTsMetadata', processedNfts)

                return processedNfts
            } catch (error) {
                commit('SET_ERROR', error.message || '加载NFT失败')
                throw error
            } finally {
                commit('SET_LOADING', false)
            }
        },

        // 加载我的NFT
        async loadMyNFTs({ commit, rootState, dispatch }) {
            commit('SET_LOADING', true)
            try {
                const address = rootState.wallet.account
                if (!address) {
                    throw new Error('钱包未连接')
                }

                // 创建要签名的消息
                const message = createSignMessage('get_my_nfts', {
                    address: address,
                    timestamp: Date.now()
                })

                // 获取签名 - 使用dispatch调用wallet模块的signMessage action
                const signature = await dispatch('wallet/signMessage', message, { root: true })

                const response = await nftService.getMyNFTs(address, signature, message)
                console.log('我的NFT数据:', response)
                const myNfts = response.nfts || []

                // 处理每个NFT的元数据
                const processedNfts = myNfts.map(nft => {
                    // 如果后端已经提供了元数据，直接使用
                    if (nft.metadata) {
                        return nft
                    }

                    // 否则，尝试从URI生成一个基本的元数据
                    return {
                        ...nft,
                        metadata: {
                            name: `NFT #${nft.tokenId}`,
                            description: 'This is an NFT created on our platform',
                            image: nft.uri,
                            attributes: [
                                { trait_type: 'Type', value: nft.contractType === 'child' ? 'Child NFT' : 'Main NFT' },
                                { trait_type: 'Rarity', value: 'Common' }
                            ]
                        }
                    }
                })

                commit('setMyNFTs', processedNfts)

                // 加载每个NFT的元数据
                dispatch('loadNFTsMetadata', processedNfts)

                return processedNfts
            } catch (error) {
                commit('SET_ERROR', error.message || '加载我的NFT失败')
                throw error
            } finally {
                commit('SET_LOADING', false)
            }
        },

        // 批量加载NFT元数据
        async loadNFTsMetadata({ dispatch }, nfts) {
            console.log('开始加载NFT元数据:', nfts)
            // 对每个NFT异步加载元数据，不阻塞主流程
            nfts.forEach(nft => {
                if (nft.uri) {
                    dispatch('loadNFTMetadata', nft)
                }
            })
        },

        // 加载单个NFT的元数据
        async loadNFTMetadata({ commit }, nft) {
            try {
                if (!nft.uri) return

                console.log('加载NFT元数据:', nft.tokenId, nft.uri)

                // 无论是否已有元数据，都尝试从URI加载最新的元数据
                // 如果URI本身是JSON字符串，或者是指向JSON的链接，都能正确处理
                const metadata = await nftService.fetchMetadataFromURI(nft.uri)
                console.log('获取到元数据:', nft.tokenId, metadata)

                if (metadata && Object.keys(metadata).length > 0) {
                    // 确保元数据有必要的字段
                    const enhancedMetadata = {
                        name: metadata.name || `NFT #${nft.tokenId}`,
                        description: metadata.description || 'This is an NFT created on our platform',
                        image: metadata.image || nft.uri,
                        attributes: metadata.attributes || [
                            { trait_type: 'Type', value: nft.contractType === 'child' ? 'Child NFT' : 'Main NFT' },
                            { trait_type: 'Rarity', value: 'Common' }
                        ]
                    }
                    commit('updateNFTMetadata', { tokenId: nft.tokenId, metadata: enhancedMetadata })
                } else {
                    // 如果无法获取元数据，创建一个基本的元数据对象
                    const basicMetadata = {
                        name: `NFT #${nft.tokenId}`,
                        description: 'This is an NFT created on our platform',
                        image: nft.uri,
                        attributes: [
                            { trait_type: 'Type', value: nft.contractType === 'child' ? 'Child NFT' : 'Main NFT' },
                            { trait_type: 'Rarity', value: 'Common' }
                        ]
                    }
                    commit('updateNFTMetadata', { tokenId: nft.tokenId, metadata: basicMetadata })
                }
            } catch (error) {
                console.error(`加载NFT ${nft.tokenId} 元数据失败:`, error)
                // 创建一个基本的元数据对象
                const fallbackMetadata = {
                    name: `NFT #${nft.tokenId}`,
                    description: 'This is an NFT created on our platform',
                    image: nft.uri,
                    attributes: [
                        { trait_type: 'Type', value: nft.contractType === 'child' ? 'Child NFT' : 'Main NFT' },
                        { trait_type: 'Rarity', value: 'Common' }
                    ]
                }
                commit('updateNFTMetadata', { tokenId: nft.tokenId, metadata: fallbackMetadata })
            }
        },

        // 铸造NFT
        async mintNFT({ commit, rootState, dispatch }, uri) {
            commit('SET_LOADING', true)
            try {
                const address = rootState.wallet.account
                if (!address) {
                    throw new Error('钱包未连接')
                }

                // 创建要签名的消息
                const message = createSignMessage('mint_nft', {
                    uri: uri,
                    timestamp: Date.now()
                })

                // 获取签名 - 使用dispatch调用wallet模块的signMessage action
                const signature = await dispatch('wallet/signMessage', message, { root: true })

                // 构建请求数据
                const mintData = {
                    address: address,
                    signature: signature,
                    message: message,
                    uri: uri
                }

                const result = await nftService.mintNFT(mintData)
                return result.transactionHash
            } catch (error) {
                commit('SET_ERROR', error.message || '铸造NFT失败')
                throw error
            } finally {
                commit('SET_LOADING', false)
            }
        },

        // 加载元数据列表
        async loadMetadata({ commit }) {
            commit('SET_LOADING', true)
            try {
                const result = await nftService.getMetadata()
                commit('setMetadata', result.metadata || [])
            } catch (error) {
                console.error('加载元数据列表出错:', error)
                commit('SET_ERROR', '加载元数据失败: ' + error.message)
            } finally {
                commit('SET_LOADING', false)
            }
        },

        // 创建元数据
        async createMetadata({ commit, dispatch }, metadata) {
            commit('SET_LOADING', true)
            try {
                const metadataData = {
                    name: metadata.name,
                    description: metadata.description,
                    external_url: metadata.externalUrl,
                    image: metadata.image,
                    policy: metadata.policy,
                    ciphertext: metadata.ciphertext
                }

                const result = await nftService.createMetadata(metadataData)

                commit('app/showSuccess', `元数据创建成功！IPFS哈希: ${result.ipfs_hash}`, { root: true })
                // 刷新元数据列表
                dispatch('nft/loadMetadata')
                return result.ipfs_hash
            } catch (error) {
                console.error('创建元数据出错:', error)
                commit('SET_ERROR', '创建元数据失败: ' + error.message)
                return null
            } finally {
                commit('SET_LOADING', false)
            }
        },

        // 加载申请列表
        async loadRequests({ commit, rootState, dispatch }) {
            commit('SET_LOADING', true)
            try {
                if (!rootState.wallet.isConnected) {
                    throw new Error('钱包未连接')
                }

                // 创建要签名的消息
                const message = createSignMessage('get_all_requests', {
                    timestamp: Date.now()
                })

                // 获取签名 - 使用dispatch调用wallet模块的signMessage action
                const signature = await dispatch('wallet/signMessage', message, { root: true })

                console.log('发送请求获取NFT申请列表，地址:', rootState.wallet.account)

                // 使用nftService获取请求
                const result = await nftService.getAllRequests(rootState.wallet.account, signature, message)
                console.log('从后端获取的请求数据:', result)

                // 确保返回的请求数据是一个数组，即使为空
                const requestsData = result.requests || []
                console.log('处理后的请求数据:', requestsData)

                commit('setRequests', requestsData)
                return requestsData
            } catch (error) {
                console.error('加载申请列表出错:', error)
                commit('SET_ERROR', '加载申请列表失败: ' + error.message)
                // 确保在错误情况下也设置一个空数组，而不是保留旧数据
                commit('setRequests', [])
                throw error
            } finally {
                commit('SET_LOADING', false)
            }
        },

        // 处理申请
        async processRequest({ commit, rootState, dispatch }, { requestId, action }) {
            commit('SET_LOADING', true)
            try {
                if (!rootState.wallet.isConnected) {
                    throw new Error('钱包未连接')
                }

                console.log('处理请求:', requestId, action);

                // 创建要签名的消息
                const message = createSignMessage('process_request', {
                    requestId: requestId,
                    decision: action,
                    timestamp: Date.now()
                })

                // 获取签名 - 使用dispatch调用wallet模块的signMessage action
                const signature = await dispatch('wallet/signMessage', message, { root: true })

                // 构建请求数据
                const requestData = {
                    address: rootState.wallet.account,
                    signature: signature,
                    message: message,
                    requestId: String(requestId), // 确保requestId作为字符串发送
                    action: action
                }

                console.log('发送请求数据:', requestData);

                // 使用nftService处理请求
                const result = await nftService.processRequest(requestData)

                const actionText = action === 'approve' ? '批准' : '拒绝'
                commit('app/showSuccess', `申请已${actionText}`, { root: true })

                // 刷新申请列表
                dispatch('loadRequests')

                return result.transactionHash
            } catch (error) {
                console.error('处理申请出错:', error)
                commit('SET_ERROR', '处理申请失败: ' + error.message)
                return null
            } finally {
                commit('SET_LOADING', false)
            }
        },

        // 申请子NFT
        async requestChildNFT({ commit, rootState, dispatch }, { parentTokenId, uri }) {
            commit('SET_LOADING', true)
            try {
                if (!rootState.wallet.isConnected) {
                    throw new Error('钱包未连接')
                }

                // 创建要签名的消息
                const message = createSignMessage('request_child_nft', {
                    parentTokenId: parentTokenId,
                    uri: uri,
                    timestamp: Date.now()
                })

                // 获取签名 - 使用dispatch调用wallet模块的signMessage action
                const signature = await dispatch('wallet/signMessage', message, { root: true })

                // 构建请求数据
                const requestData = {
                    address: rootState.wallet.account,
                    signature: signature,
                    message: message,
                    parentTokenId: parentTokenId,
                    applicantAddress: rootState.wallet.account,
                    uri: uri,
                    description: `申请创建父NFT ${parentTokenId} 的子NFT`
                }

                // 使用nftService申请子NFT
                // eslint-disable-next-line no-unused-vars
                const result = await nftService.requestChildNFT(requestData)

                commit('app/showSuccess', '子NFT申请已提交，等待审批', { root: true })

                // 刷新请求列表
                dispatch('loadRequests')

                return true
            } catch (error) {
                console.error('提交申请出错:', error)
                commit('SET_ERROR', '提交申请失败: ' + error.message)
                return false
            } finally {
                commit('SET_LOADING', false)
            }
        },

        // 创建子NFT
        async createChildNFT({ commit }, childData) {
            commit('SET_LOADING', true)
            try {
                const result = await nftService.createChildNFT(childData)
                return result
            } catch (error) {
                commit('SET_ERROR', error.message || '创建子NFT失败')
                throw error
            } finally {
                commit('SET_LOADING', false)
            }
        },

        // 更新NFT元数据
        async updateMetadata({ commit }, metadataData) {
            commit('SET_LOADING', true)
            try {
                const result = await nftService.updateMetadata(metadataData)
                return result
            } catch (error) {
                commit('SET_ERROR', error.message || '更新元数据失败')
                throw error
            } finally {
                commit('SET_LOADING', false)
            }
        },

        // 批准请求
        async approveRequest({ commit, rootState, dispatch }, requestId) {
            commit('SET_LOADING', true)
            try {
                if (!rootState.wallet.isConnected) {
                    throw new Error('钱包未连接')
                }

                return await dispatch('processRequest', {
                    requestId: requestId,
                    action: 'approve'
                })
            } catch (error) {
                console.error('批准请求出错:', error)
                commit('SET_ERROR', '批准请求失败: ' + error.message)
                return false
            } finally {
                commit('SET_LOADING', false)
            }
        },

        // 拒绝请求
        async rejectRequest({ commit, rootState, dispatch }, requestId) {
            commit('SET_LOADING', true)
            try {
                if (!rootState.wallet.isConnected) {
                    throw new Error('钱包未连接')
                }

                return await dispatch('processRequest', {
                    requestId: requestId,
                    action: 'reject'
                })
            } catch (error) {
                console.error('拒绝请求出错:', error)
                commit('SET_ERROR', '拒绝请求失败: ' + error.message)
                return false
            } finally {
                commit('SET_LOADING', false)
            }
        }
    },

    getters: {
        // 获取待处理的申请
        pendingRequests: state => {
            return state.requests.filter(req => req.status === 'pending')
        },

        // 获取我的申请
        myRequests: (state, getters, rootState) => {
            if (!rootState.wallet.account) return []
            return state.requests.filter(req => req.applicantAddress === rootState.wallet.account)
        },

        getNFTById: (state) => (id) => {
            return state.allNFTs.find(nft => nft.tokenId === id) ||
                state.myNFTs.find(nft => nft.tokenId === id)
        }
    }
} 