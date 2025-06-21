#!/bin/bash

# 安装依赖
echo "正在安装依赖..."
npm install

# 构建生产版本
echo "正在构建生产版本..."
npm run build

echo "构建完成！输出目录: ../dist" 