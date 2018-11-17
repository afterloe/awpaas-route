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
	// "time"
)

var (
	key string
	buffSize int
	needToken bool
)

func init() {
	key = "access-token"
	buffSize = 1024
	needToken = false
}

/**
	启动tcp服务

	@param: addr - 启动服务地址
	@param: serverCfg - package.JSON custom 配置中的内容
 */
func StartUpTCPServer(addr *string, serverCfg map[string]interface{}) {
	if nil != serverCfg["tokenName"] {
		key = serverCfg["tokenName"].(string)
	} 
	if nil != serverCfg["size"] {
		buffSize = int(serverCfg["size"].(float64))
	}
	if nil != serverCfg["needToken"] {
		needToken = serverCfg["needToken"].(bool)
	}
	netListen, err := net.Listen("tcp", *addr)
	defer netListen.Close()
	if nil != err {
		logger.Error(fmt.Sprintf("can't start server in %s ", *addr))
		logger.Error(err.Error())
		os.Exit(100)
	}
	logger.Info(fmt.Sprintf("auto config -- tokenName is %s, bufferSize is %d", key, buffSize))
	logger.Info("gateway service is ready ...")
	for {
		conn, err := netListen.Accept() // 获取客户端连接
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
	defer conn.Close()
	buffer := receiveData(conn) 
	if 1 < len(buffer) {
		arr := strings.Split(string(buffer), "\r\n")
		err, reqInfo := extractAuthInfo(arr) // 提取鉴权信息
		if nil != err { // 如果提取出现异常，则跳转到异常界面
			logger.Info("can't find authorize info.")
			return
		}
		flag, remote := query_whiteList(reqInfo) // 查询白名单
		if flag { // 在白名单之内，不需要鉴权即可访问
			arr[1] = fmt.Sprintf("Host: %s", remote)
			forward(reqInfo, remote, arr, conn)
			return
		}
		err = query_authInfo(reqInfo) // 查询鉴权信息
		if nil != err {
			logger.Info("authentication information query failed.")
			return
		}
		err, remote = query_mapList(reqInfo) // 查询服务映射表
		if nil != err { // 服务列表未查询到
			logger.Info("service not found.")
			return
		} 
		forward(reqInfo, remote, arr, conn)
	} else { // 返回服务异常
		logger.Info("has error")	
	}
}

/**
	提取鉴权信息

	@param：arr - tcp请求内容
 */
 func extractAuthInfo(arr []string) (error, *authentication.ReqInfo) {
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
	查询白名单信息

	@param: reqInfo - 请求信息
	@return: error - 异常信息

	@return: str - 转发地址
*/
func query_whiteList(req *authentication.ReqInfo) (bool, string) {
	return true, "127.0.0.1:5984"
}

/**
	查询鉴权信息

	@param: info - 提出的鉴权信息

	@return: error - 鉴权错误信息
 */
 func query_authInfo(info *authentication.ReqInfo) error {
	return nil
}

/**
	查询服务映射表

	@param: reqInfo - 请求信息
	@return: error - 异常信息
	@return: string - 服务映射的实际地址
*/
func query_mapList(req *authentication.ReqInfo) (error, string) {
	return nil, ""
}

/**
	转发服务

	@param：data - 转发tcp包
	@param：host - 信息内容
	@param：baseConn - 原连接信息
  */
 func forward(req *authentication.ReqInfo, addr string, content []string, baseConn net.Conn) {
	server, err := net.Dial("tcp", addr)
	defer server.Close()
	if nil != err {
		logger.Info(err)
		return
	}
	if "CONNECT" == req.Method {
		fmt.Fprint(baseConn, "HTTP/1.1 200 Connection established\r\n")
	} else {
		server.Write([]byte(strings.Join(content, "\r\n")))
	}
	go io.Copy(server, baseConn)
	io.Copy(baseConn, server)
}

/**
	接收数据统一方法

	@param：conn - tcp连接
 */
func receiveData(conn net.Conn) []byte {
	var buf bytes.Buffer
	buffer := make([]byte, buffSize)
	for {
		sizeNew, err := conn.Read(buffer)
		buf.Write(buffer[:sizeNew])
		if err == io.EOF || sizeNew < buffSize {
			break
		}
	}
	return buf.Bytes()
}