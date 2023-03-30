package main

import "fmt"

// go 从main开始运行
// go run xxx.go
func main() {
	//
	fmt.Println("hello, go")
	//
	// fmt.Printf() 也可以用来输出
	// 第一个参数，必须是字符串，可以包含 %v 这样的格式化，动词，它的值由第二个参数替代。
	// 如果指定了多个格式化动词，那么他们的值由后边的参数按顺序替换
	fmt.Printf("My weight on the surface of Mars is %v lbs.", 149.0*0.3783)
	fmt.Printf("and I would be %v years old.\n", 41*365/687)

	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 149.0)

	// 格式化动词 可以指定宽度
	// %3v  %-3v   正数，向左填充空格；负数，向右填充空格

	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
