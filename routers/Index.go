package routers

import (
	"github.com/gin-gonic/gin"
)

/**
	路由列表
 */
func Execute(route *gin.RouterGroup) {
	route.GET("/", Home)
	route.GET("/whiteList", WhiteList)
	route.PUT("/whiteList", WhiteListAppend)
	route.DELETE("/whiteList", WhiteListDel)
	route.GET("/serviceMap", ServiceMap)
	// route.PUT("/serviceMap", ServiceMapAppend)
	// route.DELETE("/whiteList", ServiceMapDel)
	route.Any("/tips/:code", Tips)
}
