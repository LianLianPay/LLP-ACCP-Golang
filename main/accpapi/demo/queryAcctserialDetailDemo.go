package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/query"
	"fmt"
	"time"
)

/**
 * 资金流水详情查询 Demo
 * 需要和资金流水列表查询配合使用
 */
func QueryAcctserialDetail() string {
	params := query.AcctserialDetailParams{}
	timestamp := time.Now().Format("20060102150405") // 获取当前时间戳
	params.Timestamp = timestamp
	params.OidPartner = "your_oid_partner"
	params.UserID = "verify-code0ddfb801-8c34-45c4-854d-81834bc27deb"
	params.UserType = "INNERUSER"
	params.JnoAcct = "20220824000004471563"
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/acctmgr/query-acctserialdetail"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
