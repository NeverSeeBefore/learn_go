package main

import "fmt"

// 在go 中 没有 继承
// 只有组合和转发
// 继承可以实现的功能 都可以通过组合和转发实现，且更灵活方便
// go 通过结构体 实现组合(composition)
// go 提供 嵌入(embedding) 特性，它可以实现方法的转发(forwarding)

type cielsius float64

type temperature struct {
	high, low cielsius
}

type location struct {
	lat, lng float64
}

// 结构体嵌入
// 如果嵌入的结构体上有方法、字段都会被，则方法会直接转发到report上
// 内置类型也可以嵌入，但是不建议这样做
type report struct {
	// 火星上的日期
	// int
	sol int
	// 最高温最低温
	temperature
	// 纬度 经度
	location
}

// 计算平均温度
func (t temperature) average() cielsius {
	return (t.high + t.low) / 2
}

// 方法转发
// func (r report) average() cielsius {
// 	return r.temperature.average()
// }

func main() {
	bradbury := location{lat: -4.5895, lng: 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{
		sol:         15,
		temperature: t,
		location:    bradbury,
	}

	fmt.Printf("%+v \n", report)
	// 访问被转发的方法
	fmt.Printf("average: %v \n", report.average())
	// 访问被转发的字段
	fmt.Printf("high: %v %v\n", report.temperature.high, report.high)

	// 如果 location 和 temterature 上都有 fn1, 则 不能直接通过 report.fn1 调用
	// 如果非要 report.fn1 则需要 主动转发，指定要调用的方法
}

// 最佳实践
// 1. 优先使用对象组合 而不是 类的继承
// 2. 对传统的继承，他不是必须的:所有使用继承能够解决的问题，都可以通过其他方式解决（比如组合）
