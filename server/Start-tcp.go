package server

import (
	"net"
	"os"
	"../integrate/logger"
	"../integrate/authentication"
	"../service/cache"
	"../service/authorize"
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
	daemonAddr string
)

func init() {
	key = "access-token"
	buffSize = 1024
	needToken = false
	daemonAddr = "127.0.0.1:8081"
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
			logger.Error("accept format exception")
			callDaemon(400, "format%20exception", conn)
			return
		}
		flag, remote := query_whiteList(reqInfo) // 查询白名单
		if flag { // 在白名单之内，不需要鉴权即可访问
			forward(reqInfo, remote, arr, conn)
			return
		} 
		if !reqInfo.Token.Flag { // 不在白名单之内则，又不存在token信息则报错
			logger.Info("can't find authorize info.")
			callDaemon(400, "can't%20find%20authorize%20info.", conn)
			return
		}
		if nil != query_authInfo(reqInfo) { // 查询鉴权信息
			logger.Info("authentication information query failed.")
			callDaemon(401, "can't%20find%20authorize%20info.", conn)
			return
		}
		flag, remote = cache.MapToAddress(reqInfo.ServerName) // 查询服务映射表
		if !flag { // 服务列表未查询到
			logger.Info("service not found.")
			callDaemon(404, "can't%20find%20" + reqInfo.ServerName + "%20info.", conn)
			return
		} 
		forward(reqInfo, remote, arr, conn)
	} else { // 返回服务异常
		logger.Error("accept error.")
		callDaemon(400, "can't%20find%20authorize%20info.", conn)
	}
}

/**
	访问守护线程获取信息

	@param: code - 返回状态码
	@param: msg - 返回信息
	@param: client - 客户端连接
*/
func callDaemon(code int, msg string, client net.Conn) {
	content := []string{"", "", "Connection: keep-alive", "User-Agent: inline", "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8", "Accept-Encoding: gzip, deflate, br", "Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,it;q=0.7", "", ""}
	content[0] = fmt.Sprintf("GET /v1/tips/%d?content=%s HTTP/1.1", code, msg)
	content[1] = fmt.Sprintf("Host: %s", daemonAddr)
	remote, err := net.Dial("tcp", daemonAddr)
	defer remote.Close()
	if nil != err {
		logger.Info(err)
		return
	}
	linkAndConnection(strings.Join(content, "\r\n"), remote, client)
	// io.Copy(baseConn, remote)
	// io.Copy(remote, baseConn)	
}

/**
	提取鉴权信息

	@param：arr - tcp请求内容
	@return: 提取异常
	@return: 鉴权信息
 */
func extractAuthInfo(arr []string) (error, *authentication.ReqInfo) {
	token := authentication.GetTokenInfo(arr, key)
	err, reqInfo := authentication.GetBaseInfo(arr[0]) // 获取请求的 server 名字 和 请求路径
	if nil == err {
		reqInfo.Token = token // 写入token信息
	}
	return err, reqInfo
}

/**
	查询白名单信息

	@param: reqInfo - 请求信息
	@return: bool - 是否存在与白名单
	@return: string - 映射的地址
*/
func query_whiteList(req *authentication.ReqInfo) (bool, string) {
	return cache.QueryWhiteList(req.ServerName + "/" + req.ReqUrl,
		req.ServerName)
}

/**
	查询鉴权信息

	@param: info - 提出的鉴权信息
	@return: error - 鉴权错误信息
 */
func query_authInfo(info *authentication.ReqInfo) error {
	var (
		token = info.Token.Value
		serviceName = info.ServerName
		url = info.ReqUrl
		requestURI = serviceName + "/" + url
	)
	flag, uid := authorize.QueryAuthorizeInfo(token, serviceName, url)
	if !flag {
		logger.Info(fmt.Sprintf("[failed] - %s access %s", token, requestURI))
		return &exceptions.Error{Msg: "token is error", Code: 400}
	}
	logger.Info(fmt.Sprintf("[success] - %s access %s", uid, requestURI))
	return nil
}

/**
	转发服务

	@param：data - 转发tcp包
	@param：host - 信息内容
	@param：client - 客户端链接
  */
func forward(req *authentication.ReqInfo, addr string, content []string, client net.Conn) {
	remote, err := net.Dial("tcp", addr)
	defer remote.Close()
	content[1] = fmt.Sprintf("Host: %s", addr) 
	content[0] = fmt.Sprintf("%s %s %s", req.Method, "/" + req.ReqUrl, req.Way)
	if nil != err {
		logger.Info(err)
		return
	}
	if "CONNECT" == req.Method {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	}
	linkAndConnection(strings.Join(content, "\r\n"), remote, client)
}

/**
	链接并转发请求

	@param: content - http报文
	@param: remote - 远程服务器
	@param: client - 客户端链接
*/
func linkAndConnection(content string, remote net.Conn, client net.Conn) {
	remote.Write([]byte(content))
	receiveBuf := receiveData(remote)
	client.Write(receiveBuf)
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