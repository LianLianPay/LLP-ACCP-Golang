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
 * 匿名用户解绑银行卡 Demo
 */
func UnLinkedAcctIndApply() string {
	params := unlinked.UnLinkedAcctIndApplyParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.UserID = "LLianPayTest-An-User-12345"
	params.TxnSeqno = "LLianPayTest" + timestamp
	params.TxnTime = timestamp
	params.NotifyURL = "https://test.lianlianpay/notify"
	// 绑定银行账号。该字段需要RSA公钥加密传输
	params.LinkedAcctno = security.LocalEncrypt("")
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v2/acctmgr/unlinkedacct-ind-apply"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
