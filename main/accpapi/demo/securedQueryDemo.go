package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/txn/secured"
	"fmt"
	"time"
)

/**
 * 担保交易信息查询 Demo
 */
func SecuredQuery() string {
	params := secured.SecuredQueryParams{}
	params.Timestamp = time.Now().Format("20060102150405") // 获取当前时间戳
	params.OidPartner = "your_oid_partner"
	params.TxnSeqNo = "202302071352032021805645"
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/txn/secured-query"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
