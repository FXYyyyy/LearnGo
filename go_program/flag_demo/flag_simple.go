package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var name string
	var age int
	//var married bool
	var delay time.Duration

	flag.StringVar(&name, "name", "KK", "姓名")
	flag.IntVar(&age, "age", 4, "年龄")
	//flag.BoolVar(&married, "married", false, "婚否")
	married := flag.Bool("married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟时间间隔")

	flag.Parse() //解析命令行的参数
	fmt.Println(name, age, married, delay)

	fmt.Println(flag.Args())  //返回命令行参数后的其他参数
	fmt.Println(flag.NArg())  //返回命令行参数后的其他参数的个数
	fmt.Println(flag.NFlag()) //返回使用的命令行参数个数
}
