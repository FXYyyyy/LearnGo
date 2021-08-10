package main

import "fmt"

func main() {
	str := "Hello, 世界"
	utfRun(str)
	fmt.Println("===========")

	unicodeRun(str)
	fmt.Println("***********")

	unicodeTransform(str)
}

//字节数组的遍历方式
func utfRun(str string){
	n := len(str)
	for i:=0;i<n;i++{
		ch := str[i]
		fmt.Println(i, ch)
	}
}

//字符遍历的方式
func unicodeRun(str string)  {
	for i, ch := range str{
		fmt.Println(i, ch)
	}
}

//将unicode转化成
func unicodeTransform(str string){
	for i, ch := range str{
		fmt.Println(i, string(ch))
	}
}