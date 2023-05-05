package main

import "fmt"

// 空接口细节

func main() {
	var userInfo = make(map[string]interface{})
	userInfo["name"] = "John Smith"
	userInfo["24"] = 24
	userInfo["hobby"] = []string{"eat", "sleep"}

	fmt.Println(userInfo)
	fmt.Println(userInfo["hobby"])
	// 1. 空接口不支持索引 cannot index userInfo["hobby"] (map index expression of type interface{})
	// fmt.Println(userInfo["hobby"][1])
	// 2. 将切片、结构体等复制给空接口，没办法通过直接获取其下的值，需要使用类型断言

	hobby, ok := userInfo["hobby"].([]string)
	fmt.Println("hobby is a slice:", ok)
	if ok {
		fmt.Println("hobby: ", hobby[0], hobby[1])
	}
}
