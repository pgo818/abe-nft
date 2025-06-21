package util

import (
	"crypto/aes"
	cbc "crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"math/big"

	"github.com/ABE/nft/nft-go-backend/internal/config"
	"github.com/fentec-project/bn256"
	"github.com/fentec-project/gofe/abe"
	"github.com/fentec-project/gofe/data"
	"github.com/fentec-project/gofe/sample"
)

// ABE 是对 config.ABE 的包装类型
type ABE struct {
	*config.ABE
}

// NewABE 初始化ABE实例
func NewABE(l int) *ABE {
	return &ABE{
		ABE: &config.ABE{
			Params: &config.ABEParams{
				P: bn256.Order, // 使用bn256曲线的阶数作为参数
			},
		},
	}
}

// Setup 初始化ABE系统，生成公钥和主密钥
func (a *ABE) Setup(gamma []string) (*config.ABEPubkey, *config.ABESeckey, error) {
	// 从均匀分布中采样随机数
	sampler := sample.NewUniform(a.Params.P)
	val, err := data.NewRandomVector(2, sampler)
	if err != nil {
		return nil, nil, err
	}

	// 生成主密钥和公钥组件
	partG2_s := new(bn256.G2).ScalarBaseMult(val[0]) // g^α
	partG2_p := new(bn256.G2).ScalarBaseMult(val[1]) // g^a
	partGT := new(bn256.GT).ScalarBaseMult(val[0])   // e(g,g)^α

	return &config.ABEPubkey{PartG2_p: partG2_p, PartGT: partGT}, &config.ABESeckey{PartG2_s: partG2_s}, nil
}

// KeyGen 根据属性集合生成用户密钥
func (a *ABE) KeyGen(gamma []string, pk *config.ABEPubkey, msk *config.ABESeckey) (*config.ABEAttribKeys, error) {
	// 随机选择t ∈ Z_p
	sampler := sample.NewUniform(a.Params.P)
	t, err := sampler.Sample()
	if err != nil {
		return nil, err
	}

	// 计算K = g^α * (g^a)^t
	K := new(bn256.G2).Add(msk.PartG2_s, new(bn256.G2).ScalarMult(pk.PartG2_p, t))
	// 计算L = g^t
	L := new(bn256.G1).ScalarBaseMult(t)

	// 为每个属性生成密钥组件
	Kx := make([]*bn256.G2, len(gamma))
	attribToI := make(map[string]int)
	for i, s := range gamma {
		// 计算H(s)^t，其中H是到G2的哈希函数
		hsi, err := bn256.HashG2(s)
		if err != nil {
			return nil, err
		}
		hsi.ScalarMult(hsi, t)

		Kx[i] = hsi
		attribToI[s] = i
	}

	return &config.ABEAttribKeys{K: K, L: L, Kx: Kx, AttribToI: attribToI}, nil
}

// Encrypt 加密消息，使用ABE封装对称密钥
func (a *ABE) Encrypt(msg string, msp *abe.MSP, pk *config.ABEPubkey) (*config.ABECipher, error) {
	if len(msp.Mat) == 0 || len(msp.Mat[0]) == 0 {
		return nil, fmt.Errorf("empty msp matrix")
	}

	// 检查属性是否重复
	attrib := make(map[string]bool)
	for _, i := range msp.RowToAttrib {
		if attrib[i] {
			return nil, fmt.Errorf("some attributes correspond to" +
				"multiple rows of the MSP struct, the scheme is not secure")
		}
		attrib[i] = true
	}

	// 生成随机对称密钥
	_, keyGt, err := bn256.RandomGT(rand.Reader)
	if err != nil {
		return nil, err
	}
	keyCBC := sha256.Sum256([]byte(keyGt.String()))

	// 初始化AES-CBC加密器
	c, err := aes.NewCipher(keyCBC[:])
	if err != nil {
		return nil, err
	}

	iv := make([]byte, c.BlockSize())
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}
	encrypterCBC := cbc.NewCBCEncrypter(c, iv)

	// PKCS7填充消息
	msgByte := []byte(msg)
	padLen := c.BlockSize() - (len(msgByte) % c.BlockSize())
	msgPad := make([]byte, len(msgByte)+padLen)
	copy(msgPad, msgByte)
	for i := len(msgByte); i < len(msgPad); i++ {
		msgPad[i] = byte(padLen)
	}

	// 执行对称加密
	symEnc := make([]byte, len(msgPad))
	encrypterCBC.CryptBlocks(symEnc, msgPad)

	// 生成共享向量v
	sampler := sample.NewUniform(a.Params.P)
	v, err := data.NewRandomVector(msp.Mat.Cols(), sampler)
	if err != nil {
		return nil, err
	}
	lamma, err := msp.Mat.MulVec(v)
	if err != nil {
		return nil, err
	}

	// 生成属性相关组件
	Ci := make([]*bn256.G2, len(msp.Mat))
	Di := make([]*bn256.G1, len(msp.Mat))
	r, err := data.NewRandomVector(len(msp.Mat), sampler)
	if err != nil {
		return nil, err
	}
	for i := 0; i < msp.Mat.Rows(); i++ {
		hsi, err := bn256.HashG2(msp.RowToAttrib[i])
		if err != nil {
			return nil, err
		}
		hsi = hsi.ScalarMult(hsi, r[i])
		hsi.Neg(hsi)

		// 计算Ci = (g^a)^λ_i * H(att(i))^(-r_i)
		if lamma[i].Sign() == -1 {
			lamma[i].Neg(lamma[i])
			partCi := new(bn256.G2).ScalarMult(pk.PartG2_p, lamma[i])
			partCi.Neg(partCi)
			partCi.Add(partCi, hsi)
			Ci[i] = partCi
		} else {
			partCi := new(bn256.G2).ScalarMult(pk.PartG2_p, lamma[i])
			partCi.Add(partCi, hsi)
			Ci[i] = partCi
		}

		// 计算Di = g^r_i
		partDi := new(bn256.G1).ScalarBaseMult(r[i])
		Di[i] = partDi
	}

	// 计算C = e(g,g)^(αs) * key
	C := new(bn256.GT).ScalarMult(pk.PartGT, v[0])
	C.Add(C, keyGt)

	// 计算C' = g^s
	CPrime := new(bn256.G1).ScalarBaseMult(v[0])

	return &config.ABECipher{C: C, CPrime: CPrime, Ci: Ci, Di: Di, Msp: msp, SymEnc: symEnc, Iv: iv}, nil
}

// Decrypt 解密消息
func (a *ABE) Decrypt(cipher *config.ABECipher, key *config.ABEAttribKeys, pk *config.ABEPubkey) (string, error) {
	// 确定用户拥有的属性
	attribMap := make(map[string]bool)
	for k := range key.AttribToI {
		attribMap[k] = true
	}

	// 筛选满足策略的属性
	countAttrib := 0
	for i := 0; i < len(cipher.Msp.Mat); i++ {
		if attribMap[cipher.Msp.RowToAttrib[i]] {
			countAttrib++
		}
	}

	// 构建用于解密的矩阵和组件
	preMatForKey := make([]data.Vector, countAttrib)
	CiForKey := make([]*bn256.G2, countAttrib)
	DiForKey := make([]*bn256.G1, countAttrib)
	rowToAttrib := make([]string, countAttrib)
	countAttrib = 0
	for i := 0; i < len(cipher.Msp.Mat); i++ {
		if attribMap[cipher.Msp.RowToAttrib[i]] {
			preMatForKey[countAttrib] = cipher.Msp.Mat[i]
			CiForKey[countAttrib] = cipher.Ci[i]
			DiForKey[countAttrib] = cipher.Di[i]
			rowToAttrib[countAttrib] = cipher.Msp.RowToAttrib[i]
			countAttrib++
		}
	}

	matForKey, err := data.NewMatrix(preMatForKey)
	if err != nil {
		return "", fmt.Errorf("the provided cipher is faulty")
	}

	if len(matForKey) == 0 {
		return "", fmt.Errorf("provided key is not sufficient for decryption")
	}

	// 解线性方程组找到重构系数
	oneVec := data.NewConstantVector(len(matForKey[0]), big.NewInt(0))
	oneVec[0].SetInt64(1)
	alpha, err := data.GaussianEliminationSolver(matForKey.Transpose(), oneVec, a.Params.P)
	if err != nil {
		return "", fmt.Errorf("provided key is not sufficient for decryption")
	}

	// 重构对称密钥
	keyGt := new(bn256.GT).Set(cipher.C)

	for i, e := range rowToAttrib {
		// 计算e(L, Ci) * e(Di, Kx)
		CiPairing := bn256.Pair(key.L, CiForKey[i])
		DiPairing := bn256.Pair(DiForKey[i], key.Kx[key.AttribToI[e]])

		if alpha[i].Sign() == -1 {
			alpha[i].Neg(alpha[i])
			partPairing := new(bn256.GT).ScalarMult(new(bn256.GT).Add(CiPairing, DiPairing), alpha[i])
			partPairing.Neg(partPairing)
			keyGt.Add(keyGt, partPairing)
		} else {
			partPairing := new(bn256.GT).ScalarMult(new(bn256.GT).Add(CiPairing, DiPairing), alpha[i])
			keyGt.Add(keyGt, partPairing)
		}
	}

	// 计算e(C', K)
	keyPairing := bn256.Pair(cipher.CPrime, key.K)
	keyPairing.Neg(keyPairing)
	keyGt.Add(keyGt, keyPairing)

	// 从重构的GT元素派生对称密钥
	keyCBC := sha256.Sum256([]byte(keyGt.String()))

	// 解密消息
	c, err := aes.NewCipher(keyCBC[:])
	if err != nil {
		return "", err
	}

	msgPad := make([]byte, len(cipher.SymEnc))
	decrypter := cbc.NewCBCDecrypter(c, cipher.Iv)
	decrypter.CryptBlocks(msgPad, cipher.SymEnc)

	// 移除填充
	padLen := int(msgPad[len(msgPad)-1])
	if (len(msgPad) - padLen) < 0 {
		return "", fmt.Errorf("failed to decrypt")
	}
	msgByte := msgPad[0:(len(msgPad) - padLen)]

	return string(msgByte), nil
}
