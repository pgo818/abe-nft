package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ABE/nft/nft-go-backend/internal/blockchain"
	"github.com/ABE/nft/nft-go-backend/internal/models"
)

// NFTHandlers NFT相关处理程序结构体
type NFTHandlers struct {
	Client *blockchain.EthClient
}

// NewNFTHandlers 创建新的NFT处理程序
func NewNFTHandlers(client *blockchain.EthClient) *NFTHandlers {
	return &NFTHandlers{
		Client: client,
	}
}

// GetAllNFTsHandler 获取所有NFT - 只显示主NFT
func (h *NFTHandlers) GetAllNFTsHandler(c *gin.Context) {
	var nfts []models.NFT
	// 只查询主NFT，排除子NFT
	result := models.DB.Where("is_child_nft = ? OR is_child_nft IS NULL", false).Find(&nfts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询NFT失败: " + result.Error.Error()})
		return
	}

	// 获取NFT元数据
	var nftResponses []models.NFTResponseWithMetadata
	for _, nft := range nfts {
		metadata, _ := h.fetchNFTMetadata(nft.URI)

		response := models.NFTResponseWithMetadata{
			NFTResponse: models.NFTResponse{
				TokenID:      nft.TokenID,
				Owner:        nft.Owner,
				URI:          nft.URI,
				IsChildNFT:   false, // 明确标识为主NFT
				ContractType: "main",
			},
			Metadata: metadata,
		}

		nftResponses = append(nftResponses, response)
	}

	c.JSON(http.StatusOK, gin.H{"nfts": nftResponses})
}

// GetUserNFTsHandler 获取指定用户的NFT
func (h *NFTHandlers) GetUserNFTsHandler(c *gin.Context) {
	address := c.Param("address")

	var nfts []models.NFT
	result := models.DB.Where("owner = ?", address).Find(&nfts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询NFT失败: " + result.Error.Error()})
		return
	}

	// 获取NFT元数据
	var nftResponses []models.NFTResponseWithMetadata
	for _, nft := range nfts {
		metadata, _ := h.fetchNFTMetadata(nft.URI)

		response := models.NFTResponseWithMetadata{
			NFTResponse: models.NFTResponse{
				TokenID:       nft.TokenID,
				Owner:         nft.Owner,
				URI:           nft.URI,
				IsChildNFT:    nft.IsChildNFT,
				ParentTokenID: nft.ParentTokenID,
				ContractType:  nft.ContractType,
			},
			Metadata: metadata,
		}

		nftResponses = append(nftResponses, response)
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

	// 查询用户拥有的主NFT
	var mainNFTs []models.NFT
	result := models.DB.Where("owner = ? AND (is_child_nft = ? OR is_child_nft IS NULL)", walletAddress.(string), false).Find(&mainNFTs)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询主NFT失败: " + result.Error.Error()})
		return
	}

	// 查询用户拥有的子NFT - 从区块链获取
	var childNFTs []models.NFT

	// 这里我们需要从区块链扫描用户拥有的子NFT
	// 由于合约限制，我们通过数据库记录来获取
	childResult := models.DB.Where("owner = ? AND is_child_nft = ?", walletAddress.(string), true).Find(&childNFTs)
	if childResult.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询子NFT失败: " + childResult.Error.Error()})
		return
	}

	// 合并主NFT和子NFT
	allNFTs := append(mainNFTs, childNFTs...)

	// 构建响应
	var nftResponses []models.NFTResponseWithMetadata
	for _, nft := range allNFTs {
		metadata, _ := h.fetchNFTMetadata(nft.URI)

		response := models.NFTResponseWithMetadata{
			NFTResponse: models.NFTResponse{
				TokenID:       nft.TokenID,
				Owner:         nft.Owner,
				URI:           nft.URI,
				TotalSupply:   nft.TotalSupply,
				IsChildNFT:    nft.IsChildNFT,
				ParentTokenID: nft.ParentTokenID,
				ContractType:  nft.ContractType,
			},
			Metadata: metadata,
		}

		nftResponses = append(nftResponses, response)
	}

	c.JSON(http.StatusOK, gin.H{"nfts": nftResponses})
}

// fetchNFTMetadata 获取NFT元数据
func (h *NFTHandlers) fetchNFTMetadata(uri string) (*models.NFTMetadata, error) {
	fmt.Printf("获取元数据，URI: %s\n", uri)

	// 处理空URI
	if uri == "" {
		fmt.Printf("警告: URI为空，返回默认元数据\n")
		return &models.NFTMetadata{
			Name:        "未命名NFT",
			Description: "This NFT has no metadata",
			Image:       "",
			Attributes: []models.Attribute{
				{TraitType: "Type", Value: "Unknown"},
				{TraitType: "Rarity", Value: "Unknown"},
			},
		}, nil
	}

	// 检查URI是否是JSON字符串
	if strings.HasPrefix(strings.TrimSpace(uri), "{") && strings.HasSuffix(strings.TrimSpace(uri), "}") {
		fmt.Printf("URI是JSON字符串，尝试直接解析\n")
		var metadata models.NFTMetadata
		err := json.Unmarshal([]byte(uri), &metadata)
		if err == nil {
			fmt.Printf("成功从URI中解析元数据: %s\n", metadata.Name)
			return &metadata, nil
		}
		fmt.Printf("解析URI中的JSON失败: %v\n", err)
	}

	// 检查URI是否为IPFS链接
	if strings.HasPrefix(uri, "ipfs://") {
		// 从URI中提取IPFS哈希
		ipfsHash := uri[len("ipfs://"):]
		fmt.Printf("提取IPFS哈希: %s，长度: %d\n", ipfsHash, len(ipfsHash))

		// 尝试从数据库中查找该IPFS哈希对应的元数据
		var metadata models.NFTMetadataDB
		result := models.DB.Where("ip_fs_hash = ?", ipfsHash).First(&metadata)
		if result.Error == nil {
			fmt.Printf("从数据库找到元数据: %s\n", metadata.Name)
			// 从数据库中找到了元数据
			return &models.NFTMetadata{
				Name:        metadata.Name,
				Description: metadata.Description,
				Image:       metadata.Image,
				Attributes: []models.Attribute{
					{TraitType: "Type", Value: "Main NFT"},
					{TraitType: "Rarity", Value: "Common"},
					{TraitType: "Policy", Value: metadata.Policy},
					{TraitType: "Encrypted_ciphertext", Value: metadata.Ciphertext},
				},
			}, nil
		}

		// 如果数据库中没有找到，创建一个基本的元数据
		fmt.Printf("数据库中未找到元数据，创建基本元数据\n")

		// 使用完整的IPFS哈希，但在显示名称时可以截取一部分
		displayHash := ipfsHash
		if len(ipfsHash) > 8 {
			displayHash = ipfsHash[:8] + "..."
		}

		// 选择一个可靠的IPFS网关作为图像URL
		imageUrl := uri // 默认使用原始URI

		// IPFS网关列表
		ipfsGateways := []string{
			"https://dweb.link/ipfs/",
			"https://cloudflare-ipfs.com/ipfs/",
			"https://gateway.pinata.cloud/ipfs/",
			"https://ipfs.io/ipfs/",
			"https://cf-ipfs.com/ipfs/",
			"https://ipfs.fleek.co/ipfs/",
		}

		// 使用第一个网关作为图像URL
		if len(ipfsGateways) > 0 {
			imageUrl = ipfsGateways[0] + ipfsHash
		}

		return &models.NFTMetadata{
			Name:        "NFT #" + displayHash,
			Description: "This is an NFT created on our platform",
			Image:       imageUrl, // 使用IPFS网关URL作为图像链接
			Attributes: []models.Attribute{
				{TraitType: "Type", Value: "Main NFT"},
				{TraitType: "Rarity", Value: "Common"},
			},
		}, nil
	}

	// 如果是HTTP链接或相对路径，创建基于URI的元数据
	if strings.HasPrefix(uri, "http") || strings.HasPrefix(uri, "/") {
		// 从URI中提取最后一部分作为ID
		parts := strings.Split(uri, "/")
		idPart := ""
		if len(parts) > 0 {
			idPart = parts[len(parts)-1]
		}
		fmt.Printf("从URI提取ID部分: %s\n", idPart)

		name := "NFT"
		if idPart != "" {
			name = "NFT #" + idPart
		}

		return &models.NFTMetadata{
			Name:        name,
			Description: "This is an NFT created on our platform",
			Image:       uri,
			Attributes: []models.Attribute{
				{TraitType: "Type", Value: "Main NFT"},
				{TraitType: "Rarity", Value: "Common"},
			},
		}, nil
	}

	// 默认元数据 - 处理其他类型的URI
	fmt.Printf("使用默认元数据处理URI: %s\n", uri)

	// 不限制URI长度，但在显示时可以截取部分
	displayUri := uri
	if len(uri) > 10 {
		displayUri = uri[:5] + "..." + uri[len(uri)-5:]
	}

	return &models.NFTMetadata{
		Name:        "NFT #" + displayUri,
		Description: "This is an NFT created on our platform",
		Image:       uri,
		Attributes: []models.Attribute{
			{TraitType: "Type", Value: "Main NFT"},
			{TraitType: "Rarity", Value: "Common"},
		},
	}, nil
}

// GetNFTHandler 获取NFT信息的处理程序
func (h *NFTHandlers) GetNFTHandler(c *gin.Context) {
	// 从URL获取参数
	tokenIDStr := c.Param("tokenId")

	// 转换tokenID为big.Int
	tokenID, ok := new(big.Int).SetString(tokenIDStr, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的token ID"})
		return
	}

	// 首先尝试从数据库获取NFT信息
	var nft models.NFT
	result := models.DB.Where("token_id = ?", tokenIDStr).First(&nft)
	if result.Error == nil {
		// 从数据库找到了NFT
		c.JSON(http.StatusOK, models.NFTResponse{
			TokenID:     nft.TokenID,
			Owner:       nft.Owner,
			URI:         nft.URI,
			TotalSupply: nft.TotalSupply,
		})
		return
	}

	// 如果数据库中没有，则从区块链获取
	owner, uri, totalSupply, err := h.Client.GetNFTInfo(tokenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取NFT信息失败: " + err.Error()})
		return
	}

	// 存储到数据库
	nft = models.NFT{
		TokenID:     tokenID.String(),
		Owner:       owner,
		URI:         uri,
		TotalSupply: totalSupply,
	}
	models.DB.Create(&nft)

	// 构造响应
	response := models.NFTResponse{
		TokenID:     tokenID.String(),
		Owner:       owner,
		URI:         uri,
		TotalSupply: totalSupply,
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

	// 铸造NFT
	txHash, err := h.Client.MintNFT(walletAddress.(string), req.URI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "铸造NFT失败: " + err.Error()})
		return
	}

	// 获取当前的总供应量作为tokenID
	totalSupply, err := h.Client.GetCurrentTotalSupply()
	if err != nil {
		// 即使获取失败，也继续处理，不影响铸造结果
		c.JSON(http.StatusOK, models.TransactionResponse{
			TransactionHash: txHash,
			Message:         "NFT铸造交易已提交，但无法保存到数据库",
		})
		return
	}

	// 新铸造的NFT的tokenID应该是当前总供应量减1
	tokenID := new(big.Int).Sub(totalSupply, big.NewInt(1))

	// 立即保存到数据库
	nft := models.NFT{
		TokenID:      tokenID.String(),
		Owner:        walletAddress.(string),
		URI:          req.URI,
		IsChildNFT:   false,
		ContractType: "main",
	}
	models.DB.Create(&nft)

	// 构造响应
	response := models.TransactionResponse{
		TransactionHash: txHash,
		Message:         "NFT铸造交易已提交",
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

	// 转换tokenID为big.Int
	tokenID, ok := new(big.Int).SetString(req.TokenID, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的token ID"})
		return
	}

	var txHash string
	// var err error

	// 根据合约类型更新元数据
	if req.ContractType == "main" {
		// 验证主NFT所有权
		owner, _, _, err := h.Client.GetNFTInfo(tokenID)
		fmt.Println("owner:", owner)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取主NFT信息失败: " + err.Error()})
			return
		}
		fmt.Println("owner:", strings.ToLower(owner))
		fmt.Println("walletAddress:", strings.ToLower(walletAddress.(string)))
		// 将owner地址转换为小写后再比较
		if strings.ToLower(owner) != strings.ToLower(walletAddress.(string)) {
			c.JSON(http.StatusForbidden, gin.H{"error": "只有NFT所有者可以更新元数据"})
			return
		}

		// 更新主NFT元数据
		txHash, err = h.Client.UpdateMainNFTMetadata(tokenID, req.NewURI)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "更新主NFT元数据失败: " + err.Error()})
			return
		}

		// 更新数据库记录
		models.DB.Model(&models.NFT{}).Where("token_id = ? AND contract_type = ?", req.TokenID, "main").
			Update("uri", req.NewURI)

	} else if req.ContractType == "child" {
		// 验证子NFT所有权
		owner, _, err := h.Client.GetChildNFTInfo(tokenID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "获取子NFT信息失败: " + err.Error()})
			return
		}

		if strings.ToLower(owner) != strings.ToLower(walletAddress.(string)) {
			c.JSON(http.StatusForbidden, gin.H{"error": "只有子NFT所有者可以更新元数据"})
			return
		}

		// 更新子NFT元数据
		txHash, err = h.Client.UpdateChildNFTMetadata(tokenID, req.NewURI)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "更新子NFT元数据失败: " + err.Error()})
			return
		}

		// 更新数据库记录
		models.DB.Model(&models.NFT{}).Where("token_id = ? AND contract_type = ?", req.TokenID, "child").
			Update("uri", req.NewURI)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的合约类型"})
		return
	}

	// 构造响应
	response := models.TransactionResponse{
		TransactionHash: txHash,
		Message:         "元数据更新交易已提交",
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

	// 转换tokenID为big.Int
	tokenID, ok := new(big.Int).SetString(req.TokenID, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的token ID"})
		return
	}

	// 验证NFT所有权
	owner, _, _, err := h.Client.GetNFTInfo(tokenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取NFT信息失败: " + err.Error()})
		return
	}

	// 将owner地址转换为小写后再比较
	if strings.ToLower(owner) != strings.ToLower(walletAddress.(string)) {
		c.JSON(http.StatusForbidden, gin.H{"error": "只有NFT所有者可以更新URI"})
		return
	}

	// 更新NFT元数据URI
	txHash, err := h.Client.UpdateMainNFTMetadata(tokenID, req.NewURI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新NFT URI失败: " + err.Error()})
		return
	}

	// 更新数据库记录
	result := models.DB.Model(&models.NFT{}).Where("token_id = ?", req.TokenID).Update("uri", req.NewURI)
	if result.Error != nil {
		// 即使数据库更新失败，区块链操作已成功，也要返回成功响应
		fmt.Printf("数据库更新失败，但区块链操作成功: %v\n", result.Error)
	}

	// 构造响应
	response := models.TransactionResponse{
		TransactionHash: txHash,
		Message:         "NFT URI更新交易已提交",
	}

	// 返回JSON响应
	c.JSON(http.StatusOK, response)
}
