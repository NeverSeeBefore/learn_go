package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://www.jianshu.com/p/de4bc02e7c72

// 竞争选举
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	c4 := make(chan int)
	c5 := make(chan int)

	go fn(c1)
	go fn(c2)
	go fn(c3)
	go fn(c4)
	go fn(c5)

	select {
	case v1 := <-c1:
		fmt.Println("the fastest is c1", v1)
	case v2 := <-c2:
		fmt.Println("the fastest is c2", v2)
	case v3 := <-c3:
		fmt.Println("the fastest is c3", v3)
	case v4 := <-c4:
		fmt.Println("the fastest is c4", v4)
	case v5 := <-c5:
		fmt.Println("the fastest is c5", v5)
	}
}

func fn(c chan int) {
	t := time.Duration(rand.Intn(4000)) * time.Millisecond
	time.Sleep(t)
	c <- int(t)
}
