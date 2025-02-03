package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DeleteSession(ctx *gin.Context) {
	s := sessions.Default(ctx)
	s.Delete("phoneNumber")
	s.Save()

	ctx.JSON(200, gin.H{"message": "login out success"})
}
