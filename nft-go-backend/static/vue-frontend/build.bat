@echo off
echo 正在安装依赖...
call npm install

echo 正在构建生产版本...
call npm run build

echo 构建完成！输出目录: ../dist 