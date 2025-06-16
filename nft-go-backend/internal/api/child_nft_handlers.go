package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	"nft-go-backend/internal/blockchain"
	"nft-go-backend/internal/models"
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
	walletAddress, _ := h.Client.GetWalletAddressFromContext(c)

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
	txHash, err := h.Client.CreateChildNFT(parentTokenID,walletAddressLower,recipient, req.URI)
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
	walletAddress, _ := h.Client.GetWalletAddressFromContext(c)

	// 解析请求体
	var req models.RequestChildNFTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 验证签名请求地址与钱包地址匹配
	if req.Address != walletAddress {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求地址与签名地址不匹配"})
		return
	}

	// 验证父NFT是否存在
	parentTokenID, ok := new(big.Int).SetString(req.ParentTokenId, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的父token ID"})
		return
	}

	// 验证父NFT所有者
	owner, _, _, err := h.Client.GetNFTInfo(parentTokenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取NFT信息失败: " + err.Error()})
		return
	}

	// 保存申请记录到数据库
	request := models.ChildNFTRequest{
		ParentTokenId:    req.ParentTokenId,
		ApplicantAddress: req.ApplicantAddress,
		URI:              req.URI,
		Status:           "pending",
	}

	result := models.DB.Create(&request)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存申请记录失败: " + result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "子NFT申请已提交，等待父NFT持有者审批",
		"requestId":    request.ID,
		"ownerAddress": owner,
	})
}

// GetAllRequestsHandler 获取所有的子NFT申请记录
func (h *ChildNFTHandlers) GetAllRequestsHandler(c *gin.Context) {
	// 获取验证后的钱包地址
	walletAddress, _ := h.Client.GetWalletAddressFromContext(c)

	// 查询该地址拥有的所有NFT
	var nfts []models.NFT
	result := models.DB.Where("owner = ? AND (is_child_nft = ? OR is_child_nft IS NULL)", walletAddress, false).Find(&nfts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询NFT失败: " + result.Error.Error()})
		return
	}

	// 收集所有tokenId
	var tokenIds []string
	for _, nft := range nfts {
		tokenIds = append(tokenIds, nft.TokenID)
	}

	// 查询所有申请记录
	var allRequests []models.ChildNFTRequest
	models.DB.Find(&allRequests)

	// 查询自己申请的子NFT
	var myRequests []models.ChildNFTRequest
	models.DB.Where("applicant_address = ?", walletAddress).Find(&myRequests)

	// 构建响应数据
	var responseRequests []models.ChildNFTRequestWithParentInfo

	// 处理别人申请自己NFT的请求
	for _, req := range allRequests {
		// 检查是否是申请自己的NFT
		isMyNFT := false
		for _, tokenId := range tokenIds {
			if tokenId == req.ParentTokenId {
				isMyNFT = true
				break
			}
		}

		if isMyNFT {
			responseReq := models.ChildNFTRequestWithParentInfo{
				ChildNFTRequest: req,
				ParentNFTOwner:  walletAddress,
				CanOperate:      true, // 作为父NFT拥有者可以操作
			}
			responseRequests = append(responseRequests, responseReq)
		}
	}

	// 处理自己申请的子NFT
	for _, req := range myRequests {
		// 获取父NFT拥有者信息
		var parentNFT models.NFT
		models.DB.Where("token_id = ?", req.ParentTokenId).First(&parentNFT)

		responseReq := models.ChildNFTRequestWithParentInfo{
			ChildNFTRequest: req,
			ParentNFTOwner:  parentNFT.Owner,
			CanOperate:      false, // 作为申请者不能操作，只能查看
		}
		responseRequests = append(responseRequests, responseReq)
	}

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

	// 解析请求体
	var req struct {
		models.SignedRequest
		RequestID uint   `json:"requestId" binding:"required"`
		Action    string `json:"action" binding:"required,oneof=approve reject"`
	}
	if err := json.Unmarshal(bodyBytes, &req); err != nil {
		fmt.Printf("解析请求体错误: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}
	fmt.Printf("解析后的请求体: %+v\n", req)

	// 验证签名请求地址与钱包地址匹配（转换为小写后比较）
	reqAddress := normalizeAddress(req.Address)
	if reqAddress != normalizedWallet {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求地址与签名地址不匹配"})
		return
	}

	// 查询申请记录
	var request models.ChildNFTRequest
	result := models.DB.First(&request, req.RequestID)
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

	if req.Action == "approve" {
		// 批准申请，创建子NFT
		recipient := common.HexToAddress(request.ApplicantAddress)
		txHash, err := h.Client.CreateChildNFT(parentTokenID,walletAddressInterface.(string), recipient, request.URI)
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
