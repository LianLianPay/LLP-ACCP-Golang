package test

import (
	"LLP-ACCP-Go/main/accpapi/config"
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const (
	serverPrivate = "MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAJJL8OLB0J/9pmzHFxpwOeigamHd3Yk6PkZdaL6reDOdlq5mOQ0/xIqXcnaWI/Q7qtT9j/b34hR74ZMyEw4Um5mbWG0C0qK7l6RbQaUExbF/gU+RiVCQ8TQW1qgw/eBh+H47Aj58hGulbfJKfeZJydzpnvTSdT9VitGR9xIJtKdHAgMBAAECgYBMmbzATnE5RGu+qyP6sOZxWoU5Rx03PCrdVw2AQHIIvKvoFxgqSshTNOc3Fngu6osRSM73pmVXCmJbWy3FAp9Rqg2FZfQoX+ds4cnj3QVpeILw6b2Sr0rI2OBkbXGFre/crM+JcjYBAkV7pnwcWRH3EyOvzLUqKs5qEkOycxTi8QJBAOUFVS8ipCnp7Qaynig6PcfJC0JP4GxpFmQu0w1OrmlzP/zezUfRwihTx1NPssJm9HD7KNiBDlgFj0PQJkGbB18CQQCjh90kBAoloAsCxe/qD4w7lbre75P16Kicb+K0FCeJsZrdXpApFhlDo60zPNUJEPph9HFptZfNBE8I8dIesHEZAkEAxe4V8Oa/ennxoBg/GAU936yhTm46R3eLIopVXOrjUb+JTcJBKBDg/Hlri1UV6W2RVRO7+WGQRAKKDtGWPpz9gQJAImZAFIVtBQEnj8vHbfsbSqVyi9blzwLEBTRcAfmDX6mmpA5yUNI/OkVB99dCEQgrQ1PCT7RNXGkdnwoPYzlGcQJBAJQQrWM8SxovyqcN7Md2wRvIjA1Ny7OJGSR8y+0eu/D0GydQbUj1rNdPX5CLNFVwvcgMwkLNUD+u+JSol5+PQHk="
	serverPublic  = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDopTeLUex6TWL0ga2vw83m3pqAIhrRnmsWvqnm8fztmAHVBTyPsUyk19hE8qksxn28V6Zkmse7VznaGt7smhZ4SkRkby3atZE0k7u6ubOoE68sVR9putj8vD8NLCnLrldAsbcZK8VMxkDfkEsX/k8jQL+a2Iv2sr7A4FnQd3EamwIDAQAB"
)

func Handle() {

	private := config.MerchantPrivateKey
	public := serverPublic

	m := make(map[string]string)
	m["param1"] = "参数1_/+-"
	m["param2"] = "参数2"
	m["param3"] = "参数3"

	fmt.Println("签名前的数据:", m)

	if err := signData(m, private); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("签名后的数据:", m)

	if err := verifyData(m, public); err != nil {
		fmt.Println("验签结果(公钥验签):", "失败", err)
	} else {
		fmt.Println("验签结果(公钥验签):", "成功")
	}
}

func signData(m map[string]string, private string) error {

	m["timestamp"] = fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	sig := url.Values{}
	for key, value := range m {
		sig.Add(key, value)
	}

	//进行转码使之可以安全的用在URL查询里
	quUrl, err := url.QueryUnescape(sig.Encode())
	if err != nil {
		fmt.Println("QueryUnescape", err)
		return err
	}

	if out, err := RsaSignWithMd5(quUrl, private); err != nil {
		return err
	} else {
		m["sig"] = out
		return nil
	}
}

func verifyData(m map[string]string, public string) error {
	sign := m["sig"]
	delete(m, "sig")

	sig := url.Values{}
	for key, value := range m {
		sig.Add(key, value)
	}

	//进行转码使之可以安全的用在URL查询里
	quUrl, _ := url.QueryUnescape(sig.Encode())

	return RsaVerifySignWithMd5(quUrl, sign, public)
}

// 签名
func RsaSignWithMd5(data string, prvKey string) (sign string, err error) {

	//如果密钥是urlSafeBase64的话需要处理下
	prvKey = Base64URLDecode(prvKey)

	keyBytes, err := base64.StdEncoding.DecodeString(prvKey)
	if err != nil {
		fmt.Println("DecodeString:", err)
		return "", err
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey", err)
		return "", err
	}

	h := md5.New()
	h.Write([]byte(data))
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.MD5, hash[:])
	if err != nil {
		fmt.Println("SignPKCS1v15:", err)
		return "", err
	}

	out := base64.RawURLEncoding.EncodeToString(signature)
	return out, nil
}

// 验签
func RsaVerifySignWithMd5(originalData, signData, pubKey string) error {

	//TODO : 验证时间

	sign, err := base64.RawURLEncoding.DecodeString(signData)
	if err != nil {
		fmt.Println("DecodeString:", err)
		return err
	}

	pubKey = Base64URLDecode(pubKey)

	public, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		fmt.Println("DecodeString")
		return err
	}

	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		fmt.Println("ParsePKIXPublicKey", err)
		return err
	}

	hash := md5.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.MD5, hash.Sum(nil), sign)
}

//因为Base64转码后可能包含有+,/,=这些不安全的URL字符串，所以要进行换字符
// '+' -> '-'
// '/' -> '_'
// '=' -> ''
// 字符串长度不足4倍的位补"="

func Base64URLDecode(data string) string {
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing) //字符串长度不足4倍的位补"="
	data = strings.Replace(data, "_", "/", -1)
	data = strings.Replace(data, "-", "+", -1)
	return data
}

func Base64UrlSafeEncode(data string) string {
	safeUrl := strings.Replace(data, "/", "_", -1)
	safeUrl = strings.Replace(safeUrl, "+", "-", -1)
	safeUrl = strings.Replace(safeUrl, "=", "", -1)
	return safeUrl
}
