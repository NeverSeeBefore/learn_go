package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func timeCost(ctx *gin.Context) {
	start := time.Now().UnixNano()
	ctx.Next()
	end := time.Now().UnixNano()
	fmt.Println("response cost:", end-start)
}

func main() {

	r := gin.Default()
	r.Use(timeCost)
	r.LoadHTMLGlob("./templates/*")

	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/user", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"code": "000000",
				"data": "user",
				"msg":  "ok",
			})
		})

		apiRouters.POST("/upload", func(ctx *gin.Context) {
			file, err := ctx.FormFile("file")
			if err != nil {
				ctx.Abort()
				return
			}
			dist := path.Join("./statics/upload", file.Filename)
			ctx.SaveUploadedFile(file, dist)
			ctx.JSON(200, gin.H{
				"code": "000000",
				"data": dist,
				"msg":  "ok",
			})
		})

		apiRouters.POST("/batchUpload", func(ctx *gin.Context) {
			// 拿到表单数据
			form, err := ctx.MultipartForm()
			if err != nil {
				ctx.Abort()
				return
			}

			fmt.Println("file--", form.File)
			// 拿到文件列表
			files := form.File["file"]
			var distArr []string

			// 生成文件夹
			day := time.Now().Format("20060102")
			dir := path.Join("./statics/upload", day)
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				ctx.Abort()
				return
			}

			// 循环保存图片
			allowExt := map[string]bool{
				".png": true,
			}
			for index, v := range files {
				ext := path.Ext(v.Filename)
				fmt.Println("--------------------------------")
				fmt.Println("file size:", v.Size)
				fmt.Println("file size:", v.Size)
				fmt.Println("file ext:", ext)
				fmt.Println("--------------------------------")
				if !allowExt[ext] {
					continue
				}

				dist := path.Join(dir, strconv.FormatInt(time.Now().UnixMicro(), 10)+"-"+strconv.FormatInt(int64(index), 10)+ext)
				ctx.SaveUploadedFile(v, dist)
				distArr = append(distArr, dist)
			}
			ctx.JSON(200, gin.H{
				"code": "000000",
				"data": distArr,
				"msg":  "ok",
			})
		})

	}

	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(ctx *gin.Context) {
			ctx.HTML(200, "index", gin.H{
				"title": "Home",
			})
		})
	}

	r.Run()

}
