# NFT+ABE+DID/VC 集成管理平台

这是一个集成了NFT（非同质化代币）管理、ABE（基于属性的加密）功能和DID/VC（去中心化身份/可验证凭证）的完整平台。该平台允许用户创建和管理NFT资产，提供基于属性的加密功能，实现对数字内容的细粒度访问控制，并支持去中心化身份和可验证凭证的创建、颁发和验证。

## 功能特点

### NFT管理平台
- NFT铸造与管理
- 元数据创建与IPFS存储
- 子NFT创建与管理
- 子NFT申请与审批流程

### ABE加密管理
- ABE系统初始化与属性设置
- 用户属性密钥生成
- 基于访问策略的数据加密
- 基于用户属性的数据解密
- 操作日志记录与监控

### DID和VC管理
- DID创建与管理（支持key、ethr、web方法）
- DID文档解析与更新
- 可验证凭证（VC）的颁发
- VC的验证与撤销
- 可验证表示（VP）的创建与验证
- 凭证管理与状态跟踪

## 系统架构

### 后端技术栈
- Go语言 + Gin框架
- 以太坊区块链集成
- GORM数据库ORM
- ABE加密库 (fentec-project/gofe)

### 前端技术栈
- HTML5 + CSS3
- Bootstrap 5
- JavaScript (原生)
- Web3.js (区块链交互)

### 数据库模型
- NFT相关表: NFT, ChildNFTRequest, NFTMetadataDB
- ABE相关表: ABESystemKey, ABEUserKey, ABECiphertext, ABEOperation
- DID相关表: DID
- VC相关表: VerifiableCredential, CredentialSchema, CredentialDefinition, VerifiablePresentation

## 快速开始

### 环境要求
- Go 1.21+
- MySQL 5.7+
- 以太坊节点 (可使用Ganache本地开发环境)
- MetaMask钱包

### 配置
1. 复制 `.env.example` 文件为 `.env`
2. 根据你的环境配置以下参数:
   ```
   ETHEREUM_RPC=http://localhost:7545
   MAIN_NFT_ADDRESS=你的主NFT合约地址
   CHILD_NFT_ADDRESS=你的子NFT合约地址
   PRIVATE_KEY=服务器账户私钥
   CHAIN_ID=1337
   PORT=9090
   DB_USER=root
   DB_PASSWORD=你的数据库密码
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=nft_db
   IPFS_ACCESS_KEY=你的IPFS访问密钥
   ```

### 安装与运行
1. 克隆代码库
   ```bash
   git clone <repository-url>
   cd nft-go-backend
   ```

2. 安装依赖
   ```bash
   go mod download
   ```

3. 构建项目
   ```bash
   go build -o nft-abe-server ./cmd/server
   ```

4. 运行服务器
   ```bash
   ./nft-abe-server
   ```

5. 访问平台
   打开浏览器访问 `http://localhost:8081`

## 使用指南

### NFT管理平台

#### 铸造NFT
1. 连接MetaMask钱包
2. 点击"NFT管理平台"
3. 在导航栏选择"铸造NFT"
4. 输入NFT的URI (可以使用元数据管理创建)
5. 点击"铸造"按钮

#### 元数据管理
1. 在导航栏选择"元数据管理"
2. 填写元数据信息 (名称、描述、图片URL等)
3. 点击"创建元数据"
4. 获取IPFS哈希，可以直接用于铸造NFT

#### 子NFT管理
1. 在"我的NFT"页面找到要创建子NFT的主NFT
2. 点击"创建子NFT"按钮
3. 输入接收者地址和URI
4. 点击"创建"按钮

#### 子NFT申请
1. 在主页找到想要申请子NFT的主NFT
2. 点击"申请子NFT"按钮
3. 输入URI
4. 点击"提交申请"
5. 等待NFT拥有者审批

#### 申请处理
1. 在导航栏选择"申请管理"
2. 查看待处理的申请
3. 点击"批准"或"拒绝"按钮

### ABE加密管理

#### 系统初始化
1. 点击"ABE加密管理"
2. 在导航栏选择"系统初始化"
3. 输入系统属性列表 (每行一个属性)
4. 点击"初始化系统"
5. 保存生成的系统密钥

#### 生成属性密钥
1. 在导航栏选择"密钥生成"
2. 输入用户属性列表
3. 输入系统公钥和主密钥
4. 点击"生成密钥"
5. 保存生成的属性密钥

#### 加密数据
1. 在导航栏选择"加密"
2. 输入要加密的消息
3. 输入访问策略 (如 "(属性1 AND 属性2) OR 属性3")
4. 输入系统公钥
5. 点击"加密"
6. 保存生成的密文

#### 解密数据
1. 在导航栏选择"解密"
2. 输入密文
3. 输入属性密钥
4. 输入系统公钥
5. 点击"解密"
6. 查看解密后的消息

#### 查看操作日志
1. 在导航栏选择"操作日志"
2. 查看所有ABE操作记录

### DID和VC管理

#### 创建DID
1. 点击"DID和VC管理"
2. 在导航栏选择"创建DID"
3. 选择DID方法（key、ethr、web）
4. 输入控制者地址
5. 点击"创建DID"
6. 保存生成的DID

#### 颁发可验证凭证
1. 在导航栏选择"颁发凭证"
2. 输入颁发者DID（需要先选择一个DID）
3. 输入主体DID（凭证持有者的DID）
4. 输入凭证类型
5. 填写凭证主体内容
6. 点击"颁发凭证"

#### 验证凭证
1. 在凭证管理页面或颁发凭证页面
2. 将凭证JSON粘贴到验证框
3. 点击"验证凭证"
4. 查看验证结果

#### 创建可验证表示
1. 在导航栏选择"创建表示"
2. 选择要包含的凭证
3. 输入验证者DID（可选）
4. 点击"创建表示"
5. 保存生成的表示

## API接口文档

### NFT相关接口
- `GET /api/nft/:tokenId` - 获取NFT信息
- `GET /api/nfts` - 获取所有NFT
- `GET /api/nfts/user/:address` - 获取用户的NFT
- `POST /api/nft/mint` - 铸造NFT
- `POST /api/nft/update-metadata` - 更新NFT元数据
- `POST /api/nft/createChild` - 创建子NFT
- `POST /api/nft/request-child` - 申请子NFT
- `POST /api/nft/process-request` - 处理子NFT申请

### 元数据相关接口
- `POST /api/metadata` - 创建元数据
- `GET /api/metadata/:hash` - 获取元数据
- `GET /api/metadata` - 获取所有元数据

### ABE相关接口
- `POST /api/abe/setup` - 初始化ABE系统
- `POST /api/abe/keygen` - 生成属性密钥
- `POST /api/abe/encrypt` - 加密数据
- `POST /api/abe/decrypt` - 解密数据

### DID相关接口
- `POST /api/did/create` - 创建DID
- `GET /api/did/resolve/:did` - 解析DID文档
- `PUT /api/did/update` - 更新DID
- `POST /api/did/revoke` - 撤销DID
- `GET /api/did/list` - 列出DID

### VC相关接口
- `POST /api/vc/issue` - 颁发可验证凭证
- `POST /api/vc/verify` - 验证凭证
- `POST /api/vc/revoke` - 撤销凭证
- `GET /api/vc/credential/:id` - 获取凭证
- `GET /api/vc/credentials` - 列出凭证
- `POST /api/vc/presentation/create` - 创建可验证表示
- `POST /api/vc/presentation/verify` - 验证表示
- `GET /api/vc/presentation/:id` - 获取表示
- `GET /api/vc/presentations` - 列出表示

## 安全注意事项

1. 在生产环境中，请确保:
   - 使用HTTPS保护API通信
   - 妥善保管ABE系统主密钥
   - 实施更严格的用户认证机制

2. ABE密钥管理:
   - 系统主密钥应由可信管理员安全保存
   - 用户属性密钥应安全传输给相应用户
   - 考虑实施密钥轮换机制

## 贡献指南

欢迎贡献代码、报告问题或提出新功能建议。请遵循以下步骤:

1. Fork项目
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建Pull Request

## 许可证

[MIT License](LICENSE)

## 联系方式

如有问题或建议，请通过以下方式联系我们:
- 项目主页: [GitHub项目地址]
- 电子邮件: [your-email@example.com]

## 医生身份DID与凭证VC系统

我们实现了一套医生身份与凭证管理系统，包括以下功能：

### 医生身份创建

- 医生（账户）通过`createDID`函数生成DID，地址自动作为身份标识
- 存储医生姓名和执业编号
- 身份信息上链并与钱包地址关联

### VC 签发流程

- 医院（固定地址0x1234...）可调用`issueVC`函数
- 传入医生DID、VC类型（如"执业资格"）和内容
- 智能合约自动生成VC唯一ID，关联医生DID
- 标记凭证为有效状态并记录上链存储

### 身份验证场景

- 医院系统可调用`verifyVC`函数，输入VC ID进行验证
- 验证VC是否有效（未被吊销）
- 验证VC持有者DID是否真实存在
- 返回验证结果

## 快速开始

访问 [http://localhost:8080/doctor-did](http://localhost:8080/doctor-did) 即可使用医生DID和VC系统。

### API 接口

#### 医生身份API

- `POST /api/did/doctor/create` - 创建医生DID
  ```json
  {
    "walletAddress": "0x...",
    "name": "张三",
    "licenseNumber": "12345678"
  }
  ```

#### 医生凭证API

- `POST /api/vc/doctor/issue` - 颁发医生凭证
  ```json
  {
    "issuerDid": "0x1234",
    "doctorDid": "did:ethr:0x...",
    "vcType": "执业资格",
    "vcContent": "凭证内容..."
  }
  ```

- `POST /api/vc/doctor/verify` - 验证医生凭证
  ```json
  {
    "vcId": "vc:uuid:..."
  }
  ```

- `POST /api/vc/doctor/list` - 获取医生所有凭证
  ```json
  {
    "doctorDid": "did:ethr:0x..."
  }
  ```

## 运行项目

1. 确保安装了Go 1.17+和MySQL
2. 配置 .env 文件
3. 运行以下命令:

```bash
# Windows
start.bat

# Linux/Mac
./start.sh
```

## 系统架构

该系统基于区块链和Web3技术栈构建:

- Go语言后端 (Gin框架)
- DID/VC 身份与凭证标准
- MySQL数据库存储
- Web前端交互界面 