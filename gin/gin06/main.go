package main

import (
	"fmt"
	"gin06/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

// 计算请求耗时
func middlewareCalcTimeCost(ctx *gin.Context) {
	fmt.Println("middlewareCalcTimeCost start")
	start := time.Now().UnixNano()
	ctx.Next()
	end := time.Now().UnixNano()
	fmt.Println("response cost:", end-start)
	fmt.Println("middlewareCalcTimeCost end")
}

func permissionCheck(ctx *gin.Context) {
	fmt.Println("permissionCheck start")
	ctx.Abort()
	ctx.JSON(200, gin.H{
		"data": "no permissions",
		"code": 999999,
		"msg":  "fail",
	})
	fmt.Println("permissionCheck end")
}

func middleware1(ctx *gin.Context) {
	ctx.Set("username", "张三")
	fmt.Println("middleware1 start")
	ctx.Next()
	fmt.Println("middleware1 end")
}
func middleware2(ctx *gin.Context) {
	fmt.Println("middleware2 start")
	fmt.Println("读取共享数据")
	username, ok := ctx.Get("username")
	if !ok {
		fmt.Println("未读取到属性 username", ok)
	}
	v, ok := username.(string)
	if !ok {
		return
	}
	var a = v
	fmt.Printf("username: %v %T\n", username, username)
	fmt.Printf("a: %v %T\n", a, a)
	ctx.Next()
	fmt.Println("middleware2 end")
}
func middleware3(ctx *gin.Context) {
	fmt.Println("middleware2 start")
	ctx.Next()
	// 中间件中可以开启 goroutine
	// 但是 goroutine 必须要操作ctx的副本而不是本身
	ctxCp := ctx.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Printf("middleware3 goroutine done! %v", ctxCp.Request.URL.Path)
	}()
	fmt.Println("middleware2 end")
}
func main() {
	r := gin.Default()

	// 全局中间件
	r.Use(middlewares.TimeCost)

	// 路由分组
	apiRouters := r.Group("/api")
	// 分组下的中间件
	apiRouters.Use(middleware1)

	apiRouters.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "/api",
			"code": 200,
			"msg":  "ok",
		})
	})

	// 中间可以通过 ctx.Abort(停止后续中间件运行并结束请求)
	apiRouters.GET("/permission", permissionCheck, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "/permission",
			"code": 200,
			"msg":  "ok",
		})
	})
	// 中间件按顺序执行。
	apiRouters.GET("/run", middleware2, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "/run",
			"code": 200,
			"msg":  "ok",
		})
	})
	// 中间件包含 goroutine
	apiRouters.GET("/go", middleware3, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "/go",
			"code": 200,
			"msg":  "ok",
		})
	})

	r.Run()
}
