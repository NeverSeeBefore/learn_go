package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now:", now)
	// 占位符
	// 2016 年
	// 01	月
	// 02	日
	// 03	时
	// 04	分
	// 05	秒
	fmt.Println("now formatted 12:", now.Format("2006-01-02 03-04-05"))
	fmt.Println("now formatted 24:", now.Format("2006-01-02 15-04-05"))

	timeObj, error := time.ParseInLocation("2006-01-02 15:04:05", "2020-04-06 15:38:20", time.Local)
	fmt.Println(timeObj, error)
	fmt.Println(timeObj.Year())
}
