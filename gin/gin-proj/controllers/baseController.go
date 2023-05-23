package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (c *BaseController) success(ctx *gin.Context, data interface{}) {
	fmt.Printf("----- data %#v", data)
	ctx.JSON(http.StatusOK, gin.H{
		"code": "000000",
		"data": data,
		"msg":  "ok",
	})
}

// func (c *BaseController) error(ctx *gin.Context, msg interface{}) {
// 	ctx.JSON(200, gin.H{
// 		"code": "999999",
// 		"msg":  msg,
// 	})
// }
