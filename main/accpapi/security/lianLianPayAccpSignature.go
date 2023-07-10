package security

import (
	"LLP-ACCP-Go/main/accpapi/config"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
)

// 加签
func Sign(source string) string {
	privateKey := config.MerchantPrivateKey
	hashedSource := md5Digest(source)
	log.Printf("加签源内容对应MD5值：%s", hex.EncodeToString(hashedSource))
	signed, err := RsaSign(privateKey, hashedSource)
	if err != nil {
		log.Println("rsa加密失败:", err)
	}
	return signed
}

// 验签
func CheckSign(source, signature string) bool {
	publicKey := config.LLianPayPublicKey
	hashedSource := md5Digest(source)
	log.Printf("验签源内容对应MD5值：%s", hex.EncodeToString(hashedSource))
	sign, err := CheckRsaSign(publicKey, hashedSource, signature)
	if err != nil {
		fmt.Printf("签名验证失败：%v\n", err)
		return false
	}
	return sign
}

func md5Digest(str string) []byte {
	hash := md5.New()
	hash.Write([]byte(str))
	return hash.Sum(nil)
}
