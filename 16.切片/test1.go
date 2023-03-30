package main

import (
	"fmt"
	"sort"
	"strings"
)

func hyperspace(words []string) {
	for i := range words {
		words[i] = strings.TrimSpace(words[i])
	}
}

func main() {
	planets := []string{"  Venus  ", "Earth  ", "   Mars"}

	// 去空格
	hyperspace(planets)

	// 排序
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)

	// 拼接字符串
	fmt.Println(strings.Join(planets, "、"))
}
