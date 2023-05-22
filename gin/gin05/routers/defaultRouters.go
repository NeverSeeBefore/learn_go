package routers

import (
	"gin05/controllers/page"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	pageController := page.PageController{}
	defaultRouters.GET("/", pageController.Home)
	defaultRouters.GET("/article", pageController.Article)
	defaultRouters.GET("/about", pageController.Article)
}
