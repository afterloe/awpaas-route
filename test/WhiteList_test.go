package test

import (
	"testing"
	"io"
	"time"
	"strings"
)

var (
	out io.Writer
	logLayout string
	whiteListCache []string
	)

func GetFormatTime() string {
	return time.Now().Format("2006-01-02 - 15:04:05")
}

func RemoveItem(item string) bool {
	index := -1
	for i, it := range whiteListCache {
		if it == item {
			index = i
			break
		}
	}
	if -1 == index {
		return false
	}
	whiteListCache = append(whiteListCache[:index], whiteListCache[index+1:]...)
	return true
}

func CleanAll() {
	whiteListCache = strings.Split("", "a")
}

func Test_whiteList(t *testing.T) {
	whiteListCache = append(whiteListCache, "/member/login")
	CleanAll()
	t.Log(whiteListCache)

	//flag := RemoveItem("/member/login")
	//t.Log(flag)
	//t.Log(whiteListCache)


	//reqUrl := "favicon.ico"
	//out = os.Stdout
	//for _, item := range whiteListCache {
	//	if 0 == strings.Index(reqUrl, item) {
	//		t.Log("find!")
	//		return
	//	}
	//}
	// for _, item := range whiteListCache {
	// 	if strings.Contains(reqUrl, item) {
	// 		t.Log("find!")
	// 		return
	// 	}
	// }
	// logLayout = "[awpaas-route][%-7s][%-5s][%v] - %-7s\n"
	// t.Log(fmt.Fprintf(out, logLayout, "daemon", "log", GetFormatTime(), "not found"))
	// t.Log(fmt.Fprintf(out, logLayout, "tcp", "error", GetFormatTime(), "not found"))
	// t.Log(fmt.Fprintf(out, logLayout, "gateway", "info", GetFormatTime(), "not found"))
}

