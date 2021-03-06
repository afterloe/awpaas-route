package test

import (
	"fmt"
	"testing"
	"../integrate/logger"
	"github.com/gomodule/redigo/redis"
	"strings"
)

type Callback func(redis.Conn)

func Test_QueryHash(t *testing.T) {
	t.Log("begin to Test.")
	getConn(func(conn redis.Conn) {
		conn.Flush()
		reply, err := redis.StringMap(conn.Do("HGETALL", "addrMap"))
		if nil != err {
			return
		}
		for key,val := range reply {
			t.Log(fmt.Sprintf("%s -> %s", key, val))
		}
	})
}

func test_QueryWhiteList(t *testing.T) {
	t.Log("begin to Test.")
	getConn(func(conn redis.Conn) {
		conn.Flush()
		reply, err := redis.String(conn.Do("GET", "w"))
		if nil != err {
			return
		}
		list := strings.Split(reply, "\\t\\n")
		for _, item := range list {
			t.Log(item)
		}
		t.Log(reply)
	})
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