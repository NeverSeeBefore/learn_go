package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixMilli()
	flag := true
	count := 0
	for i := 2; i < 120000; i++ {
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
	fmt.Println("共", count, "m个素数")
	end := time.Now().UnixMilli()
	fmt.Println("耗时", end-start, "milliseconds")
	// 共 11301 m个素数
	// 耗时 584 milliseconds
}
