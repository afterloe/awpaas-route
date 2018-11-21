package test

import (
	"fmt"
	"testing"
	"../integrate/logger"
	"github.com/gomodule/redigo/redis"
)

type Callback func(redis.Conn)

func getCache(key string) interface{} {

	return nil
}

func Test_QueryWhiteList(t *testing.T) {
	var (
		addr = ""
		flag = false
	)
	t.Log("begin to Test.")
	getConn(func(conn redis.Conn) {
		conn.Flush()
		reply, err := redis.String(conn.Do("GET", "whiteList"))
		if nil != err {
			return
		}
		addr = reply
		flag = true
		t.Log(reply)
	})
	t.Log(fmt.Sprintf("flag -> %v\t str -> %s", flag, addr))
}

func getConn(exec Callback) {
	fmt.Println("accept ...")
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if nil != err {
		logger.Error("can't connect redis service")
		return
	}
	defer c.Close()
	defer exec(c)
}