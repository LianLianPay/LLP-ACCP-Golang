package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/security"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/txn/papay"
	"fmt"
)

/**
 * 委托代扣 Demo
 */
func PaymentPapay() string {
	params := papay.PaymentPapayParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.TxnType = "GENERAL_CONSUME"
	params.UserID = "LLianPayTest-In-User-12345"
	params.ClientIP = "127.0.0.1"
	params.RiskItem = `{"frms_ware_category":"4007","goods_name":"测试商品","user_info_mercht_userno":"` + params.UserID + `","user_info_dt_register":"20220823101239","user_info_bind_phone":"13308123456","user_info_full_name":"连连测试","user_info_id_no":"123456789012345678","user_info_identify_state":"0","user_info_identify_type":"4","user_info_id_type":"0","frms_client_chnl":" 16","frms_ip_addr":"127.0.0.1","user_auth_flag":"1"}`

	orderInfo := papay.PaymentPapayOrderInfo{}
	orderInfo.TxnSeqno = "LLianPayTest" + timestamp
	orderInfo.TxnTime = timestamp
	orderInfo.TotalAmount = 1.00
	orderInfo.GoodsName = "西瓜"
	params.OrderInfo = orderInfo

	payerInfo := papay.PaymentPapayPayerInfo{}
	payerInfo.PayerID = "LLianPayTest-In-User-12345"
	payerInfo.PayerType = "USER"
	payerInfo.ContractID = security.LocalEncrypt("2022093000265008")
	params.PayerInfo = payerInfo

	payMethods := papay.PaymentPapayPayMethod{}
	payMethods.Method = "BALANCE"
	payMethods.Amount = 1.00
	params.PayMethods = payMethods

	payeeInfo := papay.PaymentPapayPayeeInfo{}
	payeeInfo.PayeeID = config.OidPartner
	payeeInfo.PayeeType = "MERCHANT"
	payeeInfo.PayeeAmount = "1.00"
	params.PayeeInfo = payeeInfo
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/txn/payment-papay"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
