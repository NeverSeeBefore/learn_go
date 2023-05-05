package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	err := ioutil.WriteFile("./write3.txt", []byte("通过ioutil写入文件"), 0666)
	if err != nil {
		fmt.Println("写入文件失败", err)
	}
}
