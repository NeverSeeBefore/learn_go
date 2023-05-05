package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

// fmt 包定义了 Stringer 接口
// type Stringer interface {
// 	String() string
// }

func (c coordinate) String() string {
	return fmt.Sprintf("%v° %v'%v''", c.d, c.m, c.s)
}

func main() {
	var lat coordinate = coordinate{4, 30, 0.0, 'N'}
	var lng coordinate = coordinate{135, 54, 0.0, 'E'}

	// fmt.Printf("Elysium Planitia is at %+v, %+v\n", lat, lng)
	fmt.Println("Elysium Planitia is at", lat, ",", lng)
	// 实现接口前
	// Elysium Planitia is at {4 30 0 78} , {135 54 0 69}
	// 实现接口后
	// Elysium Planitia is at 4° 30'0'' , 135° 54'0''
}
