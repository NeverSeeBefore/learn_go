package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

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
		}
		flag = true
	}
	exitC <- 1
}

func printPrime(primeC chan int) {
	defer wg.Done()
	count := 0
	for num := range primeC {
		count++
		fmt.Println(num)
	}
	fmt.Println("供", count, "个素数")
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
}
