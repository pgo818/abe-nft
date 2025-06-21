// ABE管理系统JavaScript文件

// 全局变量
let currentPlatform = 'home';
let abeSystemKeys = null;
let latestIPFSHash = null;
let latestCiphertext = null;
let latestUserKey = null;

// DOM加载完成后初始化ABE相关功能
document.addEventListener('DOMContentLoaded', () => {
    setupABEEventListeners();
    setupPlatformNavigation();
});

// 设置ABE事件监听器
function setupABEEventListeners() {
    const abeSetupForm = document.getElementById('abe-setup-form');
    const abeUploadForm = document.getElementById('abe-upload-form');
    const abeKeygenForm = document.getElementById('abe-keygen-form');
    const abeEncryptForm = document.getElementById('abe-encrypt-form');
    const abeDecryptForm = document.getElementById('abe-decrypt-form');

    if (abeSetupForm) {
        abeSetupForm.addEventListener('submit', handleABESetupSubmit);
    }
    if (abeUploadForm) {
        abeUploadForm.addEventListener('submit', handleABEUploadSubmit);
    }
    if (abeKeygenForm) {
        abeKeygenForm.addEventListener('submit', handleABEKeygenSubmit);
    }
    if (abeEncryptForm) {
        abeEncryptForm.addEventListener('submit', handleABEEncryptSubmit);
    }
    if (abeDecryptForm) {
        abeDecryptForm.addEventListener('submit', handleABEDecryptSubmit);
    }

    setupABENavigation();
    addExampleButtons();
}

// 设置平台切换功能
function setupPlatformNavigation() {
    const navHome = document.getElementById('nav-home');
    if (navHome) {
        navHome.addEventListener('click', showHome);
    }
}

// 显示首页
function showHome() {
    currentPlatform = 'home';
    document.getElementById('platform-selector').classList.remove('d-none');
    document.getElementById('breadcrumb-section').classList.add('d-none');
    document.getElementById('nft-platform').classList.remove('active');
    document.getElementById('abe-platform').classList.remove('active');
    updateNavigation('home');
}

// 显示指定平台
function showPlatform(platform) {
    currentPlatform = platform;
    document.getElementById('platform-selector').classList.add('d-none');
    document.getElementById('breadcrumb-section').classList.remove('d-none');

    // 隐藏所有平台
    document.getElementById('nft-platform').classList.remove('active');
    document.getElementById('abe-platform').classList.remove('active');
    document.getElementById('did-vc-platform').classList.remove('active');

    if (platform === 'nft') {
        document.getElementById('current-platform').textContent = 'NFT管理平台';
        document.getElementById('nft-platform').classList.add('active');
        showSection('all-nfts-section');
    } else if (platform === 'abe') {
        document.getElementById('current-platform').textContent = 'ABE加密管理';
        document.getElementById('abe-platform').classList.add('active');
        showABESection('abe-setup-section');
    } else if (platform === 'did') {
        document.getElementById('current-platform').textContent = 'DID和VC管理';
        document.getElementById('did-vc-platform').classList.add('active');
        showDIDVCSection('did-list-section');
        // 初始化DID和VC功能
        if (typeof initDIDVC === 'function') {
            initDIDVC();
        }
    }

    updateNavigation(platform);
}

// 显示DID和VC管理的指定部分
function showDIDVCSection(sectionId) {
    // 隐藏所有DID和VC部分
    const sections = ['did-list-section', 'vc-issue-section', 'vc-manage-section', 'vp-create-section'];
    sections.forEach(id => {
        const element = document.getElementById(id);
        if (element) {
            element.classList.add('d-none');
        }
    });

    // 显示指定部分
    const targetSection = document.getElementById(sectionId);
    if (targetSection) {
        targetSection.classList.remove('d-none');
    }

    // 更新导航状态
    updateDIDVCNavigation(sectionId);
}

// 更新DID和VC导航状态
function updateDIDVCNavigation(sectionId) {
    // 移除所有导航项的active状态
    document.querySelectorAll('#did-vc-platform .navbar-nav .nav-link').forEach(link => {
        link.classList.remove('active');
    });

    // 根据当前section设置对应导航项为active
    let navId = '';
    switch (sectionId) {
        case 'did-list-section':
            navId = 'nav-did-list';
            break;
        case 'vc-issue-section':
            navId = 'nav-vc-issue';
            break;
        case 'vc-manage-section':
            navId = 'nav-vc-manage';
            break;
        case 'vp-create-section':
            navId = 'nav-vp-create';
            break;
    }

    if (navId) {
        const navItem = document.getElementById(navId);
        if (navItem) {
            navItem.classList.add('active');
        }
    }
}

// 初始化DID和VC功能
function initDIDVC() {
    // 绑定导航事件
    const navDidList = document.getElementById('nav-did-list');
    const navDidCreate = document.getElementById('nav-did-create');
    const navVcIssue = document.getElementById('nav-vc-issue');
    const navVcManage = document.getElementById('nav-vc-manage');
    const navVpCreate = document.getElementById('nav-vp-create');

    if (navDidList) {
        navDidList.addEventListener('click', (e) => {
            e.preventDefault();
            showDIDVCSection('did-list-section');
        });
    }

    if (navDidCreate) {
        navDidCreate.addEventListener('click', (e) => {
            e.preventDefault();
            // 显示创建DID卡片
            document.getElementById('didCreateCard').scrollIntoView();
        });
    }

    if (navVcIssue) {
        navVcIssue.addEventListener('click', (e) => {
            e.preventDefault();
            showDIDVCSection('vc-issue-section');
        });
    }

    if (navVcManage) {
        navVcManage.addEventListener('click', (e) => {
            e.preventDefault();
            showDIDVCSection('vc-manage-section');
        });
    }

    if (navVpCreate) {
        navVpCreate.addEventListener('click', (e) => {
            e.preventDefault();
            showDIDVCSection('vp-create-section');
        });
    }
}

// 设置ABE导航
function setupABENavigation() {
    const navABESetup = document.getElementById('nav-abe-setup');
    const navABEUpload = document.getElementById('nav-abe-upload');
    const navABEEncrypt = document.getElementById('nav-abe-encrypt');
    const navABEDecrypt = document.getElementById('nav-abe-decrypt');
    const navABEKeygen = document.getElementById('nav-abe-keygen');
    const navABELogs = document.getElementById('nav-abe-logs');

    if (navABESetup) {
        navABESetup.addEventListener('click', () => showABESection('abe-setup-section'));
    }
    if (navABEUpload) {
        navABEUpload.addEventListener('click', () => showABESection('abe-upload-section'));
    }
    if (navABEEncrypt) {
        navABEEncrypt.addEventListener('click', () => showABESection('abe-encrypt-section'));
    }
    if (navABEDecrypt) {
        navABEDecrypt.addEventListener('click', () => showABESection('abe-decrypt-section'));
    }
    if (navABEKeygen) {
        navABEKeygen.addEventListener('click', () => showABESection('abe-keygen-section'));
    }
    if (navABELogs) {
        navABELogs.addEventListener('click', () => {
            showABESection('abe-logs-section');
            loadABELogs();
        });
    }
}

// 显示ABE部分
function showABESection(sectionId) {
    const abeSections = document.querySelectorAll('#abe-platform .section');
    abeSections.forEach(section => {
        section.classList.add('d-none');
    });

    const targetSection = document.getElementById(sectionId);
    if (targetSection) {
        targetSection.classList.remove('d-none');
    }

    updateABENavigation(sectionId);
}

// 显示DID/VC部分
function showDIDVCSection(sectionId) {
    const didvcSections = document.querySelectorAll('#did-vc-platform .section');
    didvcSections.forEach(section => {
        section.classList.add('d-none');
    });

    const targetSection = document.getElementById(sectionId);
    if (targetSection) {
        targetSection.classList.remove('d-none');
    }

    updateDIDVCNavigation(sectionId);
}

// 更新ABE导航样式
function updateABENavigation(activeSectionId) {
    const abeNavLinks = document.querySelectorAll('#abe-platform .navbar-nav .nav-link');
    abeNavLinks.forEach(link => {
        link.classList.remove('active');
    });

    const navMap = {
        'abe-setup-section': 'nav-abe-setup',
        'abe-upload-section': 'nav-abe-upload',
        'abe-encrypt-section': 'nav-abe-encrypt',
        'abe-decrypt-section': 'nav-abe-decrypt',
        'abe-keygen-section': 'nav-abe-keygen',
        'abe-logs-section': 'nav-abe-logs'
    };

    const activeNavId = navMap[activeSectionId];
    if (activeNavId) {
        const activeNav = document.getElementById(activeNavId);
        if (activeNav) {
            activeNav.classList.add('active');
        }
    }
}

// 更新DID/VC导航样式
function updateDIDVCNavigation(activeSectionId) {
    const didvcNavLinks = document.querySelectorAll('#did-vc-platform .navbar-nav .nav-link');
    didvcNavLinks.forEach(link => {
        link.classList.remove('active');
    });

    const navMap = {
        'did-list-section': 'nav-did-list',
        'did-create-section': 'nav-did-create',
        'vc-issue-section': 'nav-vc-issue',
        'vc-manage-section': 'nav-vc-manage',
        'vp-create-section': 'nav-vp-create'
    };

    const activeNavId = navMap[activeSectionId];
    if (activeNavId) {
        const activeNav = document.getElementById(activeNavId);
        if (activeNav) {
            activeNav.classList.add('active');
        }
    }
}

// 处理ABE系统初始化表单提交
async function handleABESetupSubmit(event) {
    event.preventDefault();

    const attributesText = document.getElementById('abe-attributes').value.trim();
    if (!attributesText) {
        showError('请输入属性列表');
        return;
    }

    const attributes = attributesText.split('\n').map(attr => attr.trim()).filter(attr => attr);

    showLoading(true);

    try {
        const response = await fetch(`${API_BASE_URL}/abe/setup`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                gamma: attributes
            })
        });

        const result = await response.json();

        if (response.ok) {
            abeSystemKeys = {
                pubKey: result.pub_key,
                secKey: result.sec_key,
                systemKeyId: result.system_key_id
            };

            displayABESetupResult(result);
            showSuccess('ABE系统初始化成功！系统密钥已生成。');
        } else {
            showError('ABE系统初始化失败: ' + (result.error || '未知错误'));
        }
    } catch (error) {
        console.error('ABE系统初始化错误:', error);
        showError('网络错误，请检查连接后重试');
    } finally {
        showLoading(false);
    }
}

// 显示ABE系统初始化结果
function displayABESetupResult(result) {
    const resultDiv = document.getElementById('abe-setup-result');
    resultDiv.innerHTML = `
        <div class="alert alert-success">
            <h6><i class="bi bi-check-circle me-2"></i>系统初始化成功！</h6>
            <p><strong>系统密钥ID:</strong> ${result.system_key_id}</p>
            <p><strong>创建时间:</strong> ${new Date().toLocaleString()}</p>
        </div>
        <div class="mb-3">
            <label class="form-label"><strong>公钥 (PubKey):</strong></label>
            <textarea class="form-control" rows="3" readonly onclick="this.select()">${result.pub_key}</textarea>
            <button class="btn btn-sm btn-outline-primary mt-1" onclick="copyToClipboard('${result.pub_key}')">
                <i class="bi bi-clipboard me-1"></i>复制公钥
            </button>
        </div>
        <div class="mb-3">
            <label class="form-label"><strong>主密钥 (SecKey):</strong></label>
            <textarea class="form-control" rows="3" readonly onclick="this.select()">${result.sec_key}</textarea>
            <button class="btn btn-sm btn-outline-danger mt-1" onclick="copyToClipboard('${result.sec_key}')">
                <i class="bi bi-clipboard me-1"></i>复制主密钥 (请妥善保管)
            </button>
        </div>
        <div class="alert alert-warning">
            <i class="bi bi-exclamation-triangle me-2"></i>
            <small>请妥善保管主密钥，它是生成用户属性密钥的关键。公钥可以公开分享用于加密。</small>
        </div>
    `;
}

// 处理医疗数据上传表单提交
async function handleABEUploadSubmit(event) {
    event.preventDefault();

    const textData = document.getElementById('medical-data-text').value.trim();
    const fileInput = document.getElementById('medical-data-file');
    const file = fileInput.files[0];

    if (!textData && !file) {
        showError('请输入文本数据或选择文件');
        return;
    }

    showLoading(true);

    try {
        let dataToUpload = textData;

        // 如果选择了文件，读取文件内容
        if (file) {
            dataToUpload = await readFileContent(file);
        }

        const response = await fetch(`${API_BASE_URL}/ipfs/upload`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                data: dataToUpload,
                filename: file ? file.name : 'medical-data.txt'
            })
        });

        const result = await response.json();

        if (response.ok) {
            latestIPFSHash = result.hash;
            displayABEUploadResult(result, dataToUpload);
            showSuccess('医疗数据已成功上传到IPFS！');
        } else {
            showError('上传失败: ' + (result.error || '未知错误'));
        }
    } catch (error) {
        console.error('上传医疗数据错误:', error);
        showError('网络错误，请检查连接后重试');
    } finally {
        showLoading(false);
    }
}

// 读取文件内容
function readFileContent(file) {
    return new Promise((resolve, reject) => {
        const reader = new FileReader();
        reader.onload = (e) => resolve(e.target.result);
        reader.onerror = (e) => reject(e);
        reader.readAsText(file);
    });
}

// 显示医疗数据上传结果
function displayABEUploadResult(result, originalData) {
    const resultDiv = document.getElementById('abe-upload-result');
    resultDiv.innerHTML = `
        <div class="alert alert-success">
            <h6><i class="bi bi-check-circle me-2"></i>上传成功！</h6>
            <p><strong>IPFS哈希:</strong> ${result.hash}</p>
            <p><strong>数据大小:</strong> ${originalData.length} 字符</p>
            <p><strong>上传时间:</strong> ${new Date().toLocaleString()}</p>
        </div>
        <div class="mb-3">
            <label class="form-label"><strong>IPFS哈希值:</strong></label>
            <input type="text" class="form-control" value="${result.hash}" readonly onclick="this.select()">
            <button class="btn btn-sm btn-outline-primary mt-1" onclick="copyToClipboard('${result.hash}')">
                <i class="bi bi-clipboard me-1"></i>复制哈希
            </button>
        </div>
        <div class="mb-3">
            <label class="form-label"><strong>访问URL:</strong></label>
            <input type="text" class="form-control" value="${result.url || 'ipfs://' + result.hash}" readonly onclick="this.select()">
            <button class="btn btn-sm btn-outline-primary mt-1" onclick="copyToClipboard('${result.url || 'ipfs://' + result.hash}')">
                <i class="bi bi-clipboard me-1"></i>复制URL
            </button>
        </div>
        <div class="alert alert-info">
            <i class="bi bi-info-circle me-2"></i>
            <small>医疗数据已安全存储在IPFS网络中。您可以使用此哈希值进行后续的ABE加密操作。</small>
        </div>
    `;
}

// 处理ABE密钥生成表单提交
async function handleABEKeygenSubmit(event) {
    event.preventDefault();

    const attributesText = document.getElementById('keygen-attributes').value.trim();

    if (!attributesText) {
        showError('请填写用户属性');
        return;
    }

    const attributes = attributesText.split('\n').map(attr => attr.trim()).filter(attr => attr);

    showLoading(true);

    try {
        const response = await fetch(`${API_BASE_URL}/abe/keygen`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                gamma: attributes
            })
        });

        const result = await response.json();

        if (response.ok) {
            latestUserKey = result.user_key;
            displayABEKeygenResult(result, attributes);
            showSuccess('用户属性密钥生成成功！');
        } else {
            showError('密钥生成失败: ' + (result.error || '未知错误'));
        }
    } catch (error) {
        console.error('ABE密钥生成错误:', error);
        showError('网络错误，请检查连接后重试');
    } finally {
        showLoading(false);
    }
}

// 显示ABE密钥生成结果
function displayABEKeygenResult(result, attributes) {
    const resultDiv = document.getElementById('abe-keygen-result');
    resultDiv.innerHTML = `
        <div class="alert alert-success">
            <h6><i class="bi bi-key me-2"></i>属性密钥生成成功！</h6>
            <p><strong>用户密钥ID:</strong> ${result.user_key_id}</p>
            <p><strong>属性:</strong> ${attributes.join(', ')}</p>
        </div>
        <div class="mb-3">
            <label class="form-label"><strong>属性密钥 (AttribKeys):</strong></label>
            <textarea class="form-control" rows="4" readonly onclick="this.select()">${result.attrib_keys}</textarea>
            <button class="btn btn-sm btn-outline-primary mt-1" onclick="copyToClipboard('${result.attrib_keys}')">
                <i class="bi bi-clipboard me-1"></i>复制属性密钥
            </button>
        </div>
        <div class="alert alert-info">
            <i class="bi bi-info-circle me-2"></i>
            <small>此属性密钥可用于解密符合您属性的加密数据。请妥善保管，不要泄露给他人。</small>
        </div>
    `;
}

// 处理ABE加密表单提交
async function handleABEEncryptSubmit(event) {
    event.preventDefault();

    const message = document.getElementById('encrypt-message').value.trim();
    const policy = document.getElementById('encrypt-policy').value.trim();

    if (!message || !policy) {
        showError('请填写要加密的数据和访问策略');
        return;
    }

    showLoading(true);

    try {
        const response = await fetch(`${API_BASE_URL}/abe/encrypt`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                message: message,
                policy: policy
            })
        });

        const result = await response.json();

        if (response.ok) {
            latestCiphertext = result.ciphertext;
            displayABEEncryptResult(result, policy, message.length);
            showSuccess('数据加密成功！');
        } else {
            showError('数据加密失败: ' + (result.error || '未知错误'));
        }
    } catch (error) {
        console.error('ABE加密错误:', error);
        showError('网络错误，请检查连接后重试');
    } finally {
        showLoading(false);
    }
}

// 显示ABE加密结果
function displayABEEncryptResult(result, policy, messageLength) {
    const resultDiv = document.getElementById('abe-encrypt-result');
    resultDiv.innerHTML = `
        <div class="alert alert-success">
            <h6><i class="bi bi-lock me-2"></i>数据加密成功！</h6>
            <p><strong>密文ID:</strong> ${result.ciphertext_id}</p>
            <p><strong>访问策略:</strong> ${policy}</p>
            <p><strong>原文长度:</strong> ${messageLength} 字符</p>
        </div>
        <div class="mb-3">
            <label class="form-label"><strong>密文 (Cipher):</strong></label>
            <textarea class="form-control" rows="5" readonly onclick="this.select()">${result.cipher}</textarea>
            <button class="btn btn-sm btn-outline-primary mt-1" onclick="copyToClipboard('${result.cipher}')">
                <i class="bi bi-clipboard me-1"></i>复制密文
            </button>
        </div>
        <div class="alert alert-info">
            <i class="bi bi-shield-check me-2"></i>
            <small>数据已根据指定的访问策略进行加密。只有满足策略条件的用户才能解密此数据。</small>
        </div>
    `;
}

// 处理ABE解密表单提交
async function handleABEDecryptSubmit(event) {
    event.preventDefault();

    const cipher = document.getElementById('decrypt-ciphertext').value.trim();
    const attribKeys = document.getElementById('decrypt-userkey').value.trim();

    if (!cipher || !attribKeys) {
        showError('请填写密文和用户属性密钥');
        return;
    }

    showLoading(true);

    try {
        const response = await fetch(`${API_BASE_URL}/abe/decrypt`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                cipher: cipher,
                attrib_keys: attribKeys
            })
        });

        const result = await response.json();

        if (response.ok) {
            displayABEDecryptResult(result);
            showSuccess('数据解密成功！');
        } else {
            showError('数据解密失败: ' + (result.error || '未知错误'));
            displayABEDecryptError(result.error);
        }
    } catch (error) {
        console.error('ABE解密错误:', error);
        showError('网络错误，请检查连接后重试');
    } finally {
        showLoading(false);
    }
}

// 显示ABE解密结果
function displayABEDecryptResult(result) {
    const resultDiv = document.getElementById('abe-decrypt-result');
    resultDiv.innerHTML = `
        <div class="alert alert-success">
            <h6><i class="bi bi-unlock me-2"></i>数据解密成功！</h6>
        </div>
        <div class="mb-3">
            <label class="form-label"><strong>解密后的数据:</strong></label>
            <textarea class="form-control" rows="4" readonly onclick="this.select()">${result.message}</textarea>
            <button class="btn btn-sm btn-outline-primary mt-1" onclick="copyToClipboard('${result.message}')">
                <i class="bi bi-clipboard me-1"></i>复制内容
            </button>
        </div>
        <div class="alert alert-success">
            <i class="bi bi-check-circle me-2"></i>
            <small>您的属性满足访问策略要求，解密成功。</small>
        </div>
    `;
}

// 显示ABE解密错误
function displayABEDecryptError(errorMessage) {
    const resultDiv = document.getElementById('abe-decrypt-result');
    resultDiv.innerHTML = `
        <div class="alert alert-danger">
            <h6><i class="bi bi-x-circle me-2"></i>解密失败</h6>
            <p>${errorMessage}</p>
        </div>
        <div class="alert alert-warning">
            <i class="bi bi-exclamation-triangle me-2"></i>
            <small>可能的原因：<br>
            1. 您的属性不满足访问策略要求<br>
            2. 密文格式错误<br>
            3. 属性密钥与密文不匹配<br>
            4. 公钥错误</small>
        </div>
    `;
}

// 加载ABE操作日志
async function loadABELogs() {
    showLoading(true);

    try {
        const mockLogs = [
            {
                created_at: new Date().toISOString(),
                operation_type: 'setup',
                user_id: 1,
                ip_address: '192.168.1.100',
                details: JSON.stringify({ attributes: ['age:18-25', 'department:IT'] })
            },
            {
                created_at: new Date(Date.now() - 3600000).toISOString(),
                operation_type: 'keygen',
                user_id: 1,
                ip_address: '192.168.1.100',
                details: JSON.stringify({ user_key_id: 1, attributes: ['age:22', 'department:IT'] })
            },
            {
                created_at: new Date(Date.now() - 7200000).toISOString(),
                operation_type: 'encrypt',
                user_id: 1,
                ip_address: '192.168.1.100',
                details: JSON.stringify({ ciphertext_id: 1, policy: 'age:18-25 AND department:IT' })
            }
        ];

        displayABELogs(mockLogs);
    } catch (error) {
        console.error('加载ABE日志错误:', error);
        showError('加载操作日志失败');
    } finally {
        showLoading(false);
    }
}

// 显示ABE操作日志
function displayABELogs(logs) {
    const logsList = document.getElementById('abe-logs-list');

    if (!logs || logs.length === 0) {
        logsList.innerHTML = `
            <tr>
                <td colspan="5" class="text-center text-muted">暂无操作日志</td>
            </tr>
        `;
        return;
    }

    logsList.innerHTML = logs.map(log => {
        const details = JSON.parse(log.details || '{}');
        const detailsText = Object.entries(details)
            .map(([key, value]) => `${key}: ${value}`)
            .join(', ');

        return `
            <tr>
                <td>${new Date(log.created_at).toLocaleString()}</td>
                <td>
                    <span class="badge ${getOperationBadgeClass(log.operation_type)}">
                        ${getOperationDisplayName(log.operation_type)}
                    </span>
                </td>
                <td>${log.user_id}</td>
                <td><code>${log.ip_address}</code></td>
                <td>
                    <small class="text-muted">${detailsText}</small>
                </td>
            </tr>
        `;
    }).join('');
}

// 获取操作类型的徽章样式
function getOperationBadgeClass(operationType) {
    const badgeMap = {
        'setup': 'bg-primary',
        'keygen': 'bg-success',
        'encrypt': 'bg-warning',
        'decrypt': 'bg-info'
    };
    return badgeMap[operationType] || 'bg-secondary';
}

// 获取操作类型的显示名称
function getOperationDisplayName(operationType) {
    const nameMap = {
        'setup': '系统初始化',
        'keygen': '密钥生成',
        'encrypt': '数据加密',
        'decrypt': '数据解密'
    };
    return nameMap[operationType] || operationType;
}

// 自动填充功能（从系统初始化结果中）
function autoFillFromSetup() {
    if (abeSystemKeys) {
        // 填充密钥生成表单
        const keygenPubkey = document.getElementById('keygen-pubkey');
        const keygenSeckey = document.getElementById('keygen-seckey');
        if (keygenPubkey) keygenPubkey.value = abeSystemKeys.pubKey;
        if (keygenSeckey) keygenSeckey.value = abeSystemKeys.secKey;

        // 填充加密表单
        const encryptPubkey = document.getElementById('encrypt-pubkey');
        if (encryptPubkey) encryptPubkey.value = abeSystemKeys.pubKey;

        // 填充解密表单
        const decryptPubkey = document.getElementById('decrypt-pubkey');
        if (decryptPubkey) decryptPubkey.value = abeSystemKeys.pubKey;

        showSuccess('已自动填充公钥和主密钥信息');
    }
}

// 示例数据填充功能
function fillExampleData() {
    // 根据当前显示的section填充不同的示例数据
    const currentSection = document.querySelector('#abe-platform .section:not(.d-none)');

    if (currentSection && currentSection.id === 'abe-setup-section') {
        // 系统初始化示例
        const abeAttributes = document.getElementById('abe-attributes');
        if (abeAttributes) {
            abeAttributes.value = `doctor
nurse
hospital:301
hospital:302
department:cardiology
department:neurology
role:senior
role:junior
clearance:level1
clearance:level2`;
        }
    } else if (currentSection && currentSection.id === 'abe-upload-section') {
        // 医疗数据上传示例
        const medicalDataText = document.getElementById('medical-data-text');
        if (medicalDataText) {
            medicalDataText.value = `患者基本信息：
姓名：张三
年龄：45岁
性别：男
身份证号：110101197901010001

诊断信息：
主诊断：高血压病2级
次要诊断：糖尿病前期
就诊科室：心内科
主治医生：李医生

治疗方案：
处方药物：
1. 氨氯地平片 5mg 每日一次
2. 阿司匹林肠溶片 100mg 每日一次
3. 阿托伐他汀钙片 20mg 每日一次

生活建议：
1. 低盐低脂饮食
2. 适量运动
3. 戒烟限酒
4. 定期复查

医院：301医院
日期：2024-01-15
医生签名：李医生`;
        }
    } else if (currentSection && currentSection.id === 'abe-keygen-section') {
        // 用户属性示例
        const keygenAttributes = document.getElementById('keygen-attributes');
        if (keygenAttributes) {
            keygenAttributes.value = `doctor
hospital:301
department:cardiology
role:senior`;
        }
    } else if (currentSection && currentSection.id === 'abe-encrypt-section') {
        // 加密策略示例
        const encryptPolicy = document.getElementById('encrypt-policy');
        if (encryptPolicy) {
            encryptPolicy.value = '(doctor AND hospital:301) OR role:admin';
        }

        // 加密消息示例
        const encryptMessage = document.getElementById('encrypt-message');
        if (encryptMessage) {
            if (latestIPFSHash) {
                encryptMessage.value = latestIPFSHash;
            } else {
                encryptMessage.value = 'QmExampleIPFSHashForMedicalData1234567890abcdef';
            }
        }
    }

    showSuccess('已填充示例数据');
}

// 在页面加载完成后添加示例按钮
document.addEventListener('DOMContentLoaded', () => {
    // 为各个表单添加示例按钮
    setTimeout(() => {
        addExampleButtons();
    }, 1000);
});

// 使用最新IPFS哈希
function useLatestIPFSHash() {
    if (latestIPFSHash) {
        document.getElementById('encrypt-message').value = latestIPFSHash;
        showSuccess('已填入最新的IPFS哈希值');
    } else {
        showError('请先上传医疗数据到IPFS');
    }
}

// 使用最新密文
function useLatestCiphertext() {
    if (latestCiphertext) {
        document.getElementById('decrypt-ciphertext').value = latestCiphertext;
        showSuccess('已填入最新的加密密文');
    } else {
        showError('请先进行数据加密');
    }
}

// 使用最新用户密钥
function useLatestUserKey() {
    if (latestUserKey) {
        document.getElementById('decrypt-userkey').value = latestUserKey;
        showSuccess('已填入最新生成的用户密钥');
    } else {
        showError('请先生成用户属性密钥');
    }
}

// 添加示例按钮
function addExampleButtons() {
    // 为ABE系统初始化添加示例按钮
    const setupForm = document.getElementById('abe-setup-form');
    if (setupForm) {
        const exampleBtn = document.createElement('button');
        exampleBtn.type = 'button';
        exampleBtn.className = 'btn btn-outline-light btn-sm me-2';
        exampleBtn.innerHTML = '<i class="bi bi-lightbulb me-1"></i>填充示例';
        exampleBtn.onclick = () => {
            document.getElementById('abe-attributes').value = `age:18-30
department:IT
role:developer
role:manager
clearance:level1
clearance:level2
clearance:level3
location:beijing
location:shanghai`;
        };
        setupForm.appendChild(exampleBtn);
    }

    // 为其他表单也添加类似的示例按钮...
} 