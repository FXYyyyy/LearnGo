package tools

import (
	"encoding/json"
	"firstGo/testes/model"
	"firstGo/testes/utils"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
)

// PutToEs
/* @Description: 上传一个对象To ES
*  @param object
*  @param index
*  @param Id
*  @return error
*/
func PutToEs(object interface{}, index string, Id string) error  {
	put1, err := utils.EsClient.Index().
		Index(index).
		Id(Id).
		BodyJson(object).
		Do(utils.EsCtx)
	if err != nil {
		fmt.Println("上传一个对象失败: ", err)
		return err
	}else{
		fmt.Println("上传成功", put1)
	}

	return nil
}

// GetOneEsObject
/* @Description: 获取一个Es的文档，返回一个struct
*  @param index
*  @param Id
*  @return interface{}
*  @return error
*/
func GetOneEsObject(index string, Id string) (interface{}, error) {
	get1, err := utils.EsClient.Get().
		Index(index).
		Id(Id).
		Do(utils.EsCtx)

	if err != nil {
		fmt.Println("获取一个对象失败: ", err)
		return nil, err
	}else{
		fmt.Println("获取成功")
	}

	msg := new(interface{})

	//提取文档内容，原始类型是json数据
	data, _ := get1.Source.MarshalJSON()

	//将json转换成struct
	json.Unmarshal(data, &msg)

	fmt.Println("得到的结果是", *msg)

	return msg, nil
}

// GetByWildCard
/* @Description: 根据一个字段模糊匹配
*  @param index
*  @param name
*  @param value
*  @return []interface{}
*  @return error
*/
func GetByWildCard(index string, name string, value string) ([]interface{}, error) {
	wildQuery := elastic.NewWildcardQuery(name, value)

	ret, err := utils.EsClient.Search().
		Index(index).
		Query(wildQuery).
		Do(utils.EsCtx)

	if err != nil {
		fmt.Println("模糊查询失败:", err)
		return nil, err
	}else{
		fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", ret.TookInMillis, ret.TotalHits())
	}

	var returnData []interface{}
	var obj model.Article

	if ret.TotalHits() > 0 {
		fmt.Println("搜索到的数据一共: ", ret.TotalHits(), "条")

		for _, item := range ret.Each(reflect.TypeOf(obj)) {
			if t, ok := item.(model.Article); ok {
				fmt.Println("数据是: ", t)
				returnData = append(returnData, t)
			}
		}

		return returnData,nil
	}else{
		fmt.Println("暂无搜索到的数据")
		return nil, nil
	}
}

func GetByMultiWildCard(index string, searchs map[string]string) ([]interface{}, error) {
	//创建bool查询
	boolQuery := elastic.NewBoolQuery().Must()

	//遍历条件依次创建
	for name, value := range searchs {
		fmt.Println("name: ", name, "value: ", value)
		wildQuery := elastic.NewWildcardQuery(name, value)

		boolQuery.Should(wildQuery)
	}

	ret, err := utils.EsClient.Search().
		Index(index).
		Query(boolQuery).
		Do(utils.EsCtx)

	if err != nil {
		fmt.Println("模糊查询失败:", err)
		return nil, err
	}else{
		fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", ret.TookInMillis, ret.TotalHits())
	}

	var returnData []interface{}
	var obj model.Article

	if ret.TotalHits() > 0 {
		fmt.Println("搜索到的数据一共: ", ret.TotalHits(), "条")

		for _, item := range ret.Each(reflect.TypeOf(obj)) {
			if t, ok := item.(model.Article); ok {
				fmt.Println("数据是: ", t)
				returnData = append(returnData, t)
			}
		}

		return returnData,nil
	}else{
		fmt.Println("暂无搜索到的数据")
		return nil, nil
	}
}