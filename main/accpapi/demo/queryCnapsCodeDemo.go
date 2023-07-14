package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/cnapscode"
	"fmt"
)

/**
 * 大额行号查询 Demo
 */
func QueryCnapsCode() string {
	params := cnapscode.QueryParams{}
	timestamp := utils.GetTimestamp()
	params.OidPartner = "your_oid_partner"
	params.Timestamp = timestamp
	params.BankCode = "01030000"
	params.BrabankName = "行"
	params.CityCode = "330100"

	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/acctmgr/query-cnapscode"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
