package main

import (
	"runtime"
	"os"
	"fmt"
	cli "./server"
	"./integrate/logger"
	"./config"
	"reflect"
)

var (
	defAddr string
	cpuNumber int
	pid int
)

func init() {
	defAddr = "0.0.0.0:8080"
	cpuNumber = runtime.NumCPU()
	pid = os.Getpid()
}

/**
	启动默认服务
 */
func startDefault() {
	logger.Info("main",fmt.Sprintf("listen parameter is null, will start server in %s default", defAddr))
	logger.Info("main",fmt.Sprintf("server is init success... started pid is %d", pid))
	cli.StartUpTCPServer(&defAddr, config.Get("custom").(map[string]interface{}))
}

/**
	启动缓存服务
 */
func startUpCacheService(cfg map[string]interface{}) {
	addr, port := cfg["addr"], cfg["port"]
	if nil == addr {
		addr = "127.0.0.1"
	}
	if nil == port {
		port = "6379"
	}
	addrStr := fmt.Sprintf("%s:%s", addr, port)
	logger.Info("main",fmt.Sprintf("cache server will listen to %s ", addrStr))
	cli.StartUpCacheServer(&addrStr, cfg["channel"].([]interface{}))
}

/**
	启动网关服务
 */
func startUpGatewayService(cfg map[string]interface{}) {
	addr, port := cfg["addr"], cfg["port"]
	if nil == addr {
		addr = "127.0.0.1"
	}
	if nil == port {
		port = "8080"
	}
	multiServiceCfg(cfg["multiCore"].(map[string]interface{}))
	addrStr := fmt.Sprintf("%s:%s", addr, port)
	logger.Info("main",fmt.Sprintf("gateway server will start in %s ", addrStr))
	cli.StartUpTCPServer(&addrStr, config.Get("custom").(map[string]interface{}))
}

/**
	启动守护进程
 */
func startUpDaemonService(cfg map[string]interface{}) {
	addr, port := cfg["addr"], cfg["port"]
	if nil == addr {
		addr = "127.0.0.1"
	}
	if nil == port {
		port = "8081"
	}
	addrStr := fmt.Sprintf("%s:%s", addr, port)
	logger.Info("main",fmt.Sprintf("daemon server will start in %s ", addrStr))
	cli.StartUpDaemonService(&addrStr, config.Get("custom"))
}

/**
	多核设置
 */
func multiServiceCfg(cfg map[string]interface{}) {
	flg := cfg["enable"]
	if nil == flg {
		logger.Info("main", "server not allow to use multi cpu")
		return
	}
	if flg.(bool) {
		num := config.GetByTarget(cfg, "num")
		var cpuNum = cpuNumber
		if nil != num {
			switch reflect.TypeOf(num).String() {
			case "float64":
				cpuNum = int(reflect.ValueOf(num).Float())
				break
			case "int":
				cpuNum = int(reflect.ValueOf(num).Int())
				break
			default:
				break
			}
		}
		if 0 >= cpuNum {
			cpuNum = cpuNumber
		}
		logger.Info("main", fmt.Sprintf("multi server model, server will use %v cpu", cpuNum))
		runtime.GOMAXPROCS(cpuNum * 2) // 限制go 出去的数量
	}
}

func main() {
	logger.Info("main", "service init ...")
	logger.Info("main",fmt.Sprintf("machine is %d cpus.", cpuNumber))
	serverCfg := config.Get("server").(map[string]interface{})
	finish := make(chan bool)
	if nil == serverCfg {
		startDefault()
		return
	}
	go startUpCacheService(serverCfg["cache"].(map[string]interface{})) // 缓存线程
	go startUpDaemonService(serverCfg["daemon"].(map[string]interface{})) // 守护线程
	startUpGatewayService(serverCfg["gateway"].(map[string]interface{})) // 主线程
	<-finish
}