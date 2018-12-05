package cache

import (
	"strings"
	"github.com/gomodule/redigo/redis"
	"../../integrate/logger"
)

var (
	whiteListCache []string // 白名单
	channelName4list string // 通道名
	whiteListKey string // 白名单 cache key
)

func init() {
	whiteListCache = make([]string, 0)
	channelName4list = "whiteListChange"
	whiteListKey = "whiteList"
}

/**
	刷新白名单缓存
 */
func flushWhiteListCache(list []interface{}) {
	var paramSlice []string
	for _, param := range list {
		paramSlice = append(paramSlice, param.(string))
	}
	whiteListCache = paramSlice
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
	查询白名单
 */
func QueryWhiteList(url ,serviceName string) (bool, string) {
	flag := inWhiteList(url)
	if !flag {
		return flag, ""
	}
	return flag, mapToAddress(serviceName)
}

/**
	通过本地缓存查询 白名单列表
 */
func GetWhiteListFromDisk() interface{} {
	return whiteListCache
}

/**
	通过远程缓存查询 白名单列表
 */
func GetWhiteListFromRemote(key string) bool {
	reply, err := redis.String(toRemote("GET", key))
	content, err := redis.String(reply, err)
	if nil != err {
		return false
	}
	list := strings.Split(content, "\\t\\n")
	whiteListCache = list
	return true
}

/**
	将白名单列表 更新至远程缓存
 */
func SendWhiteListToRemote(key string) {
	content := strings.Join(whiteListCache, "\\t\\n")
	reply, _ := redis.String(toRemote("SET", key, content))
	logger.Info("cache", reply)
	toRemote("PUBLISH", channelName4list, "GET\\t\\n"+ whiteListKey)
}

/**
	删除白名单
 */
func RemoveItem(item string) bool {
	index := -1
	for i, it := range whiteListCache {
		if it == item {
			index = i
			break
		}
	}
	if -1 == index {
		return false
	}
	whiteListCache = append(whiteListCache[:index], whiteListCache[index+1:]...)
	SendWhiteListToRemote(whiteListKey)
	return true
}

/**
	添加白名单
 */
func AppendItem(item string) bool {
	for _, it := range whiteListCache {
		if it == item {
			return false
		}
	}
	whiteListCache = append(whiteListCache, item)
	SendWhiteListToRemote(whiteListKey)
	return true
}
