package main

import "fmt"

func main() {

	// 1. 切片
	// 数组[起始索引:结束索引]
	arr := [8]string{"0", "1", "2", "3", "4", "5", "6", "7"}

	// 得到 arr的切片，切片包含
	var subArr = arr[0:4]
	var subArr2 = arr[4:6]

	fmt.Println(subArr)  // [0 1 2 3]
	fmt.Println(subArr2) // [4 5]

	// 如果忽略起始索引，则从0开始；忽略结束索引，则切片到最后位置
	var subArr3 = arr[:]
	var subArr4 = arr[:3]
	var subArr5 = arr[3:]

	fmt.Println(subArr3) // [0 1 2 3 4 5 6 7]
	fmt.Println(subArr4) // [0 1 2]
	fmt.Println(subArr5) // [3 4 5 6 7]

	// 2. 切分数组的语法 可以用来切分字符串
	// 切分字符串时，索引是字节数(第几个字节)
	// (一个字符，可能占一个字节、两个字节、三个字节等)
	str := "ab12你好"
	subStr := str[1:5]
	fmt.Println(subStr) // b12�
	subStr2 := str[1:6]
	fmt.Println(subStr2) // b12�
	subStr3 := str[1:7]
	fmt.Println(subStr3) // b12你

	// 3. 字符串切片，不会随着字符串变化，（大概是因为 字符串的不可变性）
	// 数组切片 得到的是 原数组的引用， 切片会随着数组的改变而变化
	str = "1好2坏qwert"
	fmt.Println(subStr, subStr2, subStr3) // 并没有随着str改变
	arr[0] = "a"
	arr[1] = "b"
	arr[2] = "c"
	fmt.Println(subArr, subArr2, subArr3, subArr4, subArr5) // 随着 arr 改变

	// 4. 直接声明切片
	// 声明切片 按照声明数组的方式，但不需要传数组的长度
	// slice := []string{"a", "b", "c"}
	// 实际上，对于这种直接声明的方式，go自动会先创建数组，在得到其全尺寸切片

	// 5. %T 输出Type
	fmt.Printf("%T \t %T \n", arr, subArr) // [8]string        []string
	fmt.Printf("%T \t %T \n", str, subStr) // string   string

	// 6. 更大的切片

	// 实际项目中需要 可变长度的数组
	// 内置函数 append， 将指定元素添加到切片中
	slice := []string{"Earth", "Mars", "Venus"}
	slice = append(slice, "Orcus", "salacia")

	// 7. 长度 和 容量
	// 切片的元素的个数 即 切片的长度 len
	// 切片底层的数组大小 即 切片的容量 cap

	dump := func(label string, slice []string) {
		fmt.Printf("%v: lenght: %v, capacity %v, %v\n", label, len(slice), cap(slice), slice)
	}

	dump("slice", slice)            // slice: lenght: 5, capacity 6, [Earth Mars Venus Orcus salacia]
	dump("dwarfs[1:2]", slice[1:2]) // dwarfs[1:2]: lenght: 1, capacity 5, [Mars]
	// 为什么容量时 6 和 5 ?
	// 当底层数组的容量不够时，会新建一个更大的数组，将内容放进去

	// 8. 验证 扩容会使用新得数组
	// 结论：1、2、3 互相独立，3、4 同步变化
	slice1 := []string{"Mars"}
	dump("slice1", slice1)

	slice2 := append(slice1, "Mars", "Venus")
	dump("slice2", slice2)

	slice3 := append(slice2, "Orcus", "salacia")
	dump("slice3", slice3)

	slice4 := append(slice3, "Venus")

	fmt.Println()

	fmt.Print("replace slice1[0] to \"A\"\n")
	slice1[0] = "A"
	dump("slice1", slice1)
	dump("slice2", slice2)
	dump("slice3", slice3)
	dump("slice4", slice4)
	fmt.Println()

	fmt.Print("replace slice2[0] to \"B\"\n")
	slice2[0] = "B"
	dump("slice1", slice1)
	dump("slice2", slice2)
	dump("slice3", slice3)
	dump("slice4", slice4)
	fmt.Println()

	fmt.Print("replace slice3[0] to \"C\"\n")
	slice3[0] = "C"
	dump("slice1", slice1)
	dump("slice2", slice2)
	dump("slice3", slice3)
	dump("slice4", slice4)
	fmt.Println()

	fmt.Print("replace slice3[0] to \"D\"\n")
	slice4[0] = "D"
	dump("slice1", slice1)
	dump("slice2", slice2)
	dump("slice3", slice3)
	dump("slice4", slice4)
	fmt.Println()

	// 9. 三索引切片操作  限制新建切片的容量
	// 用于解决： 源自子切片的修改 影响到了父切片或数组
	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	// 前5个元素的切片
	subNums := nums[0:5]
	// subNums := nums[0:5:5]
	// 向切片中增加一个元素
	subNums2 := append(subNums, "10")
	// 导致 父元素第6项的值被覆盖
	dump("nums", nums[:])
	dump("subNums", subNums)
	dump("subNums2", subNums2)

	// 10. make 函数，预分配空间
	// make(切片类型, 切片长度, 切片容量)
	// 第三个参数不传时 第二个参数 即表示长度也表示容量
	numberSlice := make([]string, 0, 8)
	numberSlice2 := make([]string, 8)
	dump("numberSlice", numberSlice)
	dump("numberSlice2", numberSlice2)

	// 11. 声明可变参数的函数
	// ...string
	fmt.Println(terraform("new", "Venus", "earth"))
	// 展开的方式 传递参数
	// nums...
	fmt.Println(terraform("new", nums...))
}

// 最后的参数 类型上加 ...
func terraform(prefix string, words ...string) []string {
	newWords := make([]string, len(words))
	for i := range words {
		newWords[i] = prefix + " " + words[i]
	}
	return newWords
}
