import api from '../utils/api';

const abeService = {
  // 获取ABE系统状态
  getSystemStatus: async () => {
    try {
      const response = await api.get('/abe/status');
      return response.data;
    } catch (error) {
      // 如果状态接口不存在，返回默认状态
      return {
        initialized: true,
        message: 'ABE系统已就绪'
      };
    }
  },

  // 初始化ABE系统
  setupSystem: async (data) => {
    const response = await api.post('/abe/setup', {
      gamma: data.gamma || data.attributes || []
    });
    return response.data;
  },

  // 生成用户密钥
  async generateKey({ wallet_address, attributes }) {
    try {
      console.log('ABE Service: 发送密钥生成请求', { wallet_address, attributes })

      // 验证钱包地址
      if (!wallet_address || wallet_address.length !== 42 || !wallet_address.startsWith('0x')) {
        throw new Error('钱包地址格式错误')
      }

      // 验证属性格式（如果提供了属性）
      if (attributes && attributes.length > 0) {
        for (const attr of attributes) {
          if (!attr.startsWith('mainNFT:')) {
            throw new Error('属性格式错误，必须是 mainNFT:NFT主地址 格式')
          }

          const nftAddress = attr.replace('mainNFT:', '')
          if (!nftAddress || nftAddress.length !== 42 || !nftAddress.startsWith('0x')) {
            throw new Error('NFT主地址格式错误')
          }
        }
      }

      const requestData = {
        wallet_address
      }

      // 如果提供了属性，则包含在请求中
      if (attributes && attributes.length > 0) {
        requestData.attributes = attributes
      }

      const response = await api.post('/abe/keygen', requestData)

      console.log('ABE Service: 密钥生成响应', response.data)
      return response.data
    } catch (error) {
      console.error('ABE Service: 生成密钥失败', error)
      if (error.response?.data?.error) {
        throw new Error(error.response.data.error)
      }
      throw new Error('生成密钥失败: ' + error.message)
    }
  },

  // 加密数据 - 验证策略格式
  encryptData: async (data) => {
    if (!data.message || !data.policy) {
      throw new Error('消息和策略不能为空');
    }

    // 验证策略格式必须是mainNFT:钱包地址
    if (!data.policy.startsWith('mainNFT:')) {
      throw new Error('访问策略格式错误，必须是 mainNFT:钱包地址 格式');
    }

    // 提取钱包地址并验证格式
    const walletAddress = data.policy.replace('mainNFT:', '');
    if (!/^0x[a-fA-F0-9]{40}$/.test(walletAddress)) {
      throw new Error('策略中的钱包地址格式错误');
    }

    const response = await api.post('/abe/encrypt', {
      message: data.message,
      policy: data.policy
    });
    return response.data;
  },

  // 解密数据
  decryptData: async (data) => {
    if (!data.cipher || !data.attrib_keys) {
      throw new Error('密文和用户密钥不能为空');
    }

    const response = await api.post('/abe/decrypt', {
      cipher: data.cipher,
      attrib_keys: data.attrib_keys
    });
    return response.data;
  },

  // 获取操作日志
  getLogs: async (page = 1, limit = 10) => {
    try {
      const response = await api.get(`/abe/logs?page=${page}&limit=${limit}`);
      return response.data;
    } catch (error) {
      // 如果日志接口不存在，返回空数组
      return { logs: [], total: 0 };
    }
  },

  // 验证钱包地址格式
  validateWalletAddress: (address) => {
    return /^0x[a-fA-F0-9]{40}$/.test(address);
  },

  // 生成策略字符串
  generatePolicy: (walletAddress) => {
    if (!abeService.validateWalletAddress(walletAddress)) {
      throw new Error('钱包地址格式错误');
    }
    return `mainNFT:${walletAddress}`;
  },

  // 上传图片到IPFS
  uploadImage: async (imageFile) => {
    try {
      if (!imageFile) {
        throw new Error('请选择图片文件');
      }

      // 验证文件类型
      const allowedTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif', 'image/webp', 'image/svg+xml', 'image/bmp'];
      if (!allowedTypes.includes(imageFile.type)) {
        throw new Error('不支持的图片格式，支持: JPG, PNG, GIF, WebP, SVG, BMP');
      }

      // 验证文件大小 (10MB)
      const maxSize = 10 * 1024 * 1024; // 10MB
      if (imageFile.size > maxSize) {
        throw new Error('图片文件大小不能超过10MB');
      }

      // 创建FormData
      const formData = new FormData();
      formData.append('image', imageFile);

      // 发送请求，注意这里不使用api实例，因为需要发送FormData
      const response = await fetch('/api/abe/upload-image', {
        method: 'POST',
        body: formData
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || '上传失败');
      }

      const result = await response.json();
      console.log('图片上传成功:', result);

      return result;
    } catch (error) {
      console.error('上传图片失败:', error);
      throw new Error('上传图片失败: ' + error.message);
    }
  }
};

export default abeService; 