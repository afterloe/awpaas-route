package routers

import (
	"github.com/gin-gonic/gin"
	"../util"
	"../service/cache"
	"net/http"
)

func ServiceMap(context *gin.Context) {
	addMap := cache.GetAddMapFromDisk()
	context.JSON(http.StatusOK, util.Success(addMap))
}

func ServiceMapAppend(context *gin.Context) {
	serviceName := context.PostForm("serviceName")
	serviceAddr := context.PostForm("serviceAddr")
	if "" == serviceName || "" == serviceAddr {
		context.JSON(http.StatusBadRequest, util.Fail(400, "lack parameter -> serviceName serviceAddr"))
		return
	}
	err := cache.AppendAddrMap(serviceName, serviceAddr)
	if nil != err {
		context.JSON(http.StatusInternalServerError, util.Error(err))
		// context.JSON(http.StatusInternalServerError, util.Fail(400, "parameter is error"))
		return
	}
	context.JSON(http.StatusOK, util.Success("append success"))
}