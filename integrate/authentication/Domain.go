package authentication

import "fmt"

/**
	对外包装的请求信息
 */
type ReqInfo struct {
	Method string // 请求方式
	ServerName string // 请求的服务名
	ReqUrl string // 请求url
	Token *tokenInfo // token信息
}

type tokenInfo struct {
	Flag bool // 是否包含token
	Value string // token 值
}

func(r *ReqInfo) String() string {
	return fmt.Sprintf("method: %s \t serverName: %s \t reqUrl: %s \t token: %s",
		r.Method, r.ServerName, r.ReqUrl, r.Token)
}

func(t *tokenInfo) String() string {
	return fmt.Sprintf("hasToken: %t \t value: %s", t.Flag, t.Value)
}