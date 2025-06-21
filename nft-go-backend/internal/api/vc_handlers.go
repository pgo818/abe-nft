package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ABE/nft/nft-go-backend/internal/models"
	"github.com/ABE/nft/nft-go-backend/internal/service"
)

// VCHandlers 可验证凭证相关处理程序结构体
type VCHandlers struct {
	Service    *service.VCService
	DIDService *service.DIDService
}

// NewVCHandlers 创建新的VC处理程序
func NewVCHandlers(vcService *service.VCService, didService *service.DIDService) *VCHandlers {
	return &VCHandlers{
		Service:    vcService,
		DIDService: didService,
	}
}

// IssueCredentialHandler 颁发凭证处理程序
func (h *VCHandlers) IssueCredentialHandler(c *gin.Context) {
	var req models.IssueCredentialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务颁发凭证
	credential, err := h.Service.IssueCredential(req.IssuerDID, req.SubjectDID, req.CredentialType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "颁发凭证失败: " + err.Error()})
		return
	}

	// 构建响应
	response := models.IssueCredentialResponse{
		Credential: models.VerifiableCredentialResponse{
			ID:       credential.CredentialID,
			Issuer:   credential.IssuerDID,
			Subject:  credential.SubjectDID,
			Type:     []string{"VerifiableCredential", credential.Type},
			IssuedAt: credential.IssuanceDate.Format("2006-01-02T15:04:05Z"),
		},
	}

	c.JSON(http.StatusOK, response)
}

// VerifyCredentialHandler 验证凭证处理程序
func (h *VCHandlers) VerifyCredentialHandler(c *gin.Context) {
	var req models.VerifyCredentialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务验证凭证
	result, err := h.Service.VerifyCredential(req.CredentialID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "验证凭证失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// RevokeCredentialHandler 撤销凭证处理程序
func (h *VCHandlers) RevokeCredentialHandler(c *gin.Context) {
	var req models.RevokeCredentialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务撤销凭证
	err := h.Service.RevokeCredential(req.CredentialID, req.IssuerDID, req.Reason)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "撤销凭证失败: " + err.Error()})
		return
	}

	// 构建响应
	response := models.RevokeCredentialResponse{
		CredentialID: req.CredentialID,
		Status:       "revoked",
		RevokedAt:    "now",
		Reason:       req.Reason,
	}

	c.JSON(http.StatusOK, response)
}

// GetCredentialHandler 获取凭证处理程序
func (h *VCHandlers) GetCredentialHandler(c *gin.Context) {
	credentialID := c.Param("id")
	if credentialID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "凭证ID参数不能为空"})
		return
	}

	credential, err := h.Service.GetCredential(credentialID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "获取凭证失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, credential)
}

// ListCredentialsHandler 列出凭证处理程序
func (h *VCHandlers) ListCredentialsHandler(c *gin.Context) {
	issuerDID := c.Query("issuer")
	subjectDID := c.Query("subject")
	status := c.Query("status")

	credentials, err := h.Service.ListCredentials(issuerDID, subjectDID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取凭证列表失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"credentials": credentials})
}

// CreatePresentationHandler 创建可验证表示处理程序
func (h *VCHandlers) CreatePresentationHandler(c *gin.Context) {
	var req models.CreatePresentationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务创建表示
	presentation, err := h.Service.CreatePresentation(req.HolderDID, req.VerifierDID, req.CredentialIDs, req.Purpose)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建表示失败: " + err.Error()})
		return
	}

	// 构建响应
	response := models.CreatePresentationResponse{
		Presentation: models.VerifiablePresentationResponse{
			Context: []string{"https://www.w3.org/2018/credentials/v1"},
			ID:      presentation.PresentationID,
			Type:    []string{"VerifiablePresentation"},
			Holder:  presentation.HolderDID,
		},
	}

	c.JSON(http.StatusOK, response)
}

// VerifyPresentationHandler 验证可验证表示处理程序
func (h *VCHandlers) VerifyPresentationHandler(c *gin.Context) {
	var req models.VerifyPresentationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务验证表示
	result, err := h.Service.VerifyPresentation(req.PresentationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "验证表示失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetPresentationHandler 获取表示处理程序
func (h *VCHandlers) GetPresentationHandler(c *gin.Context) {
	presentationID := c.Param("id")
	if presentationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "表示ID参数不能为空"})
		return
	}

	presentation, err := h.Service.GetPresentation(presentationID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "获取表示失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, presentation)
}

// ListPresentationsHandler 列出表示处理程序
func (h *VCHandlers) ListPresentationsHandler(c *gin.Context) {
	holderDID := c.Query("holder")
	verifierDID := c.Query("verifier")
	status := c.Query("status")

	presentations, err := h.Service.ListPresentations(holderDID, verifierDID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取表示列表失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"presentations": presentations})
}

// CreateDoctorDIDHandler 创建医生DID处理程序
func (h *VCHandlers) CreateDoctorDIDHandler(c *gin.Context) {
	var req models.CreateDoctorDIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务创建医生DID
	doctor, err := h.DIDService.CreateDoctorDID(req.WalletAddress, req.Name, req.LicenseNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建医生DID失败: " + err.Error()})
		return
	}

	// 构建响应
	response := models.CreateDoctorDIDResponse{
		DID:           doctor.DIDString,
		WalletAddress: doctor.WalletAddress,
		Name:          doctor.Name,
		LicenseNumber: doctor.LicenseNumber,
	}

	c.JSON(http.StatusOK, response)
}

// IssueDoctorVCHandler 颁发医生凭证处理程序
func (h *VCHandlers) IssueDoctorVCHandler(c *gin.Context) {
	var req models.IssueDoctorVCRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务颁发医生凭证
	vc, err := h.Service.IssueDoctorVC(req.IssuerDID, req.DoctorDID, req.VCType, req.VCContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "颁发医生凭证失败: " + err.Error()})
		return
	}

	// 构建响应
	response := models.IssueDoctorVCResponse{
		VCID:      vc.VCID,
		DoctorDID: vc.DoctorDID,
		IssuerDID: vc.IssuerDID,
		VCType:    vc.Type,
		IssuedAt:  vc.IssuedAt,
		ExpiresAt: vc.ExpiresAt,
	}

	c.JSON(http.StatusOK, response)
}

// VerifyDoctorVCHandler 验证医生凭证处理程序
func (h *VCHandlers) VerifyDoctorVCHandler(c *gin.Context) {
	var req models.VerifyDoctorVCRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务验证医生凭证
	result, err := h.Service.VerifyDoctorVC(req.VCID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "验证医生凭证失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetDoctorVCsHandler 获取医生凭证列表处理程序
func (h *VCHandlers) GetDoctorVCsHandler(c *gin.Context) {
	var req models.GetDoctorVCsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求体: " + err.Error()})
		return
	}

	// 调用服务获取医生凭证列表
	vcs, err := h.Service.GetDoctorVCs(req.DoctorDID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取医生凭证列表失败: " + err.Error()})
		return
	}

	// 构建响应
	response := models.GetDoctorVCsResponse{
		DoctorDID: req.DoctorDID,
		VCs:       vcs,
	}

	c.JSON(http.StatusOK, response)
}
