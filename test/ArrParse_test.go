package test

import (
	"testing"
	"strings"
	"fmt"
)

type reqInfo struct {
	Method string
	ServerName string
	ReqUrl string
	Token *tokenInfo
}

type tokenInfo struct {
	Flag bool
	Value string
}

func(r *reqInfo) String() string {
	return fmt.Sprintf("method: %s \t serverName: %s \t reqUrl: %s \t token: %s",
		r.Method, r.ServerName, r.ReqUrl, r.Token)
}

func(t *tokenInfo) String() string {
	return fmt.Sprintf("hasToken: %t \t value: %s", t.Flag, t.Value)
}

func Test_parseWord(t *testing.T)  {
	str := "GET /BaseInfo/GetVehicleByNum?NUM=%E9%97%BDAT2236%E8%93%9D&TOKEN=4C4174581C243A191329212E111A113D39152C2F352A05050339223228393A002B30E55B HTTP/1.1|cache-control: no-cache|Postman-Token: 6e0109c4-4864-4e4a-a97e-29cfbfbb7929|User-Agent: PostmanRuntime/7.4.0|Accept: */*|Host: 127.0.0.1:8080|cookie: JSESSIONID=ID-Sn_EavnqNu5RljS3vMaW8Zv3LcVw8olQMdw11; XSRF-TOKEN=ad311914-4aa6-49bb-b68c-3cfd188b4a70|accept-encoding: gzip, deflate|Connection: keep-alive||"
	//str := "GET /BaseInfo/GetVehicleByNum?NUM=%E9%97%BDAT2236%E8%93%9D&TOKEN=4C4174581C243A191329212E111A113D39152C2F352A05050339223228393A002B30E55B&access-token=sb HTTP/1.1|access-token: FDC77FC8A39AC10B984FE62C50F0B0319887EC370355A11E0EF4D1535021156ADCE6F469EF11BAD114A9B0433123B5F4|cache-control: no-cache|Postman-Token: b94ab7a3-dc4e-412c-94f3-5c8ce968bcef|User-Agent: PostmanRuntime/7.4.0|Accept: */*|Host: 127.0.0.1:8080|cookie: JSESSIONID=ID-Sn_EavnqNu5RljS3vMaW8Zv3LcVw8olQMdw11; XSRF-TOKEN=ad311914-4aa6-49bb-b68c-3cfd188b4a70|accept-encoding: gzip, deflate|Connection: keep-alive||"
	//str := "GET /docker/images/json HTTP/1.1|cache-control: no-cache|Postman-Token: 714efe78-993a-4a85-8f70-8d710c4ab0db|User-Agent: PostmanRuntime/7.4.0|Accept: */*|Host: 127.0.0.1:8080|cookie: JSESSIONID=ID-Sn_EavnqNu5RljS3vMaW8Zv3LcVw8olQMdw11; XSRF-TOKEN=ad311914-4aa6-49bb-b68c-3cfd188b4a70|accept-encoding: gzip, deflate|Connection: keep-alive||"
	//str := "ASPC \\BaseInfo HTTP/1.1"
	//str := "ASPC ///BaseInfo/GetVehicleByNum//ss/ds/ssd////ss HTTP/1.1"
	arr := strings.Split(str, "|")
	var URLIndex,HOSTIndex int
	for index, it := range arr {
		if 1 == strings.Count(it, " HTTP/") {
			URLIndex = index
		}
		if 1 == strings.Count(it, "Host: ") {
			HOSTIndex = index
		}
	}
	t.Log(arr[URLIndex])
	t.Log(arr[HOSTIndex])
}
