package api

import (
	"github.com/gin-gonic/gin"

	"nft-go-backend/internal/blockchain"
)

// Router 主路由结构体
type Router struct {
	NFTHandlers      *NFTHandlers
	ChildNFTHandlers *ChildNFTHandlers
	MetadataHandlers *MetadataHandlers
}

// NewRouter 创建新的路由实例
func NewRouter(client *blockchain.EthClient) *Router {
	return &Router{
		NFTHandlers:      NewNFTHandlers(client),
		ChildNFTHandlers: NewChildNFTHandlers(client),
		MetadataHandlers: NewMetadataHandlers(client),
	}
}

// SetupRoutes 设置所有路由
func (router *Router) SetupRoutes(r *gin.Engine) {
	// 创建API路由组
	api := r.Group("/api")

	// 不需要签名验证的路由
	api.GET("/nft/:tokenId", router.NFTHandlers.GetNFTHandler)
	api.GET("/nfts", router.NFTHandlers.GetAllNFTsHandler)
	api.GET("/nfts/user/:address", router.NFTHandlers.GetUserNFTsHandler)

	// 元数据相关路由（不需要认证）
	api.POST("/metadata", router.MetadataHandlers.CreateMetadataHandler)
	api.GET("/metadata/:hash", router.MetadataHandlers.GetMetadataHandler)
	api.GET("/metadata", router.MetadataHandlers.GetAllMetadataHandler)

	// 需要签名验证的路由
	secured := api.Group("")
	secured.Use(SignatureAuthMiddleware())
	{
		// NFT相关
		secured.POST("/nft/mint", router.NFTHandlers.MintNFTHandler)
		secured.POST("/nft/update-metadata", router.NFTHandlers.UpdateMetadataHandler)

		// 子NFT相关
		secured.POST("/nft/createChild", router.ChildNFTHandlers.CreateChildNFTHandler)
		secured.POST("/nft/request-child", router.ChildNFTHandlers.RequestChildNFTHandler)
		secured.POST("/nft/process-request", router.ChildNFTHandlers.ProcessRequestHandler)
	}

	// 需要GET请求认证的路由
	apiAuth := r.Group("/api")
	apiAuth.Use(GetRequestAuthMiddleware())
	{
		// NFT相关
		apiAuth.GET("/nft/my-nfts", router.NFTHandlers.GetMyNFTsHandler)

		// 子NFT相关
		apiAuth.GET("/nft/all-requests", router.ChildNFTHandlers.GetAllRequestsHandler)
	}
}
