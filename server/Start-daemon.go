package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../integrate/logger"
	"../integrate/notSupper"
	"../util"
	"../config"
	"../routers"
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
	engine.Use(gin.Recovery())
	engine.Use(logger.Logger())
	engine.Use(notSupper.HasError())
	engine.NoRoute(notSupper.NotFound(&notFoundStr))
	engine.NoMethod(notSupper.NotSupper(&notSupperStr))
	infoEntryPoint(engine)
	routers.Execute(engine.Group("/v1"))
	server := &http.Server{
		Addr: *addr,
		Handler: engine,
		MaxHeaderBytes: 1 << 20,
	}

	error := server.ListenAndServe()
	if nil != error {
		logger.Error("server can't to run")
		logger.Error(error.Error())
		os.Exit(102)
	}
	logger.Info("daemon service is ready ...")
}

func infoEntryPoint(c *gin.Engine) {
	info := config.Get("info").(map[string]interface{})
	c.GET("/info", func(context *gin.Context) {
		context.JSON(http.StatusOK, util.Success(info))
	})
}