package main

import "fmt"

func main() {
	// 指针 是 指向另一个变量地址的变量

	// 1. 迷途指针（Dangling pointer）
	// go 中 不会出现 迷途指针

	// 2. & 地址操作符

	var answer int = 42
	fmt.Println(&answer) // 0x14000014090

	// 无法获取字符串、数值、布尔类型的地址
	// 会报错

	// 3. * 解引用，与 & 功能相反，提供内存地址指向的值
	// * 如果后边跟一个类型，则 表示 指向该类型的一个地址
	// * 如果后边跟一个变量，则 表示 获取该变量的地址

	var address = &answer
	fmt.Println(address, *address, *&answer) // 0x14000014090 42 42

	// 可以直接 *address 改变 answer 的 值，地址时不会变化的
	*address++
	fmt.Println(address, *address, *&answer) // 0x14000014090 43 43

	// go 不允许 address++ 这样的不安全操作

	// 4. 输出地址变量的类型
	fmt.Printf("address is a %T\n", address) // address is a *int
	// *int 表示 指向int类型的指针

	// 5. -
	var canada string = "Canada"

	var home *string
	fmt.Printf("home is a %T\n", home)

	home = &canada
	fmt.Println(home, *home)

	// 6. -
	type personal struct {
		name string
		age  int
	}

	var timmy = &personal{"timmy", 25}
	// 访问字段时，解引用不是必须的
	fmt.Println(timmy, timmy.name, (*timmy).name, *&timmy.name)
	timmy.name = "Timmy"
	fmt.Println(timmy, timmy.name, (*timmy).name, *&timmy.name)

	// 7. 数组 可以自动解引用，但是string、map不能自动解引用
	var arr = &[3]string{"A", "B", "C"}
	fmt.Println(arr, arr[0], (*arr)[1]) // &[A B C] A B
	// var testSlice = &[]string{"A", "B", "C"}
	// fmt.Println(testSlice, testSlice[0]) // 报错

}
