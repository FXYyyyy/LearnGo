package main

import (
	."firstGo/testInterface/interfacerealize"
	"fmt"
)

func main() {
	testRealize()
}

/**
* 接口的赋值
* testRealize
* @Description: 
*/
func testRealize()  {
	//用实现类的实例初始化接口，则其共用的同一份内存
	r := Integer(1)
	var m1 Math1
	m1 = &r
	fmt.Println("init m1 = ", m1)
	r.PrintI()

	fmt.Println("init r =", r)
	fmt.Println("r + 2 = ", r.Add(2))
	fmt.Println("after add r = ", r)

	fmt.Println("after r add m1 = ", m1)
	m1.PrintI()
	fmt.Println("m1 + 3 = ", m1.Add(3))
	fmt.Println("after add m1 = ", m1)
	m1.PrintI()

	var mm Mathmore
	mm = &r
	fmt.Println(mm)
	mm.PrintI()
}
