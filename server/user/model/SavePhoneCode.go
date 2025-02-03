package model

// 存储code到redis数据库
func SavePhoneCode(code, phoneNumber string) error {
	//连接数据库
	//conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	//if err != nil {
	//fmt.Println("redis error", err)
	//}
	conn := RedisPool.Get()
	defer conn.Close()

	//写数据库
	_, err := conn.Do("setex", phoneNumber, 60*5, code) //时效机制 5分钟

	return err
}
