package authentication

import "fmt"

type ReqInfo struct {
	Method string
	ServerName string
	ReqUrl string
	Token *tokenInfo
}

type tokenInfo struct {
	Flag bool
	Value string
}

func(r *ReqInfo) String() string {
	return fmt.Sprintf("method: %s \t serverName: %s \t reqUrl: %s \t token: %s",
		r.Method, r.ServerName, r.ReqUrl, r.Token)
}

func(t *tokenInfo) String() string {
	return fmt.Sprintf("hasToken: %t \t value: %s", t.Flag, t.Value)
}