package main

import (
	"fmt"
	"sync"
	"time"
)

// sync 包
var wg sync.WaitGroup

func test(num int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("协程(%v)打印第%v条数据\n", num, i)
		time.Sleep(time.Second)
	}
}

func main() {
	for i := 0; i <= 6; i++ {
		wg.Add(1)
		go test(i)
	}

	wg.Wait()
	fmt.Println("关闭主线程！")
}
