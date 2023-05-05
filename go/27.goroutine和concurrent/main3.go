package main

import (
	"fmt"
	"time"
)

// 3. channel 通道
// channel 可以在多个 goroutine 之间安全的传值
// channel 可以用作变量，函数参数，结构体字段 等任何地方
// 创建通道 make(chan 传输数据的类型)
// c := make(chan int)
// <- 向通道发送值或接收值
// c <- 99
// r := <- c
// 发送操作 会等待 另一个 goroutine 尝试执行接收操作为止
// > 执行发送操作的 goroutine 在等待期间无法执行其他操作
// > 未在等待操作的 goroutine 仍然可以自由运行
// 接收操作 会等待 另一个 goroutine 尝试向通道发送数据
func main() {
	// 创建通道
	c := make(chan int)
	// 创建五个 goroutine
	for i := 0; i < 5; i++ {
		go sleepGopher(i, c)
	}

	// 读取通道信息并输出
	for i := 0; i < 5; i++ {
		gopherId := <-c
		fmt.Println("gopher", gopherId, "has finished sleeping")
	}
}

// 3秒后打鼾
func sleepGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore...")
	c <- id
}
