package main

import "fmt"

// 测试 容量增加规则

func dump(label string, slice []string) {
	fmt.Printf("%v: lenght: %v, capacity %v, %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	i := 100
	lastCapacity := 0
	slice := make([]int, 0)
	currentCapacity := 0
	for i > 0 {
		slice = append(slice, i)
		currentCapacity = cap(slice)
		if currentCapacity != lastCapacity {
			lastCapacity = currentCapacity
			fmt.Println("capacity change", currentCapacity)
		}
		i--
	}
	// capacity change 1
	// capacity change 2
	// capacity change 4
	// capacity change 8
	// capacity change 16
	// capacity change 32
	// capacity change 64
	// capacity change 128
}
