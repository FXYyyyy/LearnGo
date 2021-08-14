package main

import (
	."firstGo/testPackage/animal"	//使用 . 作为前缀，避免使用包名作为前缀来进行调用
	"fmt"
)

func main() {
	// a := animal.Animal{"中华田园猫"}
	/*
	a := Animal{"中华田园猫"}
	p := Pet{"宠物猫kk"}
	c := Cat{a, p}
	 */
	a := NewAnimal("中华田园猫")
	p := NewPet("宠物猫kk")
	c := Cat{Animal:a, Pet:p}

	fmt.Println(c.GetName())
	fmt.Println(c.Call())
	fmt.Println(c.FavorFood())
}
