##

- [go 简介](https://zhuanlan.zhihu.com/p/478076646)
- [go 基础知识博客](https://www.seoooooo.com/cms/11280.html#toc_10)
- [官方文档](https://gin-gonic.com/docs/examples/ascii-json/)

## 安装环境

1. 初始化项目
   go mod init project_name
2. 安装 gin
   go get -u github.com/gin-gonic/gin
3. 在项目内引入
   import "github.com/gin-gonic/gin"
4. 编写路由
   ```go
       r := gin.Default()
       r.GET("/news", func(c *gin.Context){
           c.String(200, "hello %v", "world")
       })
       r.Run()
   ```
5. 访问 localhost:8080

## 热加载

1. go install github.com/pilu/fresh
2. go get github.com/pilu/fresh
3. ~/.zshrc
   alias gofresh="~/go/bin/fresh"
4. gofresh 替代原本的 go run main.go

## gin [模板语法](https://www.jianshu.com/p/6c1a1fe1fbc4)

- 取值
  ```html
  {{.}} {{.title}} {{.user.id}}
  ```
- 定义变量
  ```html
  {{ $t := .title}}
  ```
- 去除空格
  ```html
  {{- .title -}}
  ```
- 比较
  ```html
  <!--
        eq  相等
        ne  不相等
        gt  大于
        ge  大于等于
        lt  小于
        le  小于等于
     -->
  {{if ge .score 80}}
  <span>优秀</span>
  {{else if ge .score 60}}
  <span> 及格 </span>
  {{else}}
  <span>不及格</span>
  {{end}}
  ```
- 条件判断
  ```html
  {{if pipeline}} T1 {{end}} {{if pipeline}} T1 {{else}} T0 {{end}} {{if
  pipeline}} T1 {{else if}} T0 {{end}}
  ```
- 循环

  ```html
  <ul>
    {{range $key, $value := .hobby}}
    <li>{{$key}}: {{$value}}</li>
    {{end}}
  </ul>
  ```

- with
  结构结构体

  ```html
  <!-- 
      news := [string]interface{}{
          "title" : "news-title",
          "content" : "news-content",
      } 
  -->
  {{with .news}}
  <span>{{.title}}</span>
  <span>{{.content}}</span>
  {{end}}
  ```

- 预定义模板函数

  ```html
  <div>len</div>
  {{len .title}}
  ```

- 自定义模板函数

  在加载模板前定义模板函数
  r.SetFuncMap(template.FuncMap{
  "UnixToTime": UnixToTime,
  })

  ```html
  <!-- 在模板中使用自定义模板函数 -->
  <div>date:</div>
  {{UnixToTime .date}}
  ```

- 模板复用

  1. 定义一个公用模板 pageHeader

     ```html
     {{define "public/pageHeader.html"}}
     <h1>common : {{.title}}</h1>
     ```

  2. 在需要引入的地方引入

     ```html
     {{template "public/pageHeader.html" .}}
     <div>content:{{.content}}</div>
     ```

- 静态文件服务

  当请求路由为 /static 时，去 ./static 下找对应资源
  r.Static("/static", "./static")

## gin 路由传值

- get
  ctx.Query("key")
  ctx.DefaultQuery("key", "default_value")

- post
  ctx.PostForm("key")
  ctx.DefaultPostForm("key", "default_value")

- 将传入的参数解析到结构体上
  var user UserInfo
  err := ctx.ShouldBind(&user)

  ShouldBind 会返回错误让用户处理，MustBind 会直接响应 400
  ShouldBind/MustBind 会自动根据 Content-Type 推断绑定类型
  如果明确知道要绑定什么，这可以使用 ShouldBingWidth, MustBingWidth

- 解析 XML 数据

  1. bindXml
  2. shouldBindXml
  3. row, err := ctx.GetRowData; err = xml.unmarshal(row, ptr);

- 动态路由
  r.Get("/api/user/:uid")
  获取路由参数
  ctx.Param("uid")
  ctx.Params

## 路由分组

apiRouter := c.Group("/api")
apiRouter.Get("/", func(ctx gin.Context) {}); // /api
apiRouter.Get("/a", func(ctx gin.Context) {}); // /api/a

## 自定义控制器

gin05

## 中间件

gin06

- handlers 可以传递多个，最后一个被称作 请求处理函数，前面所有的函数被称作中间件（可以说，所有函数都是中间件，只不过请求的响应一般要交给到最后一个中奖件处理）
  api.GET(path, ...handler)
- ctx.Next() 让后续中间件开始运行
- ctx.Abort() 终止响应，后续中间件不再允许
- 全局中间件
  r.use(...middlewares)
  group.use(...middlewares)
- 中间件数据共享
- ctx.Set(key, value)
- prop, ok := ctx.Get(key)
  prop 是 any 类型，可以使用类型断言明确类型 prop.(string)
- r := gin.Default() 此时 r 已经包含两个中间件 Logger Recovery
- r := gin.New() 可以创建不包含任何中间件的路由
- 中间件中使用 gin.Context
  goroutine 中能使用 ctx,需要使用复制 ctxCp := ctx.Copy();

## 单文件上传

gin07
首先表单设置 post multipart/form-data
然后在接口中做一下操作

1. 获取文件
   file,err := ctx.FormFile("file")
2. 拼接存储路径
   dist := path.Join("./statics/upload", file.Filename)
3. 存储文件
   ctx.SaveUploadedFile(file, dist)

## 多文件上传

post multipart/form-data

1. 获取表单
   form, err := ctx.MultipartForm();
2. 获取文件数组
   files := form.File("files");
3. 循环存储
   ctx.SaveUploadedFile(file[index], dist)

## cookie

用于登录状态、浏览历史、购物车等

- 设置 cookie
  ctx.SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
- 读取 cookie
  v, err := ctx.Cookie(name)
- 删除 cookie
  maxAge < 0 意味着 删除
  maxAge = 0 意味着 不设置

## session

gin08

当客户端第一次访问服务器时，服务端会生成一个 session 对象,值存储在服务端，key 通过 cookie 设置到客户端
当下次浏览器请求服务器时，可以通过 cookie 中的 key 确定 session。

gin 框架中使用 session

```go
  // 创建一个存储session的仓库
  store := redis.NewStore(10, "tcp", "0.0.0.0:32768", "gin-redis", []byte("session secret key"))
	if err != nil {
		fmt.Println("链接redis失败", err)
		return
	}

  // 配置session中间件
	r.Use(sessions.Sessions("mysession", store))

  // 在路由中 set 或 get session
	r.GET("/session", func(ctx *gin.Context) {
    // get
		session := sessions.Default(ctx)
		username := session.Get("username")

    // set
		session.Options(sessions.Options{
			MaxAge: 3600,
		})
    session.Set("username", ctx.DefaultQuery("username", "张三"))
		session.Save()

		ctx.JSON(200, gin.H{"username": username})
	})
```
