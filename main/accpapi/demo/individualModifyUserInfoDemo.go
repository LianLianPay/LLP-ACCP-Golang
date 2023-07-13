package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/individual"
	"fmt"
)

/**
 * 个人用户信息修改 Demo
 */
func IndividualModifyUserInfoDemo() string {
	params := individual.IndividualModifyUserInfoParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.UserID = "LLianPayTest-Api-User-12345"
	params.TxnSeqno = "LLianPayTest" + timestamp
	params.TxnTime = timestamp
	params.NotifyURL = "https://test.lianlianpay/notify"

	basicInfo := individual.IndividualModifyUserInfoBasicInfo{}
	basicInfo.IDEmblem = "202209010003370005"
	basicInfo.IDPortrait = "202209010003370005"
	params.BasicInfo = basicInfo

	url := "https://accpapi-ste.lianlianpay-inc.com/v1/acctmgr/modify-userinfo-individual"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
