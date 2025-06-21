#!/bin/bash

# NFT+ABE+DID/VC 集成后端启动脚本

echo "===== NFT+ABE+DID/VC 集成后端启动脚本 ====="

# 设置环境变量
export PORT=9090
export ETHEREUM_RPC=http://localhost:7545
export MAIN_NFT_ADDRESS=0x3b5a6b78d0625d6eb6333e0DA27b75A12Fc5F27D
export CHILD_NFT_ADDRESS=0x38C5f113b716e21C57cc24bDEE237cEd28bA866F
export PRIVATE_KEY=63435add31c605dfa2ee262dfb1dd019c985c881196309c4d194d3574a0c3fc1
export CHAIN_ID=1337
export DB_USER=root
export DB_PASSWORD=123456
export DB_HOST=localhost
export DB_PORT=3306
export DB_NAME=nft_db
export IPFS_ACCESS_KEY=NDU5RDlCQUU0NTg5NkYzRDA5Njc6dWdMSll1enZvaTBCWGNOVjZtRnNBcEY3YzVGM2FkZ3R1aWVUVUFTdTphYmUtbmZ0

echo "环境变量设置完成..."

# 检查数据库是否存在，不存在则创建
echo "检查数据库..."
mysql -u$DB_USER -p$DB_PASSWORD -e "CREATE DATABASE IF NOT EXISTS $DB_NAME;" 2>/dev/null
if [ $? -ne 0 ]; then
    echo "数据库检查失败，请确保MySQL服务已启动且用户名密码正确"
    read -p "按Enter键继续..."
    exit 1
fi
echo "数据库检查完成..."

# 构建并启动服务
echo "构建并启动服务..."
cd cmd/server
go build -o nft-abe-server
if [ $? -ne 0 ]; then
    echo "构建失败，请检查代码是否有错误"
    read -p "按Enter键继续..."
    exit 1
fi

echo "服务构建完成，正在启动..."
./nft-abe-server

# 如果服务退出，暂停以便查看错误信息
read -p "按Enter键继续..." 