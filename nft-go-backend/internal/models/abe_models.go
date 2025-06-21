package models

import (
	"time"

	"gorm.io/gorm"
)

// ABESystemKey 系统密钥表
type ABESystemKey struct {
	gorm.Model
	PubKey     string    `gorm:"type:text;not null"`
	SecKey     string    `gorm:"type:text;not null"`
	Attributes string    `gorm:"type:text;not null"`
	CreatedBy  uint      `gorm:"index"`
	ExpiresAt  time.Time `gorm:"index"`
}

// ABEUserKey 用户密钥表
type ABEUserKey struct {
	gorm.Model
	UserID      uint      `gorm:"index;not null"`
	SystemKeyID uint      `gorm:"index;not null"`
	AttribKeys  string    `gorm:"type:text;not null"`
	Attributes  string    `gorm:"type:text;not null"`
	ExpiresAt   time.Time `gorm:"index"`
}

// ABECiphertext 密文表
type ABECiphertext struct {
	gorm.Model
	Cipher      string `gorm:"type:text;not null"`
	Policy      string `gorm:"type:text;not null"`
	SystemKeyID uint   `gorm:"index;not null"`
	CreatedBy   uint   `gorm:"index"`
	NFTID       *uint  `gorm:"index"` // 关联的NFT ID
}

// ABEOperation 操作日志表
type ABEOperation struct {
	gorm.Model
	UserID        uint   `gorm:"index"`
	OperationType string `gorm:"type:varchar(50);not null"`
	Details       string `gorm:"type:text"`
	IPAddress     string `gorm:"type:varchar(50)"`
}
