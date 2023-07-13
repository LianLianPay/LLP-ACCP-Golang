package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/security"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/txn/transfer"
	"fmt"
)

func TransferMorepyee() string {
	params := transfer.TransferMorepyeeParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner

	// 商户订单信息
	orderInfo := transfer.TransferMorepyeeOrderInfo{}
	orderInfo.TxnSeqNo = "LLianPayTest" + timestamp
	orderInfo.TxnTime = timestamp
	orderInfo.TotalAmount = 300.00
	orderInfo.TxnPurpose = "其他"
	params.OrderInfo = orderInfo

	// 付款方信息
	payerInfo := transfer.TransferMorepyeePayerInfo{}
	/*
	   付款方类型。
	   用户：USER
	   平台商户：MERCHANT
	*/
	payerInfo.PayerType = "USER"
	payerInfo.PayerID = "20220930162843205648385"
	/*
	   委托代发协议号。账户+返回的的代扣协议号，委托代发时必输。该字段需要RSA公钥加密传输。
	   通过用户委托协议签约接口获取
	*/
	payerInfo.PapAgreeNo = security.LocalEncrypt("2022101800274008")
	params.PayerInfo = payerInfo

	// 收款方信息
	payeeInfo := transfer.TransferMorepyeePayeeInfo{}
	payeeInfo.PayeeType = "USER"
	// 收款方标识，收款方为用户时，为用户user_id，收款方为平台商户时，取平台商户号
	payeeInfo.PayeeID = "LLianPayTest-En-User-12345"
	payeeInfo.PayeeAmount = "300.00"
	params.PayeeInfo = []transfer.TransferMorepyeePayeeInfo{payeeInfo}

	// 测试风控参数
	params.RiskItem = "{\"frms_ware_category\":\"4007\",\"goods_name\":\"测试商品\",\"user_info_mercht_userno\":\"" + payerInfo.PayerID + "\",\"user_info_dt_register\":\"20220823101239\",\"user_info_bind_phone\":\"13308123456\",\"user_info_full_name\":\"连连测试\",\"user_info_id_no\":\"123456789012345678\",\"user_info_identify_state\":\"0\",\"user_info_identify_type\":\"4\",\"user_info_id_type\":\"0\",\"frms_client_chnl\":\" 16\",\"frms_ip_addr\":\"127.0.0.1\",\"user_auth_flag\":\"1\"}"
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/txn/transfer-morepyee"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
