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
	"net/url"
	// "time"
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
	defer netListen.Close()
	if nil != err {
		logger.Error(fmt.Sprintf("can't start server in %s ", *addr))
		logger.Error(err.Error())
		os.Exit(100)
	}
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

func doForwardConn(conn net.Conn) {
	defer conn.Close()
	buffer := receiveData(conn)
	if 1 < len(buffer) {
		// var method, host, address string
		// fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
		arr := strings.Split(string(buffer), "\r\n")
		add := "127.0.0.1:5984"
		arr[1] = fmt.Sprintf("Host: %s", add)
		newStr := strings.Join(arr, "\r\n")
		server, err := net.Dial("tcp", add)
		if nil != err {
			logger.Info(err)
			return 
		}
		// if "CONNECT" == method {
		// 	fmt.Fprint(conn, "HTTP/1.1 200 Connection established\r\n")
		// } else {
		// 	server.Write(b[:n])
		// }
		server.Write([]byte(newStr))
		go io.Copy(server, conn)
		io.Copy(conn, server)
		defer server.Close()
	}
}

/**
	请求连接转发工作

	@param: conn - 连接信息
 */
func doForwardConnBack(conn net.Conn) {
	defer conn.Close()
	var b [1024]byte
	n, err := conn.Read(b[:])
	if nil != err {
		logger.Info(err)
		return 
	}
	var method, host, address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	logger.Info(fmt.Sprintf("%s - %s - %s", method, host, address))
	hostPortURL, err := url.Parse(host)

	logger.Info(hostPortURL)
	if nil != err {
		logger.Info(err)
		return 
	}

	if hostPortURL.Opaque == "443" { //https访问
        address = hostPortURL.Scheme + ":443"
    } else { //http访问
        if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
            address = hostPortURL.Host + ":80"
        } else {
            address = hostPortURL.Host
        }
	}

	logger.Info(address)
	
	server, err := net.Dial("tcp", address)
	if nil != err {
		logger.Info(err)
		return 
	}
	if "CONNECT" == method {
		fmt.Fprint(conn, "HTTP/1.1 200 Connection established\r\n")
	} else {
		server.Write(b[:n])
	}
	go io.Copy(server, conn)
	io.Copy(conn, server)

	// buffer := receiveData(conn)
	// if 1 < len(buffer) {
	// 	arr := strings.Split(string(buffer), "\r\n")
	// 	if 1 < len(arr) {

	// 		add := "127.0.0.1:5984"
	// 		arr[1] = fmt.Sprintf("Host: %s", add)
	// 		newStr := strings.Join(arr, "\r\n")
	// 		logger.Info(newStr)
	// 		forward([]byte(newStr), add, conn)

			// err, reqInfo := auth(arr) // 提取鉴权信息
			// pathStr, flag := queryWhiteList(reqInfo) // 查询白名单
			// if flag {
			// 	// 在白名单之内，不需要鉴权即可访问
			// 	arr[1] = fmt.Sprintf("Host: %s", pathStr)
			// 	logger.Info(strings.Join(arr, "\r\n"))
			// 	forward([]byte(strings.Join(arr, "\r\n")), pathStr, conn)
			// 	return
			// }
			// err = linkAndQuery(reqInfo) // 查询鉴权信息
			// err, pathStr = linkAndList(reqInfo) // 查询服务映射表
			// if nil == err {
			// 	// TODO 转发服务
			// 	logger.Info(strings.Join(arr, "\r\n"))
			// 	logger.Info(reqInfo.Token)
			// } else {
			// 	// TODO 驳回
			// 	logger.Error(err.Error())
			// }
	// 	} else {
	// 		// TODO 驳回
	// 	}
	// }
}

/**
	查询服务映射表

	@param: reqInfo - 请求信息
	@return: error - 异常信息
	@return: string - 服务映射的实际地址
*/
func linkAndList(req *authentication.ReqInfo) (error, string) {
	return nil, ""
}

/**
	查询白名单信息

	@param: reqInfo - 请求信息
	@return: error - 异常信息

	@return: str - 转发地址
*/
func queryWhiteList(req *authentication.ReqInfo) (string, bool) {
	return "127.0.0.1:5984", true
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
	// time.Sleep(10 * time.Millisecond)
	bufferHead := receiveData(conn)
	// time.Sleep(10 * time.Millisecond)
	bufferBody := receiveData(conn)
	var buf bytes.Buffer
	buf.Write(bufferHead)
	buf.Write(bufferBody)
	baseConn.Write(buf.Bytes())
	conn.Close()
}