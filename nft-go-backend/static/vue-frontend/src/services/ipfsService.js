import api from '@/utils/api';

const ipfsService = {
    /**
     * 上传数据到IPFS
     * @param {string} data - 要上传的数据
     * @param {string} filename - 文件名（可选）
     * @param {boolean} isBinary - 是否为二进制数据（Base64编码）
     * @returns {Promise} - 包含IPFS hash的响应
     */
    async uploadToIPFS(data, filename = 'data.txt', isBinary = false) {
        try {
            const response = await api.post('/ipfs/upload', {
                data: data,
                filename: filename,
                is_binary: isBinary
            });

            return response.data;
        } catch (error) {
            console.error('上传到IPFS失败:', error);
            throw new Error(error.response?.data?.error || '上传到IPFS失败');
        }
    },

    /**
     * 从IPFS获取数据 - 修复版本：正确处理二进制数据
     * @param {string} hash - IPFS hash
     * @returns {Promise} - 数据内容（ArrayBuffer用于二进制，字符串用于文本）
     */
    async getFromIPFS(hash) {
        try {
            // 尝试多个IPFS网关
            const gateways = [
                'https://dweb.link/ipfs/',
                'https://cloudflare-ipfs.com/ipfs/',
                'https://gateway.pinata.cloud/ipfs/',
                'https://ipfs.io/ipfs/'
            ];

            let lastError = null;

            for (const gateway of gateways) {
                try {
                    console.log(`尝试从网关获取: ${gateway}${hash}`);

                    const controller = new AbortController();
                    const timeoutId = setTimeout(() => controller.abort(), 15000); // 15秒超时

                    const response = await fetch(`${gateway}${hash}`, {
                        signal: controller.signal,
                        method: 'GET',
                        headers: {
                            'Accept': '*/*'
                        }
                    });

                    clearTimeout(timeoutId);

                    if (response.ok) {
                        // 检查内容类型来决定如何处理响应
                        const contentType = response.headers.get('content-type') || '';
                        console.log(`响应内容类型: ${contentType}`);

                        // 对于二进制文件，直接返回ArrayBuffer
                        if (contentType.startsWith('image/') ||
                            contentType.includes('application/octet-stream') ||
                            contentType.includes('application/pdf') ||
                            contentType.startsWith('video/') ||
                            contentType.startsWith('audio/')) {

                            const arrayBuffer = await response.arrayBuffer();
                            return {
                                data: arrayBuffer,
                                isBinary: true,
                                contentType: contentType
                            };
                        } else {
                            // 文本内容
                            const text = await response.text();
                            return {
                                data: text,
                                isBinary: false,
                                contentType: contentType
                            };
                        }
                    } else {
                        lastError = new Error(`HTTP ${response.status}: ${response.statusText}`);
                        console.warn(`网关 ${gateway} 返回错误:`, lastError.message);
                    }
                } catch (error) {
                    lastError = error;
                    if (error.name === 'AbortError') {
                        console.warn(`网关 ${gateway} 请求超时`);
                    } else {
                        console.warn(`网关 ${gateway} 访问失败:`, error.message);
                    }
                    continue;
                }
            }

            throw lastError || new Error('所有IPFS网关都无法访问');
        } catch (error) {
            console.error('从IPFS获取数据失败:', error);
            throw new Error(`从IPFS获取数据失败: ${error.message}`);
        }
    },

    /**
     * 生成IPFS网关URL
     * @param {string} hash - IPFS hash
     * @param {number} gatewayIndex - 网关索引
     * @returns {string} - 完整的IPFS URL
     */
    getIPFSUrl(hash, gatewayIndex = 0) {
        const gateways = [
            'https://dweb.link/ipfs/',
            'https://cloudflare-ipfs.com/ipfs/',
            'https://gateway.pinata.cloud/ipfs/',
            'https://ipfs.io/ipfs/'
        ];

        return `${gateways[gatewayIndex] || gateways[0]}${hash}`;
    },

    /**
     * 通过后端API获取IPFS文件信息
     * @param {string} hash - IPFS hash
     * @returns {Promise} - 文件信息和内容
     */
    async getFromIPFSByAPI(hash) {
        try {
            const response = await api.get(`/ipfs/get/${hash}`);
            return response.data;
        } catch (error) {
            console.error('通过API获取IPFS文件失败:', error);
            throw new Error(error.response?.data?.error || '获取IPFS文件失败');
        }
    },

    /**
     * 通过后端API下载IPFS文件
     * @param {string} hash - IPFS hash
     * @returns {Promise} - 文件下载URL
     */
    getDownloadUrl(hash) {
        return `/api/ipfs/download/${hash}`;
    },

    /**
     * 验证IPFS Hash格式
     * @param {string} hash - IPFS hash
     * @returns {boolean} - 是否为有效格式
     */
    validateHash(hash) {
        if (!hash || typeof hash !== 'string') return false;

        // 基本格式验证
        return hash.length >= 40 && (hash.startsWith('Qm') || hash.startsWith('bafy'));
    },

    /**
     * 优化的从IPFS获取内容方法（优先使用后端API）
     * @param {string} hash - IPFS hash
     * @returns {Promise} - 文件内容和信息
     */
    async getFromIPFSOptimized(hash) {
        // 先验证hash格式
        if (!this.validateHash(hash)) {
            throw new Error('Hash格式不正确，请输入有效的IPFS Hash');
        }

        try {
            // 优先使用后端API
            console.log('尝试通过后端API获取IPFS内容...');
            const apiResult = await this.getFromIPFSByAPI(hash);

            // 如果后端返回了内容，直接使用
            if (apiResult.content) {
                return {
                    content: apiResult.content,
                    hash: apiResult.hash,
                    size: apiResult.size,
                    contentType: apiResult.content_type,
                    isText: apiResult.is_text,
                    isImage: apiResult.is_image,
                    url: apiResult.url,
                    downloadUrl: apiResult.download_url,
                    source: 'api'
                };
            } else {
                // 如果后端没有返回内容（如大文件），但有信息，使用网关获取内容
                console.log('后端API返回文件信息，通过网关获取内容...');
                const gatewayResult = await this.getFromIPFS(hash);

                return {
                    content: gatewayResult.data,
                    hash: apiResult.hash,
                    size: apiResult.size,
                    contentType: apiResult.content_type || gatewayResult.contentType,
                    isText: apiResult.is_text,
                    isImage: apiResult.is_image,
                    isBinary: gatewayResult.isBinary,
                    url: apiResult.url,
                    downloadUrl: apiResult.download_url,
                    source: 'gateway'
                };
            }
        } catch (apiError) {
            console.warn('后端API获取失败，尝试直接使用网关:', apiError.message);

            // 如果后端API失败，回退到直接使用网关
            try {
                const gatewayResult = await this.getFromIPFS(hash);

                // 简单的内容类型检测
                const isText = gatewayResult.isBinary ? false : this.isTextContent(gatewayResult.data);
                const isImage = this.isImageContentType(gatewayResult.contentType);
                const contentType = gatewayResult.contentType || this.detectContentType(gatewayResult.data);

                // 计算文件大小
                const size = gatewayResult.isBinary ? 
                    gatewayResult.data.byteLength : 
                    new Blob([gatewayResult.data]).size;

                return {
                    content: gatewayResult.data,
                    hash: hash,
                    size: size,
                    contentType: contentType,
                    isText: isText,
                    isImage: isImage,
                    isBinary: gatewayResult.isBinary,
                    url: this.getIPFSUrl(hash),
                    downloadUrl: this.getDownloadUrl(hash),
                    source: 'gateway-fallback'
                };
            } catch (gatewayError) {
                throw new Error(`获取IPFS内容失败: ${gatewayError.message}`);
            }
        }
    },

    /**
     * 简单的文本内容检测
     */
    isTextContent(content) {
        if (typeof content !== 'string') return false;

        const binaryCount = (content.match(/[\x00-\x08\x0E-\x1F\x7F-\xFF]/g) || []).length;
        const ratio = binaryCount / content.length;
        return ratio < 0.1;
    },

    /**
     * 基于Content-Type的图片检测
     */
    isImageContentType(contentType) {
        if (!contentType) return false;
        return contentType.startsWith('image/');
    },

    /**
     * 简单的图片内容检测
     */
    isImageContent(content) {
        if (typeof content !== 'string') return false;

        const imageSignatures = [
            '\xFF\xD8\xFF', // JPEG
            '\x89PNG', // PNG
            'GIF8', // GIF
            'RIFF', // WebP
            '<svg' // SVG
        ];

        return imageSignatures.some(sig => content.startsWith(sig));
    },

    /**
     * 简单的内容类型检测
     */
    detectContentType(content) {
        if (typeof content !== 'string') return 'application/octet-stream';

        if (content.startsWith('\xFF\xD8\xFF')) return 'image/jpeg';
        if (content.startsWith('\x89PNG')) return 'image/png';
        if (content.startsWith('GIF8')) return 'image/gif';
        if (content.startsWith('<svg')) return 'image/svg+xml';
        if (content.startsWith('RIFF')) return 'image/webp';
        if (content.startsWith('{') || content.startsWith('[')) return 'application/json';
        if (content.includes('<!DOCTYPE html') || content.includes('<html')) return 'text/html';

        return this.isTextContent(content) ? 'text/plain' : 'application/octet-stream';
    }
};

export default ipfsService; 