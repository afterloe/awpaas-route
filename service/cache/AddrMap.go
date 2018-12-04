package cache

import (
	"github.com/gomodule/redigo/redis"
	"reflect"
	"../../exceptions"
)

var (
	addressMap map[string]string // 地址映射
	addrMapKey string // 地址映射 cache key
	defAddr string // 默认地址
	channelName4map string // 服务映射通道名
)

func init() {
	addressMap = make(map[string]string)
	addrMapKey = "serviceAddrMap"
	channelName4map = "serviceDiscovery"
}

/**
	对外 服务名地址映射
*/
func MapToAddress(serviceName string) (bool, string) {
	flag, addr := inAddrMap(serviceName)
	if flag {
		return true, addr.(string)
	}
	return false, ""
}

/**
	服务名地址映射
 */
func mapToAddress(serviceName string) string {
	flag, addr := MapToAddress(serviceName)
	if !flag {
		return defAddr
	}
	return addr
}

/**
	从本地缓存读取 服务地址映射
 */
func GetAddMapFromDisk() interface{} {
	return addressMap
}

/**
	从远程缓存存读取 服务地址映射
 */
func GetAddressMapFromRemote(key string) bool {
	reply, err := redis.StringMap(toRemote("HGETALL", key))
	if nil != err {
		return false
	}
	addressMap = reply
	return true
}

/**
	判断服务映射是否存在
 */
func inAddrMap(serviceName string) (bool, interface{}) {
	v := reflect.ValueOf(addressMap)
	key := reflect.ValueOf(serviceName)
	value := v.MapIndex(key)
	if value.IsValid() {
		return true, value.Interface()
	}
	return false, ""
}

/**
	添加映射
 */
func AppendAddrMap(serviceName, serviceAddr string) error {
	flag, _ := inAddrMap(serviceName)
	if flag {
		return &exceptions.Error{Code: 400, Msg: "service has been added."}
	}
	addressMap[serviceName] = serviceAddr
	toRemote("HSET", addrMapKey, serviceName, serviceAddr)
	toRemote("PUBLISH", channelName4map, "GET\\t\\n" + addrMapKey)
	return nil
}

/**
	修改映射
 */
func ModifyAddrMap(serviceName, serviceAddr string) error {
	flag, _ := inAddrMap(serviceName)
	if !flag {
		return &exceptions.Error{Code: 400, Msg: "service not registry."}
	}
	addressMap[serviceName] = serviceAddr
	toRemote("HSET", addrMapKey, serviceName, serviceAddr)
	toRemote("PUBLISH", channelName4map, "GET\\t\\n" + addrMapKey)
	return nil
}

/**
	删除映射
 */
func DelAddrMap(serviceName string) error {
	flag, _ := inAddrMap(serviceName)
	if !flag {
		return &exceptions.Error{Code: 400, Msg: "service not registry."}
	}
	delete(addressMap, serviceName)
	toRemote("HDEL", addrMapKey, serviceName)
	toRemote("PUBLISH", channelName4map, "GET\\t\\n" + addrMapKey)
	return nil
}