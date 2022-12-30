package routers

import (
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
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ping",
		})
	})

	return router
}
