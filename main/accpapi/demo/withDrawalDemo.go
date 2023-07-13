package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/security"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/txn"
	"fmt"
	"time"
)

/**
 * 提现申请 Demo
 */
func WithDrawal() string {
	params := txn.WithDrawalParams{}
	timestamp := time.Now().Format("20060102150405")
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.NotifyURL = "https://test.lianlianpay/notify"
	params.LinkedAgrtNo = "2022081900364011"

	// 设置商户订单信息
	orderInfo := txn.WithDrawalOrderInfo{}
	orderInfo.TxnSeqNo = "LLianPayTest" + timestamp
	orderInfo.TxnTime = timestamp
	orderInfo.TotalAmount = 150.00
	orderInfo.Postscript = "用户提现"
	params.OrderInfo = orderInfo

	// 设置付款方信息
	payerInfo := txn.WithDrawalPayerInfo{}
	payerInfo.PayerType = "USER"
	payerInfo.PayerID = "LLianPayTest-In-User-12345"
	// 用户：LLianPayTest-In-User-12345 密码：qwerty，本地测试环境测试，没接入密码控件，使用本地加密方法加密密码（仅限测试环境使用）
	payerInfo.Password = security.LocalEncrypt("qwerty")
	params.PayerInfo = payerInfo

	params.RiskItem = fmt.Sprintf("{\"frms_ware_category\":\"4007\",\"goods_name\":\"测试商品\",\"user_info_mercht_userno\":\"%s\",\"user_info_dt_register\":\"20220823101239\",\"user_info_bind_phone\":\"13308123456\",\"user_info_full_name\":\"连连测试\",\"user_info_id_no\":\"123456789012345678\",\"user_info_identify_state\":\"0\",\"user_info_identify_type\":\"4\",\"user_info_id_type\":\"0\",\"frms_client_chnl\":\" 16\",\"frms_ip_addr\":\"127.0.0.1\",\"user_auth_flag\":\"1\"}", payerInfo.PayerID)

	// 测试环境URL
	url := "https://accpgw-ste.lianlianpay-inc.com/v1/acctmgr/apply-password-element"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)

	var withDrawalResult txn.WithDrawalResult
	err = utils.StringToObject(resultJsonStr, withDrawalResult)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return ""
	}
	if "0000" == withDrawalResult.RetCode {
		fmt.Println("支付成功！！！")
		return ""
	}
	// 需要输入短信验证码，调用交易二次短信验证接口
	// 用Debug模式，断点打到这里，Debug的时候把verifyCode设置成手机收到的真实验证码
	verifyCode := ""
	if "8888" == withDrawalResult.RetCode {
		validationSmsParams := txn.ValidationSmsParams{}
		validationSmsParams.Timestamp = utils.GetTimestamp()
		validationSmsParams.OidPartner = withDrawalResult.OidPartner
		validationSmsParams.PayerID = withDrawalResult.UserID
		validationSmsParams.PayerType = "USER"
		validationSmsParams.TxnSeqNo = withDrawalResult.TxnSeqNo
		validationSmsParams.TotalAmount = fmt.Sprintf("%f", withDrawalResult.TotalAmount)
		validationSmsParams.Token = withDrawalResult.Token
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
