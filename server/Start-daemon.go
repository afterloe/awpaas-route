package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../integrate/logger"
	"../integrate/notSupper"
	"../routers"
	"../config"
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
func StartUpDaemonService(addr *string, cfg interface{}) {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	initDaemonService(engine, cfg)
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

func initDaemonService(engine *gin.Engine, cfg interface{}) {
	engine.Use(gin.Recovery())
	engine.Use(logger.GinLogger())
	engine.Use(notSupper.HasError())
	engine.NoRoute(notSupper.NotFound(&notFoundStr))
	engine.NoMethod(notSupper.NotSupper(&notSupperStr))
	infoEntryPoint(engine)
	routers.Execute(engine.Group("/v1"))
	cache.LoadCache(config.GetByTarget(cfg,"whiteList").([]interface{}))
	logger.Info("daemon service is ready ...")
}

func infoEntryPoint(c *gin.Engine) {
	c.GET("/info", routers.Info)
}