package main

import "fmt"

func main() {

	// 1. 数组 元素的集合，长度固定
	// 定义数组
	// var 变量名 [长度]类型
	var arr [8]string

	// 赋值
	// 未赋值，则为对应类型的零值
	arr[0] = "aaa"

	// 取值
	fmt.Println(arr, arr[0], arr[1] == "")

	// 2. 数组越界
	// arr[8] // 编译时报错
	// for i := 0; i < 9; i++ {
	// 	fmt.Println(arr[i])
	// }
	// 运行时 panic 报错
	// panic: runtime error: index out of range [8] with length 8

	// 3. 复合字面值 初始化数组
	// arr2 := [8]string{"1", "2", "3", "4", "5", "6", "7", "8"}
	// 可以使用 ... 替代 8，go编译器 会自动计算出 8
	// arr3 := [...]string{"1", "2", "3", "4", "5", "6", "7", "8"}

	// 4. 遍历数组
	dwarfs := [...]string{"Ceres", "Ploto", "Haumea", "Makemake", "Eris"}
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}

	// 5. 数组复制
	// 无论数组是复值给变量，还是传递给函数，都会产生一个数组的拷贝
	// 所以数组作为参数传递时时是低效的
	// 所以 一般 采用 切片（slice） 进行传递

	// 6. 二维数组
	// var arr4 [8][8]string
}
