package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/txn/payment"
	"fmt"
)

/**
 * 支付结果查询 Demo
 */
func QueryPaymentDemo() string {
	params := payment.QueryPaymentParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = "your_oid_partner"
	params.AccpTxNo = "2022111417141745"
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/txn/query-payment"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
