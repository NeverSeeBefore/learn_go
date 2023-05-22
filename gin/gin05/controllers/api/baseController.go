package api

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (c *BaseController) success(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, gin.H{
		"code": "000000",
		"data": data,
		"msg":  "ok",
	})
}

func (c *BaseController) error(ctx *gin.Context, msg interface{}) {
	ctx.JSON(200, gin.H{
		"code": "999999",
		"msg":  msg,
	})
}
