package main

import "fmt"

func main() {
	// 管道默认容量0
	// 此时进行发送操作后，会等待取值后继续执行
	// make(chan int, 1) 设置容量为1
	// 此时进行发送操作化，数据可以进入管道，无序等待取值
	// 再次发送时，才会出现等待
	c := make(chan int, 1)

	c <- 10
	fmt.Printf("len %v cap %v", len(c), cap(c))

	a := <-c

	fmt.Println("a:", a)
}
