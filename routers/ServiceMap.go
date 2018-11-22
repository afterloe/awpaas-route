package routers

import (
	"github.com/gin-gonic/gin"
	"../util"
	"../service/cache"
	"net/http"
)

func ServiceMap(c *gin.Context) {
	addMap := cache.GetAddMapFromDisk()
	c.JSON(http.StatusOK, util.Success(addMap))
}
