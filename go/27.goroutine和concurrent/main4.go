package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 4. 使用 select 处理多通道
// 等待不同类型的值
// time.After() 函数，返回一个通道，这个通道会在指定时间后接受到一个值
// 通过select 可以通知等待多个通道，直到某个case准备就绪。
// 注意！即使已经停止等待 goroutine，只要 main 函数没有返回，仍在运行的 goroutine 将会继续占用内存

func main() {
	// 创建通道
	c := make(chan int)
	// 创建五个 goroutine
	for i := 0; i < 5; i++ {
		go sleepGopher(i, c)
	}

	timeout := time.After(2000 * time.Millisecond)
	// 读取通道信息并输出
	for i := 0; i < 5; i++ {
		select {
		case gopherId := <-c:
			fmt.Println("gopher", gopherId, "has finished sleeping")
		case <-timeout:
			// case time := <-timeout:
			fmt.Println("my patiance ran out")
			return
		}
	}
}

// 0-4000ms 后打鼾
func sleepGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("... ", id, " snore...")
	c <- id
}
