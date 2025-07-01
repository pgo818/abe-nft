package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/ABE/nft/nft-go-backend/internal/blockchain"
	"github.com/ABE/nft/nft-go-backend/internal/models"
	abe "github.com/ABE/nft/nft-go-backend/internal/api/abe/service"
)

// ChildNFTHandlers 子NFT相关处理程序结构体
type ChildNFTHandlers struct {
	Client *blockchain.EthClient
}

// NewChildNFTHandlers 创建新的子NFT处理程序
func NewChildNFTHandlers(client *blockchain.EthClient) *ChildNFTHandlers {
	return &ChildNFTHandlers{
		Client: client,
	}
}

// CreateChildNFTHandler 创建子NFT的处理程序
func (h *ChildNFTHandlers) CreateChildNFTHandler(c *gin.Context) {
	// 获取验证后的钱包地址
	walletAddressInterface, exists := c.Get("walletAddress")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经过身份验证"})
		return
	}

	walletAddress, ok := walletAddressInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "钱包地址格式错误"})
		return
	}

	// 将钱包地址转换为小写
	walletAddressLower := strings.ToLower(walletAddress)

	// 解析请求体
	var req models.CreateChildRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 验证签名请求地址与钱包地址匹配
	if strings.ToLower(req.Address) != walletAddressLower {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求地址与签名地址不匹配"})
		return
	}

	// 转换parentTokenID为big.Int
	parentTokenID, ok := new(big.Int).SetString(req.ParentTokenID, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的父token ID"})
		return
	}

	// 验证父NFT所有权
	owner, _, _, err := h.Client.GetNFTInfo(parentTokenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取NFT信息失败: " + err.Error()})
		return
	}
	fmt.Println("owner:", owner)
	fmt.Println("walletAddress:", walletAddress)

	// 将owner地址转换为小写后再比较
	if strings.ToLower(owner) != walletAddressLower {
		c.JSON(http.StatusForbidden, gin.H{"error": "只有NFT所有者可以创建子NFT"})
		return
	}

	// 转换recipient为以太坊地址
	recipient := common.HexToAddress(req.Recipient)

	// 创建子NFT
	txHash, err := h.Client.CreateChildNFT(parentTokenID, walletAddressLower, recipient, req.URI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建子NFT失败: " + err.Error()})
		return
	}

	// 构造响应
	response := models.TransactionResponse{
		TransactionHash: txHash,
		Message:         "子NFT创建交易已提交",
	}

	// 返回JSON响应
	c.JSON(http.StatusOK, response)
}

// RequestChildNFTHandler 申请子NFT的处理程序
func (h *ChildNFTHandlers) RequestChildNFTHandler(c *gin.Context) {
	// 获取验证后的钱包地址
	walletAddressInterface, exists := c.Get("walletAddress")
	if !exists {
		fmt.Printf("无法获取钱包地址，用户未经过身份验证\n")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经过身份验证"})
		return
	}

	walletAddress, ok := walletAddressInterface.(string)
	if !ok {
		fmt.Printf("钱包地址类型转换失败\n")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "钱包地址格式错误"})
		return
	}

	fmt.Printf("申请子NFT，钱包地址: %s\n", walletAddress)

	// 解析请求体
	var req models.RequestChildNFTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("解析请求体失败: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	fmt.Printf("请求内容: ParentTokenId=%s, ApplicantAddress=%s, URI=%s, AutoApprove=%v\n",
		req.ParentTokenId, req.ApplicantAddress, req.URI, req.AutoApprove)

	// 验证签名请求地址与钱包地址匹配
	if req.Address != walletAddress {
		fmt.Printf("地址不匹配: req.Address=%s, walletAddress=%s\n", req.Address, walletAddress)
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求地址与签名地址不匹配"})
		return
	}

	// 验证父NFT是否存在
	parentTokenID, ok := new(big.Int).SetString(req.ParentTokenId, 10)
	if !ok {
		fmt.Printf("无效的父token ID: %s\n", req.ParentTokenId)
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的父token ID"})
		return
	}

	// 验证父NFT所有者
	owner, _, _, err := h.Client.GetNFTInfo(parentTokenID)
	if err != nil {
		fmt.Printf("获取NFT信息失败: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取NFT信息失败: " + err.Error()})
		return
	}

	fmt.Printf("父NFT所有者: %s\n", owner)

	// 创建申请记录
	request := models.ChildNFTRequest{
		ParentTokenId:    req.ParentTokenId,
		ApplicantAddress: req.ApplicantAddress,
		URI:              req.URI,
		Status:           "pending",
		VCCredentials:    req.VCCredentials,
		AutoApproved:     false,
	}

	// 如果提供了VC凭证且要求自动审核，尝试进行策略验证
	var autoApproved bool = false
	var policyResult map[string]interface{}

	if req.AutoApprove && req.VCCredentials != "" {
		fmt.Printf("开始自动审核流程...\n")

		// 获取父NFT的访问策略
		accessPolicy, err := h.getAccessPolicyForNFT(req.ParentTokenId)
		if err != nil {
			fmt.Printf("获取NFT访问策略失败: %v\n", err)
		} else if accessPolicy != "" {
			fmt.Printf("NFT访问策略: %s\n", accessPolicy)

			// 创建ABE服务实例进行策略验证
			abeService := abe.NewABEService(models.DB)

			// 验证VC凭证是否满足策略
			satisfied, verificationResult, err := abeService.VerifyVCAgainstPolicy(req.VCCredentials, accessPolicy)
			if err != nil {
				fmt.Printf("策略验证过程出错: %v\n", err)
				policyResult = map[string]interface{}{
					"error":    err.Error(),
					"verified": false,
				}
			} else {
				policyResult = verificationResult

				if satisfied {
					fmt.Printf("VC凭证满足访问策略，自动审核通过\n")
					autoApproved = true
					request.Status = "auto_approved"
					request.AutoApproved = true

					// 自动创建子NFT
					recipientAddr := common.HexToAddress(req.ApplicantAddress)
					txHash, err := h.Client.CreateChildNFT(parentTokenID, owner, recipientAddr, req.URI)
					if err != nil {
						fmt.Printf("自动创建子NFT失败: %v\n", err)
						// 如果创建失败，回退到手动审核
						autoApproved = false
						request.Status = "pending"
						request.AutoApproved = false
						policyResult["auto_creation_error"] = err.Error()
					} else {
						fmt.Printf("自动创建子NFT成功，交易哈希: %s\n", txHash)
						request.ChildTokenID = fmt.Sprintf("auto_%d", time.Now().Unix())
						policyResult["transaction_hash"] = txHash

						// 将子NFT信息保存到NFT表中
						childNFT := models.NFT{
							TokenID:       request.ChildTokenID,
							Owner:         req.ApplicantAddress,
							URI:           req.URI,
							IsChildNFT:    true,
							ParentTokenID: req.ParentTokenId,
							ContractType:  "child",
						}
						models.DB.Create(&childNFT)
					}
				} else {
					fmt.Printf("VC凭证不满足访问策略，需要手动审核\n")
					policyResult["reason"] = "VC凭证不满足访问策略要求"
				}
			}
		} else {
			fmt.Printf("父NFT没有设置访问策略，回退到手动审核\n")
			policyResult = map[string]interface{}{
				"reason":                 "父NFT没有设置访问策略",
				"manual_review_required": true,
			}
		}
	}

	// 保存策略验证结果
	if policyResult != nil {
		policyResultJSON, _ := json.Marshal(policyResult)
		request.PolicyResult = string(policyResultJSON)
	}

	fmt.Printf("保存申请记录到数据库...\n")
	result := models.DB.Create(&request)
	if result.Error != nil {
		fmt.Printf("保存申请记录失败: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存申请记录失败: " + result.Error.Error()})
		return
	}

	fmt.Printf("申请记录保存成功，ID: %d\n", request.ID)

	// 构建响应
	response := gin.H{
		"requestId":    request.ID,
		"ownerAddress": owner,
		"autoApproved": autoApproved,
	}

	if autoApproved {
		response["message"] = "VC凭证验证通过，子NFT申请已自动审核并创建"
		response["transactionHash"] = policyResult["transaction_hash"]
		response["childTokenId"] = request.ChildTokenID
	} else {
		response["message"] = "子NFT申请已提交，等待父NFT持有者审批"
		if policyResult != nil {
			response["policyResult"] = policyResult
		}
	}

	c.JSON(http.StatusOK, response)
}

// getAccessPolicyForNFT 获取NFT的访问策略
func (h *ChildNFTHandlers) getAccessPolicyForNFT(tokenId string) (string, error) {
	// 第一步：根据token ID查询NFT记录，获取URI
	var nft models.NFT
	nftResult := models.DB.Where("token_id = ?", tokenId).First(&nft)
	if nftResult.Error != nil {
		fmt.Printf("Token %s 在NFT表中未找到记录: %v\n", tokenId, nftResult.Error)
		return "", fmt.Errorf("NFT %s 不存在", tokenId)
	}

	fmt.Printf("Token %s 找到NFT记录，URI: %s\n", tokenId, nft.URI)

	// 第二步：处理IPFS hash格式差异
	// NFT表存储格式：ipfs://QmVwgTpVvE56Nmx1YhFzzndwNUYw96xdmZcZ6esPzDbzfU
	// 元数据表存储格式：QmVwgTpVvE56Nmx1YhFzzndwNUYw96xdmZcZ6esPzDbzfU
	ipfsHash := nft.URI
	pureHash := ipfsHash

	// 如果URI以ipfs://开头，去掉前缀
	if strings.HasPrefix(ipfsHash, "ipfs://") {
		pureHash = strings.TrimPrefix(ipfsHash, "ipfs://")
		fmt.Printf("Token %s 转换IPFS hash: %s -> %s\n", tokenId, ipfsHash, pureHash)
	}

	// 第三步：尝试两种格式查询元数据表
	var metadata models.NFTMetadataDB
	var metadataResult *gorm.DB

	// 先尝试用纯hash查询
	metadataResult = models.DB.Where("ip_fs_hash = ?", pureHash).First(&metadata)
	if metadataResult.Error != nil {
		fmt.Printf("用纯hash %s 查询失败，尝试用完整URI %s 查询\n", pureHash, ipfsHash)
		// 再尝试用完整URI查询
		metadataResult = models.DB.Where("ip_fs_hash = ?", ipfsHash).First(&metadata)
		if metadataResult.Error != nil {
			fmt.Printf("Token %s 的URI在元数据表中未找到记录。尝试的hash格式：%s 和 %s\n", tokenId, pureHash, ipfsHash)
			return "", fmt.Errorf("NFT %s 还没有创建元数据和访问策略，请先在\"创建元数据\"页面为此NFT创建元数据", tokenId)
		}
	}

	fmt.Printf("Token %s 成功找到元数据记录\n", tokenId)

	// 第四步：检查策略是否为空
	if metadata.Policy == "" {
		fmt.Printf("Token %s 的元数据中没有设置访问策略\n", tokenId)
		return "", fmt.Errorf("NFT %s 的元数据中没有设置访问策略，请更新元数据并设置访问策略", tokenId)
	}

	// 返回现有的访问策略
	fmt.Printf("Token %s 找到访问策略: %s\n", tokenId, metadata.Policy)
	return metadata.Policy, nil
}

// GetAllRequestsHandler 获取所有的子NFT申请记录
func (h *ChildNFTHandlers) GetAllRequestsHandler(c *gin.Context) {
	// 获取验证后的钱包地址
	walletAddressInterface, exists := c.Get("walletAddress")
	if !exists {
		fmt.Printf("无法获取钱包地址，用户未经过身份验证\n")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经过身份验证"})
		return
	}

	walletAddress, ok := walletAddressInterface.(string)
	if !ok {
		fmt.Printf("钱包地址类型转换失败\n")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "钱包地址格式错误"})
		return
	}

	fmt.Printf("处理请求列表查询，钱包地址: %s\n", walletAddress)

	// 查询该地址拥有的所有NFT
	var nfts []models.NFT
	result := models.DB.Where("owner = ? AND (is_child_nft = ? OR is_child_nft IS NULL)", walletAddress, false).Find(&nfts)
	if result.Error != nil {
		fmt.Printf("查询NFT失败: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询NFT失败: " + result.Error.Error()})
		return
	}
	fmt.Printf("找到用户拥有的NFT数量: %d\n", len(nfts))

	// 收集所有tokenId
	var tokenIds []string
	for _, nft := range nfts {
		tokenIds = append(tokenIds, nft.TokenID)
		fmt.Printf("用户拥有的NFT: TokenID=%s\n", nft.TokenID)
	}

	// 查询所有申请记录 - 这里只查询与用户相关的记录
	var allRequests []models.ChildNFTRequest
	// 查询自己申请的或者申请自己NFT的记录
	models.DB.Where("applicant_address = ? OR parent_token_id IN (?)", walletAddress, tokenIds).Find(&allRequests)
	fmt.Printf("找到与用户相关的申请记录数量: %d\n", len(allRequests))

	// 构建响应数据
	var responseRequests []models.ChildNFTRequestWithParentInfo

	// 处理所有申请记录
	for _, req := range allRequests {
		// 检查是否是申请自己的NFT
		isMyNFT := false
		for _, tokenId := range tokenIds {
			if tokenId == req.ParentTokenId {
				isMyNFT = true
				break
			}
		}

		// 设置RequestId字段为ID的字符串形式
		req.RequestId = fmt.Sprintf("%d", req.ID)

		// 调试输出
		fmt.Printf("处理请求 ID=%d, RequestId=%s, 状态=%s\n", req.ID, req.RequestId, req.Status)

		var parentNFTOwner string
		var canOperate bool

		if isMyNFT {
			// 这是别人申请自己的NFT
			fmt.Printf("找到申请用户NFT的请求: ID=%d, ParentTokenId=%s, 申请者=%s\n",
				req.ID, req.ParentTokenId, req.ApplicantAddress)
			parentNFTOwner = walletAddress
			canOperate = true // 作为父NFT拥有者可以操作
		} else {
			// 这是自己申请的NFT
			fmt.Printf("找到用户自己的申请: ID=%d, ParentTokenId=%s\n", req.ID, req.ParentTokenId)
			// 获取父NFT拥有者信息
			var parentNFT models.NFT
			models.DB.Where("token_id = ?", req.ParentTokenId).First(&parentNFT)
			parentNFTOwner = parentNFT.Owner
			canOperate = false // 作为申请者不能操作，只能查看
		}

		responseReq := models.ChildNFTRequestWithParentInfo{
			ChildNFTRequest: req,
			ParentNFTOwner:  parentNFTOwner,
			CanOperate:      canOperate,
		}

		// 确保ID字段正确设置
		fmt.Printf("响应请求对象: ID=%d, RequestId=%s, CanOperate=%v\n",
			responseReq.ID, responseReq.RequestId, responseReq.CanOperate)

		responseRequests = append(responseRequests, responseReq)
	}

	fmt.Printf("返回给客户端的请求数量: %d\n", len(responseRequests))

	// 始终返回一个数组，即使为空
	c.JSON(http.StatusOK, gin.H{"requests": responseRequests})
}

// ProcessRequestHandler 处理子NFT申请
func (h *ChildNFTHandlers) ProcessRequestHandler(c *gin.Context) {
	// 获取验证后的钱包地址
	walletAddressInterface, exists := c.Get("walletAddress")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未经过身份验证"})
		return
	}

	// 转换钱包地址为小写格式
	walletAddress, ok := walletAddressInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "钱包地址格式错误"})
		return
	}
	normalizedWallet := normalizeAddress(walletAddress)

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

	// 首先解析请求体为一个通用的map，以便我们可以处理不同类型的requestId
	var rawRequest map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &rawRequest); err != nil {
		fmt.Printf("解析请求体为map错误: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 提取并转换requestId为uint
	var requestID uint
	if reqIDVal, exists := rawRequest["requestId"]; exists {
		switch v := reqIDVal.(type) {
		case float64:
			// JSON数字会被解析为float64
			requestID = uint(v)
		case string:
			// 如果是字符串，尝试转换为uint
			if id, err := strconv.ParseUint(v, 10, 32); err == nil {
				requestID = uint(id)
			} else {
				fmt.Printf("转换requestId错误: %v\n", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求ID格式"})
				return
			}
		default:
			fmt.Printf("不支持的requestId类型: %T\n", v)
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求ID类型"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少请求ID"})
		return
	}

	fmt.Printf("解析后的请求ID: %d\n", requestID)

	// 提取action字段
	action, ok := rawRequest["action"].(string)
	if !ok || (action != "approve" && action != "reject") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的action值，必须是approve或reject"})
		return
	}

	// 提取签名相关字段
	address, ok := rawRequest["address"].(string)
	if !ok || address == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少或无效的address字段"})
		return
	}

	signature, ok := rawRequest["signature"].(string)
	if !ok || signature == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少或无效的signature字段"})
		return
	}

	message, ok := rawRequest["message"].(string)
	if !ok || message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少或无效的message字段"})
		return
	}

	// 验证签名请求地址与钱包地址匹配（转换为小写后比较）
	reqAddress := normalizeAddress(address)
	if reqAddress != normalizedWallet {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求地址与签名地址不匹配"})
		return
	}

	// 查询申请记录
	var request models.ChildNFTRequest
	result := models.DB.First(&request, requestID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "申请记录不存在"})
		return
	}

	// 验证NFT所有权
	parentTokenID, _ := new(big.Int).SetString(request.ParentTokenId, 10)
	owner, _, _, err := h.Client.GetNFTInfo(parentTokenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取NFT信息失败: " + err.Error()})
		return
	}
	fmt.Println("owner:", owner)
	fmt.Println("walletAddress:", walletAddress)

	// 规范化owner地址为小写
	normalizedOwner := normalizeAddress(owner)

	// 验证当前用户是否为NFT所有者（使用小写地址比较）
	if normalizedWallet != normalizedOwner {
		c.JSON(http.StatusForbidden, gin.H{"error": "只有NFT所有者可以处理申请"})
		return
	}

	if action == "approve" {
		// 批准申请，创建子NFT
		recipient := common.HexToAddress(request.ApplicantAddress)
		txHash, err := h.Client.CreateChildNFT(parentTokenID, walletAddress, recipient, request.URI)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "创建子NFT失败: " + err.Error()})
			return
		}

		// 更新申请状态并记录子NFT的tokenID
		request.Status = "approved"
		// 这里需要获取新创建的子NFT的tokenID，暂时使用时间戳作为占位符
		// 在实际应用中，应该从交易回执中获取实际的tokenID
		request.ChildTokenID = fmt.Sprintf("child_%d", time.Now().Unix())
		models.DB.Save(&request)

		// 将子NFT信息保存到NFT表中
		childNFT := models.NFT{
			TokenID:       request.ChildTokenID,
			Owner:         request.ApplicantAddress,
			URI:           request.URI,
			IsChildNFT:    true,
			ParentTokenID: request.ParentTokenId,
			ContractType:  "child",
		}
		models.DB.Create(&childNFT)

		c.JSON(http.StatusOK, gin.H{
			"message":         "申请已批准，子NFT已创建",
			"transactionHash": txHash,
			"childTokenId":    request.ChildTokenID,
		})
	} else {
		// 拒绝申请
		request.Status = "rejected"
		models.DB.Save(&request)

		c.JSON(http.StatusOK, gin.H{
			"message": "申请已拒绝",
		})
	}
}

// normalizeAddress 将以太坊地址转换为小写格式（去除0x前缀后转小写，再添加0x前缀）
func normalizeAddress(addr string) string {
	if len(addr) < 4 {
		return addr
	}
	// 处理0x前缀
	if addr[:2] == "0x" || addr[:2] == "0X" {
		return "0x" + strings.ToLower(addr[2:])
	}
	return strings.ToLower(addr)
}
