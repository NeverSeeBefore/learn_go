package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	byteSlice, err := ioutil.ReadFile("./note.md")
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}
	fmt.Println("读取到以下内容")
	fmt.Println(string(byteSlice))
	fmt.Println("读取到结束")
}
