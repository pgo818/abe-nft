// DID和VC功能的JavaScript文件

// 全局变量
let currentDID = null;
let currentCredentials = [];
let currentPresentations = [];

// 初始化DID和VC页面
function initDIDVC() {
    // 加载DID列表
    loadDIDList();

    // 绑定导航事件
    setupDIDVCNavigation();

    // 绑定事件处理程序
    const createDIDForm = document.getElementById('create-did-form');
    const updateDIDForm = document.getElementById('update-did-form');
    const revokeDIDForm = document.getElementById('revoke-did-form');
    const issueVCForm = document.getElementById('issue-vc-form');
    const verifyVCForm = document.getElementById('verify-vc-form');
    const revokeVCForm = document.getElementById('revoke-vc-form');
    const createVPForm = document.getElementById('create-vp-form');
    const verifyVPForm = document.getElementById('verify-vp-form');

    if (createDIDForm) createDIDForm.addEventListener('submit', createDID);
    if (updateDIDForm) updateDIDForm.addEventListener('submit', updateDID);
    if (revokeDIDForm) revokeDIDForm.addEventListener('submit', revokeDID);
    if (issueVCForm) issueVCForm.addEventListener('submit', issueCredential);
    if (verifyVCForm) verifyVCForm.addEventListener('submit', verifyCredential);
    if (revokeVCForm) revokeVCForm.addEventListener('submit', revokeCredential);
    if (createVPForm) createVPForm.addEventListener('submit', createPresentation);
    if (verifyVPForm) verifyVPForm.addEventListener('submit', verifyPresentation);
}

// 设置DID/VC导航
function setupDIDVCNavigation() {
    const navDIDList = document.getElementById('nav-did-list');
    const navDIDCreate = document.getElementById('nav-did-create');
    const navVCIssue = document.getElementById('nav-vc-issue');
    const navVCManage = document.getElementById('nav-vc-manage');
    const navVPCreate = document.getElementById('nav-vp-create');

    if (navDIDList) {
        navDIDList.addEventListener('click', (e) => {
            e.preventDefault();
            showDIDVCSection('did-list-section');
            loadDIDList();
        });
    }
    if (navDIDCreate) {
        navDIDCreate.addEventListener('click', (e) => {
            e.preventDefault();
            showDIDVCSection('did-create-section');
        });
    }
    if (navVCIssue) {
        navVCIssue.addEventListener('click', (e) => {
            e.preventDefault();
            showDIDVCSection('vc-issue-section');
        });
    }
    if (navVCManage) {
        navVCManage.addEventListener('click', (e) => {
            e.preventDefault();
            showDIDVCSection('vc-manage-section');
            if (currentDID) {
                loadCredentials(currentDID);
            }
        });
    }
    if (navVPCreate) {
        navVPCreate.addEventListener('click', (e) => {
            e.preventDefault();
            showDIDVCSection('vp-create-section');
        });
    }
}

// 加载DID列表
async function loadDIDList() {
    try {
        const response = await fetch('/api/did/list');
        if (response.ok) {
            const dids = await response.json();
            const didList = document.getElementById('did-list');
            didList.innerHTML = '';

            if (dids && dids.length > 0) {
                dids.forEach(did => {
                    const card = document.createElement('div');
                    card.className = 'did-card';
                    card.innerHTML = `
                        <div class="did-info">
                            <div class="did-string">${did.didString}</div>
                            <div class="did-meta">
                                <span>方法: ${did.method}</span>
                                <span>状态: ${did.status}</span>
                                ${did.walletAddress ? `<span>钱包: ${did.walletAddress.substring(0, 6)}...${did.walletAddress.substring(38)}</span>` : ''}
                            </div>
                        </div>
                        <div class="did-actions">
                            <button class="btn btn-sm btn-primary" onclick="viewDIDDocument('${did.didString}')">
                                <i class="fas fa-eye"></i> 查看
                            </button>
                            <button class="btn btn-sm btn-success" onclick="selectDID('${did.didString}')">
                                <i class="fas fa-check"></i> 选择
                            </button>
                            ${did.status === 'active' ? `
                                <button class="btn btn-sm btn-warning" onclick="updateDID('${did.didString}')">
                                    <i class="fas fa-edit"></i> 更新
                                </button>
                                <button class="btn btn-sm btn-danger" onclick="revokeDID('${did.didString}')">
                                    <i class="fas fa-times"></i> 撤销
                                </button>
                            ` : ''}
                        </div>
                    `;
                    didList.appendChild(card);
                });
            } else {
                didList.innerHTML = '<p class="text-center text-muted">暂无DID</p>';
            }
        } else {
            console.error('Failed to load DID list');
        }
    } catch (error) {
        console.error('Error loading DID list:', error);
    }
}

// 创建DID
async function createDID(event) {
    event.preventDefault();

    const method = document.getElementById('did-method').value;
    const controllerAddress = document.getElementById('did-controller').value;

    // 如果是ethr方法，可以使用MetaMask钱包地址
    let walletAddress = controllerAddress;
    if (method === 'ethr' && !controllerAddress && window.ethereum) {
        try {
            const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
            walletAddress = accounts[0];
            document.getElementById('did-controller').value = walletAddress;
        } catch (error) {
            showNotification('请先连接MetaMask钱包', 'error');
            return;
        }
    }

    const publicKey = document.getElementById('publicKey').value;

    try {
        showLoading();
        const response = await fetch('/api/did/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'X-Wallet-Address': walletAddress || ''
            },
            body: JSON.stringify({
                method: method,
                controllerAddress: walletAddress,
                publicKey: publicKey
            })
        });

        if (response.ok) {
            const data = await response.json();
            showNotification('DID创建成功: ' + data.did, 'success');
            loadDIDList();

            // 清空表单
            document.getElementById('did-controller').value = '';
            document.getElementById('publicKey').value = '';
        } else {
            const error = await response.json();
            showNotification('创建DID失败: ' + error.error, 'error');
        }
    } catch (error) {
        showNotification('创建DID时发生错误: ' + error.message, 'error');
    } finally {
        hideLoading();
    }
}

// 查看DID
async function viewDID(didString) {
    try {
        showLoading();
        const response = await fetch(`/api/did/resolve/${encodeURIComponent(didString)}`);
        if (!response.ok) {
            throw new Error('获取DID文档失败');
        }

        const data = await response.json();

        // 显示DID文档
        const didDocumentModal = new bootstrap.Modal(document.getElementById('did-document-modal'));
        document.getElementById('did-document-content').textContent = JSON.stringify(data.didDocument, null, 2);
        didDocumentModal.show();
    } catch (error) {
        showError('查看DID失败: ' + error.message);
    } finally {
        hideLoading();
    }
}

// 选择DID
function selectDID(didString) {
    currentDID = didString;
    document.getElementById('selected-did').textContent = didString;
    document.getElementById('selected-did-container').classList.remove('d-none');

    // 更新表单中的DID字段
    document.querySelectorAll('.did-field').forEach(field => {
        field.value = didString;
    });

    // 加载与该DID相关的凭证
    loadCredentials(didString);
}

// 更新DID
async function updateDID(event) {
    event.preventDefault();

    const didString = document.getElementById('update-did-string').value;
    const controllerAddress = document.getElementById('update-did-controller').value;
    const serviceEndpoints = document.getElementById('update-did-service').value;

    if (!didString) {
        showError('请选择要更新的DID');
        return;
    }

    try {
        showLoading();
        const response = await fetch('/api/did/update', {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                did: didString,
                controllerAddress: controllerAddress,
                serviceEndpoints: serviceEndpoints,
            }),
        });

        if (!response.ok) {
            throw new Error('更新DID失败');
        }

        const data = await response.json();
        showSuccess('DID更新成功: ' + data.DID);

        // 清空表单
        document.getElementById('update-did-controller').value = '';
        document.getElementById('update-did-service').value = '';

        // 刷新DID列表
        loadDIDList();
    } catch (error) {
        showError('更新DID失败: ' + error.message);
    } finally {
        hideLoading();
    }
}

// 撤销DID
async function revokeDID(event) {
    event.preventDefault();

    const didString = document.getElementById('revoke-did-string').value;
    const controllerAddress = document.getElementById('revoke-did-controller').value;

    if (!didString || !controllerAddress) {
        showError('请填写所有必填字段');
        return;
    }

    try {
        showLoading();
        const response = await fetch('/api/did/revoke', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                did: didString,
                controllerAddress: controllerAddress,
            }),
        });

        if (!response.ok) {
            throw new Error('撤销DID失败');
        }

        const data = await response.json();
        showSuccess('DID撤销成功: ' + data.DID);

        // 清空表单
        document.getElementById('revoke-did-controller').value = '';

        // 刷新DID列表
        loadDIDList();

        // 如果撤销的是当前选择的DID，清除选择
        if (currentDID === didString) {
            currentDID = null;
            document.getElementById('selected-did').textContent = '';
            document.getElementById('selected-did-container').classList.add('d-none');
        }
    } catch (error) {
        showError('撤销DID失败: ' + error.message);
    } finally {
        hideLoading();
    }
}

// 加载凭证列表
async function loadCredentials(didString) {
    try {
        showLoading();

        // 加载作为颁发者的凭证
        const issuerResponse = await fetch(`/api/vc/credentials?issuer=${encodeURIComponent(didString)}`);
        if (!issuerResponse.ok) {
            throw new Error('加载颁发的凭证失败');
        }

        // 加载作为主体的凭证
        const subjectResponse = await fetch(`/api/vc/credentials?subject=${encodeURIComponent(didString)}`);
        if (!subjectResponse.ok) {
            throw new Error('加载持有的凭证失败');
        }

        const issuerData = await issuerResponse.json();
        const subjectData = await subjectResponse.json();

        // 合并并去重凭证
        const allCredentials = [...(issuerData.credentials || []), ...(subjectData.credentials || [])];
        const uniqueCredentials = [];
        const credentialIds = new Set();

        allCredentials.forEach(credential => {
            if (!credentialIds.has(credential.CredentialID)) {
                credentialIds.add(credential.CredentialID);
                uniqueCredentials.push(credential);
            }
        });

        currentCredentials = uniqueCredentials;

        // 更新凭证列表
        updateCredentialsList(uniqueCredentials);

        // 更新凭证选择下拉框
        updateCredentialSelect(uniqueCredentials);
    } catch (error) {
        showError('加载凭证列表失败: ' + error.message);
    } finally {
        hideLoading();
    }
}

// 更新凭证列表
function updateCredentialsList(credentials) {
    const issuedList = document.getElementById('issued-credentials-list');
    const receivedList = document.getElementById('received-credentials-list');

    issuedList.innerHTML = '';
    receivedList.innerHTML = '';

    if (credentials.length === 0) {
        issuedList.innerHTML = '<tr><td colspan="5" class="text-center">没有找到颁发的凭证</td></tr>';
        receivedList.innerHTML = '<tr><td colspan="5" class="text-center">没有找到持有的凭证</td></tr>';
        return;
    }

    // 分类凭证
    const issuedCredentials = credentials.filter(cred => cred.IssuerDID === currentDID);
    const receivedCredentials = credentials.filter(cred => cred.SubjectDID === currentDID);

    // 更新颁发的凭证列表
    if (issuedCredentials.length === 0) {
        issuedList.innerHTML = '<tr><td colspan="5" class="text-center">没有找到颁发的凭证</td></tr>';
    } else {
        issuedCredentials.forEach(cred => {
            const tr = document.createElement('tr');
            tr.innerHTML = `
                <td>${cred.CredentialID}</td>
                <td>${cred.SubjectDID}</td>
                <td>${cred.CredentialType}</td>
                <td><span class="badge ${cred.CredentialStatus === 'active' ? 'bg-success' : 'bg-danger'}">${cred.CredentialStatus}</span></td>
                <td>
                    <button class="btn btn-sm btn-info view-credential" data-id="${cred.CredentialID}">查看</button>
                    <button class="btn btn-sm btn-danger revoke-credential" data-id="${cred.CredentialID}">撤销</button>
                </td>
            `;
            issuedList.appendChild(tr);
        });
    }

    // 更新持有的凭证列表
    if (receivedCredentials.length === 0) {
        receivedList.innerHTML = '<tr><td colspan="5" class="text-center">没有找到持有的凭证</td></tr>';
    } else {
        receivedCredentials.forEach(cred => {
            const tr = document.createElement('tr');
            tr.innerHTML = `
                <td>${cred.CredentialID}</td>
                <td>${cred.IssuerDID}</td>
                <td>${cred.CredentialType}</td>
                <td><span class="badge ${cred.CredentialStatus === 'active' ? 'bg-success' : 'bg-danger'}">${cred.CredentialStatus}</span></td>
                <td>
                    <button class="btn btn-sm btn-info view-credential" data-id="${cred.CredentialID}">查看</button>
                    <button class="btn btn-sm btn-primary select-credential" data-id="${cred.CredentialID}">选择</button>
                </td>
            `;
            receivedList.appendChild(tr);
        });
    }

    // 绑定查看和选择按钮事件
    document.querySelectorAll('.view-credential').forEach(button => {
        button.addEventListener('click', () => viewCredential(button.dataset.id));
    });

    document.querySelectorAll('.select-credential').forEach(button => {
        button.addEventListener('click', () => selectCredential(button.dataset.id));
    });

    document.querySelectorAll('.revoke-credential').forEach(button => {
        button.addEventListener('click', () => confirmRevokeCredential(button.dataset.id));
    });
}

// 更新凭证选择下拉框
function updateCredentialSelect(credentials) {
    const credentialSelect = document.getElementById('vp-credential-ids');
    credentialSelect.innerHTML = '';

    // 只显示状态为active的凭证
    const activeCredentials = credentials.filter(cred => cred.CredentialStatus === 'active' && cred.SubjectDID === currentDID);

    if (activeCredentials.length === 0) {
        credentialSelect.innerHTML = '<option value="">没有可用的凭证</option>';
        return;
    }

    activeCredentials.forEach(cred => {
        const option = document.createElement('option');
        option.value = cred.CredentialID;
        option.textContent = `${cred.CredentialID} (${cred.CredentialType})`;
        credentialSelect.appendChild(option);
    });
}

// 颁发凭证
async function issueCredential(event) {
    event.preventDefault();

    const issuerDID = document.getElementById('vc-issuer-did').value;
    const subjectDID = document.getElementById('vc-subject-did').value;
    const credentialType = document.getElementById('vc-type').value.split(',').map(type => type.trim());
    const expirationDate = document.getElementById('vc-expiration').value;

    // 获取凭证主体内容
    const nameField = document.getElementById('vc-subject-name').value;
    const ageField = document.getElementById('vc-subject-age').value;
    const idNumberField = document.getElementById('vc-subject-id-number').value;

    if (!issuerDID || !subjectDID || credentialType.length === 0) {
        showError('请填写所有必填字段');
        return;
    }

    // 构建凭证主体
    const credentialSubject = {
        name: nameField,
        age: parseInt(ageField, 10),
        idNumber: idNumberField,
    };

    try {
        showLoading();
        const response = await fetch('/api/vc/issue', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                issuerDid: issuerDID,
                subjectDid: subjectDID,
                credentialType: credentialType,
                expirationDate: expirationDate || undefined,
                credentialSubject: credentialSubject,
            }),
        });

        if (!response.ok) {
            throw new Error('颁发凭证失败');
        }

        const data = await response.json();
        showSuccess('凭证颁发成功: ' + data.credentialId);

        // 清空表单
        document.getElementById('vc-subject-did').value = '';
        document.getElementById('vc-type').value = '';
        document.getElementById('vc-expiration').value = '';
        document.getElementById('vc-subject-name').value = '';
        document.getElementById('vc-subject-age').value = '';
        document.getElementById('vc-subject-id-number').value = '';

        // 刷新凭证列表
        loadCredentials(currentDID);
    } catch (error) {
        showError('颁发凭证失败: ' + error.message);
    } finally {
        hideLoading();
    }
}

// 查看凭证
async function viewCredential(credentialID) {
    try {
        showLoading();
        const response = await fetch(`/api/vc/credential/${encodeURIComponent(credentialID)}`);
        if (!response.ok) {
            throw new Error('获取凭证失败');
        }

        const credential = await response.json();

        // 显示凭证内容
        const credentialModal = new bootstrap.Modal(document.getElementById('credential-modal'));
        document.getElementById('credential-content').textContent = JSON.stringify(credential, null, 2);
        credentialModal.show();
    } catch (error) {
        showError('查看凭证失败: ' + error.message);
    } finally {
        hideLoading();
    }
}

// 选择凭证
function selectCredential(credentialID) {
    // 查找选择的凭证
    const credential = currentCredentials.find(cred => cred.CredentialID === credentialID);
    if (!credential) {
        showError('找不到选择的凭证');
        return;
    }

    // 更新表单中的凭证ID字段
    document.getElementById('revoke-vc-id').value = credentialID;

    // 如果是创建表示表单，选中对应的复选框
    const checkbox = document.querySelector(`#vp-credential-ids option[value="${credentialID}"]`);
    if (checkbox) {
        checkbox.selected = true;
    }

    showSuccess('已选择凭证: ' + credentialID);
}

// 确认撤销凭证
function confirmRevokeCredential(credentialID) {
    // 查找要撤销的凭证
    const credential = currentCredentials.find(cred => cred.CredentialID === credentialID);
    if (!credential) {
        showError('找不到要撤销的凭证');
        return;
    }

    // 更新撤销表单
    document.getElementById('revoke-vc-id').value = credentialID;
    document.getElementById('revoke-vc-issuer').value = credential.IssuerDID;

    // 显示确认对话框
    const confirmModal = new bootstrap.Modal(document.getElementById('revoke-confirm-modal'));
    document.getElementById('revoke-credential-id').textContent = credentialID;
    confirmModal.show();

    // 绑定确认按钮事件
    document.getElementById('confirm-revoke-btn').onclick = () => {
        confirmModal.hide();
        revokeCredential(null, credentialID, credential.IssuerDID);
    };
}

// 撤销凭证
async function revokeCredential(event, credentialID, issuerDID) {
    if (event) {
        event.preventDefault();
    }

    // 如果没有传入参数，从表单中获取
    if (!credentialID) {
        credentialID = document.getElementById('revoke-vc-id').value;
    }
    if (!issuerDID) {
        issuerDID = document.getElementById('revoke-vc-issuer').value;
    }

    if (!credentialID || !issuerDID) {
        showError('请填写所有必填字段');
        return;
    }

    try {
        showLoading();
        const response = await fetch('/api/vc/revoke', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                credentialId: credentialID,
                issuerDid: issuerDID,
            }),
        });

        if (!response.ok) {
            throw new Error('撤销凭证失败');
        }

        const data = await response.json();
        showSuccess('凭证撤销成功: ' + data.credentialId);

        // 清空表单
        document.getElementById('revoke-vc-id').value = '';
        document.getElementById('revoke-vc-issuer').value = '';

        // 刷新凭证列表
        loadCredentials(currentDID);
    } catch (error) {
        showError('撤销凭证失败: ' + error.message);
    } finally {
        hideLoading();
    }
}

// 验证凭证
async function verifyCredential(event) {
    event.preventDefault();

    const credentialJSON = document.getElementById('verify-vc-json').value;

    if (!credentialJSON) {
        showError('请输入要验证的凭证JSON');
        return;
    }

    try {
        // 解析凭证JSON
        const credential = JSON.parse(credentialJSON);

        showLoading();
        const response = await fetch('/api/vc/verify', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                credential: credential,
            }),
        });

        if (!response.ok) {
            throw new Error('验证凭证失败');
        }

        const data = await response.json();

        if (data.valid) {
            showSuccess('凭证验证成功');
        } else {
            showError('凭证验证失败: ' + data.reason);
        }
    } catch (error) {
        showError('验证凭证失败: ' + error.message);
    } finally {
        hideLoading();
    }
}

// 创建可验证表示
async function createPresentation(event) {
    event.preventDefault();

    const holderDID = document.getElementById('vp-holder-did').value;
    const verifierDID = document.getElementById('vp-verifier-did').value;

    // 获取选中的凭证ID
    const select = document.getElementById('vp-credential-ids');
    const selectedCredentials = Array.from(select.selectedOptions).map(option => option.value);

    if (!holderDID || selectedCredentials.length === 0) {
        showError('请填写所有必填字段并选择至少一个凭证');
        return;
    }

    try {
        showLoading();
        const response = await fetch('/api/vc/presentation/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                holderDid: holderDID,
                verifierDid: verifierDID || undefined,
                credentialIds: selectedCredentials,
            }),
        });

        if (!response.ok) {
            throw new Error('创建表示失败');
        }

        const data = await response.json();

        // 显示创建的表示
        const presentationModal = new bootstrap.Modal(document.getElementById('presentation-modal'));
        document.getElementById('presentation-content').textContent = JSON.stringify(data.presentation, null, 2);
        presentationModal.show();

        // 清空表单
        document.getElementById('vp-verifier-did').value = '';
        Array.from(select.options).forEach(option => option.selected = false);

        // 保存创建的表示到验证表单
        document.getElementById('verify-vp-json').value = JSON.stringify(data.presentation, null, 2);

        showSuccess('表示创建成功: ' + data.presentationId);
    } catch (error) {
        showError('创建表示失败: ' + error.message);
    } finally {
        hideLoading();
    }
}

// 验证可验证表示
async function verifyPresentation(event) {
    event.preventDefault();

    const presentationJSON = document.getElementById('verify-vp-json').value;

    if (!presentationJSON) {
        showError('请输入要验证的表示JSON');
        return;
    }

    try {
        // 解析表示JSON
        const presentation = JSON.parse(presentationJSON);

        showLoading();
        const response = await fetch('/api/vc/presentation/verify', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                presentation: presentation,
            }),
        });

        if (!response.ok) {
            throw new Error('验证表示失败');
        }

        const data = await response.json();

        if (data.valid) {
            showSuccess('表示验证成功');
        } else {
            showError('表示验证失败: ' + data.reason);
        }
    } catch (error) {
        showError('验证表示失败: ' + error.message);
    } finally {
        hideLoading();
    }
}

// 显示加载指示器
function showLoading() {
    document.getElementById('loading-overlay').classList.remove('d-none');
}

// 隐藏加载指示器
function hideLoading() {
    document.getElementById('loading-overlay').classList.add('d-none');
}

// 显示成功消息
function showSuccess(message) {
    const resultModal = new bootstrap.Modal(document.getElementById('result-modal'));
    document.getElementById('result-title').textContent = '成功';
    document.getElementById('result-title').className = 'modal-title text-success';
    document.getElementById('result-message').textContent = message;
    resultModal.show();
}

// 显示错误消息
function showError(message) {
    const resultModal = new bootstrap.Modal(document.getElementById('result-modal'));
    document.getElementById('result-title').textContent = '错误';
    document.getElementById('result-title').className = 'modal-title text-danger';
    document.getElementById('result-message').textContent = message;
    resultModal.show();
}

// 页面加载完成后初始化
document.addEventListener('DOMContentLoaded', function () {
    // 检查是否在DID/VC页面
    if (document.getElementById('did-vc-platform')) {
        initDIDVC();
    }
});

// 医生DID和VC管理模块
document.addEventListener('DOMContentLoaded', function () {
    // 初始化事件监听
    initDoctorDIDListeners();
});

// 初始化事件监听器
function initDoctorDIDListeners() {
    try {
        // 切换tab栏
        const tabButtons = document.querySelectorAll('.tab-button');
        tabButtons.forEach(button => {
            button.addEventListener('click', function () {
                // 移除所有激活状态
                tabButtons.forEach(btn => btn.classList.remove('active'));
                document.querySelectorAll('.tab-content').forEach(content => {
                    content.classList.remove('active');
                });

                // 激活当前tab
                this.classList.add('active');
                const tabId = this.getAttribute('data-tab');
                const targetTab = document.getElementById(tabId);
                if (targetTab) {
                    targetTab.classList.add('active');
                }
            });
        });

        // 安全地绑定事件监听器
        const connectWalletBtn = document.getElementById('connect-wallet');
        if (connectWalletBtn) {
            connectWalletBtn.addEventListener('click', connectWallet);
        }

        const createDoctorDIDBtn = document.getElementById('create-doctor-did');
        if (createDoctorDIDBtn) {
            createDoctorDIDBtn.addEventListener('click', createDoctorDID);
        }

        const issueDoctorVCBtn = document.getElementById('issue-doctor-vc');
        if (issueDoctorVCBtn) {
            issueDoctorVCBtn.addEventListener('click', issueDoctorVC);
        }

        const verifyDoctorVCBtn = document.getElementById('verify-doctor-vc');
        if (verifyDoctorVCBtn) {
            verifyDoctorVCBtn.addEventListener('click', verifyDoctorVC);
        }

        const queryDoctorVCsBtn = document.getElementById('query-doctor-vcs');
        if (queryDoctorVCsBtn) {
            queryDoctorVCsBtn.addEventListener('click', getDoctorVCs);
        }

        // 断开钱包连接按钮
        const disconnectWalletBtn = document.getElementById('disconnect-wallet');
        if (disconnectWalletBtn) {
            disconnectWalletBtn.addEventListener('click', disconnectWallet);
        }

        // 监听MetaMask账户变更
        if (window.ethereum) {
            window.ethereum.on('accountsChanged', function (accounts) {
                if (accounts.length === 0) {
                    // 用户断开了所有账户
                    disconnectWallet();
                } else if (accounts[0] !== currentAccount) {
                    // 用户切换了账户
                    currentAccount = accounts[0];
                    updateWalletUI(accounts[0]);
                    localStorage.setItem('walletAccount', accounts[0]);
                }
            });

            window.ethereum.on('chainChanged', function (chainId) {
                // 网络变更时重新加载页面
                window.location.reload();
            });
        }

        // 页面加载时检查钱包连接状态
        checkWalletConnection();

    } catch (error) {
        console.error('初始化事件监听器失败:', error);
    }
}

// 全局变量
let currentAccount = null;
let isWalletConnected = false;

// 连接钱包
async function connectWallet() {
    console.log('开始连接钱包...');

    if (window.ethereum) {
        try {
            console.log('检测到MetaMask，请求账户访问...');
            // 请求账户访问
            const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
            const account = accounts[0];

            console.log('获取到账户:', account);

            if (account) {
                currentAccount = account;
                isWalletConnected = true;
                updateWalletUI(account);
                showSuccessMessage('钱包连接成功');

                // 保存连接状态到localStorage
                localStorage.setItem('walletConnected', 'true');
                localStorage.setItem('walletAccount', account);

                console.log('钱包连接成功，状态已保存');
            }

            return account;
        } catch (error) {
            console.error('连接钱包失败:', error);
            showErrorMessage('连接钱包失败: ' + error.message);
        }
    } else {
        console.error('未检测到MetaMask');
        showErrorMessage('请安装MetaMask钱包');
    }
    return null;
}

// 断开钱包连接
function disconnectWallet() {
    currentAccount = null;
    isWalletConnected = false;
    updateWalletUI(null);

    // 清除localStorage
    localStorage.removeItem('walletConnected');
    localStorage.removeItem('walletAccount');

    showSuccessMessage('钱包已断开连接');
}

// 更新钱包UI状态
function updateWalletUI(account) {
    console.log('更新钱包UI状态，账户:', account);

    const disconnectedDiv = document.getElementById('wallet-disconnected');
    const connectedDiv = document.getElementById('wallet-connected');
    const walletAddressElement = document.getElementById('wallet-address');

    console.log('DOM元素检查:', {
        disconnectedDiv: !!disconnectedDiv,
        connectedDiv: !!connectedDiv,
        walletAddressElement: !!walletAddressElement
    });

    if (account && disconnectedDiv && connectedDiv && walletAddressElement) {
        // 显示已连接状态
        console.log('显示已连接状态');
        disconnectedDiv.style.display = 'none';
        connectedDiv.style.display = 'flex';
        walletAddressElement.textContent = shortenAddress(account);
    } else if (disconnectedDiv && connectedDiv) {
        // 显示未连接状态
        console.log('显示未连接状态');
        disconnectedDiv.style.display = 'flex';
        connectedDiv.style.display = 'none';
        if (walletAddressElement) {
            walletAddressElement.textContent = '未连接';
        }
    } else {
        console.error('缺少必要的DOM元素');
    }
}

// 检查钱包连接状态
async function checkWalletConnection() {
    if (window.ethereum) {
        try {
            const accounts = await window.ethereum.request({ method: 'eth_accounts' });
            const savedAccount = localStorage.getItem('walletAccount');
            const wasConnected = localStorage.getItem('walletConnected') === 'true';

            if (accounts.length > 0 && wasConnected && savedAccount === accounts[0]) {
                currentAccount = accounts[0];
                isWalletConnected = true;
                updateWalletUI(accounts[0]);
                console.log('自动恢复钱包连接:', accounts[0]);
            } else if (savedAccount && !accounts.includes(savedAccount)) {
                // 清除过期的连接状态
                localStorage.removeItem('walletConnected');
                localStorage.removeItem('walletAccount');
            }
        } catch (error) {
            console.error('检查钱包连接失败:', error);
        }
    }
}

// 创建医生DID
async function createDoctorDID() {
    try {
        showLoading();

        // 获取表单数据
        const name = document.getElementById('doctor-name').value;
        const licenseNumber = document.getElementById('license-number').value;

        // 检查钱包是否已连接
        let walletAddress;
        if (window.ethereum) {
            const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
            walletAddress = accounts[0];
        } else {
            throw new Error('请先连接MetaMask钱包');
        }

        // 验证表单
        if (!name || !licenseNumber) {
            throw new Error('请输入姓名和执业编号');
        }

        // 准备请求数据
        const requestData = {
            walletAddress: walletAddress,
            name: name,
            licenseNumber: licenseNumber
        };

        // 发送请求创建DID
        const response = await fetch('/api/did/doctor/create', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestData)
        });

        // 处理响应
        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || '创建DID失败');
        }

        const result = await response.json();

        // 显示结果
        const resultBox = document.getElementById('doctor-did-result');
        resultBox.innerHTML = `
            <div class="success-box">
                <h4>医生DID创建成功</h4>
                <p><strong>DID:</strong> ${result.did}</p>
                <p><strong>钱包地址:</strong> ${result.walletAddress}</p>
                <p><strong>姓名:</strong> ${result.name}</p>
                <p><strong>执业编号:</strong> ${result.licenseNumber}</p>
            </div>
        `;

        // 更新医生DID输入字段
        document.getElementById('doctor-did-input').value = result.did;

        showSuccessMessage('医生DID创建成功');
    } catch (error) {
        console.error('创建医生DID错误:', error);
        showErrorMessage('创建医生DID失败: ' + error.message);
        document.getElementById('doctor-did-result').innerHTML = `
            <div class="error-box">
                <p>创建失败: ${error.message}</p>
            </div>
        `;
    } finally {
        hideLoading();
    }
}

// 颁发医生凭证
async function issueDoctorVC() {
    try {
        showLoading();

        // 获取表单数据
        const issuerDID = document.getElementById('issuer-did').value;
        const doctorDID = document.getElementById('doctor-did-input').value;
        const vcType = document.getElementById('vc-type').value;
        const vcContent = document.getElementById('vc-content').value;

        // 验证表单
        if (!issuerDID || !doctorDID || !vcType || !vcContent) {
            throw new Error('请填写所有必填字段');
        }

        // 准备请求数据
        const requestData = {
            issuerDid: issuerDID,
            doctorDid: doctorDID,
            vcType: vcType,
            vcContent: vcContent
        };

        // 发送请求颁发凭证
        const response = await fetch('/api/vc/doctor/issue', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestData)
        });

        // 处理响应
        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || '颁发凭证失败');
        }

        const result = await response.json();

        // 显示结果
        const resultBox = document.getElementById('issue-vc-result');
        resultBox.innerHTML = `
            <div class="success-box">
                <h4>医生凭证颁发成功</h4>
                <p><strong>凭证ID:</strong> ${result.vcId}</p>
                <p><strong>医生DID:</strong> ${result.doctorDid}</p>
                <p><strong>颁发者DID:</strong> ${result.issuerDid}</p>
                <p><strong>凭证类型:</strong> ${result.vcType}</p>
                <p><strong>颁发时间:</strong> ${new Date(result.issuedAt).toLocaleString()}</p>
                <p><strong>过期时间:</strong> ${new Date(result.expiresAt).toLocaleString()}</p>
            </div>
        `;

        // 更新验证凭证ID字段
        document.getElementById('verify-vc-id').value = result.vcId;

        showSuccessMessage('医生凭证颁发成功');
    } catch (error) {
        console.error('颁发医生凭证错误:', error);
        showErrorMessage('颁发医生凭证失败: ' + error.message);
        document.getElementById('issue-vc-result').innerHTML = `
            <div class="error-box">
                <p>颁发失败: ${error.message}</p>
            </div>
        `;
    } finally {
        hideLoading();
    }
}

// 验证医生凭证
async function verifyDoctorVC() {
    try {
        showLoading();

        // 获取表单数据
        const vcId = document.getElementById('verify-vc-id').value;

        // 验证表单
        if (!vcId) {
            throw new Error('请输入凭证ID');
        }

        // 准备请求数据
        const requestData = {
            vcId: vcId
        };

        // 发送请求验证凭证
        const response = await fetch('/api/vc/doctor/verify', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestData)
        });

        // 处理响应
        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || '验证凭证失败');
        }

        const result = await response.json();

        // 显示结果
        const resultBox = document.getElementById('verify-vc-result');
        if (result.valid) {
            resultBox.innerHTML = `
                <div class="success-box">
                    <h4>凭证验证成功</h4>
                    <p><strong>医生DID:</strong> ${result.doctorDid}</p>
                    <p><strong>颁发者DID:</strong> ${result.issuerDid}</p>
                    <p><strong>凭证类型:</strong> ${result.vcType}</p>
                    <p><strong>验证结果:</strong> <span class="valid-badge">有效</span></p>
                </div>
            `;
            showSuccessMessage('凭证验证成功');
        } else {
            resultBox.innerHTML = `
                <div class="warning-box">
                    <h4>凭证验证失败</h4>
                    <p><strong>原因:</strong> ${result.reason}</p>
                    <p><strong>验证结果:</strong> <span class="invalid-badge">无效</span></p>
                </div>
            `;
            showErrorMessage('凭证验证失败: ' + result.reason);
        }
    } catch (error) {
        console.error('验证医生凭证错误:', error);
        showErrorMessage('验证医生凭证失败: ' + error.message);
        document.getElementById('verify-vc-result').innerHTML = `
            <div class="error-box">
                <p>验证失败: ${error.message}</p>
            </div>
        `;
    } finally {
        hideLoading();
    }
}

// 获取医生凭证列表
async function getDoctorVCs() {
    try {
        showLoading();

        // 获取表单数据
        const doctorDID = document.getElementById('query-doctor-did').value;

        // 验证表单
        if (!doctorDID) {
            throw new Error('请输入医生DID');
        }

        // 准备请求数据
        const requestData = {
            doctorDid: doctorDID
        };

        // 发送请求获取凭证列表
        const response = await fetch('/api/vc/doctor/list', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestData)
        });

        // 处理响应
        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || '获取凭证列表失败');
        }

        const result = await response.json();

        // 显示结果
        const resultBox = document.getElementById('query-vcs-result');
        if (result.verifiableCredentials && result.verifiableCredentials.length > 0) {
            const vcsHtml = result.verifiableCredentials.map(vc => `
                <div class="vc-card">
                    <p><strong>凭证ID:</strong> ${vc.vcId}</p>
                    <p><strong>颁发者DID:</strong> ${vc.issuerDID}</p>
                    <p><strong>凭证类型:</strong> ${vc.type}</p>
                    <p><strong>颁发时间:</strong> ${new Date(vc.issuedAt).toLocaleString()}</p>
                    <p><strong>状态:</strong> ${vc.status === 'active' ? '<span class="valid-badge">有效</span>' : '<span class="invalid-badge">已撤销</span>'}</p>
                </div>
            `).join('');

            resultBox.innerHTML = `
                <div>
                    <h4>医生凭证列表 (共${result.verifiableCredentials.length}个)</h4>
                    <div class="vc-list">
                        ${vcsHtml}
                    </div>
                </div>
            `;
            showSuccessMessage(`成功获取${result.verifiableCredentials.length}个凭证`);
        } else {
            resultBox.innerHTML = `
                <div class="info-box">
                    <p>未找到医生凭证</p>
                </div>
            `;
            showInfoMessage('未找到医生凭证');
        }
    } catch (error) {
        console.error('获取医生凭证列表错误:', error);
        showErrorMessage('获取医生凭证列表失败: ' + error.message);
        document.getElementById('query-vcs-result').innerHTML = `
            <div class="error-box">
                <p>获取凭证失败: ${error.message}</p>
            </div>
        `;
    } finally {
        hideLoading();
    }
}

// 工具函数 - 显示加载中
function showLoading() {
    // 如果有loading元素，可以在这里显示
    console.log('Loading...');
}

// 工具函数 - 隐藏加载中
function hideLoading() {
    // 如果有loading元素，可以在这里隐藏
    console.log('Loading complete');
}

// 工具函数 - 显示成功消息
function showSuccessMessage(message) {
    alert('成功: ' + message);
}

// 工具函数 - 显示错误消息
function showErrorMessage(message) {
    alert('错误: ' + message);
}

// 工具函数 - 显示信息消息
function showInfoMessage(message) {
    alert('提示: ' + message);
}

// 工具函数 - 缩短地址显示
function shortenAddress(address) {
    if (!address || typeof address !== 'string' || address.length < 10) {
        return '未连接';
    }
    return address.substring(0, 6) + '...' + address.substring(address.length - 4);
}

// 页面加载完成后初始化
document.addEventListener('DOMContentLoaded', function () {
    console.log('页面加载完成，开始初始化...');

    // 检查是否是医生DID页面
    if (document.getElementById('doctor-did') || document.getElementById('connect-wallet')) {
        console.log('检测到医生DID页面，初始化医生DID功能...');
        initDoctorDIDListeners();
    } else {
        console.log('检测到通用DID/VC页面，初始化通用功能...');
        initDIDVC();
    }
});