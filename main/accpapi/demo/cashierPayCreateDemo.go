package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/cashier"
	"fmt"
)

/**
 * 账户+收银台 Demo
 */
func CashierPayCreate() {
	// 普通消费
	generalConsume()

	// 用户充值
	//userTopup()
}

// 用户充值
func userTopup() string {
	params := cashier.CashierPayCreateParams{}

	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	// 用户充值
	params.TxnType = "USER_TOPUP"
	params.UserID = "LLianPayTest-In-User-12345"
	/*
	   用户类型。默认：注册用户。
	   注册用户：REGISTERED
	   匿名用户：ANONYMOUS
	*/
	params.UserType = "REGISTERED"
	params.NotifyURL = "https://test.lianlianpay/notify"
	params.ReturnURL = "https://open.lianlianpay.com?jump=page"
	// 交易发起渠道设置
	params.FlagChnl = "H5"
	// 测试风控参数
	params.RiskItem = `{"frms_ware_category":"4007","goods_name":"测试商品","user_info_mercht_userno":"` + params.UserID + `","user_info_dt_register":"20220823101239","user_info_bind_phone":"13308123456","user_info_full_name":"连连测试","user_info_id_no":"123456789012345678","user_info_identify_state":"0","user_info_identify_type":"4","user_info_id_type":"0","frms_client_chnl":"16","frms_ip_addr":"127.0.0.1","user_auth_flag":"1"}`

	// 设置商户订单信息
	orderInfo := cashier.CashierPayCreateOrderInfo{}
	orderInfo.TxnSeqno = "LLianPayTest" + timestamp
	orderInfo.TxnTime = timestamp
	orderInfo.TotalAmount = 100.00
	orderInfo.GoodsName = "用户充值"
	params.OrderInfo = orderInfo

	// 设置付款方信息
	payerInfo := cashier.CashierPayCreatePayerInfo{}
	payerInfo.PayerID = "LLianPayTest-In-User-12345"
	payerInfo.PayerType = "USER"
	params.PayerInfo = payerInfo

	// 测试环境请求地址
	url := "https://accpgw-ste.lianlianpay-inc.com/v1/cashier/paycreate"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}

// 普通消费
func generalConsume1() string {
	params := cashier.CashierPayCreateParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	// 普通消费
	params.TxnType = "GENERAL_CONSUME"
	params.UserID = "LLianPayTest-In-User-12345"
	/*
	   用户类型。默认：注册用户。
	   注册用户：REGISTERED
	   匿名用户：ANONYMOUS
	*/
	params.UserType = "REGISTERED"
	params.NotifyURL = "https://test.lianlianpay/notify"
	params.ReturnURL = "https://open.lianlianpay.com?jump=page"
	// 交易发起渠道设置
	params.FlagChnl = "H5"
	// 测试风控参数
	params.RiskItem = `{"frms_ware_category":"4007","goods_name":"测试商品","user_info_mercht_userno":" + params.UserID + ","user_info_dt_register":"20220823101239","user_info_bind_phone":"13308123456","user_info_full_name":"连连测试","user_info_id_no":"123456789012345678","user_info_identify_state":"0","user_info_identify_type":"4","user_info_id_type":"0","frms_client_chnl":" 16","frms_ip_addr":"127.0.0.1","user_auth_flag":"1"}`

	// 设置商户订单信息
	orderInfo := cashier.CashierPayCreateOrderInfo{
		TxnSeqno:    "LLianPayTest" + timestamp,
		TxnTime:     timestamp,
		TotalAmount: 1.00,
		GoodsName:   "西瓜",
	}
	params.OrderInfo = orderInfo

	// 设置收款方信息，消费分账场景支持上送多收款方（最多10个），收款总金额必须和订单总金额相等
	mPayeeInfo := cashier.CashierPayCreatePayeeInfo{
		PayeeID:     config.OidPartner,
		PayeeType:   "MERCHANT",
		PayeeAmount: "0.5",
		PayeeMemo:   "分账",
	}

	uPayeeInfo := cashier.CashierPayCreatePayeeInfo{
		PayeeID:     "LLianPayTest-En-User-12345",
		PayeeType:   "USER",
		PayeeAmount: "0.5",
		PayeeMemo:   "分账",
	}

	payeeInfo := []cashier.CashierPayCreatePayeeInfo{mPayeeInfo, uPayeeInfo}
	params.PayeeInfo = payeeInfo

	// 设置付款方信息
	payerInfo := cashier.CashierPayCreatePayerInfo{
		PayerID:   "LLianPayTest-In-User-12345",
		PayerType: "USER",
	}
	params.PayerInfo = payerInfo

	// 测试环境请求地址
	url := "https://accpgw-ste.lianlianpay-inc.com/v1/cashier/paycreate"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
