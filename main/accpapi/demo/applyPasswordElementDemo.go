package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/password"
	"fmt"
)

func ApplyPasswordElement() string {
	params := password.ApplyPasswordElementParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.TxnSeqno = "LLianPayTest" + timestamp
	/*
	   密码使用场景：
	   设置密码：setting_password
	   修改密码：change_password
	   换绑卡：bind_card_password
	   提现密码：cashout_password
	   支付密码：pay_password
	*/
	params.PasswordScene = "setting_password"
	params.FlagChnl = "H5"

	// 测试环境URL
	url := "https://accpgw-ste.lianlianpay-inc.com/v1/acctmgr/apply-password-element"
	//url = "http://localhost:8086/accpapi/v1/acctmgr/apply-password-element"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)

	var applyPasswordElementParams password.ApplyPasswordElementParams
	err = utils.StringToObject(resultJsonStr, applyPasswordElementParams)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return ""
	}
	return resultJsonStr
}
