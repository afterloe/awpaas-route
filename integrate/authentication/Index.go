package authentication

import (
	"strings"
	"../../exceptions"
	"net/http"
)

func ExtractToken(req *http.Request, key string) *tokenInfo {
	token := &tokenInfo{ false, ""}
	value := req.Header.Get(key)
	if "" != value {
		token.Flag = true
		token.Value = value
	}
	return token
}

/**
	获取token 信息

	@param: arr - tcp 解包内容的数组
	@param: key - token 名字
	@return: token结构体
 */
func GetTokenInfo(arr []string, key string) *tokenInfo {
	token := &tokenInfo{false, ""}
	for index, str := range arr {
		flag := strings.Contains(str, key)
		if index > 0 && flag {
			token.Flag = flag
			token.Value = strings.Split(str, ": ")[1]
		}
	}
	return token
}

/*
	替换多余"/"

	@param：str - 需要替换的字符串
	@return：替换后的字符串
 */
func eliminate(str string) string {
	strArr := strings.Split(str, "/")
	newStr := make([]string, 0)
	for _,s := range strArr {
		if "" != s {
			newStr = append(newStr, s)
		}
	}
	return strings.Join(newStr, "/")
}

/*
	获取请求信息

	@param: str - tcp解包第一条信息 包含 Method、URL等信息
	@return： 请求信息
 */
func GetBaseInfo(str string) (error, *ReqInfo) {
	str = eliminate(str)
	arr := strings.Split(str, " ")
	if 1 > len(arr) {
		return &exceptions.Error{Msg: "format Error - not http", Code: 400}, nil
	}
	url := strings.Split(arr[1], "/")
	if 2 > len(url) {
		return &exceptions.Error{Msg: "format Error - can't find server name", Code: 400}, nil
	}
	return nil, &ReqInfo{arr[0], url[1], arr[2],"/" + strings.Join(url[2:], "/"), nil}
}