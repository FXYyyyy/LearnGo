package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

/**
* main
* @Description: 一个简单的客户端
*/
func main() {
	//testByClientDo()
	//testByGet()
	testByPostForm()
}

/**
* testByClientDo
* @Description: 通过自定义的客户端请求
*/
func testByClientDo()  {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/hello?name=学院君", nil)
	if err != nil {
		fmt.Printf("初始化请求失败:%v", err)
		return
	}

	req.Header.Add("Custom-Header", "Custom-Value")

	client := &http.Client{}
	resq, err := client.Do(req)
	if err != nil {
		fmt.Printf("客户端发起请求失败: %v", err)
		return
	}

	defer resq.Body.Close()
	io.Copy(os.Stdout, resq.Body)
}

/**
* testByGet
* @Description: 通过get方法向server发起请求
*/
func testByGet()  {
	resq, err := http.Get("http://127.0.0.1:8080/hello?name=你是猪")
	if err != nil {
		fmt.Printf("客户端发起请求失败: %v", err)
		return
	}
	defer resq.Body.Close()

	fmt.Println("statusCode = ", resq.StatusCode)
	io.Copy(os.Stdout, resq.Body)
}

func testByPostForm()  {
	resq, err := http.PostForm("http://127.0.0.1:8080/hello",  url.Values{"name":{"学院君"}, "password":{"test-passwd"}})
	if err != nil {
		fmt.Printf("客户端发起请求失败: %v", err)
		return
	}

	defer resq.Body.Close()

	fmt.Println("statusCode = ", resq.StatusCode)
	io.Copy(os.Stdout, resq.Body)
}
