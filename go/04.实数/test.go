package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// 1. 将五分（0.05），1角（0.1），25分（0.25）硬币放入一个空的储蓄罐，直到罐内至少20元
	coin1 := 0.05
	coin2 := 0.1
	coin3 := 0.25

	count := 0.0
	fmt.Printf("count start with : %30.20f\n", count)

	for {
		if count > 20 || math.Abs(count-20) < 0.00001 {
			break
		}

		var coinType = rand.Intn(3)

		switch coinType {
		case 0:
			count += coin1
		case 1:
			count += coin2
		case 2:
			count += coin3
		}

		fmt.Printf("add coin%v, now count is %30.20f \n", coinType+1, count)

	}
	fmt.Printf("final count: %30.20f", count)
}
