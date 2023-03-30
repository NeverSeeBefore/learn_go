package main

import "fmt"

// 1. 声明新类型
type celsius float64
type kelvin float64

func main() {

	// type celsius float64
	// 2. 声明类型提高代码的可读性
	const degrees = 20
	var temperature celsius = degrees
	temperature += 10
	fmt.Println(temperature)
}

// 3. 可以将方法与同包中的任意类型关联
// 此时 相当于 kelvin 上居右一个方法，并且方法接收一个参数 k
func (k kelvin) celsius() {}

// 4. 声明方法
// func (接受者 接受者类型) 方法名（参数 参数类型）返回值 {}

// 5. 调用
// kelvin.方法名(接受者, 参数)
