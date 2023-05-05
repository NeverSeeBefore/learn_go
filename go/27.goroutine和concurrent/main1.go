package main

import (
	"fmt"
	"time"
)

// 地鼠 gopher
// 建设有一个地鼠工厂

// goroutine
// 在 go 中 独立的任务叫做 goroutine
// goroutine 与其他语言中的 协成、进程、线程有很多相似之处，但并不完全相同
// goroutine 创建效率非常高
// go 可以直截了当的协同多个并发（concurrent）操作

// 1. 每次使用 go 关键字，就会创建一个新的独立运行的任务
func main() {
	// 启动一个goroutine
	go sleepyGohper()
	// 等待4s，因为
	time.Sleep(4 * time.Second)
}

// 3秒后打鼾
func sleepyGohper() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")
}
