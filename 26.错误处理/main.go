package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// go中 函数和方法 可以返回多个值
// 1. 最后一个值应为错误信息，没有错误则为 nil

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// 2. error 内置错误对象
// 3. defer 关键字   保证defer 后边的语句，在函数或方法 返回之前，一定会被调用

// 更优雅的写入处理错误
type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	sw := safeWriter{w: f}

	sw.writeln("Errors are values.")
	sw.writeln("Errors are values.")
	sw.writeln("Errors are values.")
	sw.writeln("Errors are values.")
	sw.writeln("Errors are values.")

	return sw.err
}

// 4. 继承Error类
// 按照惯例，错误类应以Error结尾
// 类库中错误类可能就叫做 【Error】
type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

// 5. 类型断言
// err.(错误类型)

func testA() {
	f, err := os.Create("a")
	if err != nil {
		// 类型断言
		if errs, ok := err.(SudokuError); ok {
			// 如果 err 是 SudokuError 类型，则 err 为 errs, ok 为 true
			// 如果 err 不是 SudokuError ，  则 ok 为 false
			fmt.Println("testA", errs, ok)
		}
	}
	fmt.Println(f)

}

// 6. go 中没有异常，但是有 panic
// 其他语言
// 1. 如果函数发生异常，并且没有补货，则会冒泡到调用者，直到顶部
// 2. 异常处理是可选的
// go 语言
// 1. 错误值简单灵活,忽略错误时开发者有意识的决定，并且错误是显而易见的

// 7. 主动创建panic("i forgot my towel")
// 	panic 参数可以是任意类型
// 8. 优先使用 os.Exit ,其次使用panic
// panic后 会执行所有defer动作，os.Exit 则不会

// 9. 有时候go会panic 而不是返回错误值
// 比如 除以0
func test2() {
	var zero int
	_ = 42 / zero
}

// go中发生panic的场景：
// - 数组/切片越界
// - 空指针调用。比如访问一个 nil 结构体指针的成员
// - 过早关闭 HTTP 响应体
// - 除以 0
// - 向已经关闭的 channel 发送消息
// - 重复关闭 channel
// - 关闭未初始化的 channel
// - 未初始化 map。注意访问 map 不存在的 key 不会 panic，而是返回 map 类型对应的零值，但是不能直接赋值
// - 跨协程的 panic 处理
// - sync 计数为负数。
// - 类型断言不匹配。`var a interface{} = 1; fmt.Println(a.(string))` 会 panic，建议用 `s,ok := a.(string)`

// 10. 什么时候进行panic
// https://www.dovov.com/golangos-exitpanic.html
