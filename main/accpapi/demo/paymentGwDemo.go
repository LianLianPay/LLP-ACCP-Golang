package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/txn"
	"fmt"
)

/**
 * 网关类支付 Demo
 */
func PaymentGw() string {
	tradeCreateResult := generalConsume()
	timestamp := utils.GetTimestamp()
	params := txn.PaymentGwParams{
		Timestamp:   timestamp,
		OidPartner:  config.OidPartner,
		TxnSeqno:    tradeCreateResult.TxnSeqNo,
		TotalAmount: tradeCreateResult.TotalAmount,
		RiskItem:    `{"frms_ware_category":"4007","goods_name":"测试商品","user_info_mercht_userno":"` + tradeCreateResult.UserID + `","user_info_dt_register":"20220823101239","user_info_bind_phone":"13308123456","user_info_full_name":"连连测试","user_info_id_no":"123456789012345678","user_info_identify_state":"0","user_info_identify_type":"4","user_info_id_type":"0","frms_client_chnl":" 16","frms_ip_addr":"127.0.0.1","user_auth_flag":"1"}`,
		ClientIP:    "127.0.0.1",
	}

	payerInfo := txn.PaymentPayerInfo{
		PayerType: "USER",
		PayerID:   tradeCreateResult.UserID,
	}
	params.PayerInfo = payerInfo

	payMethods := txn.PaymentBankCardPayMethods{
		Method: "EBANK_B2B",
		Amount: tradeCreateResult.TotalAmount,
	}
	params.PayMethods = []txn.PaymentBankCardPayMethods{payMethods}
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/txn/payment-gw"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
