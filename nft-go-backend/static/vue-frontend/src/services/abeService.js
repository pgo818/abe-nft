import api from '../utils/api';

const abeService = {
  // 获取ABE系统状态
  getSystemStatus: async () => {
    return api.get('/abe/status');
  },
  
  // 初始化ABE系统
  setupSystem: async (data) => {
    return api.post('/abe/setup', data);
  },
  
  // 生成ABE密钥
  generateKey: async (data) => {
    return api.post('/abe/keygen', data);
  },
  
  // 加密数据
  encryptData: async (data) => {
    return api.post('/abe/encrypt', data);
  },
  
  // 解密数据
  decryptData: async (data) => {
    return api.post('/abe/decrypt', data);
  },
  
  // 获取操作日志
  getLogs: async (page = 1, limit = 10) => {
    return api.get(`/abe/logs?page=${page}&limit=${limit}`);
  }
};

export default abeService; 