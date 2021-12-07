package utils

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"os"

)

var (
	EsClient *elastic.Client
	err error
	EsCtx context.Context
)

func init() {
	EsClient, err = elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),	//服务器地址
		elastic.SetMaxRetries(5),			//请求失败最大重试次数
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),	//设置错误输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))			//设置info日志输出

	if err != nil {
		fmt.Println("连接ES失败: ", err)
		panic(err.Error())
	}else {
		fmt.Println("连接ES 成功")
	}

	EsCtx = context.Background()
}
