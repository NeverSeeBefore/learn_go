package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 5. 如果不使用make初始化通道，通道变量的值就是 nil
// 对 nil 通道 进行发送或接收，不会panic，但会导致永久的堵塞
// 对 nil 通道 执行 close 会导致panic
// nil 通道的用处:
//
//	 	对于包含select语句的循环，如果不希望每次循环都等待select涉及的所以通道，
//		可以先将某些通道设置为 nil，等到发送值准备就绪后，再将通道变为一个非 nil 的值，执行发送操作

func main() {
	// 创建通道
	c := make(chan int)

	fmt.Printf("%v \n%T \n%#v\n", c, c, c)
	// 0x14000102060
	// chan int
	// (chan int)(0x14000102060)

	var c2 chan int
	fmt.Printf("%v \n%T \n%#v\n", c2, c2, c2)
	// <nil>
	// chan int
	// (chan int)(nil)
}

// 0-4000ms 后打鼾
func sleepGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("... ", id, " snore...")
	c <- id
}
