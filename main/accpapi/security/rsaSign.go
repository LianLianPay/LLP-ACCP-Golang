package security

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
)

// RsaSign 签名处理
func RsaSign(privateKey string, hashedBody []byte) (string, error) {
	privateKey = "-----BEGIN PRIVATE KEY-----\n" + privateKey + "\n-----END PRIVATE KEY-----\n"
	//fmt.Println(privateKey)
	// 解码私钥
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", fmt.Errorf("failed to decode private key")
	}

	// 解析RSA私钥
	priKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("无法解析RSA私钥：%v", err)
	}

	// 转换为RSA私钥类型
	rsaPriKey, ok := priKey.(*rsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("私钥不是RSA密钥")
	}

	// 使用RSA私钥进行签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPriKey, crypto.MD5, hashedBody)
	if err != nil {
		return "", fmt.Errorf("无法签名摘要：%v", err)
	}

	// 返回Base64编码的签名结果
	return base64.StdEncoding.EncodeToString(signature), nil
}

// CheckRsaSign 签名验证
func CheckRsaSign(publicKey string, hashedBody []byte, signedStr string) (bool, error) {
	publicKey = "-----BEGIN PUBLIC KEY-----\n" + publicKey + "\n-----END PUBLIC KEY-----\n"
	// 解码公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return false, fmt.Errorf("failed to decode public key")
	}

	// 解析公钥
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, fmt.Errorf("无法解析公钥：%v", err)
	}

	// 转换为RSA公钥类型
	rsaPubKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("公钥不是RSA密钥")
	}

	// 解码已签名的字符串
	signature, _ := hex.DecodeString(signedStr)

	// 验证签名
	err = rsa.VerifyPKCS1v15(rsaPubKey, crypto.MD5, hashedBody, signature)
	if err != nil {
		return false, fmt.Errorf("签名验证失败：%v", err)
	}

	return true, nil
}
