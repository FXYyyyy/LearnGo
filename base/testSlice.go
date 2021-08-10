//切片
package main

import "fmt"

func main()  {
	testArrAndSlice()
}

/** 定义切片
*  varSlice
*  @Description:
*/
func varSlice()  {
	//定义切片
	slice1 := []string{"a", "b", "是"}

	for i, ch := range slice1{
		fmt.Println(i,ch)
	}
}

/** 测试切片与原数组的关系，为引用，改变切片的值会改变原数组
*  testArrAndSlice
*  @Description:
*/
func testArrAndSlice()  {
	arr1 := [5]int{1,2,3,4,5}

	slice1 := arr1[1:3]
	slice1[0] = 9

	//遍历切片
	for i,x := range slice1{
		fmt.Println(i, x)
	}

	fmt.Println("===========")
	//遍历原数组
	for i,x := range arr1{
		fmt.Println(i,x)
	}
}