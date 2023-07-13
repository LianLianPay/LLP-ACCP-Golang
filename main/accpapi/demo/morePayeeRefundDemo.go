package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/txn/refund"
	"fmt"
)

/**
 * 退款申请 Demo
 */
func MorePayeeRefund() string {
	params := refund.MorePayeeRefundParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.UserID = "LLianPayTest-In-User-12345abdfasjdlf"

	originalOrderInfo := refund.OriginalOrderInfo{}
	originalOrderInfo.TxnSeqNo = "LLianPayTest20221107113358"
	originalOrderInfo.TotalAmount = 150.0
	params.OriginalOrderInfo = originalOrderInfo

	refundOrderInfo := refund.RefundOrderInfo{}
	refundOrderInfo.RefundSeqNo = "LLianPayTest" + timestamp
	refundOrderInfo.RefundTime = timestamp
	refundOrderInfo.RefundAmount = 1.0
	params.RefundOrderInfo = refundOrderInfo

	pyeeRefundInfo := refund.PyeeRefundInfo{}
	pyeeRefundInfo.PayeeID = "LLianPayTest-En-User-12345"
	pyeeRefundInfo.PayeeType = "USER"
	pyeeRefundInfo.PayeeAcctType = "USEROWN"
	pyeeRefundInfo.PayeeRefundAmount = 1.0
	pyeeRefundInfo.IsAdvancePay = "N"
	params.PyeeRefundInfos = []refund.PyeeRefundInfo{pyeeRefundInfo}

	refundMethod := refund.RefundMethod{}
	refundMethod.Method = "EBANK_B2B"
	refundMethod.Amount = 1.0
	params.RefundMethods = []refund.RefundMethod{refundMethod}

	url := "https://accpapi-ste.lianlianpay-inc.com/v1/txn/more-payee-refund"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
