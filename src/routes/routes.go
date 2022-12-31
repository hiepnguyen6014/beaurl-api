package routes

import (
	"beaurl.vn/api/src/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	api := router.Group("/shorten")
	{
		api.GET("/", controllers.ShortenLink)
	}

}
