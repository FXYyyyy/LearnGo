package main

import "fmt"

func main() {
	score := 60
	switch score {
	case 90, 100:
		fmt.Println("Grade: A")	//没有写break，自己结束
	case 80:
		fmt.Println("Grade: B")
	case 70:
		fmt.Println("Grade: C")
	case 60:
		fallthrough	//继续往下执行，默认为break
	case 65:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}
}
