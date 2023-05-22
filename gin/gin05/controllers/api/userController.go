package api

import "github.com/gin-gonic/gin"

type UserController struct {
	BaseController
}

func (c UserController) Index(ctx *gin.Context) {
	c.success(ctx, gin.H{
		"api": "index",
	})
}

func (c UserController) Add(ctx *gin.Context) {
	c.success(ctx, gin.H{
		"api": "add",
	})
}

func (c UserController) Del(ctx *gin.Context) {
	c.success(ctx, gin.H{
		"api": "del",
	})
}

func (c UserController) Edit(ctx *gin.Context) {
	c.success(ctx, gin.H{
		"api": "edit",
	})

}
