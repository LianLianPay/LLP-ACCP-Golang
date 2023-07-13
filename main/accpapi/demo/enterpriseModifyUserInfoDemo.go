package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/enterprise"
	"fmt"
)

/**
 * 企业用户信息修改 Demo
 */
func EnterpriseModifyUserInfo() string {
	// 创建 ModifyUserInfoParams 结构体实例
	params := enterprise.ModifyUserInfoParams{}
	// 获取当前时间戳
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	// 设置合作伙伴标识
	params.OidPartner = config.OidPartner
	// 设置用户标识
	params.UserID = "LLianPayTest-En-User-12345"
	// 设置交易流水号
	params.TxnSeqno = "LLianPayTest" + timestamp
	// 设置交易时间
	params.TxnTime = timestamp
	// 交易结果异步通知接收地址，建议使用HTTPS协议
	params.NotifyURL = "https://test.lianlianpay/notify"
	// 基本信息更新
	basicInfo := enterprise.OpenacctApplyBasicInfo{}
	basicInfo.IDExp = "99991231"
	basicInfo.RegEmail = "158@163.com"
	//params.BasicInfo = basicInfo

	// 已有受益所有人信息替换
	uboInfos := enterprise.OpenacctApplyUboInfos{}
	uboInfos.UboName = "测试"
	uboInfos.UboNameEN = "ceshi"
	uboInfos.UboPhone = ""
	uboInfos.IDType = "ID_CARD"
	uboInfos.IDNo = ""
	uboInfos.IDEmblem = ""
	uboInfos.IDPortrait = ""
	uboInfos.IDExp = "20301010"
	uboInfos.Address = "AAA"
	uboInfos.UboType = "ubo"

	list := []enterprise.OpenacctApplyUboInfos{}
	list = append(list, uboInfos)
	params.UboInfos = list

	// 测试环境请求地址
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/acctmgr/modify-userinfo-enterprise"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}
