package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/close"
	"fmt"
	"time"
)

/**
 * 支付单关单 Demo
 */
func ClosePayment() string {
	params := close.ClosePaymentParams{}

	timestamp := time.Now().Format("20060102150405")
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.AccpTxno = "2022082716141713"
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/txn/close-payment"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
