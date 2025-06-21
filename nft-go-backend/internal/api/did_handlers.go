package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ABE/nft/nft-go-backend/internal/models"
	"github.com/ABE/nft/nft-go-backend/internal/service"
)

// DIDHandlers DID相关处理程序结构体
type DIDHandlers struct {
	Service *service.DIDService
}

// NewDIDHandlers 创建新的DID处理程序
func NewDIDHandlers(service *service.DIDService) *DIDHandlers {
	return &DIDHandlers{
		Service: service,
	}
}

// CreateDIDFromWalletHandler 从钱包创建DID处理程序
func (h *DIDHandlers) CreateDIDFromWalletHandler(c *gin.Context) {
	walletAddress := c.Param("walletAddress")
	if walletAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "钱包地址不能为空"})
		return
	}

	// 调用服务创建DID
	did, exists, err := h.Service.CreateDIDFromWallet(walletAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建DID失败: " + err.Error()})
		return
	}

	// 构建响应
	response := models.CreateDIDFromWalletResponse{
		DID:           did.DIDString,
		WalletAddress: did.WalletAddress,
		Exists:        exists,
	}

	c.JSON(http.StatusOK, response)
}

// GetDIDByWalletHandler 获取钱包DID处理程序
func (h *DIDHandlers) GetDIDByWalletHandler(c *gin.Context) {
	walletAddress := c.Param("walletAddress")
	if walletAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "钱包地址不能为空"})
		return
	}

	// 调用服务获取DID
	did, err := h.Service.GetDIDByWallet(walletAddress)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "获取DID失败: " + err.Error()})
		return
	}

	// 构建响应
	response := models.CreateDIDFromWalletResponse{
		DID:           did.DIDString,
		WalletAddress: did.WalletAddress,
		Exists:        true,
	}

	c.JSON(http.StatusOK, response)
}

// ListDIDsByWalletHandler 列出钱包DID处理程序
func (h *DIDHandlers) ListDIDsByWalletHandler(c *gin.Context) {
	walletAddress := c.Param("walletAddress")
	if walletAddress == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "钱包地址不能为空"})
		return
	}

	// 调用服务列出DID
	dids, err := h.Service.ListDIDsByWallet(walletAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取DID列表失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"dids": dids})
}

// GetAllDIDsHandler 获取所有DID处理程序
func (h *DIDHandlers) GetAllDIDsHandler(c *gin.Context) {
	// 调用服务获取所有DID
	dids, err := h.Service.GetAllDIDs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取DID列表失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"dids": dids})
}

// ResolveDIDHandler 解析DID处理程序
func (h *DIDHandlers) ResolveDIDHandler(c *gin.Context) {
	var req models.DIDResolutionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务解析DID
	doc, err := h.Service.ResolveDID(req.DID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "解析DID失败: " + err.Error()})
		return
	}

	// 构建响应
	response := models.DIDResolutionResponse{
		DIDDocument: *doc,
		Status:      "resolved",
	}

	c.JSON(http.StatusOK, response)
}

// 已弃用的方法

// CreateDIDHandler 创建DID处理程序（已弃用）
func (h *DIDHandlers) CreateDIDHandler(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "此方法已弃用，请使用 /api/did/wallet/:walletAddress 创建DID"})
}

// UpdateDIDHandler 更新DID处理程序（已弃用）
func (h *DIDHandlers) UpdateDIDHandler(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "此方法已弃用"})
}

// RevokeDIDHandler 撤销DID处理程序（已弃用）
func (h *DIDHandlers) RevokeDIDHandler(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "此方法已弃用"})
}
