import api from '../utils/api';

const didService = {
  // 获取DID列表
  getDIDs: async () => {
    return api.get('/did/list');
  },
  
  // 创建DID
  createDID: async (data) => {
    return api.post('/did/create', data);
  },
  
  // 获取DID详情
  getDIDDetails: async (did) => {
    return api.get(`/did/details?did=${encodeURIComponent(did)}`);
  },
  
  // 删除DID
  deleteDID: async (did) => {
    return api.delete(`/did/delete?did=${encodeURIComponent(did)}`);
  },
  
  // 颁发可验证凭证
  issueVC: async (data) => {
    return api.post('/did/vc/issue', data);
  },
  
  // 验证可验证凭证
  verifyVC: async (data) => {
    return api.post('/did/vc/verify', data);
  },
  
  // 获取医生DID信息
  getDoctorDID: async () => {
    return api.get('/did/doctor');
  },
  
  // 验证医生身份
  verifyDoctorIdentity: async (data) => {
    return api.post('/did/doctor/verify', data);
  }
};

export default didService; 