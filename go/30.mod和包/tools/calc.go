package tools

import (
	"fmt"
)

func PrintInfo() {
	fmt.Println("打印信息")
}

func SoryIntAsc(s []int) []int {
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		for j := i + 1; j < sLen; j++ {
			if s[i] > s[j] {
				tmp := s[i]
				s[i] = s[j]
				s[j] = tmp
			}
		}
	}
	return s
}

type compareFn func(a int, b int) bool

func Sort(s []int, compare compareFn) []int {
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		for j := i + 1; j < sLen; j++ {
			if compare(s[i], s[j]) {
				tmp := s[i]
				s[i] = s[j]
				s[j] = tmp
			}
		}
	}
	return s
}

func init() {
	fmt.Println("tools init")
}
