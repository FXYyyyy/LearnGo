package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	//testGet()
	testPostForm()
}

/**
* 测试get请求
* testGet
* @Description: 
*/
func testGet()  {
	resp, err := http.Get("https://xueyuanjun.com")
	if err != nil {
		fmt.Printf("发起请求失败：%v", err)
		return
	}

	defer resp.Body.Close()	//结束网络请求，释放资源
	//io.Copy(os.Stdout, resp.Body)	//获取实体
	fmt.Println(resp.StatusCode)	//获取状态码
	fmt.Println(resp.Header)		//获取响应头
}

/**
* testPostForm
* @Description: 测试post表单请求，需要用url。
*/
func testPostForm()  {
	resp, err := http.PostForm("https://xueyuanjun.com/login", url.Values{"name":{"学院君"}, "password":{"test-passwd"}})

	if err != nil {
		fmt.Printf("请求失败：%v", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("请求失败状态码：", resp.StatusCode)
		return
	}

	fmt.Println("状态吗：", resp.StatusCode)
	io.Copy(os.Stdout, resp.Body)
}
