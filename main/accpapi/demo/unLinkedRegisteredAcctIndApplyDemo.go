package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/security"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/unlinked"
	"fmt"
)

/**
 * 个人用户解绑银行卡 Demo
 */
func UnLinkedRegisteredAcctIndApply() string {
	params := unlinked.UnLinkedAcctIndApplyParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.UserID = "LLianPayTest-Api-User-12345"
	params.TxnSeqno = "LLianPayTest" + timestamp
	params.TxnTime = timestamp
	params.NotifyURL = "https://test.lianlianpay/notify"
	// 绑定银行账号
	params.LinkedAcctno = ""
	// 用户：LLianPayTest-Api-User-12345 密码：qwerty，本地测试环境测试，没接入密码控件，使用本地加密方法加密密码（仅限测试环境使用）
	params.Password = security.LocalEncrypt("qwerty")

	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/acctmgr/unlinkedacct-ind-apply"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
