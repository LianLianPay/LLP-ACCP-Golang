package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/security"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/txn"
	"fmt"
	"strconv"
)

/**
 * 余额支付 Demo
 */
func PaymentBalance() {
	generalConsumePay()
}

/**
 * 普通消费场景的余额支付
 */
func generalConsumePay() {
	//调用统一支付创单接口进行创单
	tradeCreateResult := generalConsume()
	//使用余额方式完成支付
	pay(tradeCreateResult)
}

/**
 * 普通消费场景的余额支付
 */
func securedConsumePay() txn.TradeCreateResult {
	tradeCreateResult := securedConsume()
	pay(tradeCreateResult)
	return tradeCreateResult
}

// 余额支付
func pay(tradeCreateResult txn.TradeCreateResult) {
	var params txn.PaymentBalanceParams
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = tradeCreateResult.OidPartner
	params.TxnSeqNo = tradeCreateResult.TxnSeqNo
	params.TotalAmount = tradeCreateResult.TotalAmount
	params.RiskItem = "{\"frms_ware_category\":\"4007\",\"goods_name\":\"测试商品\",\"user_info_mercht_userno\":\"" + tradeCreateResult.UserID + "\",\"user_info_dt_register\":\"20220823101239\",\"user_info_bind_phone\":\"13308123456\",\"user_info_full_name\":\"连连测试\",\"user_info_id_no\":\"123456789012345678\",\"user_info_identify_state\":\"0\",\"user_info_identify_type\":\"4\",\"user_info_id_type\":\"0\",\"frms_client_chnl\":\" 16\",\"frms_ip_addr\":\"127.0.0.1\",\"user_auth_flag\":\"1\"}"

	// 设置付款方信息
	var payerInfo txn.PaymentPayerInfo
	payerInfo.UserID = tradeCreateResult.UserID
	// 用户：LLianPayTest-In-User-12345 密码：qwerty，本地测试环境测试，没接入密码控件，使用本地加密方法加密密码（仅限测试环境使用）
	payerInfo.Password = security.LocalEncrypt("qwerty")
	params.PayerInfo = payerInfo
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/txn/payment-balance"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)

	var bankCardResult txn.PaymentResult
	err = utils.StringToObject(resultJsonStr, bankCardResult)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return
	}
	if "0000" == bankCardResult.RetCode {
		fmt.Println("支付成功！！！")
		return
	}
	verifyCode := ""
	if "8888" == bankCardResult.RetCode {
		validationSmsParams := txn.ValidationSmsParams{}
		validationSmsParams.Timestamp = utils.GetTimestamp()
		validationSmsParams.OidPartner = bankCardResult.OidPartner
		validationSmsParams.PayerID = bankCardResult.UserID
		validationSmsParams.PayerType = "USER"
		validationSmsParams.TxnSeqNo = bankCardResult.TxnSeqNo
		validationSmsParams.TotalAmount = strconv.FormatFloat(bankCardResult.TotalAmount, 'f', -1, 64)
		validationSmsParams.Token = bankCardResult.Token
		validationSmsParams.VerifyCode = verifyCode
		// 测试环境URL
		url := "https://accpapi-ste.lianlianpay-inc.com/v1/txn/validation-sms"
		paramsStr, err := utils.ObjectToString(params)
		if err != nil {
			fmt.Println("转换对象失败:", err)
			return
		}

		//发送请求
		resultJsonStr := client.SendRequest(url, paramsStr)
		fmt.Println(resultJsonStr)
		return
	}
}
