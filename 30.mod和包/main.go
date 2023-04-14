package main

import (
	"fmt"

	"github.com/shopspring/decimal"

	// 引入自定义包 项目名/包名
	"30.mod-pkg/calc"
	"30.mod-pkg/tools"
)

func main() {

	// 引入自定义包
	sum := calc.Add(10, 2)
	fmt.Println(sum)
	fmt.Println(calc.Sub(10, 2))

	tools.PrintInfo()
	fmt.Println(tools.SoryIntAsc([]int{1, 6, 2, 3, 5, 7}))

	fmt.Println(tools.Sort([]int{1, 6, 2, 3, 5, 7}, func(a int, b int) bool {
		swap := a-b > 0
		return swap
	}))

	// 引入第三方包
	// 已decimal为例
	// 1. 下载
	// go get github.com/shopspring/decimal
	// 另外两种下载方式
	// go mod download 根据go.mod文件下载依赖
	// go mod vendor 在包中直接使用 import 引入，执行 go mod vendor，将包复制到当前项目下的 vendor 文件夹下，如果 GOPATH 中没有则尝试下载
	// go mod tidy 删除未使用的依赖，添加未引入的依赖
	// go mod graph

	// 2. import
	// 3. 使用
	var a float64 = 0.1
	var b float64 = 0.2
	fmt.Println(a + b)
	fmt.Println(decimal.NewFromFloat(a).Add(decimal.NewFromFloat(b)))

}

// init函数
// 每个包都可以定义一个 init 函数，go 在运行时 回一次执行每个包的init 函数
// 如果 a 引入 b， b 引入 c， 则 init 函数的执行顺序为  c -> b -> a
// init 函数不接收任何参数，不返回任何值
func init() {
	fmt.Println("main init")
}
