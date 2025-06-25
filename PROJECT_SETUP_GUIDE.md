# NFT+ABE+DID/VC 集成管理平台 - 完整部署指南

## 项目概述

这是一个集成了NFT（非同质化代币）管理、ABE（基于属性的加密）功能和DID/VC（去中心化身份/可验证凭证）的完整平台。项目包含：

- **智能合约部分**：基于Hardhat框架的ERC721 NFT合约
- **后端服务**：Go语言开发的API服务，提供ABE加密、DID/VC管理功能
- **前端界面**：Vue.js开发的现代化Web界面

## 系统架构

```
项目结构:
├── contracts/              # 智能合约（Solidity）
├── nft-go-backend/         # 后端服务（Go）
│   ├── cmd/server/         # 主服务入口
│   ├── internal/           # 内部业务逻辑
│   ├── pkg/               # 智能合约Go绑定
│   └── static/vue-frontend/ # Vue前端项目
├── scripts/               # 部署脚本
└── ignition/              # 合约部署配置
```

## 一、环境要求

### 基础软件环境
- **Node.js**: 18.0+ 
- **Go**: 1.21+
- **MySQL**: 5.7+ 或 8.0+
- **Git**: 最新版本

### 区块链开发环境
- **Ganache**: 本地区块链节点
- **MetaMask**: 浏览器钱包插件

### 可选工具
- **IPFS Desktop**: 用于文件存储（如需IPFS功能）
- **Postman**: 用于API测试

## 二、环境安装

### 1. 安装Node.js和npm
```bash
# 访问 https://nodejs.org/ 下载LTS版本
# 验证安装
node --version
npm --version
```

### 2. 安装Go语言
```bash
# 访问 https://golang.org/dl/ 下载对应系统版本
# 验证安装
go version
```

### 3. 安装MySQL
```bash
# Windows: 下载MySQL Installer
# macOS: brew install mysql
# Ubuntu: sudo apt-get install mysql-server

# 启动MySQL服务并设置root密码
```

### 4. 安装Ganache
```bash
# 下载Ganache GUI版本
# 访问 https://trufflesuite.com/ganache/
# 或使用命令行版本
npm install -g ganache-cli
```

## 三、项目下载与配置

### 1. 克隆项目
```bash
git clone <项目仓库地址>
cd nft
```

### 2. 安装项目依赖

#### 安装Hardhat依赖（智能合约部分）
```bash
# 在项目根目录
npm install
```

#### 安装Go依赖（后端部分）
```bash
cd nft-go-backend
go mod download
go mod tidy
```

#### 安装Vue依赖（前端部分）
```bash
cd nft-go-backend/static/vue-frontend
npm install
```

### 3. 数据库配置

#### 创建数据库
```sql
# 连接到MySQL
mysql -u root -p

# 创建数据库
CREATE DATABASE nft_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

# 创建用户（可选）
CREATE USER 'nft_user'@'localhost' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON nft_db.* TO 'nft_user'@'localhost';
FLUSH PRIVILEGES;
```

#### 运行初始化脚本
```bash
cd nft-go-backend
mysql -u root -p nft_db < scripts/init_db.sql
```

## 四、区块链环境配置

### 1. 启动Ganache
```bash
# 方法1: 使用Ganache GUI
# 启动Ganache桌面版，配置：
# - PORT: 7545
# - CHAIN ID: 1337
# - ACCOUNTS: 10个账户，每个账户100 ETH

# 方法2: 使用命令行
ganache-cli --port 7545 --chainId 1337 --accounts 10 --deterministic
```

### 2. 配置MetaMask
1. 安装MetaMask浏览器插件
2. 添加本地网络：
   - 网络名称：Local Ganache
   - RPC URL：http://localhost:7545
   - 链ID：1337
   - 货币符号：ETH
3. 导入Ganache账户私钥到MetaMask

### 3. 编译和部署智能合约
```bash
# 在项目根目录

# 编译合约
npm run compile

# 启动本地节点（如果使用Hardhat节点）
npm run node

# 部署合约到本地网络
npm run deploy:local
```

记录部署后的合约地址，用于后端配置。

## 五、后端服务配置

### 1. 创建环境配置文件
在 `nft-go-backend` 目录下创建 `.env` 文件：

```bash
cd nft-go-backend
```

创建 `.env` 文件内容：
```env
# 服务配置
PORT=8080

# 区块链配置
ETHEREUM_RPC=http://localhost:7545
MAIN_NFT_ADDRESS=your_main_nft_contract_address
CHILD_NFT_ADDRESS=your_child_nft_contract_address
PRIVATE_KEY=your_server_account_private_key
CHAIN_ID=1337

# 数据库配置
DB_USER=root
DB_PASSWORD=your_mysql_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=nft_db

# IPFS配置（可选）
IPFS_ACCESS_KEY=your_ipfs_access_key
```

### 2. 更新start.bat文件中的配置
编辑 `nft-go-backend/start.bat`，更新以下变量：
- `MAIN_NFT_ADDRESS`：主NFT合约地址
- `CHILD_NFT_ADDRESS`：子NFT合约地址
- `PRIVATE_KEY`：服务器账户私钥
- `DB_PASSWORD`：MySQL密码

## 六、启动项目

### 1. 启动后端服务

#### Windows系统
```bash
cd nft-go-backend
start.bat
```

#### Linux/macOS系统
```bash
cd nft-go-backend
chmod +x start.sh
./start.sh
```

#### 手动启动
```bash
cd nft-go-backend
go build -o server cmd/server/main.go
./server
```

### 2. 构建前端项目（可选）
```bash
cd nft-go-backend/static/vue-frontend
npm run build
```

### 3. 访问项目
- 主界面：http://localhost:8080
- Vue前端（开发模式）：http://localhost:8081

## 七、功能使用流程

根据你提供的业务流程，使用步骤如下：

### 数据拥有者流程

#### 1. 钱包登录和身份注册
1. 访问 http://localhost:8080
2. 连接MetaMask钱包
3. 访问"DID和VC管理"页面
4. 创建医生DID信息
5. 在"颁发凭证"页面注册VC凭证信息

#### 2. 文件上传到IPFS
1. 访问"ABE加密管理" -> "IPFS上传"页面
2. 上传需要加密的文件
3. 获取IPFS存储地址hash

#### 3. 文件加密
1. 在"ABE加密管理" -> "加密"页面
2. 输入要加密的IPFS hash地址
3. 设置访问策略（如："医生 AND 主治医师"）
4. 获取加密密文

#### 4. 创建NFT元数据
1. 访问"NFT管理平台" -> "元数据管理"
2. 创建NFT元数据
3. 将密文和访问策略添加到元数据中
4. 获取元数据URI信息

#### 5. 铸造NFT
1. 在"NFT管理平台" -> "铸造NFT"页面
2. 输入步骤4获取的URI信息
3. 完成NFT铸造

### 数据使用者流程

#### 1. 钱包登录和身份注册
1. 使用不同的MetaMask账户登录
2. 重复数据拥有者的步骤1

#### 2. 申请子NFT
1. 访问"NFT管理平台" -> "所有NFT"页面
2. 找到需要申请的NFT
3. 点击"申请子NFT"
4. 输入自己的凭证信息
5. 提交申请

#### 3. 获取子NFT（需要数据拥有者审批）
1. 数据拥有者在"申请管理"页面审批申请
2. 若凭证满足策略，申请者获得子NFT

#### 4. 申请密钥
1. 访问"ABE加密管理" -> "密钥生成"页面
2. 选择自己拥有的子NFT属性
3. 生成相关解密密钥

#### 5. 解密文件
1. 在"ABE加密管理" -> "解密"页面
2. 使用密钥和子NFT中的密文进行解密
3. 获得原始IPFS hash地址

#### 6. 下载文件
1. 访问IPFS管理界面
2. 输入解密得到的hash地址
3. 下载原始文件

## 八、常见问题解决

### 1. 智能合约部署失败
```bash
# 检查Ganache是否正在运行
# 检查网络配置是否正确
# 清理并重新编译
npm run clean
npm run compile
npm run deploy:local
```

### 2. 后端服务启动失败
```bash
# 检查MySQL服务是否启动
# 检查数据库连接配置
# 检查Go版本是否符合要求
go version
```

### 3. 前端连接问题
- 检查MetaMask网络配置
- 确认合约地址配置正确
- 检查CORS配置

### 4. 数据库连接问题
```sql
# 检查MySQL服务状态
# Windows
net start mysql

# Linux/macOS
sudo service mysql start
```

## 九、开发和测试

### 运行测试
```bash
# 智能合约测试
npm run test

# 特定测试
npm run test:nft
```

### API测试
使用Postman或curl测试后端API：
```bash
# 测试健康检查
curl http://localhost:8080/api/health

# 测试NFT列表
curl http://localhost:8080/api/nfts
```

## 十、生产部署注意事项

1. **安全配置**
   - 使用HTTPS
   - 保护私钥和数据库密码
   - 配置防火墙

2. **性能优化**
   - 配置数据库连接池
   - 启用缓存
   - 配置负载均衡

3. **监控和日志**
   - 配置日志记录
   - 设置监控告警
   - 定期备份数据库

## 十一、技术支持

如果在部署过程中遇到问题，请检查：

1. 所有服务是否正常运行（MySQL、Ganache）
2. 网络配置是否正确
3. 合约地址是否正确配置
4. 账户余额是否充足（用于交易gas费）

---

**注意**：首次启动时，系统会自动创建数据库表结构。如果遇到权限问题，请确保数据库用户有足够的权限。 