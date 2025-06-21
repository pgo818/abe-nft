-- NFT+ABE+DID集成平台数据库初始化脚本
-- 版本: 2.1
-- 更新: 添加了医生DID和VC相关表支持

-- 设置字符集和连接参数
SET NAMES utf8mb4;
SET character_set_client = utf8mb4;

-- 创建数据库
CREATE DATABASE IF NOT EXISTS nft_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE nft_db;

-- 删除现有表（如果需要重新创建）
-- SET FOREIGN_KEY_CHECKS = 0;
-- DROP TABLE IF EXISTS verifiable_presentations;
-- DROP TABLE IF EXISTS verifiable_credentials;
-- DROP TABLE IF EXISTS dids;
-- SET FOREIGN_KEY_CHECKS = 1;

-- NFT表
CREATE TABLE IF NOT EXISTS nfts (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    token_id VARCHAR(255) NOT NULL,
    owner_address VARCHAR(255) NOT NULL,
    uri TEXT NOT NULL,
    contract_type VARCHAR(50) NOT NULL DEFAULT 'main',
    parent_token_id VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_token_id (token_id),
    INDEX idx_owner (owner_address),
    INDEX idx_parent (parent_token_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 子NFT申请表
CREATE TABLE IF NOT EXISTS child_nft_requests (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    parent_token_id VARCHAR(255) NOT NULL,
    applicant_address VARCHAR(255) NOT NULL,
    uri TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    child_token_id VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_parent (parent_token_id),
    INDEX idx_applicant (applicant_address),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- NFT元数据表
CREATE TABLE IF NOT EXISTS nft_metadata (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    ipfs_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT NULL,
    external_url TEXT NULL,
    image TEXT NOT NULL,
    policy TEXT NULL,
    ciphertext TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_hash (ipfs_hash)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- DID表
CREATE TABLE IF NOT EXISTS dids (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    did_string VARCHAR(255) NOT NULL COMMENT 'DID标识符',
    wallet_address VARCHAR(255) NOT NULL COMMENT '关联的钱包地址',
    status VARCHAR(50) NOT NULL DEFAULT 'active' COMMENT 'DID状态: active, revoked',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    UNIQUE INDEX idx_did_string (did_string),
    UNIQUE INDEX idx_wallet (wallet_address),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='DID身份标识表';

-- 可验证凭证表
CREATE TABLE IF NOT EXISTS verifiable_credentials (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    credential_id VARCHAR(255) NOT NULL COMMENT '凭证唯一标识',
    issuer_did VARCHAR(255) NOT NULL COMMENT '颁发者DID',
    subject_did VARCHAR(255) NOT NULL COMMENT '主体DID',
    credential_type VARCHAR(255) NOT NULL COMMENT '凭证类型',
    credential_schema TEXT NULL COMMENT '凭证模式',
    status VARCHAR(50) NOT NULL DEFAULT 'active' COMMENT '凭证状态: active, revoked, suspended',
    issuance_date TIMESTAMP NOT NULL COMMENT '颁发日期',
    expiration_date TIMESTAMP NOT NULL COMMENT '过期日期',
    claims TEXT NOT NULL COMMENT '凭证声明（JSON格式）',
    credential_subject TEXT NOT NULL COMMENT '凭证主体（JSON格式）',
    proof TEXT NOT NULL COMMENT '证明（JSON格式）',
    last_verified TIMESTAMP NULL COMMENT '最后验证时间',
    revocation_date TIMESTAMP NULL COMMENT '撤销日期',
    revocation_reason TEXT NULL COMMENT '撤销原因',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    UNIQUE INDEX idx_credential_id (credential_id),
    INDEX idx_issuer (issuer_did),
    INDEX idx_subject (subject_did),
    INDEX idx_status (status),
    INDEX idx_type (credential_type),
    INDEX idx_issuance (issuance_date),
    INDEX idx_expiration (expiration_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='可验证凭证表';

-- 可验证表示表
CREATE TABLE IF NOT EXISTS verifiable_presentations (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    presentation_id VARCHAR(255) NOT NULL COMMENT '表示唯一标识',
    holder_did VARCHAR(255) NOT NULL COMMENT '持有者DID',
    verifier_did VARCHAR(255) NULL COMMENT '验证者DID',
    credential_ids JSON NOT NULL COMMENT '包含的凭证ID列表',
    purpose TEXT NULL COMMENT '展示目的',
    challenge VARCHAR(255) NULL COMMENT '挑战值',
    presentation_date TIMESTAMP NOT NULL COMMENT '展示日期',
    last_verified TIMESTAMP NULL COMMENT '最后验证时间',
    status VARCHAR(50) NOT NULL DEFAULT 'active' COMMENT '状态: active, revoked',
    proof TEXT NOT NULL COMMENT '证明（JSON格式）',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    UNIQUE INDEX idx_presentation_id (presentation_id),
    INDEX idx_holder (holder_did),
    INDEX idx_verifier (verifier_did),
    INDEX idx_status (status),
    INDEX idx_presentation_date (presentation_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='可验证表示表';

-- 凭证模式表
CREATE TABLE IF NOT EXISTS credential_schemas (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    schema_id VARCHAR(255) NOT NULL COMMENT '模式唯一标识',
    name VARCHAR(255) NOT NULL COMMENT '模式名称',
    description TEXT NULL COMMENT '模式描述',
    version VARCHAR(50) NOT NULL COMMENT '模式版本',
    author VARCHAR(255) NULL COMMENT '模式作者',
    schema_json TEXT NOT NULL COMMENT '模式定义（JSON格式）',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    UNIQUE INDEX idx_schema_id (schema_id),
    INDEX idx_name (name),
    INDEX idx_version (version)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='凭证模式表';

-- 凭证定义表
CREATE TABLE IF NOT EXISTS credential_definitions (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    definition_id VARCHAR(255) NOT NULL COMMENT '定义唯一标识',
    schema_id VARCHAR(255) NOT NULL COMMENT '关联的模式ID',
    issuer_did VARCHAR(255) NOT NULL COMMENT '颁发者DID',
    name VARCHAR(255) NOT NULL COMMENT '定义名称',
    version VARCHAR(50) NOT NULL COMMENT '定义版本',
    tag VARCHAR(100) NULL COMMENT '标签',
    definition TEXT NOT NULL COMMENT '定义内容（JSON格式）',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    UNIQUE INDEX idx_definition_id (definition_id),
    INDEX idx_schema_id (schema_id),
    INDEX idx_issuer (issuer_did),
    INDEX idx_name (name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='凭证定义表';

-- ABE系统密钥表
CREATE TABLE IF NOT EXISTS abe_system_keys (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    pub_key TEXT NOT NULL,
    sec_key TEXT NOT NULL,
    attributes TEXT NOT NULL,
    created_by BIGINT UNSIGNED NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_created_by (created_by),
    INDEX idx_expires (expires_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ABE用户密钥表
CREATE TABLE IF NOT EXISTS abe_user_keys (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    system_key_id BIGINT UNSIGNED NOT NULL,
    attrib_keys TEXT NOT NULL,
    attributes TEXT NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_user (user_id),
    INDEX idx_system_key (system_key_id),
    INDEX idx_expires (expires_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ABE密文表
CREATE TABLE IF NOT EXISTS abe_ciphertexts (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    cipher TEXT NOT NULL,
    policy TEXT NOT NULL,
    system_key_id BIGINT UNSIGNED NOT NULL,
    created_by BIGINT UNSIGNED NOT NULL,
    nft_id BIGINT UNSIGNED NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_system_key (system_key_id),
    INDEX idx_created_by (created_by),
    INDEX idx_nft (nft_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ABE操作日志表
CREATE TABLE IF NOT EXISTS abe_operations (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    operation_type VARCHAR(50) NOT NULL,
    details TEXT NULL,
    ip_address VARCHAR(50) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    INDEX idx_user (user_id),
    INDEX idx_operation (operation_type)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 创建默认用户（用于测试）
INSERT IGNORE INTO abe_system_keys (pub_key, sec_key, attributes, created_by, expires_at)
VALUES 
('测试公钥', '测试私钥', '["admin", "user", "guest"]', 1, DATE_ADD(NOW(), INTERVAL 1 YEAR));

-- 医生表
CREATE TABLE IF NOT EXISTS doctors (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    did_string VARCHAR(255) NOT NULL COMMENT '医生DID',
    wallet_address VARCHAR(255) NOT NULL COMMENT '钱包地址',
    name VARCHAR(255) NOT NULL COMMENT '医生姓名',
    license_number VARCHAR(255) NOT NULL COMMENT '执业编号',
    status VARCHAR(50) NOT NULL DEFAULT 'active' COMMENT '状态: active, inactive',
    hospital_did VARCHAR(255) NULL COMMENT '所属医院DID',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    UNIQUE INDEX idx_did_string (did_string),
    UNIQUE INDEX idx_wallet (wallet_address),
    INDEX idx_license (license_number),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='医生表';

-- 医生凭证表
CREATE TABLE IF NOT EXISTS doctor_vcs (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    vcid VARCHAR(255) NOT NULL COMMENT '凭证ID',
    doctor_did VARCHAR(255) NOT NULL COMMENT '医生DID',
    issuer_did VARCHAR(255) NOT NULL COMMENT '颁发者DID (医院)',
    type VARCHAR(255) NOT NULL COMMENT '凭证类型',
    content TEXT NULL COMMENT '凭证内容',
    issued_at TIMESTAMP NOT NULL COMMENT '颁发时间',
    expires_at TIMESTAMP NOT NULL COMMENT '过期时间',
    status VARCHAR(50) NOT NULL DEFAULT 'active' COMMENT '状态: active, revoked',
    revocation_date TIMESTAMP NULL COMMENT '撤销日期',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL,
    UNIQUE INDEX idx_vcid (vcid),
    INDEX idx_doctor (doctor_did),
    INDEX idx_issuer (issuer_did),
    INDEX idx_type (type),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='医生可验证凭证表';

-- 创建固定医院DID（用于测试）
INSERT IGNORE INTO dids (did_string, wallet_address, status)
VALUES ('0x1234', '0x1234', 'active');

-- 检查表是否创建成功
SELECT 
    'NFT+ABE+DID集成平台数据库初始化完成' AS '结果',
    '所有表创建成功，包括DID和VC支持' AS '详情',
    NOW() AS '完成时间'; 