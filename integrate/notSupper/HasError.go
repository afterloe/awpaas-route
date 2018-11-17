package notSupper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../../util"
	"../../exceptions"
)

func HasError() func(context *gin.Context) {
	return func (c *gin.Context) {
		c.Next()
		err := c.Errors
		if nil != err {
			 e := err.Last().Err
			if val, ok := e.(*exceptions.Error); ok {
				c.JSON(http.StatusOK, util.Fail(val.Code, val.Msg))
				return
			}
			c.JSON(http.StatusOK, util.Fail(http.StatusInternalServerError, c.Errors.String()))
		}
	}
}