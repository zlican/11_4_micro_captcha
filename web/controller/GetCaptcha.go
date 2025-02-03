package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"image/png"

	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"

	getCaptcha "11_4_micro_captcha/web/proto" //proto包
)

var Uuid string

func GetCaptcha(ctx *gin.Context) {
	Uuid = ctx.Param("uuid")

	//指定consul服务发现
	consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
	)

	//Client微服务
	//初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("getcaptcha", consulService.Client())

	//调用远程函数
	resp, err := microClient.Call(context.TODO(), &getCaptcha.CallRequest{Uuid: Uuid})
	if err != nil {
		fmt.Println("call error:", err)
		return
	}

	var img captcha.Image
	json.Unmarshal(resp.Img, &img) //反序列化成二进制
	png.Encode(ctx.Writer, img)

}
