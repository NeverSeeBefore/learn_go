package main

import (
	"fmt"
	"time"
)

// 2. 表面上，goroutine 似乎是在同时运行
// 实际上由于计算机处理单元有限，goroutine 并不是同时运行的，
// > 计算机采用一种『分时』的技术,多个 goruotine 轮流执行
// 各个 goroutine 的 执行顺序时不可控的。
func main() {
	for i := 0; i < 5; i++ {
		go sleepyGohper(i)
	}
	time.Sleep(4 * time.Second)
}

// 3秒后打鼾
func sleepyGohper(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...", id)
}
