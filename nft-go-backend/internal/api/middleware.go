// 在internal/api/middleware.go中创建新文件
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
)

// GetRequestAuthMiddleware GET请求的签名认证中间件
func GetRequestAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取签名信息
		address := c.GetHeader("X-Ethereum-Address")
		signature := c.GetHeader("X-Ethereum-Signature")
		message := c.GetHeader("X-Ethereum-Message")

		if address == "" || signature == "" || message == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "缺少认证信息"})
			c.Abort()
			return
		}

		// 验证地址格式
		if !common.IsHexAddress(address) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的以太坊地址格式"})
			c.Abort()
			return
		}

		// 验证签名
		if !verifySignature(address, signature, message) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "签名验证失败"})
			c.Abort()
			return
		}

		fmt.Printf("GET请求签名验证成功，地址: %s\n", address)

		// 验证通过，保存地址到上下文
		c.Set("walletAddress", address)
		c.Next()
	}
}

// SignatureAuthMiddleware 验证以太坊签名的中间件
func SignatureAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 保存原始请求体
		bodyBytes, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "读取请求体失败"})
			c.Abort()
			return
		}

		// 重新设置请求体以便后续解析
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// 只解析签名验证所需的字段
		var authReq struct {
			Address   string `json:"address"`
			Signature string `json:"signature"`
			Message   string `json:"message"`
		}

		if err := json.Unmarshal(bodyBytes, &authReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求格式: " + err.Error()})
			c.Abort()
			return
		}

		// 验证地址格式
		if !common.IsHexAddress(authReq.Address) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的以太坊地址格式"})
			c.Abort()
			return
		}

		// 验证签名
		if !verifySignature(authReq.Address, authReq.Signature, authReq.Message) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "签名验证失败"})
			c.Abort()
			return
		}

		fmt.Printf("签名验证成功，地址: %s\n", authReq.Address)

		// 验证通过，保存地址到上下文
		c.Set("walletAddress", authReq.Address)

		// 保存完整请求体到上下文，供后续处理函数使用
		c.Set("rawRequestBody", bodyBytes)

		c.Next()
	}
}

// 验证以太坊签名
func verifySignature(address, signature, message string) bool {
	// 1. 将消息转换为以太坊可验证的格式
	msgHash := signHash([]byte(message))

	// 检查地址格式
	if !common.IsHexAddress(address) {
		return false
	}
	address = common.HexToAddress(address).Hex() // 标准化地址格式

	// 2. 从签名中提取数据
	sig := hexutil.MustDecode(signature)

	if len(sig) != 65 {
		fmt.Println("签名长度不正确")
		return false
	}

	// 检查v值并可能调整
	v := sig[64]
	if v > 1 {
		sig[64] -= 27
	}

	// 3. 使用以太坊的函数恢复出签名者地址
	sigPublicKey, err := crypto.Ecrecover(msgHash, sig)
	if err != nil {
		fmt.Println("公钥恢复错误:", err)
		return false
	}

	// 4. 从公钥派生出地址
	var addr common.Address
	copy(addr[:], crypto.Keccak256(sigPublicKey[1:])[12:])
	// 5. 比较恢复出的地址与提供的地址
	return strings.EqualFold(addr.Hex(), address)
}

// 计算消息哈希
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	fmt.Println("哈希前的消息:", msg)
	return crypto.Keccak256([]byte(msg))
}
