package main

import (
	"fmt"
	"strings"
	"time"
)

// close 函数 可以关闭通道
// 判断通道是否已被关闭
// v,ok := <- c

// 【从通道中读取值，知道关闭位置】
// 除了判断 ok 的方式
// 还可以使用 range

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbay all"} {
		time.Sleep(2 * time.Second)
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {

	for item := range upstream {
		fmt.Println(item)
	}
}
