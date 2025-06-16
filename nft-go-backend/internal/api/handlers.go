package api

import (
	"nft-go-backend/internal/blockchain"

	"github.com/gin-gonic/gin"
)

// Handlers API处理程序结构体（保留向后兼容性）
type Handlers struct {
	*Router
}

// NewHandlers 创建新的API处理程序（保留向后兼容性）
func NewHandlers(client *blockchain.EthClient) *Handlers {
	return &Handlers{
		Router: NewRouter(client),
	}
}

// SetupRoutes 设置路由（保留向后兼容性）
func (h *Handlers) SetupRoutes(r *gin.Engine) {
	h.Router.SetupRoutes(r)
}
