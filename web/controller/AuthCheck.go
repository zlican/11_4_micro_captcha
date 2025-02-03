package controller

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		if value := s.Get("phoneNumber"); value != nil {
			c.Next()
			return
		}

		c.String(http.StatusUnauthorized, "You are not allowed to. Please login")
		c.Abort()
		return
	}
}
