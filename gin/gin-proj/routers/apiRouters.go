package routers

import (
	"gin-proj/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitApiRouters(engine *gin.Engine) {
	routers := engine.Group("/api")

	userController := controllers.UserController{}
	routers.GET("/user", userController.List)

	routers.PUT("/user", userController.Edit)

	routers.GET("/session", func(ctx *gin.Context) {
		session := sessions.Default(ctx)

		ctx.JSON(200, gin.H{
			"path": "/session",
			"age":  session.Get("age"),
		})
	})
	routers.PUT("/session", func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		// session.Options(sessions.Options{MaxAge: 36000000})
		session.Set("age", ctx.DefaultPostForm("age", "100"))
		session.Save()

		ctx.JSON(200, gin.H{
			"path": "/session",
		})
	})
}
