package main

// 1. 声明函数
// func 函数名(形参 参数类型, ...){返回值 返回值类型} {
// ...
// }

// 2. 函数名如果大写，则会被导出，可被其他包调用
// 小写函数名不会被导出

// 3. 参数类型如果相同，可以只写一个
// func Unix(sec int64, nsec int64)
// func Unix(sec, nsec int64)

// 4. 多个返回值
// func Atoi(s string) (i int, err error)
// 返回值可以去掉名字，只保留类型
// func Atoi(s string) (int, error)

// 5. 可变参数
// func Println(a ...interface{}) (int, error)

// ... 表示可变
// interface{} 空接口，表示任意类型
