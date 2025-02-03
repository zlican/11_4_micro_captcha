package main

import (
	"11_4_micro_captcha/web/controller"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的gin引擎实例
	r := gin.Default()

	// 配置Redis作为session存储
	//配置session:初始化容器， 设置加密密钥
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	//设置cookie key，加密成session
	r.Use(sessions.Sessions("mysession", store))

	// 加载HTML模板文件
	r.LoadHTMLGlob("view/*")
	// 配置静态文件目录
	r.Static("/static", "./static")

	// 注册路由
	// 首页路由,需要验证登录
	r.GET("/home", controller.AuthCheck(), controller.GetHomePage)
	// 登录页面路由
	r.GET("/login", controller.GetLoginPage)

	// 用户页面路由
	r.GET("/user", controller.GetUserPage)
	// 获取验证码图片
	r.GET("/captcha/:uuid", controller.GetCaptcha)
	// 发送短信验证码
	r.GET("api/smscode/:phoneNumber", controller.SendSmsCode)
	// 获取用户信息
	r.GET("/getUserInfo", controller.GetUserInfo)
	// 删除session
	r.GET("/deleteSession", controller.DeleteSession)
	// 表单验证
	r.POST("/form", controller.VarifyForm)

	rTest := r.Group("/test")
	{
		rTest.GET("/", func(ctx *gin.Context) {
			ctx.HTML(200, "test.html", nil)
		})

		rTest.POST("/upload", func(ctx *gin.Context) {
			file, err := ctx.FormFile("image")
			if err != nil {
				ctx.JSON(400, gin.H{"message": "请选择图片"})
				return
			}
			ctx.SaveUploadedFile(file, "./static/upload/"+file.Filename)
			ctx.JSON(200, gin.H{"message": "上传成功"})
		})
	}
	// 启动服务器,监听8080端口
	r.Run(":8080")
}
