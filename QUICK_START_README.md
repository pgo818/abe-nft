# NFT+ABE+DID/VC 集成管理平台 - 快速启动

## 🚀 一键启动（推荐）

### Windows用户
```bash
# 在项目根目录运行
quick-start.bat
```

### Linux/Mac用户  
```bash
# 在项目根目录运行
chmod +x quick-start.sh
./quick-start.sh
```

## 📋 启动前准备

### 1. 环境要求
- **Node.js** 18.0+
- **Go** 1.21+  
- **MySQL** 5.7+
- **Ganache** (区块链本地节点)
- **MetaMask** (浏览器钱包)

### 2. 下载安装链接
- Node.js: https://nodejs.org/
- Go语言: https://golang.org/dl/
- MySQL: https://dev.mysql.com/downloads/installer/
- Ganache: https://trufflesuite.com/ganache/
- MetaMask: https://metamask.io/

## ⚡ 手动启动步骤

如果一键启动脚本失败，可以按以下步骤手动启动：

### 1. 安装依赖
```bash
# 安装Hardhat依赖
npm install

# 安装Go后端依赖
cd nft-go-backend
go mod download

# 安装Vue前端依赖
cd static/vue-frontend
npm install
cd ../../..
```

### 2. 配置数据库
```sql
# 连接MySQL并创建数据库
mysql -u root -p
CREATE DATABASE nft_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 3. 启动Ganache
- 启动Ganache GUI或命令行版本
- 配置端口：7545，链ID：1337

### 4. 配置MetaMask
- 添加本地网络：http://localhost:7545
- 链ID：1337
- 导入Ganache账户私钥

### 5. 部署智能合约
```bash
# 编译合约
npm run compile

# 部署到本地网络
npm run deploy:local
```

### 6. 配置后端环境
编辑 `nft-go-backend/.env` 文件：
```env
PORT=8080
ETHEREUM_RPC=http://localhost:7545
MAIN_NFT_ADDRESS=智能合约地址
CHILD_NFT_ADDRESS=子NFT合约地址
PRIVATE_KEY=Ganache账户私钥
CHAIN_ID=1337
DB_USER=root
DB_PASSWORD=你的MySQL密码
DB_HOST=localhost
DB_PORT=3306
DB_NAME=nft_db
```

### 7. 启动后端服务
```bash
cd nft-go-backend
# Windows
start.bat
# Linux/Mac
./start.sh
```

## 🌐 访问项目

启动成功后访问：http://localhost:8080

## 💼 业务流程

### 数据拥有者
1. 访问 http://localhost:8080 → 连接钱包 → 注册DID身份 → 创建VC凭证
2. 上传文件到IPFS → 获取hash地址  
3. 用ABE加密hash → 设置访问策略
4. 创建NFT元数据 → 添加密文和策略
5. 铸造NFT → 完成数据资产化

### 数据使用者
1. 访问 http://localhost:8080 → 连接钱包 → 注册DID身份 → 创建VC凭证
2. 申请子NFT → 提交凭证信息
3. 等待审批 → 获得子NFT
4. 申请密钥 → 基于子NFT属性生成
5. 解密数据 → 获取原始文件hash
6. 下载文件 → 从IPFS获取原始数据

## 📚 详细文档

如需完整的部署和配置指南，请参考：
- `PROJECT_SETUP_GUIDE.md` - 完整部署指南
- `nft-go-backend/README.md` - 后端详细文档

## ❓ 常见问题

**Q: 合约部署失败？**
A: 检查Ganache是否在7545端口运行

**Q: 后端启动失败？**  
A: 检查MySQL服务和数据库配置

**Q: 前端无法连接？**
A: 确认MetaMask网络配置正确

**Q: 数据库连接失败？**
A: 检查MySQL服务状态和密码配置

## 🛠️ 技术栈

- **智能合约**: Solidity + Hardhat
- **后端**: Go + Gin + GORM  
- **前端**: Vue.js + Bootstrap
- **区块链**: Ethereum + MetaMask
- **加密**: ABE (Attribute-Based Encryption)
- **身份**: DID/VC (去中心化身份/可验证凭证)
- **存储**: IPFS + MySQL 