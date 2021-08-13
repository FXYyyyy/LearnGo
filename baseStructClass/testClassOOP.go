package main

import "fmt"

func main() {
	a := Animal{"田园猫"}
	p := Pet{"kk"}

	//fmt.Println(a.getName())
	fmt.Println(a.call())
	fmt.Println(a.favorFood())

	c := Cat{a, p}
	fmt.Println(c.getName())
	fmt.Println(c.call())
	fmt.Println(c.favorFood())
}

//动物类
type Animal struct {
	name string
}
/*
func (a Animal) getName() string {
	return a.name
}
 */
func (a Animal) call() string {
	return "动物的叫声=="
}
func (a Animal) favorFood() string {
	return "最爱吃的食物=="
}

//宠物类
type Pet struct {
	name string
}

func (p Pet) getName() string {
	return p.name
}

//小猫，组合动物类、宠物类 ==== 继承
type Cat struct {
	Animal
	Pet
}

//重构动物类的方法 ==== 多态
//继承存在函数名字冲突时，会报错，需要组合的类重写
func (c Cat) call() string {
	return "喵喵喵"
}

func (c Cat) favorFood() string {
	return "fish"
}

/*
func (c Cat) getName() string {
	return c.Animal.name + c.Pet.name
}
 */