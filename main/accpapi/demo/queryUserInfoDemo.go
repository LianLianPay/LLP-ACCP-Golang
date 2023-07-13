package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/query"
	"fmt"
)

/**
 * 用户信息查询 Demo
 */
func QueryUserInfoDemo() string {
	params := query.UserInfoParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = "your_oid_partner"
	params.UserID = "bdb1dd105679979ca82b28edd1c8ccd2"
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/acctmgr/query-userinfo"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
