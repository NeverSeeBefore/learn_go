package main

import "fmt"

func main() {
	p := Phone{"HUAWEI"}

	p.start()

	// 实现了Usber接口的Phone变量，可以赋值给Usber变量
	var u1 Usber
	u1 = p
	u1.stop()

	c := Camera{"SONY"}
	var u2 Usber = c
	u2.start()
	u2.stop()
}

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "started")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "stopped")
}

type Camera struct {
	Name string
}

func (c Camera) start() {
	fmt.Println(c.Name, "started")
}

func (c Camera) stop() {
	fmt.Println(c.Name, "stopped")
}
