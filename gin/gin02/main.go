package main

import (
	"github.com/gin-gonic/gin"
)

// gin.Context 各种响应方式
func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./templates/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "hello %v", "go")
	})

	r.GET("/json1", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "json1",
		})
	})

	// gin.H == map[string]interface{}
	r.GET("/json2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
			"msg":     "json2",
		})
	})

	type Article struct {
		Title   string `json:"title"`
		Desc    string `json:"desc"`
		Content string `json:"content"`
	}
	article := &Article{
		Title:   "article-title",
		Desc:    "article-desc",
		Content: "article-content",
	}
	r.GET("/json3", func(ctx *gin.Context) {
		ctx.JSON(200, article)
	})

	r.GET("/jsonp", func(ctx *gin.Context) {
		ctx.JSONP(200, article)
	})
	r.GET("/xml", func(ctx *gin.Context) {
		ctx.XML(200, gin.H{
			"success": true,
			"smg":     "xml-data",
		})
	})

	r.GET("/news", func(ctx *gin.Context) {
		ctx.HTML(200, "news.html", gin.H{
			"title": "后台数据",
		})
	})
	r.Run()
}
