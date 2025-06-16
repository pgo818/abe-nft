package blockchain

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ListenToEvents 监听合约事件
func (ec *EthClient) ListenToEvents() {
	// 创建查询过滤器
	mainNFTAddress := common.HexToAddress(ec.Config.MainNFTAddress)
	childNFTAddress := common.HexToAddress(ec.Config.ChildNFTAddress)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			mainNFTAddress,
			childNFTAddress,
		},
	}

	// 创建日志通道
	logs := make(chan types.Log)

	// 订阅日志
	sub, err := ec.Client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Printf("订阅事件失败: %v", err)
		return
	}

	log.Println("开始监听区块链事件...")

	// 处理日志
	for {
		select {
		case err := <-sub.Err():
			log.Printf("事件订阅错误: %v", err)
			return
		case vLog := <-logs:
			ec.processLog(vLog)
		}
	}
}

// processLog 处理日志事件
func (ec *EthClient) processLog(vLog types.Log) {
	// 尝试解析Transfer事件
	transferEvent, err := ec.MainNFT.ParseTransfer(vLog)
	if err == nil {
		zeroAddr := common.HexToAddress("0x0000000000000000000000000000000000000000")
		if transferEvent.From == zeroAddr {
			log.Printf("新NFT铸造: Token ID %s 给 %s",
				transferEvent.TokenId.String(),
				transferEvent.To.Hex())
		} else {
			log.Printf("NFT转移: 从 %s 到 %s, Token ID %s",
				transferEvent.From.Hex(),
				transferEvent.To.Hex(),
				transferEvent.TokenId.String())
		}
		return
	}

	// 尝试解析ChildNFTCreated事件
	childEvent, err := ec.MainNFT.ParseChildNFTCreated(vLog)
	if err == nil {
		log.Printf("子NFT创建: 父Token ID %s, 接收者 %s",
			childEvent.TokenId.String(),
			childEvent.Receiver.Hex())
		return
	}

	// 尝试解析子NFT的Transfer事件
	childTransferEvent, err := ec.ChildNFT.ParseTransfer(vLog)
	if err == nil {
		zeroAddr := common.HexToAddress("0x0000000000000000000000000000000000000000")
		if childTransferEvent.From == zeroAddr {
			log.Printf("新子NFT铸造: Token ID %s 给 %s",
				childTransferEvent.TokenId.String(),
				childTransferEvent.To.Hex())
		} else {
			log.Printf("子NFT转移: 从 %s 到 %s, Token ID %s",
				childTransferEvent.From.Hex(),
				childTransferEvent.To.Hex(),
				childTransferEvent.TokenId.String())
		}
		return
	}

	// 其他事件
	log.Printf("收到未解析的事件: %v", vLog)
}

// FetchMetadata 获取NFT元数据
func (ec *EthClient) FetchMetadata(uri string) (map[string]interface{}, error) {
	// 这里可以实现HTTP请求获取元数据
	// 为简化，返回一个示例元数据
	return map[string]interface{}{
		"name":        "示例NFT",
		"description": "这是一个示例NFT元数据",
		"image":       "https://example.com/image.png",
		"attributes": []map[string]interface{}{
			{
				"trait_type": "类型",
				"value":      "示例",
			},
		},
	}, nil
}
