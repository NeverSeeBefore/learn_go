package main

import "fmt"

func main() {
	// 加密串 前移3位解密
	var message = "L fdph, L vdz,L frqtxhuhg"
	// A: 65
	// a: 97
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'A' && c <= 'z' {
			c -= 3
			if c <= 'A' {
				c += 46
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
	// fmt.Println('A', 'a')
}
