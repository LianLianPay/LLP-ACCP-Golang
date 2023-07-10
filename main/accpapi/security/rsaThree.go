package security

import (
	"LLP-ACCP-Go/main/accpapi/config"
	"fmt"
	"github.com/yuchenfw/gocrypt"
	"github.com/yuchenfw/gocrypt/rsa"
)

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
