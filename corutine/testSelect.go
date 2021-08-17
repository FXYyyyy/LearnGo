package main

import (
	"fmt"
	"math/rand"
)

func main() {
	simpleSelect()
}

/**
* select的简单实现
* simpleSelect
* @Description:
*/
func simpleSelect()  {
	chs := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	randInt := rand.Intn(3)	//随机生成0-2的整型
	fmt.Println("randInt = ", randInt)
	//chs[randInt] <- randInt

	select {
	case <-chs[0]:
		fmt.Println("First Chan")
		case <-chs[1]:
			fmt.Println("Second Chan")
			case num := <-chs[2]:
				fmt.Println("Third Chan : ", num)
	default:
		fmt.Println("No Chan Select")
	}
}