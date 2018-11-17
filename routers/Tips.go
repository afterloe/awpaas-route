package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../util"
	"strconv"
)

func Tips(c *gin.Context) {
	codeStr, content := c.Param("code"), c.Query("content")
	code, err := strconv.Atoi(codeStr)
	if nil != err {
		code = 200
	}
	c.JSON(http.StatusOK, util.Build(nil, code, content))
}