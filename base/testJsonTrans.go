package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	WebSite string
	Age uint
	Male bool
	Skills []string
	id uint
}

func main() {
	uJson := testJsonMarshal()

	//var user2 User

	//将uJson进行解码
	//testJsonUnMarshal(uJson, &user2)

	testJsonUnKnow(uJson)
}

/**
* testJsonMarshal
* @Description:将数据类型转化成json
* @return ret
*/
func testJsonMarshal() (ret string) {
	user := User{
		"学院君",
		"https://xueyuanjun.com",
		18,
		true,
		[]string{"Golang", "PHP", "C", "Java", "Python"},
		1,	//仅首字母大写的变量可以转化，小写的不行
	}

	uJson, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("JSON 编码失败：%v\n", err)
		return
	}

	fmt.Println("uJsonSelf = ", uJson)		//打印出来是unicode的编码
	fmt.Println("uJson = ", string(uJson))	//需对uJson进行string的转化，否则打印出来的都是unicode的数字
	fmt.Printf("uJson = %s\n", uJson)

	return string(uJson)
}

/**
* testJsonUnMarshal
* @Description: 将byte转化成对应的数据类型
* @param uJson
* @param user
*/
func testJsonUnMarshal(uJson string, user *User)  {
	uJsonByte := []byte(uJson)
	err := json.Unmarshal(uJsonByte, user)
	if err != nil {
		fmt.Printf("JSON 解码失败：%v\n", err)
		return
	}

	fmt.Println("user = ", user)
	fmt.Printf("JSON 解码结果: %#v\n", user)
}

/**
* testJsonUnKnow
* @Description: 在未知结构时，使用interface
* @param uJson
*/
func testJsonUnKnow(uJson string){
	uJsonByte := []byte(uJson)

	var x interface{}

	json.Unmarshal(uJsonByte, &x)

	xnew,ok := x.(map[string]interface{})

	if !ok {
		fmt.Println("转换成map失败")
		return
	}
	for k,v := range xnew {
		fmt.Println("v =====", v)
		switch v2 := v.(type) {
		case string:
			fmt.Println(k, "is string", v2)
		case int:
			fmt.Println(k, "is int", v2)
		case bool:
			fmt.Println(k, "is bool", v2)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, iv := range v2 {
				fmt.Println(i, iv)
			}
		default:
			fmt.Println(k, "is another type not handle yet")
		}
	}

	fmt.Println("x ==== ", x)
	fmt.Println("xnew=====", xnew)
}