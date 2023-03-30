package main

import "fmt"

// nil 是 各种类型的零值
// nil 会导致 panic

type person struct {
	age int
}

func (p *person) birthday() {
	if p == nil {
		return
	}
	p.age++
}
func main() {
	var nobody *person
	fmt.Println(nobody)
	nobody.birthday()

	// 定义一个函数变量，但是没有赋值
	var fn func(a, b int) int
	fmt.Println("fn == nil", fn == nil)

	// slice 声明后，没有使用字面值初始化或没有使用make分配空间，其值为nil
	// rang len append 等一些内置函数 可以正常处理 nil
	// 对值 为 nil 的 map 取值 不会报错
	// 对值 为 nil 的 map 赋值 会报错
	var soup map[string]string
	fmt.Println("soup == nil", soup == nil)
	onion, ok := soup["union"]
	fmt.Println("onion", onion)
	fmt.Println("ok", ok)
	// soup["union"] = "aaa" // 报错

	// 接口
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil) // <nil> <nil> true
	// 当 接口类型变量被赋值后，接口就会在内部指向该类型和值
	// 接口类型的变量，只有在类型和值都为nil的时候才为nil
	fmt.Printf("%#v\n", v) // <nil>
	var p *int
	v = p
	fmt.Printf("%#v\n", v) // (*int)(nil)
}
