package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	// 1. 只读方式打开文件
	file, err := os.Open("./note.md")
	defer file.Close()
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}

	reader := bufio.NewReader(file)

	var content string
	for {
		str, err := reader.ReadString('\n') // 一次读取一行
		if err == io.EOF {
			content += str
			fmt.Println("read file done:", err)
			break
		}
		if err != nil {
			fmt.Println("read file error:", err)
			return
		}
		content += str
	}

	fmt.Println("读取到以下内容")
	fmt.Println(content)
	fmt.Println("读取到结束")
}
