package main

import (
	"fmt"
	"strings"
)

// 接口
// 接口关注 可以做什么

// 1. 定义接口
// 变量方式
//
//	var t interface {
//		talk() string
//	}
//
// 类型方式（为了复用，通常会将接口声明为接口类型）
// 接口类型通常以 er 为结尾
type talker interface {
	talk() string
}

// 2. 实现接口
// go 中 实现接口不是显示的实现，
// 只要有对应接口中的方法即可

type martain struct{}

func (m martain) talk() string {
	return "nack nack"
}

type leser int

func (l leser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type starShip struct {
	leser
}

func main() {
	var m = martain{}
	shout(m)
	shout(leser(2))

	s := starShip{leser: leser(3)}
	fmt.Println(s.talk())
	shout(s)
}
