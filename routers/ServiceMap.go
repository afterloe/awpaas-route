package routers

import (
	"github.com/gin-gonic/gin"
	"../util"
	"../service/cache"
	"net/http"
	"fmt"
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
		return
	}
	context.JSON(http.StatusOK, util.Success("append success"))
}

func ServiceMapModify(context *gin.Context) {
	serviceName := context.PostForm("serviceName")
	serviceAddr := context.PostForm("serviceAddr")
	if "" == serviceName || "" == serviceAddr {
		context.JSON(http.StatusBadRequest, util.Fail(400, "lack parameter -> serviceName serviceAddr"))
		return
	}
	err := cache.ModifyAddrMap(serviceName, serviceAddr)
	if nil != err {
		context.JSON(http.StatusInternalServerError, util.Error(err))
		return
	}
	context.JSON(http.StatusOK, util.Success("modify success"))
}

func ServiceMapDel(context *gin.Context) {
	serviceName := context.DefaultQuery("serviceName", "unknow")
	if "unknow" == serviceName {
		context.JSON(http.StatusBadRequest, util.Fail(400, "lack parameter -> serviceName"))
		return
	}
	err := cache.DelAddrMap(serviceName)
	if nil != err {
		context.JSON(http.StatusInternalServerError, util.Error(err))
		return
	}
	context.JSON(http.StatusOK, util.Success(fmt.Sprintf("delete %s success", serviceName)))
}