package main

import "fmt"

func main(){
	a := [5]int{1,2,3,4}
	//使用range遍历
	for _, x := range a{
		fmt.Println(x)
	}
}