package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/offlinetxn"
	"fmt"
)

/**
 * 电子回单生成 Demo
 */
func ReceiptProduce() string {
	params := offlinetxn.ReceiptProduceParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = "your_oid_partner"
	params.TxnSeqno = "LLianPayTest" + timestamp
	params.TxnTime = timestamp
	params.TradeAccpTxno = "2022092116412310"
	params.TradeBillType = "PAYBILL"
	params.TotalAmount = 202.00
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/offlinetxn/receipt-produce"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
