package page

import "github.com/gin-gonic/gin"

type PageController struct{}

func (p PageController) Home(ctx *gin.Context) {
	ctx.String(200, "Home")
}
func (p PageController) Article(ctx *gin.Context) {
	ctx.String(200, "Article")
}
func (p PageController) About(ctx *gin.Context) {
	ctx.String(200, "About")
}
