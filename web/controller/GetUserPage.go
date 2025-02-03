package controller

import "github.com/gin-gonic/gin"

func GetUserPage(ctx *gin.Context) {
	ctx.HTML(200, "user.html", gin.H{})
}
