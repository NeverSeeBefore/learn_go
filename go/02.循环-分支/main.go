package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")

	var command = "walk inside"
	var exit = strings.Contains(command, "outside")

	fmt.Println("You leave the cave:", exit)

	// 比较运算符 == != >= <= > <
	var age = 26
	var minor = age < 18
	fmt.Printf("My age is %v, am I minor? %v\n", age, minor)

	// 条件判断
	if command == "go east" {
		fmt.Println("You head further up the  mountain.")
	} else if command == "walk inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("Didn't quite get that.")
	}

	// 逻辑运算符 || &&
	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	fmt.Printf("The year is %v, is a leap year? %v \n", year, leap)

	// switch
	// 默认情况下，执行完case对应的语句，不会继续往下执行， 只有加了 fallthrough ，才会执行下一个case
	switch command {
	case "go east":
		fmt.Println("You head further up the  mountain.")
	case "walk inside":
		fmt.Println("You enter the cave where you live out the rest of your life.")
		fallthrough
	default:
		fmt.Println("Didn't quite get that.")
	}

	// for
	var count = 10
	for count >= 0 {
		fmt.Println("count", count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff!")

	// for  第二种方式
	count = 10
	for {
		if count < 0 {
			break
		}
		fmt.Println("count", count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Liftoff!")

	// 短声明, 同样适用于switch，if
	for count := 10; count >= 0; count-- {
		fmt.Println("count", count)
		time.Sleep(time.Second)
	}
}
