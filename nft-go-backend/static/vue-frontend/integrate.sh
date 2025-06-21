#!/bin/bash

# 构建Vue前端
echo "正在构建Vue前端..."
npm run build

# 确保目标目录存在
echo "正在准备集成..."
mkdir -p ../dist

# 复制构建文件到后端静态目录
echo "正在集成到后端..."
cp -r dist/* ../dist/

echo "集成完成！"
echo "现在可以启动Go后端服务器提供完整应用。" 