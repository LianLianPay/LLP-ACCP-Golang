package security

import (
	"LLP-ACCP-Go/main/accpapi/config"
	"encoding/hex"
	"fmt"
	"log"
)

// SignWithDefaultPrivateKey 签名处理，默认使用lianLianPayConstant中配置的私钥MerchantPrivateKey
func SignWithDefaultPrivateKey(body string) string {
	return Sign(config.MerchantPrivateKey, body)
}

// Sign 签名处理
func Sign(privateKey, body string) string {
	hashedBody := Md5Digest([]byte(body))
	log.Printf("签名处理中，签名源内容：%s，对应MD5值：%s", body, hex.EncodeToString(hashedBody))
	// 将md5值进行rsa加密
	signed, err := RsaSign(privateKey, hashedBody)
	if err != nil {
		log.Println("rsa加密失败:", err)
	}
	return signed
}

// CheckSignWithDefaultPublicKey 签名验证，默认使用lianLianPayConstant中配置的公钥LLianPayPublicKey
func CheckSignWithDefaultPublicKey(body, signedStr string) bool {
	return CheckSign(config.LLianPayPublicKey, body, signedStr)
}

// CheckSign 签名验证
func CheckSign(publicKey, body, signedStr string) bool {
	hashedBody := Md5Digest([]byte(body))
	log.Printf("签名验证中，源串：%s，对应MD5值：%s", body, hex.EncodeToString(hashedBody))
	sign, err := CheckRsaSign(publicKey, hashedBody, signedStr)
	if err != nil {
		fmt.Printf("签名验证失败：%v\n", err)
		return false
	}
	return sign
}
