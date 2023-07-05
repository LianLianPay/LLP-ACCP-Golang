package client

import (
	"LLP-ACCP-Go/main/accpapi/security"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func SendRequest(url, body string) string {
	if url == "" {
		panic("请求URL非法")
	}
	//生成数字签名
	signature := security.Sign(body)

	log.Printf("请求URL：%s\n", url)
	log.Printf("请求签名值：%s\n", signature)
	log.Printf("请求参数：%s\n", body)

	client := http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	if err != nil {
		log.Fatal("创建请求失败：", err)
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Signature-Type", "RSA")
	req.Header.Set("Signature-Data", signature)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("发送请求失败：", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic("请求结果异常，响应状态码为：" + strconv.Itoa(resp.StatusCode))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	respBodyStr := string(respBody)
	if err != nil {
		log.Fatal("读取响应失败：", err)
	}
	log.Printf("响应结果：%s\n", respBodyStr)

	respSignatureData := resp.Header.Get("Signature-Data")
	log.Printf("响应签名：%s\n", respSignatureData)

	//检验数字签名
	checkSign := security.CheckSign(respBodyStr, respSignatureData)
	if !checkSign {
		log.Printf("返回响应验证签名异常，请核实！")
	} else {
		log.Printf("响应验签通过！")
	}

	return respBodyStr
}
