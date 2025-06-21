package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"github.com/ABE/nft/nft-go-backend/internal/config"
	"github.com/ABE/nft/nft-go-backend/pkg/childnft"
	"github.com/ABE/nft/nft-go-backend/pkg/mainnft"
)

// EthClient 以太坊客户端结构体
type EthClient struct {
	Client     *ethclient.Client
	MainNFT    *mainnft.Mainnft
	ChildNFT   *childnft.Childnft
	Auth       *bind.TransactOpts
	CallOpts   *bind.CallOpts
	PrivateKey *ecdsa.PrivateKey
	Config     *config.Config
}

// NewEthClient 创建新的以太坊客户端
func NewEthClient(cfg *config.Config) (*EthClient, error) {
	// 连接到以太坊节点
	client, err := ethclient.Dial(cfg.EthereumRPC)
	if err != nil {
		return nil, fmt.Errorf("无法连接到以太坊客户端: %v", err)
	}

	// 加载私钥
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("无法加载私钥: %v", err)
	}

	// 从私钥获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("无法获取公钥")
	}

	// 从公钥获取地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	log.Println("使用地址:", fromAddress.Hex())

	// 创建交易选项
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(cfg.ChainID))
	if err != nil {
		return nil, fmt.Errorf("无法创建交易选项: %v", err)
	}

	// 设置默认gas限制
	auth.GasLimit = uint64(3000000)

	// 创建只读调用选项
	callOpts := &bind.CallOpts{
		Pending: false,
		From:    fromAddress,
		Context: context.Background(),
	}

	// 创建合约实例
	mainNFTAddress := common.HexToAddress(cfg.MainNFTAddress)
	childNFTAddress := common.HexToAddress(cfg.ChildNFTAddress)

	mainNFT, err := mainnft.NewMainnft(mainNFTAddress, client)
	if err != nil {
		return nil, fmt.Errorf("无法创建MainNFT实例: %v", err)
	}

	childNFT, err := childnft.NewChildnft(childNFTAddress, client)
	if err != nil {
		return nil, fmt.Errorf("无法创建ChildNFT实例: %v", err)
	}

	return &EthClient{
		Client:     client,
		MainNFT:    mainNFT,
		ChildNFT:   childNFT,
		Auth:       auth,
		CallOpts:   callOpts,
		PrivateKey: privateKey,
		Config:     cfg,
	}, nil
}

// UpdateTransactOpts 更新交易选项（nonce和gasPrice）
func (ec *EthClient) UpdateTransactOpts() error {
	// 获取nonce
	nonce, err := ec.Client.PendingNonceAt(context.Background(), ec.Auth.From)
	if err != nil {
		return fmt.Errorf("无法获取nonce: %v", err)
	}
	ec.Auth.Nonce = big.NewInt(int64(nonce))

	// 获取当前gas价格
	gasPrice, err := ec.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("无法获取gas价格: %v", err)
	}
	ec.Auth.GasPrice = gasPrice

	return nil
}

// CheckTokenExists 检查token是否存在
func (ec *EthClient) CheckTokenExists(tokenID *big.Int) (bool, error) {
	// 先检查总供应量
	totalSupply, err := ec.MainNFT.TotalSupply(ec.CallOpts)
	if err != nil {
		return false, fmt.Errorf("无法获取总供应量: %v", err)
	}

	// 如果tokenID >= totalSupply，则token不存在
	if tokenID.Cmp(totalSupply) >= 0 {
		return false, nil
	}

	// 尝试获取owner，如果失败则token不存在
	_, err = ec.MainNFT.OwnerOf(ec.CallOpts, tokenID)
	if err != nil {
		// 如果错误信息包含"nonexistent"或"invalid"，说明token不存在
		return false, nil
	}

	return true, nil
}

// GetNFTInfo 获取NFT信息（改进版，增加token存在性检查）
func (ec *EthClient) GetNFTInfo(tokenID *big.Int) (string, string, string, error) {
	// 首先检查token是否存在
	exists, err := ec.CheckTokenExists(tokenID)
	if err != nil {
		return "", "", "", fmt.Errorf("检查token存在性失败: %v", err)
	}

	if !exists {
		return "", "", "", fmt.Errorf("Token ID %s 不存在", tokenID.String())
	}

	// 获取owner
	owner, err := ec.MainNFT.OwnerOf(ec.CallOpts, tokenID)
	fmt.Println("owner:", owner)
	if err != nil {
		return "", "", "", fmt.Errorf("获取owner失败: %v", err)
	}

	// 获取URI
	uri, err := ec.MainNFT.TokenURI(ec.CallOpts, tokenID)
	if err != nil {
		return "", "", "", fmt.Errorf("获取URI失败: %v", err)
	}

	// 获取总供应量
	totalSupply, err := ec.MainNFT.TotalSupply(ec.CallOpts)
	if err != nil {
		return "", "", "", fmt.Errorf("获取总供应量失败: %v", err)
	}

	return owner.Hex(), uri, totalSupply.String(), nil
}

// GetNextTokenID 获取下一个将要铸造的TokenID
func (ec *EthClient) GetNextTokenID() (*big.Int, error) {
	totalSupply, err := ec.MainNFT.TotalSupply(ec.CallOpts)
	if err != nil {
		return nil, fmt.Errorf("无法获取总供应量: %v", err)
	}

	// 下一个TokenID通常等于当前总供应量（从0开始计数）
	return totalSupply, nil
}

// ListAllTokens 列出所有已铸造的tokens
func (ec *EthClient) ListAllTokens() ([]*big.Int, error) {
	totalSupply, err := ec.MainNFT.TotalSupply(ec.CallOpts)
	if err != nil {
		return nil, fmt.Errorf("无法获取总供应量: %v", err)
	}

	var tokens []*big.Int
	for i := int64(0); i < totalSupply.Int64(); i++ {
		tokenID := big.NewInt(i)
		exists, err := ec.CheckTokenExists(tokenID)
		if err != nil {
			log.Printf("检查token %d 时出错: %v", i, err)
			continue
		}
		if exists {
			tokens = append(tokens, tokenID)
		}
	}

	return tokens, nil
}

// PerformContractOperation 使用平台私钥执行合约操作
func (ec *EthClient) PerformContractOperation(operation func(*bind.TransactOpts) (common.Hash, error)) (string, error) {
	// 更新交易选项
	if err := ec.UpdateTransactOpts(); err != nil {
		return "", err
	}

	// 执行操作
	txHash, err := operation(ec.Auth)
	if err != nil {
		return "", err
	}

	return txHash.Hex(), nil
}

// MintNFT 铸造NFT - 使用mintTo函数直接铸造给指定用户
func (ec *EthClient) MintNFT(userAddress string, uri string) (string, error) {
	// 验证用户地址格式
	if !common.IsHexAddress(userAddress) {
		return "", fmt.Errorf("无效的用户地址格式: %s", userAddress)
	}

	operation := func(auth *bind.TransactOpts) (common.Hash, error) {
		// 使用mintTo函数，将NFT直接铸造给指定用户
		tx, err := ec.MainNFT.MintTo(auth, common.HexToAddress(userAddress), uri)
		if err != nil {
			return common.Hash{}, fmt.Errorf("铸造NFT失败: %v", err)
		}
		return tx.Hash(), nil
	}

	return ec.PerformContractOperation(operation)
}

// CreateChildNFT 创建子NFT - 平台代表用户创建
func (ec *EthClient) CreateChildNFT(parentTokenID *big.Int, userAddress string, recipient common.Address, uri string) (string, error) {
	// 验证用户地址格式
	if !common.IsHexAddress(userAddress) {
		return "", fmt.Errorf("无效的用户地址格式: %s", userAddress)
	}
	operation := func(auth *bind.TransactOpts) (common.Hash, error) {
		tx, err := ec.MainNFT.CreateChildNFTWithURI(auth, common.HexToAddress(userAddress), parentTokenID, recipient, uri)
		if err != nil {
			return common.Hash{}, fmt.Errorf("创建子NFT失败: %v", err)
		}
		return tx.Hash(), nil
	}

	return ec.PerformContractOperation(operation)
}

// GetWalletAddressFromContext 从Gin上下文获取钱包地址
func (ec *EthClient) GetWalletAddressFromContext(c *gin.Context) (string, bool) {
	walletAddress, exists := c.Get("walletAddress")
	if !exists {
		return "", false
	}
	return walletAddress.(string), true
}

// GetCurrentTotalSupply 获取当前NFT总供应量
func (ec *EthClient) GetCurrentTotalSupply() (*big.Int, error) {
	return ec.MainNFT.TotalSupply(ec.CallOpts)
}

// GetServerAddress 获取服务器钱包地址
func (ec *EthClient) GetServerAddress() string {
	return ec.Auth.From.Hex()
}

// GetCallerAddress 获取当前账户地址
func (c *EthClient) GetCallerAddress(ctx *gin.Context) string {
	// 这里可以从JWT token或其他认证方式获取地址
	// 简单示例，从请求头获取
	return ctx.GetHeader("X-Wallet-Address")
}

// GetChildNFTInfo 获取子NFT信息
func (ec *EthClient) GetChildNFTInfo(tokenID *big.Int) (string, string, error) {
	// 获取子NFT所有者
	owner, err := ec.ChildNFT.OwnerOf(ec.CallOpts, tokenID)
	if err != nil {
		return "", "", fmt.Errorf("获取子NFT所有者失败: %v", err)
	}

	// 获取子NFT URI
	uri, err := ec.ChildNFT.TokenURI(ec.CallOpts, tokenID)
	if err != nil {
		return "", "", fmt.Errorf("获取子NFT URI失败: %v", err)
	}

	return owner.Hex(), uri, nil
}

// UpdateMainNFTMetadata 更新主NFT元数据
func (ec *EthClient) UpdateMainNFTMetadata(tokenID *big.Int, newURI string) (string, error) {
	operation := func(auth *bind.TransactOpts) (common.Hash, error) {
		tx, err := ec.MainNFT.SetSpecificTokenURI(auth, auth.From, tokenID, newURI)
		if err != nil {
			return common.Hash{}, fmt.Errorf("更新主NFT元数据失败: %v", err)
		}
		return tx.Hash(), nil
	}

	return ec.PerformContractOperation(operation)
}

// UpdateChildNFTMetadata 更新子NFT元数据
func (ec *EthClient) UpdateChildNFTMetadata(tokenID *big.Int, newURI string) (string, error) {
	operation := func(auth *bind.TransactOpts) (common.Hash, error) {
		tx, err := ec.ChildNFT.SetSpecificTokenURI(auth, tokenID, newURI)
		if err != nil {
			return common.Hash{}, fmt.Errorf("更新子NFT元数据失败: %v", err)
		}
		return tx.Hash(), nil
	}

	return ec.PerformContractOperation(operation)
}

// MintNFTToSelf 铸造NFT给平台自己（原有的mint函数保留，以防某些场景需要）
func (ec *EthClient) MintNFTToSelf(uri string) (string, error) {
	operation := func(auth *bind.TransactOpts) (common.Hash, error) {
		tx, err := ec.MainNFT.MintTo(auth, auth.From, uri)
		if err != nil {
			return common.Hash{}, fmt.Errorf("铸造NFT失败: %v", err)
		}
		return tx.Hash(), nil
	}

	return ec.PerformContractOperation(operation)
}

