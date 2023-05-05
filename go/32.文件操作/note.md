## 文件操作

### 读取文件

1. file.Read

   - 打开文件
     file, err := os.Open(filepath)
   - 读取指定字节长度的数据
     n, err := file.Read(byteSlice)
   - 关闭文件流
     file.Close()

2. bufio

   - 打开文件
     file, err := os.Open(filepath)
   - 创建 Reader
     reader := bufio.NewReader(file)
   - 读取文件
     str, err := reader.ReadString('\n') // 参数为结束标记
   - 关闭文件流
     file.Close()

3. ioutil
   byteSlice, err := ioutil.ReadFile(filepath)

<!-- 充数内容ABCDEFGHIJKLMNOPQRETUVWXYZabcdefghijklmnopqrstuvwxyz1234567890-=[]\';'/.,,./,≈√∫µµ≤≤≥…¬˚∆˙©ƒ∂ßåœ∑®†¥øπ“‘æ…≥÷«$ -->

### 写入文件

1. file

   - 打开文件
     file, err := os.Open(filepath, os.O_CREATE|os.O_RDWR, 0666) // 文件路径， 创建、读写权限， 0666 为 linux 系统下的权限
     O_CREATE 创建
     O_TRUNC 清空
     O_WRONLY 只写
     O_APPEND 追加
     O_RDWR 可读可写
     多个权限可以使用【|】或运算符
   - 两种写入方式
     file.Write([]byte)
     file.writeString(string)
   - 关闭文件流
     file.Close()

2. ioutil
   ioutil.WriteFile(filePath, string, 0666)

### 复制

    基于流的方式，读一点写一点
    基于文件的方式，读取文件全部内容，写入到另一个文件

### 文件夹

    - 创建文件夹
      err = os.Mkdir("./aaa", 0777)
      err = os.MkdirAll("./aaa/bbb/ccc", 0777)

    - 删除文件夹
      err = os.Remove("./aaa")
      err = os.RemoveAll("./aaa")

    - 重命名文件/文件夹
      err = os.Rename("./aaa", "./xxx")
