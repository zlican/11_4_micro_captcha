package controller

import (
	"11_4_micro_captcha/web/user-proto"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
)

var phoneNumber string
var Code string

func SendSmsCode(ctx *gin.Context) {
	phoneNumber = ctx.Param("phoneNumber")
	captchaCode := ctx.Query("captchaCode")
	uuid := ctx.Query("uuid")

	data := make(map[string]string)

	//启用微服务
	consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
	)

	microClient := user.NewUserService("user", consulService.Client())
	resp, err := microClient.SendSmsCode(context.TODO(), &user.SendSmsCodeRequest{PhoneNumber: phoneNumber, CaptCode: captchaCode, Uuid: uuid})
	if err != nil {
		fmt.Println("SendSmsCode micro error:", err)
	}

	data["message"] = resp.Message
	data["ok"] = resp.Ok

	ctx.JSON(200, data)

}
