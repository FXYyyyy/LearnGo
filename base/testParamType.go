package main

import (
	"fmt"
	"reflect"
)

func main() {
	testMultiParam()
}

/**
* 测试入参类型对实参的影响
* testChangeParam
* @Description:
*/
func testChangeParam()  {
	sliceTest := []int{1,2,3}
	changeSlice(sliceTest)
	fmt.Println(sliceTest)

	arrTest := [5]int{1,2,3}
	changeArr(arrTest)
	fmt.Println("arrTest in main = ", arrTest)
}
/**
* 入参为slice类型，引用类型，为引用传递，会改变实参的值
* change
* @Description:
* @param slice1
*/
func changeSlice(slice1 []int)  {
	slice1[0] = 10
}

/**
* 入参为数组类型，值类型，为值传递，不改变实参的值
* changeArr
* @Description:
* @param arr1
*/
func changeArr(arr1 [5]int)  {
	arr1[0] = 10
	fmt.Println("arr1 in func = ", arr1)
}

// ==============================================
/**
* 测试变长的入参
* testMultiParam
* @Description: 
*/
func testMultiParam(){
	slice1 := []int{1,2,3,4,5}
	sameTypeMultiParam(slice1[1:3]...)

	diffTypeMultiParam(1, "1", [1]int{1}, true)
}

/**
* 相同类型的多个传参，行参实际为slice
* sameTypeMultiParam
* @Description:
* @param numbers
*/
func sameTypeMultiParam(numbers ...int)  {
	for k,v := range numbers{
		fmt.Println(k,v)
	}
}

func diffTypeMultiParam(args ...interface{})  {
	for _,arg := range args{
		switch reflect.TypeOf(arg).Kind() {
		case reflect.Int:
			fmt.Println(arg, "is an int value")
		case reflect.String:
			fmt.Printf("\"%s\" is a string value.\n", arg)
		case reflect.Array:
			fmt.Println(arg, "is an array type.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}