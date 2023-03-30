package main

import "fmt"

func main() {
	// 凯撒加密算法
	// 每个字母移动 K 位， K相当于密钥
	// 维吉尼亚加密算法
	// 密钥时一个词，假设密钥为 GOLANG
	// 每一个字母移动的位数 根据密钥确定，比如第一、七个字母移动 6位（G）

	var str = "hello world abcdefghjklmnopqrstuvwxyz"
	var key = "GOLANG"
	var keyLen = len(key)
	// var keys = [6, 14, 11, 0, 13, 6]
	// encode
	for i, c := range str {
		if c >= 'a' && c <= 'z' {
			// 计算对应的密钥位置
			indexK := i % keyLen
			// 找到密钥字符 k
			charK := key[indexK]
			// 找到 k 对应的数字
			offset := int8(charK - 'A')
			// 计算新得字符编码
			encodeC := c + rune(offset)
			if encodeC > 'z' {
				encodeC -= 26
			}
			fmt.Printf("%c", encodeC)
		} else {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()

	// decode
	str2 := "nswlb ccclq gpndrlmvukystcaqeyzigwkef"
	for i, c := range str2 {
		if c >= 'a' && c <= 'z' {
			// 计算对应的密钥位置
			indexK := i % keyLen
			// 找到密钥字符 k
			charK := key[indexK]
			// 找到 k 对应的数字
			offset := int8(charK - 'A')
			// 计算新得字符编码
			decodeC := c - rune(offset)
			if decodeC < 'a' {
				decodeC += 26
			}
			fmt.Printf("%c", decodeC)
		} else {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
}
