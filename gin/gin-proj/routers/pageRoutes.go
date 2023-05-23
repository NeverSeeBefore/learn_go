package routers

import (
	"gin-proj/controllers"

	"github.com/gin-gonic/gin"
)

func InitPageRouters(engine *gin.Engine) {
	routers := engine.Group("/")

	p := controllers.PageController{}
	routers.GET("/", p.Home)
}
