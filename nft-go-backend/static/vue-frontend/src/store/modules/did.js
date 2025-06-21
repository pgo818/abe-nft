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