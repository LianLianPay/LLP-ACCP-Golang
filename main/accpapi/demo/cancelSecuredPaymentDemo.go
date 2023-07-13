package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/txn/secured"
	"fmt"
)

/**
 * 担保确认撤销 Demo
 */
func CancelSecuredPayment() string {
	params := secured.CancelSecuredPaymentParams{
		Timestamp:  utils.GetTimestamp(),
		OidPartner: "LLianPayConstant.OidPartner",
		TxnSeqNo:   "LLianPayTest20230110124354",
	}

	// 测试环境请求地址
	url := "https://192.168.110.93/v1/txn/cancel-secured-payment"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
