package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

// 路由传值
func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./templates/*")

	// 获取get参数
	r.GET("/", func(ctx *gin.Context) {
		username := ctx.Query("username")
		age := ctx.DefaultQuery("age", "1")
		ctx.HTML(200, "index.html", gin.H{
			"title":    "首页gin05",
			"username": username,
			"age":      age,
		})
	})

	// 获取post参数
	r.POST("/api/user", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		age := ctx.DefaultPostForm("age", "1")
		ctx.JSON(200, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	type UserInfo struct {
		Username string `json:"username" form:"username"`
		Password string `json:"password" form:"password"`
		Age      int    `json:"age" form:"age"`
	}

	// 绑定参数到结构体
	r.POST("/api/userStruct", func(ctx *gin.Context) {
		var user UserInfo
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(200, gin.H{
				"status": "fail",
				"msg":    err.Error(),
				"data":   "",
			})
			return
		}

		ctx.JSON(200, gin.H{
			"status": "ok",
			"msg":    "",
			"data":   user,
		})
	})

	// 绑定xml
	r.POST("/api/xml", func(ctx *gin.Context) {
		var user UserInfo

		// 自动分析解析类型
		// if err := ctx.ShouldBind(&user); err != nil {
		// 	ctx.JSON(200, gin.H{
		// 		"status": "fail",
		// 		"msg":    err.Error(),
		// 		"data":   "",
		// 	})
		// 	return
		// }

		// 指定已xml方式绑定
		// if err := ctx.ShouldBindXML(&user); err != nil {
		// 	ctx.JSON(200, gin.H{
		// 		"status": "fail",
		// 		"msg":    err.Error(),
		// 		"data":   "",
		// 	})
		// 	return
		// }

		// 指定已xml方式绑定
		// if err := ctx.BindXML(&user); err != nil {
		// 	ctx.JSON(200, gin.H{
		// 		"status": "fail",
		// 		"msg":    err.Error(),
		// 		"data":   "",
		// 	})
		// 	return
		// }

		// 通过获取原始数据
		rowSlice, err := ctx.GetRawData()

		if err != nil {
			ctx.JSON(200, gin.H{
				"status": "fail",
				"msg":    "",
				"data":   err.Error(),
			})
			return
		}

		// 在使用xml包解析
		err = xml.Unmarshal(rowSlice, &user)
		if err != nil {
			ctx.JSON(200, gin.H{
				"status": "fail",
				"msg":    "",
				"data":   err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"status": "ok",
			"msg":    "",
			"data":   user,
		})
	})

	// 动态路由
	r.GET("/api/user/:uid", func(ctx *gin.Context) {

		ctx.JSON(200, gin.H{
			"status": "ok",
			"msg":    "",
			"data": map[string]interface{}{
				"uid":    ctx.Param("uid"),
				"params": ctx.Params,
			},
		})
	})

	r.Run()
}
