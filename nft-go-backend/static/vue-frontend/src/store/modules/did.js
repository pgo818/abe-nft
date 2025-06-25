import didService from '@/services/didService';

// DID状态模块
export default {
    namespaced: true,

    state: {
        dids: [],
        selectedDID: null,
        vcs: [],
        selectedVC: null,
        doctorDID: null,
        doctorVCs: [],
        isLoading: false,
        error: null
    },

    mutations: {
        SET_DIDS(state, dids) {
            state.dids = dids;
        },
        SET_SELECTED_DID(state, did) {
            state.selectedDID = did;
        },
        SET_VCS(state, vcs) {
            state.vcs = vcs;
        },
        SET_SELECTED_VC(state, vc) {
            state.selectedVC = vc;
        },
        SET_DOCTOR_DID(state, doctorDID) {
            state.doctorDID = doctorDID;
        },
        SET_DOCTOR_VCS(state, doctorVCs) {
            state.doctorVCs = doctorVCs;
        },
        SET_LOADING(state, isLoading) {
            state.isLoading = isLoading;
        },
        SET_ERROR(state, error) {
            state.error = error;
        },
        ADD_DID(state, did) {
            state.dids.push(did);
        },
        REMOVE_DID(state, didId) {
            state.dids = state.dids.filter(did => did.id !== didId);
        }
    },

    actions: {
        // 获取DID列表
        async getDIDs({ commit }) {
            commit('SET_LOADING', true);
            try {
                const dids = await didService.getDIDs();
                commit('SET_DIDS', dids);
                return dids;
            } catch (error) {
                commit('SET_ERROR', error.message || '获取DID列表失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 创建DID
        async createDID({ commit }, didData) {
            commit('SET_LOADING', true);
            try {
                const result = await didService.createDID(didData);
                commit('ADD_DID', result);
                return result;
            } catch (error) {
                commit('SET_ERROR', error.message || '创建DID失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 获取DID详情
        async getDIDDetails({ commit }, did) {
            commit('SET_LOADING', true);
            try {
                const details = await didService.getDIDDetails(did);
                commit('SET_SELECTED_DID', details);
                return details;
            } catch (error) {
                commit('SET_ERROR', error.message || '获取DID详情失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 删除DID
        async deleteDID({ commit }, did) {
            commit('SET_LOADING', true);
            try {
                await didService.deleteDID(did);
                commit('REMOVE_DID', did.id);
                return true;
            } catch (error) {
                commit('SET_ERROR', error.message || '删除DID失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 颁发可验证凭证
        async issueVC({ commit }, vcData) {
            commit('SET_LOADING', true);
            try {
                const result = await didService.issueVC(vcData);
                return result;
            } catch (error) {
                commit('SET_ERROR', error.message || '颁发可验证凭证失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 验证可验证凭证
        async verifyVC({ commit }, vcData) {
            commit('SET_LOADING', true);
            try {
                const result = await didService.verifyVC(vcData);
                return result;
            } catch (error) {
                commit('SET_ERROR', error.message || '验证可验证凭证失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 获取医生DID信息
        async getDoctorDID({ commit }) {
            commit('SET_LOADING', true);
            try {
                const doctorDID = await didService.getDoctorDID();
                commit('SET_DOCTOR_DID', doctorDID);
                return doctorDID;
            } catch (error) {
                commit('SET_ERROR', error.message || '获取医生DID信息失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 验证医生身份
        async verifyDoctorIdentity({ commit }, data) {
            commit('SET_LOADING', true);
            try {
                const result = await didService.verifyDoctorIdentity(data);
                return result;
            } catch (error) {
                commit('SET_ERROR', error.message || '验证医生身份失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 创建医生DID
        async createDoctorDID({ commit, rootState }, doctorData) {
            commit('SET_LOADING', true);
            try {
                // 获取当前钱包地址
                const walletAddress = rootState.wallet.account;
                if (!walletAddress) {
                    throw new Error('请先连接钱包');
                }

                // 准备请求数据
                const requestData = {
                    walletAddress: walletAddress,
                    name: doctorData.name,
                    licenseNumber: doctorData.license || doctorData.doctorId
                };

                const result = await didService.createDoctorDID(requestData);

                // 自动颁发医生凭证
                if (result.data) {
                    const vcData = {
                        issuerDid: "0x1234", // 默认医院DID
                        doctorDid: result.data.did,
                        vcType: "医生执业资格",
                        vcContent: JSON.stringify({
                            name: doctorData.name,
                            licenseNumber: requestData.licenseNumber,
                            department: doctorData.department,
                            hospital: doctorData.hospital,
                            title: doctorData.title,
                            specialties: doctorData.specialties
                        })
                    };

                    try {
                        await didService.issueDoctorVC(vcData);
                    } catch (vcError) {
                        console.warn('自动颁发凭证失败:', vcError);
                    }
                }

                return result.data;
            } catch (error) {
                commit('SET_ERROR', error.response?.data?.error || error.message || '创建医生DID失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 颁发医生凭证
        async issueDoctorCredential({ commit }, credentialData) {
            commit('SET_LOADING', true);
            try {
                const vcData = {
                    issuerDid: "0x1234", // 默认医院DID
                    doctorDid: credentialData.did,
                    vcType: "医生执业资格",
                    vcContent: JSON.stringify({
                        name: credentialData.name,
                        doctorId: credentialData.doctorId,
                        department: credentialData.department,
                        hospital: credentialData.hospital,
                        title: credentialData.title,
                        license: credentialData.license,
                        specialties: credentialData.specialties
                    })
                };

                const result = await didService.issueDoctorVC(vcData);
                return result.data;
            } catch (error) {
                commit('SET_ERROR', error.response?.data?.error || error.message || '颁发医生凭证失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 加载医生DID列表
        async loadDoctorDIDs({ commit }) {
            commit('SET_LOADING', true);
            try {
                const result = await didService.getDoctorDIDs();
                // 后端返回格式为 {doctors: [...]}
                const doctors = Array.isArray(result.data?.doctors) ? result.data.doctors : [];
                return doctors;
            } catch (error) {
                console.warn('获取医生DID列表失败:', error);
                // 暂时返回空数组，避免阻塞页面
                if (error.response?.status === 404) {
                    return [];
                }
                commit('SET_ERROR', error.response?.data?.error || error.message || '获取医生DID列表失败');
                return []; // 确保出错时也返回空数组
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 获取VC详情
        async getVC({ commit }, vcId) {
            commit('SET_LOADING', true);
            try {
                // 这里需要根据实际API实现
                const result = await didService.verifyVC({ vcId });
                return result.data;
            } catch (error) {
                commit('SET_ERROR', error.response?.data?.error || error.message || '获取凭证失败');
                throw error;
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 获取钱包相关的DID列表
        async getDIDsByWallet({ commit, rootState }) {
            commit('SET_LOADING', true);
            try {
                const walletAddress = rootState.wallet.account;
                if (!walletAddress) {
                    throw new Error('请先连接钱包');
                }

                const result = await didService.getDIDsByWallet(walletAddress);
                // 后端返回格式为 {dids: [...]}
                const dids = Array.isArray(result.data?.dids) ? result.data.dids : [];
                commit('SET_DIDS', dids);
                return dids;
            } catch (error) {
                console.warn('获取DID列表失败:', error);
                // 暂时返回空数组，避免阻塞页面
                if (error.response?.status === 404) {
                    return [];
                }
                commit('SET_ERROR', error.response?.data?.error || error.message || '获取DID列表失败');
                return []; // 确保出错时也返回空数组
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 获取医生凭证列表
        async getDoctorVCs({ commit }, doctorDID) {
            commit('SET_LOADING', true);
            try {
                const result = await didService.getDoctorVCs(doctorDID);
                // 后端返回格式为 {doctorDid: "...", verifiableCredentials: [...]}
                const vcs = Array.isArray(result.data?.verifiableCredentials) ? result.data.verifiableCredentials : [];
                commit('SET_DOCTOR_VCS', vcs);
                return vcs;
            } catch (error) {
                console.warn('获取医生凭证失败:', error);
                commit('SET_ERROR', error.response?.data?.error || error.message || '获取医生凭证失败');
                return []; // 确保出错时也返回空数组
            } finally {
                commit('SET_LOADING', false);
            }
        },

        // 通过钱包地址加载医生VC凭证
        async loadDoctorVCs({ commit, dispatch }, { walletAddress }) {
            commit('SET_LOADING', true);
            try {
                console.log('开始加载医生VC凭证，钱包地址:', walletAddress);

                // 首先获取医生DID列表
                const doctors = await dispatch('loadDoctorDIDs');
                console.log('获取到的医生列表:', doctors);

                // 找到当前钱包地址对应的医生DID
                const currentDoctor = doctors.find(doctor =>
                    doctor.walletAddress && doctor.walletAddress.toLowerCase() === walletAddress.toLowerCase()
                );

                if (!currentDoctor) {
                    console.log('没有找到对应的医生DID，返回空数组');
                    commit('SET_DOCTOR_VCS', []);
                    return [];
                }

                console.log('找到医生DID:', currentDoctor.didString);

                // 获取该医生的VC凭证
                const vcs = await dispatch('getDoctorVCs', currentDoctor.didString);
                console.log('获取到的VC凭证:', vcs);

                return vcs;
            } catch (error) {
                console.error('加载医生VC凭证失败:', error);
                commit('SET_ERROR', error.message || '加载医生VC凭证失败');
                commit('SET_DOCTOR_VCS', []);
                return [];
            } finally {
                commit('SET_LOADING', false);
            }
        }
    },

    getters: {
        getDIDById: (state) => (id) => {
            return state.dids.find(did => did.id === id);
        },
        getDIDByDID: (state) => (did) => {
            return state.dids.find(item => item.did === did);
        },
        getVCById: (state) => (id) => {
            return state.vcs.find(vc => vc.id === id);
        }
    }
} 