// 全局变量
let currentAccount = null;
let isConnected = false;
const API_BASE_URL = window.location.origin + '/api';

// 保存钱包连接状态到localStorage
function saveWalletState(account) {
    localStorage.setItem('walletConnected', 'true');
    localStorage.setItem('walletAccount', account);
}

// 清除钱包连接状态
function clearWalletState() {
    localStorage.removeItem('walletConnected');
    localStorage.removeItem('walletAccount');
}

// 从localStorage恢复钱包状态
function restoreWalletState() {
    const wasConnected = localStorage.getItem('walletConnected') === 'true';
    const savedAccount = localStorage.getItem('walletAccount');

    if (wasConnected && savedAccount) {
        return savedAccount;
    }
    return null;
}

// 初始化钱包连接
function initializeWalletConnection() {
    console.log('初始化钱包连接...');
    // 这个函数用于初始化钱包连接状态
    // 主要逻辑在checkIfWalletConnected中处理
}

// DOM加载完成后初始化
document.addEventListener('DOMContentLoaded', () => {
    // 初始化钱包连接
    initializeWalletConnection();

    // 隐藏所有不相关的界面元素
    hideIrrelevantSections();

    // 检查MetaMask是否已安装
    if (!checkIfMetaMaskInstalled()) {
        return;
    }

    // 绑定事件处理函数
    setupEventListeners();

    // 检查是否已连接钱包
    checkIfWalletConnected();
    loadAllNFTs(); // 加载所有NFT

    // 设置MetaMask事件监听器
    if (window.ethereum) {
        window.ethereum.on('accountsChanged', handleAccountsChanged);
        window.ethereum.on('chainChanged', () => {
            // 网络变化时重新加载页面
            window.location.reload();
        });
    }

    // 在页面加载时检查是否有保存的钱包和DID
    const walletAddress = localStorage.getItem('userWallet');
    const did = localStorage.getItem('userDID');

    if (walletAddress && did) {
        updateWalletInfo(walletAddress, did);
    }

    // 添加连接钱包按钮
    const header = document.querySelector('header');
    if (header) {
        const walletButton = document.createElement('button');
        walletButton.id = 'connect-wallet';
        walletButton.className = 'btn btn-primary';
        walletButton.innerHTML = '<i class="fas fa-wallet"></i> 连接钱包创建DID';
        walletButton.onclick = createDIDFromWallet;
        header.appendChild(walletButton);
    }
});

// 检查MetaMask是否已安装
function checkIfMetaMaskInstalled() {
    if (typeof window.ethereum === 'undefined') {
        showError('请安装MetaMask钱包以使用本应用');
        document.getElementById('connect-wallet').disabled = true;
        document.getElementById('connect-wallet-btn').disabled = true;
        return false;
    }
    return true;
}

// 绑定事件处理函数
function setupEventListeners() {
    // 连接钱包按钮
    const connectWalletBtn = document.getElementById('connect-wallet');
    const connectWalletPromptBtn = document.getElementById('connect-wallet-btn');

    if (connectWalletBtn) {
        connectWalletBtn.addEventListener('click', connectWallet);
    }
    if (connectWalletPromptBtn) {
        connectWalletPromptBtn.addEventListener('click', connectWallet);
    }

    // 添加断开连接按钮监听器
    const disconnectBtn = document.getElementById('disconnect-wallet');
    if (disconnectBtn) {
        disconnectBtn.addEventListener('click', disconnectWallet);
    }

    // 导航事件
    const navAllNfts = document.getElementById('nav-all-nfts');
    const navMint = document.getElementById('nav-mint');
    const navMetadata = document.getElementById('nav-metadata');
    const navMyNfts = document.getElementById('nav-my-nfts');
    const navRequests = document.getElementById('nav-requests');

    if (navAllNfts) {
        navAllNfts.addEventListener('click', () => {
            showSection('all-nfts-section');
            loadAllNFTs();
        });
    }
    if (navMint) {
        navMint.addEventListener('click', () => {
            showSection('mint-nft-section');
        });
    }
    if (navMetadata) {
        navMetadata.addEventListener('click', () => {
            showSection('metadata-section');
            loadMetadataList();
            updateMetadataPreview();
        });
    }
    if (navMyNfts) {
        navMyNfts.addEventListener('click', () => {
            showSection('my-nfts-section');
            loadMyNFTs();
        });
    }
    if (navRequests) {
        navRequests.addEventListener('click', () => {
            showSection('requests-section');
            loadPendingRequests();
        });
    }

    // 表单提交事件
    const mintForm = document.getElementById('mint-form');
    const metadataForm = document.getElementById('metadata-form');
    const requestForm = document.getElementById('request-form');
    const requestChildForm = document.getElementById('request-child-form');
    const createChildForm = document.getElementById('create-child-form');
    const updateMetadataForm = document.getElementById('update-metadata-form');

    if (mintForm) {
        mintForm.addEventListener('submit', handleMintFormSubmit);
    }
    if (metadataForm) {
        metadataForm.addEventListener('submit', handleMetadataFormSubmit);
        // 添加输入字段的事件监听器，用于实时预览
        const metadataInputs = ['metadata-name', 'metadata-description', 'metadata-external-url',
            'metadata-image', 'metadata-policy', 'metadata-ciphertext'];
        metadataInputs.forEach(id => {
            const element = document.getElementById(id);
            if (element) {
                element.addEventListener('input', updateMetadataPreview);
            }
        });
    }
    if (requestForm) {
        requestForm.addEventListener('submit', handleRequestFormSubmit);
    }
    if (requestChildForm) {
        requestChildForm.addEventListener('submit', handleRequestChildFormSubmit);
    }
    if (createChildForm) {
        createChildForm.addEventListener('submit', handleCreateChildFormSubmit);
    }
    if (updateMetadataForm) {
        updateMetadataForm.addEventListener('submit', handleUpdateMetadataFormSubmit);
    }
}

// 检查钱包是否已连接
async function checkIfWalletConnected() {
    if (!checkIfMetaMaskInstalled()) {
        return;
    }

    try {
        // 首先尝试从localStorage恢复状态
        const savedAccount = restoreWalletState();

        // 检查MetaMask当前连接的账户
        const accounts = await window.ethereum.request({
            method: 'eth_accounts'
        });

        if (accounts.length > 0) {
            // MetaMask有连接的账户
            const currentMetaMaskAccount = accounts[0];

            if (savedAccount && savedAccount === currentMetaMaskAccount) {
                // localStorage中保存的账户与MetaMask当前账户一致
                currentAccount = currentMetaMaskAccount;
                isConnected = true;
                updateUIForConnectedWallet();
                console.log('自动连接钱包成功:', currentAccount);
            } else if (savedAccount) {
                // 账户不一致，清除旧状态，使用新账户
                currentAccount = currentMetaMaskAccount;
                isConnected = true;
                saveWalletState(currentAccount);
                updateUIForConnectedWallet();
                console.log('检测到账户变更，已更新:', currentAccount);
            } else {
                // 没有保存的状态，但MetaMask有连接，自动连接
                currentAccount = currentMetaMaskAccount;
                isConnected = true;
                saveWalletState(currentAccount);
                updateUIForConnectedWallet();
                console.log('自动连接检测到的钱包:', currentAccount);
            }
        } else if (savedAccount) {
            // MetaMask没有连接账户，但localStorage有保存，清除过期状态
            clearWalletState();
            console.log('清除过期的钱包连接状态');
        }
    } catch (error) {
        console.error('检查钱包连接失败:', error);
        // 清除可能损坏的状态
        clearWalletState();
    }
}

// 连接钱包
async function connectWallet() {
    if (!window.ethereum) {
        alert("请安装MetaMask钱包！");
        return null;
    }

    try {
        const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
        const walletAddress = accounts[0];
        console.log("钱包连接成功:", walletAddress);
        return walletAddress;
    } catch (error) {
        console.error("连接钱包失败:", error);
        alert("连接钱包失败: " + error.message);
        return null;
    }
}

// 从钱包创建DID
async function createDIDFromWallet() {
    const walletAddress = await connectWallet();
    if (!walletAddress) return;

    try {
        const response = await fetch(`/api/did/wallet/${walletAddress}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        const data = await response.json();
        if (response.ok) {
            if (data.exists) {
                alert(`已找到现有DID: ${data.did}`);
            } else {
                alert(`创建DID成功: ${data.did}`);
            }
            // 保存DID到本地存储
            localStorage.setItem('userDID', data.did);
            localStorage.setItem('userWallet', walletAddress);

            // 更新UI显示
            updateWalletInfo(walletAddress, data.did);
        } else {
            alert("创建DID失败: " + (data.error || "未知错误"));
        }
    } catch (error) {
        console.error("创建DID错误:", error);
        alert("创建DID时发生错误: " + error.message);
    }
}

// 更新钱包信息显示
function updateWalletInfo(walletAddress, did) {
    const walletInfoDiv = document.getElementById('wallet-info');
    if (!walletInfoDiv) {
        // 创建钱包信息显示区域
        const container = document.querySelector('.container');
        const newWalletInfo = document.createElement('div');
        newWalletInfo.id = 'wallet-info';
        newWalletInfo.className = 'wallet-info';
        newWalletInfo.innerHTML = `
            <div class="wallet-address">
                <strong>钱包地址:</strong> ${shortenAddress(walletAddress)}
            </div>
            <div class="did-info">
                <strong>DID:</strong> ${did}
            </div>
        `;
        container.insertBefore(newWalletInfo, container.firstChild);
    } else {
        // 更新现有钱包信息
        walletInfoDiv.innerHTML = `
            <div class="wallet-address">
                <strong>钱包地址:</strong> ${shortenAddress(walletAddress)}
            </div>
            <div class="did-info">
                <strong>DID:</strong> ${did}
            </div>
        `;
    }
}

// 缩短地址显示
function shortenAddress(address) {
    return address ? `${address.substring(0, 6)}...${address.substring(address.length - 4)}` : '';
}

// 处理账户变更
function handleAccountsChanged(accounts) {
    if (accounts.length === 0) {
        // 用户断开了所有账户
        isConnected = false;
        currentAccount = null;
        clearWalletState();
        updateUIForDisconnectedWallet();
        console.log('钱包已断开连接');
    } else if (accounts[0] !== currentAccount) {
        // 用户切换了账户
        currentAccount = accounts[0];
        saveWalletState(currentAccount);
        updateUIForConnectedWallet();
        console.log('钱包账户已切换:', currentAccount);
    }
}

// 更新已连接钱包的UI
function updateUIForConnectedWallet() {
    // 显示钱包地址
    const displayAddress = `${currentAccount.substring(0, 6)}...${currentAccount.substring(38)}`;
    const walletAddressElement = document.getElementById('wallet-address');
    if (walletAddressElement) {
        walletAddressElement.textContent = displayAddress;
        walletAddressElement.title = currentAccount;
    } else {
        console.error('未找到wallet-address元素');
    }

    // 更新UI元素
    const connectWalletBtn = document.getElementById('connect-wallet');
    const walletInfoDiv = document.getElementById('wallet-info');
    const connectWalletPrompt = document.getElementById('connect-wallet-prompt');

    if (connectWalletBtn) connectWalletBtn.classList.add('d-none');
    if (walletInfoDiv) walletInfoDiv.classList.remove('d-none');
    if (connectWalletPrompt) connectWalletPrompt.classList.add('d-none');

    // 更新钱包信息显示
    updateWalletInfo(currentAccount, localStorage.getItem('userDID') || '');

    // 默认显示主界面而不是铸造NFT页面
    showSection('all-nfts-section');
    updateNavigation(''); // 不高亮任何导航

    // 加载所有NFT
    loadAllNFTs();
}

// 更新断开连接钱包的UI
function updateUIForDisconnectedWallet() {
    // 更新UI元素
    document.getElementById('connect-wallet').classList.remove('d-none');
    document.getElementById('wallet-info').classList.add('d-none');
    document.getElementById('connect-wallet-prompt').classList.remove('d-none');

    // 隐藏所有功能区域
    hideAllSections();
}

// 显示指定区域，隐藏其他区域
function showSection(sectionId) {
    if (!isConnected && sectionId !== 'all-nfts-section') {
        return; // 如果未连接钱包且不是主页面，不执行
    }

    // 隐藏所有区域
    hideAllSections();

    // 显示指定区域
    document.getElementById(sectionId).classList.remove('d-none');

    // 更新导航状态
    updateNavigation(sectionId);
}

// 隐藏所有功能区域
function hideAllSections() {
    document.querySelectorAll('.section').forEach(section => {
        section.classList.add('d-none');
    });
}

// 显示/隐藏加载指示器
function showLoading(show) {
    const loadingOverlay = document.getElementById('loading-overlay');
    if (show) {
        loadingOverlay.classList.remove('d-none');
    } else {
        loadingOverlay.classList.add('d-none');
    }
}

// 处理铸造NFT表单提交
async function handleMintFormSubmit(event) {
    event.preventDefault();
    if (!isConnected) return;

    const uri = document.getElementById('nft-uri').value.trim();
    if (!uri) {
        showError('请输入有效的URI');
        return;
    }

    showLoading(true);
    try {
        // 创建要签名的消息
        const message = JSON.stringify({
            action: 'mint_nft',
            uri: uri,
            timestamp: Date.now()
        });

        // 获取签名
        const signature = await signMessage(message);

        // 发送铸造请求
        const response = await fetch(`${API_BASE_URL}/nft/mint`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                address: currentAccount,
                signature: signature,
                message: message,
                uri: uri
            })
        });

        const result = await response.json();

        if (response.ok) {
            // 显示成功消息
            showSuccess('NFT铸造交易已提交', result.transactionHash);
            // 重置表单
            document.getElementById('mint-form').reset();
            // 重新加载NFT列表
            loadAllNFTs();
            if (document.getElementById('my-nfts-section').classList.contains('d-none') === false) {
                loadMyNFTs();
            }
        } else {
            showError(result.error || '铸造NFT失败');
        }
    } catch (error) {
        console.error('铸造NFT出错:', error);
        showError('铸造NFT失败: ' + error.message);
    } finally {
        showLoading(false);
    }
}

// 处理申请子NFT表单提交
async function handleRequestFormSubmit(event) {
    event.preventDefault();
    if (!isConnected) return;

    const parentTokenId = document.getElementById('parent-token-id').value.trim();
    const uri = document.getElementById('request-uri').value.trim();

    if (!parentTokenId || !uri) {
        showError('请填写所有必填字段');
        return;
    }

    showLoading(true);
    try {
        // 创建要签名的消息
        const message = JSON.stringify({
            action: 'request_child_nft',
            parentTokenId: parentTokenId,
            uri: uri,
            timestamp: Date.now()
        });

        // 获取签名
        const signature = await signMessage(message);

        // 发送申请请求
        const response = await fetch(`${API_BASE_URL}/nft/request-child`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                address: currentAccount,
                signature: signature,
                message: message,
                parentTokenId: parentTokenId,
                applicantAddress: currentAccount,
                uri: uri
            })
        });

        const result = await response.json();

        if (response.ok) {
            // 显示成功消息
            showSuccess('子NFT申请已提交，等待审批');
            // 重置表单
            document.getElementById('request-form').reset();
        } else {
            showError(result.error || '提交申请失败');
        }
    } catch (error) {
        console.error('提交申请出错:', error);
        showError('提交申请失败: ' + error.message);
    } finally {
        showLoading(false);
    }
}

// 加载我的NFT列表
async function loadMyNFTs() {
    if (!isConnected) return;

    const nftList = document.getElementById('my-nfts-list');
    nftList.innerHTML = '<div class="col-12 text-center"><p>加载中...</p></div>';

    try {
        // 创建要签名的消息
        const message = JSON.stringify({
            action: 'get_my_nfts',
            timestamp: Date.now()
        });

        // 获取签名
        const signature = await signMessage(message);

        // 发送获取我的NFT请求
        const response = await fetch(`${API_BASE_URL}/nft/my-nfts`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'X-Ethereum-Address': currentAccount,
                'X-Ethereum-Signature': signature,
                'X-Ethereum-Message': message
            }
        });

        const result = await response.json();

        if (response.ok && result.nfts) {
            if (result.nfts.length === 0) {
                nftList.innerHTML = '<div class="col-12 text-center"><p>您还没有NFT</p></div>';
                return;
            }

            nftList.innerHTML = result.nfts.map(nft => `
                <div class="col-md-4">
                    <div class="card nft-card">
                        <img src="${nft.metadata?.image || nft.uri}" class="card-img-top" alt="${nft.metadata?.name || 'NFT'}">
                        <div class="card-body">
                            <h5 class="card-title">
                                ${nft.metadata?.name || 'NFT #' + nft.tokenId}
                                ${nft.isChildNft ? '<span class="badge bg-secondary ms-2">子NFT</span>' : '<span class="badge bg-primary ms-2">主NFT</span>'}
                            </h5>
                            <p class="card-text text-truncate">${nft.metadata?.description || 'No description'}</p>
                            <p class="card-text"><small class="text-muted">Token ID: ${nft.tokenId}</small></p>
                            ${nft.isChildNft && nft.parentTokenId ? `<p class="card-text"><small class="text-muted">父NFT ID: ${nft.parentTokenId}</small></p>` : ''}
                            <div class="btn-group w-100" role="group">
                                ${!nft.isChildNft ? `<button class="btn btn-primary btn-sm" onclick="createChildNFT('${nft.tokenId}')">创建子NFT</button>` : ''}
                                <button class="btn btn-secondary btn-sm" onclick="updateMetadata('${nft.tokenId}', '${nft.contractType || (nft.isChildNft ? 'child' : 'main')}')">更新元数据</button>
                            </div>
                        </div>
                    </div>
                </div>
            `).join('');
        } else {
            nftList.innerHTML = '<div class="col-12 text-center"><p class="text-danger">加载NFT失败</p></div>';
        }

    } catch (error) {
        console.error('加载NFT列表出错:', error);
        nftList.innerHTML = '<div class="col-12 text-center"><p class="text-danger">加载NFT失败</p></div>';
    }
}

// 加载所有NFT列表（主页面）
async function loadAllNFTs() {
    try {
        const response = await fetch(`${API_BASE_URL}/nfts`);
        const result = await response.json();

        if (response.ok && result.nfts) {
            const nftContainer = document.getElementById('all-nfts-container');
            if (!nftContainer) return;

            if (result.nfts.length === 0) {
                nftContainer.innerHTML = '<div class="col-12 text-center"><p>暂无NFT</p></div>';
                return;
            }

            nftContainer.innerHTML = result.nfts.map(nft => `
                <div class="col-md-4">
                    <div class="card nft-card">
                        <img src="${nft.metadata?.image || nft.uri}" class="card-img-top" alt="${nft.metadata?.name || 'NFT'}">
                        <div class="card-body">
                            <h5 class="card-title">${nft.metadata?.name || 'NFT #' + nft.tokenId}</h5>
                            <p class="card-text text-truncate">${nft.metadata?.description || 'No description'}</p>
                            <p class="card-text"><small class="text-muted">Token ID: ${nft.tokenId}</small></p>
                            <p class="card-text"><small class="text-muted">Owner: ${nft.owner.substring(0, 6)}...${nft.owner.substring(38)}</small></p>
                            <div class="d-flex gap-2 mt-3">
                                ${isConnected ?
                    `<button class="btn btn-outline-primary btn-sm" onclick="requestChildNFT('${nft.tokenId}')">申请子NFT</button>` :
                    ''
                }
                            </div>
                        </div>
                    </div>
                </div>
            `).join('');
        }
    } catch (error) {
        console.error('加载所有NFT失败:', error);
    }
}

// 修改加载待处理申请的函数
async function loadPendingRequests() {
    console.log('loadPendingRequests 被调用');
    console.log('isConnected:', isConnected);
    console.log('currentAccount:', currentAccount);

    if (!isConnected) {
        console.log('用户未连接钱包，无法加载请求');
        return;
    }

    const pendingList = document.getElementById('pending-requests-list');
    if (!pendingList) {
        console.error('找不到pending-requests-list元素');
        return;
    }

    pendingList.innerHTML = '<tr><td colspan="6" class="text-center">加载中...</td></tr>';

    try {
        const message = JSON.stringify({
            action: 'get_pending_requests',
            timestamp: Date.now()
        });

        const signature = await signMessage(message);

        console.log('发送请求到API...');
        const response = await fetch(`${API_BASE_URL}/nft/all-requests`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'X-Ethereum-Address': currentAccount,
                'X-Ethereum-Signature': signature,
                'X-Ethereum-Message': message
            }
        });

        console.log('API响应状态:', response.status);
        const result = await response.json();

        // 添加调试信息
        console.log("API返回的完整响应:", result);
        console.log("响应是否OK:", response.ok);
        console.log("请求数量:", result.requests ? result.requests.length : 0);

        if (response.ok && result.requests) {
            console.log("开始处理请求数据...");

            if (result.requests.length === 0) {
                console.log("没有请求数据");
                pendingList.innerHTML = '<tr><td colspan="6" class="text-center">没有待处理的申请</td></tr>';
                return;
            }

            // 添加调试信息
            console.log("第一条请求数据:", result.requests[0]);
            console.log("所有请求数据:", result.requests);

            try {
                const htmlContent = result.requests.map((req, index) => {
                    // 调试每个请求对象
                    console.log(`处理第${index + 1}个请求:`, req);
                    console.log("请求ID:", req.ID, "或", req.id);

                    // 使用正确的ID字段（兼容大小写）
                    const requestId = req.ID || req.id || req.RequestId || req.requestId;
                    console.log("最终使用的请求ID:", requestId);

                    if (!requestId) {
                        console.error("请求ID为空:", req);
                        return '';
                    }

                    // 状态显示
                    let statusBadge = '';
                    switch (req.status) {
                        case 'pending':
                            statusBadge = '<span class="badge bg-warning">待处理</span>';
                            break;
                        case 'approved':
                            statusBadge = '<span class="badge bg-success">已完成</span>';
                            break;
                        case 'rejected':
                            statusBadge = '<span class="badge bg-danger">已拒绝</span>';
                            break;
                        default:
                            statusBadge = '<span class="badge bg-secondary">未知</span>';
                    }

                    // 操作按钮 - 只有在可操作且状态为pending时才显示
                    let actionButtons = '';
                    if (req.canOperate && req.status === 'pending') {
                        actionButtons = `
                            <button class="btn btn-success btn-sm me-1" onclick="processRequest(${requestId}, 'approve')">批准</button>
                            <button class="btn btn-danger btn-sm" onclick="processRequest(${requestId}, 'reject')">拒绝</button>
                        `;
                    } else if (req.status === 'approved') {
                        actionButtons = '<span class="text-success">已完成</span>';
                    } else if (req.status === 'rejected') {
                        actionButtons = '<span class="text-danger">已拒绝</span>';
                    } else {
                        actionButtons = '<span class="text-muted">无操作权限</span>';
                    }

                    const rowHtml = `
                        <tr class="${req.status === 'approved' ? 'table-success' : req.status === 'rejected' ? 'table-danger' : ''}">
                            <td>${requestId}</td>
                            <td>${req.parentTokenId || 'N/A'}</td>
                            <td class="wallet-address-short" title="${req.applicantAddress || 'N/A'}">
                                ${req.applicantAddress ? req.applicantAddress.substring(0, 10) + '...' + req.applicantAddress.substring(req.applicantAddress.length - 8) : 'N/A'}
                            </td>
                            <td class="text-truncate" style="max-width: 200px;" title="${req.uri || 'N/A'}">
                                ${req.uri || 'N/A'}
                            </td>
                            <td>${statusBadge}</td>
                            <td>${actionButtons}</td>
                        </tr>
                    `;
                    console.log(`第${index + 1}个请求的HTML:`, rowHtml);
                    return rowHtml;
                }).join('');

                console.log("生成的完整HTML内容:", htmlContent);

                if (htmlContent.trim() === '') {
                    console.error("生成的HTML内容为空");
                    pendingList.innerHTML = '<tr><td colspan="6" class="text-center text-danger">数据处理失败</td></tr>';
                } else {
                    pendingList.innerHTML = htmlContent;
                    console.log("HTML内容已设置到表格中");
                }

            } catch (htmlError) {
                console.error("生成HTML时出错:", htmlError);
                pendingList.innerHTML = '<tr><td colspan="6" class="text-center text-danger">数据处理失败</td></tr>';
            }

        } else {
            console.error("API返回错误或没有请求数据");
            console.error("API返回错误:", result);
            pendingList.innerHTML = '<tr><td colspan="6" class="text-center text-danger">加载失败: ' + (result.error || '未知错误') + '</td></tr>';
        }
    } catch (error) {
        console.error('加载待处理申请出错:', error);
        pendingList.innerHTML = '<tr><td colspan="6" class="text-center text-danger">加载失败: ' + error.message + '</td></tr>';
    }
}

// 简化的测试函数 - 直接显示测试数据
function testDisplayRequests() {
    console.log('测试显示请求数据...');
    const pendingList = document.getElementById('pending-requests-list');
    if (!pendingList) {
        console.error('找不到pending-requests-list元素');
        alert('找不到请求列表元素！');
        return;
    }

    // 创建测试数据
    const testData = [
        {
            ID: 8,
            id: 8,
            parentTokenId: "7",
            applicantAddress: "0x651e0fd49c7dbb5cca8b5be0319d92773443b711",
            uri: "ipfs://test123",
            status: "pending",
            canOperate: true
        },
        {
            ID: 9,
            id: 9,
            parentTokenId: "6",
            applicantAddress: "0xaf97631f96007bbde9c7803b3bea096f4a5a5561",
            uri: "ipfs://test456",
            status: "pending",
            canOperate: true
        }
    ];

    console.log('使用测试数据:', testData);

    try {
        const htmlContent = testData.map((req, index) => {
            const requestId = req.ID || req.id;

            return `
                <tr>
                    <td>${requestId}</td>
                    <td>${req.parentTokenId}</td>
                    <td class="wallet-address-short" title="${req.applicantAddress}">
                        ${req.applicantAddress.substring(0, 10)}...${req.applicantAddress.substring(req.applicantAddress.length - 8)}
                    </td>
                    <td class="text-truncate" style="max-width: 200px;" title="${req.uri}">
                        ${req.uri}
                    </td>
                    <td><span class="badge bg-warning">待处理</span></td>
                    <td>
                        <button class="btn btn-success btn-sm me-1" onclick="alert('批准请求 ${requestId}')">批准</button>
                        <button class="btn btn-danger btn-sm" onclick="alert('拒绝请求 ${requestId}')">拒绝</button>
                    </td>
                </tr>
            `;
        }).join('');

        console.log('生成的测试HTML:', htmlContent);
        pendingList.innerHTML = htmlContent;
        console.log('测试数据已显示在表格中');
        alert('测试数据已显示！检查表格是否有内容。');

    } catch (error) {
        console.error('测试显示出错:', error);
        alert('测试显示出错: ' + error.message);
    }
}

// 处理申请
async function processRequest(requestId, action) {
    if (!isConnected) return;

    console.log(`处理申请 ID=${requestId}, 操作=${action}`);

    showLoading(true);
    try {
        // 创建要签名的消息
        const message = JSON.stringify({
            action: 'process_request',
            requestId: requestId,
            decision: action,
            timestamp: Date.now()
        });

        // 获取签名
        const signature = await signMessage(message);

        // 发送处理申请的请求
        const response = await fetch(`${API_BASE_URL}/nft/process-request`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                address: currentAccount,
                signature: signature,
                message: message,
                requestId: requestId,
                action: action
            })
        });

        const result = await response.json();

        if (response.ok) {
            // 显示成功消息
            const actionText = action === 'approve' ? '批准' : '拒绝';
            showSuccess(`申请已${actionText}`, result.transactionHash);
            // 重新加载待处理申请
            loadPendingRequests();
        } else {
            showError(result.error || '处理申请失败');
        }
    } catch (error) {
        console.error('处理申请出错:', error);
        showError('处理申请失败: ' + error.message);
    } finally {
        showLoading(false);
    }
}

// 使用MetaMask签名消息
async function signMessage(message) {
    try {
        return await window.ethereum.request({
            method: 'personal_sign',
            params: [message, currentAccount]
        });
    } catch (error) {
        console.error('签名消息失败:', error);
        throw new Error('签名消息失败: ' + error.message);
    }
}

// 显示成功消息
function showSuccess(message, txHash = null) {
    const modal = new bootstrap.Modal(document.getElementById('result-modal'));
    document.getElementById('result-title').textContent = '操作成功';
    document.getElementById('result-message').textContent = message;

    const txHashContainer = document.getElementById('tx-hash-container');
    if (txHash) {
        document.getElementById('tx-hash').textContent = txHash;
        txHashContainer.classList.remove('d-none');
    } else {
        txHashContainer.classList.add('d-none');
    }

    modal.show();
}

// 显示错误消息
function showError(message) {
    const modal = new bootstrap.Modal(document.getElementById('result-modal'));
    document.getElementById('result-title').textContent = '操作失败';
    document.getElementById('result-message').textContent = message;
    document.getElementById('tx-hash-container').classList.add('d-none');
    modal.show();
}

// 新增导航状态更新函数
function updateNavigation(sectionId) {
    // 对于主导航，只处理平台级别的切换
    if (sectionId === 'home' || sectionId === 'nft' || sectionId === 'abe') {
        // 移除主导航所有活跃状态
        document.querySelectorAll('nav .navbar-nav .nav-link').forEach(link => {
            link.classList.remove('active', 'text-white');
            link.classList.add('text-light');
        });

        // 设置主页导航为活跃（如果是首页）
        if (sectionId === 'home') {
            const homeNav = document.getElementById('nav-home');
            if (homeNav) {
                homeNav.classList.add('active', 'text-white');
                homeNav.classList.remove('text-light');
            }
        }
        return;
    }

    // 对于NFT平台内的导航
    if (currentPlatform === 'nft') {
        // 移除NFT平台内所有导航项的active状态
        document.querySelectorAll('#nft-platform .navbar-nav .nav-link').forEach(link => {
            link.classList.remove('active');
        });

        // 根据当前section设置对应导航项为active
        let navId = '';
        switch (sectionId) {
            case 'all-nfts-section':
                navId = 'nav-all-nfts';
                break;
            case 'mint-nft-section':
                navId = 'nav-mint';
                break;
            case 'metadata-section':
                navId = 'nav-metadata';
                break;
            case 'my-nfts-section':
                navId = 'nav-my-nfts';
                break;
            case 'requests-section':
                navId = 'nav-requests';
                break;
        }

        if (navId) {
            const navItem = document.getElementById(navId);
            if (navItem) {
                navItem.classList.add('active');
            }
        }
    }
}

// 添加手动断开连接功能
function disconnectWallet() {
    isConnected = false;
    currentAccount = null;
    clearWalletState();
    updateUIForDisconnectedWallet();
    console.log('钱包连接已断开');
}

// 申请子NFT函数
async function requestChildNFT(parentTokenId) {
    if (!isConnected) {
        showError('请先连接钱包');
        return;
    }

    // 显示申请子NFT模态框
    document.getElementById('request-parent-token-id').value = parentTokenId;
    const modal = new bootstrap.Modal(document.getElementById('request-child-modal'));
    modal.show();
}

// 创建子NFT函数（为拥有者提供）
async function createChildNFT(parentTokenId) {
    if (!isConnected) {
        showError('请先连接钱包');
        return;
    }

    // 显示创建子NFT模态框
    document.getElementById('create-parent-token-id').value = parentTokenId;
    const modal = new bootstrap.Modal(document.getElementById('create-child-modal'));
    modal.show();
}

// 处理申请子NFT表单提交
async function handleRequestChildFormSubmit(event) {
    event.preventDefault();
    if (!isConnected) return;

    const parentTokenId = document.getElementById('request-parent-token-id').value.trim();
    const uri = document.getElementById('request-child-uri').value.trim();

    if (!parentTokenId || !uri) {
        showError('请填写所有必填字段');
        return;
    }

    showLoading(true);
    try {
        // 创建要签名的消息
        const message = JSON.stringify({
            action: 'request_child_nft',
            parentTokenId: parentTokenId,
            uri: uri,
            timestamp: Date.now()
        });

        // 获取签名
        const signature = await signMessage(message);

        // 发送申请请求
        const response = await fetch(`${API_BASE_URL}/nft/request-child`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                address: currentAccount,
                signature: signature,
                message: message,
                parentTokenId: parentTokenId,
                applicantAddress: currentAccount,
                uri: uri
            })
        });

        const result = await response.json();

        if (response.ok) {
            // 隐藏模态框
            const modal = bootstrap.Modal.getInstance(document.getElementById('request-child-modal'));
            modal.hide();

            // 显示成功消息
            showSuccess('子NFT申请已提交，等待审批');

            // 重置表单
            document.getElementById('request-child-form').reset();
        } else {
            showError(result.error || '提交申请失败');
        }
    } catch (error) {
        console.error('提交申请出错:', error);
        showError('提交申请失败: ' + error.message);
    } finally {
        showLoading(false);
    }
}

// 处理创建子NFT表单提交
async function handleCreateChildFormSubmit(event) {
    event.preventDefault();
    if (!isConnected) return;

    const parentTokenId = document.getElementById('create-parent-token-id').value.trim();
    const recipient = document.getElementById('create-recipient').value.trim();
    const uri = document.getElementById('create-child-uri').value.trim();

    if (!parentTokenId || !recipient || !uri) {
        showError('请填写所有必填字段');
        return;
    }

    showLoading(true);
    try {
        // 创建要签名的消息
        const message = JSON.stringify({
            action: 'create_child_nft',
            parentTokenId: parentTokenId,
            recipient: recipient,
            uri: uri,
            timestamp: Date.now()
        });

        // 获取签名
        const signature = await signMessage(message);

        // 发送创建请求
        const response = await fetch(`${API_BASE_URL}/nft/createChild`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                address: currentAccount,
                signature: signature,
                message: message,
                parentTokenId: parentTokenId,
                recipient: recipient,
                uri: uri
            })
        });

        const result = await response.json();

        if (response.ok) {
            // 隐藏模态框
            const modal = bootstrap.Modal.getInstance(document.getElementById('create-child-modal'));
            modal.hide();

            // 显示成功消息
            showSuccess('子NFT创建交易已提交', result.transactionHash);

            // 重置表单
            document.getElementById('create-child-form').reset();
        } else {
            showError(result.error || '创建子NFT失败');
        }
    } catch (error) {
        console.error('创建子NFT出错:', error);
        showError('创建子NFT失败: ' + error.message);
    } finally {
        showLoading(false);
    }
}

// 处理元数据表单提交
async function handleMetadataFormSubmit(event) {
    event.preventDefault();

    const name = document.getElementById('metadata-name').value.trim();
    const description = document.getElementById('metadata-description').value.trim();
    const externalUrl = document.getElementById('metadata-external-url').value.trim();
    const image = document.getElementById('metadata-image').value.trim();
    const policy = document.getElementById('metadata-policy').value.trim();
    const ciphertext = document.getElementById('metadata-ciphertext').value.trim();

    if (!name || !description || !image || !policy || !ciphertext) {
        showError('请填写所有必填字段');
        return;
    }

    showLoading(true);
    try {
        const response = await fetch(`${API_BASE_URL}/metadata`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: name,
                description: description,
                external_url: externalUrl,
                image: image,
                policy: policy,
                ciphertext: ciphertext
            })
        });

        const result = await response.json();

        if (response.ok) {
            showSuccess(`元数据创建成功！IPFS哈希: ${result.ipfs_hash}`, result.ipfs_hash);
            // 重置表单
            document.getElementById('metadata-form').reset();
            updateMetadataPreview();
            // 刷新元数据列表
            loadMetadataList();
        } else {
            showError(result.error || '创建元数据失败');
        }
    } catch (error) {
        console.error('创建元数据出错:', error);
        showError('创建元数据失败: ' + error.message);
    } finally {
        showLoading(false);
    }
}

// 更新元数据预览
function updateMetadataPreview() {
    const name = document.getElementById('metadata-name').value.trim();
    const description = document.getElementById('metadata-description').value.trim();
    const externalUrl = document.getElementById('metadata-external-url').value.trim();
    const image = document.getElementById('metadata-image').value.trim();
    const policy = document.getElementById('metadata-policy').value.trim();
    const ciphertext = document.getElementById('metadata-ciphertext').value.trim();

    const metadata = {
        description: description,
        external_url: externalUrl,
        image: image,
        name: name,
        attributes: [
            {
                trait_type: "Policy",
                value: policy
            },
            {
                trait_type: "Encrypted_ciphertext",
                value: ciphertext
            }
        ]
    };

    document.getElementById('metadata-preview').textContent = JSON.stringify(metadata, null, 2);
}

// 加载元数据列表
async function loadMetadataList() {
    try {
        const response = await fetch(`${API_BASE_URL}/metadata`);
        const result = await response.json();

        const metadataList = document.getElementById('metadata-list');

        if (response.ok && result.metadata) {
            if (result.metadata.length === 0) {
                metadataList.innerHTML = '<tr><td colspan="5" class="text-center">暂无元数据</td></tr>';
                return;
            }

            metadataList.innerHTML = result.metadata.map(metadata => `
                <tr>
                    <td>${metadata.name}</td>
                    <td title="${metadata.description}">${metadata.description.length > 50 ? metadata.description.substring(0, 50) + '...' : metadata.description}</td>
                    <td>
                        <code class="small">${metadata.ipfs_hash}</code>
                        <button class="btn btn-sm btn-outline-primary ms-1" onclick="copyToClipboard('${metadata.ipfs_hash}')">复制</button>
                        <button class="btn btn-sm btn-outline-secondary ms-1" onclick="copyIPFSUrl('${metadata.ipfs_hash}')">复制IPFS链接</button>
                    </td>
                    <td>${new Date(metadata.CreatedAt).toLocaleString()}</td>
                    <td>
                        <button class="btn btn-sm btn-primary" onclick="viewMetadata('${metadata.ipfs_hash}')">查看</button>
                        <button class="btn btn-sm btn-success" onclick="useAsURI('${metadata.ipfs_hash}')">用作URI</button>
                    </td>
                </tr>
            `).join('');
        } else {
            metadataList.innerHTML = '<tr><td colspan="5" class="text-center text-danger">加载元数据失败</td></tr>';
        }
    } catch (error) {
        console.error('加载元数据列表出错:', error);
        document.getElementById('metadata-list').innerHTML = '<tr><td colspan="5" class="text-center text-danger">加载元数据失败</td></tr>';
    }
}

// 复制到剪贴板
function copyToClipboard(text) {
    navigator.clipboard.writeText(text).then(() => {
        showSuccess('已复制到剪贴板');
    }).catch(err => {
        console.error('复制失败:', err);
        showError('复制失败');
    });
}

// 复制IPFS链接
function copyIPFSUrl(hash) {
    const ipfsUrl = `ipfs://${hash}`;
    copyToClipboard(ipfsUrl);
}

// 查看元数据详情
async function viewMetadata(hash) {
    try {
        const response = await fetch(`${API_BASE_URL}/metadata/${hash}`);
        const metadata = await response.json();

        if (response.ok) {
            // 显示元数据详情的模态框或者新窗口
            alert(`元数据详情:\n${JSON.stringify(metadata, null, 2)}`);
        } else {
            showError('获取元数据详情失败');
        }
    } catch (error) {
        console.error('获取元数据详情出错:', error);
        showError('获取元数据详情失败');
    }
}

// 使用哈希作为URI（用于铸造NFT）
function useAsURI(hash) {
    const ipfsUrl = `ipfs://${hash}`;

    // 如果当前在铸造NFT页面，直接填入URI字段
    const uriInput = document.getElementById('nft-uri');
    if (uriInput) {
        uriInput.value = ipfsUrl;
        showSuccess('已将IPFS链接填入铸造NFT的URI字段');
        showSection('mint-nft-section');
    } else {
        // 否则复制到剪贴板
        copyIPFSUrl(hash);
    }
}

// 更新元数据函数
async function updateMetadata(tokenId, contractType) {
    if (!isConnected) {
        showError('请先连接钱包');
        return;
    }

    // 显示更新元数据模态框
    document.getElementById('update-token-id').value = tokenId;
    document.getElementById('update-contract-type').value = contractType;
    const modal = new bootstrap.Modal(document.getElementById('update-metadata-modal'));
    modal.show();
}

// 处理更新元数据表单提交
async function handleUpdateMetadataFormSubmit(event) {
    event.preventDefault();
    if (!isConnected) return;

    const tokenId = document.getElementById('update-token-id').value.trim();
    const contractType = document.getElementById('update-contract-type').value.trim();
    const newUri = document.getElementById('update-new-uri').value.trim();

    if (!tokenId || !contractType || !newUri) {
        showError('请填写所有必填字段');
        return;
    }

    showLoading(true);
    try {
        // 创建要签名的消息
        const message = JSON.stringify({
            action: 'update_metadata',
            tokenId: tokenId,
            contractType: contractType,
            newUri: newUri,
            timestamp: Date.now()
        });

        // 获取签名
        const signature = await signMessage(message);

        // 发送更新请求
        const response = await fetch(`${API_BASE_URL}/nft/update-metadata`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                address: currentAccount,
                signature: signature,
                message: message,
                tokenId: tokenId,
                contractType: contractType,
                newUri: newUri
            })
        });

        const result = await response.json();

        if (response.ok) {
            // 隐藏模态框
            const modal = bootstrap.Modal.getInstance(document.getElementById('update-metadata-modal'));
            modal.hide();

            // 显示成功消息
            showSuccess('元数据更新交易已提交', result.transactionHash);

            // 重置表单并刷新NFT列表
            document.getElementById('update-metadata-form').reset();
            loadMyNFTs();
        } else {
            showError(result.error || '更新元数据失败');
        }
    } catch (error) {
        console.error('更新元数据出错:', error);
        showError('更新元数据失败: ' + error.message);
    } finally {
        showLoading(false);
    }
}

// 显示申请管理模态框
async function showRequestsModal() {
    if (!isConnected) {
        showError('请先连接钱包');
        return;
    }

    const modal = new bootstrap.Modal(document.getElementById('requests-modal'));
    modal.show();
    await loadAllRequests();
}

// 加载所有申请记录
async function loadAllRequests() {
    if (!isConnected) return;

    const requestsList = document.getElementById('requests-list');
    requestsList.innerHTML = '<div class="text-center"><p>加载中...</p></div>';

    try {
        // 创建要签名的消息
        const message = JSON.stringify({
            action: 'get_all_requests',
            timestamp: Date.now()
        });

        // 获取签名
        const signature = await signMessage(message);

        // 发送获取申请记录的请求
        const response = await fetch(`${API_BASE_URL}/nft/all-requests`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                'X-Ethereum-Address': currentAccount,
                'X-Ethereum-Signature': signature,
                'X-Ethereum-Message': message
            }
        });

        const result = await response.json();

        if (response.ok && result.requests) {
            if (result.requests.length === 0) {
                requestsList.innerHTML = '<div class="text-center"><p>没有申请记录</p></div>';
                return;
            }

            const requestsHtml = result.requests.map(req => {
                const statusBadge = req.status === 'pending' ?
                    '<span class="badge bg-warning">待处理</span>' :
                    req.status === 'approved' ?
                        '<span class="badge bg-success">已完成</span>' :
                        '<span class="badge bg-danger">已拒绝</span>';

                // 操作按钮 - 只有在可操作且状态为pending时才显示
                const actionButtons = req.canOperate && req.status === 'pending' ?
                    `<button class="btn btn-success btn-sm me-2" onclick="processRequest(${req.ID}, 'approve')">批准</button>
                     <button class="btn btn-danger btn-sm" onclick="processRequest(${req.ID}, 'reject')">拒绝</button>` :
                    req.status === 'approved' ?
                        '<span class="text-success">已完成</span>' :
                        req.status === 'rejected' ?
                            '<span class="text-danger">已拒绝</span>' :
                            '<span class="text-muted">无操作权限</span>';

                return `
                    <div class="card mb-3 ${req.status === 'approved' ? 'border-success' : req.status === 'rejected' ? 'border-danger' : ''}">
                        <div class="card-body">
                            <div class="row">
                                <div class="col-md-6">
                                    <h6 class="card-title">申请 #${req.ID}</h6>
                                    <p class="card-text"><small class="text-muted">父NFT ID: ${req.parentTokenId}</small></p>
                                    <p class="card-text"><small class="text-muted">申请人: ${req.applicantAddress.substring(0, 10)}...</small></p>
                                    ${req.childTokenId ? `<p class="card-text"><small class="text-muted">子NFT ID: ${req.childTokenId}</small></p>` : ''}
                                </div>
                                <div class="col-md-6 text-end">
                                    <div class="mb-2">${statusBadge}</div>
                                    <div>${actionButtons}</div>
                                    <p class="card-text"><small class="text-muted">创建时间: ${new Date(req.CreatedAt).toLocaleString()}</small></p>
                                </div>
                            </div>
                            <div class="row mt-2">
                                <div class="col-12">
                                    <p class="card-text text-truncate"><small>URI: ${req.uri}</small></p>
                                </div>
                            </div>
                        </div>
                    </div>
                `;
            }).join('');

            requestsList.innerHTML = requestsHtml;
        } else {
            requestsList.innerHTML = '<div class="text-center text-danger"><p>加载申请记录失败</p></div>';
        }
    } catch (error) {
        console.error('加载申请记录出错:', error);
        requestsList.innerHTML = '<div class="text-center text-danger"><p>加载申请记录失败</p></div>';
    }
}

// 隐藏不相关的界面元素
function hideIrrelevantSections() {
    console.log('隐藏不相关的界面元素');

    // 只隐藏DID和VC相关的界面元素，保留NFT和ABE界面
    const elementsToHide = [
        'didCreateCard',
        'did-list-section',
        'vc-issue-section',
        'vc-manage-section',
        'vp-create-section'
    ];

    // 遍历并隐藏元素
    elementsToHide.forEach(id => {
        const element = document.getElementById(id);
        if (element) {
            console.log(`隐藏元素: ${id}`);
            element.classList.add('d-none');
        } else {
            console.log(`未找到元素: ${id}`);
        }
    });

    // 确保NFT和ABE相关界面可见
    const nftSections = [
        'all-nfts-section',
        'mint-nft-section',
        'metadata-section',
        'my-nfts-section',
        'requests-section'
    ];

    // 默认显示全部NFT界面
    const allNftsSection = document.getElementById('all-nfts-section');
    if (allNftsSection) {
        allNftsSection.classList.remove('d-none');
    }
} 