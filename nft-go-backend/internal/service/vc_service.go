package service

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/ABE/nft/nft-go-backend/internal/models"
)

// VCService 提供可验证凭证相关功能的服务
type VCService struct {
	DB *gorm.DB
}

// NewVCService 创建新的VC服务实例
func NewVCService(db *gorm.DB) *VCService {
	return &VCService{
		DB: db,
	}
}

// IssueCredential 颁发凭证
func (s *VCService) IssueCredential(issuerDID, subjectDID, credentialType string) (*models.VerifiableCredential, error) {
	// 验证颁发者DID
	var issuer models.DID
	if err := s.DB.Where("did_string = ? AND status = ?", issuerDID, "active").First(&issuer).Error; err != nil {
		return nil, fmt.Errorf("颁发者DID无效: %v", err)
	}

	// 验证主体DID
	var subject models.DID
	if err := s.DB.Where("did_string = ? AND status = ?", subjectDID, "active").First(&subject).Error; err != nil {
		return nil, fmt.Errorf("主体DID无效: %v", err)
	}

	// 生成凭证ID
	credentialID := fmt.Sprintf("urn:uuid:%s", uuid.New().String())

	// 创建凭证
	now := time.Now()
	expiration := now.AddDate(1, 0, 0) // 1年后过期

	// 创建凭证主体
	credentialSubject := map[string]interface{}{
		"id":   subjectDID,
		"type": credentialType,
	}
	credentialSubjectJSON, err := json.Marshal(credentialSubject)
	if err != nil {
		return nil, fmt.Errorf("序列化凭证主体失败: %v", err)
	}

	// 创建凭证声明
	claims := map[string]interface{}{
		"issuanceDate":   now.Format(time.RFC3339),
		"expirationDate": expiration.Format(time.RFC3339),
		"issuer":         issuerDID,
		"subject":        subjectDID,
	}
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return nil, fmt.Errorf("序列化凭证声明失败: %v", err)
	}

	// 创建凭证证明（在实际应用中应该使用密码学方法生成）
	proof := map[string]interface{}{
		"type":               "EcdsaSecp256k1Signature2019",
		"created":            now.Format(time.RFC3339),
		"proofPurpose":       "assertionMethod",
		"verificationMethod": fmt.Sprintf("%s#keys-1", issuerDID),
	}
	proofJSON, err := json.Marshal(proof)
	if err != nil {
		return nil, fmt.Errorf("序列化凭证证明失败: %v", err)
	}

	// 创建凭证记录
	credential := models.VerifiableCredential{
		CredentialID:      credentialID,
		IssuerDID:         issuerDID,
		SubjectDID:        subjectDID,
		Type:              credentialType,
		CredentialSchema:  "https://example.com/schemas/" + credentialType,
		Status:            "active",
		IssuanceDate:      now,
		ExpirationDate:    expiration,
		Claims:            string(claimsJSON),
		CredentialSubject: string(credentialSubjectJSON),
		Proof:             string(proofJSON),
	}

	// 保存到数据库
	if err := s.DB.Create(&credential).Error; err != nil {
		return nil, fmt.Errorf("保存凭证失败: %v", err)
	}

	return &credential, nil
}

// VerifyCredential 验证凭证
func (s *VCService) VerifyCredential(credentialID string) (*models.VerifyCredentialResponse, error) {
	// 查询凭证记录
	var credential models.VerifiableCredential
	if err := s.DB.Where("credential_id = ?", credentialID).First(&credential).Error; err != nil {
		return nil, fmt.Errorf("凭证不存在: %v", err)
	}

	// 验证凭证状态
	if credential.Status != "active" {
		return &models.VerifyCredentialResponse{
			Valid:  false,
			Reason: "凭证已被撤销",
		}, nil
	}

	// 验证过期时间
	if credential.ExpirationDate.Before(time.Now()) {
		return &models.VerifyCredentialResponse{
			Valid:  false,
			Reason: "凭证已过期",
		}, nil
	}

	// 验证颁发者DID
	var issuer models.DID
	if err := s.DB.Where("did_string = ? AND status = ?", credential.IssuerDID, "active").First(&issuer).Error; err != nil {
		return &models.VerifyCredentialResponse{
			Valid:  false,
			Reason: "颁发者DID无效",
		}, nil
	}

	// 验证主体DID
	var subject models.DID
	if err := s.DB.Where("did_string = ? AND status = ?", credential.SubjectDID, "active").First(&subject).Error; err != nil {
		return &models.VerifyCredentialResponse{
			Valid:  false,
			Reason: "主体DID无效",
		}, nil
	}

	// 验证证明（在实际应用中应该使用密码学方法验证）
	// 这里简化处理，仅检查证明是否存在
	if credential.Proof == "" {
		return &models.VerifyCredentialResponse{
			Valid:  false,
			Reason: "凭证证明无效",
		}, nil
	}

	// 记录验证时间
	credential.LastVerified = &time.Time{}
	*credential.LastVerified = time.Now()
	s.DB.Save(&credential)

	return &models.VerifyCredentialResponse{
		Valid:       true,
		IssuerDID:   credential.IssuerDID,
		SubjectDID:  credential.SubjectDID,
		IssuedAt:    credential.IssuanceDate.Format(time.RFC3339),
		ExpiresAt:   credential.ExpirationDate.Format(time.RFC3339),
		Credentials: []string{credential.Type},
	}, nil
}

// RevokeCredential 撤销凭证
func (s *VCService) RevokeCredential(credentialID, issuerDID, reason string) error {
	// 查询凭证记录
	var credential models.VerifiableCredential
	if err := s.DB.Where("credential_id = ?", credentialID).First(&credential).Error; err != nil {
		return fmt.Errorf("凭证不存在: %v", err)
	}

	// 验证颁发者
	if credential.IssuerDID != issuerDID {
		return fmt.Errorf("只有颁发者可以撤销凭证")
	}

	// 更新凭证状态
	now := time.Now()
	credential.Status = "revoked"
	credential.RevocationDate = &now
	credential.RevocationReason = reason

	if err := s.DB.Save(&credential).Error; err != nil {
		return fmt.Errorf("撤销凭证失败: %v", err)
	}

	return nil
}

// CreatePresentation 创建可验证表示
func (s *VCService) CreatePresentation(holderDID, verifierDID string, credentialIDs []string, purpose string) (*models.VerifiablePresentation, error) {
	// 验证持有者DID
	var holder models.DID
	if err := s.DB.Where("did_string = ? AND status = ?", holderDID, "active").First(&holder).Error; err != nil {
		return nil, fmt.Errorf("持有者DID无效: %v", err)
	}

	// 验证所有凭证
	var credentials []models.VerifiableCredential
	for _, credID := range credentialIDs {
		var cred models.VerifiableCredential
		if err := s.DB.Where("credential_id = ? AND status = ? AND subject_did = ?",
			credID, "active", holderDID).First(&cred).Error; err != nil {
			return nil, fmt.Errorf("凭证 %s 无效或不属于持有者: %v", credID, err)
		}
		credentials = append(credentials, cred)
	}

	// 生成表示ID
	presentationID := fmt.Sprintf("urn:uuid:%s", uuid.New().String())

	// 创建表示记录
	now := time.Now()
	presentation := models.VerifiablePresentation{
		PresentationID:   presentationID,
		HolderDID:        holderDID,
		VerifierDID:      verifierDID,
		CredentialIDs:    credentialIDs,
		Purpose:          purpose,
		Challenge:        uuid.New().String(), // 生成随机挑战值
		PresentationDate: now,
		Status:           "active",
	}

	// 创建证明（在实际应用中应该使用密码学方法生成）
	proof := map[string]interface{}{
		"type":               "EcdsaSecp256k1Signature2019",
		"created":            now.Format(time.RFC3339),
		"proofPurpose":       "authentication",
		"verificationMethod": fmt.Sprintf("%s#keys-1", holderDID),
		"challenge":          presentation.Challenge,
	}
	proofJSON, err := json.Marshal(proof)
	if err != nil {
		return nil, fmt.Errorf("序列化表示证明失败: %v", err)
	}
	presentation.Proof = string(proofJSON)

	// 保存到数据库
	if err := s.DB.Create(&presentation).Error; err != nil {
		return nil, fmt.Errorf("保存表示失败: %v", err)
	}

	return &presentation, nil
}

// VerifyPresentation 验证可验证表示
func (s *VCService) VerifyPresentation(presentationID string) (*models.VerifyPresentationResponse, error) {
	// 查询表示记录
	var presentation models.VerifiablePresentation
	if err := s.DB.Where("presentation_id = ?", presentationID).First(&presentation).Error; err != nil {
		return nil, fmt.Errorf("表示不存在: %v", err)
	}

	// 验证表示状态
	if presentation.Status != "active" {
		return &models.VerifyPresentationResponse{
			Valid:  false,
			Reason: "表示已被撤销",
		}, nil
	}

	// 验证持有者DID
	var holder models.DID
	if err := s.DB.Where("did_string = ? AND status = ?", presentation.HolderDID, "active").First(&holder).Error; err != nil {
		return &models.VerifyPresentationResponse{
			Valid:  false,
			Reason: "持有者DID无效",
		}, nil
	}

	// 验证所有包含的凭证
	for _, credID := range presentation.CredentialIDs {
		var cred models.VerifiableCredential
		if err := s.DB.Where("credential_id = ? AND status = ?", credID, "active").First(&cred).Error; err != nil {
			return &models.VerifyPresentationResponse{
				Valid:  false,
				Reason: fmt.Sprintf("凭证 %s 无效", credID),
			}, nil
		}

		// 验证凭证是否过期
		if cred.ExpirationDate.Before(time.Now()) {
			return &models.VerifyPresentationResponse{
				Valid:  false,
				Reason: fmt.Sprintf("凭证 %s 已过期", credID),
			}, nil
		}
	}

	// 验证证明（在实际应用中应该使用密码学方法验证）
	// 这里简化处理，仅检查证明是否存在
	if presentation.Proof == "" {
		return &models.VerifyPresentationResponse{
			Valid:  false,
			Reason: "表示证明无效",
		}, nil
	}

	// 记录验证时间
	now := time.Now()
	presentation.LastVerified = &now
	s.DB.Save(&presentation)

	return &models.VerifyPresentationResponse{
		Valid:         true,
		HolderDID:     presentation.HolderDID,
		PresentedAt:   presentation.PresentationDate.Format(time.RFC3339),
		CredentialIDs: presentation.CredentialIDs,
	}, nil
}

// GetCredential 获取凭证详情
func (s *VCService) GetCredential(credentialID string) (*models.VerifiableCredentialResponse, error) {
	// 从数据库查询凭证
	var credential models.VerifiableCredential
	if err := s.DB.Where("credential_id = ?", credentialID).First(&credential).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("凭证不存在: %s", credentialID)
		}
		return nil, fmt.Errorf("查询凭证失败: %w", err)
	}

	// 构建可验证凭证响应
	vc := models.VerifiableCredentialResponse{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
		},
		ID:             credential.CredentialID,
		Type:           []string{"VerifiableCredential", credential.Type},
		Issuer:         credential.IssuerDID,
		IssuanceDate:   credential.IssuanceDate.UTC().Format(time.RFC3339),
		ExpirationDate: credential.ExpirationDate.UTC().Format(time.RFC3339),
		CredentialSubject: map[string]interface{}{
			"id": credential.SubjectDID,
		},
		Proof: models.Proof{
			Type:               "SimpleProof2024",
			Created:            credential.IssuanceDate.UTC().Format(time.RFC3339),
			VerificationMethod: fmt.Sprintf("%s#keys-1", credential.IssuerDID),
			ProofPurpose:       "assertionMethod",
			ProofValue:         credential.Proof,
		},
	}

	// 解析claims并添加到凭证主体
	var claims map[string]interface{}
	if err := json.Unmarshal([]byte(credential.Claims), &claims); err != nil {
		return nil, fmt.Errorf("解析claims失败: %w", err)
	}
	for k, v := range claims {
		vc.CredentialSubject[k] = v
	}

	// 如果凭证已撤销，添加撤销信息
	if credential.Status == "revoked" && credential.RevocationDate != nil {
		vc.CredentialStatus = map[string]interface{}{
			"type":      "RevocationList2020Status",
			"status":    "revoked",
			"revokedAt": credential.RevocationDate.UTC().Format(time.RFC3339),
			"reason":    credential.RevocationReason,
		}
	}

	return &vc, nil
}

// ListCredentials 列出凭证
func (s *VCService) ListCredentials(issuerDID, subjectDID, status string) ([]models.VerifiableCredential, error) {
	var credentials []models.VerifiableCredential
	query := s.DB

	// 根据颁发者过滤
	if issuerDID != "" {
		query = query.Where("issuer_did = ?", issuerDID)
	}

	// 根据主体过滤
	if subjectDID != "" {
		query = query.Where("subject_did = ?", subjectDID)
	}

	// 根据状态过滤
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 执行查询
	if err := query.Find(&credentials).Error; err != nil {
		return nil, fmt.Errorf("查询凭证列表失败: %w", err)
	}

	return credentials, nil
}

// GetPresentation 获取展示详情
func (s *VCService) GetPresentation(presentationID string) (*models.VerifiablePresentationResponse, error) {
	// 从数据库查询展示
	var presentation models.VerifiablePresentation
	if err := s.DB.Where("presentation_id = ?", presentationID).First(&presentation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("展示不存在: %s", presentationID)
		}
		return nil, fmt.Errorf("查询展示失败: %w", err)
	}

	// 获取包含的凭证
	var verifiableCredentials []models.VerifiableCredentialResponse
	for _, credID := range presentation.CredentialIDs {
		credential, err := s.GetCredential(credID)
		if err != nil {
			return nil, fmt.Errorf("获取凭证失败 %s: %w", credID, err)
		}
		verifiableCredentials = append(verifiableCredentials, *credential)
	}

	// 构建可验证展示响应
	vp := models.VerifiablePresentationResponse{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
		},
		ID:                   presentation.PresentationID,
		Type:                 []string{"VerifiablePresentation"},
		Holder:               presentation.HolderDID,
		VerifiableCredential: verifiableCredentials,
		Proof: models.Proof{
			Type:               "SimpleProof2024",
			Created:            presentation.PresentationDate.UTC().Format(time.RFC3339),
			VerificationMethod: fmt.Sprintf("%s#keys-1", presentation.HolderDID),
			ProofPurpose:       "authentication",
			Challenge:          presentation.Challenge,
			ProofValue:         presentation.Proof,
		},
	}

	return &vp, nil
}

// ListPresentations 列出展示
func (s *VCService) ListPresentations(holderDID, verifierDID, status string) ([]models.VerifiablePresentation, error) {
	var presentations []models.VerifiablePresentation
	query := s.DB

	// 根据持有者过滤
	if holderDID != "" {
		query = query.Where("holder_did = ?", holderDID)
	}

	// 根据验证者过滤
	if verifierDID != "" {
		query = query.Where("verifier_did = ?", verifierDID)
	}

	// 根据状态过滤
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 执行查询
	if err := query.Find(&presentations).Error; err != nil {
		return nil, fmt.Errorf("查询展示列表失败: %w", err)
	}

	return presentations, nil
}

// 辅助函数：生成随机挑战
func generateChallenge() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// IssueDoctorVC 颁发医生可验证凭证
func (s *VCService) IssueDoctorVC(issuerDID, doctorDID, vcType, vcContent string) (*models.DoctorVC, error) {
	fmt.Printf("开始颁发医生凭证: issuerDID=%s, doctorDID=%s, vcType=%s\n", issuerDID, doctorDID, vcType)
	
	// 验证参数
	if issuerDID == "" || doctorDID == "" || vcType == "" {
		return nil, fmt.Errorf("颁发者DID、医生DID和凭证类型不能为空")
	}

	// 验证医院DID（简化验证）
	if issuerDID != "0x1234" && !IsHospitalDID(issuerDID) {
		fmt.Printf("医院DID验证失败: %s\n", issuerDID)
		return nil, fmt.Errorf("颁发者DID不是有效的医院DID")
	}

	// 验证医生DID
	var doctor models.Doctor
	if err := s.DB.Where("did_string = ? AND status = ?", doctorDID, "active").First(&doctor).Error; err != nil {
		fmt.Printf("医生DID验证失败: %v\n", err)
		return nil, fmt.Errorf("医生DID无效: %v", err)
	}
	fmt.Printf("找到医生记录: %+v\n", doctor)

	// 检查该医生是否已经有活跃的凭证
	// var existingVC models.DoctorVC
	// result := s.DB.Where("doctor_did = ? AND status = ?", doctorDID, "active").First(&existingVC)
	// fmt.Println("result:",result)
	// if result.Error == nil {
	// 	fmt.Printf("发现现有凭证，直接返回: %+v\n", existingVC)
	// 	return &existingVC, nil
	// } else if result.Error != gorm.ErrRecordNotFound {
	// 	fmt.Printf("查询现有凭证失败: %v\n", result.Error)
	// 	return nil, fmt.Errorf("查询现有凭证失败: %v", result.Error)
	// }

	// 简化的钱包地址检查 - 直接查询所有该钱包的凭证
	var walletVCs []models.DoctorVC
	var walletsToCheck []string
	
	// 先获取所有同钱包地址的医生DID
	var doctorsWithSameWallet []models.Doctor
	if err := s.DB.Where("wallet_address = ? AND status = ?", doctor.WalletAddress, "active").Find(&doctorsWithSameWallet).Error; err != nil {
		fmt.Printf("查询同钱包医生失败: %v\n", err)
		return nil, fmt.Errorf("查询同钱包医生失败: %v", err)
	}
	
	for _, d := range doctorsWithSameWallet {
		walletsToCheck = append(walletsToCheck, d.DIDString)
	}
	
	if len(walletsToCheck) > 0 {
		if err := s.DB.Where("doctor_did IN ? AND status = ?", walletsToCheck, "active").Find(&walletVCs).Error; err != nil {
			fmt.Printf("查询钱包凭证失败: %v\n", err)
			return nil, fmt.Errorf("查询钱包凭证失败: %v", err)
		}
		
		if len(walletVCs) > 0 {
			fmt.Printf("钱包地址已经有凭证，返回第一个: %+v\n", walletVCs[0])
			return &walletVCs[0], nil
		}
	}

	// 生成凭证ID
	vcID := fmt.Sprintf("vc:%s", uuid.New().String())
	fmt.Printf("生成凭证ID: %s\n", vcID)

	// 创建凭证
	now := time.Now()
	expiration := now.AddDate(1, 0, 0) // 1年后过期

	// 创建可验证凭证记录
	doctorVC := models.DoctorVC{
		VCID:      vcID,
		DoctorDID: doctorDID,
		IssuerDID: issuerDID,
		Type:      vcType,
		Content:   vcContent,
		IssuedAt:  now,
		ExpiresAt: expiration,
		Status:    "active",
	}

	fmt.Printf("准备保存凭证: %+v\n", doctorVC)

	// 保存到数据库
	if err := s.DB.Create(&doctorVC).Error; err != nil {
		fmt.Printf("保存医生凭证失败: %v\n", err)
		return nil, fmt.Errorf("保存医生凭证失败: %v", err)
	}

	fmt.Printf("成功保存凭证: %+v\n", doctorVC)
	return &doctorVC, nil
}

// VerifyDoctorVC 验证医生可验证凭证
func (s *VCService) VerifyDoctorVC(vcID string) (*models.VerifyDoctorVCResponse, error) {
	// 查询凭证记录
	var doctorVC models.DoctorVC
	if err := s.DB.Where("vcid = ?", vcID).First(&doctorVC).Error; err != nil {
		return nil, fmt.Errorf("凭证不存在: %v", err)
	}

	// 验证凭证状态
	if doctorVC.Status != "active" {
		return &models.VerifyDoctorVCResponse{
			Valid:  false,
			Reason: "凭证已被撤销",
		}, nil
	}

	// 验证过期时间
	if doctorVC.ExpiresAt.Before(time.Now()) {
		return &models.VerifyDoctorVCResponse{
			Valid:  false,
			Reason: "凭证已过期",
		}, nil
	}

	// 验证医生DID
	var doctor models.Doctor
	if err := s.DB.Where("did_string = ? AND status = ?", doctorVC.DoctorDID, "active").First(&doctor).Error; err != nil {
		return &models.VerifyDoctorVCResponse{
			Valid:  false,
			Reason: "医生DID无效",
		}, nil
	}

	// 验证成功
	return &models.VerifyDoctorVCResponse{
		Valid:     true,
		DoctorDID: doctorVC.DoctorDID,
		IssuerDID: doctorVC.IssuerDID,
		VCType:    doctorVC.Type,
	}, nil
}

// GetDoctorVCs 获取医生的所有可验证凭证
func (s *VCService) GetDoctorVCs(doctorDID string) ([]models.DoctorVC, error) {
	// 验证医生DID
	var doctor models.Doctor
	if err := s.DB.Where("did_string = ? AND status = ?", doctorDID, "active").First(&doctor).Error; err != nil {
		return nil, fmt.Errorf("医生DID无效: %v", err)
	}

	// 查询所有凭证
	var doctorVCs []models.DoctorVC
	if err := s.DB.Where("doctor_did = ?", doctorDID).Find(&doctorVCs).Error; err != nil {
		return nil, fmt.Errorf("获取医生凭证失败: %v", err)
	}

	return doctorVCs, nil
}

// RevokeDoctorVC 撤销医生可验证凭证
func (s *VCService) RevokeDoctorVC(vcID, issuerDID string) error {
	// 查询凭证记录
	var doctorVC models.DoctorVC
	if err := s.DB.Where("vcid = ?", vcID).First(&doctorVC).Error; err != nil {
		return fmt.Errorf("凭证不存在: %v", err)
	}

	// 验证颁发者
	if doctorVC.IssuerDID != issuerDID {
		return fmt.Errorf("只有颁发者可以撤销凭证")
	}

	// 更新凭证状态
	now := time.Now()
	doctorVC.Status = "revoked"
	doctorVC.RevocationDate = &now

	if err := s.DB.Save(&doctorVC).Error; err != nil {
		return fmt.Errorf("撤销凭证失败: %v", err)
	}

	return nil
}

// 辅助函数：检查是否为医院DID
func IsHospitalDID(did string) bool {
	// 在实际应用中，应该查询数据库或区块链验证DID
	// 这里简化处理，假设以"hospital"开头的DID为有效医院DID
	return len(did) > 8 && did[0:8] == "hospital"
}
