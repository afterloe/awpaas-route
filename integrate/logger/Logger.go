package logger

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"os"
	"io"
)

var (
	out, err io.Writer
	ginLogLayout, logLayout, timeFormat string
)

func init() {
	out = os.Stdout
	err = os.Stderr
	timeFormat = "2006-01-02 - 15:04:05"
	ginLogLayout = "[awpaas-route][%-7s][%-5s][%v] - %3d | %13v | %15s | %-7s %s\n"
	logLayout = "[awpaas-route][%-7s][%-5s][%v] - %-7s\n"
}

func getFormatTime() string {
	return time.Now().Format(timeFormat)
}

func Error(args ...interface{}) {
	fmt.Fprintf(err, logLayout, "error", args[0], getFormatTime(), args[1])
}

func Info(args ...interface{}){
	fmt.Fprintf(err, logLayout, "info", args[0], getFormatTime(), args[1])
}

func Logger(args ...interface{}) {
	fmt.Fprintf(err, logLayout, "log", args[0], getFormatTime(), args[1])
}

func GinLogger() func(*gin.Context) {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		//clientIP := c.ClientIP()
		clientIP := c.Request.Header["X-Real-IP"]
		if 0 == len(clientIP) {
			clientIP = append(clientIP, "127.0.0.1")
		}
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if "" != raw {
			path = path + "?" + raw
		}

		fmt.Fprintf(out, ginLogLayout, "daemon", "log", end.Format(timeFormat), statusCode,
			latency, clientIP, method, path)
	}
}
