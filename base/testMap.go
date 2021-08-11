//字典是一个无序的集合
package main

import (
	"fmt"
)

func main()  {
	var testMap map[string]int
	testMap = initMap()
	printMap(testMap)
	findKeyValue(testMap, "hh")
	traverseMap(testMap)
}

/**
* 初始化一个map
* initMap
* @Description:
*/
func initMap() map[string]int{
	var map1 map[string]int
	map1 = map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
	}
	printMap(map1)

	map2 := map[string]int{
		"kk" : 3,
		"xixi" : 2,
		"liu" : 1,
	}
	printMap(map2)

	//通过make函数定义
	var map3 = make(map[string]int, 100)
	map3["hh"] = 23
	map3["zhi"] = 26
	printMap(map3)

	return map3
}

/**
* 打印一个字典
* printMap
* @Description:
* @param mapTest
*/
func printMap(mapTest map[string]int)  {
	fmt.Println(mapTest)
}

/**
* 找一个键的值
* findKeyValue
* @Description:
* @param mapTest
*/
func findKeyValue(mapTest map[string]int, f string)  {
	value, ok := mapTest[f]
	if ok{
		fmt.Println(value)
	}else{
		fmt.Println("false", ok)
	}
}

/**
* 遍历map
* traverseMap
* @Description:
* @param mapTest
*/
func traverseMap(mapTest map[string]int)  {
	for k, v :=range mapTest{
		fmt.Println(k, v)
	}
}

//键值对调、对key排序、对值排序，都需要自己写

