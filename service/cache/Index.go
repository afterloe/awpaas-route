package cache

import (
	"../../integrate/logger"
	"github.com/gomodule/redigo/redis"
)

type Callback func(redis.Conn)

func getCache(key string) interface{} {

	return nil
}

func QueryWhiteList(service_name string) (bool, string) {
	var (
		addr = ""
		flag = false
	)
	getConn(func(conn redis.Conn) {
		reply, err := redis.String(conn.Do("GET", "whiteList"))
		if nil != err {
			return
		}
		addr = reply
		
	})
	return flag, addr
}

func getConn(exec Callback) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if nil != err {
		logger.Error("can't connect redis service")
		return
	}
	defer c.Close()
	defer exec(c)
}