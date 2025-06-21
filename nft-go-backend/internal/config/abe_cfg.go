package config

import (
	"math/big"

	"github.com/fentec-project/bn256"
	"github.com/fentec-project/gofe/abe"
)

// ABEParams 定义ABE方案的参数
type ABEParams struct {
	P *big.Int // 素数阶，通常使用bn256的阶数
}

// ABE 定义基于属性的加密方案
type ABE struct {
	Params *ABEParams
}

// ABESeckey 主密钥结构
type ABESeckey struct {
	PartG2_s *bn256.G2 // α ∈ Z_p, g^α ∈ G2
}

// ABEPubkey 公钥结构
type ABEPubkey struct {
	PartG2_p *bn256.G2 // a ∈ Z_p, g^a ∈ G2
	PartGT   *bn256.GT // e(g,g)^α ∈ GT
}

// ABEAttribKeys 属性密钥结构
type ABEAttribKeys struct {
	K         *bn256.G2      // K = g^α * (g^a)^t ∈ G2
	L         *bn256.G1      // L = g^t ∈ G1
	Kx        []*bn256.G2    // 每个属性对应的密钥组件
	AttribToI map[string]int // 属性到索引的映射
}

// ABECipher 密文结构
type ABECipher struct {
	C      *bn256.GT   // C = e(g,g)^(αs) * key
	CPrime *bn256.G1   // C' = g^s
	Ci     []*bn256.G2 // 属性相关组件
	Di     []*bn256.G1 // 属性相关组件
	Msp    *abe.MSP    // 访问策略
	SymEnc []byte      // 对称加密的消息
	Iv     []byte      // 初始向量
}
