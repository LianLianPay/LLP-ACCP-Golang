package client

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
)

func Handle() {
	// 加载根证书
	caCert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatalf("Failed to load CA certificate: %v", err)
	}

	// 创建自定义的TLS配置
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true, // 跳过证书验证
		RootCAs:            x509.NewCertPool(),
	}

	// 将根证书添加到根证书池中
	if !tlsConfig.RootCAs.AppendCertsFromPEM(caCert) {
		log.Fatalf("Failed to append CA certificate")
	}

	// 创建自定义的TCP连接
	dialer := &net.Dialer{
		Timeout:   30,
		KeepAlive: 30,
	}

	conn, err := tls.DialWithDialer(dialer, "tcp", "example.com:443", tlsConfig)
	if err != nil {
		log.Fatalf("Failed to establish TLS connection: %v", err)
	}
	defer conn.Close()

	// 使用连接进行后续操作
	// ...

	log.Println("TLS connection established successfully")
}
