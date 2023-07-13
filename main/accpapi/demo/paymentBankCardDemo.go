package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/security"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/txn"
	"fmt"
)

/**
 * 银行卡快捷支付 Demo
 */
func PaymentBankCard() string {
	// 调用统一支付创单接口进行创单
	tradeCreateResult := UserTopup()
	// 使用银行卡快捷方式完成支付
	params := txn.PaymentBankCardParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = tradeCreateResult.OidPartner
	params.TxnSeqNo = tradeCreateResult.TxnSeqNo
	params.TotalAmount = tradeCreateResult.TotalAmount
	params.RiskItem = `{"frmsWareCategory":"4007","goodsName":"西瓜","userInfoMerchtUserno":"LLianPayTest-In-User-12345","userInfoDtRegister":"20220823101239","userInfoBindPhone":"13208123456","userInfoFullName":"连连测试","userInfoIdNo":"","userInfoIdentifyState":"0","userInfoIdentifyType":"4","userInfoIdType":"0","frmsClientChnl":" H5","frmsIpAddr":"127.0.0.1","userAuthFlag":"1"}`

	// 设置付款方信息
	payerInfo := txn.PaymentPayerInfo{}
	payerInfo.UserID = tradeCreateResult.UserID
	// 用户：LLianPayTest-In-User-12345 密码：qwerty，本地测试环境测试，没接入密码控件，使用本地加密方法加密密码（仅限测试环境使用）
	payerInfo.Password = security.LocalEncrypt("qwerty")
	params.PayerInfo = payerInfo

	// 设置付款银行卡信息
	bankCardInfo := txn.PaymentBankCardInfo{}
	// 使用该用户的绑卡协议号
	bankCardInfo.LinkedAgrtno = "2022081900364011"
	params.BankCardInfo = bankCardInfo

	// 设置付款方式信息，支持组合支付，传入数组
	payMethod := txn.PaymentBankCardPayMethods{}
	// 协议支付借记卡
	payMethod.Method = "AGRT_DEBIT_CARD"
	payMethod.Amount = tradeCreateResult.TotalAmount
	params.PayMethods = []txn.PaymentBankCardPayMethods{payMethod}
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/txn/payment-bankcard"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)

	var bankCardResult txn.PaymentBankCardResult
	err = utils.StringToObject(resultJsonStr, bankCardResult)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return ""
	}

	// 小额免验，不需要验证码，直接返回0000
	if "0000" == bankCardResult.RetCode {
		fmt.Println("支付成功！！！")
		return ""
	}
	// 需要输入短信验证码，调用交易二次短信验证接口
	// 用Debug模式，断点打到这里，Debug的时候把verifyCode设置成手机收到的真实验证码
	verifyCode := ""
	if "8888" == bankCardResult.RetCode {
		validationSmsParams := txn.ValidationSmsParams{}
		validationSmsParams.Timestamp = utils.GetTimestamp()
		validationSmsParams.OidPartner = bankCardResult.OidPartner
		validationSmsParams.PayerID = bankCardResult.UserID
		validationSmsParams.PayerType = "USER"
		validationSmsParams.TxnSeqNo = bankCardResult.TxnSeqNo
		validationSmsParams.TotalAmount = fmt.Sprintf("", bankCardResult.TotalAmount)
		validationSmsParams.Token = bankCardResult.Token
		validationSmsParams.VerifyCode = verifyCode

		// 测试环境URL
		url := "https://accpapi-ste.lianlianpay-inc.com/v1/txn/validation-sms"
		paramsStr, err := utils.ObjectToString(params)
		if err != nil {
			fmt.Println("转换对象失败:", err)
			return ""
		}

		//发送请求
		resultJsonStr := client.SendRequest(url, paramsStr)
		return resultJsonStr
	}
	return ""
}
