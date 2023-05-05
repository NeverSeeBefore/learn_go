package main

import (
	"fmt"
)

func main() {
	// 1. 字符串
	// peace := "peace"
	// var peace = "peace"
	var peace string = "peace"
	fmt.Println(peace)

	// 2. 字符串零值是 ""

	// 3. 字面值  和 原始字符串字面值
	fmt.Println("peace be upon you \nupon you be peace")
	fmt.Println(`string can span multiple line with the \n escape sequence`)
	fmt.Println(`
		peace be upon you
		upon you be peace
	`)

	// 4. unicode 联盟为超过一百万个字符分配了不同的 数值，这个数叫做 code point
	// ASCII 是 unicode 的子集，一共可以标识128个字符
	// rune 类型，是 int32 的类型别名 用于 unicode code point
	// byte 类型，是 uint8 的类型别名 用于二进制数据

	// 5. 自定义类型别名
	// type usignint8 = uint8

	// 6. 打印字符，code point
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang rune = 33

	// %v 输出 code point
	fmt.Printf("%v \t %v \t %v \t %v \n", pi, alpha, omega, bang)
	// %c 输出 字符串
	fmt.Printf("%c \t %c \t %c \t %c \n", pi, alpha, omega, bang)
	// 960      940     969     33
	// π        ά       ω       !

	// 7. 字符字面量  使用单引号
	var grade rune = 'A'
	fmt.Printf("%c %d\n", grade, grade)

	// 8. 可以给变量赋值不同的字符串值，但是字符串本身是不可变的，
	// 很多语言都是这样的
	message := "shalom"
	fmt.Printf("%s \n", message)
	c := message[1]
	fmt.Printf("%c \n", c) // h

	// message[1] = 'A'
	// 报错

	// 9. 凯撒加密法 caersar cipher
	// 对加密信息，每个字母都移动固定长度的位置
	// 如  a -> d  b -> e

	var c1 rune = 'y'
	c1 += 3
	if c1 > 'z' {
		c1 -= 26
	}
	fmt.Printf("y move right 3 step is %c \n", c1)

	// 10. c = c - 'a' + 'A'
	// 如果 c 是 'g', c - 'a' + 'A' = ?

	// 11. ROT 13
	// 凯撒加密法的变体，把字母后移13位
	// 加密解密都是一个算法
	str := "uv vagreangvbany fcnpr fgngvba"
	fmt.Println("decode message: ", str)
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()

	// 12. go 内置函数，不需要import
	// len 返回 字符串的字节数

	fmt.Println(len("message"))

	// 13. utf8 是 unicode code point 的编码方式之一
	// utf8 是一种有效率的，可变长度的编码 每个code point 可以是 8位 16位 32 位
	// 由于 可变长度，通过 len 得到的结果很可能跟实际的字符长度不符
	// 此时可以通过 utf-8 这个包,他可以按照rune计算长度

	// utf8.DecodeRuneInString()

	// 14. 使用 range 关键字 可以遍历各种集合

	var str2 = "stringh"
	for i, c := range str2 {
		fmt.Printf("%c \t %d \n", c, i)
	}

	// 如果不需要索引 可以使用 空标识符 【_】 代替
	// for _, c := range str2 {}
}
