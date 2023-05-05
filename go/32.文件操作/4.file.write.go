package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./write1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println("open file failed: ", err)
		return
	}

	file.WriteString("通过file.WriteString 写入内容\r\n第二行内容\r\n")
	var str = "通过file.Write 写入内容\r\naaa"
	file.Write([]byte(str))
	fmt.Println()
}
