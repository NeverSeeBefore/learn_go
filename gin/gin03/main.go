package main

import (
	"fmt"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

// gin 模板语法
func main() {
	r := gin.Default()

	r.Static("/static", "./static")

	// 配置自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})

	// 加载html模板 /** 表示一层目录
	r.LoadHTMLGlob("./templates/**/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "default/index.html", gin.H{
			"title": "首页",
		})
	})

	r.GET("/news", func(ctx *gin.Context) {
		type Article struct {
			Title   string `json:"title"`
			Content string `json:"content"`
			Desc    string `json:"description"`
		}
		ctx.HTML(200, "default/news.html", gin.H{
			"title": "新闻页",
			"news": Article{
				Title:   "article-title",
				Content: "article-content",
				Desc:    "article-description",
			},
		})
	})

	r.GET("/admin", func(ctx *gin.Context) {
		ctx.HTML(200, "admin/index.html", gin.H{
			"title": "admin首页",
			"score": 88,
			"hobby": []string{"吃饭", "睡觉", "打游戏"},
			"news": map[string]interface{}{
				"title":   "news-title",
				"content": "news-content",
			},
			"date": time.Now().UnixMilli(),
		})
	})

	r.GET("/admin/news", func(ctx *gin.Context) {
		ctx.HTML(200, "admin/news.html", gin.H{
			"title": "admin新闻页",
		})
	})

	r.Run()
}

func UnixToTime(timestamp int64) string {
	t := time.UnixMilli(timestamp)
	s := t.Format("2006-01-02 15:04:05")
	fmt.Println("UnixToTime", s)
	return s
}
