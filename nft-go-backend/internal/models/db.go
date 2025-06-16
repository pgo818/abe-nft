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

	// 自动迁移数据库模式
	err = DB.AutoMigrate(&NFT{}, &ChildNFTRequest{}, &NFTMetadataDB{})
	if err != nil {
		return fmt.Errorf("自动迁移失败: %w", err)
	}

	log.Println("数据库连接和迁移成功")
	return nil
}
