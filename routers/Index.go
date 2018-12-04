package routers

import (
	"github.com/gin-gonic/gin"
)

/**
	路由列表
 */
func Execute(route *gin.RouterGroup) {
	route.GET("/whiteList", WhiteList)
	route.PUT("/whiteList", WhiteListAppend)
	route.DELETE("/whiteList", WhiteListDel)
	route.GET("/serviceMap", ServiceMap)
	route.POST("/serviceMap", ServiceMapAppend)
	route.PUT("/serviceMap", ServiceMapModify)
	route.DELETE("/serviceMap", ServiceMapDel)
	route.Any("/tips/:code", Tips)
}
