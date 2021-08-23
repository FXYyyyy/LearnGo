package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/httpstest", HandleProcess)

		err := http.ListenAndServeTLS(":8443", "./ca.crt", "./ca.key", nil)	//开启监听
	if err!=nil{
		log.Fatalf("启动HTTPS服务器失败：%v", err)
	}
}

func HandleProcess(writer http.ResponseWriter, request *http.Request)  {
	params := request.URL.Query()
	fmt.Fprintf(writer, "hell, world, this Name = %s", params.Get("name"))
}
