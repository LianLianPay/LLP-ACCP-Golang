package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/security"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/individual"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/regphone"
	"fmt"
)

/**
 * 个人用户API开户流程 Demo
 */
func IndividualOpenacct() {
	userId := "LLianPayTest-Api-User-12345"
	// 绑定手机验证码申请
	applyResult := phoneVerifyCodeApply(userId)
	if applyResult.RetCode != "0000" {
		fmt.Println("绑定手机号验证申请失败，请检查！")
		return
	}

	// 触发绑定手机验证码申请时，会下发短信验证码
	// 用Debug模式，断点打到这里，Debug的时候把verifyCode设置成手机收到的真实验证码
	verifyCode := ""
	// 绑定手机验证码验证
	verifyResult := PhoneVerifyCodeVerify(userId, verifyCode)
	if verifyResult.RetCode != "0000" {
		fmt.Println("绑定手机号验证验证失败，请检查！")
		return
	}

	// 个人开户申请
	openacctApplyResult := OpenacctApply(userId)

	// 随机密码因子获取接口
	// 测试环境没有接密码控件的话，可以用连连公钥进行加密调试接口
	password := "qwerty"
	encryptStr := security.LocalEncrypt(password)
	// 个人开户验证
	// 用Debug模式，断点打到这里，Debug的时候把verifyCode设置成手机收到的真实验证码
	openacctVerifyCode := ""
	openacctVerify(openacctApplyResult, openacctVerifyCode, encryptStr)
}

// 绑定手机验证码申请
func phoneVerifyCodeApply(userId string) regphone.VerifyCodeResult {
	params := regphone.VerifyCodeApplyParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.UserID = userId
	// 设置手机号
	params.RegPhone = ""

	// 测试环境URL
	url := "https://accpgw-ste.lianlianpay-inc.com/v1/acctmgr/apply-password-element"
	var verifyCodeResult regphone.VerifyCodeResult
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return verifyCodeResult
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)

	err = utils.StringToObject(resultJsonStr, verifyCodeResult)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return verifyCodeResult
	}
	return verifyCodeResult
}

// 绑定手机验证码验证
func PhoneVerifyCodeVerify(userId string, verifyCode string) regphone.VerifyCodeResult {
	params := regphone.VerifyCodeVerifyParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.UserID = userId
	// 设置手机号
	params.RegPhone = ""
	params.VerifyCode = verifyCode
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/acctmgr/regphone-verifycode-verify"
	var verifyCodeResult regphone.VerifyCodeResult
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return verifyCodeResult
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)

	err = utils.StringToObject(resultJsonStr, verifyCodeResult)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return verifyCodeResult
	}
	return verifyCodeResult
}

// 个人开户申请
func OpenacctApply(userId string) individual.OpenacctApplyResult {
	params := individual.OpenacctApplyParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.UserID = userId
	params.TxnSeqno = "LLianPayTest" + timestamp
	params.TxnTime = timestamp
	params.NotifyURL = "https://test.lianlianpay/notify"

	// 设置开户基本信息
	basicInfo := individual.OpenacctApplyBasicInfo{}
	// 设置手机号
	basicInfo.RegPhone = ""
	// 设置用户名称，企业用户传企业名称，个体工商户传营业执照的名称，个人用户传姓名（正式环境会进行审核的）
	basicInfo.UserName = ""
	basicInfo.IDType = "ID_CARD"
	// 设置证件号
	basicInfo.IDNo = ""
	// 设置证件有效期
	basicInfo.IDExp = "99991231"
	// 设置地址
	basicInfo.Address = ""
	// 设置职业
	basicInfo.Occupation = "01"
	params.BasicInfo = basicInfo

	// 开户账户申请信息
	accountInfo := individual.OpenacctApplyAccountInfo{}
	// 个人用户建议设置成个人支付账户类型
	accountInfo.AccountType = "PERSONAL_PAYMENT_ACCOUNT"
	accountInfo.AccountNeedLevel = "V3"
	params.AccountInfo = accountInfo

	// 测试环境URL
	url := "https://accpgw-ste.lianlianpay-inc.com/v1/acctmgr/apply-password-element"
	var openacctApplyResult individual.OpenacctApplyResult
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return openacctApplyResult
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)

	err = utils.StringToObject(resultJsonStr, openacctApplyResult)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return openacctApplyResult
	}
	return openacctApplyResult
}

// 个人开户验证
func openacctVerify(openacctApplyResult individual.OpenacctApplyResult, verifyCode string, password string) {
	params := individual.OpenacctVerifyParams{
		Timestamp:  utils.GetTimestamp(),
		OidPartner: config.OidPartner,
		UserID:     openacctApplyResult.UserID,
		TxnSeqno:   openacctApplyResult.TxnSeqno,
		Token:      openacctApplyResult.Token,
		// Verify_code: verifyCode,
		Password: password,
	}
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/acctmgr/openacct-verify-individual"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	fmt.Println(resultJsonStr)
}
