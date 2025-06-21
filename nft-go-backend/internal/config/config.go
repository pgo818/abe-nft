package config

import (
	"os"
	"strconv"
	"fmt"
	"github.com/joho/godotenv"
)

// Config 结构体
type Config struct {
	EthereumRPC      string
	MainNFTAddress   string
	ChildNFTAddress  string
	PrivateKey       string
	ChainID          int64
	Port             string

	// 数据库配置
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string

	//ipfs
	AcccessKey string
}

// LoadConfig 加载配置
func LoadConfig() (*Config, error) {
	// 加载.env文件
	godotenv.Load()


	return &Config{
		EthereumRPC:     getEnv("ETHEREUM_RPC", "http://localhost:7545"),
		MainNFTAddress:  getEnv("MAIN_NFT_ADDRESS", "0x3b5a6b78d0625d6eb6333e0DA27b75A12Fc5F27D"),
		ChildNFTAddress: getEnv("CHILD_NFT_ADDRESS", "0x38C5f113b716e21C57cc24bDEE237cEd28bA866F"),
		PrivateKey:      getEnv("PRIVATE_KEY", "63435add31c605dfa2ee262dfb1dd019c985c881196309c4d194d3574a0c3fc1"),
		ChainID:         getEnvAsInt64("CHAIN_ID", 1337),
		Port:            getEnv("PORT", "8080"),
		// 数据库配置
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "123456"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBName:     getEnv("DB_NAME", "nft_db"),

		//ipfs
		AcccessKey: getEnv("IPFS_ACCESS_KEY", "NDU5RDlCQUU0NTg5NkYzRDA5Njc6dWdMSll1enZvaTBCWGNOVjZtRnNBcEY3YzVGM2FkZ3R1aWVUVUFTdTphYmUtbmZ0"),
		
	}, nil
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// getEnvAsInt64 获取环境变量并转换为int64
func getEnvAsInt64(key string, defaultValue int64) int64 {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}
	
	return intValue
} 

// GetDSN 返回数据库连接字符串
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}