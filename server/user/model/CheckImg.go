package model

import (
	"github.com/gomodule/redigo/redis"
)

func CheckImg(uuid, captchaCode string) bool {
	//连接数据库
	//conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//fmt.Println("redis error", err)
	//}

	conn := RedisPool.Get()
	defer conn.Close()

	rsp, _ := redis.String(conn.Do("get", uuid))

	return rsp == captchaCode
}
