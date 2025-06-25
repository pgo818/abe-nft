# 故障排除指南

## 🚨 常见问题及解决方案

### 1. 环境相关问题

#### Node.js版本问题
```bash
# 错误：Node.js版本过低
Error: Unsupported Node.js version

# 解决方案：
# 1. 卸载旧版本Node.js
# 2. 从 https://nodejs.org/ 下载最新LTS版本
# 3. 验证安装
node --version  # 应该显示 v18.0.0 或更高版本
npm --version
```

#### Go语言环境问题
```bash
# 错误：Go版本不兼容
go: cannot find module

# 解决方案：
# 1. 更新Go到1.21+版本
go version  # 检查版本
# 2. 设置GOPATH和GOROOT环境变量
# 3. 清理模块缓存
go clean -modcache
go mod tidy
```

#### MySQL连接问题
```bash
# 错误：数据库连接失败
Error: dial tcp [::1]:3306: connect: connection refused

# 解决方案：
# Windows:
net start mysql
# macOS:
brew services start mysql
# Linux:
sudo service mysql start

# 检查MySQL状态
mysql -u root -p -e "SELECT 1"
```

### 2. 区块链相关问题

#### Ganache启动问题
```bash
# 错误：无法连接到区块链网络
Error: connect ECONNREFUSED 127.0.0.1:7545

# 解决方案：
# 1. 确保Ganache在正确端口运行
# GUI版本：检查设置中的端口配置
# CLI版本：
ganache-cli --port 7545 --chainId 1337 --deterministic

# 2. 检查防火墙设置
# 3. 验证网络连接
curl http://localhost:7545
```

#### 智能合约部署失败
```bash
# 错误：合约部署失败
Error: Transaction reverted

# 解决方案：
# 1. 检查账户余额
# 2. 清理之前的部署
npm run clean
# 3. 重新编译
npm run compile
# 4. 重新部署
npm run deploy:local

# 如果仍然失败，尝试重启Ganache
```

#### MetaMask连接问题
```bash
# 错误：MetaMask无法连接

# 解决方案：
# 1. 检查网络配置
网络名称：Local Ganache
RPC URL：http://localhost:7545
链ID：1337
货币符号：ETH

# 2. 重置MetaMask连接
# 3. 清除浏览器缓存
# 4. 重新导入账户私钥
```

### 3. 后端服务问题

#### Go模块依赖问题
```bash
# 错误：模块下载失败
go: module download failed

# 解决方案：
cd nft-go-backend
# 1. 清理模块缓存
go clean -modcache
# 2. 设置代理（中国用户）
go env -w GOPROXY=https://goproxy.cn,direct
# 3. 重新下载依赖
go mod download
go mod tidy
```

#### 端口占用问题
```bash
# 错误：端口8080已被占用
Error: listen tcp :8080: bind: address already in use

# 解决方案：
# 1. 查找占用进程
# Windows:
netstat -ano | findstr :8080
taskkill /PID <PID> /F
# Linux/Mac:
lsof -i :8080
kill -9 <PID>

# 2. 或修改端口配置
# 编辑 nft-go-backend/.env
PORT=8081
```

#### 数据库表不存在
```bash
# 错误：Table doesn't exist
Error: Table 'nft_db.nfts' doesn't exist

# 解决方案：
# 1. 检查数据库是否创建
mysql -u root -p -e "SHOW DATABASES;"
# 2. 手动创建数据库
mysql -u root -p -e "CREATE DATABASE nft_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
# 3. 运行初始化脚本
mysql -u root -p nft_db < nft-go-backend/scripts/init_db.sql
# 4. 重启后端服务（会自动创建表）
```

### 4. 前端相关问题

#### npm安装失败
```bash
# 错误：npm依赖安装失败
npm ERR! peer dep missing

# 解决方案：
# 1. 清除npm缓存
npm cache clean --force
# 2. 删除node_modules
rm -rf node_modules package-lock.json
# 3. 重新安装
npm install
# 4. 如果仍失败，尝试使用yarn
npm install -g yarn
yarn install
```

#### Vue编译错误
```bash
# 错误：Vue编译失败
Module not found: Error: Can't resolve

# 解决方案：
cd nft-go-backend/static/vue-frontend
# 1. 检查依赖
npm list
# 2. 重新安装依赖
npm install
# 3. 清除缓存
npm run clean  # 如果有这个命令
# 4. 重新构建
npm run build
```

### 5. 特定功能问题

#### ABE加密失败
```bash
# 错误：ABE加密操作失败

# 解决方案：
# 1. 检查属性格式
# 属性应该用英文，避免特殊字符
# 2. 检查策略格式
# 正确格式："(医生 AND 主治医师) OR 专家"
# 3. 检查系统密钥是否正确初始化
```

#### IPFS上传失败
```bash
# 错误：IPFS上传失败

# 解决方案：
# 1. 检查IPFS配置
# 2. 验证访问密钥
# 3. 检查网络连接
# 4. 尝试使用公共IPFS网关
```

#### DID创建失败
```bash
# 错误：DID创建失败

# 解决方案：
# 1. 检查钱包连接
# 2. 确保账户有足够ETH（gas费）
# 3. 检查网络连接
# 4. 验证合约地址配置
```

### 6. 性能优化

#### 启动速度慢
```bash
# 问题：项目启动缓慢

# 优化方案：
# 1. 使用SSD硬盘
# 2. 增加内存
# 3. 关闭不必要的后台程序
# 4. 使用国内镜像源
npm config set registry https://registry.npm.taobao.org/
```

#### 交易确认慢
```bash
# 问题：区块链交易确认慢

# 解决方案：
# 1. 增加gas价格
# 2. 检查Ganache配置
# 3. 重启Ganache
# 4. 清除MetaMask历史记录
```

### 7. 调试技巧

#### 查看日志
```bash
# 后端日志
cd nft-go-backend
# 启动服务并查看详细日志

# 前端日志
# 打开浏览器开发者工具查看控制台

# 区块链日志
# 查看Ganache窗口或控制台输出
```

#### API测试
```bash
# 测试后端API
curl http://localhost:8080/api/health
curl http://localhost:8080/api/nfts

# 使用Postman或其他API测试工具
```

#### 数据库检查
```sql
-- 检查数据库连接
mysql -u root -p -e "SELECT 1"

-- 查看表结构
USE nft_db;
SHOW TABLES;
DESCRIBE nfts;

-- 查看数据
SELECT * FROM nfts LIMIT 5;
```

### 8. 重置项目

#### 完全重置
```bash
# 如果遇到无法解决的问题，可以完全重置项目

# 1. 停止所有服务
# 2. 清理数据库
mysql -u root -p -e "DROP DATABASE IF EXISTS nft_db;"
# 3. 清理缓存
npm cache clean --force
go clean -cache
# 4. 删除依赖
rm -rf node_modules
rm -rf nft-go-backend/static/vue-frontend/node_modules
# 5. 重新开始
npm install
cd nft-go-backend && go mod download
cd static/vue-frontend && npm install
```

### 9. 获取帮助

如果以上方案都无法解决问题：

1. **检查日志**：仔细查看错误信息
2. **搜索错误**：在GitHub Issues或Stack Overflow搜索相关错误
3. **查看文档**：阅读 PROJECT_SETUP_GUIDE.md 了解详细配置
4. **重新部署**：按照快速启动指南重新部署项目

### 10. 系统要求检查

最低系统要求：
- **操作系统**：Windows 10/macOS 10.15/Ubuntu 18.04+
- **内存**：8GB RAM
- **硬盘**：10GB 可用空间
- **网络**：稳定的互联网连接

推荐配置：
- **内存**：16GB RAM
- **硬盘**：SSD
- **CPU**：4核以上 