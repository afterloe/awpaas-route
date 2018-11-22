package cache

import (
	"github.com/gomodule/redigo/redis"
	"strings"
	"reflect"
	"../../config"
	"../../integrate/logger"
	"fmt"
)

type Callback func(redis.Conn)
var (
	whiteListCache []string
	addressMap map[string]string
	defAddr string
)

func init() {
	whiteListCache = make([]string, 0)
	addressMap = make(map[string]string)
	gateway := config.GetByTarget(config.Get("server"), "gateway")
	defAddr = fmt.Sprintf("%s:%s",
		config.GetByTarget(gateway, "addr"),
		config.GetByTarget(gateway, "port"))
}

/**
	刷新白名单缓存
 */
func FlushWhiteListCache(list []interface{}) {
	var paramSlice []string
	for _, param := range list {
		paramSlice = append(paramSlice, param.(string))
	}
	whiteListCache = paramSlice
}

/**
	服务名地址映射
 */
func mapToAddress(serviceName string) string {
	address := reflect.ValueOf(addressMap[serviceName])
	if !address.IsValid() {
		return address.String()
	}
	return defAddr
}

/**
	判断是否在白名单之中
 */
func inWhiteList(reqUrl string) bool {
	for _, item := range whiteListCache {
		if strings.Contains(reqUrl, item) {
			return true
		}
	}
	return false
}

/**
	白名单映射
 */
func QueryWhiteList(url ,serviceName string) (bool, string) {
	flag := inWhiteList(url)
	if !flag {
		return flag, ""
	}
	return flag, mapToAddress(serviceName)
}

func GetWhiteListFromDisk() interface{} {
	return whiteListCache
}

func GetWhiteListFromRemote(key string) bool {
	var flag = false
	getConn(func(conn redis.Conn) {
		reply, err := redis.String(conn.Do("GET", key))
		if nil != err {
			return
		}
		flag = true
		list := strings.Split(reply, "\\t\\n")
		whiteListCache = list
	})
	return flag
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
