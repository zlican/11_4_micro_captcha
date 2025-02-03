package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	//连接数据库
	//conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//fmt.Println("redis error", err)
	//}
	//defer conn.Close()

	conn := RedisPool.Get()
	conn.Do("set", "zlican", "chen")
	rsp, _ := redis.String(conn.Do("get", "zlican"))
	fmt.Println(rsp)
}
