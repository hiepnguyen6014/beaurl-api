package configs

import (
	"log"
	"os"
	"strings"
	"time"

	"beaurl.vn/api/src/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetTrustedProxies(router *gin.Engine) {

	origins := GetOrigins()

	GIN_MODE := os.Getenv("GIN_MODE")

	if GIN_MODE == gin.ReleaseMode {
		router.SetTrustedProxies(origins)
	}

	if GIN_MODE == gin.DebugMode {
		router.SetTrustedProxies([]string{"127.0.0.1"})
	}

}

func GetOrigins() []string {
	GIN_MODE := os.Getenv("GIN_MODE")

	if GIN_MODE == gin.DebugMode {
		return []string{"http://localhost:3000"}
	} else {

		origin := os.Getenv("ORIGIN")
		origins := strings.Split(origin, ",")

		return origins
	}
}

func Cors(router *gin.Engine) {

	origins := GetOrigins()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Content-Type", "Access-Control-Allow-Methods"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

}

func Router() *gin.Engine {
	LoadEnv()

	GIN_MODE := os.Getenv("GIN_MODE")

	if GIN_MODE == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	Cors(router)
	SetTrustedProxies(router)

	routes.Routes(router)

	log.Printf("--- GIN_MODE %v ---", os.Getenv("GIN_MODE"))
	log.Printf("--- PORT %v ---", os.Getenv("PORT"))

	return router
}
