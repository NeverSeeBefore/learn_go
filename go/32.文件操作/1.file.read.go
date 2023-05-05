package main

import (
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

	fmt.Println("file ptr:", file)

	var strSlice []byte
	var tempSlice = make([]byte, 128)

	for {
		// 使用file.Read 读取文件
		n, err := file.Read(tempSlice)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完成", err)
				break
			} else {
				fmt.Println("file.Read error:", err)
				return
			}
		}
		fmt.Printf("读取到%v字节数据\n %v\n", n, string(tempSlice))
		strSlice = append(strSlice, tempSlice[:n]...)
	}

	fmt.Println("读取到以下内容")
	fmt.Println(string(strSlice))
	fmt.Println("读取到结束")

	// 写入文件
}
