package main

import "fmt"

// 结构体组合

type address struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	sex  string
}

type user struct {
	// 匿名属性，可以通过 .person. 访问，也可以 . 访问
	// 如果包含结构体各自有同名属性，则不能直接 . 访问
	person
	// 命名属性 不能直接 . 访问
	Address address
}

func main() {
	var user = user{
		person{"john", 18, "female"},
		address{"jilin", "changchung"},
	}

	fmt.Printf("%#v\n", user)
	// fmt.Println(user.person.name,  user.name, user.Address.province, user.)
}
