package main

import (
	"gin05/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 路由分组
	routers.ApiRoutersInit(r)
	routers.DefaultRoutersInit(r)

	r.Run()
}
