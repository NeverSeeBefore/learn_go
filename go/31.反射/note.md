1.  reflect.TypeOf 获取变量类型
    返回值 t
    reflect.TypeOf(x).Name() 类型名称（自定义类型名称）
    reflect.TypeOf(x).Kind 类型种类(底层类型)
2.  reflect.ValueOf 获取变量值
    返回空接口类型变量的值
    reflect.ValueOf(x).kind() = reflect.TypeOf(x).Kind(), 但缺少.Name()
3.  使用 reflect.ValueOf 进行断言

    ```go
        switch t.Kind() {
        case reflect.Int:
        	fmt.Println("x是int", v.Int()+1)
        	break
        case reflect.Float64:
        	fmt.Println("x是float64", v.Float()+10.3)
        case reflect.String:
        	fmt.Println("x是string", v.String()+"--")
        }
    ```

4.  使用反射修改变量值

    ```go
        // x 需要是 int64 指针
        func setValue(x interface{}, newV int64) {
         	var v = reflect.ValueOf(x)
         	// var t = reflect.TypeOf(x)

         	fmt.Println("修改完成", v.Elem().Kind())
         	if v.Elem().Kind() == reflect.Int64 {
         		v.Elem().SetInt(newV)
         	}
         }
    ```

5.  isNil() 和 isValid()
    isNil() 是否为 nil
    isValid() 是否为有效值（非零值）

6.  reflect.TypeOf(x).Kind() == reflect.Struct
    是否为结构体

7.  结构体反射

获取属性
t.Field(index).
t.FieldByName(fnName)
t.NumField()

修改属性
v.Elem().FieldByName(fdName).SetString(newValue)

获取方法
t.Method(index)
t.MethodByName(mdName)
t.NumMethod()

执行方法
v.Method(index).Call(params)
v.MethodByName(mdName).Call(params)

只有指着类型参数才可以修改值，只有指针类型才能获取 v.Elem()，指针类型变量 t.Kind() == reflect.Ptr
