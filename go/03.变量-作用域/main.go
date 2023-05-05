package main

import (
	"fmt"
	"math/rand"
)

var ear = "AD"

func main() {
	// 声明变量 var
	// 短声明 count := 10

	// 短声明, 同样适用于switch，if
	// for count := 10; count >= 0; count-- {
	// 	fmt.Println("count", count)
	// 	time.Sleep(time.Second)
	// }

	// 变量的作用域，说声明的 {} 内

	// 展示随机日期

	for count := 0; count < 10; count++ {
		year := rand.Intn(2022) + 1
		month := rand.Intn(12)
		daysInMonth := 31

		var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

		switch month {
		case 2:
			if leap {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}
		case 4, 6, 9, 11:
			daysInMonth = 30
		}

		var day = rand.Intn(daysInMonth) + 1
		fmt.Println(ear, year, month, day)
	}
}
