package security

import (
	"LLP-ACCP-Go/main/accpapi/config"
	"encoding/hex"
	"fmt"
	"log"
)

// Sign 加签
func Sign(body string) string {
	privateKey := config.MerchantPrivateKey
	privateKey = "MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBALLZeToAWlGYaCy+rN4RdFEB9mMr18rS/70LOglInqb9Q7KGhRs0dtT43zm35r/deEPwHKkI/u7GSz0S5lxO0LnvgvxqWqpfD8VZ1JYPHW6hNd1IN1uL5IQz8ThYLkFeJxe8/kGDjIL3UQlEq2xZRjFmamMI0GlHtZveBfJo2sAbAgMBAAECgYAUaPsLonSsMR9aU9w4b3d9syKiCKPUsqWatkhGPPJWDjsGiZFxHFifT1NyK8l6hdSutB0hMTSenSSC8MA3MothK47l5kNvW633a/TfsNps+BU8uc5jHjFpFJr6h1XK6Z/irxNpR4LN3Qi1nivEd+euBR+KwQ5lPUk/2cWdcXIMGQJBAO8zKdNSryFl3Oozjh2WN8h7MPEA9Xyru4L5GDJU8FTt9R7vtiidjwOuD0jPOe7KZiiP9DXmWRv7kRy29mXUlzcCQQC/aTR62V/NYYKBE2ccI7dTVL7J6RyeNpV74LVOz2p0OwHchZTOHJoSSU7XkGRDzK4VyS+Df1xXdHMM+Sodzwg9AkATFNXFUdc+ps2UsUApBA63I3yn/ReYNmri69QOT9BMNh9TtHOnkcoS1RPLDxzj+6limvk0Q4fgP5U34TCOtwLJAkEAqNEX6RWYyiIrIL5FHWtev88pahkAFYYctyILppOG1zZHwP/LqzFMrb4cHEEMPUPjcfF7x2VnwvZAlnuXKq0zHQJBAJugRT4AcL9zLD2Czm7pJ+9LOoO9MjKDTJsWWF8zjlPLgIAQVVk7lzf/xrQPcCOLZ7X71XoJdQZkVkmwyJUtGk4="
	hashedBody := Md5Digest(body)
	log.Printf("签名处理中，签名源内容：%s，对应MD5值：%s", body, hex.EncodeToString(hashedBody))
	signed, err := RsaSign(privateKey, hashedBody)
	if err != nil {
		log.Println("rsa加密失败:", err)
	}
	return signed
}

// CheckSign 验签
func CheckSign(body, signature string) bool {
	publicKey := config.LLianPayPublicKey
	hashedBody := Md5Digest(body)
	log.Printf("签名验证中，源串：%s，对应MD5值：%s", body, hex.EncodeToString(hashedBody))
	sign, err := CheckRsaSign(publicKey, hashedBody, signature)
	if err != nil {
		fmt.Printf("签名验证失败：%v\n", err)
		return false
	}
	return sign
}
