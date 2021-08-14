package main

import (
	."firstGo/testPackage/animal"	//使用 . 作为前缀，避免使用包名作为前缀来进行调用
	"fmt"
)

func main() {
	//testClass()
	testAssertion()
}

/**
* 测试类的实例和组合
* testClass
* @Description:
*/
func testClass()  {
	// a := animal.Animal{"中华田园猫"}
	/*
		a := Animal{"中华田园猫"}
		p := Pet{"宠物猫kk"}
		c := Cat{a, p}
	*/
	a := NewAnimal("中华田园猫")
	p := NewPet("宠物猫kk")
	c := Cat{Animal:a, Pet:p}

	//fmt.Println(c.GetName())
	fmt.Println(c.Call())
	fmt.Println(c.FavorFood())
}

/**
* 测试类型断言，断言语法. 的左侧必须是接口类型
* testAssertion
* @Description:
*/
func testAssertion()  {
	animal := NewAnimal("中华田园猫")
	pet := NewPet("kk")
	//cat := Cat{Animal: animal, Pet:pet}
	//var inAnimal InterfaceAnimal = cat
	var inAnimal InterfaceAnimal = NewCat(animal, pet)

	fmt.Println(inAnimal.FavorFood())

	//断言是否Cat类, true
	if x, ok := inAnimal.(Cat); ok{
		fmt.Println(x)
		fmt.Println("is Cat Type")
	}else{
		fmt.Println("not Cat Type")
	}

	//断言是否是Animal, false
	if x, ok := inAnimal.(Animal); ok{
		fmt.Println(x)
		fmt.Println("is Animal Type")
	}else{
		fmt.Println("not Animal Type")
	}

	//断言是否是Animal, true
	if x, ok := inAnimal.(InterfaceAnimal); ok{
		fmt.Println(x)
		fmt.Println("is InterfaceAnimal Type")
	}else{
		fmt.Println("not InterfaceAnimal Type")
	}
}