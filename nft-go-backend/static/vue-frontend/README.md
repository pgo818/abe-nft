# NFT+ABE集成管理平台 - Vue前端

这是NFT+ABE集成管理平台的Vue前端实现，提供了NFT管理、ABE加密系统和DID/VC管理功能。

## 功能特性

- NFT管理
  - 查看所有NFT
  - 铸造NFT
  - 管理元数据
  - 创建子NFT
  - 申请管理

- ABE加密系统
  - 系统初始化
  - 密钥生成
  - 数据加密
  - 数据解密
  
- DID和VC管理
  - 创建和管理DID
  - 颁发和验证VC
  - 医生DID和凭证管理

## 技术栈

- Vue 3
- Vue Router
- Axios
- Bootstrap 5
- Web3.js (用于区块链交互)

## 开发指南

### 安装依赖
```bash
npm install
```

### 启动开发服务器
```bash
npm run serve
```

### 构建生产版本
```bash
npm run build
```

## 项目结构

```
vue-frontend/
├── public/
├── src/
│   ├── assets/          # 静态资源
│   ├── components/      # 可复用组件
│   │   ├── common/      # 通用组件
│   │   ├── nft/         # NFT相关组件
│   │   ├── abe/         # ABE相关组件
│   │   └── did/         # DID/VC相关组件
│   ├── views/           # 页面视图
│   ├── router/          # 路由配置
│   ├── services/        # API服务
│   ├── store/           # 状态管理
│   ├── utils/           # 工具函数
│   ├── App.vue          # 根组件
│   └── main.js          # 入口文件
└── package.json         # 项目配置
```

## 接口文档

详细的API接口文档请参考后端文档。 