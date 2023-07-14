package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/txn"
	"fmt"
)

/**
 * 用户委托协议申请/修改 Demo
 */
func PapAgreeApplyAndModify() {
	//apply()
	modify()
}

/**
 * 用户委托协议申请
 */
func apply() string {
	params := txn.PapAgreeApplyAndModifyParams{}

	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.TxnSeqNo = fmt.Sprintf("LLianPayTest%s", timestamp)
	params.TxnTime = timestamp
	params.UserID = "LLianPayTest-In-User-12345"
	params.FlagChnl = "H5"
	params.NotifyURL = "https://test.lianlianpay/notify"
	params.ReturnURL = "https://open.lianlianpay.com?jump=page"

	// 设置签约信息
	papAgreeSignInfo := txn.PapAgreeSignInfo{
		SignStartTime:   "20221001",
		SignInvalidTime: "20231001",
		SingleLimit:     100.00,
		DailyLimit:      1000.00,
		MonthlyLimit:    5000.00,
		AgreementType:   "WITH_HOLD",
	}
	params.PapSignInfo = papAgreeSignInfo

	url := "https://accpgw-ste.lianlianpay-inc.com/v1/txn/pap-agree-apply"
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
 * 用户委托协议修改
 */
func modify() string {
	params := txn.PapAgreeApplyAndModifyParams{}

	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.TxnSeqNo = fmt.Sprintf("LLianPayTest%s", timestamp)
	params.TxnTime = timestamp
	params.UserID = "LLianPayTest-In-User-12345"
	params.PapAgreeNo = "2022093000265008"
	params.FlagChnl = "H5"
	params.NotifyURL = "https://test.lianlianpay/notify"
	params.ReturnURL = "https://open.lianlianpay.com?jump=page"

	url := "https://accpgw-ste.lianlianpay-inc.com/v1/txn/pap-agree-modify"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
