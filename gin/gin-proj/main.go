package main

import (
	"fmt"
	"gin-proj/middlewares"
	"gin-proj/routers"
	"text/template"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})

	engine.LoadHTMLGlob("./templates/*")

	store, err := redis.NewStore(10, "tcp", "0.0.0.0:32795", "gin-redis", []byte("session secret key"))
	if err != nil {
		fmt.Println("链接redis失败", err)
		return
	}
	engine.Use(sessions.Sessions("sessionid", store))
	engine.Static("/statics", "./statics")

	engine.Use(middlewares.TimeCost)

	routers.InitApiRouters(engine)
	routers.InitPageRouters(engine)

	engine.Run()

}

func UnixToTime(timestamp int64) string {
	t := time.UnixMilli(timestamp)
	s := t.Format("2006-01-02 15:04:05")
	fmt.Println("UnixToTime", s)
	return s
}
