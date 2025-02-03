package model

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 存储图片id到redis数据库
func SaveImgCode(code, uuid string) error {
	//连接数据库
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis error", err)
	}
	defer conn.Close()

	//写数据库
	_, err = conn.Do("setex", uuid, 60*5, code) //时效机制 5分钟

	return err
}
