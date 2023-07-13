package demo

import (
	"LLP-ACCP-Go/main/accpapi/client"
	"LLP-ACCP-Go/main/accpapi/config"
	"LLP-ACCP-Go/main/accpapi/utils"
	"LLP-ACCP-Go/main/accpapi/v1/offlinetxn"
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/**
 * 电子回单下载 Demo
 */
func ReceiptDownload() {
	params := offlinetxn.ReceiptDownloadParams{
		Timestamp:       utils.GetTimestamp(),
		OidPartner:      config.OidPartner,
		Token:           "6a456fbd3c6ffaecccdc3cadba1a3621",
		ReceiptAccpTxno: "2022092116412310",
	}
	// 测试环境URL
	url := "https://accpapi-ste.lianlianpay-inc.com/v1/offlinetxn/receipt-download"
	paramsStr, err := utils.ObjectToString(params)
	if err != nil {
		fmt.Println("转换对象失败:", err)
		return
	}

	//发送请求
	resultJsonStr := client.SendRequest(url, paramsStr)

	var receiptDownloadResult offlinetxn.ReceiptDownloadResult
	err = utils.StringToObject(resultJsonStr, receiptDownloadResult)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return
	}
	fmt.Println(receiptDownloadResult)
	base64ToFile(receiptDownloadResult.ReceiptSumFile, "D:\\Receipts\\")
}

/**
 * 原base64压缩文件转文件
 *
 * @param base64 原Base64压缩文件
 * @param path   文件保存路径
 * @throws RuntimeException
 */
func base64ToFile(base64Data string, outputPath string) error {
	fmt.Println("解压文件地址" + outputPath)

	byteBase64, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return err
	}

	byteArray := bytes.NewReader(byteBase64)
	zipReader, err := zip.NewReader(byteArray, int64(len(byteBase64)))
	if err != nil {
		return err
	}

	for _, file := range zipReader.File {
		fmt.Println("文件名称:", file.Name)

		path := filepath.Join(outputPath, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
			continue
		}

		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}

		outputFile, err := os.Create(path)
		if err != nil {
			return err
		}

		zipFile, err := file.Open()
		if err != nil {
			outputFile.Close()
			return err
		}

		_, err = io.Copy(outputFile, zipFile)
		if err != nil {
			outputFile.Close()
			zipFile.Close()
			return err
		}

		outputFile.Close()
		zipFile.Close()
	}

	return nil
}
