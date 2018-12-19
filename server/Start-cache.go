package server

import (
	"../integrate/logger"
	"../service/cache"
	"github.com/gomodule/redigo/redis"
	"fmt"
	"strings"
	"time"
	"os"
)

/**
	服务发现 信息监听处理与分发

*/
func serviceDiscovery(action, key string) {
	logger.Logger("cache", fmt.Sprintf("do %-7s by %-7s", action, key))
	switch action {
	case "GET":
		cache.GetAddressMapFromRemote(key)
	}
}

/**
	白名单 信息监听与分发
*/
func whiteListChange(action, key string) {
	logger.Logger("cache", fmt.Sprintf("do %-7s by %-7s", action, key))
	switch action {
	case "GET":
		cache.GetWhiteListFromRemote(key)
	case "CLEAN":
		cache.GetWhiteListFromRemote(action)
	}
}

/**
	消息分发主线程

	ACTION\\t\\nKEY
*/
func handleMessage(channel string, data []byte) {
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

/**
	每30秒 判断远程服务器是否能够启动
*/
func reConn(addr *string) {
	timer := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-timer.C:
			func() {
				_, err := redis.Dial("tcp", *addr, redis.DialConnectTimeout(3 * time.Second))
				if nil != err {
					os.Exit(10)
				} else {
					reConn(addr)
				}
			}()
		}
	}
}

/**
	缓存同步模块主逻辑
*/
func StartUpCacheServer(addr *string, channel []interface{}) {
	conn, err := redis.Dial("tcp", *addr, redis.DialConnectTimeout(3 * time.Second))
	if nil != err {
		logger.Error("cache", "can't get any from remote.. please check network -> " + *addr)
		return
	}
	defer func() {
		logger.Error("cache", "linked service is down")
		// reload
		reConn(addr)
	}()
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
			logger.Logger("cache", fmt.Sprintf("%s: %s count %d", v.Channel, v.Kind, v.Count))
		case error:
			return
		}
	}
}