package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/**
* main
* @Description: 一个简单的server端
*/
func main() {
	http.HandleFunc("/hello", HandleRequest)

	err := http.ListenAndServe(":8080", nil)	//开启监听
	if err != nil {
		log.Fatalf("启动http服务器失败：%v", err)
	}
}

func HandleRequest(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Method = ", request.Method)
	params := request.URL.Query()
	fmt.Println("param: ", params)
	fmt.Println("Method = ", request.Method)

	if request.Method == http.MethodPost {
		fmt.Println("this is a post request")
		con,_ := ioutil.ReadAll(request.Body)
		//fmt.Fprintf(writer, "application/json ->"+string(con))
		paD := json.Decoder{}
		paran := paD.Decode(con)
		fmt.Println("pa = ", paran)
	}
	fmt.Fprintf(writer, "hello, zhizhi, %s", params.Get("name"))
}