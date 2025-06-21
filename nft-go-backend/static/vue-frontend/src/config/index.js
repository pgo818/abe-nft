// 环境变量
const env = process.env.NODE_ENV || 'development';

// 基础配置
const baseConfig = {
  // API基础URL
  apiBaseUrl: '/api',
  
  // 区块链配置
  blockchain: {
    // 以太坊网络ID
    networkId: 1,
    // 智能合约地址
    contracts: {
      nft: '0x0000000000000000000000000000000000000000',
      childNft: '0x0000000000000000000000000000000000000000'
    }
  },
  
  // IPFS配置
  ipfs: {
    gateway: 'https://ipfs.io/ipfs/'
  }
};

// 环境特定配置
const envConfig = {
  development: {
    apiBaseUrl: '/api',
    blockchain: {
      networkId: 5, // Goerli测试网
      contracts: {
        nft: process.env.VUE_APP_NFT_CONTRACT_ADDRESS,
        childNft: process.env.VUE_APP_CHILD_NFT_CONTRACT_ADDRESS
      }
    }
  },
  production: {
    apiBaseUrl: '/api',
    blockchain: {
      networkId: 1, // 以太坊主网
      contracts: {
        nft: process.env.VUE_APP_NFT_CONTRACT_ADDRESS,
        childNft: process.env.VUE_APP_CHILD_NFT_CONTRACT_ADDRESS
      }
    }
  }
};

// 合并配置
const config = {
  ...baseConfig,
  ...envConfig[env]
};

export default config; 