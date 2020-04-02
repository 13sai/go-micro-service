package web

import (
	"net/http"

	"github.com/spf13/viper"
    "github.com/gin-gonic/gin"
)

func StartHttp() *gin.Engine {
	gin.SetMode(viper.GetString("runmode"))
	g := gin.New()
	g = Load(g)
	return g
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
