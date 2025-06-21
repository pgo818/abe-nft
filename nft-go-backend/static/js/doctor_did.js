// 医生DID和VC管理功能

// 确保API_BASE_URL已定义
if (typeof API_BASE_URL === 'undefined') {
    const API_BASE_URL = window.location.origin + '/api';
}

// 创建医生DID
async function createDoctorDID() {
    const doctorName = document.getElementById('doctor-name').value.trim();
    const licenseNumber = document.getElementById('license-number').value.trim();

    if (!doctorName || !licenseNumber) {
        if (typeof showError === 'function') {
            showError('请填写所有必填字段');
        } else {
            alert('请填写所有必填字段');
        }
        return;
    }

    try {
        const response = await fetch(`${API_BASE_URL}/did/doctor/create`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: doctorName,
                licenseNumber: licenseNumber
            })
        });

        const result = await response.json();

        if (response.ok) {
            const resultDiv = document.getElementById('doctor-did-result');
            if (resultDiv) {
                resultDiv.innerHTML = `
                    <div class="alert alert-success">
                        <h6><i class="bi bi-check-circle me-2"></i>医生DID创建成功</h6>
                        <p><strong>医生DID:</strong> <code>${result.did}</code></p>
                        <p><strong>医生姓名:</strong> ${result.name}</p>
                        <p><strong>执业编号:</strong> ${result.licenseNumber}</p>
                        <button class="btn btn-sm btn-outline-primary" onclick="copyToClipboard('${result.did}')">
                            <i class="bi bi-clipboard me-1"></i>复制DID
                        </button>
                    </div>
                `;
            }

            // 清空表单
            const form = document.getElementById('doctor-did-form');
            if (form) {
                form.reset();
            }

            if (typeof showSuccess === 'function') {
                showSuccess('医生DID创建成功');
            } else {
                alert('医生DID创建成功');
            }
        } else {
            if (typeof showError === 'function') {
                showError(result.error || '创建医生DID失败');
            } else {
                alert(result.error || '创建医生DID失败');
            }
        }
    } catch (error) {
        console.error('创建医生DID出错:', error);
        if (typeof showError === 'function') {
            showError('创建医生DID失败: ' + error.message);
        } else {
            alert('创建医生DID失败: ' + error.message);
        }
    }
}

// 颁发医生凭证
async function issueDoctorVC() {
    const issuerDID = document.getElementById('issuer-did').value.trim();
    const doctorDID = document.getElementById('doctor-did-input').value.trim();
    const vcType = document.getElementById('vc-type-select').value;
    const vcContent = document.getElementById('vc-content').value.trim();

    if (!issuerDID || !doctorDID || !vcType || !vcContent) {
        if (typeof showError === 'function') {
            showError('请填写所有必填字段');
        } else {
            alert('请填写所有必填字段');
        }
        return;
    }

    try {
        const response = await fetch(`${API_BASE_URL}/vc/doctor/issue`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                issuerDID: issuerDID,
                doctorDID: doctorDID,
                credentialType: vcType,
                content: vcContent
            })
        });

        const result = await response.json();

        if (response.ok) {
            const resultDiv = document.getElementById('issue-vc-result');
            if (resultDiv) {
                resultDiv.innerHTML = `
                    <div class="alert alert-success">
                        <h6><i class="bi bi-award me-2"></i>医生凭证颁发成功</h6>
                        <p><strong>凭证ID:</strong> <code>${result.credentialId}</code></p>
                        <p><strong>凭证类型:</strong> ${result.credentialType}</p>
                        <p><strong>颁发者:</strong> ${result.issuerDID}</p>
                        <p><strong>医生DID:</strong> ${result.doctorDID}</p>
                        <button class="btn btn-sm btn-outline-primary" onclick="copyToClipboard('${result.credentialId}')">
                            <i class="bi bi-clipboard me-1"></i>复制凭证ID
                        </button>
                    </div>
                `;
            }

            // 清空表单
            const form = document.getElementById('doctor-vc-form');
            if (form) {
                form.reset();
            }

            if (typeof showSuccess === 'function') {
                showSuccess('医生凭证颁发成功');
            } else {
                alert('医生凭证颁发成功');
            }
        } else {
            if (typeof showError === 'function') {
                showError(result.error || '颁发医生凭证失败');
            } else {
                alert(result.error || '颁发医生凭证失败');
            }
        }
    } catch (error) {
        console.error('颁发医生凭证出错:', error);
        if (typeof showError === 'function') {
            showError('颁发医生凭证失败: ' + error.message);
        } else {
            alert('颁发医生凭证失败: ' + error.message);
        }
    }
}

// 验证医生凭证
async function verifyDoctorVC() {
    const vcId = document.getElementById('verify-vc-id').value.trim();

    if (!vcId) {
        if (typeof showError === 'function') {
            showError('请输入凭证ID');
        } else {
            alert('请输入凭证ID');
        }
        return;
    }

    try {
        const response = await fetch(`${API_BASE_URL}/vc/doctor/verify`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                credentialId: vcId
            })
        });

        const result = await response.json();

        if (response.ok) {
            const resultDiv = document.getElementById('verify-vc-result');
            if (resultDiv) {
                const statusClass = result.valid ? 'alert-success' : 'alert-danger';
                const statusIcon = result.valid ? 'bi-shield-check' : 'bi-shield-x';
                const statusText = result.valid ? '验证通过' : '验证失败';

                resultDiv.innerHTML = `
                    <div class="alert ${statusClass}">
                        <h6><i class="${statusIcon} me-2"></i>凭证${statusText}</h6>
                        <p><strong>凭证ID:</strong> ${result.credentialId}</p>
                        <p><strong>验证状态:</strong> ${result.valid ? '有效' : '无效'}</p>
                        ${result.credential ? `
                            <p><strong>凭证类型:</strong> ${result.credential.credentialType}</p>
                            <p><strong>颁发者:</strong> ${result.credential.issuerDID}</p>
                            <p><strong>医生DID:</strong> ${result.credential.doctorDID}</p>
                            <p><strong>颁发时间:</strong> ${new Date(result.credential.issuedAt).toLocaleString()}</p>
                        ` : ''}
                        ${result.error ? `<p><strong>错误信息:</strong> ${result.error}</p>` : ''}
                    </div>
                `;
            }
        } else {
            if (typeof showError === 'function') {
                showError(result.error || '验证凭证失败');
            } else {
                alert(result.error || '验证凭证失败');
            }
        }
    } catch (error) {
        console.error('验证凭证出错:', error);
        if (typeof showError === 'function') {
            showError('验证凭证失败: ' + error.message);
        } else {
            alert('验证凭证失败: ' + error.message);
        }
    }
}

// 查询医生凭证
async function queryDoctorVCs() {
    const doctorDID = document.getElementById('query-doctor-did').value.trim();

    if (!doctorDID) {
        if (typeof showError === 'function') {
            showError('请输入医生DID');
        } else {
            alert('请输入医生DID');
        }
        return;
    }

    try {
        const response = await fetch(`${API_BASE_URL}/vc/doctor/${encodeURIComponent(doctorDID)}`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            }
        });

        const result = await response.json();

        if (response.ok) {
            const resultDiv = document.getElementById('query-vcs-result');
            if (resultDiv) {
                if (result.credentials && result.credentials.length > 0) {
                    const credentialsHtml = result.credentials.map(cred => `
                        <div class="card mb-2">
                            <div class="card-body">
                                <h6 class="card-title">${cred.credentialType}</h6>
                                <p class="card-text"><small class="text-muted">凭证ID: ${cred.id}</small></p>
                                <p class="card-text"><small class="text-muted">颁发者: ${cred.issuerDID}</small></p>
                                <p class="card-text"><small class="text-muted">颁发时间: ${new Date(cred.issuedAt).toLocaleString()}</small></p>
                                <button class="btn btn-sm btn-outline-primary" onclick="copyToClipboard('${cred.id}')">
                                    <i class="bi bi-clipboard me-1"></i>复制ID
                                </button>
                            </div>
                        </div>
                    `).join('');

                    resultDiv.innerHTML = `
                        <div class="alert alert-info">
                            <h6><i class="bi bi-info-circle me-2"></i>找到 ${result.credentials.length} 个凭证</h6>
                        </div>
                        ${credentialsHtml}
                    `;
                } else {
                    resultDiv.innerHTML = `
                        <div class="alert alert-warning">
                            <h6><i class="bi bi-exclamation-triangle me-2"></i>未找到凭证</h6>
                            <p>该医生DID下没有找到任何凭证</p>
                        </div>
                    `;
                }
            }
        } else {
            if (typeof showError === 'function') {
                showError(result.error || '查询凭证失败');
            } else {
                alert(result.error || '查询凭证失败');
            }
        }
    } catch (error) {
        console.error('查询凭证出错:', error);
        if (typeof showError === 'function') {
            showError('查询凭证失败: ' + error.message);
        } else {
            alert('查询凭证失败: ' + error.message);
        }
    }
}

// 复制到剪贴板
function copyToClipboard(text) {
    navigator.clipboard.writeText(text).then(() => {
        if (typeof showSuccess === 'function') {
            showSuccess('已复制到剪贴板');
        } else {
            alert('已复制到剪贴板');
        }
    }).catch(err => {
        console.error('复制失败:', err);
        if (typeof showError === 'function') {
            showError('复制失败');
        } else {
            alert('复制失败');
        }
    });
}
