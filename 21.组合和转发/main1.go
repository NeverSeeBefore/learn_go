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

// 结构体组合
type report struct {
	// 火星上的日期
	sol int
	// 最高温最低温
	temperature temperature
	// 纬度 经度
	location location
}

// 计算平均温度
func (t temperature) average() cielsius {
	return (t.high + t.low) / 2
}

// 方法转发（主动转发）
func (r report) average() cielsius {
	return r.temperature.average()
}

func main() {
	bradbury := location{lat: -4.5895, lng: 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{
		sol:         15,
		temperature: t,
		location:    bradbury,
	}

	fmt.Printf("%+v \n", report)
}
