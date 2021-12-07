package main

import (
	_const "firstGo/testes/const"
	"firstGo/testes/tools"
)

func main() {
	/*art := model.Article{
		Id: 5,
		Title:"lili",
		Content: "hhhhha, xxxx, xixi",
		Category: []int{1, 2},
		CategoryName: []string{"C++", "散文"},
	}
	str := strconv.Itoa(art.Id)
	fmt.Println(str)
	tools.PutToEs(art, _const.MyEsTest, str)*/

	tools.GetOneEsObject(_const.MyEsTest, "1")

	//tools.GetByWildCard(_const.MyEsTest, "Title", "*zhi*")
	searchs := map[string]string{
		"Title": "*zhi*",
		"Content": "*lo*",
	}
	tools.GetByMultiWildCard(_const.MyEsTest, searchs)
}
