package service

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/ABE/nft/nft-go-backend/internal/models"
)

// DIDService DID服务结构体
type DIDService struct {
	DB *gorm.DB
}

// NewDIDService 创建新的DID服务
func NewDIDService(db *gorm.DB) *DIDService {
	return &DIDService{
		DB: db,
	}
}

// CreateDID 创建DID（已弃用，只保留兼容性）
func (s *DIDService) CreateDID(method, controllerAddress, publicKey string) (*models.DIDDocument, error) {
	return nil, fmt.Errorf("此方法已弃用，请使用钱包地址创建DID")
}

// ResolveDID 解析DID
func (s *DIDService) ResolveDID(didString string) (*models.DIDDocument, error) {
	// 查询DID记录
	var did models.DID
	if err := s.DB.Where("did_string = ?", didString).First(&did).Error; err != nil {
		return nil, fmt.Errorf("DID不存在: %v", err)
	}

	// 创建DID文档
	doc := models.DIDDocument{
		Context: []string{"https://www.w3.org/ns/did/v1"},
		ID:      did.DIDString,
		Controller: []string{
			did.WalletAddress,
		},
		VerificationMethod: []models.VerificationMethod{
			{
				ID:           fmt.Sprintf("%s#keys-1", did.DIDString),
				Type:         "EcdsaSecp256k1VerificationKey2019",
				Controller:   did.DIDString,
				PublicKeyHex: "", // 这里应该从数据库获取公钥
			},
		},
		Authentication: []string{
			fmt.Sprintf("%s#keys-1", did.DIDString),
		},
		Created: did.CreatedAt.UTC().Format(time.RFC3339),
		Updated: did.UpdatedAt.UTC().Format(time.RFC3339),
	}

	return &doc, nil
}

// UpdateDID 更新DID（已弃用）
func (s *DIDService) UpdateDID(didString, controllerAddress, publicKey string) (*models.DIDDocument, error) {
	return nil, fmt.Errorf("此方法已弃用")
}

// RevokeDID 撤销DID（已弃用）
func (s *DIDService) RevokeDID(didString, controllerAddress string) error {
	return fmt.Errorf("此方法已弃用")
}

// CreateDIDFromWallet 从钱包地址创建DID
func (s *DIDService) CreateDIDFromWallet(walletAddress string) (*models.DID, bool, error) {
	// 验证钱包地址
	if walletAddress == "" {
		return nil, false, fmt.Errorf("钱包地址不能为空")
	}

	// 检查是否已存在
	var existingDID models.DID
	result := s.DB.Where("wallet_address = ? AND status = ?", walletAddress, "active").First(&existingDID)
	if result.Error == nil {
		// 已存在
		return &existingDID, true, nil
	} else if result.Error != gorm.ErrRecordNotFound {
		// 查询出错
		return nil, false, fmt.Errorf("查询DID失败: %v", result.Error)
	}

	// 不存在，创建新DID
	didID := fmt.Sprintf("did:ethr:%s", walletAddress)

	// 确保所有必需字段都有值
	newDID := models.DID{
		DIDString:     didID,
		WalletAddress: walletAddress,
		Status:        "active",
	}

	// 使用事务确保数据一致性
	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, false, fmt.Errorf("开始事务失败: %v", err)
	}

	if err := tx.Create(&newDID).Error; err != nil {
		tx.Rollback()
		return nil, false, fmt.Errorf("创建DID失败: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, false, fmt.Errorf("提交事务失败: %v", err)
	}

	return &newDID, false, nil
}

// GetDIDByWallet 通过钱包地址获取DID
func (s *DIDService) GetDIDByWallet(walletAddress string) (*models.DID, error) {
	var did models.DID
	if err := s.DB.Where("wallet_address = ? AND status = ?", walletAddress, "active").First(&did).Error; err != nil {
		return nil, fmt.Errorf("未找到DID: %v", err)
	}
	return &did, nil
}

// ListDIDsByWallet 获取指定钱包地址创建的所有DID
func (s *DIDService) ListDIDsByWallet(walletAddress string) ([]models.DID, error) {
	if walletAddress == "" {
		return nil, fmt.Errorf("钱包地址不能为空")
	}

	var dids []models.DID
	if err := s.DB.Where("wallet_address = ? AND status = ?", walletAddress, "active").Find(&dids).Error; err != nil {
		return nil, fmt.Errorf("获取DID列表失败: %v", err)
	}
	return dids, nil
}

// GetAllDIDs 获取所有活跃的DID
func (s *DIDService) GetAllDIDs() ([]models.DID, error) {
	var dids []models.DID
	if err := s.DB.Where("status = ?", "active").Find(&dids).Error; err != nil {
		return nil, fmt.Errorf("获取DID列表失败: %v", err)
	}

	return dids, nil
}

// CreateDoctorDID 创建医生DID
func (s *DIDService) CreateDoctorDID(walletAddress, name, licenseNumber string) (*models.Doctor, error) {
	// 验证参数
	if walletAddress == "" || name == "" || licenseNumber == "" {
		return nil, fmt.Errorf("钱包地址、姓名和执业编号不能为空")
	}

	// 检查是否已存在
	var existingDoctor models.Doctor
	result := s.DB.Where("wallet_address = ? AND status = ?", walletAddress, "active").First(&existingDoctor)
	if result.Error == nil {
		// 已存在
		return &existingDoctor, nil
	} else if result.Error != gorm.ErrRecordNotFound {
		// 查询出错
		return nil, fmt.Errorf("查询医生DID失败: %v", result.Error)
	}

	// 不存在，创建新医生DID
	didID := fmt.Sprintf("did:ethr:%s", walletAddress)

	// 使用事务确保数据一致性
	tx := s.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, fmt.Errorf("开始事务失败: %v", err)
	}

	// 创建DID记录
	newDID := models.DID{
		DIDString:     didID,
		WalletAddress: walletAddress,
		Status:        "active",
	}

	if err := tx.Create(&newDID).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("创建DID记录失败: %v", err)
	}

	// 创建医生记录
	newDoctor := models.Doctor{
		DIDString:     didID,
		WalletAddress: walletAddress,
		Name:          name,
		LicenseNumber: licenseNumber,
		Status:        "active",
	}

	if err := tx.Create(&newDoctor).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("创建医生记录失败: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("提交事务失败: %v", err)
	}

	return &newDoctor, nil
}

// GetDoctorByDID 通过DID获取医生信息
func (s *DIDService) GetDoctorByDID(doctorDID string) (*models.Doctor, error) {
	var doctor models.Doctor
	if err := s.DB.Where("did_string = ? AND status = ?", doctorDID, "active").First(&doctor).Error; err != nil {
		return nil, fmt.Errorf("未找到医生: %v", err)
	}
	return &doctor, nil
}

// GetDoctorByWallet 通过钱包地址获取医生信息
func (s *DIDService) GetDoctorByWallet(walletAddress string) (*models.Doctor, error) {
	var doctor models.Doctor
	if err := s.DB.Where("wallet_address = ? AND status = ?", walletAddress, "active").First(&doctor).Error; err != nil {
		return nil, fmt.Errorf("未找到医生: %v", err)
	}
	return &doctor, nil
}
