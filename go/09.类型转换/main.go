package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	age := 41
	marsDays := 687
	earthDays := 365.2425

	fmt.Println("I am ", age*int(earthDays)/marsDays, " years of mars old")

	// 1. 转换 使用目标类型 包裹需要转换的类型
	// 如果浮点类型转换为整型 小数位被截断，而不是舍入

	// 2. 无符号、有符号 整数类型之间需要转换, 不同大小的整数类型也需要转换

	// 3. 『环绕』行为
	var bh float64 = 32768
	var h = int16(bh)
	fmt.Println(h) // -32767
	// 对于int16来说，最大值即32767，不能存储32767以上的数值
	// 可以通过math 中的 max min 判断能否换行
	if bh <= math.MaxInt16 || bh >= math.MinInt16 {
		// ... in range
	}

	// 4. byte rune 转为 string
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang rune = 33
	fmt.Print(string(pi), string(alpha), string(omega), string(bang), "\n")

	// 5. 数值 -> 字符串，它的值必须能转化为 code point，
	fmt.Print(string(65), string(8755), string(432131411412353211414), "\n")
	// 包 strconv 中的 Itoa 可以转换数值为 字符串, Itoa => Integer to ASCII
	fmt.Println("Launch in T minus", strconv.Itoa(10), "seconds")

	// fmt.Sprintf 也可以数值转化为 string
	fmt.Println(fmt.Sprintf("Launch in T minus %v seconds", 9))

	// 6. 字符串 -> 整数
	// nil 相当于 null
	countdown, err := strconv.Atoi("8js")
	if err != nil {
		// fmt.Println("fail", countdown, err)
		fmt.Println("fail", countdown, err.Error())
	} else {
		fmt.Println("success", countdown, err)
	}

	// 7. go 是静态类型，一旦定义，不可改变

	// 8. bool -> string
	fmt.Println("bool to string: ", fmt.Sprintf("%v", true))
	fmt.Println("bool to string: ", fmt.Sprintf("%v", false))

}
