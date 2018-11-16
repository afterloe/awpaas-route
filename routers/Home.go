package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../util"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, util.Success("data is ready"))
}