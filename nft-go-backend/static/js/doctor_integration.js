// 医生DID和VC功能整合到主界面

// 全局变量
const API_BASE_URL = window.location.origin + '/api';

document.addEventListener('DOMContentLoaded', () => {
    // 设置医生DID和VC的导航事件
    setupDoctorNavigation();

    // 创建医生DID和VC的界面元素
    createDoctorDIDSection();
    createDoctorVCSection();

    // 初始化事件监听器
    initEventListeners();
});

// 初始化事件监听器
function initEventListeners() {
    // 确保在showPlatform函数中能够正确处理医生DID平台
    const originalShowPlatform = window.showPlatform;
    if (typeof originalShowPlatform === 'function') {
        window.showPlatform = function (platform) {
            originalShowPlatform(platform);

            // 如果选择了DID平台，确保医生DID和VC的导航事件正常工作
            if (platform === 'did') {
                setTimeout(() => {
                    setupDoctorNavigation();
                    console.log('医生DID导航已设置');
                }, 500);
            }
        };
    }

    // 检查是否已经在DID平台
    if (document.getElementById('did-vc-platform') &&
        document.getElementById('did-vc-platform').classList.contains('active')) {
        setupDoctorNavigation();
        console.log('页面加载时医生DID导航已设置');
    }
}

// 设置医生DID和VC的导航
function setupDoctorNavigation() {
    console.log('开始设置医生DID和VC导航事件');

    const navDoctorDID = document.getElementById('nav-doctor-did');
    const navDoctorVC = document.getElementById('nav-doctor-vc');

    console.log('导航元素:', { navDoctorDID, navDoctorVC });

    if (navDoctorDID) {
        // 移除所有现有的事件监听器
        const newNavDoctorDID = navDoctorDID.cloneNode(true);
        navDoctorDID.parentNode.replaceChild(newNavDoctorDID, navDoctorDID);

        // 添加新的事件监听器
        newNavDoctorDID.addEventListener('click', function (e) {
            e.preventDefault();
            console.log('医生DID导航被点击');
            showDoctorDIDSection();
        });

        // 添加直接点击处理
        newNavDoctorDID.onclick = function (e) {
            e.preventDefault();
            console.log('医生DID导航被直接点击');
            showDoctorDIDSection();
        };
    } else {
        console.error('未找到医生DID导航元素');
    }

    if (navDoctorVC) {
        // 移除所有现有的事件监听器
        const newNavDoctorVC = navDoctorVC.cloneNode(true);
        navDoctorVC.parentNode.replaceChild(newNavDoctorVC, navDoctorVC);

        // 添加新的事件监听器
        newNavDoctorVC.addEventListener('click', function (e) {
            e.preventDefault();
            console.log('医生凭证导航被点击');
            showDoctorVCSection();
        });

        // 添加直接点击处理
        newNavDoctorVC.onclick = function (e) {
            e.preventDefault();
            console.log('医生凭证导航被直接点击');
            showDoctorVCSection();
        };
    } else {
        console.error('未找到医生凭证导航元素');
    }

    // 添加全局函数以便直接从HTML调用
    window.showDoctorDIDSection = showDoctorDIDSection;
    window.showDoctorVCSection = showDoctorVCSection;

    console.log('医生DID和VC导航事件设置完成');
}

// 显示医生DID界面
function showDoctorDIDSection() {
    // 隐藏所有其他部分
    document.querySelectorAll('.section').forEach(section => {
        section.classList.add('d-none');
    });

    // 显示医生DID部分
    let doctorDIDSection = document.getElementById('doctor-did-section');
    if (doctorDIDSection) {
        doctorDIDSection.classList.remove('d-none');
    } else {
        createDoctorDIDSection();
        doctorDIDSection = document.getElementById('doctor-did-section');
        if (doctorDIDSection) {
            doctorDIDSection.classList.remove('d-none');
        }
    }

    console.log('显示医生DID界面');
}

// 显示医生VC界面
function showDoctorVCSection() {
    // 隐藏所有其他部分
    document.querySelectorAll('.section').forEach(section => {
        section.classList.add('d-none');
    });

    // 显示医生VC部分
    let doctorVCSection = document.getElementById('doctor-vc-section');
    if (doctorVCSection) {
        doctorVCSection.classList.remove('d-none');
    } else {
        createDoctorVCSection();
        doctorVCSection = document.getElementById('doctor-vc-section');
        if (doctorVCSection) {
            doctorVCSection.classList.remove('d-none');
        }
    }

    console.log('显示医生VC界面');
}

// 创建医生DID界面
function createDoctorDIDSection() {
    const container = document.querySelector('.container');
    if (!container) return;

    // 检查是否已存在
    let doctorDIDSection = document.getElementById('doctor-did-section');
    if (doctorDIDSection) {
        return;
    }

    // 创建医生DID部分
    doctorDIDSection = document.createElement('div');
    doctorDIDSection.id = 'doctor-did-section';
    doctorDIDSection.className = 'section d-none';

    doctorDIDSection.innerHTML = `
        <h2><i class="bi bi-person-badge me-2"></i>医生DID管理</h2>
        <div class="row">
            <div class="col-md-6">
                <div class="card">
                    <div class="card-header">
                        <h5><i class="bi bi-plus-circle me-2"></i>创建医生DID</h5>
                    </div>
                    <div class="card-body">
                        <form id="doctor-did-form">
                            <div class="mb-3">
                                <label for="doctor-name" class="form-label">医生姓名</label>
                                <input type="text" class="form-control" id="doctor-name" required>
                            </div>
                            <div class="mb-3">
                                <label for="license-number" class="form-label">执业证号</label>
                                <input type="text" class="form-control" id="license-number" required>
                            </div>
                            <button type="button" class="btn btn-primary" onclick="createDoctorDID()">
                                <i class="bi bi-person-plus me-1"></i>创建医生DID
                            </button>
                        </form>
                    </div>
                </div>
            </div>
            <div class="col-md-6">
                <div class="card">
                    <div class="card-header">
                        <h5><i class="bi bi-check-circle me-2"></i>创建结果</h5>
                    </div>
                    <div class="card-body">
                        <div id="doctor-did-result">
                            <p>请填写表单并创建医生DID</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    `;

    container.appendChild(doctorDIDSection);
}

// 创建医生VC界面
function createDoctorVCSection() {
    const container = document.querySelector('.container');
    if (!container) return;

    // 检查是否已存在
    let doctorVCSection = document.getElementById('doctor-vc-section');
    if (doctorVCSection) {
        return;
    }

    // 创建医生VC部分
    doctorVCSection = document.createElement('div');
    doctorVCSection.id = 'doctor-vc-section';
    doctorVCSection.className = 'section d-none';

    doctorVCSection.innerHTML = `
        <h2><i class="bi bi-prescription2 me-2"></i>医生凭证管理</h2>
        
        <!-- 颁发医生凭证 -->
        <div class="card mb-4">
            <div class="card-header">
                <h5><i class="bi bi-award me-2"></i>颁发医生凭证</h5>
            </div>
            <div class="card-body">
                <div class="row">
                    <div class="col-md-6">
                        <form id="doctor-vc-form">
                            <div class="mb-3">
                                <label for="issuer-did" class="form-label">颁发者DID</label>
                                <input type="text" class="form-control" id="issuer-did" required>
                                <small class="text-muted">医院DID（固定地址：0x1234...）</small>
                            </div>
                            <div class="mb-3">
                                <label for="doctor-did-input" class="form-label">医生DID</label>
                                <input type="text" class="form-control" id="doctor-did-input" required>
                            </div>
                            <div class="mb-3">
                                <label for="vc-type-select" class="form-label">凭证类型</label>
                                <select class="form-control" id="vc-type-select" required>
                                    <option value="MedicalLicense">医师执业证书</option>
                                    <option value="SpecialistCertificate">专科医师证书</option>
                                    <option value="HospitalAffiliation">医院工作证明</option>
                                </select>
                            </div>
                            <div class="mb-3">
                                <label for="vc-content" class="form-label">凭证内容</label>
                                <textarea class="form-control" id="vc-content" rows="3" required></textarea>
                                <small class="text-muted">JSON格式的凭证内容</small>
                            </div>
                            <button type="button" class="btn btn-primary" onclick="issueDoctorVC()">
                                <i class="bi bi-award me-1"></i>颁发凭证
                            </button>
                        </form>
                    </div>
                    <div class="col-md-6">
                        <div id="issue-vc-result">
                            <p>请填写表单并颁发医生凭证</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        
        <!-- 验证医生凭证 -->
        <div class="card mb-4">
            <div class="card-header">
                <h5><i class="bi bi-shield-check me-2"></i>验证医生凭证</h5>
            </div>
            <div class="card-body">
                <div class="row">
                    <div class="col-md-6">
                        <form id="verify-vc-form">
                            <div class="mb-3">
                                <label for="verify-vc-id" class="form-label">凭证ID</label>
                                <input type="text" class="form-control" id="verify-vc-id" required>
                            </div>
                            <button type="button" class="btn btn-primary" onclick="verifyDoctorVC()">
                                <i class="bi bi-shield-check me-1"></i>验证凭证
                            </button>
                        </form>
                    </div>
                    <div class="col-md-6">
                        <div id="verify-vc-result">
                            <p>请输入凭证ID并验证</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        
        <!-- 查询医生凭证 -->
        <div class="card">
            <div class="card-header">
                <h5><i class="bi bi-search me-2"></i>查询医生凭证</h5>
            </div>
            <div class="card-body">
                <div class="row">
                    <div class="col-md-6">
                        <form id="query-vcs-form">
                            <div class="mb-3">
                                <label for="query-doctor-did" class="form-label">医生DID</label>
                                <input type="text" class="form-control" id="query-doctor-did" required>
                            </div>
                            <button type="button" class="btn btn-primary" onclick="queryDoctorVCs()">
                                <i class="bi bi-search me-1"></i>查询凭证
                            </button>
                        </form>
                    </div>
                    <div class="col-md-6">
                        <div id="query-vcs-result">
                            <p>请输入医生DID并查询</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    `;

    container.appendChild(doctorVCSection);
}

// 显示错误消息
function showError(message) {
    // 检查是否存在全局函数
    if (typeof window.showError === 'function') {
        window.showError(message);
        return;
    }

    // 否则使用alert
    alert('错误: ' + message);
    console.error('错误:', message);
}

// 显示成功消息
function showSuccess(message) {
    // 检查是否存在全局函数
    if (typeof window.showSuccess === 'function') {
        window.showSuccess(message);
        return;
    }

    // 否则使用alert
    alert('成功: ' + message);
    console.log('成功:', message);
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