package main

import (
	"fmt"
	"math"
)

func main() {

	// go 中没有class
	// 需要通过 结构体、类型、方法 达成类似的功能

	// go 中没有构造函数
	// (new|New) + 类型名 这样命名的方法 一般会用作 构造函数
	// 有个构造函数会 直接命名为 New

	lat := coordite{4, 35, 22.2, 'S'}
	long := coordite{137, 26, 30.12, 'E'}

	loc := NewLocation(lat, long)
	fmt.Printf("%+v\n", loc)

	// 计算两点间距离
	var mars = world{redius: 3389.5}
	var spirit = location{Lat: -14.5684, Long: 175.472632}
	opportunity := location{Lat: -1.9462, Long: 354.4734}
	dist := world.distance(mars, spirit, opportunity)

	fmt.Printf("%.2f km\n", dist)
	fmt.Println()
	fmt.Printf("%.2f km\n", mars.distance(spirit, opportunity))

}

type location struct {
	Lat  float64
	Long float64
}

type coordite struct {
	d, m, s float64
	h       rune
}

func (c coordite) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}

	return sign * (c.d + c.m/60 + c.s/3600)
}

func NewLocation(lat, long coordite) location {

	loc := location{coordite.decimal(lat), coordite.decimal(long)}
	return loc
}

type world struct {
	redius float64
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.Lat))
	s2, c2 := math.Sincos(rad(p2.Lat))
	clong := math.Cos(rad(p1.Long - p2.Long))
	return w.redius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}
