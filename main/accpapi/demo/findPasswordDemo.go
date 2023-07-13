package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/find"
	"fmt"
)

/**
 * 找回密码申请&验证 Demo
 */
func FindPassword() string {
	// 创建 FindPasswordApplyParams 结构体实例
	params := find.FindPasswordApplyParams{}
	// 获取当前时间戳
	params.Timestamp = utils.GetTimestamp()
	// 设置合作伙伴标识
	params.OidPartner = config.OidPartner
	// 设置用户标识
	params.UserID = "LLianPayTest-Api-User-12345abcd"
	// 绑定银行账号。个人用户绑定的银行卡号，若未绑卡，则不传。
	//params.LinkedAcctno = ""

	// 测试环境请求地址
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/acctmgr/modify-userinfo-enterprise"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}
	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	var findPasswordApplyResult find.FindPasswordApplyResult
	err = utils.StringToObject(resultJsonStr, findPasswordApplyResult)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return ""
	}
	if "0000" == findPasswordApplyResult.RetCode {
		// 用Debug模式，断点打到这里，Debug的时候把verifyCode设置成手机收到的真实验证码
		// 验证码
		verifyCode := ""
		// 与加密密文对应的随机因子random_key
		randomKey := ""
		// 加密密文
		password := ""

		// 找回密码验证
		verifyParams := find.FindPasswordVerifyParams{
			Timestamp:  utils.GetTimestamp(),
			OidPartner: config.OidPartner,
			UserID:     findPasswordApplyResult.UserID,
			Token:      findPasswordApplyResult.Token,
			/*
			   短信验证码。企业和个体工商户验证注册绑定手机号。
			   个人验证银行预留手机号，未绑卡用户验证注册手机号。
			*/
			VerifyCode: verifyCode,
			RandomKey:  randomKey,
			Password:   password,
		}
		// 测试环境请求地址
		url := "https://accpapi-ste.lianlianpay-inc.com/v1/acctmgr/modify-userinfo-enterprise"
		paramsStr2, err := utils.ObjectToString(verifyParams)
		if err != nil {
			fmt.Println("转换对象失败:", err)
			return ""
		}
		//发送请求
		resultJsonStr2 := client.SendRequest(url, paramsStr2)
		return resultJsonStr2
	}
	return ""
}
