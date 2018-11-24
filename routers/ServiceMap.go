package routers

import (
	"github.com/gin-gonic/gin"
	"../util"
	"../service/cache"
	"net/http"
	"../exceptions"
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
	if err {
		err = err.(exceptions.Error)
		context.JSON(http.StatusInternalServerError, util.Fail(err.Code, err.Msg))
		return
	}
	context.JSON(http.StatusOK, util.Success("append success"))
}