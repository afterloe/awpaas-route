package server

import (
	"../integrate/logger"
	"../service/cache"
	"github.com/gomodule/redigo/redis"
	"fmt"
	"strings"
	"time"
)

func serviceDiscovery(action, key string) {
	logger.Logger("cache", fmt.Sprintf("do %-7s by %-7s", action, key))
	switch action {
	case "GET":
		cache.GetAddressMapFromRemote(key)
	}
}

func whiteListChange(action, key string) {
	logger.Logger("cache", fmt.Sprintf("do %-7s by %-7s", action, key))
	switch action {
	case "GET":
		cache.GetWhiteListFromRemote(key)
	}
}

func handleMessage(channel string, data []byte) {
	// ACTION\\t\\nKEY
	context := string(data)
	content := strings.Split(context, "\\t\\n")
	logger.Logger("cache", fmt.Sprintf("receive %-7s from %-7s", context, channel))
	if 2 < len(content) {
		logger.Error("format msg fail ->" + context)
		return
	}
	switch channel {
	case "serviceDiscovery":
		serviceDiscovery(content[0], content[1])
	case "whiteListChange":
		whiteListChange(content[0], content[1])
	default:
		return
	}
}

func StartUpCacheServer(addr *string, channel []interface{}) {
	conn, err := redis.Dial("tcp", *addr, redis.DialConnectTimeout(3 * time.Second),
		redis.DialReadTimeout(3 * time.Second), redis.DialWriteTimeout(3 * time.Second))
	if nil != err {
		fmt.Println(err)
		logger.Error("cache", "can't get any from remote.. please check network -> " + *addr)
		return
	}
	defer conn.Close()
	if nil != err {
		logger.Error("cache", fmt.Sprintf("can't link cache server to %s ", *addr))
	}
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(channel...)
	logger.Info("cache", "cache service linked is ready ...")
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			handleMessage(v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			return
		}
	}
}