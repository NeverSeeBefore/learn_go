package main

import "fmt"

// 接口嵌套
type Walker interface {
	walk()
}

type Caller interface {
	setName(string)
	call()
}
type Animal interface {
	Walker
	Caller
}

type Dog struct {
	name string
}

func (d Dog) walk() {
	fmt.Println("step step step...")
}
func (d Dog) call() {
	fmt.Println(d.name, ": wang wang")
}
func (d *Dog) setName(name string) {
	d.name = name
}

func main() {
	var dog = &Dog{"egg"}
	var w Walker = dog
	w.walk()

	var c Caller = dog
	c.call()

	var a Animal = dog
	a.setName("dog egg")
	a.call()

	dog.call()
}
