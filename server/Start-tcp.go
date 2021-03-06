package server

import (
	"../integrate/logger"
	"../integrate/authentication"
	"../integrate/soaClient"
	"../service/authorize"
	"../exceptions"
	"../service/cache"
	"fmt"
	"strings"

	"io"
	"net/http"
)

var (
	key string
	needToken bool
	daemonAddr string
)

func init() {
	key = "access-token"
	needToken = false
	daemonAddr = "127.0.0.1:8081"
}

/**
	提取信息
*/
func extractInfo(req *http.Request) *authentication.ReqInfo {
	urlArr := strings.Split(req.RequestURI, "/")
	schema := req.Header.Get("X-Forwarded-Proto")
	if "ws" == schema {
		schema = "ws" // 支持webSocket
	} else {
		schema = "http"
	}
	return &authentication.ReqInfo{
		Method: req.Method,
		ServerName: urlArr[1],
		Way: req.Proto,
		Scheme: schema,
		ReqUrl: "/" + strings.Join(urlArr[2:], "/"),
	}
}

/**
	启动转发服务
*/
func sendForward(req *http.Request, rw http.ResponseWriter, addr string, client *authentication.ReqInfo) {
	remote, err := http.NewRequest(req.Method, fmt.Sprintf("%s://%s%s", client.Scheme ,addr, client.ReqUrl), req.Body)
	if nil != err {
		sendDaemonForward(500, fmt.Sprintf("service+%s+is+down", client.ServerName), req, rw)
	}
	forward(req, rw, remote)
}

/**
	转发服务业务逻辑
*/
func forward(r *http.Request, w http.ResponseWriter, remote *http.Request) {
	for key, value := range r.Header {
		for _, v := range value {
			remote.Header.Add(key, v)
		}
	}
	response, err := soaClient.GeneratorClient().Do(remote)
	if err != nil && response == nil {
		logger.Error("gateway", fmt.Sprintf("forward %+v", err))
	} else {
		defer response.Body.Close()
		for key, value := range response.Header {
			for _, v := range value {
				w.Header().Add(key, v)
			}
		}
		logger.Logger("gateway", fmt.Sprintf("%3d | %15s | %-7s | %s", response.StatusCode,
			r.Header.Get("X-Real-IP"), r.Method, r.RequestURI))
		io.Copy(w, response.Body)
	}
}

/**
	转发服务到守护进程
*/
func sendDaemonForward(code int, msg string, req *http.Request, rw http.ResponseWriter)  {
	remote, err := http.NewRequest("GET", fmt.Sprintf("http://%s/v1/tips/%d?content=%s", daemonAddr, code, msg), nil)
	if nil != err {
		logger.Error("gateway", fmt.Sprintf("forward %+v", err))
	}
	forward(req, rw, remote)
}

/**
	网关主逻辑

	@param: rw
	@param: req
*/
func doGateway(rw http.ResponseWriter, req *http.Request) {
	//日志 缺少访问 计时 以及 谁访问了什么路径
	if "/favicon.ico" == req.RequestURI {
		return
	}
	client := extractInfo(req) // 提取网关请求信息
	flag, addr := cache.QueryWhiteList(req.RequestURI, client.ServerName) // 查询请求是否在白名单之内
	if flag {
		// 存在即转发
		sendForward(req, rw, addr, client)
		return
	}
	if "" == req.Header.Get(key) {
		logger.Info("gateway", "can't find authorize info.")
		sendDaemonForward(400, "can't%20find%20authorize%20info.", req, rw)
		return
	}
	// 提取鉴权信息
	client.Token = authentication.ExtractToken(req, key)
	if nil != queryAuthInfo(client) { // 查询鉴权信息
		logger.Info("gateway", "authentication information query failed.")
		sendDaemonForward(401, "can't%20find%20authorize%20info.", req, rw)
		return
	}
	// 服务查询服务影射
	flag, addr = cache.MapToAddress(client.ServerName) // 查询服务映射表
	if !flag { // 服务列表未查询到
		logger.Info("gateway", "service not found.")
		sendDaemonForward(404, "can't%20find%20" + client.ServerName + "%20info.", req, rw)
		return
	}
	// 转发服务
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
	if nil != serverCfg["needToken"] {
		needToken = serverCfg["needToken"].(bool)
	}
	http.HandleFunc("/", doGateway)
	http.ListenAndServe(*addr, nil)
}
