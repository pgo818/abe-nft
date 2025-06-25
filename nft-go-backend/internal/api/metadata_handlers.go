package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ABE/nft/nft-go-backend/internal/blockchain"
	"github.com/ABE/nft/nft-go-backend/internal/models"
)

// MetadataHandlers 元数据相关处理程序结构体
type MetadataHandlers struct {
	Client *blockchain.EthClient
}

// NewMetadataHandlers 创建新的元数据处理程序
func NewMetadataHandlers(client *blockchain.EthClient) *MetadataHandlers {
	return &MetadataHandlers{
		Client: client,
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

	// 构造元数据对象
	metadata := map[string]interface{}{
		"description":  req.Description,
		"external_url": req.ExternalURL,
		"image":        req.Image,
		"name":         req.Name,
		"attributes": []map[string]interface{}{
			{
				"trait_type": "Policy",
				"value":      req.Policy,
			},
			{
				"trait_type": "Encrypted_ciphertext",
				"value":      req.Ciphertext,
			},
		},
	}

	// 将元数据转换为JSON
	metadataJSON, err := json.Marshal(metadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JSON序列化失败: " + err.Error()})
		return
	}

	// 上传到本地IPFS
	ipfsHash, err := h.uploadToLocalIPFS(metadataJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "上传到本地IPFS失败: " + err.Error()})
		return
	}

	// 尝试保存到数据库，但忽略错误
	metadataDB := models.NFTMetadataDB{
		Name:        req.Name,
		Description: req.Description,
		ExternalURL: req.ExternalURL,
		Image:       req.Image,
		Policy:      req.Policy,
		Ciphertext:  req.Ciphertext,
		IPFSHash:    ipfsHash,
	}

	// 尝试保存，但忽略错误
	models.DB.Create(&metadataDB)
	// 不检查错误，直接继续

	// 返回响应
	response := models.MetadataResponse{
		IPFSHash: ipfsHash,
		Message:  "元数据已成功创建并保存到IPFS",
	}

	c.JSON(http.StatusOK, response)
}

// GetMetadataHandler 根据IPFS哈希获取元数据
func (h *MetadataHandlers) GetMetadataHandler(c *gin.Context) {
	ipfsHash := c.Param("hash")

	var metadata models.NFTMetadataDB
	result := models.DB.Where("ipfs_hash = ?", ipfsHash).First(&metadata)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "元数据不存在"})
		return
	}

	// 构造响应
	response := map[string]interface{}{
		"description":  metadata.Description,
		"external_url": metadata.ExternalURL,
		"image":        metadata.Image,
		"name":         metadata.Name,
		"attributes": []map[string]interface{}{
			{
				"trait_type": "Policy",
				"value":      metadata.Policy,
			},
			{
				"trait_type": "Encrypted_ciphertext",
				"value":      metadata.Ciphertext,
			},
		},
	}

	c.JSON(http.StatusOK, response)
}

// GetAllMetadataHandler 获取所有元数据
func (h *MetadataHandlers) GetAllMetadataHandler(c *gin.Context) {
	var metadataList []models.NFTMetadataDB
	result := models.DB.Find(&metadataList)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询元数据失败: " + result.Error.Error()})
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

	// 如果没有提供文件名，使用默认值
	if req.Filename == "" {
		if req.IsBinary {
			req.Filename = "data.bin"
		} else {
			req.Filename = "data.txt"
		}
	}

	var data []byte
	var err error

	if req.IsBinary {
		// 处理二进制数据：Base64解码
		data, err = base64.StdEncoding.DecodeString(req.Data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Base64解码失败: " + err.Error()})
			return
		}
	} else {
		// 处理文本数据
		data = []byte(req.Data)
	}

	// 上传到本地IPFS
	ipfsHash, err := h.uploadToLocalIPFSWithFilename(data, req.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "上传到IPFS失败: " + err.Error()})
		return
	}

	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"hash":      ipfsHash,
		"url":       "ipfs://" + ipfsHash,
		"message":   "数据已成功上传到IPFS",
		"filename":  req.Filename,
		"is_binary": req.IsBinary,
		"size":      len(data),
	})
}

// uploadToLocalIPFSWithFilename 上传到本地IPFS节点（带文件名）
func (h *MetadataHandlers) uploadToLocalIPFSWithFilename(data []byte, filename string) (string, error) {
	// 本地IPFS节点的API地址（默认端口5001）
	ipfsAPIURL := "http://localhost:5001/api/v0/add"

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// 创建文件字段
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return "", fmt.Errorf("创建表单文件失败: %w", err)
	}

	// 写入数据
	_, err = part.Write(data)
	if err != nil {
		return "", fmt.Errorf("写入数据失败: %w", err)
	}

	// 关闭writer
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("关闭writer失败: %w", err)
	}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", ipfsAPIURL, &buf)
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置Content-Type
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求到本地IPFS节点
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("发送请求到本地IPFS失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("本地IPFS API返回错误 %d: %s", resp.StatusCode, string(body))
	}

	// 解析响应
	var result struct {
		Hash string `json:"Hash"`
		Name string `json:"Name"`
		Size string `json:"Size"`
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %w", err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("解析响应失败: %w", err)
	}

	return result.Hash, nil
}

// uploadToLocalIPFS 上传到本地IPFS节点（兼容性包装器）
func (h *MetadataHandlers) uploadToLocalIPFS(data []byte) (string, error) {
	return h.uploadToLocalIPFSWithFilename(data, "metadata.json")
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

	// 从本地IPFS节点获取文件
	content, contentType, err := h.getFromLocalIPFS(hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "从IPFS获取文件失败: " + err.Error()})
		return
	}

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
