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
func WhiteList(content *gin.Context) {
	whiteList := cache.GetWhiteListFromDisk()
	content.JSON(http.StatusOK, util.Success(whiteList))
}

/**
	添加一条进入白名单
*/
func WhiteListAppend(content *gin.Context) {
	item := content.PostForm("item")
	if "" == item {
		content.JSON(http.StatusBadRequest, util.Fail(400, "lack parameter -> item"))
		return
	}
	flag := cache.AppednItem(item)
	if flag {
		content.JSON(http.StatusOK, util.Success("append success"))
		return
	}
	content.JSON(http.StatusInternalServerError, util.Fail(500, "item has been added."))
}