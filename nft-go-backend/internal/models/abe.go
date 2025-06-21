package models

// SetupRequest 初始化ABE系统的请求
type SetupRequest struct {
	Gamma []string `json:"gamma" binding:"required"` // 属性列表
}

// SetupResponse 初始化ABE系统的响应
type SetupResponse struct {
	PubKey string `json:"pub_key"` // 序列化后的公钥
	SecKey string `json:"sec_key"` // 序列化后的主密钥
}

// KeyGenRequest 生成属性密钥的请求
type KeyGenRequest struct {
	Gamma  []string `json:"gamma" binding:"required"`   // 属性列表
	PubKey string   `json:"pub_key" binding:"required"` // 序列化后的公钥
	SecKey string   `json:"sec_key" binding:"required"` // 序列化后的主密钥
}

// KeyGenResponse 生成属性密钥的响应
type KeyGenResponse struct {
	AttribKeys string `json:"attrib_keys"` // 序列化后的属性密钥
}

// EncryptRequest 加密消息的请求
type EncryptRequest struct {
	Message string `json:"message" binding:"required"` // 要加密的消息
	Policy  string `json:"policy" binding:"required"`  // 访问策略（布尔表达式）
	PubKey  string `json:"pub_key" binding:"required"` // 序列化后的公钥
}

// EncryptResponse 加密消息的响应
type EncryptResponse struct {
	Cipher string `json:"cipher"` // 序列化后的密文
}

// DecryptRequest 解密消息的请求
type DecryptRequest struct {
	Cipher     string `json:"cipher" binding:"required"`      // 序列化后的密文
	AttribKeys string `json:"attrib_keys" binding:"required"` // 序列化后的属性密钥
	PubKey     string `json:"pub_key" binding:"required"`     // 序列化后的公钥
}

// DecryptResponse 解密消息的响应
type DecryptResponse struct {
	Message string `json:"message"` // 解密后的消息
}

// ErrorResponse API错误响应
type ErrorResponse struct {
	Error string `json:"error"`
}


