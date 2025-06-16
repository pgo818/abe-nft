package main

import (
	"log"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"nft-go-backend/internal/api"
	"nft-go-backend/internal/blockchain"
	"nft-go-backend/internal/config"
	"nft-go-backend/internal/models"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库
	err = models.InitDB(cfg.GetDSN())
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 初始化区块链客户端
	client, err := blockchain.NewEthClient(cfg)
	if err != nil {
		log.Fatalf("初始化以太坊客户端失败: %v", err)
	}

	// 创建Gin路由引擎
	r := gin.Default()

	// 添加中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 添加CORS中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 获取当前文件的目录
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)

	// 使用绝对路径
	templatesPath := filepath.Join(currentDir, "../../templates/*")
	staticPath := filepath.Join(currentDir, "../../static")

	// 添加静态文件服务
	r.Static("/static", staticPath)

	// 设置HTML模板目录
	r.LoadHTMLGlob(templatesPath)

	// 主页路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "NFT平台",
		})
	})

	// 初始化API处理程序
	handlers := api.NewHandlers(client)

	// 设置路由
	handlers.SetupRoutes(r)

	// 启动事件监听（非阻塞）
	go client.ListenToEvents()

	// 启动服务器
	port := ":" + cfg.Port
	log.Printf("API服务器启动在 %s 端口", port)
	log.Fatal(r.Run(port))
}
