package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"

	"nft-go-backend/internal/blockchain"
	"nft-go-backend/internal/models"
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

// CreateMetadataHandler 创建NFT元数据并保存到本地IPFS
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

	// 保存到数据库
	metadataDB := models.NFTMetadataDB{
		Name:        req.Name,
		Description: req.Description,
		ExternalURL: req.ExternalURL,
		Image:       req.Image,
		Policy:      req.Policy,
		Ciphertext:  req.Ciphertext,
		IPFSHash:    ipfsHash,
	}

	result := models.DB.Create(&metadataDB)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存到数据库失败: " + result.Error.Error()})
		return
	}

	// 返回响应
	response := models.MetadataResponse{
		IPFSHash: ipfsHash,
		Message:  "元数据已成功创建并保存到本地IPFS",
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

// uploadToLocalIPFS 上传到本地IPFS节点
func (h *MetadataHandlers) uploadToLocalIPFS(data []byte) (string, error) {
	// 本地IPFS节点的API地址（默认端口5001）
	ipfsAPIURL := "http://localhost:5001/api/v0/add"

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// 创建文件字段
	part, err := writer.CreateFormFile("file", "metadata.json")
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
