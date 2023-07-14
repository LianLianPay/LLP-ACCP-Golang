package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/query"
	"fmt"
)

/**
 * 资金流水列表查询 Demo
 */
func QueryAcctserial() string {
	params := query.AcctserialParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = "your_oid_partner"
	params.UserID = "verify-code0ddfb801-8c34-45c4-854d-81834bc27deb"
	params.UserType = "INNERUSER"
	params.AcctType = "USEROWN_PSETTLE"
	params.DateStart = "20220824000000"
	params.DateEnd = timestamp
	params.PageNo = "1"
	params.PageSize = "10"

	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/acctmgr/query-acctserial"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
