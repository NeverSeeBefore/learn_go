package routers

import (
	"gin05/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	userController := api.UserController{}
	apiRouters.GET("/user", userController.Index)
	apiRouters.POST("/user", userController.Add)
	apiRouters.DELETE("/user", userController.Del)
	apiRouters.PUT("/user", userController.Edit)
}
