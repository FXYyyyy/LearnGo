//切片
package main

import "fmt"

func main()  {
	//testArrAndSlice()
	//sliceAdd()
	copySlice()
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

/** 测试切片与原数组的关系，为引用，改变切片的值会改变原数组，存在内容共享，可以append重新分配空间，解决该问题
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

/** 切片元素的添加
*  sliceAdd
*  @Description:
*/
func sliceAdd()  {
	slice1 := []int{1,2,3,4,5}
	printSliceInfo(slice1)

	//追加append，动态扩容，但是底层仍旧是数组的扩容，重新分配空间，需要进行内存的移动和复制
	slice2 := append(slice1, 6,7,8,9,10,11)
	printSliceInfo(slice2)
}

/** copy操作，以元素较少的切片复制
*  copySlice
*  @Description:
*/
func copySlice()  {
	sliceMore := []int{1,2,3,4,5}
	sliceLess := []int{7,8,9}

	//将less复制到more，结果more为[7,8,9,4,5]
	copy(sliceMore, sliceLess)

	printSliceInfo(sliceMore)
}

/** 打印切片的信息
*  printSliceInfo
*  @Description:
*  @param slice
*/
func printSliceInfo(slice []int)  {
	fmt.Println("len = ", len(slice), "en")
	fmt.Println("cap = ", cap(slice), "===")	//切片的容量

	for i, x := range slice{
		fmt.Println(i, x)
	}
}