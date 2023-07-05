package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/txn"
	"fmt"
)

func TradeCreate() {
	// 用户充值
	//userTopup();
	// 普通消费
	result := generalConsume()
	fmt.Println(result)
}

// 普通消费
func generalConsume() string {
	params := txn.TradeCreateParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	//params.Timestamp = "20230704145205"
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

	// 设置商户订单信息
	orderInfo := txn.TradeCreateOrderInfo{
		TxnSeqno:    "LLianPayTest" + timestamp,
		TxnTime:     timestamp,
		TotalAmount: 150.00,
		GoodsName:   "西瓜",
	}
	params.OrderInfo = orderInfo

	// 设置收款方信息，消费分账场景支持上送多收款方（最多10个），收款总金额必须和订单总金额相等
	mPayeeInfo := txn.TradeCreatePayeeInfo{
		PayeeID:     config.OidPartner,
		PayeeType:   "MERCHANT",
		PayeeAmount: "50.00",
		PayeeMemo:   "分账",
	}

	uPayeeInfo := txn.TradeCreatePayeeInfo{
		PayeeID:     "LLianPayTest-En-User-12345",
		PayeeType:   "USER",
		PayeeAmount: "100.00",
		PayeeMemo:   "分账",
	}
	params.PayeeInfo = []txn.TradeCreatePayeeInfo{mPayeeInfo, uPayeeInfo}

	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/txn/tradecreate"
	url = "http://localhost:8086/accpapi/v1/txn/tradecreate"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)

	var tradeCreateResult txn.TradeCreateResult
	err = utils.StringToObject(resultJsonStr, tradeCreateResult)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return ""
	}
	return resultJsonStr
}
