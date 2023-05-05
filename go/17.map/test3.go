package main

import (
	"fmt"
	"sort"
)

// map 实现 set
func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	set := make(map[float64]bool, 0)
	for _, t := range temperatures {
		set[t] = true
	}
	fmt.Println(temperatures)
	fmt.Println(set)

	unique := make([]float64, 0, len(set))
	for _t := range set {
		unique = append(unique, _t)
	}
	sort.Float64s(unique)
	fmt.Println(unique)
}
