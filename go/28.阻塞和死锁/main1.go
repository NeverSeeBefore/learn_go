package main

// 1. 阻塞
// 当 goroutine 在等待通道接收和发送时，我们就说他阻塞了
// 此时： 除了 goroutine 本身占用的少量内存外，被阻塞的 goroutine 不消耗任何其他资源
// 2. 死锁
// 当一个或者多个 goroutine 因为一个永远无法发生的事情导致阻塞时，我们就说他 死锁，
// 死锁的程序通常会 崩溃 或 挂起

func main() {
	c := make(chan int)
	// go func() { c <- 2 }()
	<-c
	// 因为不会有 goroutine 向通道传值，导致程序死锁，报错
	// fatal error: all goroutines are asleep - deadlock!
}
