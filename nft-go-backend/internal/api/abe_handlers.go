package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ABEHandlers ABE相关处理程序结构体
type ABEHandlers struct {
	Service *ABEService
}

// NewABEHandlers 创建新的ABE处理程序
func NewABEHandlers(service *ABEService) *ABEHandlers {
	return &ABEHandlers{
		Service: service,
	}
}

// SetupABE 初始化ABE系统处理程序
func (h *ABEHandlers) SetupABE(c *gin.Context) {
	var req struct {
		Attributes []string `json:"attributes" binding:"required"`
		UserID     uint     `json:"userId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务初始化ABE系统
	systemKey, err := h.Service.SetupABE(req.Attributes, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "初始化ABE系统失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"systemKeyId": systemKey.ID,
		"message":     "ABE系统初始化成功",
	})
}

// KeyGenABE 生成用户密钥处理程序
func (h *ABEHandlers) KeyGenABE(c *gin.Context) {
	var req struct {
		SystemKeyID    uint     `json:"systemKeyId" binding:"required"`
		UserID         uint     `json:"userId" binding:"required"`
		UserAttributes []string `json:"userAttributes" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务生成用户密钥
	userKey, err := h.Service.KeyGenABE(req.SystemKeyID, req.UserID, req.UserAttributes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成用户密钥失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userKeyId": userKey.ID,
		"message":   "用户密钥生成成功",
	})
}

// EncryptABE 加密数据处理程序
func (h *ABEHandlers) EncryptABE(c *gin.Context) {
	var req struct {
		SystemKeyID uint   `json:"systemKeyId" binding:"required"`
		Message     string `json:"message" binding:"required"`
		Policy      string `json:"policy" binding:"required"`
		UserID      uint   `json:"userId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务加密数据
	ciphertext, err := h.Service.EncryptABE(req.SystemKeyID, req.Message, req.Policy, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "加密数据失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ciphertextId": ciphertext.ID,
		"policy":       ciphertext.Policy,
		"message":      "数据加密成功",
	})
}

// DecryptABE 解密数据处理程序
func (h *ABEHandlers) DecryptABE(c *gin.Context) {
	var req struct {
		CiphertextID uint `json:"ciphertextId" binding:"required"`
		UserKeyID    uint `json:"userKeyId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务解密数据
	message, err := h.Service.DecryptABE(req.CiphertextID, req.UserKeyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "解密数据失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       message,
		"decryptStatus": "success",
	})
}
