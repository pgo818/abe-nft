/**
 * 带签名的API请求工具函数
 */

import axios from 'axios';

// 创建axios实例
const api = axios.create({
  baseURL: '/api', // 与后端API的基础URL保持一致
  timeout: 10000,  // 请求超时时间
  headers: {
    'Content-Type': 'application/json'
  }
});

// 请求拦截器
api.interceptors.request.use(
  config => {
    // 从localStorage获取token
    const token = localStorage.getItem('token');
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// 响应拦截器
api.interceptors.response.use(
  response => {
    return response;
  },
  error => {
    if (error.response) {
      // 处理错误响应
      switch (error.response.status) {
        case 401:
          // 未授权，清除token并重定向到登录页
          localStorage.removeItem('token');
          window.location.href = '/';
          break;
        case 403:
          // 权限不足
          console.error('权限不足');
          break;
        case 500:
          // 服务器错误
          console.error('服务器错误');
          break;
        default:
          console.error(`请求错误: ${error.response.status}`);
      }
    } else if (error.request) {
      // 请求已发出但未收到响应
      console.error('网络错误，无法连接到服务器');
    } else {
      // 请求配置出错
      console.error('请求配置错误:', error.message);
    }
    return Promise.reject(error);
  }
);

export default api;

/**
 * 发送带签名的GET请求
 * @param {string} url - 请求URL
 * @param {Object} options - 请求选项
 * @param {string} options.address - 钱包地址
 * @param {string} options.signature - 签名
 * @param {string} options.message - 签名的消息
 * @returns {Promise} - Fetch Promise
 */
export const fetchWithSignature = (url, options) => {
    const { method = 'GET', address, signature, message } = options

    // 构建请求头
    const headers = {
        'Content-Type': 'application/json',
        'X-Ethereum-Address': address,
        'X-Ethereum-Signature': signature,
        'X-Ethereum-Message': message
    }

    // 发送请求
    return fetch(url, {
        method,
        headers
    })
}

/**
 * 构建带签名的请求体
 * @param {string} action - 操作类型
 * @param {Object} data - 请求数据
 * @param {string} address - 钱包地址
 * @param {string} signature - 签名
 * @param {string} message - 签名的消息
 * @returns {Object} - 请求体对象
 */
export const buildSignedRequest = (action, data, address, signature, message) => {
    return {
        address,
        signature,
        message,
        ...data
    }
}

/**
 * 创建签名消息
 * @param {string} action - 操作类型
 * @param {Object} data - 消息数据
 * @returns {string} - JSON字符串
 */
export const createSignMessage = (action, data) => {
    return JSON.stringify({
        action,
        ...data,
        timestamp: Date.now()
    })
}

/**
 * 处理API响应
 * @param {Response} response - Fetch响应对象
 * @returns {Promise} - 处理后的Promise
 */
export const handleApiResponse = async (response) => {
    const data = await response.json()

    if (!response.ok) {
        throw new Error(data.error || '请求失败')
    }

    return data
}

/**
 * 格式化钱包地址（简短显示）
 * @param {string} address - 钱包地址
 * @returns {string} - 格式化后的地址
 */
export const formatAddress = (address) => {
    if (!address) return ''
    return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`
} 