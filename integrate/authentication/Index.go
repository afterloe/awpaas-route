package authentication

import (
	"strings"
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
