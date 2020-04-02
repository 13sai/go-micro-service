package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.GET("/health", func(c *gin.Context){
		SendResponse(c, 0, "success", nil)
	})

	return g
}