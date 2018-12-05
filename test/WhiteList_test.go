package test

import (
	"testing"
	"strings"
	"io"
	"fmt"
	"time"
	"os"
)

var (
	out io.Writer
	logLayout string
)

func GetFormatTime() string {
	return time.Now().Format("2006-01-02 - 15:04:05")
}

func Test_whiteList(t *testing.T) {
	whiteListCache := [...]string{"/member/login", "/fs/preview", "docker/images/json"}
	reqUrl := "*/images/json"
	out = os.Stdout
	for _, item := range whiteListCache {
		if strings.Contains(reqUrl, item) {
			t.Log("find!")
			return
		}
	}
	logLayout = "[awpaas-route][%-7s][%-5s][%v] - %-7s\n"
	t.Log(fmt.Fprintf(out, logLayout, "daemon", "log", GetFormatTime(), "not found"))
	t.Log(fmt.Fprintf(out, logLayout, "tcp", "error", GetFormatTime(), "not found"))
	t.Log(fmt.Fprintf(out, logLayout, "gateway", "info", GetFormatTime(), "not found"))
}

