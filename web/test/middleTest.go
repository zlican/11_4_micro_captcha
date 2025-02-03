package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func middleTest1(ctx *gin.Context) {
	ctx.Writer.Write([]byte("before middleware"))
	t := time.Now()
	ctx.Next() //调用下一个中间件
	ctx.Writer.Write([]byte("after middleware"))
	fmt.Print(time.Now().Sub(t)) //中间件用来测试程序消耗时间
}

func main() {
	r := gin.Default()
	//r.Use(middleTest1) //作用于以下所有路由

	r.GET("/test", middleTest1, func(ctx *gin.Context) {
		time.Sleep(1 * time.Second)
		ctx.Writer.Write([]byte("this is test"))
	})
	r.GET("/test2", middleTest1, func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("this is test2"))
	})

	r.Run(":8080")
}
