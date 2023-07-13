package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/txn/secured"
	"fmt"
)

/**
 * 担保确认 Demo
 */
func SecuredConfirm() string {
	tradeCreateResult := securedConsumePay()

	timestamp := utils.GetTimestamp()
	params := secured.SecuredConfirmParams{}
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.UserID = tradeCreateResult.UserID
	params.ConfirmMode = "PART"

	originalOrderInfo := secured.SecuredConfirmOriginalOrderInfo{}
	originalOrderInfo.TxnSeqNo = tradeCreateResult.TxnSeqNo
	originalOrderInfo.TotalAmount = tradeCreateResult.TotalAmount
	params.OriginalOrderInfo = originalOrderInfo

	orderInfo := secured.SecuredConfirmOrderInfo{}
	orderInfo.ConfirmSeqNo = "LLianPayTest" + string(timestamp)
	orderInfo.ConfirmTime = timestamp
	orderInfo.ConfirmAmount = tradeCreateResult.TotalAmount
	params.ConfirmOrderInfo = orderInfo

	payeeInfo := secured.SecuredConfirmPayeeInfo{}
	payeeInfo.PayeeID = config.OidPartner
	payeeInfo.PayeeType = "MERCHANT"
	payeeInfo.PayeeAmount = "10.00"

	payeeInfo1 := secured.SecuredConfirmPayeeInfo{}
	payeeInfo1.PayeeID = "LLianPayTest-En-User-12345"
	payeeInfo1.PayeeType = "USER"
	payeeInfo1.PayeeAmount = "20.00"

	payeeInfo2 := secured.SecuredConfirmPayeeInfo{}
	payeeInfo2.PayeeID = "ll627544175777419265"
	payeeInfo2.PayeeType = "USER"
	payeeInfo2.PayeeAmount = "20.00"

	params.PayeeInfo = []secured.SecuredConfirmPayeeInfo{payeeInfo, payeeInfo1, payeeInfo2}
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/txn/secured-confirm"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
