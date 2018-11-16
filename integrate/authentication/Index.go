package authentication

import (
	"strings"
	"../../exceptions"
)

/*
	获取token 信息
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
	return nil, &ReqInfo{arr[0], url[1], strings.Join(url[2:], "/"), nil}
}