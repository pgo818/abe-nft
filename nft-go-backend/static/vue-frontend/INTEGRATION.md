# 后端集成指南

本文档介绍如何将Vue前端与Go后端集成。

## 开发环境

在开发环境中，Vue前端和Go后端分别运行在不同的端口上：

- Vue前端: http://localhost:8080
- Go后端: http://localhost:3000

Vue开发服务器配置了代理，将所有 `/api` 请求转发到Go后端。这在 `vue.config.js` 中配置：

```js
devServer: {
  proxy: {
    '/api': {
      target: 'http://localhost:3000',
      changeOrigin: true
    }
  }
}
```

## 生产环境

在生产环境中，Vue前端会被构建成静态文件，并由Go后端提供服务。

### 构建前端

运行以下命令构建Vue前端：

```bash
cd nft-go-backend/static/vue-frontend
npm run build
```

构建后的文件将位于 `nft-go-backend/static/dist` 目录中。

### 配置Go后端

Go后端需要配置为提供静态文件服务，并处理前端路由。在 `main.go` 中添加以下代码：

```go
// 配置静态文件服务
r.Static("/assets", "./static/dist/assets")
r.StaticFile("/favicon.ico", "./static/dist/favicon.ico")

// 处理前端路由
r.NoRoute(func(c *gin.Context) {
    // 检查是否是API请求
    if strings.HasPrefix(c.Request.URL.Path, "/api/") {
        c.JSON(http.StatusNotFound, gin.H{"error": "API not found"})
        return
    }
    
    // 返回index.html以支持前端路由
    c.File("./static/dist/index.html")
})
```

## API接口

前端通过以下方式调用后端API：

1. 直接调用：
```js
const response = await fetch('/api/nfts')
const data = await response.json()
```

2. 带签名的请求（需要身份验证的API）：
```js
// 创建签名消息
const message = JSON.stringify({
  action: 'get_my_nfts',
  timestamp: Date.now()
})

// 获取签名
const signature = await window.ethereum.request({
  method: 'personal_sign',
  params: [message, account]
})

// 发送请求
const response = await fetch('/api/nft/my-nfts', {
  method: 'GET',
  headers: {
    'Content-Type': 'application/json',
    'X-Ethereum-Address': account,
    'X-Ethereum-Signature': signature,
    'X-Ethereum-Message': message
  }
})
```

## 注意事项

1. 确保Go后端允许跨域请求（在开发环境中）
2. 所有API路由应以 `/api` 开头
3. 前端路由由Vue Router处理，后端应将所有非API请求路由到index.html 