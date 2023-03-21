package main

import (
	"flag"
	"fmt"
	"strings"
)

type SliceValue []string

func NewSliceValue(vals []string, p *[]string) *SliceValue {
	*p = vals
	return (*SliceValue)(p)
}

func (s *SliceValue) Set(val string) error {
	*s = SliceValue(strings.Split(val, ","))
	return nil
}

func (s *SliceValue) Get() interface{} {
	return []string(*s)
}

func (s *SliceValue) String() string {
	return strings.Join([]string(*s), "@")
}

// main 实现自定义的flag
func main() {
	var myCat []string
	flag.Var(NewSliceValue([]string{}, &myCat), "catName", "My cat Name")

	flag.Parse()
	fmt.Println(myCat)
}
