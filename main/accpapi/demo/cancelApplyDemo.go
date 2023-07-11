package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/security"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/canncel"
	"fmt"
	"time"
)

func CancelApply() string {
	params := canncel.CancelApplyParams{
		Timestamp:  time.Now().Format("20060102150405"),
		OidPartner: config.OidPartner,
		UserID:     "LLianPayTest-Api-User-12345",
		TxnSeqno:   "LLianPayTest" + time.Now().Format("20060102150405"),
		TxnTime:    time.Now().Format("20060102150405"),
		NotifyURL:  "https://test.lianlianpay/notify",
		UserName:   "API开户测试",
		IdType:     "ID_CARD",
		IdNo:       "",
		RegPhone:   "",
		Password:   security.LocalEncrypt("qwerty"),
	}

	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/acctmgr/cancel-apply"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)

	var cancelApplyParams canncel.CancelApplyParams
	err = utils.StringToObject(resultJsonStr, cancelApplyParams)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return ""
	}
	return resultJsonStr
}
