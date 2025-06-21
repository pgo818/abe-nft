package api

import (
	"github.com/ABE/nft/nft-go-backend/internal/blockchain"
	"github.com/ABE/nft/nft-go-backend/internal/models"
	"github.com/ABE/nft/nft-go-backend/internal/service"

	"github.com/gin-gonic/gin"
)

// Handlers 是API处理程序的集合
type Handlers struct {
	Router *Router
}

// SetupRoutes 设置API路由
func (h *Handlers) SetupRoutes(r *gin.Engine) {
	h.Router.SetupRoutes(r)
}

// NewHandlers 创建所有API处理程序
func NewHandlers(client *blockchain.EthClient) *Router {
	// 获取数据库连接
	db := models.GetDB()

	// 创建ABE服务
	abeService := NewABEService(db)

	// 创建DID服务
	didService := service.NewDIDService(db)

	// 创建VC服务
	vcService := service.NewVCService(db)

	// 创建路由
	router := &Router{
		NFTHandlers:      NewNFTHandlers(client),
		ChildNFTHandlers: NewChildNFTHandlers(client),
		MetadataHandlers: NewMetadataHandlers(client),
		ABEHandlers:      NewABEHandlers(abeService),
		DIDHandlers:      NewDIDHandlers(didService),
		VCHandlers:       NewVCHandlers(vcService,didService),
	}

	return router
}
