package main

import "fmt"

func main() {
	// map

	// 1. 声明map 必须指定 key 、value 的类型
	// var temperature map[string]int

	// 符合字面值
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	fmt.Printf("%v %T \n", temperature, temperature)

	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Printf("%v\n", temperature)

	// 访问不存在的属性，得到的时对应类型的零值
	moon := temperature["moon"]
	fmt.Println("moon : ", moon, temperature["moon"])

	// ,ok 写法
	// 如果 temperature 存在 moon， 则 ok 为 true
	// 如果 temperature 不存在 moon， 则 moon2 为 零值，ok 为 false
	if moon2, ok := temperature["moon"]; ok {
		fmt.Println("moon temperature is :", moon2, ok)
	} else {
		fmt.Println("where is moon", moon2, ok)
	}

	// 主动赋值为 0，ok 为 true
	temperature["moon"] = 0
	if moon2, ok := temperature["moon"]; ok {
		fmt.Println("moon temperature is :", moon2, ok)
	} else {
		fmt.Println("where is moon", moon2, ok)
	}

	// 2. map 传递给新变量、用作参数传递， 不会被复制。

	// 3. delete(map, "key")
	// 删除map中的key
	delete(temperature, "moon")

	// 4. 除非使用 符合字面值 初始化map，否则必须使用make分配内存空间
	// make(map类型,  key数量)，第二个参数的作用是，未指定数量的key分配内存空间
	temperature2 := make(map[float64]int, 8)
	fmt.Println("len", len(temperature2)) // 0

	// 5. map 计数
	// test1.go

	// 6. range 遍历map 顺序时无法保证的
	for k, v := range temperature {
		fmt.Println(k, v)
	}

	// 7. map 分组
	// test2.go

	// 8. 用 map 实现 set
	// (go中没有set)
	// test3.go

}
