package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

type User struct {
	Username string `json:"username"`
}

var user = User{Username: ""}

func (c *UserController) List(ctx *gin.Context) {
	c.success(ctx, user)
}

func (c *UserController) Edit(ctx *gin.Context) {
	user.Username = ctx.DefaultPostForm("username", "unknown")
	c.success(ctx, user)
}
