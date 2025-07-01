package service

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/ABE/nft/nft-go-backend/internal/blockchain"
	"github.com/ABE/nft/nft-go-backend/internal/models"
)

// MetadataService 元数据业务逻辑服务
type MetadataService struct {
	Client *blockchain.EthClient
}

// NewMetadataService 创建新的元数据服务
func NewMetadataService(client *blockchain.EthClient) *MetadataService {
	return &MetadataService{
		Client: client,
	}
}

// CreateMetadata 创建NFT元数据并保存到IPFS
func (s *MetadataService) CreateMetadata(req models.CreateMetadataRequest) (*models.MetadataResponse, error) {
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
		return nil, fmt.Errorf("JSON序列化失败: %w", err)
	}

	// 上传到本地IPFS
	ipfsHash, err := s.uploadToLocalIPFS(metadataJSON)
	if err != nil {
		return nil, fmt.Errorf("上传到本地IPFS失败: %w", err)
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

	// 返回响应
	response := &models.MetadataResponse{
		IPFSHash: ipfsHash,
		Message:  "元数据已成功创建并保存到IPFS",
	}

	return response, nil
}

// GetMetadata 根据IPFS哈希获取元数据
func (s *MetadataService) GetMetadata(ipfsHash string) (map[string]interface{}, error) {
	var metadata models.NFTMetadataDB
	result := models.DB.Where("ipfs_hash = ?", ipfsHash).First(&metadata)
	if result.Error != nil {
		return nil, fmt.Errorf("元数据不存在")
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

	return response, nil
}

// GetAllMetadata 获取所有元数据
func (s *MetadataService) GetAllMetadata() ([]models.NFTMetadataDB, error) {
	var metadataList []models.NFTMetadataDB
	result := models.DB.Find(&metadataList)
	if result.Error != nil {
		return nil, fmt.Errorf("查询元数据失败: %w", result.Error)
	}

	return metadataList, nil
}

// UploadToIPFS 通用IPFS上传 - 支持二进制数据
func (s *MetadataService) UploadToIPFS(data string, filename string, isBinary bool) (map[string]interface{}, error) {
	// 如果没有提供文件名，使用默认值
	if filename == "" {
		if isBinary {
			filename = "data.bin"
		} else {
			filename = "data.txt"
		}
	}

	var dataBytes []byte
	var err error

	if isBinary {
		// 处理二进制数据：Base64解码
		dataBytes, err = base64.StdEncoding.DecodeString(data)
		if err != nil {
			return nil, fmt.Errorf("Base64解码失败: %w", err)
		}
	} else {
		// 处理文本数据
		dataBytes = []byte(data)
	}

	// 上传到本地IPFS
	ipfsHash, err := s.uploadToLocalIPFSWithFilename(dataBytes, filename)
	if err != nil {
		return nil, fmt.Errorf("上传到IPFS失败: %w", err)
	}

	// 返回响应
	response := map[string]interface{}{
		"hash":      ipfsHash,
		"url":       "ipfs://" + ipfsHash,
		"message":   "数据已成功上传到IPFS",
		"filename":  filename,
		"is_binary": isBinary,
		"size":      len(dataBytes),
	}

	return response, nil
}

// GetFromIPFS 从IPFS获取文件
func (s *MetadataService) GetFromIPFS(ipfsHash string) ([]byte, error) {
	// 本地IPFS节点的API地址
	ipfsAPIURL := fmt.Sprintf("http://localhost:8080/ipfs/%s", ipfsHash)

	// 发送GET请求
	resp, err := http.Get(ipfsAPIURL)
	if err != nil {
		return nil, fmt.Errorf("从本地IPFS获取文件失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("本地IPFS返回错误 %d", resp.StatusCode)
	}

	// 读取响应体
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	return data, nil
}

// uploadToLocalIPFSWithFilename 上传到本地IPFS节点（带文件名）
func (s *MetadataService) uploadToLocalIPFSWithFilename(data []byte, filename string) (string, error) {
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
func (s *MetadataService) uploadToLocalIPFS(data []byte) (string, error) {
	return s.uploadToLocalIPFSWithFilename(data, "metadata.json")
}
