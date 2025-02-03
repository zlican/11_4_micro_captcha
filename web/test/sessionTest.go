package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//初始化容器， 设置加密密钥

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	//设置cookie key，加密成session
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		//调用session，设置session数据
		session := sessions.Default(c)
		var phoneNumber = "15659589183"
		session.Set("phoneNumber", phoneNumber)
		//修改session时需要save函数配合，否则不生效
		session.Save()
		stroeValue := session.Get("phoneNumber")
		c.JSON(200, gin.H{phoneNumber: stroeValue})
	})
	r.Run(":8000")
}
