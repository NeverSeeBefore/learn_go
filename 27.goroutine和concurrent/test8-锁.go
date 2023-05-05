package main

// go build ./main.go
// go build -race ./main.go

// 同样还是 开启多个协程计算素数的代码
// 这次设置一个全局变量用于统计素数个数

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var count int = 0

// 互斥锁
// 同时只有一个 goroutine 进行操作
var mutex sync.Mutex

// 读写锁
// 写操作 同时只有一个 goroutine 进行操作
// 读操作 如果有写操作则所有 goroutine 都不能读，如果没有写操作 则所有 goroutine 都可以读
// var rwMutex sync.RWMutex

func increaseCount() {
	mutex.Lock()
	defer mutex.Unlock()
	count++
}

func putNum(inC chan int) {
	defer wg.Done()
	for i := 2; i < 120000; i++ {
		inC <- i
	}
	close(inC)
}

func primeNum(inC, primeC, exitC chan int) {
	defer wg.Done()
	flag := true
	for num := range inC {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeC <- num
			increaseCount()
		}
		flag = true
	}
	exitC <- 1
}

func printPrime(primeC chan int) {
	defer wg.Done()
	for num := range primeC {
		fmt.Println(num)
	}
}

func main() {

	start := time.Now().UnixMilli()

	primeCCount := 4
	inC := make(chan int, 1000)
	primeC := make(chan int, 1000)
	exitC := make(chan int, primeCCount)

	wg.Add(1)
	go putNum(inC)

	for i := 0; i < primeCCount; i++ {
		wg.Add(1)
		go primeNum(inC, primeC, exitC)
	}

	wg.Add(1)
	go printPrime(primeC)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < primeCCount; i++ {
			<-exitC
		}
		close(primeC)
		close(exitC)
	}()

	wg.Wait()

	end := time.Now().UnixMilli()
	fmt.Println("执行完毕")
	fmt.Println("耗时", end-start, "milliseconds")
	fmt.Println("共", count, "个素数")
}
