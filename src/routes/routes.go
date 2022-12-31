package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	api := router.Group("/api")
	{
		api.GET("/ping", func(ctx *gin.Context) {
			fmt.Printf("ClientIP: %s\n", ctx.ClientIP())
			fmt.Printf("Domain: %s\n", ctx.Request.Host)
			ctx.JSON(http.StatusOK, gin.H{
				"message": "ping",
			})
		})
	}

	return router
}
