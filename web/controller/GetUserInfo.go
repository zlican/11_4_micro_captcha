package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetUserInfo(ctx *gin.Context) {

	dbconf := "root:Aa123456@tcp(127.0.0.1:3306)/test?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dbconf), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var GetUserInfo UserInfo
	db.Where("phone_number = ?", PhoneNumber).Select("name").Find(&GetUserInfo)

	ctx.JSON(200, gin.H{"UserName": GetUserInfo.Name})
}
