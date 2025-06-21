package api

import (
	"github.com/gin-gonic/gin"

	"github.com/ABE/nft/nft-go-backend/internal/blockchain"
	"github.com/ABE/nft/nft-go-backend/internal/models"
	"github.com/ABE/nft/nft-go-backend/internal/service"
)

// Router 主路由结构体
type Router struct {
	NFTHandlers      *NFTHandlers
	ChildNFTHandlers *ChildNFTHandlers
	MetadataHandlers *MetadataHandlers
	ABEHandlers      *ABEHandlers
	DIDHandlers      *DIDHandlers
	VCHandlers       *VCHandlers
}

// NewRouter 创建新的路由实例
func NewRouter(client *blockchain.EthClient) *Router {
	// 获取数据库连接
	db := models.GetDB()
	abeService := NewABEService(db)

	// 创建DID服务
	didService := service.NewDIDService(db)
	// 创建VC服务
	vcService := service.NewVCService(db)

	return &Router{
		NFTHandlers:      NewNFTHandlers(client),
		ChildNFTHandlers: NewChildNFTHandlers(client),
		MetadataHandlers: NewMetadataHandlers(client),
		ABEHandlers:      NewABEHandlers(abeService),
		DIDHandlers:      NewDIDHandlers(didService),
		VCHandlers:       NewVCHandlers(vcService, didService),
	}
}

// SetupRoutes 设置所有路由
func (router *Router) SetupRoutes(r *gin.Engine) {
	// 创建API路由组
	api := r.Group("/api")

	// 静态文件路径和模板已在main.go中设置

	// 前端页面路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "首页",
		})
	})

	// 医生DID和VC页面
	r.GET("/doctor-did", func(c *gin.Context) {
		c.HTML(200, "doctor_did.html", gin.H{
			"title": "医生DID和可验证凭证",
		})
	})

	// 健康检查
	api.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 不需要签名验证的路由
	api.GET("/nft/:tokenId", router.NFTHandlers.GetNFTHandler)
	api.GET("/nfts", router.NFTHandlers.GetAllNFTsHandler)
	api.GET("/nfts/user/:address", router.NFTHandlers.GetUserNFTsHandler)

	// 元数据相关路由（不需要认证）
	api.POST("/metadata", router.MetadataHandlers.CreateMetadataHandler)
	api.GET("/metadata/:hash", router.MetadataHandlers.GetMetadataHandler)
	api.GET("/metadata", router.MetadataHandlers.GetAllMetadataHandler)

	// IPFS相关路由
	ipfs := api.Group("/ipfs")
	{
		ipfs.POST("/upload", router.MetadataHandlers.UploadToIPFSHandler)
	}

	// ABE相关路由（不需要认证，用于测试）
	abe := api.Group("/abe")
	{
		abe.POST("/setup", router.ABEHandlers.SetupABE)
		abe.POST("/keygen", router.ABEHandlers.KeyGenABE)
		abe.POST("/encrypt", router.ABEHandlers.EncryptABE)
		abe.POST("/decrypt", router.ABEHandlers.DecryptABE)
	}

	// DID路由 - 仅支持钱包地址创建和管理
	did := api.Group("/did")
	{
		// 通用DID操作
		did.GET("/list", router.DIDHandlers.GetAllDIDsHandler) // 获取所有DID列表

		// 钱包相关的DID操作
		did.POST("/wallet/:walletAddress", router.DIDHandlers.CreateDIDFromWalletHandler) // 通过钱包地址创建DID
		did.GET("/wallet/:walletAddress", router.DIDHandlers.GetDIDByWalletHandler)       // 获取钱包的DID信息
		did.GET("/list/:walletAddress", router.DIDHandlers.ListDIDsByWalletHandler)       // 列出钱包的所有DID

		// 医生DID相关操作
		did.POST("/doctor/create", router.VCHandlers.CreateDoctorDIDHandler) // 创建医生DID

		// 保留DID解析功能
		did.POST("/resolve", router.DIDHandlers.ResolveDIDHandler) // 解析DID文档

		// 已弃用的方法（返回错误提示）
		did.POST("/create", router.DIDHandlers.CreateDIDHandler) // 已弃用
		did.POST("/update", router.DIDHandlers.UpdateDIDHandler) // 已弃用
		did.POST("/revoke", router.DIDHandlers.RevokeDIDHandler) // 已弃用
	}

	// VC路由
	vc := api.Group("/vc")
	{
		// 通用VC操作
		vc.POST("/issue", router.VCHandlers.IssueCredentialHandler)
		vc.POST("/verify", router.VCHandlers.VerifyCredentialHandler)
		vc.POST("/revoke", router.VCHandlers.RevokeCredentialHandler)
		vc.POST("/presentation/create", router.VCHandlers.CreatePresentationHandler)
		vc.POST("/presentation/verify", router.VCHandlers.VerifyPresentationHandler)

		// 医生VC相关操作
		vc.POST("/doctor/issue", router.VCHandlers.IssueDoctorVCHandler)    // 颁发医生凭证
		vc.POST("/doctor/verify", router.VCHandlers.VerifyDoctorVCHandler)  // 验证医生凭证
		vc.GET("/doctor/:doctorDID", router.VCHandlers.GetDoctorVCsHandler) // 获取医生凭证列表
	}

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

		// 集成NFT+ABE相关
		// secured.POST("/nft/mint-encrypted", router.NFTHandlers.MintEncryptedNFTHandler)
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
