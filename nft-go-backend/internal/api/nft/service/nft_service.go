package service

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ABE/nft/nft-go-backend/internal/blockchain"
	"github.com/ABE/nft/nft-go-backend/internal/models"
)

// NFTService NFT业务逻辑服务
type NFTService struct {
	Client *blockchain.EthClient
}

// NewNFTService 创建新的NFT服务
func NewNFTService(client *blockchain.EthClient) *NFTService {
	return &NFTService{
		Client: client,
	}
}

// GetAllMainNFTs 获取所有主NFT
func (s *NFTService) GetAllMainNFTs() ([]models.NFTResponseWithMetadata, error) {
	var nfts []models.NFT
	// 只查询主NFT，排除子NFT
	result := models.DB.Where("is_child_nft = ? OR is_child_nft IS NULL", false).Find(&nfts)
	if result.Error != nil {
		return nil, result.Error
	}

	// 获取NFT元数据
	var nftResponses []models.NFTResponseWithMetadata
	for _, nft := range nfts {
		metadata, _ := s.FetchNFTMetadata(nft.URI)

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

	return nftResponses, nil
}

// GetUserNFTs 获取指定用户的NFT
func (s *NFTService) GetUserNFTs(address string) ([]models.NFTResponseWithMetadata, error) {
	var nfts []models.NFT
	result := models.DB.Where("owner = ?", address).Find(&nfts)
	if result.Error != nil {
		return nil, result.Error
	}

	// 获取NFT元数据
	var nftResponses []models.NFTResponseWithMetadata
	for _, nft := range nfts {
		metadata, _ := s.FetchNFTMetadata(nft.URI)

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

	return nftResponses, nil
}

// GetMyNFTs 获取当前用户的NFT (包含子NFT)
func (s *NFTService) GetMyNFTs(walletAddress string) ([]models.NFTResponseWithMetadata, error) {
	// 查询用户拥有的主NFT
	var mainNFTs []models.NFT
	result := models.DB.Where("owner = ? AND (is_child_nft = ? OR is_child_nft IS NULL)", walletAddress, false).Find(&mainNFTs)
	if result.Error != nil {
		return nil, fmt.Errorf("查询主NFT失败: %v", result.Error)
	}

	// 查询用户拥有的子NFT
	var childNFTs []models.NFT
	childResult := models.DB.Where("owner = ? AND is_child_nft = ?", walletAddress, true).Find(&childNFTs)
	if childResult.Error != nil {
		return nil, fmt.Errorf("查询子NFT失败: %v", childResult.Error)
	}

	// 合并主NFT和子NFT
	allNFTs := append(mainNFTs, childNFTs...)

	// 构建响应
	var nftResponses []models.NFTResponseWithMetadata
	for _, nft := range allNFTs {
		metadata, _ := s.FetchNFTMetadata(nft.URI)

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

	return nftResponses, nil
}

// GetNFTInfo 获取NFT信息
func (s *NFTService) GetNFTInfo(tokenIDStr string) (*models.NFTResponse, error) {
	// 转换tokenID为big.Int
	tokenID, ok := new(big.Int).SetString(tokenIDStr, 10)
	if !ok {
		return nil, fmt.Errorf("无效的token ID")
	}

	// 首先尝试从数据库获取NFT信息
	var nft models.NFT
	result := models.DB.Where("token_id = ?", tokenIDStr).First(&nft)
	if result.Error == nil {
		// 从数据库找到了NFT
		return &models.NFTResponse{
			TokenID:     nft.TokenID,
			Owner:       nft.Owner,
			URI:         nft.URI,
			TotalSupply: nft.TotalSupply,
		}, nil
	}

	// 如果数据库中没有，则从区块链获取
	owner, uri, totalSupply, err := s.Client.GetNFTInfo(tokenID)
	if err != nil {
		return nil, fmt.Errorf("获取NFT信息失败: %v", err)
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
	response := &models.NFTResponse{
		TokenID:     tokenID.String(),
		Owner:       owner,
		URI:         uri,
		TotalSupply: totalSupply,
	}

	return response, nil
}

// MintNFT 铸造NFT
func (s *NFTService) MintNFT(walletAddress, uri string) (*models.TransactionResponse, error) {
	// 铸造NFT
	txHash, err := s.Client.MintNFT(walletAddress, uri)
	if err != nil {
		return nil, fmt.Errorf("铸造NFT失败: %v", err)
	}

	// 获取当前的总供应量作为tokenID
	totalSupply, err := s.Client.GetCurrentTotalSupply()
	if err != nil {
		// 即使获取失败，也继续处理，不影响铸造结果
		return &models.TransactionResponse{
			TransactionHash: txHash,
			Message:         "NFT铸造交易已提交，但无法保存到数据库",
		}, nil
	}

	// 新铸造的NFT的tokenID应该是当前总供应量减1
	tokenID := new(big.Int).Sub(totalSupply, big.NewInt(1))

	// 立即保存到数据库
	nft := models.NFT{
		TokenID:      tokenID.String(),
		Owner:        walletAddress,
		URI:          uri,
		IsChildNFT:   false,
		ContractType: "main",
	}
	models.DB.Create(&nft)

	// 构造响应
	response := &models.TransactionResponse{
		TransactionHash: txHash,
		Message:         "NFT铸造交易已提交",
	}

	return response, nil
}

// UpdateMetadata 更新NFT元数据
func (s *NFTService) UpdateMetadata(walletAddress, tokenID, contractType, newURI string) (*models.TransactionResponse, error) {
	// 转换tokenID为big.Int
	tokenIDBigInt, ok := new(big.Int).SetString(tokenID, 10)
	if !ok {
		return nil, fmt.Errorf("无效的token ID")
	}

	var txHash string

	// 根据合约类型更新元数据
	if contractType == "main" {
		// 验证主NFT所有权
		owner, _, _, err := s.Client.GetNFTInfo(tokenIDBigInt)
		if err != nil {
			return nil, fmt.Errorf("获取主NFT信息失败: %v", err)
		}

		// 将owner地址转换为小写后再比较
		if strings.ToLower(owner) != strings.ToLower(walletAddress) {
			return nil, fmt.Errorf("只有NFT所有者可以更新元数据")
		}

		// 更新主NFT元数据
		txHash, err = s.Client.UpdateMainNFTMetadata(tokenIDBigInt, newURI)
		if err != nil {
			return nil, fmt.Errorf("更新主NFT元数据失败: %v", err)
		}

		// 更新数据库记录
		models.DB.Model(&models.NFT{}).Where("token_id = ? AND contract_type = ?", tokenID, "main").
			Update("uri", newURI)

	} else if contractType == "child" {
		// 验证子NFT所有权
		owner, _, err := s.Client.GetChildNFTInfo(tokenIDBigInt)
		if err != nil {
			return nil, fmt.Errorf("获取子NFT信息失败: %v", err)
		}

		if strings.ToLower(owner) != strings.ToLower(walletAddress) {
			return nil, fmt.Errorf("只有子NFT所有者可以更新元数据")
		}

		// 更新子NFT元数据
		txHash, err = s.Client.UpdateChildNFTMetadata(tokenIDBigInt, newURI)
		if err != nil {
			return nil, fmt.Errorf("更新子NFT元数据失败: %v", err)
		}

		// 更新数据库记录
		models.DB.Model(&models.NFT{}).Where("token_id = ? AND contract_type = ?", tokenID, "child").
			Update("uri", newURI)

	} else {
		return nil, fmt.Errorf("无效的合约类型")
	}

	// 构造响应
	response := &models.TransactionResponse{
		TransactionHash: txHash,
		Message:         "元数据更新交易已提交",
	}

	return response, nil
}

// UpdateNFTURI 更新NFT元数据URI
func (s *NFTService) UpdateNFTURI(walletAddress, tokenID, newURI string) (*models.TransactionResponse, error) {
	// 转换tokenID为big.Int
	tokenIDBigInt, ok := new(big.Int).SetString(tokenID, 10)
	if !ok {
		return nil, fmt.Errorf("无效的token ID")
	}

	// 验证NFT所有权
	owner, _, _, err := s.Client.GetNFTInfo(tokenIDBigInt)
	if err != nil {
		return nil, fmt.Errorf("获取NFT信息失败: %v", err)
	}

	// 将owner地址转换为小写后再比较
	if strings.ToLower(owner) != strings.ToLower(walletAddress) {
		return nil, fmt.Errorf("只有NFT所有者可以更新URI")
	}

	// 更新NFT元数据URI
	txHash, err := s.Client.UpdateMainNFTMetadata(tokenIDBigInt, newURI)
	if err != nil {
		return nil, fmt.Errorf("更新NFT URI失败: %v", err)
	}

	// 更新数据库记录
	result := models.DB.Model(&models.NFT{}).Where("token_id = ?", tokenID).Update("uri", newURI)
	if result.Error != nil {
		// 即使数据库更新失败，区块链操作已成功，也要返回成功响应
		fmt.Printf("数据库更新失败，但区块链操作成功: %v\n", result.Error)
	}

	// 构造响应
	response := &models.TransactionResponse{
		TransactionHash: txHash,
		Message:         "NFT URI更新交易已提交",
	}

	return response, nil
}

// FetchNFTMetadata 获取NFT元数据
func (s *NFTService) FetchNFTMetadata(uri string) (*models.NFTMetadata, error) {
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
		return s.handleIPFSMetadata(uri)
	}

	// 如果是HTTP链接或相对路径，创建基于URI的元数据
	if strings.HasPrefix(uri, "http") || strings.HasPrefix(uri, "/") {
		return s.handleHTTPMetadata(uri)
	}

	// 默认元数据 - 处理其他类型的URI
	return s.handleDefaultMetadata(uri)
}

// handleIPFSMetadata 处理IPFS类型的元数据
func (s *NFTService) handleIPFSMetadata(uri string) (*models.NFTMetadata, error) {
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

// handleHTTPMetadata 处理HTTP类型的元数据
func (s *NFTService) handleHTTPMetadata(uri string) (*models.NFTMetadata, error) {
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

// handleDefaultMetadata 处理默认类型的元数据
func (s *NFTService) handleDefaultMetadata(uri string) (*models.NFTMetadata, error) {
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