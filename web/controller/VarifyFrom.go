package controller

import (
	"11_4_micro_captcha/web/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func VarifyForm(ctx *gin.Context) {
	phoneNumber := ctx.PostForm("phoneNumber")
	phoneCode := ctx.PostForm("verification")

	boolS := model.CheckCode(phoneNumber, phoneCode)
	if boolS {
		SaveUserInfo(phoneNumber)
		s := sessions.Default(ctx)
		s.Set("phoneNumber", phoneNumber) //相当于在给用户返回的那个session里设定key-value
		s.Save()
		ctx.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/home")

	} else {
		ctx.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/login")
	}
}
