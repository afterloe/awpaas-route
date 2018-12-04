package server

import (
	"../integrate/logger"
	"../service/cache"
	"github.com/gomodule/redigo/redis"
	"fmt"
	"strings"
)

func serviceDiscovery(action, key string) {
	fmt.Printf("reflush service address map %s %s\r\n", action, key)
	switch action {
	case "GET":
		cache.GetAddressMapFromRemote(key)
	}
}

func whiteListChange(action, key string) {
	fmt.Printf("reflush white list %s %s\r\n", action, key)
	switch action {
	case "GET":
		cache.GetWhiteListFromRemote(key)
	}
}

func handleMessage(channel string, data []byte) {
	// ACTION\\t\\nKEY
	content := strings.Split(string(data), "\\t\\n")
	if 2 < len(content) {
		logger.Error("format msg fail ->" + string(data))
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
	conn, err := redis.Dial("tcp", *addr, redis.DialConnectTimeout(3000),
		redis.DialReadTimeout(3000), redis.DialWriteTimeout(3000))
	if nil != err {
		logger.Error("can't get any from remote.. please check network -> " + *addr)
		return
	}
	defer conn.Close()
	if nil != err {
		logger.Error(fmt.Sprintf("can't link cache server to %s ", *addr))
	}
	psc := redis.PubSubConn{Conn: conn}
	psc.Subscribe(channel...)
	logger.Info("cache service linked is ready ...")
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