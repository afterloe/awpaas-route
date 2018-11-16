package test

import (
	"testing"
	"strings"
	"fmt"
	"../exceptions"
)

const key = "access-token"

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
	//str := "GET /BaseInfo/GetVehicleByNum?NUM=%E9%97%BDAT2236%E8%93%9D&TOKEN=4C4174581C243A191329212E111A113D39152C2F352A05050339223228393A002B30E55B HTTP/1.1|cache-control: no-cache|Postman-Token: 6e0109c4-4864-4e4a-a97e-29cfbfbb7929|User-Agent: PostmanRuntime/7.4.0|Accept: */*|Host: 127.0.0.1:8080|cookie: JSESSIONID=ID-Sn_EavnqNu5RljS3vMaW8Zv3LcVw8olQMdw11; XSRF-TOKEN=ad311914-4aa6-49bb-b68c-3cfd188b4a70|accept-encoding: gzip, deflate|Connection: keep-alive||"
	str := "GET /BaseInfo/GetVehicleByNum?NUM=%E9%97%BDAT2236%E8%93%9D&TOKEN=4C4174581C243A191329212E111A113D39152C2F352A05050339223228393A002B30E55B&access-token=sb HTTP/1.1|access-token: FDC77FC8A39AC10B984FE62C50F0B0319887EC370355A11E0EF4D1535021156ADCE6F469EF11BAD114A9B0433123B5F4|cache-control: no-cache|Postman-Token: b94ab7a3-dc4e-412c-94f3-5c8ce968bcef|User-Agent: PostmanRuntime/7.4.0|Accept: */*|Host: 127.0.0.1:8080|cookie: JSESSIONID=ID-Sn_EavnqNu5RljS3vMaW8Zv3LcVw8olQMdw11; XSRF-TOKEN=ad311914-4aa6-49bb-b68c-3cfd188b4a70|accept-encoding: gzip, deflate|Connection: keep-alive||"
	//str := "ASPC \\BaseInfo HTTP/1.1"
	//str := "ASPC ///BaseInfo/GetVehicleByNum//ss/ds/ssd////ss HTTP/1.1"
	arr := strings.Split(str, "|")
	token := getTokenInfo(arr)// 查询token
	if !token.Flag {
		fmt.Println("token is null")
		return
	}
	err, req := getBaseInfo(arr[0]) // 获取请求的 server 名字 和 请求路径
	if nil != err {
		t.Error(err)
	}
	req.Token = token
	fmt.Println(req)
}

func getTokenInfo(arr []string) *tokenInfo {
	token := &tokenInfo{false, ""}
	for index, str := range arr {
		flag := strings.Contains(str, key)
		if index > 0 && flag {
			token.Flag = true
			token.Value = strings.Split(str, ": ")[1]
		}
	}
	return token
}

func replic(str string) string {
	strArr := strings.Split(str, "/")
	newStr := make([]string, 0)
	for _,s := range strArr {
		if "" != s {
			newStr = append(newStr, s)
		}
	}
	return strings.Join(newStr, "/")
}

func getBaseInfo(str string) (error, *reqInfo) {
	str = replic(str)
	//str = strings.Replace(str, "//", "/", -1)
	arr := strings.Split(str, " ")
	if 1 > len(arr) {
		return &exceptions.Error{Msg: "format Error - not http", Code: 400}, nil
	}
	url := strings.Split(arr[1], "/")
	if 2 > len(url) {
		return &exceptions.Error{Msg: "format Error - can't find server name", Code: 400}, nil
	}
	return nil, &reqInfo{arr[0], url[1], strings.Join(url[2:], "/"), nil}
}
