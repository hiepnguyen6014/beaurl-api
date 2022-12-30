package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		fmt.Printf("ClientIP: %s\n", ctx.ClientIP())
		fmt.Printf("Domain: %s\n", ctx.Request.Host)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	return router
}
