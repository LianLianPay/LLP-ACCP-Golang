package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/documents"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func documentsUpload() string {
	// 创建 UploadParams 结构体实例
	params := documents.UploadParams{}
	// 获取当前时间戳
	timestamp := utils.GetTimestamp()
	params.Timestamp = timestamp
	// 设置合作伙伴标识
	params.OidPartner = config.OidPartner
	// 设置用户标识
	params.UserID = "LLianPayTest-In-User-12345"
	// 设置交易流水号
	params.TxnSeqno = "LLianPayTest" + timestamp
	// 设置交易时间
	params.TxnTime = timestamp
	// 文件类型，支持bmp、png、jpeg、jpg、gif
	params.FileType = "jpeg"
	// 上下文类型
	params.ContextType = "USER_IMAGE"
	// 文件名
	params.FileName = "照片"
	// 文件内容
	imageBase64, err := getImageBase64()
	if err != nil {
		fmt.Println("getImageBase64 fail:", err)
		return ""
	}
	params.FileContext = imageBase64

	// 测试环境请求地址
	url := "https://accpfile-ste.lianlianpay-inc.com/v1/documents/upload"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return ""
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)
	return resultJsonStr
}

/**
 * 读取resource下的一张示例图片，转化成BASE64字符串，不需要加前缀data:image/png;base64
 */
func getImageBase64() (string, error) {
	file, err := os.Open("documents-upload-example.jpeg")
	if err != nil {
		return "", err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	encodedString := base64.StdEncoding.EncodeToString(data)
	return encodedString, nil
}
