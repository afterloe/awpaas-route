package test

import (
	"testing"
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func Test_redisPubSub(t *testing.T) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer c.Close()
	if nil != err {
		t.Errorf("find error:\t %s\r\n", err)
		return
	}
	psc := redis.PubSubConn{Conn: c}
	//psc.PSubscribe("example")
	psc.Subscribe("exp")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return
		}
	}
	//c.Send("SUBSCRIBE", "example")
	//c.Flush()
	//for {
	//	reply, err := c.Receive()
	//	byts,_ := redis.ByteSlices(reply, err)
	//	t.Log(string(byts[0]))
	//	t.Log(string(byts[1]))
	//	t.Log(string(byts[2]))
	//	if err != nil {
	//		t.Error(err)
	//		return
	//	}
	//	//time.Sleep(3000)
	//	//c.Send("UNSUBSCRIBE", "example", "192.168.3.3-help")
	//	//c.Flush()
	//
	//}
}