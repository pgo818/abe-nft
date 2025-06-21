package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fentec-project/gofe/abe"
	"gorm.io/gorm"

	"github.com/ABE/nft/nft-go-backend/internal/models"
)

// ABEService ABE服务结构体
type ABEService struct {
	DB *gorm.DB
}

// NewABEService 创建新的ABE服务
func NewABEService(db *gorm.DB) *ABEService {
	return &ABEService{
		DB: db,
	}
}

// SetupABE 初始化ABE系统
func (s *ABEService) SetupABE(attributes []string, userID uint) (*models.ABESystemKey, error) {
	// 创建FAME密码系统
	fameScheme := abe.NewFAME()

	// 生成主密钥
	pubKey, secKey, err := fameScheme.GenerateMasterKeys()
	if err != nil {
		return nil, fmt.Errorf("生成主密钥失败: %v", err)
	}

	// 序列化密钥
	pubKeyBytes, err := json.Marshal(pubKey)
	if err != nil {
		return nil, fmt.Errorf("序列化公钥失败: %v", err)
	}

	secKeyBytes, err := json.Marshal(secKey)
	if err != nil {
		return nil, fmt.Errorf("序列化私钥失败: %v", err)
	}

	// 序列化属性
	attributesBytes, err := json.Marshal(attributes)
	if err != nil {
		return nil, fmt.Errorf("序列化属性失败: %v", err)
	}

	// 创建系统密钥记录
	systemKey := models.ABESystemKey{
		PubKey:     string(pubKeyBytes),
		SecKey:     string(secKeyBytes),
		Attributes: string(attributesBytes),
		CreatedBy:  userID,
		ExpiresAt:  time.Now().AddDate(1, 0, 0), // 1年后过期
	}

	// 保存到数据库
	if err := s.DB.Create(&systemKey).Error; err != nil {
		return nil, fmt.Errorf("保存系统密钥失败: %v", err)
	}

	return &systemKey, nil
}

// KeyGenABE 生成用户密钥
func (s *ABEService) KeyGenABE(systemKeyID uint, userID uint, userAttributes []string) (*models.ABEUserKey, error) {
	// 获取系统密钥
	var systemKey models.ABESystemKey
	if err := s.DB.First(&systemKey, systemKeyID).Error; err != nil {
		return nil, fmt.Errorf("获取系统密钥失败: %v", err)
	}

	// 反序列化系统密钥
	var pubKey abe.FAMEPubKey
	if err := json.Unmarshal([]byte(systemKey.PubKey), &pubKey); err != nil {
		return nil, fmt.Errorf("反序列化公钥失败: %v", err)
	}

	var secKey abe.FAMESecKey
	if err := json.Unmarshal([]byte(systemKey.SecKey), &secKey); err != nil {
		return nil, fmt.Errorf("反序列化私钥失败: %v", err)
	}

	// 创建FAME密码系统
	fameScheme := abe.NewFAME()

	// 生成用户密钥
	attribKeys, err := fameScheme.GenerateAttribKeys(userAttributes, &secKey)
	if err != nil {
		return nil, fmt.Errorf("生成用户密钥失败: %v", err)
	}

	// 序列化用户密钥
	attribKeysBytes, err := json.Marshal(attribKeys)
	if err != nil {
		return nil, fmt.Errorf("序列化用户密钥失败: %v", err)
	}

	// 序列化用户属性
	userAttributesBytes, err := json.Marshal(userAttributes)
	if err != nil {
		return nil, fmt.Errorf("序列化用户属性失败: %v", err)
	}

	// 创建用户密钥记录
	userKey := models.ABEUserKey{
		UserID:      userID,
		SystemKeyID: systemKeyID,
		AttribKeys:  string(attribKeysBytes),
		Attributes:  string(userAttributesBytes),
		ExpiresAt:   systemKey.ExpiresAt, // 与系统密钥同时过期
	}

	// 保存到数据库
	if err := s.DB.Create(&userKey).Error; err != nil {
		return nil, fmt.Errorf("保存用户密钥失败: %v", err)
	}

	return &userKey, nil
}

// EncryptABE 加密数据
func (s *ABEService) EncryptABE(systemKeyID uint, message string, policy string, userID uint) (*models.ABECiphertext, error) {
	// 获取系统密钥
	var systemKey models.ABESystemKey
	if err := s.DB.First(&systemKey, systemKeyID).Error; err != nil {
		return nil, fmt.Errorf("获取系统密钥失败: %v", err)
	}

	// 反序列化公钥
	var pubKey abe.FAMEPubKey
	if err := json.Unmarshal([]byte(systemKey.PubKey), &pubKey); err != nil {
		return nil, fmt.Errorf("反序列化公钥失败: %v", err)
	}

	// 创建FAME密码系统
	fameScheme := abe.NewFAME()

	// 将策略字符串转换为MSP结构
	msp, err := abe.BooleanToMSP(policy, false)
	if err != nil {
		return nil, fmt.Errorf("转换策略失败: %v", err)
	}

	// 加密消息
	cipher, err := fameScheme.Encrypt(message, msp, &pubKey)
	if err != nil {
		return nil, fmt.Errorf("加密失败: %v", err)
	}

	// 序列化密文
	cipherBytes, err := json.Marshal(cipher)
	if err != nil {
		return nil, fmt.Errorf("序列化密文失败: %v", err)
	}

	// 创建密文记录
	ciphertext := models.ABECiphertext{
		Cipher:      string(cipherBytes),
		Policy:      policy,
		SystemKeyID: systemKeyID,
		CreatedBy:   userID,
	}

	// 保存到数据库
	if err := s.DB.Create(&ciphertext).Error; err != nil {
		return nil, fmt.Errorf("保存密文失败: %v", err)
	}

	return &ciphertext, nil
}

// DecryptABE 解密数据
func (s *ABEService) DecryptABE(ciphertextID uint, userKeyID uint) (string, error) {
	// 获取密文
	var ciphertext models.ABECiphertext
	if err := s.DB.First(&ciphertext, ciphertextID).Error; err != nil {
		return "", fmt.Errorf("获取密文失败: %v", err)
	}

	// 获取用户密钥
	var userKey models.ABEUserKey
	if err := s.DB.First(&userKey, userKeyID).Error; err != nil {
		return "", fmt.Errorf("获取用户密钥失败: %v", err)
	}

	// 验证系统密钥ID是否匹配
	if userKey.SystemKeyID != ciphertext.SystemKeyID {
		return "", fmt.Errorf("用户密钥与密文不匹配")
	}

	// 获取系统密钥
	var systemKey models.ABESystemKey
	if err := s.DB.First(&systemKey, ciphertext.SystemKeyID).Error; err != nil {
		return "", fmt.Errorf("获取系统密钥失败: %v", err)
	}

	// 反序列化公钥
	var pubKey abe.FAMEPubKey
	if err := json.Unmarshal([]byte(systemKey.PubKey), &pubKey); err != nil {
		return "", fmt.Errorf("反序列化公钥失败: %v", err)
	}

	// 反序列化用户密钥
	var attribKeys abe.FAMEAttribKeys
	if err := json.Unmarshal([]byte(userKey.AttribKeys), &attribKeys); err != nil {
		return "", fmt.Errorf("反序列化用户密钥失败: %v", err)
	}

	// 反序列化密文
	var cipher abe.FAMECipher
	if err := json.Unmarshal([]byte(ciphertext.Cipher), &cipher); err != nil {
		return "", fmt.Errorf("反序列化密文失败: %v", err)
	}

	// 创建FAME密码系统
	fameScheme := abe.NewFAME()

	// 解密消息
	message, err := fameScheme.Decrypt(&cipher, &attribKeys, &pubKey)
	if err != nil {
		return "", fmt.Errorf("解密失败: %v", err)
	}

	// 记录操作日志
	operationLog := models.ABEOperation{
		UserID:        userKey.UserID,
		OperationType: "decrypt",
		Details:       fmt.Sprintf("解密密文ID: %d", ciphertextID),
	}
	s.DB.Create(&operationLog)

	return message, nil
}

// LogOperation 记录操作日志
func (s *ABEService) LogOperation(userID uint, operationType string, details map[string]interface{}, ipAddress string) error {
	detailsJSON, err := json.Marshal(details)
	if err != nil {
		return err
	}

	log := &models.ABEOperation{
		UserID:        userID,
		OperationType: operationType,
		Details:       string(detailsJSON),
		IPAddress:     ipAddress,
	}
	return s.DB.Create(log).Error
}

// GetSystemKey 获取系统密钥
func (s *ABEService) GetSystemKey(id uint) (*models.ABESystemKey, error) {
	var systemKey models.ABESystemKey
	if err := s.DB.First(&systemKey, id).Error; err != nil {
		return nil, err
	}
	return &systemKey, nil
}

// GetLatestSystemKey 获取最新的系统密钥
func (s *ABEService) GetLatestSystemKey() (*models.ABESystemKey, error) {
	var systemKey models.ABESystemKey
	if err := s.DB.Order("created_at desc").First(&systemKey).Error; err != nil {
		return nil, err
	}
	return &systemKey, nil
}

// 其他ABE数据库操作方法...
