package main

import "fmt"

func main() {
	users := []user{
		{
			name: "zhizhi",
			age: 18,
		},
		{
			name: "lili",
			age: 35,
		},
	}

	//管道，一层一层的计算，流式调用
	ret := sumAges(users, filterUser, mapToString)
	fmt.Println("sum age = ", ret)
}

/**
* 定义一个结构体
*  user
*  @Description:
*/
type user struct {
	name string
	age int
}

/**
* 过滤器
* filterUser
* @Description:
* @param users
* @return interface{}
*/
func filterUser(users []user) interface{} {
	newSlice := make([]user, len(users))
	for _,v := range users {
		if v.age >=18 && v.age <=30{
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}

/**
* 将集合转化为需处理元素的集合
* mapToString
* @Description:
* @param users
* @return interface{}
*/
func mapToString(users []user) interface{} {
	newSlice := make([]int, len(users))
	for _,user := range users {
		newSlice = append(newSlice, user.age)
	}

	return newSlice
}

/**
* 将流式处理的每个步骤加入进来，依次计算，需要将不同的步骤抽象成相同的入参和出参，定义了interface
* sumAges
* @Description:
* @param users
* @param pipe
* @return int
*/
func sumAges(users []user, pipe ...func([]user)interface{}) int {
	var ages []int
	for _,f := range pipe{
		ret := f(users)
		switch ret.(type) {
		case []user:
			users = ret.([]user)
		case []int:
			ages = ret.([]int)
		}
	}

	if len(ages) == 0{
		fmt.Println("there is no get age slice pipe")
		return 0
	}

	sum := 0
	for _,age := range ages{
		sum += age
	}

	return sum
}
