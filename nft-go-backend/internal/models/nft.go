package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

// NFT 表示NFT实体的数据库模型
type NFT struct {
	gorm.Model
	TokenID       string `json:"tokenId" gorm:"unique"`
	Owner         string `json:"owner"`
	URI           string `json:"uri"`
	TotalSupply   string `json:"totalSupply,omitempty"`
	IsChildNFT    bool   `json:"isChildNft" gorm:"default:false"`  // 标识是否为子NFT
	ParentTokenID string `json:"parentTokenId,omitempty"`          // 父NFT的TokenID（仅子NFT有效）
	ContractType  string `json:"contractType" gorm:"default:main"` // 合约类型：main或child
}

// SignedRequest 表示签名请求的基础结构
type SignedRequest struct {
	Address   string `json:"address" binding:"required"`   // 钱包地址
	Signature string `json:"signature" binding:"required"` // 签名
	Message   string `json:"message" binding:"required"`   // 签名的消息
}

// NFTResponse 表示NFT信息的响应结构
type NFTResponse struct {
	TokenID       string `json:"tokenId"`
	Owner         string `json:"owner"`
	URI           string `json:"uri"`
	TotalSupply   string `json:"totalSupply,omitempty"`
	IsChildNFT    bool   `json:"isChildNft"`
	ParentTokenID string `json:"parentTokenId,omitempty"`
	ContractType  string `json:"contractType"`
}

// MintRequest 表示铸造NFT的请求结构
type MintRequest struct {
	SignedRequest
	URI string `json:"uri" binding:"required"`
}

// CreateChildRequest 表示创建子NFT的请求结构
type CreateChildRequest struct {
	SignedRequest
	ParentTokenID string `json:"parentTokenId" binding:"required"`
	Recipient     string `json:"recipient" binding:"required"`
	URI           string `json:"uri" binding:"required"`
}

// ChildNFTRequest 表示申请子NFT的数据库记录结构
type ChildNFTRequest struct {
	gorm.Model
	RequestId        string `json:"requestId" gorm:"-"` // 虚拟字段，用于API响应
	ChildTokenID     string `json:"childTokenId"`
	ParentTokenId    string `json:"parentTokenId"`
	ApplicantAddress string `json:"applicantAddress"`
	URI              string `json:"uri"`
	Status           string `json:"status" gorm:"default:pending"`     // pending, approved, rejected, auto_approved
	VCCredentials    string `json:"vcCredentials"`                     // 提交的VC凭证（JSON）
	AutoApproved     bool   `json:"autoApproved" gorm:"default:false"` // 是否自动审核通过
	PolicyResult     string `json:"policyResult"`                      // 策略验证结果（JSON）
}

// MarshalJSON 自定义JSON序列化，确保ID字段被正确包含
func (c ChildNFTRequest) MarshalJSON() ([]byte, error) {
	type Alias ChildNFTRequest
	return json.Marshal(&struct {
		ID uint `json:"id"`
		*Alias
	}{
		ID:    c.ID,
		Alias: (*Alias)(&c),
	})
}

// RequestChildNFTRequest 表示申请子NFT的API请求结构
type RequestChildNFTRequest struct {
	SignedRequest
	ParentTokenId    string `json:"parentTokenId" binding:"required"`
	ApplicantAddress string `json:"applicantAddress" binding:"required"`
	URI              string `json:"uri" binding:"required"`
	VCCredentials    string `json:"vcCredentials,omitempty"` // VC凭证JSON字符串（可选）
	AutoApprove      bool   `json:"autoApprove,omitempty"`   // 是否尝试自动审核
}

// TransactionResponse 表示交易响应的结构
type TransactionResponse struct {
	TransactionHash string `json:"transactionHash"`
	Message         string `json:"message"`
}

// NFTMetadata 表示NFT元数据的结构
type NFTMetadata struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Attributes  []Attribute `json:"attributes"`
}

// Attribute 表示NFT属性的结构
type Attribute struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}

// NFTMetadataRequest 表示NFT元数据的请求结构
type NFTMetadataRequest struct {
	ParentTokenID string `json:"parentTokenId" binding:"required"`
	Recipient     string `json:"recipient" binding:"required"`
	URI           string `json:"uri" binding:"required"`
}

// NFTResponseWithMetadata 表示包含元数据的NFT响应结构
type NFTResponseWithMetadata struct {
	NFTResponse
	Metadata *NFTMetadata `json:"metadata,omitempty"`
}

// NFTMetadataDB 表示NFT元数据的数据库模型
type NFTMetadataDB struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	ExternalURL string `json:"external_url"`
	Image       string `json:"image" gorm:"not null"`
	Policy      string `json:"policy"`
	Ciphertext  string `json:"ciphertext"`
	IPFSHash    string `json:"ipfs_hash" gorm:"unique;not null"`
}

// CreateMetadataRequest 表示创建元数据的请求结构
type CreateMetadataRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ExternalURL string `json:"external_url"`
	Image       string `json:"image" binding:"required"`
	Policy      string `json:"policy" binding:"required"`
	Ciphertext  string `json:"ciphertext" binding:"required"`
}

// MetadataResponse 表示元数据响应结构
type MetadataResponse struct {
	IPFSHash string `json:"ipfs_hash"`
	Message  string `json:"message"`
}

// UpdateMetadataRequest 表示更新元数据的请求结构
type UpdateMetadataRequest struct {
	SignedRequest
	TokenID      string `json:"tokenId" binding:"required"`
	NewURI       string `json:"newUri" binding:"required"`
	ContractType string `json:"contractType" binding:"required"` // "main" 或 "child"
}

// UpdateURIRequest 表示更新NFT URI的请求结构
type UpdateURIRequest struct {
	SignedRequest
	TokenID string `json:"tokenId" binding:"required"`
	NewURI  string `json:"newUri" binding:"required"`
}

// GetAllRequestsResponse 表示获取所有申请记录的响应结构
type GetAllRequestsResponse struct {
	Requests []ChildNFTRequestWithParentInfo `json:"requests"`
}

// ChildNFTRequestWithParentInfo 表示包含父NFT信息的子NFT申请记录
type ChildNFTRequestWithParentInfo struct {
	ChildNFTRequest
	ParentNFTOwner string `json:"parentNftOwner"`
	CanOperate     bool   `json:"canOperate"` // 当前用户是否可以操作此申请
}

// MarshalJSON 自定义JSON序列化，确保ID字段被正确包含
func (c ChildNFTRequestWithParentInfo) MarshalJSON() ([]byte, error) {
	type Alias ChildNFTRequestWithParentInfo
	return json.Marshal(&struct {
		ID uint `json:"id"`
		*Alias
	}{
		ID:    c.ID,
		Alias: (*Alias)(&c),
	})
}
