package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ABE/nft/nft-go-backend/internal/api/nft/service"
	"github.com/ABE/nft/nft-go-backend/internal/blockchain"
	"github.com/ABE/nft/nft-go-backend/internal/models"
)

// MetadataHandlers 元数据相关处理程序结构体
type MetadataHandlers struct {
	Service *service.MetadataService
}

// NewMetadataHandlers 创建新的元数据处理程序
func NewMetadataHandlers(client *blockchain.EthClient) *MetadataHandlers {
	return &MetadataHandlers{
		Service: service.NewMetadataService(client),
	}
}

// CreateMetadataHandler 创建NFT元数据并保存到IPFS
func (h *MetadataHandlers) CreateMetadataHandler(c *gin.Context) {
	// 解析请求体
	var req models.CreateMetadataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	response, err := h.Service.CreateMetadata(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetMetadataHandler 根据IPFS哈希获取元数据
func (h *MetadataHandlers) GetMetadataHandler(c *gin.Context) {
	ipfsHash := c.Param("hash")

	response, err := h.Service.GetMetadata(ipfsHash)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetAllMetadataHandler 获取所有元数据
func (h *MetadataHandlers) GetAllMetadataHandler(c *gin.Context) {
	metadataList, err := h.Service.GetAllMetadata()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"metadata": metadataList})
}

// UploadToIPFSHandler 通用IPFS上传处理程序 - 修复版本：支持二进制数据
func (h *MetadataHandlers) UploadToIPFSHandler(c *gin.Context) {
	var req struct {
		Data     string `json:"data" binding:"required"`
		Filename string `json:"filename"`
		IsBinary bool   `json:"is_binary"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数: " + err.Error()})
		return
	}

	response, err := h.Service.UploadToIPFS(req.Data, req.Filename, req.IsBinary)
	if err != nil {
		if strings.Contains(err.Error(), "Base64解码失败") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, response)
}

// DownloadFromIPFSHandler 通过Hash从IPFS下载文件
func (h *MetadataHandlers) DownloadFromIPFSHandler(c *gin.Context) {
	hash := c.Param("hash")

	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash参数不能为空"})
		return
	}

	// 验证Hash格式（基本验证）
	if len(hash) < 40 || (!strings.HasPrefix(hash, "Qm") && !strings.HasPrefix(hash, "bafy")) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash格式不正确"})
		return
	}

	content, err := h.Service.GetFromIPFS(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "从IPFS获取文件失败: " + err.Error()})
		return
	}

	// 检测内容类型
	contentType := h.detectContentType(content)

	// 设置响应头
	c.Header("Content-Type", contentType)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=ipfs_%s", hash[:12]))
	c.Header("X-IPFS-Hash", hash)

	// 返回文件内容
	c.Data(http.StatusOK, contentType, content)
}

// GetFromIPFSHandler 通过Hash从IPFS获取文件信息和内容
func (h *MetadataHandlers) GetFromIPFSHandler(c *gin.Context) {
	hash := c.Param("hash")

	if hash == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash参数不能为空"})
		return
	}

	// 验证Hash格式
	if len(hash) < 40 || (!strings.HasPrefix(hash, "Qm") && !strings.HasPrefix(hash, "bafy")) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash格式不正确"})
		return
	}

	// 从本地IPFS节点获取文件
	content, contentType, err := h.getFromLocalIPFS(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "从IPFS获取文件失败: " + err.Error()})
		return
	}

	// 检测文件类型和大小
	size := len(content)
	isText := h.isTextContent(content, contentType)
	isImage := strings.HasPrefix(contentType, "image/")

	// 如果是文本内容，包含在响应中；如果是二进制文件，只返回信息
	response := gin.H{
		"hash":         hash,
		"size":         size,
		"content_type": contentType,
		"is_text":      isText,
		"is_image":     isImage,
		"url":          fmt.Sprintf("https://dweb.link/ipfs/%s", hash),
		"download_url": fmt.Sprintf("/api/ipfs/download/%s", hash),
		"message":      "文件信息获取成功",
	}

	// 如果是文本文件且大小合理，包含内容
	if isText && size < 1024*1024 { // 1MB限制
		response["content"] = string(content)
	}

	c.JSON(http.StatusOK, response)
}

// getFromLocalIPFS 从本地IPFS节点获取文件
func (h *MetadataHandlers) getFromLocalIPFS(hash string) ([]byte, string, error) {
	// 本地IPFS节点的API地址
	ipfsAPIURL := fmt.Sprintf("http://localhost:5001/api/v0/cat?arg=%s", hash)

	// 创建HTTP请求
	resp, err := http.Get(ipfsAPIURL)
	if err != nil {
		return nil, "", fmt.Errorf("连接IPFS节点失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, "", fmt.Errorf("IPFS API返回错误 %d: %s", resp.StatusCode, string(body))
	}

	// 读取文件内容
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("读取文件内容失败: %w", err)
	}

	// 检测内容类型
	contentType := h.detectContentType(content)

	return content, contentType, nil
}

// detectContentType 检测文件内容类型
func (h *MetadataHandlers) detectContentType(content []byte) string {
	if len(content) == 0 {
		return "application/octet-stream"
	}

	// 检查文件头部特征
	if len(content) >= 3 && content[0] == 0xFF && content[1] == 0xD8 && content[2] == 0xFF {
		return "image/jpeg"
	}
	if len(content) >= 8 && string(content[:8]) == "\x89PNG\r\n\x1a\n" {
		return "image/png"
	}
	if len(content) >= 6 && (string(content[:6]) == "GIF87a" || string(content[:6]) == "GIF89a") {
		return "image/gif"
	}
	if len(content) >= 4 && string(content[:4]) == "RIFF" {
		return "image/webp"
	}
	if len(content) >= 4 && string(content[:4]) == "<svg" {
		return "image/svg+xml"
	}
	if len(content) >= 4 && string(content[:4]) == "%PDF" {
		return "application/pdf"
	}

	// 检查是否为文本内容
	if h.isTextContent(content, "") {
		// 进一步检测文本类型
		contentStr := string(content)
		if (strings.HasPrefix(contentStr, "{") && strings.HasSuffix(contentStr, "}")) ||
			(strings.HasPrefix(contentStr, "[") && strings.HasSuffix(contentStr, "]")) {
			return "application/json"
		}
		if strings.Contains(contentStr, "<!DOCTYPE html") || strings.Contains(contentStr, "<html") {
			return "text/html"
		}
		if strings.Contains(contentStr, "<?xml") {
			return "application/xml"
		}
		return "text/plain"
	}

	return "application/octet-stream"
}

// isTextContent 检测是否为文本内容
func (h *MetadataHandlers) isTextContent(content []byte, contentType string) bool {
	if contentType != "" && strings.HasPrefix(contentType, "text/") {
		return true
	}

	// 检查二进制字符比例
	if len(content) == 0 {
		return true
	}

	// 统计二进制字符
	binaryCount := 0
	for _, b := range content {
		// 检查是否为控制字符（除了常见的换行、制表符等）
		if (b < 0x20 && b != 0x09 && b != 0x0A && b != 0x0D) || b == 0x7F {
			binaryCount++
		}
	}

	// 如果二进制字符比例小于10%，认为是文本
	ratio := float64(binaryCount) / float64(len(content))
	return ratio < 0.1
}
