package main

import (
	"fmt"
	"reflect"
)

// 结构体反射

type Student struct {
	Name  string `json:"name" form:"username"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str string = fmt.Sprintf("姓名：%v，年龄：%v，分数：%v。", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Println("这是一个打印方法")
}

// 通过结构体反射获取结构体中的字段
func PrintStructField(x interface{}) {
	// 判断参数是不是结构体类型
	t := reflect.TypeOf(x)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("PrintStructField err: not a struct")
		return
	}

	// 1. 通过类型变量里的 Field 获取结构体字段
	fmt.Println("通过类型变量里的 Field 获取结构体字段")
	field0 := t.Field(0)
	// field0 := t.Elem().Field(0)
	// {
	// 	Name:"Name",
	// 	PkgPath:"",
	// 	Type:(*reflect.rtype)(0x100957b00),
	// 	Tag:"json:\"name\"",
	// 	Offset:0x0,
	// 	Index:[]int{0},
	// 	Anonymous:false
	// }
	fmt.Printf("%#v\n", field0)
	fmt.Println("字段名称：", field0.Name)
	fmt.Println("字段类型：", field0.Type)
	fmt.Println("字段Tag json：", field0.Tag.Get("json"))
	fmt.Println("字段Tag form：", field0.Tag.Get("form"))

	// 2.通过类型变量里的 fieldByName 获取结构体字段
	fmt.Println("通过类型变量里的 FieldByName 获取结构体字段")
	field1, ok := t.FieldByName("Age")
	if !ok {
		fmt.Println("t.FieldByName(Age); 失败")
	} else {
		fmt.Println("字段名称：", field1.Name)
		fmt.Println("字段类型：", field1.Type)
		fmt.Println("字段Tag json：", field1.Tag.Get("json"))
		fmt.Println("字段Tag form：", field1.Tag.Get("form"))
	}

	// 3. 通过类型变量里的 NumField 获取结构体有哪些字段
	fmt.Println("通过类型变量里的 NumField 获取结构体有哪些字段")
	num := t.NumField()
	fmt.Printf("结构体下有%d个字段\n", num)

	// 4. 遍历所有属性
	fmt.Println("遍历所有属性")
	for i := 0; i < num; i++ {
		fmt.Printf("name: %v, type: %v, json: %v\n", t.Field(i).Name, t.Field(i).Type, t.Field(i).Tag.Get("json"))
	}
}

func PrintStructFn(x interface{}) {
	// 判断参数是不是结构体类型
	t := reflect.TypeOf(x)
	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("PrintStructField err: not a struct")
		return
	}
	// 1. 通过 Method(i) 获取结构体上的方法
	method0 := t.Method(0) // 函数的顺序是按照ASCII值相关，与定义顺序无关
	fmt.Println("method0 函数名称", method0.Name)
	fmt.Println("method0 函数类型", method0.Type)
	fmt.Println()

	// 2. 通过 MethodByName(string) 获取结构体上的方法
	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println("method1 函数名称", method1.Name)
		fmt.Println("method1 函数类型", method1.Type)
	} else {
		fmt.Println("Print is not x's method")
	}

	//3. 通过 NumMethod 获取结构体上方法的数量
	num := t.NumMethod()
	fmt.Printf("结构体下有%d个方法\n", num)

	// 4. 执行方法
	v := reflect.ValueOf(x)
	// 4.1
	v.Method(1).Call(nil)
	// 4.2
	v.MethodByName("Print").Call(nil)

	// 执行 GetInfo
	fmt.Println("run GetInfo")
	info := v.MethodByName("GetInfo").Call(nil)
	fmt.Printf("MethodByName(\"GetInfo\").Call(nil) returns :\n%v %T\n", info, info)
	fmt.Printf("len(info): %d, cap(info): %d\n", len(info), cap(info))

	// 执行 SetInfo
	fmt.Println("run SetInfo")
	var params []reflect.Value
	params = append(params, reflect.ValueOf("小刚"))
	params = append(params, reflect.ValueOf(15))
	params = append(params, reflect.ValueOf(76))
	v.MethodByName("SetInfo").Call(params)

	// 执行 GetInfo
	info = v.MethodByName("GetInfo").Call(nil)
	fmt.Printf("MethodByName(\"GetInfo\").Call(nil) returns :\n%v %T\n", info, info)
	fmt.Printf("len(info): %d, cap(info): %d\n", len(info), cap(info))
}

func ReflectChangeField(x interface{}) {
	// 判断参数是不是结构体类型
	t := reflect.TypeOf(x)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		fmt.Println("PrintStructField err: not a struct pointer")
		return
	}

	v := reflect.ValueOf(x)
	name := v.Elem().FieldByName("Name")
	fmt.Println("beforeChange", x)
	name.SetString("小智")
	fmt.Println("afterChange", x)

}

func main() {
	s := Student{"小明", 12, 88}
	PrintStructField(s)
	fmt.Println()
	PrintStructFn(&s)

	ReflectChangeField(&s)
}
