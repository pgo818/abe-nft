package models

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB(dsn string) error {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("连接数据库失败: %w", err)
	}

	// 自动迁移数据库模式 - 包含NFT、ABE和DID/VC模型
	err = DB.AutoMigrate(
		// NFT相关模型
		&NFT{},
		&ChildNFTRequest{},
		&NFTMetadataDB{},
		// ABE相关模型
		&ABESystemKey{},
		&ABEUserKey{},
		&ABECiphertext{},
		&ABEOperation{},
		// DID/VC相关模型
		&DID{},
		&VerifiableCredential{},
		&VerifiablePresentation{},
		&CredentialSchema{},
		&CredentialDefinition{},
		// 医生身份和凭证模型
		&Doctor{},
		&DoctorVC{},
	)
	if err != nil {
		return fmt.Errorf("自动迁移失败: %w", err)
	}

	log.Println("数据库连接和迁移成功")
	return nil
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return DB
}
