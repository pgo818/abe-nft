# NFT项目 - Hardhat框架

这是一个基于Hardhat框架的完整NFT项目，包含ERC721合约、测试和部署脚本。

## 功能特点

- 🎨 完整的ERC721 NFT合约
- 💰 付费铸造功能
- 🔢 批量铸造支持
- 👑 管理员免费铸造
- 💸 资金提取功能
- 🔧 可配置的参数（价格、供应量、BaseURI）
- ✅ 完整的测试覆盖
- 🚀 自动化部署脚本

## 快速开始

### 安装依赖
```bash
npm install
```

### 编译合约
```bash
npm run compile
```

### 运行测试
```bash
npm run test
```

### 启动本地节点
```bash
npm run node
```

### 部署到本地网络（新终端）
```bash
npm run deploy:local
```

## 合约功能

### 主要功能
- `safeMint(address to)` - 付费铸造单个NFT
- `safeMintBatch(address to, uint256 quantity)` - 付费批量铸造
- `ownerMint(address to, uint256 quantity)` - 管理员免费铸造
- `withdraw()` - 提取合约资金

### 配置功能（仅管理员）
- `setMintPrice(uint256 _mintPrice)` - 设置铸造价格
- `setMaxSupply(uint256 _maxSupply)` - 设置最大供应量
- `setBaseURI(string memory baseURI)` - 设置元数据URI

### 查询功能
- `totalSupply()` - 当前总供应量
- `maxSupply()` - 最大供应量
- `mintPrice()` - 当前铸造价格

## 部署说明

1. 配置网络（在hardhat.config.js中）
2. 设置环境变量（如需要）
3. 运行部署命令

### 本地部署
```bash
npm run node              # 启动本地节点
npm run deploy:local      # 部署到本地
```

### 测试网部署
```bash
npm run deploy:sepolia    # 部署到Sepolia测试网
```

## 测试

项目包含全面的测试用例：
- 部署测试
- 铸造功能测试
- 权限控制测试
- 资金管理测试

运行测试：
```bash
npm run test
```

## 项目结构

```
├── contracts/          # 智能合约
│   ├── MyToken.sol     # 主NFT合约
│   └── Lock.sol        # 示例合约
├── test/               # 测试文件
│   └── MyToken.test.js # NFT合约测试
├── ignition/           # 部署脚本
│   └── modules/
│       └── MyToken.js  # NFT部署模块
├── hardhat.config.js   # Hardhat配置
└── package.json        # 项目配置
```

## 技术栈

- **Solidity ^0.8.20** - 智能合约语言
- **Hardhat** - 开发框架
- **OpenZeppelin** - 智能合约库
- **Ethers.js** - 以太坊库
- **Chai** - 测试框架

## 📚 项目文档

为了帮助用户快速上手，项目提供了以下详细文档：

- **[QUICK_START_README.md](QUICK_START_README.md)** - 快速启动指南
- **[PROJECT_SETUP_GUIDE.md](PROJECT_SETUP_GUIDE.md)** - 完整部署指南  
- **[TROUBLESHOOTING.md](TROUBLESHOOTING.md)** - 故障排除指南
- **[PORT_CONFIGURATION.md](PORT_CONFIGURATION.md)** - 端口配置说明

## 🚀 快速启动

### Windows用户
```bash
quick-start.bat
```

### Linux/Mac用户
```bash
chmod +x quick-start.sh
./quick-start.sh
```

## 📋 项目端口配置

- **Go后端服务**: 8080端口
- **Vue前端开发**: 8081端口  
- **Ganache区块链**: 7545端口
- **MySQL数据库**: 3306端口

详细配置说明请查看 [PORT_CONFIGURATION.md](PORT_CONFIGURATION.md)

## 许可证

MIT 


## 更新NFT

# 重新编译合约
npx hardhat compile

# 重新生成Go绑定文件
abigen --abi=./artifacts/contracts/MainNFT.sol/MainNFT.abi --pkg=mainnft --out=./nft-go-backend/pkg/mainnft/MainNFT.go
abigen --abi=./artifacts/contracts/ChildNFT.sol/ChildNFT.abi --pkg=childnft --out=./nft-go-backend/pkg/childnft/ChildNFT.go

# 重新部署合约
npx hardhat run scripts/deploy.js --network localhost