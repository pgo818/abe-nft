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

	// 执行安全的数据库迁移
	if err := safeMigration(); err != nil {
		return fmt.Errorf("数据库迁移失败: %w", err)
	}

	log.Println("数据库连接和迁移成功")
	return nil
}

// safeMigration 安全的数据库迁移
func safeMigration() error {
	log.Println("开始安全数据库迁移...")

	// 强制重建NFT元数据表以修复表结构问题
	if err := fixNFTMetadataTable(); err != nil {
		log.Printf("修复NFT元数据表失败: %v", err)
		// 继续执行，不中断迁移
	}

	// 使用安全的迁移方式，不删除现有数据
	if err := migrateDIDTable(); err != nil {
		return fmt.Errorf("迁移DID表失败: %w", err)
	}
	log.Println("DID表迁移完成")

	if err := migrateDoctorTable(); err != nil {
		return fmt.Errorf("迁移Doctor表失败: %w", err)
	}
	log.Println("Doctor表迁移完成")

	// 迁移DoctorVC表
	if err := migrateDoctorVCTable(); err != nil {
		return fmt.Errorf("迁移DoctorVC表失败: %w", err)
	}
	log.Println("DoctorVC表迁移完成")

	// 自动迁移其他表 - 这些都是安全的
	err := DB.AutoMigrate(
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
		&VerifiableCredential{},
		&VerifiablePresentation{},
		&CredentialSchema{},
		&CredentialDefinition{},
	)
	if err != nil {
		return fmt.Errorf("自动迁移其他表失败: %w", err)
	}

	log.Println("所有表迁移完成")
	return nil
}

// migrateDIDTable 迁移DID表
func migrateDIDTable() error {
	// 检查表是否存在
	if !DB.Migrator().HasTable(&DID{}) {
		// 表不存在，直接创建
		return DB.AutoMigrate(&DID{})
	}

	// 表存在，检查列是否存在并且正确
	if !DB.Migrator().HasColumn(&DID{}, "did_string") {
		// 添加did_string列
		if err := DB.Migrator().AddColumn(&DID{}, "did_string"); err != nil {
			return err
		}
	}

	if !DB.Migrator().HasColumn(&DID{}, "wallet_address") {
		// 添加wallet_address列
		if err := DB.Migrator().AddColumn(&DID{}, "wallet_address"); err != nil {
			return err
		}
	}

	if !DB.Migrator().HasColumn(&DID{}, "status") {
		// 添加status列
		if err := DB.Migrator().AddColumn(&DID{}, "status"); err != nil {
			return err
		}
	}

	// 最后执行完整的迁移以确保索引和约束正确
	return DB.AutoMigrate(&DID{})
}

// migrateDoctorTable 迁移Doctor表
func migrateDoctorTable() error {
	// 检查表是否存在
	if !DB.Migrator().HasTable(&Doctor{}) {
		// 表不存在，直接创建
		return DB.AutoMigrate(&Doctor{})
	}

	// 表存在，检查列是否存在并且正确
	if !DB.Migrator().HasColumn(&Doctor{}, "did_string") {
		if err := DB.Migrator().AddColumn(&Doctor{}, "did_string"); err != nil {
			return err
		}
	}

	if !DB.Migrator().HasColumn(&Doctor{}, "wallet_address") {
		if err := DB.Migrator().AddColumn(&Doctor{}, "wallet_address"); err != nil {
			return err
		}
	}

	if !DB.Migrator().HasColumn(&Doctor{}, "name") {
		if err := DB.Migrator().AddColumn(&Doctor{}, "name"); err != nil {
			return err
		}
	}

	if !DB.Migrator().HasColumn(&Doctor{}, "license_number") {
		if err := DB.Migrator().AddColumn(&Doctor{}, "license_number"); err != nil {
			return err
		}
	}

	if !DB.Migrator().HasColumn(&Doctor{}, "status") {
		if err := DB.Migrator().AddColumn(&Doctor{}, "status"); err != nil {
			return err
		}
	}

	if !DB.Migrator().HasColumn(&Doctor{}, "hospital_did") {
		if err := DB.Migrator().AddColumn(&Doctor{}, "hospital_did"); err != nil {
			return err
		}
	}

	// 最后执行完整的迁移以确保索引和约束正确
	return DB.AutoMigrate(&Doctor{})
}

// migrateDoctorVCTable 迁移DoctorVC表
func migrateDoctorVCTable() error {
	// 检查表是否存在
	if !DB.Migrator().HasTable(&DoctorVC{}) {
		// 表不存在，直接创建
		return DB.AutoMigrate(&DoctorVC{})
	}

	// 表存在，检查列是否存在并且正确
	if !DB.Migrator().HasColumn(&DoctorVC{}, "vcid") {
		if err := DB.Migrator().AddColumn(&DoctorVC{}, "vcid"); err != nil {
			return err
		}
	}

	if !DB.Migrator().HasColumn(&DoctorVC{}, "doctor_did") {
		if err := DB.Migrator().AddColumn(&DoctorVC{}, "doctor_did"); err != nil {
			return err
		}
	}

	if !DB.Migrator().HasColumn(&DoctorVC{}, "type") {
		if err := DB.Migrator().AddColumn(&DoctorVC{}, "type"); err != nil {
			return err
		}
	}

	if !DB.Migrator().HasColumn(&DoctorVC{}, "content") {
		if err := DB.Migrator().AddColumn(&DoctorVC{}, "content"); err != nil {
			return err
		}
	}

	if !DB.Migrator().HasColumn(&DoctorVC{}, "issued_at") {
		if err := DB.Migrator().AddColumn(&DoctorVC{}, "issued_at"); err != nil {
			return err
		}
	}

	if !DB.Migrator().HasColumn(&DoctorVC{}, "status") {
		if err := DB.Migrator().AddColumn(&DoctorVC{}, "status"); err != nil {
			return err
		}
	}

	// 最后执行完整的迁移以确保索引和约束正确
	return DB.AutoMigrate(&DoctorVC{})
}

// fixNFTMetadataTable 修复NFT元数据表结构
func fixNFTMetadataTable() error {
	log.Println("开始修复NFT元数据表...")

	// 检查表是否存在
	if !DB.Migrator().HasTable(&NFTMetadataDB{}) {
		log.Println("NFT元数据表不存在，创建新表...")
		return DB.AutoMigrate(&NFTMetadataDB{})
	}

	// 检查关键列是否存在
	requiredColumns := []string{"name", "description", "external_url", "image", "policy", "ciphertext", "ipfs_hash"}

	for _, column := range requiredColumns {
		if !DB.Migrator().HasColumn(&NFTMetadataDB{}, column) {
			log.Printf("NFT元数据表缺少列: %s，正在添加...", column)
			if err := DB.Migrator().AddColumn(&NFTMetadataDB{}, column); err != nil {
				log.Printf("添加列 %s 失败: %v", column, err)
				// 继续处理其他列
			}
		}
	}

	// 执行完整迁移
	if err := DB.AutoMigrate(&NFTMetadataDB{}); err != nil {
		log.Printf("NFT元数据表完整迁移失败: %v", err)
		return err
	}

	log.Println("NFT元数据表修复完成")
	return nil
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return DB
}
