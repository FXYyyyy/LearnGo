package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// MyValue 实现了一个自定义的flag.Value，入参以,分割的字符串，解析后复制给struct
type MyValue struct {
	Name string
	Age  int
}

func NewMyValue(a MyValue, p *MyValue) *MyValue {
	p.Name = a.Name
	p.Age = a.Age

	return p
}

func (m *MyValue) Set(val string) error {
	x := strings.Split(val, ",")
	m.Name = x[0]
	if len(x) < 2 {
		return fmt.Errorf("参数需要有两个及以上")
	}
	age, err := strconv.Atoi(x[1])
	if err != nil {
		return err
	}
	m.Age = age
	return nil
}

func (m *MyValue) String() string {
	return fmt.Sprintf("this is my struct name = %s, age = %d", m.Name, m.Age)
}

func main() {
	var myCat MyValue
	flag.Var(NewMyValue(MyValue{}, &myCat), "myValue", "define my value struct")

	flag.Parse()
	fmt.Println(myCat.String())
}
