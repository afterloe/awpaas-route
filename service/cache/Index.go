package cache

import (
	"github.com/gomodule/redigo/redis"
	"strings"
	"reflect"
	"../../config"
	"../../exceptions"
	"../../integrate/logger"
	"fmt"
)

var (
	whiteListCache []string
	addressMap map[string]string
	defAddr string
	whiteListKey string
	addrMapKey string
)

func init() {
	whiteListCache = make([]string, 0)
	addressMap = make(map[string]string)
	gateway := config.GetByTarget(config.Get("server"), "gateway")
	defAddr = fmt.Sprintf("%s:%s",
		config.GetByTarget(gateway, "addr"),
		config.GetByTarget(gateway, "port"))
	whiteListKey = "whiteList"
	addrMapKey = "serviceAddrMap"
}

func LoadCache(list []interface{}) {
	GetAddressMapFromRemote("addrMap")
	flag := GetWhiteListFromRemote("whiteList")
	if !flag || 0 == len(whiteListCache) {
		flushWhiteListCache(list)
		SendWhiteListToRemote(whiteListKey)
	}
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
	对外 服务名地址映射
*/
func MapToAddress(serviceName string) (bool, string) {
	address := reflect.ValueOf(addressMap[serviceName])
	if !address.IsValid() {
		return true, address.String()
	}
	return false, ""
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

func GetAddMapFromDisk() interface{} {
	return addressMap
}

func GetWhiteListFromDisk() interface{} {
	return whiteListCache
}

func GetAddressMapFromRemote(key string) bool {
	reply, err := redis.StringMap(toRemote("HGETALL", key))
	if nil != err {
		return false
	}
	addressMap = reply
	return true
}

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

func SendWhiteListToRemote(key string) {
	content := strings.Join(whiteListCache, "\\t\\n")
	reply, _ := redis.String(toRemote("SET", key, content))
	logger.Info(reply)
	toRemote("PUBLISH", "whiteListChange", "GET\t\n"+ whiteListKey)
}

func toRemote(action string, key ...interface{}) (interface{}, error) {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	defer conn.Close()
	if nil != err {
		logger.Error("can't connect redis service")
		return nil, err
	}
	return conn.Do(action, key...)
}

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

func SendAddrMapToRemote(key string) {
	args := make([]interface{}, 0)
	args = append(args, key)
	for field, value := range addressMap {
		args = append(args, field, value)
	}
	reply, err := redis.String(toRemote("HMSET", args...))
	if nil != err || "OK" != reply {
		logger.Error(err)
	}
}

func AppendAddrMap(serviceName, serviceAddr string) error {
	v := reflect.ValueOf(addressMap)
	key := reflect.ValueOf(serviceName)
	value := v.MapIndex(key)
	if value.IsValid() {
		return &exceptions.Error{Code: 400, Msg: "service has been added."}
	}
	v.SetMapIndex(key, reflect.ValueOf(serviceAddr))
	SendAddrMapToRemote(addrMapKey)
	return nil
}

func ServiceMapModify(serviceName, serviceAddr string) error {
	v := reflect.ValueOf(addressMap)
	key := reflect.ValueOf(serviceName)
	value := v.MapIndex(key)
	if !value.IsValid() {
		return &exceptions.Error{Code: 400, Msg: "service not registry."}
	}
	v.SetMapIndex(key, reflect.ValueOf(serviceAddr))
	SendAddrMapToRemote(addrMapKey)
	return nil
}