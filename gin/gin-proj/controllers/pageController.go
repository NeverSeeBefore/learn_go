package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageController struct {
	BaseController
}

func (c *PageController) Home(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index", gin.H{
		"title":    "home",
		"username": ctx.DefaultQuery("username", "sunny"),
	})
}
