package main

import (
	// 导入session包

	// 导入session存储引擎
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./templates/*")

	// 创建基于cookie的存储引擎，参数为密钥 随便写
	// store := cookie.NewStore([]byte("session secret key"))
	// store := redis.NewStore(10, "tcp", "0.0.0.0:32768", "gin-redis", []byte("session secret key"))
	store, err := redis.NewStore(10, "tcp", "0.0.0.0:32768", "gin-redis", []byte("session secret key"))
	if err != nil {
		fmt.Println("链接redis失败", err)
		return
	}

	// 配置session中间件
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/news", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		session.Options(sessions.Options{
			MaxAge: 3600,
		})
		session.Set("username", ctx.DefaultQuery("username", "张三"))
		session.Save()

		ctx.HTML(200, "news", gin.H{
			"title":    "News",
			"main":     "测试session",
			"username": session.Get("username"),
		})
	})

	r.GET("/session", func(ctx *gin.Context) {
		sessions := sessions.Default(ctx)
		username := sessions.Get("username")

		ctx.JSON(200, gin.H{"username": username})
	})

	r.Run()
}
