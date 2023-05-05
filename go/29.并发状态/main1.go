package main

import "sync"

// 共享值
// 竞争条件（race condition）
// go 尝试在编译中发现竞争条件

// 互斥锁 mutex (mutual exclusive)
var mu sync.Mutex

func main() {
	mu.Lock()
	defer mu.Unlock()
}
