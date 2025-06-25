import api from '@/utils/api';

const nftService = {
  // 获取所有NFT
  getAllNFTs: async () => {
    const response = await api.get('/nfts');
    return response.data;
  },

  // 获取我的NFT
  getMyNFTs: async (address, signature, message) => {
    const response = await api.get('/nft/my-nfts', {
      headers: {
        'X-Ethereum-Address': address,
        'X-Ethereum-Signature': signature,
        'X-Ethereum-Message': message
      }
    });
    return response.data;
  },

  // 铸造NFT
  mintNFT: async (data) => {
    const response = await api.post('/nft/mint', data);
    return response.data;
  },

  // 获取NFT详情
  getNFTDetails: async (tokenId) => {
    return api.get(`/nft/${tokenId}`);
  },

  // 创建子NFT
  createChildNFT: async (data) => {
    return api.post('/nft/createChild', data);
  },

  // 更新NFT元数据
  updateMetadata: async (data) => {
    return api.post('/nft/update-metadata', data);
  },

  // 更新NFT元数据URI
  updateNFTMetadataURI: async (data) => {
    const response = await api.post('/nft/update-uri', data);
    return response.data;
  },

  // 获取用户的NFT
  getUserNFTs: async (address) => {
    return api.get(`/nfts/user/${address}`);
  },

  // 获取所有请求
  getAllRequests: async (address, signature, message) => {
    const response = await api.get('/nft/all-requests', {
      headers: {
        'X-Ethereum-Address': address,
        'X-Ethereum-Signature': signature,
        'X-Ethereum-Message': message
      }
    });
    return response.data;
  },

  // 处理子NFT请求
  processRequest: async (data) => {
    return api.post('/nft/process-request', data);
  },

  // 申请子NFT
  requestChildNFT: async (data) => {
    console.log('申请子NFT请求数据:', data)
    const response = await api.post('/nft/request-child', data);
    console.log('申请子NFT响应:', response.data)
    return response.data;
  },

  // 获取元数据列表
  getMetadata: async () => {
    return api.get('/metadata');
  },

  // 创建元数据
  createMetadata: async (metadata) => {
    return api.post('/metadata', metadata);
  },

  // 获取NFT元数据
  getNFTMetadata: async (tokenId) => {
    const response = await api.get(`/nft/${tokenId}`);
    return response.data;
  },

  // 从URI获取NFT元数据
  async fetchMetadataFromURI(uri) {
    try {
      console.log('开始获取元数据，URI:', uri);

      // 处理空URI
      if (!uri || uri.trim() === '') {
        console.log('URI为空，返回默认元数据');
        return {
          name: '未命名NFT',
          description: 'This NFT has no metadata',
          image: '',
          attributes: [
            { trait_type: 'Type', value: 'Unknown' },
            { trait_type: 'Rarity', value: 'Unknown' }
          ]
        };
      }

      // 如果URI本身就是JSON字符串，直接解析
      if (uri.trim().startsWith('{') && uri.trim().endsWith('}')) {
        console.log('URI本身是JSON字符串，尝试直接解析');
        try {
          const metadata = JSON.parse(uri);
          console.log('成功解析JSON字符串:', metadata);

          // 如果图片URL是IPFS链接，转换为HTTP链接
          if (metadata.image && metadata.image.startsWith('ipfs://')) {
            metadata.image = metadata.image.replace('ipfs://', 'https://ipfs.io/ipfs/');
            console.log('转换图片IPFS链接为HTTP:', metadata.image);
          }

          return metadata;
        } catch (parseError) {
          console.error('解析URI中的JSON失败:', parseError);
        }
      }

      // IPFS网关列表 - 按优先级排序
      const ipfsGateways = [
        'https://dweb.link/ipfs/',
        'https://cloudflare-ipfs.com/ipfs/',
        'https://gateway.pinata.cloud/ipfs/',
        'https://ipfs.io/ipfs/',
        'https://cf-ipfs.com/ipfs/',
        'https://ipfs.fleek.co/ipfs/',
        'https://ipfs.infura.io/ipfs/'
      ];

      // 如果URI是IPFS链接，尝试多个网关
      if (uri.startsWith('ipfs://')) {
        // 从URI中提取IPFS哈希
        const ipfsHash = uri.substring('ipfs://'.length);
        console.log('提取IPFS哈希:', ipfsHash, '长度:', ipfsHash.length);

        // 使用完整哈希，但显示时可以截取
        const displayHash = ipfsHash.length > 8 ? `${ipfsHash.substring(0, 8)}...` : ipfsHash;

        const fallbackMetadata = {
          name: `NFT #${displayHash}`,
          description: 'This is an NFT created on our platform',
          image: uri,
          attributes: [
            { trait_type: 'Type', value: 'Main NFT' },
            { trait_type: 'Rarity', value: 'Common' }
          ]
        };

        // 尝试所有IPFS网关，直到成功或全部失败
        for (const gateway of ipfsGateways) {
          try {
            const gatewayUrl = `${gateway}${ipfsHash}`;
            console.log(`尝试IPFS网关: ${gateway}`);

            // 设置超时和重定向处理
            const controller = new AbortController();
            const timeoutId = setTimeout(() => controller.abort(), 10000); // 10秒超时

            const response = await fetch(gatewayUrl, {
              signal: controller.signal,
              headers: { 'Accept': 'application/json' },
              redirect: 'follow' // 自动跟随重定向
            });
            clearTimeout(timeoutId);

            if (!response.ok) {
              console.warn(`网关 ${gateway} 返回错误状态码: ${response.status}`);
              continue; // 尝试下一个网关
            }

            // 检查是否发生了重定向
            if (response.redirected) {
              console.log(`网关 ${gateway} 重定向到: ${response.url}`);
            }

            const contentType = response.headers.get('content-type');
            console.log(`网关 ${gateway} 响应内容类型:`, contentType);

            // 尝试获取响应文本
            const text = await response.text();
            console.log(`网关 ${gateway} 响应文本:`, text.substring(0, 200) + (text.length > 200 ? '...' : ''));

            try {
              // 尝试将文本解析为JSON
              const metadata = JSON.parse(text);
              console.log(`网关 ${gateway} 成功解析元数据:`, metadata);

              // 如果图片URL是IPFS链接，也转换为HTTP链接
              if (metadata.image && metadata.image.startsWith('ipfs://')) {
                metadata.image = metadata.image.replace('ipfs://', gateway);
                console.log('转换图片IPFS链接为HTTP:', metadata.image);
              }

              return metadata;
            } catch (parseError) {
              console.error(`网关 ${gateway} 解析JSON失败:`, parseError);
              // 尝试检查响应是否是HTML而不是JSON
              if (text.includes('<!DOCTYPE html>') || text.includes('<html>')) {
                console.warn(`网关 ${gateway} 返回了HTML而不是JSON，尝试下一个网关`);
              }
              // 继续尝试下一个网关
            }
          } catch (fetchError) {
            console.error(`网关 ${gateway} 获取失败:`, fetchError);
            // 继续尝试下一个网关
          }
        }

        console.warn('所有IPFS网关都失败，使用基本元数据');
        return fallbackMetadata;
      }

      // 如果URI是相对路径，添加基础URL
      let fetchUrl = uri;
      if (uri.startsWith('/')) {
        fetchUrl = `${window.location.origin}${uri}`;
        console.log('转换相对路径为绝对路径:', fetchUrl);
      }

      // 如果URI已经是完整的HTTP/HTTPS URL，直接使用
      console.log('使用URL获取元数据:', fetchUrl);

      try {
        // 设置超时
        const controller = new AbortController();
        const timeoutId = setTimeout(() => controller.abort(), 10000); // 10秒超时

        const response = await fetch(fetchUrl, { signal: controller.signal });
        clearTimeout(timeoutId);

        if (!response.ok) {
          console.error(`获取元数据失败，状态码: ${response.status}`);
          return this.createBasicMetadata(uri);
        }

        const contentType = response.headers.get('content-type');
        console.log('响应内容类型:', contentType);

        // 尝试获取响应文本
        const text = await response.text();
        console.log('响应文本:', text.substring(0, 200) + (text.length > 200 ? '...' : ''));

        try {
          // 尝试将文本解析为JSON
          const metadata = JSON.parse(text);
          console.log('成功解析元数据:', metadata);

          // 如果图片URL是IPFS链接，也转换为HTTP链接
          if (metadata.image && metadata.image.startsWith('ipfs://')) {
            // 使用第一个IPFS网关
            metadata.image = metadata.image.replace('ipfs://', ipfsGateways[0]);
            console.log('转换图片IPFS链接为HTTP:', metadata.image);
          }

          return metadata;
        } catch (parseError) {
          console.error('解析JSON失败:', parseError);
          // 返回基本元数据
          return this.createBasicMetadata(uri);
        }
      } catch (error) {
        console.error('获取元数据失败:', error);
        // 返回一个基本的元数据对象
        return this.createBasicMetadata(uri);
      }
    } catch (error) {
      console.error('处理元数据过程中发生错误:', error);
      // 返回一个基本的元数据对象
      return this.createBasicMetadata(uri);
    }
  },

  // 创建基本的元数据对象
  createBasicMetadata(uri) {
    console.log('创建基本元数据对象，URI:', uri);

    // 从URI中提取ID或名称
    let name = 'NFT';
    if (uri) {
      // 尝试从URI中提取最后一部分作为名称
      const parts = uri.split('/');
      const lastPart = parts[parts.length - 1];
      if (lastPart) {
        name = `NFT #${lastPart.substring(0, 8)}`;
      }
    }

    return {
      name: name,
      description: 'This is an NFT created on our platform',
      image: uri,
      attributes: [
        { trait_type: 'Type', value: 'Main NFT' },
        { trait_type: 'Rarity', value: 'Common' }
      ]
    };
  },

  // 获取NFT请求列表
  async getNFTRequests() {
    const response = await api.get('/nft/all-requests')
    return response.data
  },

  // 批准NFT请求
  async approveNFTRequest(requestId, data) {
    const response = await api.post(`/nft/requests/${requestId}/approve`, data)
    return response.data
  },

  // 拒绝NFT请求
  async rejectNFTRequest(requestId, data) {
    const response = await api.post(`/nft/requests/${requestId}/reject`, data)
    return response.data
  },

  // 创建NFT请求
  async createNFTRequest(data) {
    const response = await api.post('/nft/requests', data)
    return response.data
  }
};

export default nftService; 