package main

import "fmt"

// 编写一个温度转换程序表
//  第一列摄氏度 第二列华氏度
// -40°C到100°C 间隔5°C
// 负责画线和填充值的代码应该都是可复用的,画表格和计算温度应该是不同的方法。
// 实现一个drawTable参数，接收一个一等函数作为参数，

type rowDataFn = func(row int) (float64, float64)
type headerDataFn = func() (string, string)

type celsius float64
type fahrenheit float64
type kelvin float64

func (val celsius) fahrenheit() fahrenheit {
	val = val*9.0/5.0 + 32
	return fahrenheit(val)
}

func (val celsius) kelvin() kelvin {
	val += 263.15
	return kelvin(val)
}

func (val fahrenheit) celsius() celsius {
	val = (val - 32) * 5.0 / 9.0
	return celsius(val)
}

func (val fahrenheit) kelvin() kelvin {
	c := fahrenheit.celsius(val)
	k := celsius.kelvin(c)
	return k
}

func (val kelvin) celsius() celsius {
	val -= 273.15
	c := celsius(val)
	return c
}

func (val kelvin) fahrenheit() fahrenheit {
	c := kelvin.celsius(val)
	f := celsius.fahrenheit(c)
	return f
}

func drawTable(rowCount int, rowDataFn rowDataFn, headerDataFn headerDataFn) {
	drawSplitLine := func() {
		fmt.Print("=================================\n")
	}

	drawSplitLine()
	v1, v2 := headerDataFn()
	fmt.Printf("|\t%v\t|\t%v\t|\n", v1, v2)
	drawSplitLine()

	drawRow := func(row int) {
		v1, v2 := rowDataFn(row)
		fmt.Printf("|\t%.2f\t|\t%.2f\t|\n", v1, v2)
	}

	var index int = 1
	for index <= rowCount {
		drawRow(index)
		index++
	}
	drawSplitLine()
}

func main() {

	const rowCount int = (100-(-40))/5 + 1
	var rowDataFn1 rowDataFn = func(row int) (float64, float64) {
		v1 := -40 + 5*(row-1)
		v2 := celsius.fahrenheit(celsius(v1))
		return float64(v1), float64(v2)
	}
	drawTable(rowCount, rowDataFn1, func() (string, string) { return "°C", "°F" })
}
