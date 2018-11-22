package test

import (
	"testing"
	"github.com/gomodule/redigo/redis"
	"fmt"
)

type callback func(...interface{}) string

type channelChain struct {
	fn callback
	next *channelChain
}

func (this *channelChain) SetNextSuccess(success *channelChain) *channelChain {
	this.next = success
	return this.next
}

func (this *channelChain) PassChannel(args ...interface{}) string {
	let := this.fn(args...)
	if "next" == let {
		return this.next.PassChannel(args...)
	}
	return let
}

var serviceDiscovery = &channelChain{func(i ...interface{}) string {
	var (
		channel = i[0]
		data = i[1]
		)
	if "serviceDiscovery" != channel {
		return "next"
	}
	fmt.Printf("receive serviceDiscovery ... %s \r\n", data)
	return ""
}, nil}

var whiteListChange = &channelChain{func(i ...interface{}) string {
	var (
		channel = i[0]
		data = i[1]
		)
	if "whiteListChange" != channel {
		return "next"
	}
	fmt.Printf("receive list whiteListChange ... %s \r\n", data)
	return ""
}, nil}

func Test_redisPubSub(t *testing.T) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	serviceDiscovery.SetNextSuccess(whiteListChange)
	defer c.Close()
	if nil != err {
		t.Errorf("find error:\t %s\r\n", err)
		return
	}
	psc := redis.PubSubConn{Conn: c}
	psc.Subscribe("serviceDiscovery", "whiteListChange")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			serviceDiscovery.PassChannel(v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return
		}
	}
}