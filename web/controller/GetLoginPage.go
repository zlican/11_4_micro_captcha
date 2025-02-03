package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{"title": "login"})
}
