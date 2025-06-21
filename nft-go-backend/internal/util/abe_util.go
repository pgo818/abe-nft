package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ABE/nft/nft-go-backend/internal/config"

	"github.com/fentec-project/gofe/abe"
)

// ABEUtil ABE工具类
type ABEUtil struct {
	abe *ABE
}

// NewABEUtil 创建新的ABE工具实例
func NewABEUtil(securityParam int) *ABEUtil {
	return &ABEUtil{
		abe: NewABE(securityParam),
	}
}

// SetupABE 初始化ABE系统
func (a *ABEUtil) SetupABE(gamma []string) (*config.ABEPubkey, *config.ABESeckey, error) {
	return a.abe.Setup(gamma)
}

// KeyGenABE 生成属性密钥
func (a *ABEUtil) KeyGenABE(gamma []string, pk *config.ABEPubkey, msk *config.ABESeckey) (*config.ABEAttribKeys, error) {
	return a.abe.KeyGen(gamma, pk, msk)
}

// EncryptABE 加密消息
func (a *ABEUtil) EncryptABE(message, policy string, pk *config.ABEPubkey) (*config.ABECipher, error) {
	// 将策略字符串转换为MSP
	msp, err := abe.BooleanToMSP(policy, false)
	if err != nil {
		return nil, fmt.Errorf("策略解析失败: %v", err)
	}

	return a.abe.Encrypt(message, msp, pk)
}

// DecryptABE 解密消息
func (a *ABEUtil) DecryptABE(cipher *config.ABECipher, key *config.ABEAttribKeys, pk *config.ABEPubkey) (string, error) {
	return a.abe.Decrypt(cipher, key, pk)
}

// SerializePubKey 序列化公钥
func (a *ABEUtil) SerializePubKey(pk *config.ABEPubkey) (string, error) {
	data, err := json.Marshal(pk)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// DeserializePubKey 反序列化公钥
func (a *ABEUtil) DeserializePubKey(pkStr string) (*config.ABEPubkey, error) {
	data, err := base64.StdEncoding.DecodeString(pkStr)
	if err != nil {
		return nil, err
	}

	var pk config.ABEPubkey
	err = json.Unmarshal(data, &pk)
	if err != nil {
		return nil, err
	}

	return &pk, nil
}

// SerializeSecKey 序列化私钥
func (a *ABEUtil) SerializeSecKey(sk *config.ABESeckey) (string, error) {
	data, err := json.Marshal(sk)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// DeserializeSecKey 反序列化私钥
func (a *ABEUtil) DeserializeSecKey(skStr string) (*config.ABESeckey, error) {
	data, err := base64.StdEncoding.DecodeString(skStr)
	if err != nil {
		return nil, err
	}

	var sk config.ABESeckey
	err = json.Unmarshal(data, &sk)
	if err != nil {
		return nil, err
	}

	return &sk, nil
}

// SerializeAttribKeys 序列化属性密钥
func (a *ABEUtil) SerializeAttribKeys(keys *config.ABEAttribKeys) (string, error) {
	data, err := json.Marshal(keys)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// DeserializeAttribKeys 反序列化属性密钥
func (a *ABEUtil) DeserializeAttribKeys(keysStr string) (*config.ABEAttribKeys, error) {
	data, err := base64.StdEncoding.DecodeString(keysStr)
	if err != nil {
		return nil, err
	}

	var keys config.ABEAttribKeys
	err = json.Unmarshal(data, &keys)
	if err != nil {
		return nil, err
	}

	return &keys, nil
}

// SerializeCipher 序列化密文
func (a *ABEUtil) SerializeCipher(cipher *config.ABECipher) (string, error) {
	data, err := json.Marshal(cipher)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

// DeserializeCipher 反序列化密文
func (a *ABEUtil) DeserializeCipher(cipherStr string) (*config.ABECipher, error) {
	data, err := base64.StdEncoding.DecodeString(cipherStr)
	if err != nil {
		return nil, err
	}

	var cipher config.ABECipher
	err = json.Unmarshal(data, &cipher)
	if err != nil {
		return nil, err
	}

	return &cipher, nil
}
