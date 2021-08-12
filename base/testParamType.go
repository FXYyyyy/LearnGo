package main

import "fmt"

func main() {
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