package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/security"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/individual"
	"fmt"
)

/**
 * 个人用户新增绑卡 Demo
 */
func IndividualBindCardDemo() string {
	//绑卡申请
	params := individual.IndividualBindCardApplyParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.UserID = "LLianPayTest-Api-User-12345"
	params.TxnSeqno = "LLianPayTest" + timestamp
	params.TxnTime = timestamp
	params.NotifyURL = "https://test.lianlianpay/notify"
	// 设置银行卡号
	params.LinkedAcctNo = ""
	// 设置绑卡手机号
	params.LinkedPhone = ""
	// 设置钱包密码，正式环境要接密码控件，调试API可以用连连公钥加密密码
	params.Password = security.LocalEncrypt("")

	// 测试环境请求地址
	url := "https://accpgw-ste.lianlianpay-inc.com/v1/acctmgr/openacct-apply"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)

	var bindCardApplyResult individual.IndividualBindCardApplyResult
	err = utils.StringToObject(resultJsonStr, bindCardApplyResult)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return ""
	}

	//绑卡验证
	if !("0000" == bindCardApplyResult.RetCode) {
		fmt.Println("绑卡申请失败，请检查！")
		return ""
	}

	params2 := individual.IndividualBindCardVerifyParams{}
	timestamp = utils.GetTimestamp()
	params2.Timestamp = timestamp
	params2.OidPartner = bindCardApplyResult.OidPartner
	params2.UserID = bindCardApplyResult.UserID
	params2.TxnSeqno = bindCardApplyResult.TxnSeqno
	params2.Token = bindCardApplyResult.Token
	// 测试环境首次绑卡，不下发短信验证码，任意6位数字
	params2.VerifyCode = "123456"

	url2 := "https://accpgw-ste.lianlianpay-inc.com/v1/acctmgr/openacct-apply"
	paramsStr2, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr2 := client.SendRequest(url2, paramsStr2)
	return resultJsonStr2
}
