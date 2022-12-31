package configs

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetOrigins() []string {
	origin := os.Getenv("ORIGIN")
	origins := strings.Split(origin, ",")

	return origins
}

func Cors(router *gin.Engine) {

	origins := GetOrigins()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

}

func SetTrustedProxies(router *gin.Engine) {

	origins := GetOrigins()

	GIN_MODE := os.Getenv("GIN_MODE")

	if GIN_MODE == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
		router.SetTrustedProxies(origins)
	}

	if GIN_MODE == gin.DebugMode {
		router.SetTrustedProxies([]string{"127.0.0.1"})
	}

}

func CommonConfigs(router *gin.Engine) {
	Cors(router)
	SetTrustedProxies(router)
}
