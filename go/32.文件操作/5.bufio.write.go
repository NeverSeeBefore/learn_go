package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./write2.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println("open file failed: ", err)
		return
	}
	writer := bufio.NewWriter(file)
	// 写入缓存
	writer.WriteString("通过bufio写入内容\r\n第二行")
	// 将缓存内容写入文件
	writer.Flush()
}
