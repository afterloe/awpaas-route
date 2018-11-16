package server

import (
	"net"
	"os"
	"../integrate/logger"
	"../integrate/authentication"
	"../exceptions"
	"fmt"
	"bytes"
	"io"
	"strings"
)

var notFoundStr, notSupperStr,key string

func init() {
	notFoundStr = "route is not defined."
	notSupperStr = "method is not supper"
}

func StartUpTCPServer(addr *string,serverCfg map[string]interface{}) {
	if nil != serverCfg["tokenName"] {
		key = serverCfg["tokenName"].(string)
	} else {
		key = "access-token"
	}
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

// 转发连接
func forwardConn(conn net.Conn) {
	buffer := receiveData(conn)
	defer conn.Close()
	if 1 < len(buffer) {
		arr := strings.Split(string(buffer), "\r\n")
		if 1 < len(arr) {
			err, reqInfo := auth(arr) // 进行鉴权，并提取信息
			if nil == err {
				// 转发服务
				logger.Info(strings.Join(arr, "\r\n"))
				logger.Info(reqInfo.Token)
			} else {
				// TODO 驳回
				logger.Error(err.Error())
			}
		} else {
			// TODO 驳回
		}
	}
}

// 服务鉴权
func auth(arr []string) (error, *authentication.ReqInfo) {
	token := authentication.GetTokenInfo(arr, key)
	if !token.Flag {
		return &exceptions.Error{Msg: "token is null", Code: 400}, nil
	}
	err, reqInfo := authentication.GetBaseInfo(arr[0]) // 获取请求的 server 名字 和 请求路径
	if nil == err {
		reqInfo.Token = token // 写入token信息
	}
	return err, reqInfo
}

//接收数据统一方法
func receiveData(conn net.Conn) []byte {
	var buf bytes.Buffer
	buffer := make([]byte, 8192)
	for {
		sizeNew, err := conn.Read(buffer)
		buf.Write(buffer[:sizeNew])
		if err == io.EOF || sizeNew < 8192 {
			break
		}
	}
	return buf.Bytes()
}

// 转发服务
func forward(data []byte, host string, baseconn net.Conn) {
	conn, _ := net.Dial("tcp", host)
	conn.Write(data)
	//time.Sleep(10 * time.Millisecond)
	bufferHead := receiveData(conn)
	//time.Sleep(10 * time.Millisecond)
	bufferBody := receiveData(conn)
	var buf bytes.Buffer
	buf.Write(bufferHead)
	buf.Write(bufferBody)
	baseconn.Write(buf.Bytes())
	conn.Close()
}