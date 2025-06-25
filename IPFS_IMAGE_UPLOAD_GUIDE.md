# ABE系统图片上传功能指南

## 功能概述

在ABE（属性基加密）系统中，我们新增了图片上传功能，允许用户将图片直接上传到IPFS网络，并在创建NFT元数据时使用这些图片。

## 新增的API端点

### 后端API
- **端点**: `POST /api/abe/upload-image`
- **功能**: 上传图片文件到IPFS
- **支持格式**: JPG, JPEG, PNG, GIF, WebP, SVG, BMP
- **文件大小限制**: 最大10MB
- **返回**: IPFS Hash和多个网关访问URL

### 前端服务
- **服务**: `abeService.uploadImage(imageFile)`
- **功能**: 前端图片上传封装
- **验证**: 自动验证文件类型和大小

## 使用方法

### 1. 在ABE加密页面使用

1. 访问ABE加密页面（`/abe/encrypt`）
2. 启用"完整IPFS+NFT工作流程"选项
3. 在NFT元数据部分：
   - 点击"选择图片文件"按钮
   - 选择本地图片文件（支持拖拽）
   - 系统会自动上传到IPFS并显示预览
   - 也可以手动输入图片URL

### 2. 图片预览功能

- 上传成功后会显示图片预览
- 支持移除已选择的图片
- 自动处理IPFS链接转换为HTTP网关链接
- 图片加载失败时显示占位符

### 3. IPFS存储特性

- 图片会上传到本地IPFS节点（端口5001）
- 返回多个IPFS网关URL供选择：
  - `https://dweb.link/ipfs/`
  - `https://cloudflare-ipfs.com/ipfs/`
  - `https://gateway.pinata.cloud/ipfs/`
  - `https://ipfs.io/ipfs/`

## 技术实现

### 后端实现 (`abe_handlers.go`)

```go
func (h *ABEHandlers) UploadImageABE(c *gin.Context) {
    // 1. 解析multipart表单
    // 2. 验证文件类型和大小
    // 3. 上传到本地IPFS节点
    // 4. 返回IPFS Hash和访问URL
}
```

### 前端实现 (`abeService.js`)

```javascript
uploadImage: async (imageFile) => {
    // 1. 验证文件类型和大小
    // 2. 创建FormData
    // 3. 发送POST请求到后端
    // 4. 返回上传结果
}
```

### Vue组件集成 (`ABEEncrypt.vue`)

- 文件选择器组件
- 图片预览组件
- 上传进度指示
- 错误处理和用户反馈

## 使用示例

### 基本用法

1. **选择图片**: 点击"选择图片文件"按钮
2. **自动上传**: 选择文件后自动上传到IPFS
3. **预览确认**: 查看上传的图片预览
4. **继续流程**: 填写其他NFT信息并执行加密

### 手动URL输入

如果不想上传文件，也可以直接输入图片URL：
- HTTP/HTTPS链接
- IPFS链接（`ipfs://QmHash...`）
- 其他公开可访问的图片链接

## 错误处理

### 文件验证错误
- 不支持的文件格式
- 文件大小超过限制
- 文件读取失败

### 网络错误
- IPFS节点连接失败
- 上传超时
- 网络中断

### 用户反馈
- 成功上传显示绿色提示和IPFS Hash
- 失败时显示红色错误信息
- 上传过程中显示加载动画

## 注意事项

1. **IPFS节点要求**: 需要本地运行IPFS节点（默认端口5001）
2. **文件持久性**: 上传到IPFS的文件需要节点保持在线才能访问
3. **网关可用性**: 不同IPFS网关的可用性可能不同
4. **隐私考虑**: 上传到IPFS的内容是公开可访问的

## 集成工作流程

当启用"完整IPFS+NFT工作流程"时：

1. 用户选择或输入图片
2. 上传原始数据到IPFS
3. 加密文件Hash
4. 上传密文到IPFS
5. 创建包含图片的NFT元数据
6. 上传元数据到IPFS

这样确保了整个过程中图片、数据和元数据都存储在分布式网络中。 