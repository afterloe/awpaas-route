package test

import (
	"testing"
	"../integrate/soaClient"
	"../config"
)

// http://pi.awpaas.cn/manager/info
func Test_Call(t *testing.T) {
	reply, err := soaClient.Call("GET", "pi.awpaas.cn", "/manager/info", nil, nil)
	if nil != err {
		t.Error(err)
	}
	t.Log(reply)
	t.Log(reply["data"])
	info := config.GetByTarget(reply["data"], "remarks")
	t.Log(info)
}