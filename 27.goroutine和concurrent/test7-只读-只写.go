package main

import "fmt"

// 只读只写管道可以用在函数参数上，以限制函数内操作权限

func main() {
	// 默认情况下为双向管道
	c1 := make(chan int, 2)
	c1 <- 1
	c1 <- 2
	m1 := <-c1
	m2 := <-c1
	fmt.Println(m1, m2)

	// 只写管道
	c2 := make(chan<- int, 2)
	c2 <- 3
	c2 <- 4
	// m3 := <-c2
	// fmt.Println(m3)
	// ./test7-只读-只写.go:18:10: invalid operation: cannot receive from send-only channel c2 (variable of type chan<- int)

	// 只读管道
	// c3 := make(<-chan int, 2)
	// c3 <- 5
	// ./test7-只读-只写.go:25:2: invalid operation: cannot send to receive-only channel c3 (variable of type <-chan int)
}

// [https://blog.csdn.net/weixin_43893483/article/details/119818846]golang之channel注意细节（只读、只写、select、recover）
