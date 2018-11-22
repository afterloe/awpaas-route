package routers

import (
	"github.com/gin-gonic/gin"
	"../util"
	"../service/cache"
	"net/http"
)

func WhiteList(c *gin.Context) {
	whiteList := cache.GetWhiteListFromDisk()
	c.JSON(http.StatusOK, util.Success(whiteList))
}