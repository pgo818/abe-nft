#!/bin/bash

echo "===== NFT+ABE+DID/VC 项目快速启动脚本 ====="
echo

# 检查是否在项目根目录
if [ ! -f "package.json" ]; then
    echo "❌ 错误：请在项目根目录运行此脚本！"
    echo "当前目录：$(pwd)"
    exit 1
fi

echo "⏳ 正在检查环境依赖..."

# 检查Node.js
if ! command -v node &> /dev/null; then
    echo "❌ 未检测到Node.js，请先安装Node.js 18.0+"
    echo "下载地址：https://nodejs.org/"
    exit 1
fi
echo "✅ Node.js 已安装"

# 检查Go
if ! command -v go &> /dev/null; then
    echo "❌ 未检测到Go语言，请先安装Go 1.21+"
    echo "下载地址：https://golang.org/dl/"
    exit 1
fi
echo "✅ Go语言 已安装"

# 检查MySQL
if ! command -v mysql &> /dev/null; then
    echo "⚠️  未检测到MySQL，请确保MySQL已安装并启动"
    echo "安装命令："
    echo "  macOS: brew install mysql"
    echo "  Ubuntu: sudo apt-get install mysql-server"
    echo
    read -p "是否继续？(y/n): " continue
    if [[ ! "$continue" =~ ^[Yy]$ ]]; then
        exit 1
    fi
else
    echo "✅ MySQL 已安装"
fi

echo
echo "===== 第一步：安装项目依赖 ====="

# 安装主项目依赖
echo "🔄 安装Hardhat依赖..."
npm install
if [ $? -ne 0 ]; then
    echo "❌ Hardhat依赖安装失败！"
    exit 1
fi
echo "✅ Hardhat依赖安装完成"

# 安装Go依赖
echo "🔄 安装Go后端依赖..."
cd nft-go-backend
go mod download
go mod tidy
if [ $? -ne 0 ]; then
    echo "❌ Go依赖安装失败！"
    exit 1
fi
echo "✅ Go依赖安装完成"

# 安装Vue依赖
echo "🔄 安装Vue前端依赖..."
cd static/vue-frontend
npm install
if [ $? -ne 0 ]; then
    echo "❌ Vue依赖安装失败！"
    exit 1
fi
echo "✅ Vue依赖安装完成"

cd ../../..

echo
echo "===== 第二步：检查配置文件 ====="

# 检查.env文件
if [ ! -f "nft-go-backend/.env" ]; then
    echo "⚠️  未找到.env配置文件"
    echo "📝 正在创建默认配置文件..."
    
    cat > nft-go-backend/.env << EOL
# 服务配置
PORT=8080

# 区块链配置
ETHEREUM_RPC=http://localhost:7545
MAIN_NFT_ADDRESS=请更新为实际合约地址
CHILD_NFT_ADDRESS=请更新为实际合约地址
PRIVATE_KEY=请更新为实际私钥
CHAIN_ID=1337

# 数据库配置
DB_USER=root
DB_PASSWORD=123456
DB_HOST=localhost
DB_PORT=3306
DB_NAME=nft_db

# IPFS配置（可选）
IPFS_ACCESS_KEY=your_ipfs_access_key
EOL
    
    echo "✅ 默认配置文件已创建"
    echo "⚠️  请编辑 nft-go-backend/.env 文件，更新实际的配置信息"
else
    echo "✅ 配置文件已存在"
fi

echo
echo "===== 第三步：提示用户配置区块链环境 ====="
echo
echo "🔗 区块链环境配置提醒："
echo "1. 启动Ganache（GUI版本或命令行版本）"
echo "   - 端口：7545"
echo "   - 链ID：1337"
echo "   - 账户数：10个"
echo
echo "2. 配置MetaMask"
echo "   - 网络名称：Local Ganache"
echo "   - RPC URL：http://localhost:7545"
echo "   - 链ID：1337"
echo "   - 货币符号：ETH"
echo
echo "3. 导入Ganache账户私钥到MetaMask"
echo

read -p "✅ Ganache是否已启动？(y/n): " ganache_ready
if [[ ! "$ganache_ready" =~ ^[Yy]$ ]]; then
    echo "⚠️  请先启动Ganache再继续"
    exit 1
fi

echo
echo "===== 第四步：编译和部署智能合约 ====="

echo "🔄 编译智能合约..."
npm run compile
if [ $? -ne 0 ]; then
    echo "❌ 合约编译失败！"
    exit 1
fi
echo "✅ 合约编译完成"

echo "🔄 部署智能合约到本地网络..."
npm run deploy:local
if [ $? -ne 0 ]; then
    echo "❌ 合约部署失败！请检查Ganache是否正在运行"
    exit 1
fi
echo "✅ 合约部署完成"

echo
echo "⚠️  重要：请记录合约部署地址，并更新 nft-go-backend/.env 文件中的："
echo "  - MAIN_NFT_ADDRESS"
echo "  - CHILD_NFT_ADDRESS"
echo "  - PRIVATE_KEY（使用Ganache中的一个账户私钥）"
echo

read -p "✅ 配置是否已更新？(y/n): " config_updated
if [[ ! "$config_updated" =~ ^[Yy]$ ]]; then
    echo "⚠️  请先更新配置文件再继续"
    echo "📝 编辑文件：nft-go-backend/.env"
    exit 1
fi

echo
echo "===== 第五步：初始化数据库 ====="

echo "🔄 创建数据库..."
mysql -u root -p123456 -e "CREATE DATABASE IF NOT EXISTS nft_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;" 2>/dev/null
if [ $? -ne 0 ]; then
    echo "⚠️  数据库创建可能失败，请手动创建数据库 'nft_db'"
    echo "SQL命令：CREATE DATABASE nft_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"
fi
echo "✅ 数据库准备完成"

echo
echo "===== 第六步：启动后端服务 ====="

echo "🚀 正在启动后端服务..."
echo "⚠️  服务启动后请保持此终端打开"
echo "🌐 访问地址：http://localhost:8080"
echo
echo "按Enter键启动服务..."
read

cd nft-go-backend
chmod +x start.sh
./start.sh

echo
echo "===== 启动完成 ====="
echo
echo "🎉 项目启动完成！"
echo "🌐 请访问：http://localhost:8080"
echo
echo "📋 下一步操作："
echo "1. 连接MetaMask钱包"
echo "2. 访问DID和VC管理页面注册身份"
echo "3. 开始使用NFT+ABE+DID/VC功能"
echo
echo "📚 详细使用说明请参考：PROJECT_SETUP_GUIDE.md"
echo 