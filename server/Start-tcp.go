package server

import (
	"../integrate/logger"
	"../integrate/authentication"
	"../service/authorize"
	"../exceptions"
	"../service/cache"
	"fmt"
	"strings"

	"io"
	"net/http"
	"log"
)

var (
	key string
	capacity int
	needToken bool
	daemonAddr string
)

func init() {
	key = "access-token"
	capacity = 1024
	needToken = false
	daemonAddr = "127.0.0.1:8081"
}

func extractInfo(req *http.Request) *authentication.ReqInfo {
	urlArr := strings.Split(req.RequestURI, "/")
	return &authentication.ReqInfo{
		Method: req.Method,
		ServerName: urlArr[1],
		Way: req.Proto,
		ReqUrl: "/" + strings.Join(urlArr[2:], "/"),
	}
}

func sendForward(req *http.Request, rw http.ResponseWriter, addr string, client *authentication.ReqInfo) {
	for _, v := range req.Trailer {
		fmt.Println(v)
	}
	remote, err := http.NewRequest(req.Method, fmt.Sprintf("http://%s%s", addr, client.ReqUrl), req.Body)
	for key, value := range req.Header {
		for _, v := range value {
			remote.Header.Add(key, v)
		}
	}
	response, err := http.DefaultClient.Do(remote)
	if err != nil && response == nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	} else {
		// Close the connection to reuse it
		defer response.Body.Close()
		for key, value := range response.Header {
			for _, v := range value {
				rw.Header().Add(key, v)
			}
		}
		io.Copy(rw, response.Body)
	}
}

func sendDaemonForward(code int, msg string, req *http.Request, rw http.ResponseWriter)  {
	remote, err := http.NewRequest("GET", fmt.Sprintf("http://%s/v1/tips/%d?content=%s", daemonAddr, code, msg), nil)
	for key, value := range req.Header {
		for _, v := range value {
			remote.Header.Add(key, v)
		}
	}
	response, err := http.DefaultClient.Do(remote)
	if err != nil && response == nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
		//sendDaemonForward(502, "service%20inaccessibility", client)
	} else {
		// Close the connection to reuse it
		defer response.Body.Close()
		for key, value := range response.Header {
			for _, v := range value {
				rw.Header().Add(key, v)
			}
		}
		io.Copy(rw, response.Body)
	}
}

type Pxy struct {}

func (*Pxy)ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	client := extractInfo(req)
	flag, addr := cache.QueryWhiteList(req.RequestURI, client.ServerName)
	if flag {
		sendForward(req, rw, addr, client)
		return
	}
	if "" == req.Header.Get(key) {
		logger.Info("gateway", "can't find authorize info.")
		sendDaemonForward(400, "can't%20find%20authorize%20info.", req, rw)
		return
	}
	client.Token = authentication.ExtractToken(req, key)
	if nil != queryAuthInfo(client) { // 查询鉴权信息
		logger.Info("gateway", "authentication information query failed.")
		sendDaemonForward(401, "can't%20find%20authorize%20info.", req, rw)
		return
	}
	flag, addr = cache.MapToAddress(client.ServerName) // 查询服务映射表
	if !flag { // 服务列表未查询到
		logger.Info("gateway", "service not found.")
		sendDaemonForward(404, "can't%20find%20" + client.ServerName + "%20info.", req, rw)
		return
	}
	sendForward(req, rw, addr, client)
}

/**
	查询鉴权信息
	@param: info - 提出的鉴权信息
	@return: error - 鉴权错误信息
 */
func queryAuthInfo(info *authentication.ReqInfo) error {
	var (
		token = info.Token.Value
		serviceName = info.ServerName
		url = info.ReqUrl
		requestURI = serviceName + url
	)
	flag, uid := authorize.QueryAuthorizeInfo(token, serviceName, url)
	if !flag {
		logger.Info("gateway", fmt.Sprintf("[failed] - %s access %s", token, requestURI))
		return &exceptions.Error{Msg: "token is error", Code: 400}
	}
	logger.Info("gateway", fmt.Sprintf("[success] - %s access %s", uid, requestURI))
	return nil
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
		capacity = int(serverCfg["size"].(float64))
	}
	if nil != serverCfg["needToken"] {
		needToken = serverCfg["needToken"].(bool)
	}
	http.Handle("/", &Pxy{})
	http.ListenAndServe(*addr, nil)
}
