package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/webapge"
	"fmt"
	"time"
)

/**
 * 用户开户申请(页面接入) Demo
 */
func H5Openacct() {
	enterprise()
	innerUser()
}

/**
 * 企业用户（页面）开户 Demo
 */
func enterprise() string {
	params := webapge.OpenacctApplyParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.UserID = "LLianPayTest-En-User-12345"
	params.TxnSeqno = "LLianPayTest" + timestamp
	params.TxnTime = timestamp
	/*
	   交易发起渠道。
	   ANDROID
	   IOS
	   H5
	   PC
	*/
	params.FlagChnl = "H5"
	// 交易完成回跳页面地址，H5/PC渠道必传。
	params.ReturnURL = "https://open.lianlianpay.com?jump=page"
	// 交易结果异步通知接收地址，建议HTTPS协议。
	params.NotifyURL = "https://test.lianlianpay/notify"
	/*
	   用户类型。
	   INNERUSER：个人用户
	   INNERCOMPANY：企业用户
	*/
	params.UserType = "INNERCOMPANY"
	/*
	   开户类型。
	   OpenActivatedUser：待激活开户
	   OpenNormalUser：正常开户，默认为正常开户
	*/
	params.CustTradeSerialType = "OpenNormalUser"

	// 设置开户账户申请信息
	accountInfo := webapge.OpenacctApplyAccountInfo{}
	/*
	   个人支付账户  PERSONAL_PAYMENT_ACCOUNT
	   企业支付账户  ENTERPRISE_PAYMENT_ACCOUNT
	*/
	accountInfo.AccountType = "ENTERPRISE_PAYMENT_ACCOUNT"
	//accountInfo.AccountNeedLevel = "V3"
	params.AccountInfo = accountInfo

	// 测试环境请求地址
	url := "https://accpgw-ste.lianlianpay-inc.com/v1/acctmgr/openacct-apply"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}

/**
 * 个人用户（页面）开户 Demo
 */
func innerUser() string {
	params := webapge.OpenacctApplyParams{}
	timestamp := time.Now().Format("20060102150405")
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.UserID = "LLianPayTest-In-User-12345"
	params.TxnSeqno = "LLianPayTest" + timestamp
	params.TxnTime = timestamp
	/*
		交易发起渠道。
		ANDROID
		IOS
		H5
		PC
	*/
	params.FlagChnl = "H5"
	// 交易完成回跳页面地址，H5/PC渠道必传。
	params.ReturnURL = "https://open.lianlianpay.com?jump=page"
	// 交易结果异步通知接收地址，建议HTTPS协议。
	params.NotifyURL = "https://test.lianlianpay/notify"
	/*
		用户类型。
		INNERUSER：个人用户
		INNERCOMPANY：企业用户
	*/
	params.UserType = "INNERUSER"
	/*
		开户类型。
		OpenActivatedUser：待激活开户
		OpenNormalUser：正常开户，默认为正常开户
	*/
	params.CustTradeSerialType = "OpenActivatedUser"

	// 设置开户账户申请信息
	accountInfo := webapge.OpenacctApplyAccountInfo{}
	/*
		个人支付账户  PERSONAL_PAYMENT_ACCOUNT
		企业支付账户  ENTERPRISE_PAYMENT_ACCOUNT
	*/
	accountInfo.AccountType = "PERSONAL_PAYMENT_ACCOUNT"
	accountInfo.AccountNeedLevel = "V3"
	params.AccountInfo = accountInfo

	// 测试环境请求地址
	url := "https://accpgw-ste.lianlianpay-inc.com/v1/acctmgr/openacct-apply"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
