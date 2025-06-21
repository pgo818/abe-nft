# DID功能修复指南

## 问题总结

1. **使用MetaMask钱包创建DID失败**：报错 `Error 1364 (HY000)`
2. **需要限制DID创建方式**：只允许通过钱包地址创建DID
3. **DID列表功能**：只显示自己钱包创建的DID信息

## 解决方案

### 1. 数据库修复

已更新 `scripts/init_db.sql` 文件，添加了完整的DID和VC相关表结构：

- `dids` 表：存储DID身份信息
- `verifiable_credentials` 表：存储可验证凭证
- `verifiable_presentations` 表：存储可验证表示
- `credential_schemas` 表：存储凭证模式
- `credential_definitions` 表：存储凭证定义

**重新初始化数据库**：
```bash
mysql -u root -p < scripts/init_db.sql
```

### 2. 服务层修复

已修复 `internal/service/did_service.go`：

- **修复了CreateDIDFromWallet方法**：
  - 添加了事务支持
  - 改进了错误处理
  - 确保所有必需字段都有值
  
- **添加了ListDIDsByWallet方法**：
  - 获取指定钱包地址的所有DID
  - 只返回活跃状态的DID

- **弃用了其他创建方法**：
  - CreateDID, UpdateDID, RevokeDID 都已标记为弃用

### 3. API处理器修复

已修复 `internal/api/did_handlers.go`：

- **改进了CreateDIDFromWalletHandler**：
  - 更好的响应格式
  - 区分创建和已存在的情况
  
- **添加了ListDIDsByWalletHandler**：
  - 列出指定钱包的所有DID
  - 返回详细的DID信息

- **弃用了其他处理器**：
  - 返回适当的错误消息，引导使用正确的API

### 4. 路由配置修复

已修复 `internal/api/router.go`：

```go
// DID路由 - 仅支持钱包地址创建和管理
did := api.Group("/did")
{
    // 钱包相关的DID操作
    did.POST("/wallet/:walletAddress", router.DIDHandlers.CreateDIDFromWalletHandler)  // 通过钱包地址创建DID
    did.GET("/wallet/:walletAddress", router.DIDHandlers.GetDIDByWalletHandler)        // 获取钱包的DID信息
    did.GET("/list/:walletAddress", router.DIDHandlers.ListDIDsByWalletHandler)        // 列出钱包的所有DID
    
    // 保留DID解析功能
    did.POST("/resolve", router.DIDHandlers.ResolveDIDHandler)                         // 解析DID文档
    
    // 已弃用的方法（返回错误提示）
    did.POST("/create", router.DIDHandlers.CreateDIDHandler)                           // 已弃用
    did.POST("/update", router.DIDHandlers.UpdateDIDHandler)                           // 已弃用
    did.POST("/revoke", router.DIDHandlers.RevokeDIDHandler)                           // 已弃用
}
```

### 5. 模型修复

已修复 `internal/models/did.go` 和 `internal/models/db.go`：

- 修复了VerifiableCredentialResponse结构体的字段类型
- 添加了自动迁移支持
- 修复了VC handlers中的类型兼容性问题

## API使用说明

### 创建DID（通过钱包地址）
```http
POST /api/did/wallet/{walletAddress}
```

响应示例：
```json
{
  "message": "DID创建成功",
  "data": {
    "did": "did:ethr:0x1234567890abcdef1234567890abcdef12345678",
    "walletAddress": "0x1234567890abcdef1234567890abcdef12345678",
    "exists": false
  }
}
```

### 获取钱包的DID信息
```http
GET /api/did/wallet/{walletAddress}
```

### 列出钱包的所有DID
```http
GET /api/did/list/{walletAddress}
```

响应示例：
```json
{
  "walletAddress": "0x1234567890abcdef1234567890abcdef12345678",
  "dids": [
    {
      "id": 1,
      "didString": "did:ethr:0x1234567890abcdef1234567890abcdef12345678",
      "walletAddress": "0x1234567890abcdef1234567890abcdef12345678",
      "status": "active",
      "createdAt": "2024-01-01T00:00:00Z",
      "updatedAt": "2024-01-01T00:00:00Z"
    }
  ],
  "total": 1,
  "message": "获取DID列表成功"
}
```

## 测试步骤

1. **重新初始化数据库**：
   ```bash
   mysql -u root -p < scripts/init_db.sql
   ```

2. **启动服务**：
   ```bash
   go run cmd/server/main.go
   ```

3. **测试创建DID**：
   ```bash
   curl -X POST http://localhost:8080/api/did/wallet/0x1234567890abcdef1234567890abcdef12345678
   ```

4. **测试获取DID列表**：
   ```bash
   curl -X GET http://localhost:8080/api/did/list/0x1234567890abcdef1234567890abcdef12345678
   ```

## 注意事项

1. **数据库兼容性**：确保使用MySQL 5.7+，支持JSON字段类型
2. **钱包地址格式**：使用标准的以太坊地址格式（0x开头的42位十六进制字符串）
3. **并发安全**：使用了数据库事务确保创建DID的原子性
4. **唯一性约束**：一个钱包地址只能对应一个活跃的DID

## 故障排除

如果仍然遇到问题：

1. 检查数据库连接配置
2. 确认MySQL版本支持JSON字段
3. 查看应用日志获取详细错误信息
4. 确认钱包地址格式正确 