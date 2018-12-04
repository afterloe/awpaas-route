package cache

import (
	"github.com/gomodule/redigo/redis"
	"../../config"
	"../../exceptions"
	"../../integrate/logger"
	"fmt"
	"time"
)

var (
	redisAddr string // redis 连接地址
	fuse bool // 熔断标识
	startTime time.Time // 熔断启动时间
)

func init() {
	gateway := config.GetByTarget(config.Get("server"), "gateway")
	cache := config.GetByTarget(config.Get("server"), "cache")
	defAddr = fmt.Sprintf("%s:%s",
		config.GetByTarget(gateway, "addr"),
		config.GetByTarget(gateway, "port"))
	redisAddr = fmt.Sprintf("%s:%s",
		config.GetByTarget(cache, "addr"),
		config.GetByTarget(cache, "port"))
	fuse = false
	startTime = time.Now()
}

/**
	加载缓存
*/
func LoadCache(list []interface{}) {
	GetAddressMapFromRemote(addrMapKey)
	flag := GetWhiteListFromRemote(whiteListKey)
	if !flag || 0 == len(whiteListCache) {
		flushWhiteListCache(list)
		SendWhiteListToRemote(whiteListKey)
	}
}

/**
	远程连接
*/
func toRemote(action string, key ...interface{}) (interface{}, error) {
	// 如果熔断开启 并且 两次间隔没有超过30秒 直接返回
	if fuse && 30 * 1000 > time.Now().Sub(startTime) {
		return nil, &exceptions.Error{Msg: "fuse is open", Code: 500}
	}
	conn, err := redis.Dial("tcp", redisAddr, redis.DialConnectTimeout(3000),
		redis.DialReadTimeout(3000), redis.DialWriteTimeout(3000))
	if nil != err {
		logger.Error("can't connect redis service")
		fuse = true
		startTime = time.Now()
		return nil, err
	}
	defer conn.Close()
	return conn.Do(action, key...)
}
