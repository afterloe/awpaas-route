package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../util"
	"../config"
)

func Info(context *gin.Context) {
	info := config.Get("info").(map[string]interface{})
	context.JSON(http.StatusOK, util.Success(info))
}
