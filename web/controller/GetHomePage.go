package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var PhoneNumber string

func GetHomePage(ctx *gin.Context) {
	s := sessions.Default(ctx)
	value := s.Get("phoneNumber")
	PhoneNumber = value.(string)

	ctx.HTML(http.StatusOK, "home.html", gin.H{})

}
