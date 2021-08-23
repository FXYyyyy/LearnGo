package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	requestHttps()
}

func requestHttps()  {
	req, err := http.NewRequest("GET", "https://127.0.0.1:8443/httpstest?name=你是猪", nil)
	if err != nil {
		fmt.Printf("请求初始化失败: %v", err)
		return
	}

	//设置跳过不安全的https
	tlsllTransport := &http.Transport{
		MaxIdleConnsPerHost: 10,
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS11,
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: tlsllTransport,
	}

	resq,err := client.Do(req)
	if err != nil{
		fmt.Printf("客户端发起请求失败: %v", err)
		return
	}

	defer resq.Body.Close()
	io.Copy(os.Stdout, resq.Body)
}
