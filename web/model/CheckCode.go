package model

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func CheckCode(phoneNumber, code string) bool {
	//连接数据库
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis error", err)
	}
	defer conn.Close()

	res, _ := redis.String(conn.Do("get", phoneNumber))
	fmt.Println(res, code)
	return res == code
}
