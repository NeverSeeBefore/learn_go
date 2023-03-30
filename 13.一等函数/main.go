package main

import (
	"fmt"
	"math/rand"
)

// 1. 在 go 中，函数是头等的，他可以用作其他整数、字符串或其他类型能用到的地方
// 将函数赋值给变量
// 将函数作为其他函数的参数
// 函数作为其他函数的返回类型

type kelvin float64

func fakeSensor(index int) kelvin {
	return kelvin(rand.Intn(151) + 150 + index)
}

// 传递函数
func measureTemperature(samples int, sensor func(index int) kelvin) {
	sensor(samples)
}

// 2. 定义函数类型 有助于精简代码
type sensorfn func(index int) kelvin

func measureTemperature1(samples int, sensor func(index int) kelvin) {}
func measureTemperature2(samples int, sensor sensorfn)               {}

func main() {

	// 传递到变量
	sensor := fakeSensor
	fmt.Println(sensor(1))

	// 传递到参数
	measureTemperature(1, sensor)

	// 3. 匿名函数 也叫 函数字面值
	f := func(message string) {
		fmt.Print(message)
	}
	f("Go to the party.\n")

	// 4. 立即执行函数
	func() {
		fmt.Print("IIFE\n")
	}()

	// 5. 匿名函数， 都是闭包的

	k := 294
	f2 := func() kelvin {
		return kelvin(k)
	}

	fmt.Println(f2())
	k++
	fmt.Println(f2())
}
