# 端口配置说明

## 📋 正确的端口配置

经过确认，项目的正确端口配置如下：

### 🔧 服务端口分配

| 服务名称 | 端口 | 协议 | 说明 |
|---------|------|------|------|
| **Go后端服务** | `8080` | HTTP | 主要API服务端口 |
| **Vue前端开发服务器** | `8081` | HTTP | Vue开发服务器端口 |
| **Ganache区块链节点** | `7545` | HTTP/RPC | 本地以太坊网络 |
| **MySQL数据库** | `3306` | TCP | 数据库服务端口 |

### 🔄 服务间通信

```
用户浏览器 → Vue前端(8081) → 代理转发 → Go后端(8080)
                                           ↓
Go后端(8080) → Ganache(7545) + MySQL(3306)
```

### 📝 配置文件更新

所有相关文档和脚本已更新为正确的端口配置：

#### 后端配置 (`nft-go-backend/.env`)
```env
PORT=8080  # Go后端端口
ETHEREUM_RPC=http://localhost:7545  # Ganache端口
DB_PORT=3306  # MySQL端口
```

#### Vue前端配置 (`vue.config.js`)
```javascript
devServer: {
    port: 8081,  // Vue开发服务器端口
    proxy: {
        '/api': {
            target: 'http://localhost:8080',  // 代理到Go后端
            changeOrigin: true
        }
    }
}
```

### 🌐 访问地址

- **主应用访问**: http://localhost:8080
- **Vue开发模式**: http://localhost:8081 (仅开发时)
- **API接口**: http://localhost:8080/api/*

### ⚠️ 重要说明

1. **生产环境**: 只需要访问 `http://localhost:8080`，Vue前端会作为静态文件被Go后端服务
2. **开发环境**: 可以同时运行Vue开发服务器 (`8081`) 和Go后端 (`8080`)
3. **端口冲突**: 如果8080端口被占用，可以修改 `.env` 文件中的 `PORT` 配置

### 🚀 启动顺序

1. **Ganache** (7545端口) - 区块链网络
2. **MySQL** (3306端口) - 数据库服务  
3. **Go后端** (8080端口) - API服务
4. **Vue前端** (8081端口) - 可选，仅开发时

### 🔧 故障排除

#### 端口占用检查
```bash
# Windows
netstat -ano | findstr :8080
netstat -ano | findstr :8081

# Linux/Mac  
lsof -i :8080
lsof -i :8081
```

#### 修改端口配置
如果需要修改端口，请更新以下文件：
- `nft-go-backend/.env` - Go后端端口
- `nft-go-backend/static/vue-frontend/vue.config.js` - Vue前端端口
- `nft-go-backend/start.bat` / `start.sh` - 启动脚本中的环境变量

---

**注意**: 文档中之前错误标注的9090端口已全部更正为8080端口。 