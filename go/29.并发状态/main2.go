package main

// worker
// 长时间运行的工作进程
// 通常会被写成包含select的for循环

func worker() {
	for {
		select {
		// cahnnels
		}
	}
}

func main() {
	go worker()
}
