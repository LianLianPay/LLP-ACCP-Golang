package security

import (
	"LLP-ACCP-Go/main/accpapi/config"
	"fmt"
	"github.com/yuchenfw/gocrypt"
	"github.com/yuchenfw/gocrypt/rsa"
)

/*
*
该文件使用第三方开源库实现加签，经测试与手动实现的一致
github链接: https://github.com/yuchenfw/gocrypt
*/
func Handle(source string) string {
	secretInfo := rsa.RSASecret{
		PublicKey:          config.LLianPayPublicKey,
		PublicKeyDataType:  gocrypt.Base64,
		PrivateKey:         config.MerchantPrivateKey,
		PrivateKeyType:     gocrypt.PKCS8,
		PrivateKeyDataType: gocrypt.Base64,
	}

	handleRSA := rsa.NewRSACrypt(secretInfo) //RSA

	sign, err := handleRSA.Sign(source, gocrypt.MD5, gocrypt.Base64)
	if err != nil {
		fmt.Println("sign error :", err)
		return ""
	}
	fmt.Println("sign data :", sign)
	return sign
}
