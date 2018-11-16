package server

import (
	"net"
	"os"
	"../integrate/logger"
	"fmt"
	"bytes"
	"io"
	"strings"
)

var notFoundStr, notSupperStr string

func init() {
	notFoundStr = "route is not defined."
	notSupperStr = "method is not supper"
}

func StartUpTCPServer(addr *string) {
	netListen, err := net.Listen("tcp", *addr)
	if nil != err {
		logger.Error(fmt.Sprintf("can't start server in %s ", *addr))
		logger.Error(err.Error())
		os.Exit(100)
	}
	defer netListen.Close()
	logger.Info("waiting request ...")
	for {
		conn, err := netListen.Accept()
		if nil != err {
			continue
		}
		if nil != conn{
			go forwardConn(conn)
		}
	}
}
func forwardConn(conn net.Conn) {
	buffer := receiveData(conn)
	defer conn.Close()
	if 1 < len(buffer) {
		arr := strings.Split(string(buffer), "\r\n")
		if 1 < len(arr) {
			logger.Info(strings.Join(arr, "\r\n"))
		}
	}
}

func receiveData(conn net.Conn) []byte {
	var buf bytes.Buffer
	buffer := make([]byte, 8192)
	for {
		sizenew, err := conn.Read(buffer)
		buf.Write(buffer[:sizenew])
		if err == io.EOF || sizenew < 8192 {
			break
		}
	}
	return buf.Bytes()
}
