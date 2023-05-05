package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixMilli()
	count := 0
	c := make(chan int)
	go search(2, 30000, c)
	go search(30001, 60000, c)
	go search(60001, 90000, c)
	go search(90001, 120000, c)

	flag := 0
	for {
		count = count + <-c
		flag++
		if flag == 4 {
			break
		}
	}

	fmt.Println("共", count, "m个素数")
	end := time.Now().UnixMilli()
	fmt.Println("耗时", end-start, "milliseconds")
	// 共 11301 m个素数
	// 耗时 257 milliseconds
}

func search(start, end int, c chan int) {
	count := 0
	flag := true
	for i := start; i < end; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			count++
			// fmt.Println(i)
		}
		flag = true
	}
	c <- count
}
