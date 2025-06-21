package models

import (
	"time"

	"gorm.io/gorm"
)

// DID 表示去中心化身份标识符的数据库模型
type DID struct {
	gorm.Model
	DIDString     string `json:"didString" gorm:"unique;not null"` // 完整的DID字符串
	WalletAddress string `json:"walletAddress" gorm:"unique"`      // 关联的钱包地址
	Status        string `json:"status" gorm:"default:active"`     // DID状态：active, revoked
}

// TableName 指定表名
func (DID) TableName() string {
	return "dids"
}

// Doctor 医生身份表示数据库模型
type Doctor struct {
	gorm.Model
	DIDString     string `json:"didString" gorm:"unique;not null"` // 医生DID
	WalletAddress string `json:"walletAddress" gorm:"unique"`      // 钱包地址
	Name          string `json:"name"`                             // 医生姓名
	LicenseNumber string `json:"licenseNumber"`                    // 执业编号
	Status        string `json:"status" gorm:"default:active"`     // 状态
	HospitalDID   string `json:"hospitalDID"`                      // 所属医院DID
}

// TableName 指定表名
func (Doctor) TableName() string {
	return "doctors"
}

// DoctorVC 医生可验证凭证数据库模型
type DoctorVC struct {
	gorm.Model
	VCID           string     `json:"vcId" gorm:"unique;not null"`  // 凭证ID
	DoctorDID      string     `json:"doctorDID" gorm:"not null"`    // 医生DID
	IssuerDID      string     `json:"issuerDID" gorm:"not null"`    // 颁发者DID (医院)
	Type           string     `json:"type"`                         // 凭证类型，如"执业资格"
	Content        string     `json:"content" gorm:"type:text"`     // 凭证内容
	IssuedAt       time.Time  `json:"issuedAt"`                     // 颁发时间
	ExpiresAt      time.Time  `json:"expiresAt"`                    // 过期时间
	Status         string     `json:"status" gorm:"default:active"` // 状态：active, revoked
	RevocationDate *time.Time `json:"revocationDate"`               // 撤销日期
}

// TableName 指定表名
func (DoctorVC) TableName() string {
	return "doctor_vcs"
}

// DIDDocument 表示DID文档的结构
type DIDDocument struct {
	Context              []string             `json:"@context"`
	ID                   string               `json:"id"`
	Controller           []string             `json:"controller,omitempty"`
	VerificationMethod   []VerificationMethod `json:"verificationMethod"`
	Authentication       []string             `json:"authentication"`
	AssertionMethod      []string             `json:"assertionMethod,omitempty"`
	KeyAgreement         []string             `json:"keyAgreement,omitempty"`
	CapabilityInvocation []string             `json:"capabilityInvocation,omitempty"`
	CapabilityDelegation []string             `json:"capabilityDelegation,omitempty"`
	Service              []Service            `json:"service,omitempty"`
	Created              string               `json:"created"`
	Updated              string               `json:"updated,omitempty"`
}

// VerificationMethod 表示DID文档中的验证方法
type VerificationMethod struct {
	ID              string                 `json:"id"`
	Type            string                 `json:"type"`
	Controller      string                 `json:"controller"`
	PublicKeyJwk    map[string]interface{} `json:"publicKeyJwk,omitempty"`
	PublicKeyBase58 string                 `json:"publicKeyBase58,omitempty"`
	PublicKeyHex    string                 `json:"publicKeyHex,omitempty"`
}

// Service 表示DID文档中的服务端点
type Service struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

// VerifiableCredential 表示可验证凭证的数据库模型
type VerifiableCredential struct {
	gorm.Model
	CredentialID      string     `json:"credentialId" gorm:"unique;not null"` // 凭证ID
	IssuerDID         string     `json:"issuerDid" gorm:"not null"`           // 颁发者DID
	SubjectDID        string     `json:"subjectDid" gorm:"not null"`          // 主体DID
	Type              string     `json:"type" gorm:"column:credential_type"`  // 凭证类型（避免与关键字冲突）
	CredentialSchema  string     `json:"credentialSchema"`                    // 凭证模式
	Status            string     `json:"status" gorm:"default:active"`        // 凭证状态：active, revoked, suspended
	IssuanceDate      time.Time  `json:"issuanceDate" gorm:"not null"`        // 颁发日期
	ExpirationDate    time.Time  `json:"expirationDate"`                      // 过期日期
	Claims            string     `json:"claims" gorm:"type:text"`             // 凭证声明（JSON格式）
	CredentialSubject string     `json:"credentialSubject" gorm:"type:text"`  // 凭证主体（JSON格式）
	Proof             string     `json:"proof" gorm:"type:text"`              // 证明（JSON格式）
	LastVerified      *time.Time `json:"lastVerified"`                        // 最后验证时间
	RevocationDate    *time.Time `json:"revocationDate"`                      // 撤销日期
	RevocationReason  string     `json:"revocationReason"`                    // 撤销原因
}

// TableName 指定表名
func (VerifiableCredential) TableName() string {
	return "verifiable_credentials"
}

// CredentialSchema 表示凭证模式的数据库模型
type CredentialSchema struct {
	gorm.Model
	SchemaID    string `json:"schemaId" gorm:"unique;not null"`      // 模式ID
	Name        string `json:"name" gorm:"not null"`                 // 模式名称
	Description string `json:"description"`                          // 模式描述
	Version     string `json:"version" gorm:"not null"`              // 模式版本
	Author      string `json:"author"`                               // 作者
	SchemaJSON  string `json:"schemaJson" gorm:"type:text;not null"` // 模式定义（JSON格式）
}

// TableName 指定表名
func (CredentialSchema) TableName() string {
	return "credential_schemas"
}

// CredentialDefinition 表示凭证定义的数据库模型
type CredentialDefinition struct {
	gorm.Model
	DefinitionID string `json:"definitionId" gorm:"unique;not null"`  // 定义ID
	SchemaID     string `json:"schemaId" gorm:"not null"`             // 关联的模式ID
	IssuerDID    string `json:"issuerDid" gorm:"not null"`            // 颁发者DID
	Name         string `json:"name" gorm:"not null"`                 // 定义名称
	Version      string `json:"version" gorm:"not null"`              // 定义版本
	Tag          string `json:"tag"`                                  // 标签
	Definition   string `json:"definition" gorm:"type:text;not null"` // 定义内容（JSON格式）
}

// TableName 指定表名
func (CredentialDefinition) TableName() string {
	return "credential_definitions"
}

// VerifiablePresentation 表示可验证表示的数据库模型
type VerifiablePresentation struct {
	gorm.Model
	PresentationID   string     `json:"presentationId" gorm:"unique;not null"` // 表示ID
	HolderDID        string     `json:"holderDid" gorm:"not null"`             // 持有者DID
	VerifierDID      string     `json:"verifierDid"`                           // 验证者DID
	CredentialIDs    []string   `json:"credentialIds" gorm:"serializer:json"`  // 包含的凭证ID列表
	Purpose          string     `json:"purpose"`                               // 展示目的
	Challenge        string     `json:"challenge"`                             // 挑战值
	PresentationDate time.Time  `json:"presentationDate"`                      // 展示日期
	LastVerified     *time.Time `json:"lastVerified"`                          // 最后验证时间
	Status           string     `json:"status" gorm:"default:active"`          // 状态
	Proof            string     `json:"proof" gorm:"type:text"`                // 证明（JSON格式）
}

// TableName 指定表名
func (VerifiablePresentation) TableName() string {
	return "verifiable_presentations"
}

// DIDResolutionRequest 表示DID解析请求
type DIDResolutionRequest struct {
	DID string `json:"did" binding:"required"` // 要解析的DID
}

// DIDResolutionResponse 表示DID解析响应
type DIDResolutionResponse struct {
	DIDDocument DIDDocument `json:"didDocument"` // DID文档
	Status      string      `json:"status"`      // 解析状态
}

// CreateDIDRequest 表示创建DID的请求
type CreateDIDRequest struct {
	Method            string `json:"method" binding:"required"`            // DID方法
	ControllerAddress string `json:"controllerAddress" binding:"required"` // 控制者地址
	PublicKey         string `json:"publicKey"`                            // 公钥（可选）
}

// CreateDIDResponse 表示创建DID的响应
type CreateDIDResponse struct {
	DID         string      `json:"did"`         // 创建的DID
	DIDDocument DIDDocument `json:"didDocument"` // DID文档
}

// UpdateDIDRequest 表示更新DID的请求
type UpdateDIDRequest struct {
	DID               string `json:"did" binding:"required"` // 要更新的DID
	ControllerAddress string `json:"controllerAddress"`      // 新的控制者地址
	PublicKey         string `json:"publicKey"`              // 新的公钥
	ServiceEndpoints  string `json:"serviceEndpoints"`       // 新的服务端点
}

// UpdateDIDResponse 表示更新DID的响应
type UpdateDIDResponse struct {
	DID         string      `json:"did"`         // 更新的DID
	DIDDocument DIDDocument `json:"didDocument"` // 更新后的DID文档
}

// RevokeDIDRequest 表示撤销DID的请求
type RevokeDIDRequest struct {
	DID               string `json:"did" binding:"required"`               // 要撤销的DID
	ControllerAddress string `json:"controllerAddress" binding:"required"` // 控制者地址
}

// RevokeDIDResponse 表示撤销DID的响应
type RevokeDIDResponse struct {
	DID    string `json:"did"`    // 撤销的DID
	Status string `json:"status"` // 撤销状态
}

// IssueCredentialRequest 表示颁发凭证的请求
type IssueCredentialRequest struct {
	IssuerDID      string `json:"issuerDid" binding:"required"`      // 颁发者DID
	SubjectDID     string `json:"subjectDid" binding:"required"`     // 主体DID
	CredentialType string `json:"credentialType" binding:"required"` // 凭证类型
}

// IssueCredentialResponse 表示颁发凭证的响应
type IssueCredentialResponse struct {
	Credential VerifiableCredentialResponse `json:"credential"` // 完整凭证
}

// VerifiableCredentialResponse 表示可验证凭证的响应格式
type VerifiableCredentialResponse struct {
	Context           []string               `json:"@context,omitempty"`
	ID                string                 `json:"id"`
	Type              []string               `json:"type"`
	Issuer            string                 `json:"issuer"`
	Subject           string                 `json:"subject,omitempty"` // 用于兼容vc_handlers.go
	IssuanceDate      string                 `json:"issuanceDate,omitempty"`
	IssuedAt          string                 `json:"issuedAt,omitempty"` // 用于兼容vc_handlers.go
	ExpirationDate    string                 `json:"expirationDate,omitempty"`
	CredentialSubject map[string]interface{} `json:"credentialSubject,omitempty"`
	CredentialStatus  map[string]interface{} `json:"credentialStatus,omitempty"`
	Proof             Proof                  `json:"proof,omitempty"`
}

// VerifyCredentialRequest 表示验证凭证的请求
type VerifyCredentialRequest struct {
	CredentialID string `json:"credentialId" binding:"required"` // 要验证的凭证ID
}

// VerifyCredentialResponse 表示验证凭证的响应
type VerifyCredentialResponse struct {
	Valid       bool     `json:"valid"`                 // 是否有效
	Reason      string   `json:"reason"`                // 无效原因（如果有）
	IssuerDID   string   `json:"issuerDid,omitempty"`   // 颁发者DID
	SubjectDID  string   `json:"subjectDid,omitempty"`  // 主体DID
	IssuedAt    string   `json:"issuedAt,omitempty"`    // 颁发时间
	ExpiresAt   string   `json:"expiresAt,omitempty"`   // 过期时间
	Credentials []string `json:"credentials,omitempty"` // 凭证类型列表
}

// RevokeCredentialRequest 表示撤销凭证的请求
type RevokeCredentialRequest struct {
	CredentialID string `json:"credentialId" binding:"required"` // 要撤销的凭证ID
	IssuerDID    string `json:"issuerDid" binding:"required"`    // 颁发者DID
	Reason       string `json:"reason"`                          // 撤销原因
}

// RevokeCredentialResponse 表示撤销凭证的响应
type RevokeCredentialResponse struct {
	CredentialID string `json:"credentialId"` // 撤销的凭证ID
	Status       string `json:"status"`       // 撤销状态
	RevokedAt    string `json:"revokedAt"`    // 撤销时间
	Reason       string `json:"reason"`       // 撤销原因
}

// CreatePresentationRequest 表示创建表示的请求
type CreatePresentationRequest struct {
	HolderDID     string   `json:"holderDid" binding:"required"`     // 持有者DID
	VerifierDID   string   `json:"verifierDid"`                      // 验证者DID
	CredentialIDs []string `json:"credentialIds" binding:"required"` // 要包含的凭证ID列表
	Purpose       string   `json:"purpose" binding:"required"`       // 展示目的
}

// CreatePresentationResponse 表示创建表示的响应
type CreatePresentationResponse struct {
	Presentation VerifiablePresentationResponse `json:"presentation"` // 完整表示
}

// VerifiablePresentationResponse 表示可验证展示的响应格式
type VerifiablePresentationResponse struct {
	Context              []string                       `json:"@context"`
	ID                   string                         `json:"id"`
	Type                 []string                       `json:"type"`
	Holder               string                         `json:"holder"`
	VerifiableCredential []VerifiableCredentialResponse `json:"verifiableCredential"`
	Proof                Proof                          `json:"proof"`
}

// VerifyPresentationRequest 表示验证表示的请求
type VerifyPresentationRequest struct {
	PresentationID string `json:"presentationId" binding:"required"` // 要验证的表示ID
}

// VerifyPresentationResponse 表示验证表示的响应
type VerifyPresentationResponse struct {
	Valid         bool     `json:"valid"`                   // 是否有效
	Reason        string   `json:"reason"`                  // 无效原因（如果有）
	HolderDID     string   `json:"holderDid,omitempty"`     // 持有者DID
	PresentedAt   string   `json:"presentedAt,omitempty"`   // 展示时间
	CredentialIDs []string `json:"credentialIds,omitempty"` // 包含的凭证ID
}

// CreateDIDFromWalletRequest 表示从钱包创建DID的请求
type CreateDIDFromWalletRequest struct {
	WalletAddress string `json:"walletAddress" binding:"required"` // 钱包地址
}

// CreateDIDFromWalletResponse 表示从钱包创建DID的响应
type CreateDIDFromWalletResponse struct {
	DID           string `json:"did"`           // 创建的DID
	WalletAddress string `json:"walletAddress"` // 钱包地址
	Exists        bool   `json:"exists"`        // 是否已存在
}

// Proof 表示凭证或表示的证明
type Proof struct {
	Type               string `json:"type"`
	Created            string `json:"created"`
	VerificationMethod string `json:"verificationMethod"`
	ProofPurpose       string `json:"proofPurpose"`
	Challenge          string `json:"challenge,omitempty"`
	ProofValue         string `json:"proofValue"`
}

// CreateDoctorDIDRequest 创建医生DID的请求
type CreateDoctorDIDRequest struct {
	WalletAddress string `json:"walletAddress" binding:"required"` // 钱包地址
	Name          string `json:"name" binding:"required"`          // 医生姓名
	LicenseNumber string `json:"licenseNumber" binding:"required"` // 执业编号
}

// CreateDoctorDIDResponse 创建医生DID的响应
type CreateDoctorDIDResponse struct {
	DID           string `json:"did"`           // 创建的DID
	WalletAddress string `json:"walletAddress"` // 钱包地址
	Name          string `json:"name"`          // 医生姓名
	LicenseNumber string `json:"licenseNumber"` // 执业编号
}

// IssueDoctorVCRequest 颁发医生凭证的请求
type IssueDoctorVCRequest struct {
	IssuerDID string `json:"issuerDid" binding:"required"` // 颁发者DID (医院)
	DoctorDID string `json:"doctorDid" binding:"required"` // 医生DID
	VCType    string `json:"vcType" binding:"required"`    // 凭证类型
	VCContent string `json:"vcContent" binding:"required"` // 凭证内容
}

// IssueDoctorVCResponse 颁发医生凭证的响应
type IssueDoctorVCResponse struct {
	VCID      string    `json:"vcId"`      // 凭证ID
	DoctorDID string    `json:"doctorDid"` // 医生DID
	IssuerDID string    `json:"issuerDid"` // 颁发者DID
	VCType    string    `json:"vcType"`    // 凭证类型
	IssuedAt  time.Time `json:"issuedAt"`  // 颁发时间
	ExpiresAt time.Time `json:"expiresAt"` // 过期时间
}

// VerifyDoctorVCRequest 验证医生凭证的请求
type VerifyDoctorVCRequest struct {
	VCID string `json:"vcId" binding:"required"` // 凭证ID
}

// VerifyDoctorVCResponse 验证医生凭证的响应
type VerifyDoctorVCResponse struct {
	Valid     bool   `json:"valid"`               // 是否有效
	DoctorDID string `json:"doctorDid,omitempty"` // 医生DID
	IssuerDID string `json:"issuerDid,omitempty"` // 颁发者DID
	VCType    string `json:"vcType,omitempty"`    // 凭证类型
	Reason    string `json:"reason,omitempty"`    // 无效原因
}

// GetDoctorVCsRequest 获取医生凭证列表请求
type GetDoctorVCsRequest struct {
	DoctorDID string `json:"doctorDid" binding:"required"` // 医生DID
}

// GetDoctorVCsResponse 获取医生凭证列表响应
type GetDoctorVCsResponse struct {
	DoctorDID string     `json:"doctorDid"`             // 医生DID
	VCs       []DoctorVC `json:"verifiableCredentials"` // 凭证列表
}
