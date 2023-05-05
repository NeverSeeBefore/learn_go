package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. go 有十种整数类型 五种有符号类型，五种无符号类型，通过取值范围确认使用那种
	// int
	// uint
	// int8				-128 to 127			8bit(one byte)
	// uint8			0 to 255			8bit
	// int16								16bit
	// uint16
	// int32
	// uint32
	// int64
	// uint64
	// Tip: int 和 uint 不是 某个类型的别名， 其与系统架构有关
	// int 和 uint 是针对目标设备优化的类型，所以在一些老设备上，int 和 uint 是 32 位的，其他设备是 64 位的

	// uint8 的 范围时 0-255，可以用来表示颜色
	// 用更大的 类型也可以，但会浪费许多内存

	// 2. 数字前面加 0x，可以表示十六进制的数
	// 输出可以使用 %x

	var red, green, blue uint8 = 0, 141, 255
	fmt.Printf("coloris #%02x%02x%02x;\n", red, green, blue)
	// %02x  位数不够填充0，展示位数为2位

	// 3.【环绕】所有整数类型都有取值范围，超出这个范围时就会触发环绕
	var number uint8 = 255
	fmt.Println(number) // 255
	number++
	fmt.Println(number) // 0

	// 4. 输出二进制
	var number1 int8 = 3
	fmt.Printf("number1 %08b\n", number1) // 00000011
	number1++
	fmt.Printf("number1 %08b\n", number1) // 00000100

	// 5. math 为与架构无关的一些类型定义了最大值和最小值
	fmt.Println("max int8 is ", math.MaxInt8)
	fmt.Println("min int8 is ", math.MinInt8)

}
