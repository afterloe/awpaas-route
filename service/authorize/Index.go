package authorize

import (
	"../../config"
	"../../integrate/soaClient"
	"../cache"
	"../../integrate/logger"
)

var (
	who, access, serviceName string
	enable bool
)

func init() {
	cfg := config.Get("author").(map[string]interface{})
	enable = false
	if cfg["enable"].(bool) {
		enable = cfg["enable"].(bool)
	}
	who = cfg["who"].(string)
	access = cfg["access"].(string)
	serviceName = cfg["serviceName"].(string)
}

/*
	查询鉴权信息
*/
func QueryAuthorizeInfo(token, accessServiceName, url string) (bool, string) {
	if !enable {
		return true, ""
	}
	registry, addr := cache.MapToAddress(serviceName)
	if !registry {
		return false, ""
	}
	callUrl := access + "?" + soaClient.Encode(map[string]interface{}{
		token: token,
		serviceName: accessServiceName,
		url: url,
	})
	reply, err := soaClient.Call("GET", addr, callUrl, nil, nil)
	if nil != err {
		return false, ""
	}
	logger.Info("authorize", reply)
	return true, "34dd9907628b4ae2b274764028f95e3a"
}