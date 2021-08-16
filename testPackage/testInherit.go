package main

import (
	. "firstGo/testPackage/animal" //使用 . 作为前缀，避免使用包名作为前缀来进行调用
	"fmt"
	"reflect"
)

func main() {
	//testClass()
	//testAssertion()
	testReflect()
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
	var inAnimal InterfaceAnimal = &Cat{animal, pet}
	//var inAnimal InterfaceAnimal = animal

	fmt.Println(inAnimal.FavorFood())
	fmt.Println(reflect.TypeOf(inAnimal))	//基于反射动态进行类型断言

	//断言是否Cat类, true
	if x, ok := inAnimal.(*Cat); ok{	//Cat包含指针类型的方法
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

	//i := 10
	//fmt.Println(i.(type))	必须用在switch语句中
}

func testReflect()  {
	animal := NewAnimal("中华田园猫")
	pet := NewPet("kk")
	cat := Cat{Animal: animal, Pet:pet}

	//reflect反射，判断类型
	catType := reflect.TypeOf(cat)
	fmt.Println("cat Type", catType)

	//reflect反射，获取类型值
	catValue := reflect.ValueOf(cat)
	//catValue := reflect.ValueOf(&cat)	//在下面获取为Cat* 类型，获取其方法会引起panic
	fmt.Println("cat Value: ", catValue)

	//reflect反射，获取实例的所有属性
	for i:=0; i< catValue.NumField(); i++ {
		//获取属性名
		cname := catValue.Type().Field(i).Name
		//获取属性类型
		ctype := catValue.Type().Field(i).Type
		//获取属性值
		cvalue := catValue.Type().Field(i)

		fmt.Println("cname, ctype, cValue: ", cname, ctype, cvalue)
	}

	//reflect反射，获取实例的所有方法，仅获取到了值方法
	for i := 0; i<catValue.NumMethod(); i++{
		//获取方法名
		fName := catValue.Type().Method(i).Name
		fmt.Println("fName: ", fName)

		//获取方法类型
		fType := catValue.Type().Method(i).Type
		fmt.Println("fType: ", fType)

		//调用该方法
		fmt.Println(catValue.Method(i).Call([]reflect.Value{}))
	}
}