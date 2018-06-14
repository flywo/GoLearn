package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string
	Authors []string
	Publisher string
	IsPublished bool
	Price float64
}

func main() {
	//json.Marshal()对一组数据进行JSON格式的编码
	gobook := Book{"编程语言", []string{"x", "y","z"}, "li.com.cn", true, 9.9}
	b, err := json.Marshal(gobook)//转换过程如果发现指针，则会转化成将指针所指向的值，如果指针指向的是零值，那么Null将作为转化后的结果输出。
	//map必须是[string]T才能被转化
	if err!=nil {
		fmt.Println("发生了错误")
	}
	fmt.Println(string(b))

	//json.Unmarshal()解析json字符串
	var debook Book
	err = nil
	err = json.Unmarshal(b, &debook)
	if err!=nil {
		fmt.Println("解析发生了错误")
	}
	fmt.Println(debook, debook.Title)

	//使用未知类型
	var r interface{}
	err = nil
	err = json.Unmarshal(b, &r)
	fmt.Println(r)
	gobook2, ok := r.(map[string]interface{})
	if ok {
		for k, v := range gobook2 {
			switch v2 := v.(type) {
			case string:
				fmt.Println(v2)
			case int:
				fmt.Println(v2)
			case bool:
				fmt.Println(v2)
			case []interface{}:
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println(k, "别的key")
			}
		}
	}
}
