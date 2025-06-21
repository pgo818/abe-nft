@echo off
echo 正在构建Vue前端...
call npm run build

echo 正在准备集成...
if not exist "..\dist" mkdir "..\dist"

echo 正在集成到后端...
xcopy "dist\*" "..\dist\" /E /Y

echo 集成完成！
echo 现在可以启动Go后端服务器提供完整应用。 