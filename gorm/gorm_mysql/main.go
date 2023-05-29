package main

import (
	"fmt"
	"gin-proj/models"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// https://gorm.io/docs/query.html

	engine := gin.Default()
	fmt.Println("-------------------------")
	fmt.Println(time.Now().UnixMilli())
	fmt.Println("-------------------------")

	engine.GET("/users", func(ctx *gin.Context) {
		userList := []models.User{}
		result := models.DB.Preload("Cate").Find(&userList)
		ctx.JSON(200, gin.H{
			"path":         ctx.Request.URL.Path,
			"userList":     userList,
			"rowsAffected": result.RowsAffected,
			"error":        result.Error,
		})
	})

	engine.GET("/user/add", func(ctx *gin.Context) {
		user := models.User{
			Name:     "sunny",
			Age:      18,
			Email:    "sunny@gmail.com",
			AddTime:  time.Now(),
			EditTime: time.Now(),
		}
		result := models.DB.Create(&user)
		ctx.JSON(200, gin.H{
			"path":         ctx.Request.URL.Path,
			"status":       "ok",
			"rowsAffected": result.RowsAffected,
			"error":        result.Error,
		})
	})

	// Get the first record ordered by primary key
	engine.GET("/first", func(ctx *gin.Context) {
		user := models.User{}
		result := models.DB.First(&user)
		ctx.JSON(200, gin.H{
			"path":         ctx.Request.URL.Path,
			"description":  "SELECT * FROM users ORDER BY id LIMIT 1",
			"user":         user,
			"rowsAffected": result.RowsAffected,
			"error":        result.Error,
		})
	})

	engine.GET("/last", func(ctx *gin.Context) {
		user := models.User{}
		result := models.DB.Last(&user)
		ctx.JSON(200, gin.H{
			"path":         ctx.Request.URL.Path,
			"description":  "// Get the first record ordered by primary key",
			"user":         user,
			"rowsAffected": result.RowsAffected,
			"error":        result.Error,
		})
	})

	// Get one record, no specified order
	// SELECT * FROM users LIMIT 1
	engine.GET("/take", func(ctx *gin.Context) {
		user := models.User{}
		result := models.DB.Take(&user)
		ctx.JSON(200, gin.H{
			"path":         ctx.Request.URL.Path,
			"description":  "// Get the first record ordered by primary key",
			"user":         user,
			"rowsAffected": result.RowsAffected,
			"error":        result.Error,
		})
	})

	engine.Run()

}
