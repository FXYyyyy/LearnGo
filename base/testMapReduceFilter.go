package main

import (
	"fmt"
	"strconv"
)

func main() {
	users := []map[string]string{
		{
			"name": "zhizhi",
			"age": "18",
			"year": "2",
		},
		{
			"name": "lili",
			"age": "30",
			"year": "3",
		},
	}

	sumAge(users)	//原始方法

	// ===== 使用map-reduce-filter进行计算，利于扩展 ======
	//获取关键字函数
	getFunc := func(filed string, m map[string]string) string{
		value,ok := m[filed]
		if ok{
			return value
		}else{
			return "0"
		}
	}

	//过滤器函数
	filterFunc := func(filed string, m map[string]string)bool{
		value, ok := m[filed]
		if !ok {
			return false
		}
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		if valueInt < 18 || valueInt > 24 {
			return false
		}
		return true
	}

	//开始计算
	filed := "age"
	//过滤
	users = filterMap(users, filed, filterFunc)
	//将集合转化为slice
	reduceSlice := mapToStringSlice(users, filed, getFunc)
	//对转化后的slice进行计算
	sumRet := filedSum(reduceSlice, func(s string) int {
		num,_ := strconv.Atoi(s)
		return num
	})

	fmt.Println(sumRet)
}

/**
* 对集合进行求年龄和
* sumAge
* @Description: 
* @param users
* @return int
*/
func sumAge(users []map[string]string) int {
	sum := 0
	for _,user := range users{
		age, ok := user["age"]
		if ok{
			num,_ := strconv.Atoi(age)
			sum += num
		}else{
			fmt.Println("there is no age")
		}
	}

	fmt.Println("age sum = ", sum)
	return sum
}

/**
* 对map的过滤器
* filterMap
* @Description:
* @param users
* @param filter
* @return []map[string]string
*/
func filterMap(users []map[string]string,filed string, filter func(string, map[string]string) bool) []map[string]string {
	newSlice := make([]map[string]string, len(users))
	for _,user := range users{
		if filter(filed, user){
			newSlice = append(newSlice, user)
		}else{
			fmt.Println("this is not valid", user)
		}
	}

	return newSlice
}
/**
* reduce --- 获取需要处理的集合作为一个平铺的切片
* mapToStringSlice
* @Description:
* @param users
* @param f
* @return []string
*/
func mapToStringSlice(users []map[string]string, filed string,f func(string, map[string]string) string) []string {
	newSlice := make([]string, len(users))	//指定分片的容量
	for _,user := range users{
		needString := f(filed, user)
		newSlice = append(newSlice, needString)
	}

	return newSlice
}

/**
* 对reduce处理后的slice，进行进一步的计算
* filedSum
* @Description:
* @param slice
* @param f
* @return int
*/
func filedSum(slice []string, f func(string)int) int  {
	sum := 0
	for _,v := range slice{
		num := f(v)
		sum += num
	}

	return sum
}