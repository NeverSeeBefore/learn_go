package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func birthdayWithPointer(p *person) {
	p.age++
}

// 不接受指针参数
func birthdayNoPointer(p person) {
	p.age++
}

// 方法设置接收指针类型参数
func (p *person) birthday() {
	p.age++
}

func main() {
	// 8. 实现修改

	// go 中 函数的参数是按值传递的,想要函数改变穿进去的参数，需要传递参数的指针
	timmy := person{name: "timmy", age: 27}
	fmt.Printf("timmy: %+v\n", timmy) // timmy: {name:timmy age:27}
	birthdayWithPointer(&timmy)
	fmt.Printf("timmy: %+v\n", timmy) // timmy: {name:timmy age:28}
	birthdayNoPointer(timmy)
	fmt.Printf("timmy: %+v\n", timmy) // timmy: {name:timmy age:28}
	// 这里 timmy不是指针类型，但是因为birthday的接收者是指针，所以go会自动将timmy的指针传入
	timmy.birthday()
	fmt.Printf("timmy: %+v\n", timmy) // timmy: {name:timmy age:29}

	// 使用指针类型作为接收者的策略应该始终如一
	// 如果一种类型上的某些方法需要使用指针作为接收者，就应该为这种类型的所有方法指定指针作为接收者

	// 9. 内部指针
	// 即 结构体 某个字段的内存地址
	// &personal.name
	fmt.Println("timmy", &timmy.name, *&timmy.name)

	// 10. map 是一种隐式的指针，被赋值或被参数传递时就在暗中使用指针
	// slice 也使用了指针，每次slice的内部，都会包含三个元素的结构，数组的指针，切片的容量，切片的长度

	// 11. 指针和接口
	// 如果一个方法的接收者 不是指针， 那参数即可以传参数，也可以传指向参数的指针
	shout(martain{})  // NACK NACK
	shout(&martain{}) // NACK NACK
	// shout(leser(2))
	pew := leser(2)
	// shout(pew) // 报错
	shout(&pew) // 不报错

}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martain struct{}

func (m martain) talk() string {
	return "nack nack"
}

type leser int

func (l *leser) talk() string {
	return strings.Repeat("pew ", int(*l))
}
