package main

import (
	"fmt"
	"strings"
	"time"
)

// close 函数 可以关闭通道
// 通道被关闭后无法写入任何值，尝试写入会 panic
// 尝试读取被关闭的通道，会得到对应的0值
// 注意！ 如果循环读取已关闭的通道，那么循环可能会一致运行下去，会耗费大量的 CPU 时间
// 判断通道是否已被关闭
// v,ok := <- c

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
	for {
		item, ok := <-upstream
		if !ok {
			fmt.Println("通道被关闭 item:", item, "ok:", ok)
			close(downstream)
			return
		}
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

func printGopher(upstream chan string) {

	for {
		item, ok := <-upstream
		if !ok {
			return
		}
		fmt.Println(item)
	}
}
