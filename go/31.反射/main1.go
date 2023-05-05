package main

import (
	"fmt"
	"reflect"
)

type myInt int
type Person struct {
	name string
	age  int
}

func main() {
	a := 10
	b := 23.4
	c := false
	d := "hello world"
	e := 'S'
	var f myInt = 33
	var g = Person{name: "张三", age: 20}
	reflectFn(a)  //int 10
	reflectFn(b)  //float64 23.4
	reflectFn(c)  //bool false
	reflectFn(d)  //string hello world
	reflectFn(e)  //int32 83
	reflectFn(f)  //main.myInt 33
	reflectFn(g)  //main.Person {张三 20}
	reflectFn(&a) //*int 0x14000126008
	reflectFn(&g) //*main.Person &{张三 20}

	fmt.Println("通过反射修改变量值")
	var h int64 = 123
	fmt.Println(h)
	setValue(&h, 321)
	fmt.Println(h)
}

func reflectFn(x interface{}) {
	// 1. reflect.TypeOf 获取变量类型
	// 返回值 t
	// t.Name() 类型名称（自定义类型名称）
	// t.Kind 类型种类(底层类型)
	// 2. reflect.ValueOf 获取变量值
	// 返回空接口类型变量的值
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(t, t.Name(), t.Kind(), v)

	// 3. 使用 reflect.ValueOf 进行断言
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("x是int", v.Int()+1)
		break
	case reflect.Float64:
		fmt.Println("x是float64", v.Float()+10.3)
	case reflect.String:
		fmt.Println("x是string", v.String()+"--")
	}

	fmt.Println()
}

func setValue(x interface{}, newV int64) {
	var v = reflect.ValueOf(x)
	// var t = reflect.TypeOf(x)

	fmt.Println("修改完成", v.Elem().Kind())
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(newV)
	}
}
