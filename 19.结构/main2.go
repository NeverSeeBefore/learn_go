package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}

func (p *person) setAge(age int) {
	p.age = age
}

func main() {
	// 初始化结构体

	// 通过 new 创建，得到的时指针
	var p1 = new(person)
	p1.name = "John"
	p1.age = 20
	p1.gender = "Female"

	// 直接字面量创建
	var p2 = person{
		name:   "John",
		age:    20,
		gender: "Female",
	}

	// 创建同时取地址
	var p3 = &person{
		name:   "John",
		age:    20,
		gender: "Female",
	}
	fmt.Printf("%#v\n%#v\n%#v\n", p1, p2, p3)

	fmt.Println()
	p1.setAge(33)
	fmt.Printf("%#v\n%#v\n%#v\n", p1, p2, p3)

}
