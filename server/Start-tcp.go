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

/**
	启动tcp服务

	@param: addr - 启动服务地址
	@param: serverCfg - package.JSON custom 配置中的内容
 */
func StartUpTCPServer(addr *string, serverCfg map[string]interface{}) {
	if nil != serverCfg["tokenName"] {
		key = serverCfg["tokenName"].(string)
	} else {
		// 默认 token名
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
			go doForwardConn(conn) // 异步处理
		}
	}
}

/**
	请求连接转发工作

	@param: conn - 连接信息
 */
func doForwardConn(conn net.Conn) {
	buffer := receiveData(conn)
	defer conn.Close()
	if 1 < len(buffer) {
		arr := strings.Split(string(buffer), "\r\n")
		if 1 < len(arr) {
			err, reqInfo := auth(arr) // 提取鉴权信息
			err = queryWhiteList(reqInfo) // 查询白名单
			if nil != err {
				//
			}
			err = linkAndQuery(reqInfo) // 查询鉴权信息
			err = linkAndList(reqInfo) // 查询服务映射表
			if nil == err {
				// TODO 转发服务
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

/**
	查询鉴权信息

	@param: info - 提出的鉴权信息
 */
func linkAndQuery(info *authentication.ReqInfo) error {
	return nil
}

/**
	提取鉴权信息

	@param：arr - tcp请求内容
 */
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

/**
	接收数据统一方法

	@param：conn - tcp连接
 */
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

/**
	转发服务

	@param：data - 转发tcp包
	@param：host - 信息内容
	@param：baseConn - 原连接信息
  */
func forward(data []byte, host string, baseConn net.Conn) {
	conn, _ := net.Dial("tcp", host)
	conn.Write(data)
	//time.Sleep(10 * time.Millisecond)
	bufferHead := receiveData(conn)
	//time.Sleep(10 * time.Millisecond)
	bufferBody := receiveData(conn)
	var buf bytes.Buffer
	buf.Write(bufferHead)
	buf.Write(bufferBody)
	baseConn.Write(buf.Bytes())
	conn.Close()
}