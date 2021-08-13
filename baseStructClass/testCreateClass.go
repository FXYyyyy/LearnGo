package main

import (
	"fmt"
)

func main() {
	stu1 := NewStudent(1, "zhizhi", true, 95.5)
	fmt.Println(stu1.GetName())

	fmt.Printf("%T", stu1)

	stu1.setScore(90)	//stu1是一个指针类型，但是在使用时，会自动的解引用
	fmt.Println(stu1)
}

type Student struct {
	id uint
	name string
	male bool
	score float64
}

/**
* 初始化一个struct
* NewStudent
* @Description: 
* @param id
* @param name
* @param male
* @param score
* @return *Student
*/
func NewStudent(id uint, name string, male bool, score float64) *Student {
	return &Student{id, name, male, score}
}

/**
* 获取一个名字（值方法）
* GetName
* @Description:
* @receiver s
* @return string
*/
func (s Student) GetName() string {
	return s.name
}

/**
* 设置分数（指针方法）
* setScore
* @Description:
* @receiver s
* @param newScore
*/
func (s *Student) setScore(newScore float64) {
	s.score = newScore
}

/**
* 以字符串格式打印类的实例，无需显式调用 String 方法，Go 语言会自动调用该方法来打印
* String
* @Description:
* @receiver s
* @return string
*/
func (s Student) String() string {
	return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}",
		s.id, s.name, s.male, s.score)
}