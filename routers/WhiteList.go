package routers

import (
	"github.com/gin-gonic/gin"
	"../util"
	"../service/cache"
	"net/http"
)

/**
	获取白名单列表
*/
func WhiteList(context *gin.Context) {
	whiteList := cache.GetWhiteListFromDisk()
	context.JSON(http.StatusOK, util.Success(whiteList))
}

/**
	添加一条进入白名单
*/
func WhiteListAppend(context *gin.Context) {
	item := context.PostForm("item")
	if "" == item {
		context.JSON(http.StatusBadRequest, util.Fail(400, "lack parameter -> item"))
		return
	}
	flag := cache.AppednItem(item)
	if flag {
		context.JSON(http.StatusOK, util.Success("append success"))
		return
	}
	context.JSON(http.StatusInternalServerError, util.Fail(500, "item has been added."))
}

func WhiteListDel(context *gin.Context) {
	item := context.PostForm("item")
	if "" == item {
		context.JSON(http.StatusBadRequest, util.Fail(400, "lack parameter -> item"))
		return
	}
	flag := cache.RemoveItem(item)
	if flag {
		context.JSON(http.StatusOK, util.Success("remove success"))
		return
	}
	context.JSON(http.StatusInternalServerError, util.Fail(500, "item has been remove."))
}	