package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/acctmge/random"
	"fmt"
)

/**
 * 随机因子获取 Demo
 */
func GetRandom() string {
	params := random.GetRandomParams{}
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	params.OidPartner = config.OidPartner
	params.UserID = "LLianPayTest-In-User-12345"
	/*
	   交易发起渠道。
	   ANDROID
	   IOS
	   H5
	   PC
	*/
	params.FlagChnl = "H5"
	// 测试环境都传test，正式环境传真实域名/包名
	params.PkgName = "test"
	// 测试环境都传test，正式环境传真实域名/应用名
	params.AppName = "test"
	params.EncryptAlgorithm = "RSA"

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
