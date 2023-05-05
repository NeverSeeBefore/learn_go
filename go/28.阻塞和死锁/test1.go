package main

import "fmt"

func main() {
	c := make(chan int)
	// 1. 读取已关闭的通道
	close(c)
	fmt.Println("读取已关闭的通道:", <-c)
	// 读取已关闭的通道: 0
	// 读取到对应 零值

	// 2. 向关闭的通道写入
	c <- 9
	// panic: send on closed channel
	// 会导致 panic

}
