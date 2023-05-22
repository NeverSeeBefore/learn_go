package main

import "github.com/gin-gonic/gin"

// gin 创建项目
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值:%v", "你好gin")
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "%v", "我是新闻页面")
	})

	r.GET("/aaa", func(c *gin.Context) {
		c.String(200, "path: %v;", "aaa")
	})

	// 默认在 0.0.0.0:8080启动项目
	// r.Run()
	// 主动指定8000为监听的端口
	r.Run(":8000")
}
