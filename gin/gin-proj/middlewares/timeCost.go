package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeCost(ctx *gin.Context) {
	start := time.Now().UnixNano()
	ctx.Next()
	end := time.Now().UnixNano()
	fmt.Println("time cost", end-start)
}
