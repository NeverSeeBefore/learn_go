package main

import (
	"fmt"
	"strings"
	"time"
)

// 地鼠装配线
// sourceGopher 产生字符串
// filterGopher 过滤掉包含 bad 的字符串
// printGopher 打印字符串
// 当字符串为 "" 时，结束
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
	downstream <- ""
}

func filterGopher(upstream, downstream chan string) {
	for {
		item := <-upstream
		if item == "" {
			downstream <- item
			return
		}
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}

func printGopher(upstream chan string) {

	for {
		item := <-upstream
		if item == "" {
			return
		}
		fmt.Println(item)
	}
}
