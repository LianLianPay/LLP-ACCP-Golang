package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/virtualno"
	"fmt"
	"time"
)

/**
 * 虚拟卡申请 Demo
 */
func VirtualNoApplyDemo() string {
	params := virtualno.VirtualNoApplyParams{}
	timestamp := time.Now().Format("20060102150405")
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.TxnSeqno = "LLianPayTest" + timestamp
	params.TxnTime = timestamp

	// 虚拟卡参数
	virtualInfo := virtualno.VirtualInfo{}
	/*
		所属客户类型。
		用户：USER
		平台商户：MERCHANT
	*/
	virtualInfo.CustType = "USER"
	// 申请虚拟卡的请求方ID，若为消费类虚拟卡，则必填。
	virtualInfo.UserID = "LLianPayTest-In-User-12345"
	virtualInfo.ApplyUserID = "LLianPayTest-In-User-12345-afafafd"
	// 申请虚拟卡的客户姓名。会尝试检查通过虚拟卡转账的他行卡户名是否一致。
	virtualInfo.ApplyUsername = "个人开户测试"
	/*
		申请虚拟卡场景。消费：consume
		现金上缴：deposit_cash
		费用上缴：deposit_fee
		其他：other
		（payee_id+applay_userid+apply_sence只能一张卡）
	*/
	virtualInfo.ApplySence = "consume"
	/*
		虚拟卡业务类型。
		消费业务：OFFLINE_CONSUME
	*/
	virtualInfo.BusinessType = "OFFLINE_CONSUME"
	// 虚拟卡入账通知URL
	virtualInfo.BBocNotify = "https://test.lianlianpay/notify"
	params.VirtualInfo = virtualInfo

	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/acctmgr/virtualno-apply"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
