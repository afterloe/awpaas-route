package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../integrate/logger"
	"../integrate/notSupper"
	"../routers"
	"../service/cache"
	"os"
)

var notFoundStr, notSupperStr string

func init() {
	notFoundStr = "route is not defined."
	notSupperStr = "method is not supper"
}

/**
	启动守护进程

*/
func StartUpDaemonService(addr *string, serverCfg map[string]interface{}) {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	initDaemonService(engine, serverCfg)
	server := &http.Server{
		Addr: *addr,
		Handler: engine,
		MaxHeaderBytes: 1 << 20,
	}
	e := server.ListenAndServe()
	if nil != e {
		logger.Error("server can't to run")
		logger.Error(e.Error())
		os.Exit(102)
	}
}

func initDaemonService(engine *gin.Engine, serverCfg map[string]interface{}) {
	engine.Use(gin.Recovery())
	engine.Use(logger.Logger())
	engine.Use(notSupper.HasError())
	engine.NoRoute(notSupper.NotFound(&notFoundStr))
	engine.NoMethod(notSupper.NotSupper(&notSupperStr))
	infoEntryPoint(engine)
	routers.Execute(engine.Group("/v1"))
	cache.FlushWhiteListCache(serverCfg["whiteList"].([]string))
	logger.Info("daemon service is ready ...")
}

func infoEntryPoint(c *gin.Engine) {
	c.GET("/info", routers.Info)
}