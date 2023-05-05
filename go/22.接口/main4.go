package main

import "fmt"

type Phone struct {
}

// 空接口 标识任意类型
type Any interface{}

func main() {

	// 任意类型 可以赋值给空接口变量
	var a Any = 20
	show(a)

	a = "你好"
	show(a)

	a = true
	show(a)

	a = []Any{1, "2", true, 4, 5}
	show(a)

	// person := make(map[string]Any)
	person := make(map[string]Any)
	person["username"] = "Bob"
	person["age"] = 20
	person["female"] = true
	a = person
	show(a)

	// 类型断言
	// 变量值，ok := 变量.(类型)
	// 如果断言成功，则 ok 为 nil
	a = "你好"
	v, ok := a.(string)
	if !ok {
		fmt.Println("a is not string")
	} else {
		fmt.Println("a is string", v)
	}

}

// 有了空接口，函数可以接收任意类型参数
// func show(parmas ...interface{}) {}

func show(v Any) {

	fmt.Printf("%v %T, type: %v type: %v\n", v, v, typeOf(v), TypeOf(v))
}

func typeOf(v interface{}) string {
	if _, ok := v.(int); ok {
		return "int"
	}
	if _, ok := v.(string); ok {
		return "string"
	}
	if _, ok := v.(bool); ok {
		return "bool"
	}
	if _, ok := v.([]Any); ok {
		return "slice"
	}
	// 这个不好使，估计时因 v 是 map[string]Any 类型,不是 map[Any]Any 类型
	if _, ok := v.(map[string]Any); ok {
		return "map"
	}
	if _, ok := v.(Phone); ok {
		return "phone"
	}
	return "unknown"
	// return fmt.Sprintf("%T", v)
}

// 通过 x.(type) 判断类型，只能用在 switch-case 中
func TypeOf(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case []interface{}:
		return "slice"
	default:
		return "unknown"
	}
	// return fmt.Sprintf("%T", v)
}
