package main

import (
	"fmt"
	"math"
)

// map 分组

func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	groups := make(map[float64][]float64)

	for _, t := range temperatures {
		// 整数部分  math.Trunc
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}

	for k, v := range groups {
		fmt.Printf("%v: %v\n", k, v)
	}

}
