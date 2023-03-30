package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 1.一个数可能很大，甚至超出uint64的范围
	// 此时可以用float64类型
	// 通过指数形式赋值的变量，默认是float64类型的
	var distance = 24e18
	fmt.Println(distance)

	// 2. 大数更好的选择是使用big包
	// 超过10e18的数  big.Int
	// 任意精度的浮点数 big.Float
	// 分数	big.Rat
	var lightSpeed = big.NewInt(299792)
	fmt.Println(lightSpeed)
	// 另一种得到bigInt的类型变量的方法
	var lightSpeed2 = new(big.Int)
	fmt.Println(lightSpeed2)
	lightSpeed2.SetString("2400000000000000000000000000000000", 10) // 将 10进制 数字字符串 转为 bigInt
	fmt.Println(lightSpeed2)

	// 2. 常量可以不指定类型，也可以指定类型
	// const distance2 = 2400000000000000000000000000000000
	// fmt.Println(distance2)
	// 上边代码会报错，因为赋值数字时，distabce默认是int型，
	// const distance2 = 2400000000000000000000000000000000 / 1000000000000000
	// fmt.Println(distance2)
	// 上边代码不会报错，因为，go中的字面值和常量是在编译阶段完成的。运行时已经符合int的范围了。
}
