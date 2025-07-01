package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ABE/nft/nft-go-backend/internal/api/nft/service"
	"github.com/ABE/nft/nft-go-backend/internal/blockchain"
	"github.com/ABE/nft/nft-go-backend/internal/models"
)

// NFTHandlers NFT相关处理程序结构体
type NFTHandlers struct {
	Service *service.NFTService
}

// NewNFTHandlers 创建新的NFT处理程序
func NewNFTHandlers(client *blockchain.EthClient) *NFTHandlers {
	return &NFTHandlers{
		Service: service.NewNFTService(client),
	}
}

// GetAllNFTsHandler 获取所有NFT - 只显示主NFT
func (h *NFTHandlers) GetAllNFTsHandler(c *gin.Context) {
	nftResponses, err := h.Service.GetAllMainNFTs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询NFT失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"nfts": nftResponses})
}

// GetUserNFTsHandler 获取指定用户的NFT
func (h *NFTHandlers) GetUserNFTsHandler(c *gin.Context) {
	address := c.Param("address")

	nftResponses, err := h.Service.GetUserNFTs(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询NFT失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"nfts": nftResponses})
}

// GetMyNFTsHandler 获取当前用户的NFT (需要认证) - 包含子NFT
func (h *NFTHandlers) GetMyNFTsHandler(c *gin.Context) {
	walletAddress, exists := c.Get("walletAddress")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经过身份验证"})
		return
	}

	nftResponses, err := h.Service.GetMyNFTs(walletAddress.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"nfts": nftResponses})
}

// GetNFTHandler 获取NFT信息的处理程序
func (h *NFTHandlers) GetNFTHandler(c *gin.Context) {
	// 从URL获取参数
	tokenIDStr := c.Param("tokenId")

	response, err := h.Service.GetNFTInfo(tokenIDStr)
	if err != nil {
		if err.Error() == "无效的token ID" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// 返回JSON响应
	c.JSON(http.StatusOK, response)
}

// MintNFTHandler 铸造NFT的处理程序
func (h *NFTHandlers) MintNFTHandler(c *gin.Context) {
	// 获取验证后的钱包地址
	walletAddress, exists := c.Get("walletAddress")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经过身份验证"})
		return
	}

	// 获取保存的原始请求体
	rawBody, exists := c.Get("rawRequestBody")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取请求体"})
		return
	}

	bodyBytes, ok := rawBody.([]byte)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "请求体格式错误"})
		return
	}

	// 解析请求体
	var req models.MintRequest
	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		fmt.Printf("解析请求体错误: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 验证签名请求地址与钱包地址匹配
	if req.Address != walletAddress.(string) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求地址与签名地址不匹配"})
		return
	}

	response, err := h.Service.MintNFT(walletAddress.(string), req.URI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回JSON响应
	c.JSON(http.StatusOK, response)
}

// UpdateMetadataHandler 更新NFT元数据的处理程序
func (h *NFTHandlers) UpdateMetadataHandler(c *gin.Context) {
	// 获取验证后的钱包地址
	walletAddress, exists := c.Get("walletAddress")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经过身份验证"})
		return
	}

	// 解析请求体
	var req models.UpdateMetadataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 验证签名请求地址与钱包地址匹配
	if req.Address != walletAddress.(string) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求地址与签名地址不匹配"})
		return
	}

	response, err := h.Service.UpdateMetadata(walletAddress.(string), req.TokenID, req.ContractType, req.NewURI)
	if err != nil {
		if strings.Contains(err.Error(), "只有") {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else if strings.Contains(err.Error(), "无效的") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateNFTURIHandler 更新NFT元数据URI的处理程序
func (h *NFTHandlers) UpdateNFTURIHandler(c *gin.Context) {
	// 获取验证后的钱包地址
	walletAddress, exists := c.Get("walletAddress")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经过身份验证"})
		return
	}

	// 获取保存的原始请求体
	rawBody, exists := c.Get("rawRequestBody")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取请求体"})
		return
	}

	bodyBytes, ok := rawBody.([]byte)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "请求体格式错误"})
		return
	}

	// 解析请求体
	var req models.UpdateURIRequest
	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 验证签名请求地址与钱包地址匹配
	if req.Address != walletAddress.(string) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求地址与签名地址不匹配"})
		return
	}

	response, err := h.Service.UpdateNFTURI(walletAddress.(string), req.TokenID, req.NewURI)
	if err != nil {
		if strings.Contains(err.Error(), "只有") {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		} else if strings.Contains(err.Error(), "无效的") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// 返回JSON响应
	c.JSON(http.StatusOK, response)
}
