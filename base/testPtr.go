package main

import "fmt"

func main() {
	initPtr()
}

func initPtr()  {
	var ptr1 *int
	a := 100
	ptr1 = &a
	printfPtr(ptr1)

	ptr2 := new(int)
	*ptr2 = 200
	printfPtr(ptr2)
}

/**
* 打印指针本身和指向的值
* printfPtr
* @Description:
* @param ptr
*/
func printfPtr(ptr *int)  {
	fmt.Printf("ptr1Self = %p\n", ptr)
	fmt.Printf("ptrValue = %d\n", *ptr)
}
