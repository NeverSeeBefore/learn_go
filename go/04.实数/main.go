package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. 当定义一个变量，值包含小数时，其类型就是float64
	var answer float64 = 365.2425
	fmt.Println(answer)

	// float64	双精度	64位浮点类型,占8字节类型，某些语言叫做double
	// float32	单精度	32位浮类型,占四字节,float 32 是用来节省空间的

	// answer2 := 365
	// 2. 当定义一个数 不包含小数时，默认是 int

	// 3. 主动转换为 float64

	var pi64 = math.Pi
	var pi32 float32 = math.Pi
	var pi32_2 = float32(math.Pi)

	fmt.Println(pi64)   // 3.141592653589793
	fmt.Println(pi32)   // 3.1415927
	fmt.Println(pi32_2) // 3.1415927

	// 4. 每个类型都有一个默认值，被称作 零值
	// 声明变量却没有初始化时，它的值就是零值
	var price float64
	fmt.Println(price) // 0

	// 5. print println 会尽可能多的展示小数
	// 如果需要控制展示的位数，可以通过 printf 控制

	var third = 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)    // %v 通用占位符
	fmt.Printf("%f\n", third)    // %f 浮点型变量占位符
	fmt.Printf("%.3f\n", third)  // %f 浮点型变量占位符 	3位小数
	fmt.Printf("%5.2f\n", third) // %f 浮点型变量占位符		2位小数， 正数 + . + 小数 5位，未到5位时 填充空格

	// 6. 浮点类型的精度, 不适用于金融类计算
	fmt.Println(third + third + third) // 1
	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)                        // 0.30000000000000004
	fmt.Println(piggyBank == 0.3)                 // false
	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001) // true 临时解决办法

	// 7. 为避免最小化舍入错误，应先做乘法，在做除法
	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, "° F\n") // 69.80000000000001° F
	fmt.Print((9.0/5.0*celsius)+32, "° F\n") // 69.80000000000001° F
	fmt.Print((celsius*9.0/5.0)+32, "° F\n") // 69.8° F

}
