package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

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
		Gamma []string `json:"gamma"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 使用提供的属性或默认属性
	attributes := req.Gamma
	if len(attributes) == 0 {
		// 使用默认属性
		attributes = []string{
			"department:HR",
			"department:IT",
			"department:FINANCE",
			"role:admin",
			"role:manager",
			"role:employee",
			"role:doctor",
			"level:1",
			"level:2",
			"level:3",
			"level:4",
			"level:5",
			"admin:true",
			"admin:false",
			"hospital:301",
			"specialty:cardiology",
			"clearance:high",
			"clearance:medium",
			"clearance:low",
		}
	}

	// 调用服务初始化ABE系统 (使用默认用户ID 1)
	systemKey, err := h.Service.SetupABE(attributes, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "初始化ABE系统失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"system_key_id": systemKey.ID,
		"pub_key":       systemKey.PubKey,
		"sec_key":       systemKey.SecKey,
		"message":       "ABE系统初始化成功",
	})
}

// KeyGenABE 生成用户密钥处理程序
func (h *ABEHandlers) KeyGenABE(c *gin.Context) {
	var req struct {
		WalletAddress string   `json:"wallet_address" binding:"required"`
		Attributes    []string `json:"attributes"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 验证钱包地址格式
	if len(req.WalletAddress) != 42 || !strings.HasPrefix(req.WalletAddress, "0x") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "钱包地址格式错误，必须是有效的以太坊地址"})
		return
	}

	// 自动获取或创建系统密钥
	systemKey, err := h.Service.GetOrCreateSystemKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取系统密钥失败: " + err.Error()})
		return
	}

	// 确定用户属性
	var userAttributes []string
	if len(req.Attributes) > 0 {
		// 使用前端传来的NFT属性
		userAttributes = req.Attributes

		// 验证属性格式
		for _, attr := range userAttributes {
			if !strings.HasPrefix(attr, "mainNFT:") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "属性格式错误，必须是 mainNFT:NFT主地址 格式"})
				return
			}

			// 提取NFT主地址并验证
			nftAddress := strings.TrimPrefix(attr, "mainNFT:")
			if len(nftAddress) != 42 || !strings.HasPrefix(nftAddress, "0x") {
				c.JSON(http.StatusBadRequest, gin.H{"error": "NFT主地址格式错误，必须是有效的以太坊地址"})
				return
			}
		}
	} else {
		// 如果没有提供属性，使用默认的钱包地址属性（向后兼容）
		userAttributes = []string{
			"mainNFT:" + req.WalletAddress,
		}
	}

	// 调用服务生成用户密钥 (使用默认用户ID 1)
	userKey, err := h.Service.KeyGenABE(systemKey.ID, 1, userAttributes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成用户密钥失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_key_id":    userKey.ID,
		"attrib_keys":    userKey.AttribKeys,
		"wallet_address": req.WalletAddress,
		"attributes":     userAttributes,
		"nft_count":      len(userAttributes),
		"message":        "用户密钥生成成功",
	})
}

// EncryptABE 加密数据处理程序
func (h *ABEHandlers) EncryptABE(c *gin.Context) {
	var req struct {
		Message string `json:"message" binding:"required"`
		Policy  string `json:"policy" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 验证Policy格式必须是mainNFT:钱包地址
	if !strings.HasPrefix(req.Policy, "mainNFT:") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "访问策略格式错误，必须是 mainNFT:钱包地址 格式"})
		return
	}

	// 提取钱包地址
	walletAddress := strings.TrimPrefix(req.Policy, "mainNFT:")
	if len(walletAddress) != 42 || !strings.HasPrefix(walletAddress, "0x") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "钱包地址格式错误，必须是有效的以太坊地址"})
		return
	}

	// 自动获取或创建系统密钥
	systemKey, err := h.Service.GetOrCreateSystemKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取系统密钥失败: " + err.Error()})
		return
	}

	// 调用服务加密数据 (使用默认用户ID 1)
	ciphertext, err := h.Service.EncryptABE(systemKey.ID, req.Message, req.Policy, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "加密数据失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ciphertext_id":  ciphertext.ID,
		"cipher":         ciphertext.Cipher,
		"policy":         ciphertext.Policy,
		"wallet_address": walletAddress,
		"message":        "数据加密成功",
	})
}

// DecryptABE 解密数据处理程序
func (h *ABEHandlers) DecryptABE(c *gin.Context) {
	var req struct {
		Cipher     string `json:"cipher" binding:"required"`
		AttribKeys string `json:"attrib_keys" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务解密数据（直接解密，不依赖数据库记录）
	message, err := h.Service.DecryptABEDirect(req.Cipher, req.AttribKeys)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "解密数据失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"status":  "success",
	})
}

// UploadImageABE 上传图片到IPFS处理程序
func (h *ABEHandlers) UploadImageABE(c *gin.Context) {
	// 解析multipart表单
	err := c.Request.ParseMultipartForm(32 << 20) // 32MB
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "解析表单失败: " + err.Error()})
		return
	}

	file, handler, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取图片文件失败: " + err.Error()})
		return
	}
	defer file.Close()

	// 验证文件类型
	allowedTypes := []string{".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg", ".bmp"}
	ext := strings.ToLower(filepath.Ext(handler.Filename))
	isValidType := false
	for _, allowedType := range allowedTypes {
		if ext == allowedType {
			isValidType = true
			break
		}
	}

	if !isValidType {
		c.JSON(http.StatusBadRequest, gin.H{"error": "不支持的图片格式，支持: jpg, jpeg, png, gif, webp, svg, bmp"})
		return
	}

	// 验证文件大小 (10MB限制)
	if handler.Size > 10<<20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "图片文件大小不能超过10MB"})
		return
	}

	// 读取文件内容
	fileContent, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "读取文件内容失败: " + err.Error()})
		return
	}

	// 上传到IPFS
	ipfsHash, err := h.uploadToLocalIPFS(fileContent, handler.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "上传到IPFS失败: " + err.Error()})
		return
	}

	// 生成多个访问URL
	ipfsGateways := []string{
		"https://dweb.link/ipfs/",
		"https://cloudflare-ipfs.com/ipfs/",
		"https://gateway.pinata.cloud/ipfs/",
		"https://ipfs.io/ipfs/",
	}

	var urls []string
	for _, gateway := range ipfsGateways {
		urls = append(urls, gateway+ipfsHash)
	}

	c.JSON(http.StatusOK, gin.H{
		"hash":         ipfsHash,
		"url":          "ipfs://" + ipfsHash,
		"http_urls":    urls,
		"primary_url":  urls[0], // 推荐使用的URL
		"filename":     handler.Filename,
		"size":         handler.Size,
		"content_type": handler.Header.Get("Content-Type"),
		"message":      "图片已成功上传到IPFS",
	})
}

// uploadToLocalIPFS 上传文件到本地IPFS节点
func (h *ABEHandlers) uploadToLocalIPFS(data []byte, filename string) (string, error) {
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
