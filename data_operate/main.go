package main

import (
	"fmt"
	"strings"
)

func StringProcess(list []string, chain []func(string2 string) string) {
	for index, str := range list {
		result := str
		for _, proc :=  range chain {
			result = proc(result)
		}

		list[index] = result
	}
}

// removePrifix
/* @Description: 一个自定义的函数
*  @param str
*  @return string
*/
func removePrefix(str string) string  {
	return strings.TrimPrefix(str, "go")
}

func main() {
	list := []string{
		"go lili",
		"go zhizhi",
		"go xixi",
	}
	chain := []func(string)string {
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}
	StringProcess(list, chain)

	for _, str := range list {
		fmt.Println(str)
	}
}
