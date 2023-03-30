package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// struct
	// 将分类零件组成一个完整的结构体

	// 1. struct 变量
	// var 类型名 struct
	// var curiosity struct {
	// 	lat  float64
	// 	long float64
	// }

	// curiosity.lat = -4.5895
	// curiosity.long = 137.4417

	// fmt.Println(curiosity)

	// 2. struct 类型
	// type location struct {
	// 	lat  float64
	// 	long float64
	// }
	// 简写
	type location struct {
		lat, long float64
	}

	// 3. 初始化
	// 通过成对k、v进行初始化
	// 缺点是 需要写 key
	// 优点是，初始化时不必关注 结构体属性的顺序，如果有结构体属性没有赋初始值，则为对应零值
	opportunity := location{lat: -1.9462, long: 354.4734}
	// 按字段定义的顺序进行初始化
	// 优点 只需写下字段对应的值
	// 缺点 赋值顺序必须时声明顺序，且必需为所有属性赋值
	spirit := location{-14.5684, 175.472636}

	fmt.Println(spirit, opportunity)
	// %v 只打印 value
	fmt.Printf("%v \n", opportunity)
	// %+v 打印 key: value
	fmt.Printf("%+v \n", opportunity)

	// 4. struct 赋值给 其他变量 或 用作参数传递
	// 相当于将 struct 完整的复制了一遍，互不影响
	bradbury := opportunity
	bradbury.lat = 111

	fmt.Printf("%+v \n", bradbury)
	fmt.Printf("%+v \n", opportunity)

	changeLong := func(l location) {
		l.long = 0.001
	}
	changeLong(bradbury)
	fmt.Println("changeLong")
	fmt.Printf("%+v \n", bradbury)
	fmt.Printf("%+v \n", opportunity)

	// 5. struct 切片
	locations := []location{
		{lat: 1, long: 0.001},
		{lat: 2, long: 0.002},
		{lat: 3, long: 0.003},
	}
	fmt.Printf("location类型的切片 %+v \n", locations)

	// 6. struct 转 json
	// json.Marshal 转化json时 只会关注结构体中 可导出的字段
	// (大写字母开头表示可导出)

	bytes, err := json.Marshal(opportunity)
	exutOnErr(err)
	fmt.Println("小写属性 转换为json", err)
	fmt.Println("bytes", bytes)
	fmt.Println("bytes to string", string(bytes))
	// 打印 {}，因为 struct 中的属性都以 小写开头
	type locat struct {
		Lat  float64
		Long float64
	}
	locat1 := locat{0.1, 2.0}
	bytes, err = json.Marshal(locat1)
	exutOnErr(err)
	fmt.Println("大写属性 转换为json", err)
	fmt.Println("bytes", bytes)
	fmt.Println("bytes to string", string(bytes))

	// 7. struct tag
	// 通过tag 指定 json转化后的 key
	type locatTag struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

}

func exutOnErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
