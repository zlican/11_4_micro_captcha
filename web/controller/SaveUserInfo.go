package controller

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}

func SaveUserInfo(phoneNumer string) {
	//gorm 操作 mysql
	dbconf := "root:Aa123456@tcp(127.0.0.1:3306)/test?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dbconf), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&UserInfo{})

	db.Create(&UserInfo{PhoneNumber: phoneNumer, Name: phoneNumer})
}
