package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/fentec-project/bn256"
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

// FAMEPubKeyData 用于JSON序列化的公钥数据结构
type FAMEPubKeyData struct {
	PartG2 [][2][]byte `json:"part_g2"`
	PartGT [][2][]byte `json:"part_gt"`
}

// FAMESecKeyData 用于JSON序列化的私钥数据结构
type FAMESecKeyData struct {
	PartInt [][]byte    `json:"part_int"`
	PartG1  [][3][]byte `json:"part_g1"`
}

// FAMEAttribKeysData 用于JSON序列化的属性密钥数据结构
type FAMEAttribKeysData struct {
	K0        [][3][]byte    `json:"k0"`
	K         [][][3][]byte  `json:"k"`
	KPrime    [][3][]byte    `json:"k_prime"`
	AttribToI map[string]int `json:"attrib_to_i"`
}

// FAMECipherData 用于JSON序列化的密文数据结构
type FAMECipherData struct {
	Ct0     [][3][]byte   `json:"ct0"`
	Ct      [][][3][]byte `json:"ct"`
	CtPrime []byte        `json:"ct_prime"`
	MspData []byte        `json:"msp_data"`
	SymEnc  []byte        `json:"sym_enc"`
	Iv      []byte        `json:"iv"`
}

// serializeFAMEPubKey 序列化FAME公钥
func serializeFAMEPubKey(pubKey *abe.FAMEPubKey) (string, error) {
	// 序列化PartG2数组
	partG2Data := make([][2][]byte, 2)
	for i := 0; i < 2; i++ {
		if pubKey.PartG2[i] != nil {
			partG2Data[i] = [2][]byte{pubKey.PartG2[i].Marshal()}
		} else {
			partG2Data[i] = [2][]byte{[]byte{}}
		}
	}

	// 序列化PartGT数组
	partGTData := make([][2][]byte, 2)
	for i := 0; i < 2; i++ {
		if pubKey.PartGT[i] != nil {
			partGTData[i] = [2][]byte{pubKey.PartGT[i].Marshal()}
		} else {
			partGTData[i] = [2][]byte{[]byte{}}
		}
	}

	data := FAMEPubKeyData{
		PartG2: partG2Data,
		PartGT: partGTData,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(jsonData), nil
}

// serializeFAMESecKey 序列化FAME私钥
func serializeFAMESecKey(secKey *abe.FAMESecKey) (string, error) {
	// 序列化PartInt数组
	partIntData := make([][]byte, 4)
	for i := 0; i < 4; i++ {
		if secKey.PartInt[i] != nil {
			partIntData[i] = []byte(secKey.PartInt[i].String())
		} else {
			partIntData[i] = []byte{}
		}
	}

	// 序列化PartG1数组
	partG1Data := make([][3][]byte, 3)
	for i := 0; i < 3; i++ {
		if secKey.PartG1[i] != nil {
			partG1Data[i] = [3][]byte{secKey.PartG1[i].Marshal()}
		} else {
			partG1Data[i] = [3][]byte{[]byte{}}
		}
	}

	data := FAMESecKeyData{
		PartInt: partIntData,
		PartG1:  partG1Data,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(jsonData), nil
}

// serializeFAMEAttribKeys 序列化FAME属性密钥
func serializeFAMEAttribKeys(attribKeys *abe.FAMEAttribKeys) (string, error) {
	// 序列化K0数组
	k0Data := make([][3][]byte, 3)
	for i := 0; i < 3; i++ {
		if attribKeys.K0[i] != nil {
			k0Data[i] = [3][]byte{attribKeys.K0[i].Marshal()}
		} else {
			k0Data[i] = [3][]byte{[]byte{}}
		}
	}

	// 序列化K数组
	kData := make([][][3][]byte, len(attribKeys.K))
	for i := 0; i < len(attribKeys.K); i++ {
		kData[i] = make([][3][]byte, 3)
		for j := 0; j < 3; j++ {
			if attribKeys.K[i][j] != nil {
				kData[i][j] = [3][]byte{attribKeys.K[i][j].Marshal()}
			} else {
				kData[i][j] = [3][]byte{[]byte{}}
			}
		}
	}

	// 序列化KPrime数组
	kPrimeData := make([][3][]byte, 3)
	for i := 0; i < 3; i++ {
		if attribKeys.KPrime[i] != nil {
			kPrimeData[i] = [3][]byte{attribKeys.KPrime[i].Marshal()}
		} else {
			kPrimeData[i] = [3][]byte{[]byte{}}
		}
	}

	data := FAMEAttribKeysData{
		K0:        k0Data,
		K:         kData,
		KPrime:    kPrimeData,
		AttribToI: attribKeys.AttribToI,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(jsonData), nil
}

// serializeFAMECipher 序列化FAME密文
func serializeFAMECipher(cipher *abe.FAMECipher) (string, error) {
	// 序列化Ct0数组
	ct0Data := make([][3][]byte, 3)
	for i := 0; i < 3; i++ {
		if i < len(cipher.Ct0) && cipher.Ct0[i] != nil {
			ct0Data[i] = [3][]byte{cipher.Ct0[i].Marshal()}
		} else {
			ct0Data[i] = [3][]byte{[]byte{}}
		}
	}

	// 序列化Ct数组
	ctData := make([][][3][]byte, len(cipher.Ct))
	for i := 0; i < len(cipher.Ct); i++ {
		ctData[i] = make([][3][]byte, 3)
		for j := 0; j < 3; j++ {
			if i < len(cipher.Ct) && j < len(cipher.Ct[i]) && cipher.Ct[i][j] != nil {
				ctData[i][j] = [3][]byte{cipher.Ct[i][j].Marshal()}
			} else {
				ctData[i][j] = [3][]byte{[]byte{}}
			}
		}
	}

	// 序列化CtPrime
	var ctPrimeData []byte
	if cipher.CtPrime != nil {
		ctPrimeData = cipher.CtPrime.Marshal()
	} else {
		ctPrimeData = []byte{}
	}

	// 序列化MSP
	var mspData []byte
	var err error
	if cipher.Msp != nil {
		mspData, err = json.Marshal(cipher.Msp)
		if err != nil {
			return "", fmt.Errorf("序列化MSP失败: %v", err)
		}
	} else {
		mspData = []byte("{}")
	}

	data := FAMECipherData{
		Ct0:     ct0Data,
		Ct:      ctData,
		CtPrime: ctPrimeData,
		MspData: mspData,
		SymEnc:  cipher.SymEnc,
		Iv:      cipher.Iv,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(jsonData), nil
}

// deserializeFAMEPubKey 反序列化FAME公钥
func deserializeFAMEPubKey(data string) (*abe.FAMEPubKey, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	var keyData FAMEPubKeyData
	if err := json.Unmarshal(decoded, &keyData); err != nil {
		return nil, err
	}

	pubKey := &abe.FAMEPubKey{}

	// 初始化PartG2数组
	pubKey.PartG2 = [2]*bn256.G2{new(bn256.G2), new(bn256.G2)}

	// 反序列化PartG2
	for i := 0; i < 2; i++ {
		if i < len(keyData.PartG2) && len(keyData.PartG2[i]) > 0 && len(keyData.PartG2[i][0]) > 0 {
			if _, err := pubKey.PartG2[i].Unmarshal(keyData.PartG2[i][0]); err != nil {
				return nil, fmt.Errorf("反序列化PartG2[%d]失败: %v", i, err)
			}
		}
	}

	// 初始化PartGT数组
	pubKey.PartGT = [2]*bn256.GT{new(bn256.GT), new(bn256.GT)}

	// 反序列化PartGT
	for i := 0; i < 2; i++ {
		if i < len(keyData.PartGT) && len(keyData.PartGT[i]) > 0 && len(keyData.PartGT[i][0]) > 0 {
			if _, err := pubKey.PartGT[i].Unmarshal(keyData.PartGT[i][0]); err != nil {
				return nil, fmt.Errorf("反序列化PartGT[%d]失败: %v", i, err)
			}
		}
	}

	return pubKey, nil
}

// deserializeFAMESecKey 反序列化FAME私钥
func deserializeFAMESecKey(data string) (*abe.FAMESecKey, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	var keyData FAMESecKeyData
	if err := json.Unmarshal(decoded, &keyData); err != nil {
		return nil, err
	}

	secKey := &abe.FAMESecKey{}

	// 初始化PartInt数组
	secKey.PartInt = [4]*big.Int{new(big.Int), new(big.Int), new(big.Int), new(big.Int)}

	// 反序列化PartInt
	for i := 0; i < 4; i++ {
		if i < len(keyData.PartInt) && len(keyData.PartInt[i]) > 0 {
			var ok bool
			secKey.PartInt[i], ok = secKey.PartInt[i].SetString(string(keyData.PartInt[i]), 10)
			if !ok {
				return nil, fmt.Errorf("反序列化PartInt[%d]失败", i)
			}
		}
	}

	// 初始化PartG1数组
	secKey.PartG1 = [3]*bn256.G1{new(bn256.G1), new(bn256.G1), new(bn256.G1)}

	// 反序列化PartG1
	for i := 0; i < 3; i++ {
		if i < len(keyData.PartG1) && len(keyData.PartG1[i]) > 0 && len(keyData.PartG1[i][0]) > 0 {
			if _, err := secKey.PartG1[i].Unmarshal(keyData.PartG1[i][0]); err != nil {
				return nil, fmt.Errorf("反序列化PartG1[%d]失败: %v", i, err)
			}
		}
	}

	return secKey, nil
}

// deserializeFAMEAttribKeys 反序列化FAME属性密钥
func deserializeFAMEAttribKeys(data string) (*abe.FAMEAttribKeys, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	var keyData FAMEAttribKeysData
	if err := json.Unmarshal(decoded, &keyData); err != nil {
		return nil, err
	}

	attribKeys := &abe.FAMEAttribKeys{}

	// 初始化K0数组
	attribKeys.K0 = [3]*bn256.G2{new(bn256.G2), new(bn256.G2), new(bn256.G2)}

	// 反序列化K0
	for i := 0; i < 3; i++ {
		if i < len(keyData.K0) && len(keyData.K0[i]) > 0 && len(keyData.K0[i][0]) > 0 {
			if _, err := attribKeys.K0[i].Unmarshal(keyData.K0[i][0]); err != nil {
				return nil, fmt.Errorf("反序列化K0[%d]失败: %v", i, err)
			}
		}
	}

	// 反序列化K
	attribKeys.K = make([][3]*bn256.G1, len(keyData.K))
	for i := range keyData.K {
		attribKeys.K[i] = [3]*bn256.G1{new(bn256.G1), new(bn256.G1), new(bn256.G1)}
		for j := 0; j < 3; j++ {
			if j < len(keyData.K[i]) && len(keyData.K[i][j]) > 0 && len(keyData.K[i][j][0]) > 0 {
				if _, err := attribKeys.K[i][j].Unmarshal(keyData.K[i][j][0]); err != nil {
					return nil, fmt.Errorf("反序列化K[%d][%d]失败: %v", i, j, err)
				}
			}
		}
	}

	// 初始化KPrime数组
	attribKeys.KPrime = [3]*bn256.G1{new(bn256.G1), new(bn256.G1), new(bn256.G1)}

	// 反序列化KPrime
	for i := 0; i < 3; i++ {
		if i < len(keyData.KPrime) && len(keyData.KPrime[i]) > 0 && len(keyData.KPrime[i][0]) > 0 {
			if _, err := attribKeys.KPrime[i].Unmarshal(keyData.KPrime[i][0]); err != nil {
				return nil, fmt.Errorf("反序列化KPrime[%d]失败: %v", i, err)
			}
		}
	}

	// 复制属性映射
	attribKeys.AttribToI = keyData.AttribToI

	return attribKeys, nil
}

// deserializeFAMECipher 反序列化FAME密文
func deserializeFAMECipher(data string) (*abe.FAMECipher, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	var cipherData FAMECipherData
	if err := json.Unmarshal(decoded, &cipherData); err != nil {
		return nil, err
	}

	cipher := &abe.FAMECipher{}

	// 初始化Ct0数组
	cipher.Ct0 = [3]*bn256.G2{new(bn256.G2), new(bn256.G2), new(bn256.G2)}

	// 反序列化Ct0
	for i := 0; i < 3; i++ {
		if i < len(cipherData.Ct0) && len(cipherData.Ct0[i]) > 0 && len(cipherData.Ct0[i][0]) > 0 {
			if _, err := cipher.Ct0[i].Unmarshal(cipherData.Ct0[i][0]); err != nil {
				return nil, fmt.Errorf("反序列化Ct0[%d]失败: %v", i, err)
			}
		}
	}

	// 反序列化Ct
	cipher.Ct = make([][3]*bn256.G1, len(cipherData.Ct))
	for i := range cipherData.Ct {
		cipher.Ct[i] = [3]*bn256.G1{new(bn256.G1), new(bn256.G1), new(bn256.G1)}
		for j := 0; j < 3; j++ {
			if j < len(cipherData.Ct[i]) && len(cipherData.Ct[i][j]) > 0 && len(cipherData.Ct[i][j][0]) > 0 {
				if _, err := cipher.Ct[i][j].Unmarshal(cipherData.Ct[i][j][0]); err != nil {
					return nil, fmt.Errorf("反序列化Ct[%d][%d]失败: %v", i, j, err)
				}
			}
		}
	}

	// 反序列化CtPrime
	cipher.CtPrime = new(bn256.GT)
	if len(cipherData.CtPrime) > 0 {
		if _, err := cipher.CtPrime.Unmarshal(cipherData.CtPrime); err != nil {
			return nil, fmt.Errorf("反序列化CtPrime失败: %v", err)
		}
	}

	// 反序列化MSP
	if len(cipherData.MspData) > 0 {
		var msp abe.MSP
		if err := json.Unmarshal(cipherData.MspData, &msp); err != nil {
			return nil, fmt.Errorf("反序列化MSP失败: %v", err)
		}
		cipher.Msp = &msp
	} else {
		cipher.Msp = &abe.MSP{}
	}

	// 复制对称加密数据
	cipher.SymEnc = cipherData.SymEnc
	cipher.Iv = cipherData.Iv

	return cipher, nil
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

	// 使用自定义序列化方法
	pubKeyStr, err := serializeFAMEPubKey(pubKey)
	if err != nil {
		return nil, fmt.Errorf("序列化公钥失败: %v", err)
	}

	secKeyStr, err := serializeFAMESecKey(secKey)
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
		PubKey:     pubKeyStr,
		SecKey:     secKeyStr,
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

	// 反序列化主密钥 (密钥生成只需要主密钥，不需要公钥)
	secKey, err := deserializeFAMESecKey(systemKey.SecKey)
	if err != nil {
		return nil, fmt.Errorf("反序列化私钥失败: %v", err)
	}

	// 创建FAME密码系统
	fameScheme := abe.NewFAME()

	// 生成用户密钥 - FAME算法的密钥生成只需要主密钥和用户属性
	// 注意：在FAME密钥生成阶段不需要使用公钥
	// 公钥主要用于加密和解密阶段
	attribKeys, err := fameScheme.GenerateAttribKeys(userAttributes, secKey)
	if err != nil {
		return nil, fmt.Errorf("生成用户密钥失败: %v", err)
	}

	// 序列化用户密钥
	attribKeysStr, err := serializeFAMEAttribKeys(attribKeys)
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
		AttribKeys:  attribKeysStr,
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
	pubKey, err := deserializeFAMEPubKey(systemKey.PubKey)
	if err != nil {
		return nil, fmt.Errorf("反序列化公钥失败: %v", err)
	}

	// 检查公钥是否正确初始化
	if pubKey == nil ||
		pubKey.PartG2[0] == nil || pubKey.PartG2[1] == nil ||
		pubKey.PartGT[0] == nil || pubKey.PartGT[1] == nil {
		// 如果公钥未正确初始化，重新生成系统密钥
		fameScheme := abe.NewFAME()
		newPubKey, newSecKey, err := fameScheme.GenerateMasterKeys()
		if err != nil {
			return nil, fmt.Errorf("生成新密钥失败: %v", err)
		}

		// 使用新生成的密钥
		pubKey = newPubKey

		// 序列化并更新数据库中的密钥
		pubKeyStr, err := serializeFAMEPubKey(newPubKey)
		if err != nil {
			return nil, fmt.Errorf("序列化新公钥失败: %v", err)
		}

		secKeyStr, err := serializeFAMESecKey(newSecKey)
		if err != nil {
			return nil, fmt.Errorf("序列化新私钥失败: %v", err)
		}

		// 更新数据库中的系统密钥
		systemKey.PubKey = pubKeyStr
		systemKey.SecKey = secKeyStr
		if err := s.DB.Save(&systemKey).Error; err != nil {
			return nil, fmt.Errorf("更新系统密钥失败: %v", err)
		}
	}
	fmt.Println("反序列化公钥成功")
	// 创建FAME密码系统
	fameScheme := abe.NewFAME()

	// 将策略字符串转换为MSP结构
	msp, err := abe.BooleanToMSP(policy, false)
	if err != nil {
		return nil, fmt.Errorf("转换策略失败: %v", err)
	}
	fmt.Println("message:", message)

	// 加密消息
	cipher, err := fameScheme.Encrypt(message, msp, pubKey)
	if err != nil {
		return nil, fmt.Errorf("加密失败: %v", err)
	}

	// 序列化密文
	cipherStr, err := serializeFAMECipher(cipher)
	if err != nil {
		return nil, fmt.Errorf("序列化密文失败: %v", err)
	}

	// 创建密文记录
	ciphertext := models.ABECiphertext{
		Cipher:      cipherStr,
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
	pubKey, err := deserializeFAMEPubKey(systemKey.PubKey)
	if err != nil {
		return "", fmt.Errorf("反序列化公钥失败: %v", err)
	}

	// 检查公钥是否正确初始化
	if pubKey == nil ||
		pubKey.PartG2[0] == nil || pubKey.PartG2[1] == nil ||
		pubKey.PartGT[0] == nil || pubKey.PartGT[1] == nil {
		return "", fmt.Errorf("公钥未正确初始化，请重新生成系统密钥")
	}

	// 反序列化用户密钥
	attribKeys, err := deserializeFAMEAttribKeys(userKey.AttribKeys)
	if err != nil {
		return "", fmt.Errorf("反序列化用户密钥失败: %v", err)
	}

	// 检查用户密钥是否正确初始化
	if attribKeys == nil ||
		attribKeys.K0[0] == nil || attribKeys.K0[1] == nil || attribKeys.K0[2] == nil ||
		attribKeys.KPrime[0] == nil || attribKeys.KPrime[1] == nil || attribKeys.KPrime[2] == nil {
		return "", fmt.Errorf("用户密钥未正确初始化，请重新生成用户密钥")
	}

	// 反序列化密文
	cipher, err := deserializeFAMECipher(ciphertext.Cipher)
	if err != nil {
		return "", fmt.Errorf("反序列化密文失败: %v", err)
	}

	// 检查密文是否正确初始化
	if cipher == nil ||
		cipher.Ct0[0] == nil || cipher.Ct0[1] == nil || cipher.Ct0[2] == nil ||
		cipher.CtPrime == nil || cipher.Msp == nil {
		return "", fmt.Errorf("密文未正确初始化，请重新加密")
	}

	// 创建FAME密码系统
	fameScheme := abe.NewFAME()

	// 解密消息
	message, err := fameScheme.Decrypt(cipher, attribKeys, pubKey)
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

// DecryptABEDirect 直接解密数据（不依赖数据库记录）
func (s *ABEService) DecryptABEDirect(cipherStr string, attribKeysStr string) (string, error) {
	// 自动获取或创建系统密钥
	systemKey, err := s.GetOrCreateSystemKey()
	if err != nil {
		return "", fmt.Errorf("获取系统密钥失败: %v", err)
	}

	// 反序列化公钥
	pubKey, err := deserializeFAMEPubKey(systemKey.PubKey)
	if err != nil {
		return "", fmt.Errorf("反序列化公钥失败: %v", err)
	}

	// 检查公钥是否正确初始化
	if pubKey == nil ||
		pubKey.PartG2[0] == nil || pubKey.PartG2[1] == nil ||
		pubKey.PartGT[0] == nil || pubKey.PartGT[1] == nil {
		return "", fmt.Errorf("公钥未正确初始化，请重新生成系统密钥")
	}

	// 反序列化用户密钥
	attribKeys, err := deserializeFAMEAttribKeys(attribKeysStr)
	if err != nil {
		return "", fmt.Errorf("反序列化用户密钥失败: %v", err)
	}

	// 检查用户密钥是否正确初始化
	if attribKeys == nil ||
		attribKeys.K0[0] == nil || attribKeys.K0[1] == nil || attribKeys.K0[2] == nil ||
		attribKeys.KPrime[0] == nil || attribKeys.KPrime[1] == nil || attribKeys.KPrime[2] == nil {
		return "", fmt.Errorf("用户密钥未正确初始化，请重新生成用户密钥")
	}

	// 反序列化密文
	cipher, err := deserializeFAMECipher(cipherStr)
	if err != nil {
		return "", fmt.Errorf("反序列化密文失败: %v", err)
	}

	// 检查密文是否正确初始化
	if cipher == nil ||
		cipher.Ct0[0] == nil || cipher.Ct0[1] == nil || cipher.Ct0[2] == nil ||
		cipher.CtPrime == nil || cipher.Msp == nil {
		return "", fmt.Errorf("密文未正确初始化，请重新加密")
	}

	// 创建FAME密码系统
	fameScheme := abe.NewFAME()

	// 解密消息
	message, err := fameScheme.Decrypt(cipher, attribKeys, pubKey)
	if err != nil {
		return "", fmt.Errorf("解密失败: %v", err)
	}

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
	if err := s.DB.Order("created_at DESC").First(&systemKey).Error; err != nil {
		return nil, err
	}
	return &systemKey, nil
}

// GetOrCreateSystemKey 获取或创建系统密钥
func (s *ABEService) GetOrCreateSystemKey() (*models.ABESystemKey, error) {
	// 首先尝试获取现有的系统密钥
	systemKey, err := s.GetLatestSystemKey()
	if err == nil {
		// 验证现有系统密钥的完整性
		pubKey, pubErr := deserializeFAMEPubKey(systemKey.PubKey)
		secKey, secErr := deserializeFAMESecKey(systemKey.SecKey)

		if pubErr != nil || secErr != nil ||
			pubKey == nil || secKey == nil ||
			pubKey.PartG2[0] == nil || pubKey.PartG2[1] == nil ||
			pubKey.PartGT[0] == nil || pubKey.PartGT[1] == nil {
			// 如果现有密钥损坏，删除并重新创建
			fmt.Printf("检测到损坏的系统密钥，正在重新生成...\n")
			s.DB.Delete(&systemKey)
		} else {
			return systemKey, nil
		}
	}

	// 创建新的系统密钥
	fmt.Printf("创建新的ABE系统密钥...\n")
	defaultAttributes := []string{
		"mainNFT:0x651e0fd49C7dbB5cca8b5Be0319d92773443b711",
		"mainNFT:0xAF97631F96007bbde9C7803B3BeA096f4A5a5561",
		"mainNFT:0x8ac134A862BD7279406852ebe9736f23D4eae444",
		"mainNFT:0xE5B4c33E9cb5D7BfcdEA781e24D301924fF1B987",
		"mainNFT:0x47b749D07Ac4Ad1FB366A8B12F6D8BD53EF92831",
		"mainNFT:0x948eD518e2dEF88c8BFB2742fe8dEb6843a7D3D4",
		"mainNFT:0xF73756057652D9158e6A52135151A96e783cB570",
		"mainNFT:0x871343E7F8C63aF0989398f32A5Cc24B1274D748",
		"mainNFT:0xe1814486DFf863e3e15a6d2c2E2Ca58D7dd59dcF",
		"mainNFT:0xfDF080f6D103e896F77d51127d3a73557Ad551dA",
	}

	return s.SetupABE(defaultAttributes, 1) // 使用默认用户ID 1
}

// GenerateUserKeyAuto 自动生成用户密钥
func (s *ABEService) GenerateUserKeyAuto(userAttributes []string) (*models.ABEUserKey, error) {
	// 自动获取或创建系统密钥
	systemKey, err := s.GetOrCreateSystemKey()
	if err != nil {
		return nil, fmt.Errorf("获取系统密钥失败: %v", err)
	}

	// 生成用户密钥
	return s.KeyGenABE(systemKey.ID, 1, userAttributes) // 使用默认用户ID 1
}

// extractAttributesFromMap 从map中提取VC属性的辅助函数
func extractAttributesFromMap(data map[string]interface{}, vcAttributes map[string]string) {
	// 支持的属性列表
	supportedAttributes := []string{
		"wallet", "name", "department", "hospital", "title",
		"licenseNumber", "specialty", "did", "specialties",
	}

	for _, attr := range supportedAttributes {
		if value, ok := data[attr].(string); ok {
			vcAttributes[attr] = value
		}
	}
}

// VerifyVCAgainstPolicy 验证VC凭证是否满足访问策略
func (s *ABEService) VerifyVCAgainstPolicy(vcContent string, policy string) (bool, map[string]interface{}, error) {
	fmt.Printf("开始验证VC凭证策略: VC=%s, Policy=%s\n", vcContent, policy)

	// 解析VC凭证内容
	var vcData map[string]interface{}
	if err := json.Unmarshal([]byte(vcContent), &vcData); err != nil {
		return false, nil, fmt.Errorf("解析VC凭证失败: %v", err)
	}

	// 提取VC中的属性
	vcAttributes := make(map[string]string)

	// 方法1：尝试从credentialSubject中提取（标准VC格式）
	if credentialSubject, ok := vcData["credentialSubject"].(map[string]interface{}); ok {
		fmt.Printf("使用标准VC格式提取属性\n")
		extractAttributesFromMap(credentialSubject, vcAttributes)
	} else {
		// 方法2：直接从根级别提取属性（当前VC格式）
		fmt.Printf("使用直接格式提取属性\n")
		extractAttributesFromMap(vcData, vcAttributes)
	}

	fmt.Printf("提取的VC属性: %+v\n", vcAttributes)

	// 解析策略字符串
	result, failedConditions, err := s.evaluatePolicyWithDetails(policy, vcAttributes)
	if err != nil {
		return false, nil, fmt.Errorf("策略验证失败: %v", err)
	}

	// 构建详细结果
	verificationResult := map[string]interface{}{
		"vcAttributes":     vcAttributes,
		"policy":           policy,
		"satisfied":        result,
		"verificationTime": time.Now().Format(time.RFC3339),
	}

	// 如果验证失败，添加失败详情
	if !result && len(failedConditions) > 0 {
		verificationResult["failedConditions"] = failedConditions
		verificationResult["detailedReason"] = s.buildDetailedFailureReason(failedConditions, vcAttributes)
	}

	fmt.Printf("策略验证结果: %v\n", result)
	return result, verificationResult, nil
}

// evaluatePolicyWithDetails 评估策略并返回失败的具体条件
func (s *ABEService) evaluatePolicyWithDetails(policy string, attributes map[string]string) (bool, []map[string]interface{}, error) {
	var failedConditions []map[string]interface{}

	// 处理简单的策略格式，如 "department:心内科"
	policy = strings.TrimSpace(policy)

	// 如果策略包含逻辑操作符，需要递归解析
	if strings.Contains(policy, " AND ") || strings.Contains(policy, " OR ") {
		return s.evaluateComplexPolicyWithDetails(policy, attributes)
	}

	// 处理单个条件
	result, err := s.evaluateSimpleCondition(policy, attributes)
	if err != nil {
		return false, nil, err
	}

	if !result {
		// 解析条件以获取详细信息
		parts := strings.SplitN(policy, ":", 2)
		if len(parts) == 2 {
			attributeName := strings.TrimSpace(parts[0])
			expectedValue := strings.TrimSpace(strings.Trim(parts[1], "\"'"))
			actualValue, exists := attributes[attributeName]

			failedCondition := map[string]interface{}{
				"condition":     policy,
				"attribute":     attributeName,
				"expectedValue": expectedValue,
				"actualValue":   "",
				"hasAttribute":  exists,
			}

			if exists {
				failedCondition["actualValue"] = actualValue
			}

			failedConditions = append(failedConditions, failedCondition)
		}
	}

	return result, failedConditions, nil
}

// evaluateComplexPolicyWithDetails 评估复杂策略（包含AND/OR）并返回失败详情
func (s *ABEService) evaluateComplexPolicyWithDetails(policy string, attributes map[string]string) (bool, []map[string]interface{}, error) {
	var allFailedConditions []map[string]interface{}

	// 去除外层括号
	policy = strings.TrimSpace(policy)
	if strings.HasPrefix(policy, "(") && strings.HasSuffix(policy, ")") {
		policy = policy[1 : len(policy)-1]
	}

	// 分割AND条件
	andParts := s.splitPolicy(policy, " AND ")

	allSatisfied := true
	for _, andPart := range andParts {
		andPart = strings.TrimSpace(andPart)

		// 检查是否包含OR
		if strings.Contains(andPart, " OR ") {
			// 处理OR条件
			orResult, orFailed, err := s.evaluateOrConditionWithDetails(andPart, attributes)
			if err != nil {
				return false, nil, err
			}
			if !orResult {
				allSatisfied = false
				allFailedConditions = append(allFailedConditions, orFailed...)
			}
		} else {
			// 处理简单条件
			condResult, condFailed, err := s.evaluatePolicyWithDetails(andPart, attributes)
			if err != nil {
				return false, nil, err
			}
			if !condResult {
				allSatisfied = false
				allFailedConditions = append(allFailedConditions, condFailed...)
			}
		}
	}

	return allSatisfied, allFailedConditions, nil
}

// evaluateOrConditionWithDetails 评估OR条件并返回失败详情
func (s *ABEService) evaluateOrConditionWithDetails(orPolicy string, attributes map[string]string) (bool, []map[string]interface{}, error) {
	var allFailedConditions []map[string]interface{}

	// 去除括号
	orPolicy = strings.TrimSpace(orPolicy)
	if strings.HasPrefix(orPolicy, "(") && strings.HasSuffix(orPolicy, ")") {
		orPolicy = orPolicy[1 : len(orPolicy)-1]
	}

	// 分割OR条件
	orParts := s.splitPolicy(orPolicy, " OR ")

	for _, orPart := range orParts {
		orPart = strings.TrimSpace(orPart)

		// 递归处理每个OR部分
		result, failedConds, err := s.evaluatePolicyWithDetails(orPart, attributes)
		if err != nil {
			return false, nil, err
		}
		if result {
			return true, nil, nil // OR条件只要有一个为true就可以
		}
		allFailedConditions = append(allFailedConditions, failedConds...)
	}

	return false, allFailedConditions, nil
}

// buildDetailedFailureReason 构建详细的失败原因说明
func (s *ABEService) buildDetailedFailureReason(failedConditions []map[string]interface{}, vcAttributes map[string]string) string {
	if len(failedConditions) == 0 {
		return "未知原因"
	}

	var reasons []string

	for _, condition := range failedConditions {
		attribute, _ := condition["attribute"].(string)
		expectedValue, _ := condition["expectedValue"].(string)
		actualValue, _ := condition["actualValue"].(string)
		hasAttribute, _ := condition["hasAttribute"].(bool)

		var reason string
		if !hasAttribute {
			reason = fmt.Sprintf("缺少必需的属性：%s（期望值：%s）", attribute, expectedValue)
		} else {
			// 属性名称中文化
			attributeNameCN := s.getAttributeNameCN(attribute)
			reason = fmt.Sprintf("%s不匹配：期望%s，实际%s", attributeNameCN, expectedValue, actualValue)
		}
		reasons = append(reasons, reason)
	}

	return strings.Join(reasons, "；")
}

// getAttributeNameCN 获取属性的中文名称
func (s *ABEService) getAttributeNameCN(attribute string) string {
	attributeNames := map[string]string{
		"department":    "科室",
		"hospital":      "医院",
		"title":         "职称",
		"specialty":     "专长",
		"name":          "姓名",
		"licenseNumber": "执业编号",
		"wallet":        "钱包地址",
		"did":           "DID",
	}

	if cnName, exists := attributeNames[attribute]; exists {
		return cnName
	}
	return attribute
}

// splitPolicy 安全地分割策略字符串，考虑括号嵌套
func (s *ABEService) splitPolicy(policy string, separator string) []string {
	var parts []string
	var current strings.Builder
	var depth int

	for i := 0; i < len(policy); {
		if policy[i] == '(' {
			depth++
			current.WriteByte(policy[i])
			i++
		} else if policy[i] == ')' {
			depth--
			current.WriteByte(policy[i])
			i++
		} else if depth == 0 && strings.HasPrefix(policy[i:], separator) {
			// 在顶层遇到分隔符
			parts = append(parts, current.String())
			current.Reset()
			i += len(separator)
		} else {
			current.WriteByte(policy[i])
			i++
		}
	}

	if current.Len() > 0 {
		parts = append(parts, current.String())
	}

	return parts
}

// evaluateSimpleCondition 评估简单条件
func (s *ABEService) evaluateSimpleCondition(condition string, attributes map[string]string) (bool, error) {
	condition = strings.TrimSpace(condition)

	// 去除括号
	if strings.HasPrefix(condition, "(") && strings.HasSuffix(condition, ")") {
		condition = condition[1 : len(condition)-1]
	}

	// 解析条件格式：attribute:value 或 attribute operator value
	parts := strings.SplitN(condition, ":", 2)
	if len(parts) != 2 {
		return false, fmt.Errorf("无效的条件格式: %s", condition)
	}

	attributeName := strings.TrimSpace(parts[0])
	expectedValue := strings.TrimSpace(parts[1])

	// 移除引号
	expectedValue = strings.Trim(expectedValue, "\"'")

	// 获取VC中的属性值
	actualValue, exists := attributes[attributeName]
	if !exists {
		fmt.Printf("VC中缺少属性: %s\n", attributeName)
		return false, nil
	}

	// 比较值（不区分大小写）
	result := strings.EqualFold(actualValue, expectedValue)
	fmt.Printf("条件评估: %s:%s == %s -> %v\n", attributeName, actualValue, expectedValue, result)

	return result, nil
}

// 其他ABE数据库操作方法...
