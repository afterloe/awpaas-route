package test

import (
	"testing"
	"strings"
)

func Test_whiteList(t *testing.T) {
	whiteListCache := [...]string{"/member/login", "/fs/preview", "docker/images/json"}
	reqUrl := "*/images/json"
	for _, item := range whiteListCache {
		if strings.Contains(reqUrl, item) {
			t.Log("find!")
			return
		}
	}
	t.Log("not find!")
}

